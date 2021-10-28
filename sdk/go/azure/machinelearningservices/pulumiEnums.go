


package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationSharingPolicy string

const (
	ApplicationSharingPolicyPersonal = ApplicationSharingPolicy("Personal")
	ApplicationSharingPolicyShared   = ApplicationSharingPolicy("Shared")
)

func (ApplicationSharingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationSharingPolicy)(nil)).Elem()
}

func (e ApplicationSharingPolicy) ToApplicationSharingPolicyOutput() ApplicationSharingPolicyOutput {
	return pulumi.ToOutput(e).(ApplicationSharingPolicyOutput)
}

func (e ApplicationSharingPolicy) ToApplicationSharingPolicyOutputWithContext(ctx context.Context) ApplicationSharingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApplicationSharingPolicyOutput)
}

func (e ApplicationSharingPolicy) ToApplicationSharingPolicyPtrOutput() ApplicationSharingPolicyPtrOutput {
	return e.ToApplicationSharingPolicyPtrOutputWithContext(context.Background())
}

func (e ApplicationSharingPolicy) ToApplicationSharingPolicyPtrOutputWithContext(ctx context.Context) ApplicationSharingPolicyPtrOutput {
	return ApplicationSharingPolicy(e).ToApplicationSharingPolicyOutputWithContext(ctx).ToApplicationSharingPolicyPtrOutputWithContext(ctx)
}

func (e ApplicationSharingPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationSharingPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationSharingPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApplicationSharingPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ApplicationSharingPolicyOutput struct{ *pulumi.OutputState }

func (ApplicationSharingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationSharingPolicy)(nil)).Elem()
}

func (o ApplicationSharingPolicyOutput) ToApplicationSharingPolicyOutput() ApplicationSharingPolicyOutput {
	return o
}

func (o ApplicationSharingPolicyOutput) ToApplicationSharingPolicyOutputWithContext(ctx context.Context) ApplicationSharingPolicyOutput {
	return o
}

func (o ApplicationSharingPolicyOutput) ToApplicationSharingPolicyPtrOutput() ApplicationSharingPolicyPtrOutput {
	return o.ToApplicationSharingPolicyPtrOutputWithContext(context.Background())
}

func (o ApplicationSharingPolicyOutput) ToApplicationSharingPolicyPtrOutputWithContext(ctx context.Context) ApplicationSharingPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationSharingPolicy) *ApplicationSharingPolicy {
		return &v
	}).(ApplicationSharingPolicyPtrOutput)
}

func (o ApplicationSharingPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApplicationSharingPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationSharingPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApplicationSharingPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationSharingPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationSharingPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationSharingPolicyPtrOutput struct{ *pulumi.OutputState }

func (ApplicationSharingPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationSharingPolicy)(nil)).Elem()
}

func (o ApplicationSharingPolicyPtrOutput) ToApplicationSharingPolicyPtrOutput() ApplicationSharingPolicyPtrOutput {
	return o
}

func (o ApplicationSharingPolicyPtrOutput) ToApplicationSharingPolicyPtrOutputWithContext(ctx context.Context) ApplicationSharingPolicyPtrOutput {
	return o
}

func (o ApplicationSharingPolicyPtrOutput) Elem() ApplicationSharingPolicyOutput {
	return o.ApplyT(func(v *ApplicationSharingPolicy) ApplicationSharingPolicy {
		if v != nil {
			return *v
		}
		var ret ApplicationSharingPolicy
		return ret
	}).(ApplicationSharingPolicyOutput)
}

func (o ApplicationSharingPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationSharingPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApplicationSharingPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ApplicationSharingPolicyInput interface {
	pulumi.Input

	ToApplicationSharingPolicyOutput() ApplicationSharingPolicyOutput
	ToApplicationSharingPolicyOutputWithContext(context.Context) ApplicationSharingPolicyOutput
}

var applicationSharingPolicyPtrType = reflect.TypeOf((**ApplicationSharingPolicy)(nil)).Elem()

type ApplicationSharingPolicyPtrInput interface {
	pulumi.Input

	ToApplicationSharingPolicyPtrOutput() ApplicationSharingPolicyPtrOutput
	ToApplicationSharingPolicyPtrOutputWithContext(context.Context) ApplicationSharingPolicyPtrOutput
}

type applicationSharingPolicyPtr string

func ApplicationSharingPolicyPtr(v string) ApplicationSharingPolicyPtrInput {
	return (*applicationSharingPolicyPtr)(&v)
}

func (*applicationSharingPolicyPtr) ElementType() reflect.Type {
	return applicationSharingPolicyPtrType
}

func (in *applicationSharingPolicyPtr) ToApplicationSharingPolicyPtrOutput() ApplicationSharingPolicyPtrOutput {
	return pulumi.ToOutput(in).(ApplicationSharingPolicyPtrOutput)
}

func (in *applicationSharingPolicyPtr) ToApplicationSharingPolicyPtrOutputWithContext(ctx context.Context) ApplicationSharingPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApplicationSharingPolicyPtrOutput)
}

type BatchLoggingLevel string

const (
	BatchLoggingLevelInfo    = BatchLoggingLevel("Info")
	BatchLoggingLevelWarning = BatchLoggingLevel("Warning")
	BatchLoggingLevelDebug   = BatchLoggingLevel("Debug")
)

func (BatchLoggingLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchLoggingLevel)(nil)).Elem()
}

func (e BatchLoggingLevel) ToBatchLoggingLevelOutput() BatchLoggingLevelOutput {
	return pulumi.ToOutput(e).(BatchLoggingLevelOutput)
}

func (e BatchLoggingLevel) ToBatchLoggingLevelOutputWithContext(ctx context.Context) BatchLoggingLevelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BatchLoggingLevelOutput)
}

func (e BatchLoggingLevel) ToBatchLoggingLevelPtrOutput() BatchLoggingLevelPtrOutput {
	return e.ToBatchLoggingLevelPtrOutputWithContext(context.Background())
}

func (e BatchLoggingLevel) ToBatchLoggingLevelPtrOutputWithContext(ctx context.Context) BatchLoggingLevelPtrOutput {
	return BatchLoggingLevel(e).ToBatchLoggingLevelOutputWithContext(ctx).ToBatchLoggingLevelPtrOutputWithContext(ctx)
}

func (e BatchLoggingLevel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BatchLoggingLevel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BatchLoggingLevel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BatchLoggingLevel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BatchLoggingLevelOutput struct{ *pulumi.OutputState }

func (BatchLoggingLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchLoggingLevel)(nil)).Elem()
}

func (o BatchLoggingLevelOutput) ToBatchLoggingLevelOutput() BatchLoggingLevelOutput {
	return o
}

func (o BatchLoggingLevelOutput) ToBatchLoggingLevelOutputWithContext(ctx context.Context) BatchLoggingLevelOutput {
	return o
}

func (o BatchLoggingLevelOutput) ToBatchLoggingLevelPtrOutput() BatchLoggingLevelPtrOutput {
	return o.ToBatchLoggingLevelPtrOutputWithContext(context.Background())
}

func (o BatchLoggingLevelOutput) ToBatchLoggingLevelPtrOutputWithContext(ctx context.Context) BatchLoggingLevelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BatchLoggingLevel) *BatchLoggingLevel {
		return &v
	}).(BatchLoggingLevelPtrOutput)
}

func (o BatchLoggingLevelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BatchLoggingLevelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BatchLoggingLevel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BatchLoggingLevelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BatchLoggingLevelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BatchLoggingLevel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BatchLoggingLevelPtrOutput struct{ *pulumi.OutputState }

func (BatchLoggingLevelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BatchLoggingLevel)(nil)).Elem()
}

func (o BatchLoggingLevelPtrOutput) ToBatchLoggingLevelPtrOutput() BatchLoggingLevelPtrOutput {
	return o
}

func (o BatchLoggingLevelPtrOutput) ToBatchLoggingLevelPtrOutputWithContext(ctx context.Context) BatchLoggingLevelPtrOutput {
	return o
}

func (o BatchLoggingLevelPtrOutput) Elem() BatchLoggingLevelOutput {
	return o.ApplyT(func(v *BatchLoggingLevel) BatchLoggingLevel {
		if v != nil {
			return *v
		}
		var ret BatchLoggingLevel
		return ret
	}).(BatchLoggingLevelOutput)
}

func (o BatchLoggingLevelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BatchLoggingLevelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BatchLoggingLevel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BatchLoggingLevelInput interface {
	pulumi.Input

	ToBatchLoggingLevelOutput() BatchLoggingLevelOutput
	ToBatchLoggingLevelOutputWithContext(context.Context) BatchLoggingLevelOutput
}

var batchLoggingLevelPtrType = reflect.TypeOf((**BatchLoggingLevel)(nil)).Elem()

type BatchLoggingLevelPtrInput interface {
	pulumi.Input

	ToBatchLoggingLevelPtrOutput() BatchLoggingLevelPtrOutput
	ToBatchLoggingLevelPtrOutputWithContext(context.Context) BatchLoggingLevelPtrOutput
}

type batchLoggingLevelPtr string

func BatchLoggingLevelPtr(v string) BatchLoggingLevelPtrInput {
	return (*batchLoggingLevelPtr)(&v)
}

func (*batchLoggingLevelPtr) ElementType() reflect.Type {
	return batchLoggingLevelPtrType
}

func (in *batchLoggingLevelPtr) ToBatchLoggingLevelPtrOutput() BatchLoggingLevelPtrOutput {
	return pulumi.ToOutput(in).(BatchLoggingLevelPtrOutput)
}

func (in *batchLoggingLevelPtr) ToBatchLoggingLevelPtrOutputWithContext(ctx context.Context) BatchLoggingLevelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BatchLoggingLevelPtrOutput)
}

type BatchOutputAction string

const (
	BatchOutputActionSummaryOnly = BatchOutputAction("SummaryOnly")
	BatchOutputActionAppendRow   = BatchOutputAction("AppendRow")
)

func (BatchOutputAction) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchOutputAction)(nil)).Elem()
}

func (e BatchOutputAction) ToBatchOutputActionOutput() BatchOutputActionOutput {
	return pulumi.ToOutput(e).(BatchOutputActionOutput)
}

func (e BatchOutputAction) ToBatchOutputActionOutputWithContext(ctx context.Context) BatchOutputActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BatchOutputActionOutput)
}

func (e BatchOutputAction) ToBatchOutputActionPtrOutput() BatchOutputActionPtrOutput {
	return e.ToBatchOutputActionPtrOutputWithContext(context.Background())
}

func (e BatchOutputAction) ToBatchOutputActionPtrOutputWithContext(ctx context.Context) BatchOutputActionPtrOutput {
	return BatchOutputAction(e).ToBatchOutputActionOutputWithContext(ctx).ToBatchOutputActionPtrOutputWithContext(ctx)
}

func (e BatchOutputAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BatchOutputAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BatchOutputAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BatchOutputAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BatchOutputActionOutput struct{ *pulumi.OutputState }

func (BatchOutputActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchOutputAction)(nil)).Elem()
}

func (o BatchOutputActionOutput) ToBatchOutputActionOutput() BatchOutputActionOutput {
	return o
}

func (o BatchOutputActionOutput) ToBatchOutputActionOutputWithContext(ctx context.Context) BatchOutputActionOutput {
	return o
}

func (o BatchOutputActionOutput) ToBatchOutputActionPtrOutput() BatchOutputActionPtrOutput {
	return o.ToBatchOutputActionPtrOutputWithContext(context.Background())
}

func (o BatchOutputActionOutput) ToBatchOutputActionPtrOutputWithContext(ctx context.Context) BatchOutputActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BatchOutputAction) *BatchOutputAction {
		return &v
	}).(BatchOutputActionPtrOutput)
}

func (o BatchOutputActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BatchOutputActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BatchOutputAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BatchOutputActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BatchOutputActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BatchOutputAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BatchOutputActionPtrOutput struct{ *pulumi.OutputState }

func (BatchOutputActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BatchOutputAction)(nil)).Elem()
}

func (o BatchOutputActionPtrOutput) ToBatchOutputActionPtrOutput() BatchOutputActionPtrOutput {
	return o
}

func (o BatchOutputActionPtrOutput) ToBatchOutputActionPtrOutputWithContext(ctx context.Context) BatchOutputActionPtrOutput {
	return o
}

func (o BatchOutputActionPtrOutput) Elem() BatchOutputActionOutput {
	return o.ApplyT(func(v *BatchOutputAction) BatchOutputAction {
		if v != nil {
			return *v
		}
		var ret BatchOutputAction
		return ret
	}).(BatchOutputActionOutput)
}

func (o BatchOutputActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BatchOutputActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BatchOutputAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BatchOutputActionInput interface {
	pulumi.Input

	ToBatchOutputActionOutput() BatchOutputActionOutput
	ToBatchOutputActionOutputWithContext(context.Context) BatchOutputActionOutput
}

var batchOutputActionPtrType = reflect.TypeOf((**BatchOutputAction)(nil)).Elem()

type BatchOutputActionPtrInput interface {
	pulumi.Input

	ToBatchOutputActionPtrOutput() BatchOutputActionPtrOutput
	ToBatchOutputActionPtrOutputWithContext(context.Context) BatchOutputActionPtrOutput
}

type batchOutputActionPtr string

func BatchOutputActionPtr(v string) BatchOutputActionPtrInput {
	return (*batchOutputActionPtr)(&v)
}

func (*batchOutputActionPtr) ElementType() reflect.Type {
	return batchOutputActionPtrType
}

func (in *batchOutputActionPtr) ToBatchOutputActionPtrOutput() BatchOutputActionPtrOutput {
	return pulumi.ToOutput(in).(BatchOutputActionPtrOutput)
}

func (in *batchOutputActionPtr) ToBatchOutputActionPtrOutputWithContext(ctx context.Context) BatchOutputActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BatchOutputActionPtrOutput)
}

type ClusterPurpose string

const (
	ClusterPurposeFastProd  = ClusterPurpose("FastProd")
	ClusterPurposeDenseProd = ClusterPurpose("DenseProd")
	ClusterPurposeDevTest   = ClusterPurpose("DevTest")
)

func (ClusterPurpose) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterPurpose)(nil)).Elem()
}

func (e ClusterPurpose) ToClusterPurposeOutput() ClusterPurposeOutput {
	return pulumi.ToOutput(e).(ClusterPurposeOutput)
}

func (e ClusterPurpose) ToClusterPurposeOutputWithContext(ctx context.Context) ClusterPurposeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ClusterPurposeOutput)
}

func (e ClusterPurpose) ToClusterPurposePtrOutput() ClusterPurposePtrOutput {
	return e.ToClusterPurposePtrOutputWithContext(context.Background())
}

func (e ClusterPurpose) ToClusterPurposePtrOutputWithContext(ctx context.Context) ClusterPurposePtrOutput {
	return ClusterPurpose(e).ToClusterPurposeOutputWithContext(ctx).ToClusterPurposePtrOutputWithContext(ctx)
}

func (e ClusterPurpose) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterPurpose) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterPurpose) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ClusterPurpose) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ClusterPurposeOutput struct{ *pulumi.OutputState }

func (ClusterPurposeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterPurpose)(nil)).Elem()
}

func (o ClusterPurposeOutput) ToClusterPurposeOutput() ClusterPurposeOutput {
	return o
}

func (o ClusterPurposeOutput) ToClusterPurposeOutputWithContext(ctx context.Context) ClusterPurposeOutput {
	return o
}

func (o ClusterPurposeOutput) ToClusterPurposePtrOutput() ClusterPurposePtrOutput {
	return o.ToClusterPurposePtrOutputWithContext(context.Background())
}

func (o ClusterPurposeOutput) ToClusterPurposePtrOutputWithContext(ctx context.Context) ClusterPurposePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterPurpose) *ClusterPurpose {
		return &v
	}).(ClusterPurposePtrOutput)
}

func (o ClusterPurposeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ClusterPurposeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterPurpose) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ClusterPurposeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterPurposeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterPurpose) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ClusterPurposePtrOutput struct{ *pulumi.OutputState }

func (ClusterPurposePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterPurpose)(nil)).Elem()
}

func (o ClusterPurposePtrOutput) ToClusterPurposePtrOutput() ClusterPurposePtrOutput {
	return o
}

func (o ClusterPurposePtrOutput) ToClusterPurposePtrOutputWithContext(ctx context.Context) ClusterPurposePtrOutput {
	return o
}

func (o ClusterPurposePtrOutput) Elem() ClusterPurposeOutput {
	return o.ApplyT(func(v *ClusterPurpose) ClusterPurpose {
		if v != nil {
			return *v
		}
		var ret ClusterPurpose
		return ret
	}).(ClusterPurposeOutput)
}

func (o ClusterPurposePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterPurposePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ClusterPurpose) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ClusterPurposeInput interface {
	pulumi.Input

	ToClusterPurposeOutput() ClusterPurposeOutput
	ToClusterPurposeOutputWithContext(context.Context) ClusterPurposeOutput
}

var clusterPurposePtrType = reflect.TypeOf((**ClusterPurpose)(nil)).Elem()

type ClusterPurposePtrInput interface {
	pulumi.Input

	ToClusterPurposePtrOutput() ClusterPurposePtrOutput
	ToClusterPurposePtrOutputWithContext(context.Context) ClusterPurposePtrOutput
}

type clusterPurposePtr string

func ClusterPurposePtr(v string) ClusterPurposePtrInput {
	return (*clusterPurposePtr)(&v)
}

func (*clusterPurposePtr) ElementType() reflect.Type {
	return clusterPurposePtrType
}

func (in *clusterPurposePtr) ToClusterPurposePtrOutput() ClusterPurposePtrOutput {
	return pulumi.ToOutput(in).(ClusterPurposePtrOutput)
}

func (in *clusterPurposePtr) ToClusterPurposePtrOutputWithContext(ctx context.Context) ClusterPurposePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ClusterPurposePtrOutput)
}

type ComputeEnvironmentType string

const (
	ComputeEnvironmentTypeACI = ComputeEnvironmentType("ACI")
	ComputeEnvironmentTypeAKS = ComputeEnvironmentType("AKS")
)

func (ComputeEnvironmentType) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeEnvironmentType)(nil)).Elem()
}

func (e ComputeEnvironmentType) ToComputeEnvironmentTypeOutput() ComputeEnvironmentTypeOutput {
	return pulumi.ToOutput(e).(ComputeEnvironmentTypeOutput)
}

func (e ComputeEnvironmentType) ToComputeEnvironmentTypeOutputWithContext(ctx context.Context) ComputeEnvironmentTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ComputeEnvironmentTypeOutput)
}

func (e ComputeEnvironmentType) ToComputeEnvironmentTypePtrOutput() ComputeEnvironmentTypePtrOutput {
	return e.ToComputeEnvironmentTypePtrOutputWithContext(context.Background())
}

func (e ComputeEnvironmentType) ToComputeEnvironmentTypePtrOutputWithContext(ctx context.Context) ComputeEnvironmentTypePtrOutput {
	return ComputeEnvironmentType(e).ToComputeEnvironmentTypeOutputWithContext(ctx).ToComputeEnvironmentTypePtrOutputWithContext(ctx)
}

func (e ComputeEnvironmentType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComputeEnvironmentType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComputeEnvironmentType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ComputeEnvironmentType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ComputeEnvironmentTypeOutput struct{ *pulumi.OutputState }

func (ComputeEnvironmentTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeEnvironmentType)(nil)).Elem()
}

func (o ComputeEnvironmentTypeOutput) ToComputeEnvironmentTypeOutput() ComputeEnvironmentTypeOutput {
	return o
}

func (o ComputeEnvironmentTypeOutput) ToComputeEnvironmentTypeOutputWithContext(ctx context.Context) ComputeEnvironmentTypeOutput {
	return o
}

func (o ComputeEnvironmentTypeOutput) ToComputeEnvironmentTypePtrOutput() ComputeEnvironmentTypePtrOutput {
	return o.ToComputeEnvironmentTypePtrOutputWithContext(context.Background())
}

func (o ComputeEnvironmentTypeOutput) ToComputeEnvironmentTypePtrOutputWithContext(ctx context.Context) ComputeEnvironmentTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeEnvironmentType) *ComputeEnvironmentType {
		return &v
	}).(ComputeEnvironmentTypePtrOutput)
}

func (o ComputeEnvironmentTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ComputeEnvironmentTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComputeEnvironmentType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ComputeEnvironmentTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComputeEnvironmentTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComputeEnvironmentType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ComputeEnvironmentTypePtrOutput struct{ *pulumi.OutputState }

func (ComputeEnvironmentTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeEnvironmentType)(nil)).Elem()
}

func (o ComputeEnvironmentTypePtrOutput) ToComputeEnvironmentTypePtrOutput() ComputeEnvironmentTypePtrOutput {
	return o
}

func (o ComputeEnvironmentTypePtrOutput) ToComputeEnvironmentTypePtrOutputWithContext(ctx context.Context) ComputeEnvironmentTypePtrOutput {
	return o
}

func (o ComputeEnvironmentTypePtrOutput) Elem() ComputeEnvironmentTypeOutput {
	return o.ApplyT(func(v *ComputeEnvironmentType) ComputeEnvironmentType {
		if v != nil {
			return *v
		}
		var ret ComputeEnvironmentType
		return ret
	}).(ComputeEnvironmentTypeOutput)
}

func (o ComputeEnvironmentTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComputeEnvironmentTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ComputeEnvironmentType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ComputeEnvironmentTypeInput interface {
	pulumi.Input

	ToComputeEnvironmentTypeOutput() ComputeEnvironmentTypeOutput
	ToComputeEnvironmentTypeOutputWithContext(context.Context) ComputeEnvironmentTypeOutput
}

var computeEnvironmentTypePtrType = reflect.TypeOf((**ComputeEnvironmentType)(nil)).Elem()

type ComputeEnvironmentTypePtrInput interface {
	pulumi.Input

	ToComputeEnvironmentTypePtrOutput() ComputeEnvironmentTypePtrOutput
	ToComputeEnvironmentTypePtrOutputWithContext(context.Context) ComputeEnvironmentTypePtrOutput
}

type computeEnvironmentTypePtr string

func ComputeEnvironmentTypePtr(v string) ComputeEnvironmentTypePtrInput {
	return (*computeEnvironmentTypePtr)(&v)
}

func (*computeEnvironmentTypePtr) ElementType() reflect.Type {
	return computeEnvironmentTypePtrType
}

func (in *computeEnvironmentTypePtr) ToComputeEnvironmentTypePtrOutput() ComputeEnvironmentTypePtrOutput {
	return pulumi.ToOutput(in).(ComputeEnvironmentTypePtrOutput)
}

func (in *computeEnvironmentTypePtr) ToComputeEnvironmentTypePtrOutputWithContext(ctx context.Context) ComputeEnvironmentTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ComputeEnvironmentTypePtrOutput)
}

type ComputeInstanceAuthorizationType string

const (
	ComputeInstanceAuthorizationTypePersonal = ComputeInstanceAuthorizationType("personal")
)

func (ComputeInstanceAuthorizationType) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeInstanceAuthorizationType)(nil)).Elem()
}

func (e ComputeInstanceAuthorizationType) ToComputeInstanceAuthorizationTypeOutput() ComputeInstanceAuthorizationTypeOutput {
	return pulumi.ToOutput(e).(ComputeInstanceAuthorizationTypeOutput)
}

func (e ComputeInstanceAuthorizationType) ToComputeInstanceAuthorizationTypeOutputWithContext(ctx context.Context) ComputeInstanceAuthorizationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ComputeInstanceAuthorizationTypeOutput)
}

func (e ComputeInstanceAuthorizationType) ToComputeInstanceAuthorizationTypePtrOutput() ComputeInstanceAuthorizationTypePtrOutput {
	return e.ToComputeInstanceAuthorizationTypePtrOutputWithContext(context.Background())
}

func (e ComputeInstanceAuthorizationType) ToComputeInstanceAuthorizationTypePtrOutputWithContext(ctx context.Context) ComputeInstanceAuthorizationTypePtrOutput {
	return ComputeInstanceAuthorizationType(e).ToComputeInstanceAuthorizationTypeOutputWithContext(ctx).ToComputeInstanceAuthorizationTypePtrOutputWithContext(ctx)
}

func (e ComputeInstanceAuthorizationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComputeInstanceAuthorizationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComputeInstanceAuthorizationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ComputeInstanceAuthorizationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ComputeInstanceAuthorizationTypeOutput struct{ *pulumi.OutputState }

func (ComputeInstanceAuthorizationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeInstanceAuthorizationType)(nil)).Elem()
}

func (o ComputeInstanceAuthorizationTypeOutput) ToComputeInstanceAuthorizationTypeOutput() ComputeInstanceAuthorizationTypeOutput {
	return o
}

func (o ComputeInstanceAuthorizationTypeOutput) ToComputeInstanceAuthorizationTypeOutputWithContext(ctx context.Context) ComputeInstanceAuthorizationTypeOutput {
	return o
}

func (o ComputeInstanceAuthorizationTypeOutput) ToComputeInstanceAuthorizationTypePtrOutput() ComputeInstanceAuthorizationTypePtrOutput {
	return o.ToComputeInstanceAuthorizationTypePtrOutputWithContext(context.Background())
}

func (o ComputeInstanceAuthorizationTypeOutput) ToComputeInstanceAuthorizationTypePtrOutputWithContext(ctx context.Context) ComputeInstanceAuthorizationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeInstanceAuthorizationType) *ComputeInstanceAuthorizationType {
		return &v
	}).(ComputeInstanceAuthorizationTypePtrOutput)
}

func (o ComputeInstanceAuthorizationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ComputeInstanceAuthorizationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComputeInstanceAuthorizationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ComputeInstanceAuthorizationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComputeInstanceAuthorizationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComputeInstanceAuthorizationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ComputeInstanceAuthorizationTypePtrOutput struct{ *pulumi.OutputState }

func (ComputeInstanceAuthorizationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeInstanceAuthorizationType)(nil)).Elem()
}

func (o ComputeInstanceAuthorizationTypePtrOutput) ToComputeInstanceAuthorizationTypePtrOutput() ComputeInstanceAuthorizationTypePtrOutput {
	return o
}

func (o ComputeInstanceAuthorizationTypePtrOutput) ToComputeInstanceAuthorizationTypePtrOutputWithContext(ctx context.Context) ComputeInstanceAuthorizationTypePtrOutput {
	return o
}

func (o ComputeInstanceAuthorizationTypePtrOutput) Elem() ComputeInstanceAuthorizationTypeOutput {
	return o.ApplyT(func(v *ComputeInstanceAuthorizationType) ComputeInstanceAuthorizationType {
		if v != nil {
			return *v
		}
		var ret ComputeInstanceAuthorizationType
		return ret
	}).(ComputeInstanceAuthorizationTypeOutput)
}

func (o ComputeInstanceAuthorizationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComputeInstanceAuthorizationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ComputeInstanceAuthorizationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ComputeInstanceAuthorizationTypeInput interface {
	pulumi.Input

	ToComputeInstanceAuthorizationTypeOutput() ComputeInstanceAuthorizationTypeOutput
	ToComputeInstanceAuthorizationTypeOutputWithContext(context.Context) ComputeInstanceAuthorizationTypeOutput
}

var computeInstanceAuthorizationTypePtrType = reflect.TypeOf((**ComputeInstanceAuthorizationType)(nil)).Elem()

type ComputeInstanceAuthorizationTypePtrInput interface {
	pulumi.Input

	ToComputeInstanceAuthorizationTypePtrOutput() ComputeInstanceAuthorizationTypePtrOutput
	ToComputeInstanceAuthorizationTypePtrOutputWithContext(context.Context) ComputeInstanceAuthorizationTypePtrOutput
}

type computeInstanceAuthorizationTypePtr string

func ComputeInstanceAuthorizationTypePtr(v string) ComputeInstanceAuthorizationTypePtrInput {
	return (*computeInstanceAuthorizationTypePtr)(&v)
}

func (*computeInstanceAuthorizationTypePtr) ElementType() reflect.Type {
	return computeInstanceAuthorizationTypePtrType
}

func (in *computeInstanceAuthorizationTypePtr) ToComputeInstanceAuthorizationTypePtrOutput() ComputeInstanceAuthorizationTypePtrOutput {
	return pulumi.ToOutput(in).(ComputeInstanceAuthorizationTypePtrOutput)
}

func (in *computeInstanceAuthorizationTypePtr) ToComputeInstanceAuthorizationTypePtrOutputWithContext(ctx context.Context) ComputeInstanceAuthorizationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ComputeInstanceAuthorizationTypePtrOutput)
}

type ComputeType string

const (
	ComputeTypeAKS               = ComputeType("AKS")
	ComputeTypeAmlCompute        = ComputeType("AmlCompute")
	ComputeTypeComputeInstance   = ComputeType("ComputeInstance")
	ComputeTypeDataFactory       = ComputeType("DataFactory")
	ComputeTypeVirtualMachine    = ComputeType("VirtualMachine")
	ComputeTypeHDInsight         = ComputeType("HDInsight")
	ComputeTypeDatabricks        = ComputeType("Databricks")
	ComputeTypeDataLakeAnalytics = ComputeType("DataLakeAnalytics")
)

func (ComputeType) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeType)(nil)).Elem()
}

func (e ComputeType) ToComputeTypeOutput() ComputeTypeOutput {
	return pulumi.ToOutput(e).(ComputeTypeOutput)
}

func (e ComputeType) ToComputeTypeOutputWithContext(ctx context.Context) ComputeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ComputeTypeOutput)
}

func (e ComputeType) ToComputeTypePtrOutput() ComputeTypePtrOutput {
	return e.ToComputeTypePtrOutputWithContext(context.Background())
}

func (e ComputeType) ToComputeTypePtrOutputWithContext(ctx context.Context) ComputeTypePtrOutput {
	return ComputeType(e).ToComputeTypeOutputWithContext(ctx).ToComputeTypePtrOutputWithContext(ctx)
}

func (e ComputeType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComputeType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComputeType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ComputeType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ComputeTypeOutput struct{ *pulumi.OutputState }

func (ComputeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeType)(nil)).Elem()
}

func (o ComputeTypeOutput) ToComputeTypeOutput() ComputeTypeOutput {
	return o
}

func (o ComputeTypeOutput) ToComputeTypeOutputWithContext(ctx context.Context) ComputeTypeOutput {
	return o
}

func (o ComputeTypeOutput) ToComputeTypePtrOutput() ComputeTypePtrOutput {
	return o.ToComputeTypePtrOutputWithContext(context.Background())
}

func (o ComputeTypeOutput) ToComputeTypePtrOutputWithContext(ctx context.Context) ComputeTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeType) *ComputeType {
		return &v
	}).(ComputeTypePtrOutput)
}

func (o ComputeTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ComputeTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComputeType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ComputeTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComputeTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComputeType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ComputeTypePtrOutput struct{ *pulumi.OutputState }

func (ComputeTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeType)(nil)).Elem()
}

func (o ComputeTypePtrOutput) ToComputeTypePtrOutput() ComputeTypePtrOutput {
	return o
}

func (o ComputeTypePtrOutput) ToComputeTypePtrOutputWithContext(ctx context.Context) ComputeTypePtrOutput {
	return o
}

func (o ComputeTypePtrOutput) Elem() ComputeTypeOutput {
	return o.ApplyT(func(v *ComputeType) ComputeType {
		if v != nil {
			return *v
		}
		var ret ComputeType
		return ret
	}).(ComputeTypeOutput)
}

func (o ComputeTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComputeTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ComputeType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ComputeTypeInput interface {
	pulumi.Input

	ToComputeTypeOutput() ComputeTypeOutput
	ToComputeTypeOutputWithContext(context.Context) ComputeTypeOutput
}

var computeTypePtrType = reflect.TypeOf((**ComputeType)(nil)).Elem()

type ComputeTypePtrInput interface {
	pulumi.Input

	ToComputeTypePtrOutput() ComputeTypePtrOutput
	ToComputeTypePtrOutputWithContext(context.Context) ComputeTypePtrOutput
}

type computeTypePtr string

func ComputeTypePtr(v string) ComputeTypePtrInput {
	return (*computeTypePtr)(&v)
}

func (*computeTypePtr) ElementType() reflect.Type {
	return computeTypePtrType
}

func (in *computeTypePtr) ToComputeTypePtrOutput() ComputeTypePtrOutput {
	return pulumi.ToOutput(in).(ComputeTypePtrOutput)
}

func (in *computeTypePtr) ToComputeTypePtrOutputWithContext(ctx context.Context) ComputeTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ComputeTypePtrOutput)
}

type ContainerType string

const (
	ContainerTypeStorageInitializer = ContainerType("StorageInitializer")
	ContainerTypeInferenceServer    = ContainerType("InferenceServer")
)

func (ContainerType) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerType)(nil)).Elem()
}

func (e ContainerType) ToContainerTypeOutput() ContainerTypeOutput {
	return pulumi.ToOutput(e).(ContainerTypeOutput)
}

func (e ContainerType) ToContainerTypeOutputWithContext(ctx context.Context) ContainerTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContainerTypeOutput)
}

func (e ContainerType) ToContainerTypePtrOutput() ContainerTypePtrOutput {
	return e.ToContainerTypePtrOutputWithContext(context.Background())
}

func (e ContainerType) ToContainerTypePtrOutputWithContext(ctx context.Context) ContainerTypePtrOutput {
	return ContainerType(e).ToContainerTypeOutputWithContext(ctx).ToContainerTypePtrOutputWithContext(ctx)
}

func (e ContainerType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContainerType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContainerType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContainerTypeOutput struct{ *pulumi.OutputState }

func (ContainerTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerType)(nil)).Elem()
}

func (o ContainerTypeOutput) ToContainerTypeOutput() ContainerTypeOutput {
	return o
}

func (o ContainerTypeOutput) ToContainerTypeOutputWithContext(ctx context.Context) ContainerTypeOutput {
	return o
}

func (o ContainerTypeOutput) ToContainerTypePtrOutput() ContainerTypePtrOutput {
	return o.ToContainerTypePtrOutputWithContext(context.Background())
}

func (o ContainerTypeOutput) ToContainerTypePtrOutputWithContext(ctx context.Context) ContainerTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerType) *ContainerType {
		return &v
	}).(ContainerTypePtrOutput)
}

func (o ContainerTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContainerTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContainerTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContainerType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContainerTypePtrOutput struct{ *pulumi.OutputState }

func (ContainerTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerType)(nil)).Elem()
}

func (o ContainerTypePtrOutput) ToContainerTypePtrOutput() ContainerTypePtrOutput {
	return o
}

func (o ContainerTypePtrOutput) ToContainerTypePtrOutputWithContext(ctx context.Context) ContainerTypePtrOutput {
	return o
}

func (o ContainerTypePtrOutput) Elem() ContainerTypeOutput {
	return o.ApplyT(func(v *ContainerType) ContainerType {
		if v != nil {
			return *v
		}
		var ret ContainerType
		return ret
	}).(ContainerTypeOutput)
}

func (o ContainerTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContainerTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContainerType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ContainerTypeInput interface {
	pulumi.Input

	ToContainerTypeOutput() ContainerTypeOutput
	ToContainerTypeOutputWithContext(context.Context) ContainerTypeOutput
}

var containerTypePtrType = reflect.TypeOf((**ContainerType)(nil)).Elem()

type ContainerTypePtrInput interface {
	pulumi.Input

	ToContainerTypePtrOutput() ContainerTypePtrOutput
	ToContainerTypePtrOutputWithContext(context.Context) ContainerTypePtrOutput
}

type containerTypePtr string

func ContainerTypePtr(v string) ContainerTypePtrInput {
	return (*containerTypePtr)(&v)
}

func (*containerTypePtr) ElementType() reflect.Type {
	return containerTypePtrType
}

func (in *containerTypePtr) ToContainerTypePtrOutput() ContainerTypePtrOutput {
	return pulumi.ToOutput(in).(ContainerTypePtrOutput)
}

func (in *containerTypePtr) ToContainerTypePtrOutputWithContext(ctx context.Context) ContainerTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContainerTypePtrOutput)
}

type DataBindingMode string

const (
	DataBindingModeMount    = DataBindingMode("Mount")
	DataBindingModeDownload = DataBindingMode("Download")
	DataBindingModeUpload   = DataBindingMode("Upload")
)

func (DataBindingMode) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBindingMode)(nil)).Elem()
}

func (e DataBindingMode) ToDataBindingModeOutput() DataBindingModeOutput {
	return pulumi.ToOutput(e).(DataBindingModeOutput)
}

func (e DataBindingMode) ToDataBindingModeOutputWithContext(ctx context.Context) DataBindingModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DataBindingModeOutput)
}

func (e DataBindingMode) ToDataBindingModePtrOutput() DataBindingModePtrOutput {
	return e.ToDataBindingModePtrOutputWithContext(context.Background())
}

func (e DataBindingMode) ToDataBindingModePtrOutputWithContext(ctx context.Context) DataBindingModePtrOutput {
	return DataBindingMode(e).ToDataBindingModeOutputWithContext(ctx).ToDataBindingModePtrOutputWithContext(ctx)
}

func (e DataBindingMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataBindingMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataBindingMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataBindingMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DataBindingModeOutput struct{ *pulumi.OutputState }

func (DataBindingModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataBindingMode)(nil)).Elem()
}

func (o DataBindingModeOutput) ToDataBindingModeOutput() DataBindingModeOutput {
	return o
}

func (o DataBindingModeOutput) ToDataBindingModeOutputWithContext(ctx context.Context) DataBindingModeOutput {
	return o
}

func (o DataBindingModeOutput) ToDataBindingModePtrOutput() DataBindingModePtrOutput {
	return o.ToDataBindingModePtrOutputWithContext(context.Background())
}

func (o DataBindingModeOutput) ToDataBindingModePtrOutputWithContext(ctx context.Context) DataBindingModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataBindingMode) *DataBindingMode {
		return &v
	}).(DataBindingModePtrOutput)
}

func (o DataBindingModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DataBindingModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataBindingMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DataBindingModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataBindingModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataBindingMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DataBindingModePtrOutput struct{ *pulumi.OutputState }

func (DataBindingModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataBindingMode)(nil)).Elem()
}

func (o DataBindingModePtrOutput) ToDataBindingModePtrOutput() DataBindingModePtrOutput {
	return o
}

func (o DataBindingModePtrOutput) ToDataBindingModePtrOutputWithContext(ctx context.Context) DataBindingModePtrOutput {
	return o
}

func (o DataBindingModePtrOutput) Elem() DataBindingModeOutput {
	return o.ApplyT(func(v *DataBindingMode) DataBindingMode {
		if v != nil {
			return *v
		}
		var ret DataBindingMode
		return ret
	}).(DataBindingModeOutput)
}

func (o DataBindingModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataBindingModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DataBindingMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DataBindingModeInput interface {
	pulumi.Input

	ToDataBindingModeOutput() DataBindingModeOutput
	ToDataBindingModeOutputWithContext(context.Context) DataBindingModeOutput
}

var dataBindingModePtrType = reflect.TypeOf((**DataBindingMode)(nil)).Elem()

type DataBindingModePtrInput interface {
	pulumi.Input

	ToDataBindingModePtrOutput() DataBindingModePtrOutput
	ToDataBindingModePtrOutputWithContext(context.Context) DataBindingModePtrOutput
}

type dataBindingModePtr string

func DataBindingModePtr(v string) DataBindingModePtrInput {
	return (*dataBindingModePtr)(&v)
}

func (*dataBindingModePtr) ElementType() reflect.Type {
	return dataBindingModePtrType
}

func (in *dataBindingModePtr) ToDataBindingModePtrOutput() DataBindingModePtrOutput {
	return pulumi.ToOutput(in).(DataBindingModePtrOutput)
}

func (in *dataBindingModePtr) ToDataBindingModePtrOutputWithContext(ctx context.Context) DataBindingModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DataBindingModePtrOutput)
}

type DatasetType string

const (
	DatasetTypeTabular = DatasetType("tabular")
	DatasetTypeFile    = DatasetType("file")
)

func (DatasetType) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetType)(nil)).Elem()
}

func (e DatasetType) ToDatasetTypeOutput() DatasetTypeOutput {
	return pulumi.ToOutput(e).(DatasetTypeOutput)
}

func (e DatasetType) ToDatasetTypeOutputWithContext(ctx context.Context) DatasetTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DatasetTypeOutput)
}

func (e DatasetType) ToDatasetTypePtrOutput() DatasetTypePtrOutput {
	return e.ToDatasetTypePtrOutputWithContext(context.Background())
}

func (e DatasetType) ToDatasetTypePtrOutputWithContext(ctx context.Context) DatasetTypePtrOutput {
	return DatasetType(e).ToDatasetTypeOutputWithContext(ctx).ToDatasetTypePtrOutputWithContext(ctx)
}

func (e DatasetType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatasetType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatasetType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DatasetType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DatasetTypeOutput struct{ *pulumi.OutputState }

func (DatasetTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetType)(nil)).Elem()
}

func (o DatasetTypeOutput) ToDatasetTypeOutput() DatasetTypeOutput {
	return o
}

func (o DatasetTypeOutput) ToDatasetTypeOutputWithContext(ctx context.Context) DatasetTypeOutput {
	return o
}

func (o DatasetTypeOutput) ToDatasetTypePtrOutput() DatasetTypePtrOutput {
	return o.ToDatasetTypePtrOutputWithContext(context.Background())
}

func (o DatasetTypeOutput) ToDatasetTypePtrOutputWithContext(ctx context.Context) DatasetTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatasetType) *DatasetType {
		return &v
	}).(DatasetTypePtrOutput)
}

func (o DatasetTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DatasetTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatasetType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DatasetTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatasetTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatasetType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DatasetTypePtrOutput struct{ *pulumi.OutputState }

func (DatasetTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatasetType)(nil)).Elem()
}

func (o DatasetTypePtrOutput) ToDatasetTypePtrOutput() DatasetTypePtrOutput {
	return o
}

func (o DatasetTypePtrOutput) ToDatasetTypePtrOutputWithContext(ctx context.Context) DatasetTypePtrOutput {
	return o
}

func (o DatasetTypePtrOutput) Elem() DatasetTypeOutput {
	return o.ApplyT(func(v *DatasetType) DatasetType {
		if v != nil {
			return *v
		}
		var ret DatasetType
		return ret
	}).(DatasetTypeOutput)
}

func (o DatasetTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatasetTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DatasetType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DatasetTypeInput interface {
	pulumi.Input

	ToDatasetTypeOutput() DatasetTypeOutput
	ToDatasetTypeOutputWithContext(context.Context) DatasetTypeOutput
}

var datasetTypePtrType = reflect.TypeOf((**DatasetType)(nil)).Elem()

type DatasetTypePtrInput interface {
	pulumi.Input

	ToDatasetTypePtrOutput() DatasetTypePtrOutput
	ToDatasetTypePtrOutputWithContext(context.Context) DatasetTypePtrOutput
}

type datasetTypePtr string

func DatasetTypePtr(v string) DatasetTypePtrInput {
	return (*datasetTypePtr)(&v)
}

func (*datasetTypePtr) ElementType() reflect.Type {
	return datasetTypePtrType
}

func (in *datasetTypePtr) ToDatasetTypePtrOutput() DatasetTypePtrOutput {
	return pulumi.ToOutput(in).(DatasetTypePtrOutput)
}

func (in *datasetTypePtr) ToDatasetTypePtrOutputWithContext(ctx context.Context) DatasetTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DatasetTypePtrOutput)
}

type DatastoreTypeArm string

const (
	DatastoreTypeArmBlob       = DatastoreTypeArm("blob")
	DatastoreTypeArmAdls       = DatastoreTypeArm("adls")
	DatastoreTypeArm_Adls_gen2 = DatastoreTypeArm("adls-gen2")
	DatastoreTypeArmDbfs       = DatastoreTypeArm("dbfs")
	DatastoreTypeArmFile       = DatastoreTypeArm("file")
	DatastoreTypeArmMysqldb    = DatastoreTypeArm("mysqldb")
	DatastoreTypeArmSqldb      = DatastoreTypeArm("sqldb")
	DatastoreTypeArmPsqldb     = DatastoreTypeArm("psqldb")
)

func (DatastoreTypeArm) ElementType() reflect.Type {
	return reflect.TypeOf((*DatastoreTypeArm)(nil)).Elem()
}

func (e DatastoreTypeArm) ToDatastoreTypeArmOutput() DatastoreTypeArmOutput {
	return pulumi.ToOutput(e).(DatastoreTypeArmOutput)
}

func (e DatastoreTypeArm) ToDatastoreTypeArmOutputWithContext(ctx context.Context) DatastoreTypeArmOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DatastoreTypeArmOutput)
}

func (e DatastoreTypeArm) ToDatastoreTypeArmPtrOutput() DatastoreTypeArmPtrOutput {
	return e.ToDatastoreTypeArmPtrOutputWithContext(context.Background())
}

func (e DatastoreTypeArm) ToDatastoreTypeArmPtrOutputWithContext(ctx context.Context) DatastoreTypeArmPtrOutput {
	return DatastoreTypeArm(e).ToDatastoreTypeArmOutputWithContext(ctx).ToDatastoreTypeArmPtrOutputWithContext(ctx)
}

func (e DatastoreTypeArm) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatastoreTypeArm) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatastoreTypeArm) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DatastoreTypeArm) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DatastoreTypeArmOutput struct{ *pulumi.OutputState }

func (DatastoreTypeArmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatastoreTypeArm)(nil)).Elem()
}

func (o DatastoreTypeArmOutput) ToDatastoreTypeArmOutput() DatastoreTypeArmOutput {
	return o
}

func (o DatastoreTypeArmOutput) ToDatastoreTypeArmOutputWithContext(ctx context.Context) DatastoreTypeArmOutput {
	return o
}

func (o DatastoreTypeArmOutput) ToDatastoreTypeArmPtrOutput() DatastoreTypeArmPtrOutput {
	return o.ToDatastoreTypeArmPtrOutputWithContext(context.Background())
}

func (o DatastoreTypeArmOutput) ToDatastoreTypeArmPtrOutputWithContext(ctx context.Context) DatastoreTypeArmPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatastoreTypeArm) *DatastoreTypeArm {
		return &v
	}).(DatastoreTypeArmPtrOutput)
}

func (o DatastoreTypeArmOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DatastoreTypeArmOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatastoreTypeArm) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DatastoreTypeArmOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatastoreTypeArmOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatastoreTypeArm) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DatastoreTypeArmPtrOutput struct{ *pulumi.OutputState }

func (DatastoreTypeArmPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatastoreTypeArm)(nil)).Elem()
}

func (o DatastoreTypeArmPtrOutput) ToDatastoreTypeArmPtrOutput() DatastoreTypeArmPtrOutput {
	return o
}

func (o DatastoreTypeArmPtrOutput) ToDatastoreTypeArmPtrOutputWithContext(ctx context.Context) DatastoreTypeArmPtrOutput {
	return o
}

func (o DatastoreTypeArmPtrOutput) Elem() DatastoreTypeArmOutput {
	return o.ApplyT(func(v *DatastoreTypeArm) DatastoreTypeArm {
		if v != nil {
			return *v
		}
		var ret DatastoreTypeArm
		return ret
	}).(DatastoreTypeArmOutput)
}

func (o DatastoreTypeArmPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatastoreTypeArmPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DatastoreTypeArm) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DatastoreTypeArmInput interface {
	pulumi.Input

	ToDatastoreTypeArmOutput() DatastoreTypeArmOutput
	ToDatastoreTypeArmOutputWithContext(context.Context) DatastoreTypeArmOutput
}

var datastoreTypeArmPtrType = reflect.TypeOf((**DatastoreTypeArm)(nil)).Elem()

type DatastoreTypeArmPtrInput interface {
	pulumi.Input

	ToDatastoreTypeArmPtrOutput() DatastoreTypeArmPtrOutput
	ToDatastoreTypeArmPtrOutputWithContext(context.Context) DatastoreTypeArmPtrOutput
}

type datastoreTypeArmPtr string

func DatastoreTypeArmPtr(v string) DatastoreTypeArmPtrInput {
	return (*datastoreTypeArmPtr)(&v)
}

func (*datastoreTypeArmPtr) ElementType() reflect.Type {
	return datastoreTypeArmPtrType
}

func (in *datastoreTypeArmPtr) ToDatastoreTypeArmPtrOutput() DatastoreTypeArmPtrOutput {
	return pulumi.ToOutput(in).(DatastoreTypeArmPtrOutput)
}

func (in *datastoreTypeArmPtr) ToDatastoreTypeArmPtrOutputWithContext(ctx context.Context) DatastoreTypeArmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DatastoreTypeArmPtrOutput)
}

type DistributionType string

const (
	DistributionTypePyTorch    = DistributionType("PyTorch")
	DistributionTypeTensorFlow = DistributionType("TensorFlow")
	DistributionTypeMpi        = DistributionType("Mpi")
)

func (DistributionType) ElementType() reflect.Type {
	return reflect.TypeOf((*DistributionType)(nil)).Elem()
}

func (e DistributionType) ToDistributionTypeOutput() DistributionTypeOutput {
	return pulumi.ToOutput(e).(DistributionTypeOutput)
}

func (e DistributionType) ToDistributionTypeOutputWithContext(ctx context.Context) DistributionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DistributionTypeOutput)
}

func (e DistributionType) ToDistributionTypePtrOutput() DistributionTypePtrOutput {
	return e.ToDistributionTypePtrOutputWithContext(context.Background())
}

func (e DistributionType) ToDistributionTypePtrOutputWithContext(ctx context.Context) DistributionTypePtrOutput {
	return DistributionType(e).ToDistributionTypeOutputWithContext(ctx).ToDistributionTypePtrOutputWithContext(ctx)
}

func (e DistributionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DistributionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DistributionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DistributionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DistributionTypeOutput struct{ *pulumi.OutputState }

func (DistributionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DistributionType)(nil)).Elem()
}

func (o DistributionTypeOutput) ToDistributionTypeOutput() DistributionTypeOutput {
	return o
}

func (o DistributionTypeOutput) ToDistributionTypeOutputWithContext(ctx context.Context) DistributionTypeOutput {
	return o
}

func (o DistributionTypeOutput) ToDistributionTypePtrOutput() DistributionTypePtrOutput {
	return o.ToDistributionTypePtrOutputWithContext(context.Background())
}

func (o DistributionTypeOutput) ToDistributionTypePtrOutputWithContext(ctx context.Context) DistributionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DistributionType) *DistributionType {
		return &v
	}).(DistributionTypePtrOutput)
}

func (o DistributionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DistributionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DistributionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DistributionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DistributionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DistributionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DistributionTypePtrOutput struct{ *pulumi.OutputState }

func (DistributionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DistributionType)(nil)).Elem()
}

func (o DistributionTypePtrOutput) ToDistributionTypePtrOutput() DistributionTypePtrOutput {
	return o
}

func (o DistributionTypePtrOutput) ToDistributionTypePtrOutputWithContext(ctx context.Context) DistributionTypePtrOutput {
	return o
}

func (o DistributionTypePtrOutput) Elem() DistributionTypeOutput {
	return o.ApplyT(func(v *DistributionType) DistributionType {
		if v != nil {
			return *v
		}
		var ret DistributionType
		return ret
	}).(DistributionTypeOutput)
}

func (o DistributionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DistributionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DistributionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DistributionTypeInput interface {
	pulumi.Input

	ToDistributionTypeOutput() DistributionTypeOutput
	ToDistributionTypeOutputWithContext(context.Context) DistributionTypeOutput
}

var distributionTypePtrType = reflect.TypeOf((**DistributionType)(nil)).Elem()

type DistributionTypePtrInput interface {
	pulumi.Input

	ToDistributionTypePtrOutput() DistributionTypePtrOutput
	ToDistributionTypePtrOutputWithContext(context.Context) DistributionTypePtrOutput
}

type distributionTypePtr string

func DistributionTypePtr(v string) DistributionTypePtrInput {
	return (*distributionTypePtr)(&v)
}

func (*distributionTypePtr) ElementType() reflect.Type {
	return distributionTypePtrType
}

func (in *distributionTypePtr) ToDistributionTypePtrOutput() DistributionTypePtrOutput {
	return pulumi.ToOutput(in).(DistributionTypePtrOutput)
}

func (in *distributionTypePtr) ToDistributionTypePtrOutputWithContext(ctx context.Context) DistributionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DistributionTypePtrOutput)
}

type DockerSpecificationType string

const (
	DockerSpecificationTypeBuild = DockerSpecificationType("Build")
	DockerSpecificationTypeImage = DockerSpecificationType("Image")
)

func (DockerSpecificationType) ElementType() reflect.Type {
	return reflect.TypeOf((*DockerSpecificationType)(nil)).Elem()
}

func (e DockerSpecificationType) ToDockerSpecificationTypeOutput() DockerSpecificationTypeOutput {
	return pulumi.ToOutput(e).(DockerSpecificationTypeOutput)
}

func (e DockerSpecificationType) ToDockerSpecificationTypeOutputWithContext(ctx context.Context) DockerSpecificationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DockerSpecificationTypeOutput)
}

func (e DockerSpecificationType) ToDockerSpecificationTypePtrOutput() DockerSpecificationTypePtrOutput {
	return e.ToDockerSpecificationTypePtrOutputWithContext(context.Background())
}

func (e DockerSpecificationType) ToDockerSpecificationTypePtrOutputWithContext(ctx context.Context) DockerSpecificationTypePtrOutput {
	return DockerSpecificationType(e).ToDockerSpecificationTypeOutputWithContext(ctx).ToDockerSpecificationTypePtrOutputWithContext(ctx)
}

func (e DockerSpecificationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DockerSpecificationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DockerSpecificationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DockerSpecificationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DockerSpecificationTypeOutput struct{ *pulumi.OutputState }

func (DockerSpecificationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DockerSpecificationType)(nil)).Elem()
}

func (o DockerSpecificationTypeOutput) ToDockerSpecificationTypeOutput() DockerSpecificationTypeOutput {
	return o
}

func (o DockerSpecificationTypeOutput) ToDockerSpecificationTypeOutputWithContext(ctx context.Context) DockerSpecificationTypeOutput {
	return o
}

func (o DockerSpecificationTypeOutput) ToDockerSpecificationTypePtrOutput() DockerSpecificationTypePtrOutput {
	return o.ToDockerSpecificationTypePtrOutputWithContext(context.Background())
}

func (o DockerSpecificationTypeOutput) ToDockerSpecificationTypePtrOutputWithContext(ctx context.Context) DockerSpecificationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DockerSpecificationType) *DockerSpecificationType {
		return &v
	}).(DockerSpecificationTypePtrOutput)
}

func (o DockerSpecificationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DockerSpecificationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DockerSpecificationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DockerSpecificationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DockerSpecificationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DockerSpecificationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DockerSpecificationTypePtrOutput struct{ *pulumi.OutputState }

func (DockerSpecificationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DockerSpecificationType)(nil)).Elem()
}

func (o DockerSpecificationTypePtrOutput) ToDockerSpecificationTypePtrOutput() DockerSpecificationTypePtrOutput {
	return o
}

func (o DockerSpecificationTypePtrOutput) ToDockerSpecificationTypePtrOutputWithContext(ctx context.Context) DockerSpecificationTypePtrOutput {
	return o
}

func (o DockerSpecificationTypePtrOutput) Elem() DockerSpecificationTypeOutput {
	return o.ApplyT(func(v *DockerSpecificationType) DockerSpecificationType {
		if v != nil {
			return *v
		}
		var ret DockerSpecificationType
		return ret
	}).(DockerSpecificationTypeOutput)
}

func (o DockerSpecificationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DockerSpecificationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DockerSpecificationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DockerSpecificationTypeInput interface {
	pulumi.Input

	ToDockerSpecificationTypeOutput() DockerSpecificationTypeOutput
	ToDockerSpecificationTypeOutputWithContext(context.Context) DockerSpecificationTypeOutput
}

var dockerSpecificationTypePtrType = reflect.TypeOf((**DockerSpecificationType)(nil)).Elem()

type DockerSpecificationTypePtrInput interface {
	pulumi.Input

	ToDockerSpecificationTypePtrOutput() DockerSpecificationTypePtrOutput
	ToDockerSpecificationTypePtrOutputWithContext(context.Context) DockerSpecificationTypePtrOutput
}

type dockerSpecificationTypePtr string

func DockerSpecificationTypePtr(v string) DockerSpecificationTypePtrInput {
	return (*dockerSpecificationTypePtr)(&v)
}

func (*dockerSpecificationTypePtr) ElementType() reflect.Type {
	return dockerSpecificationTypePtrType
}

func (in *dockerSpecificationTypePtr) ToDockerSpecificationTypePtrOutput() DockerSpecificationTypePtrOutput {
	return pulumi.ToOutput(in).(DockerSpecificationTypePtrOutput)
}

func (in *dockerSpecificationTypePtr) ToDockerSpecificationTypePtrOutputWithContext(ctx context.Context) DockerSpecificationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DockerSpecificationTypePtrOutput)
}

type EarlyTerminationPolicyType string

const (
	EarlyTerminationPolicyTypeBandit              = EarlyTerminationPolicyType("Bandit")
	EarlyTerminationPolicyTypeMedianStopping      = EarlyTerminationPolicyType("MedianStopping")
	EarlyTerminationPolicyTypeTruncationSelection = EarlyTerminationPolicyType("TruncationSelection")
)

func (EarlyTerminationPolicyType) ElementType() reflect.Type {
	return reflect.TypeOf((*EarlyTerminationPolicyType)(nil)).Elem()
}

func (e EarlyTerminationPolicyType) ToEarlyTerminationPolicyTypeOutput() EarlyTerminationPolicyTypeOutput {
	return pulumi.ToOutput(e).(EarlyTerminationPolicyTypeOutput)
}

func (e EarlyTerminationPolicyType) ToEarlyTerminationPolicyTypeOutputWithContext(ctx context.Context) EarlyTerminationPolicyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EarlyTerminationPolicyTypeOutput)
}

func (e EarlyTerminationPolicyType) ToEarlyTerminationPolicyTypePtrOutput() EarlyTerminationPolicyTypePtrOutput {
	return e.ToEarlyTerminationPolicyTypePtrOutputWithContext(context.Background())
}

func (e EarlyTerminationPolicyType) ToEarlyTerminationPolicyTypePtrOutputWithContext(ctx context.Context) EarlyTerminationPolicyTypePtrOutput {
	return EarlyTerminationPolicyType(e).ToEarlyTerminationPolicyTypeOutputWithContext(ctx).ToEarlyTerminationPolicyTypePtrOutputWithContext(ctx)
}

func (e EarlyTerminationPolicyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EarlyTerminationPolicyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EarlyTerminationPolicyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EarlyTerminationPolicyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EarlyTerminationPolicyTypeOutput struct{ *pulumi.OutputState }

func (EarlyTerminationPolicyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EarlyTerminationPolicyType)(nil)).Elem()
}

func (o EarlyTerminationPolicyTypeOutput) ToEarlyTerminationPolicyTypeOutput() EarlyTerminationPolicyTypeOutput {
	return o
}

func (o EarlyTerminationPolicyTypeOutput) ToEarlyTerminationPolicyTypeOutputWithContext(ctx context.Context) EarlyTerminationPolicyTypeOutput {
	return o
}

func (o EarlyTerminationPolicyTypeOutput) ToEarlyTerminationPolicyTypePtrOutput() EarlyTerminationPolicyTypePtrOutput {
	return o.ToEarlyTerminationPolicyTypePtrOutputWithContext(context.Background())
}

func (o EarlyTerminationPolicyTypeOutput) ToEarlyTerminationPolicyTypePtrOutputWithContext(ctx context.Context) EarlyTerminationPolicyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EarlyTerminationPolicyType) *EarlyTerminationPolicyType {
		return &v
	}).(EarlyTerminationPolicyTypePtrOutput)
}

func (o EarlyTerminationPolicyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EarlyTerminationPolicyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EarlyTerminationPolicyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EarlyTerminationPolicyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EarlyTerminationPolicyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EarlyTerminationPolicyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EarlyTerminationPolicyTypePtrOutput struct{ *pulumi.OutputState }

func (EarlyTerminationPolicyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EarlyTerminationPolicyType)(nil)).Elem()
}

func (o EarlyTerminationPolicyTypePtrOutput) ToEarlyTerminationPolicyTypePtrOutput() EarlyTerminationPolicyTypePtrOutput {
	return o
}

func (o EarlyTerminationPolicyTypePtrOutput) ToEarlyTerminationPolicyTypePtrOutputWithContext(ctx context.Context) EarlyTerminationPolicyTypePtrOutput {
	return o
}

func (o EarlyTerminationPolicyTypePtrOutput) Elem() EarlyTerminationPolicyTypeOutput {
	return o.ApplyT(func(v *EarlyTerminationPolicyType) EarlyTerminationPolicyType {
		if v != nil {
			return *v
		}
		var ret EarlyTerminationPolicyType
		return ret
	}).(EarlyTerminationPolicyTypeOutput)
}

func (o EarlyTerminationPolicyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EarlyTerminationPolicyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EarlyTerminationPolicyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EarlyTerminationPolicyTypeInput interface {
	pulumi.Input

	ToEarlyTerminationPolicyTypeOutput() EarlyTerminationPolicyTypeOutput
	ToEarlyTerminationPolicyTypeOutputWithContext(context.Context) EarlyTerminationPolicyTypeOutput
}

var earlyTerminationPolicyTypePtrType = reflect.TypeOf((**EarlyTerminationPolicyType)(nil)).Elem()

type EarlyTerminationPolicyTypePtrInput interface {
	pulumi.Input

	ToEarlyTerminationPolicyTypePtrOutput() EarlyTerminationPolicyTypePtrOutput
	ToEarlyTerminationPolicyTypePtrOutputWithContext(context.Context) EarlyTerminationPolicyTypePtrOutput
}

type earlyTerminationPolicyTypePtr string

func EarlyTerminationPolicyTypePtr(v string) EarlyTerminationPolicyTypePtrInput {
	return (*earlyTerminationPolicyTypePtr)(&v)
}

func (*earlyTerminationPolicyTypePtr) ElementType() reflect.Type {
	return earlyTerminationPolicyTypePtrType
}

func (in *earlyTerminationPolicyTypePtr) ToEarlyTerminationPolicyTypePtrOutput() EarlyTerminationPolicyTypePtrOutput {
	return pulumi.ToOutput(in).(EarlyTerminationPolicyTypePtrOutput)
}

func (in *earlyTerminationPolicyTypePtr) ToEarlyTerminationPolicyTypePtrOutputWithContext(ctx context.Context) EarlyTerminationPolicyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EarlyTerminationPolicyTypePtrOutput)
}

type EncryptionStatus string

const (
	EncryptionStatusEnabled  = EncryptionStatus("Enabled")
	EncryptionStatusDisabled = EncryptionStatus("Disabled")
)

func (EncryptionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionStatus)(nil)).Elem()
}

func (e EncryptionStatus) ToEncryptionStatusOutput() EncryptionStatusOutput {
	return pulumi.ToOutput(e).(EncryptionStatusOutput)
}

func (e EncryptionStatus) ToEncryptionStatusOutputWithContext(ctx context.Context) EncryptionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EncryptionStatusOutput)
}

func (e EncryptionStatus) ToEncryptionStatusPtrOutput() EncryptionStatusPtrOutput {
	return e.ToEncryptionStatusPtrOutputWithContext(context.Background())
}

func (e EncryptionStatus) ToEncryptionStatusPtrOutputWithContext(ctx context.Context) EncryptionStatusPtrOutput {
	return EncryptionStatus(e).ToEncryptionStatusOutputWithContext(ctx).ToEncryptionStatusPtrOutputWithContext(ctx)
}

func (e EncryptionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncryptionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncryptionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EncryptionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EncryptionStatusOutput struct{ *pulumi.OutputState }

func (EncryptionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionStatus)(nil)).Elem()
}

func (o EncryptionStatusOutput) ToEncryptionStatusOutput() EncryptionStatusOutput {
	return o
}

func (o EncryptionStatusOutput) ToEncryptionStatusOutputWithContext(ctx context.Context) EncryptionStatusOutput {
	return o
}

func (o EncryptionStatusOutput) ToEncryptionStatusPtrOutput() EncryptionStatusPtrOutput {
	return o.ToEncryptionStatusPtrOutputWithContext(context.Background())
}

func (o EncryptionStatusOutput) ToEncryptionStatusPtrOutputWithContext(ctx context.Context) EncryptionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionStatus) *EncryptionStatus {
		return &v
	}).(EncryptionStatusPtrOutput)
}

func (o EncryptionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EncryptionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncryptionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EncryptionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncryptionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncryptionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EncryptionStatusPtrOutput struct{ *pulumi.OutputState }

func (EncryptionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionStatus)(nil)).Elem()
}

func (o EncryptionStatusPtrOutput) ToEncryptionStatusPtrOutput() EncryptionStatusPtrOutput {
	return o
}

func (o EncryptionStatusPtrOutput) ToEncryptionStatusPtrOutputWithContext(ctx context.Context) EncryptionStatusPtrOutput {
	return o
}

func (o EncryptionStatusPtrOutput) Elem() EncryptionStatusOutput {
	return o.ApplyT(func(v *EncryptionStatus) EncryptionStatus {
		if v != nil {
			return *v
		}
		var ret EncryptionStatus
		return ret
	}).(EncryptionStatusOutput)
}

func (o EncryptionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncryptionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EncryptionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EncryptionStatusInput interface {
	pulumi.Input

	ToEncryptionStatusOutput() EncryptionStatusOutput
	ToEncryptionStatusOutputWithContext(context.Context) EncryptionStatusOutput
}

var encryptionStatusPtrType = reflect.TypeOf((**EncryptionStatus)(nil)).Elem()

type EncryptionStatusPtrInput interface {
	pulumi.Input

	ToEncryptionStatusPtrOutput() EncryptionStatusPtrOutput
	ToEncryptionStatusPtrOutputWithContext(context.Context) EncryptionStatusPtrOutput
}

type encryptionStatusPtr string

func EncryptionStatusPtr(v string) EncryptionStatusPtrInput {
	return (*encryptionStatusPtr)(&v)
}

func (*encryptionStatusPtr) ElementType() reflect.Type {
	return encryptionStatusPtrType
}

func (in *encryptionStatusPtr) ToEncryptionStatusPtrOutput() EncryptionStatusPtrOutput {
	return pulumi.ToOutput(in).(EncryptionStatusPtrOutput)
}

func (in *encryptionStatusPtr) ToEncryptionStatusPtrOutputWithContext(ctx context.Context) EncryptionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EncryptionStatusPtrOutput)
}

type EndpointAuthMode string

const (
	EndpointAuthModeAMLToken = EndpointAuthMode("AMLToken")
	EndpointAuthModeKey      = EndpointAuthMode("Key")
	EndpointAuthModeAADToken = EndpointAuthMode("AADToken")
)

func (EndpointAuthMode) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointAuthMode)(nil)).Elem()
}

func (e EndpointAuthMode) ToEndpointAuthModeOutput() EndpointAuthModeOutput {
	return pulumi.ToOutput(e).(EndpointAuthModeOutput)
}

func (e EndpointAuthMode) ToEndpointAuthModeOutputWithContext(ctx context.Context) EndpointAuthModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EndpointAuthModeOutput)
}

func (e EndpointAuthMode) ToEndpointAuthModePtrOutput() EndpointAuthModePtrOutput {
	return e.ToEndpointAuthModePtrOutputWithContext(context.Background())
}

func (e EndpointAuthMode) ToEndpointAuthModePtrOutputWithContext(ctx context.Context) EndpointAuthModePtrOutput {
	return EndpointAuthMode(e).ToEndpointAuthModeOutputWithContext(ctx).ToEndpointAuthModePtrOutputWithContext(ctx)
}

func (e EndpointAuthMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EndpointAuthMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EndpointAuthMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EndpointAuthMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EndpointAuthModeOutput struct{ *pulumi.OutputState }

func (EndpointAuthModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointAuthMode)(nil)).Elem()
}

func (o EndpointAuthModeOutput) ToEndpointAuthModeOutput() EndpointAuthModeOutput {
	return o
}

func (o EndpointAuthModeOutput) ToEndpointAuthModeOutputWithContext(ctx context.Context) EndpointAuthModeOutput {
	return o
}

func (o EndpointAuthModeOutput) ToEndpointAuthModePtrOutput() EndpointAuthModePtrOutput {
	return o.ToEndpointAuthModePtrOutputWithContext(context.Background())
}

func (o EndpointAuthModeOutput) ToEndpointAuthModePtrOutputWithContext(ctx context.Context) EndpointAuthModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointAuthMode) *EndpointAuthMode {
		return &v
	}).(EndpointAuthModePtrOutput)
}

func (o EndpointAuthModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EndpointAuthModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EndpointAuthMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EndpointAuthModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EndpointAuthModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EndpointAuthMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EndpointAuthModePtrOutput struct{ *pulumi.OutputState }

func (EndpointAuthModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointAuthMode)(nil)).Elem()
}

func (o EndpointAuthModePtrOutput) ToEndpointAuthModePtrOutput() EndpointAuthModePtrOutput {
	return o
}

func (o EndpointAuthModePtrOutput) ToEndpointAuthModePtrOutputWithContext(ctx context.Context) EndpointAuthModePtrOutput {
	return o
}

func (o EndpointAuthModePtrOutput) Elem() EndpointAuthModeOutput {
	return o.ApplyT(func(v *EndpointAuthMode) EndpointAuthMode {
		if v != nil {
			return *v
		}
		var ret EndpointAuthMode
		return ret
	}).(EndpointAuthModeOutput)
}

func (o EndpointAuthModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EndpointAuthModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EndpointAuthMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EndpointAuthModeInput interface {
	pulumi.Input

	ToEndpointAuthModeOutput() EndpointAuthModeOutput
	ToEndpointAuthModeOutputWithContext(context.Context) EndpointAuthModeOutput
}

var endpointAuthModePtrType = reflect.TypeOf((**EndpointAuthMode)(nil)).Elem()

type EndpointAuthModePtrInput interface {
	pulumi.Input

	ToEndpointAuthModePtrOutput() EndpointAuthModePtrOutput
	ToEndpointAuthModePtrOutputWithContext(context.Context) EndpointAuthModePtrOutput
}

type endpointAuthModePtr string

func EndpointAuthModePtr(v string) EndpointAuthModePtrInput {
	return (*endpointAuthModePtr)(&v)
}

func (*endpointAuthModePtr) ElementType() reflect.Type {
	return endpointAuthModePtrType
}

func (in *endpointAuthModePtr) ToEndpointAuthModePtrOutput() EndpointAuthModePtrOutput {
	return pulumi.ToOutput(in).(EndpointAuthModePtrOutput)
}

func (in *endpointAuthModePtr) ToEndpointAuthModePtrOutputWithContext(ctx context.Context) EndpointAuthModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EndpointAuthModePtrOutput)
}

type EndpointComputeType string

const (
	EndpointComputeTypeManaged        = EndpointComputeType("Managed")
	EndpointComputeTypeK8S            = EndpointComputeType("K8S")
	EndpointComputeTypeAzureMLCompute = EndpointComputeType("AzureMLCompute")
)

func (EndpointComputeType) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointComputeType)(nil)).Elem()
}

func (e EndpointComputeType) ToEndpointComputeTypeOutput() EndpointComputeTypeOutput {
	return pulumi.ToOutput(e).(EndpointComputeTypeOutput)
}

func (e EndpointComputeType) ToEndpointComputeTypeOutputWithContext(ctx context.Context) EndpointComputeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EndpointComputeTypeOutput)
}

func (e EndpointComputeType) ToEndpointComputeTypePtrOutput() EndpointComputeTypePtrOutput {
	return e.ToEndpointComputeTypePtrOutputWithContext(context.Background())
}

func (e EndpointComputeType) ToEndpointComputeTypePtrOutputWithContext(ctx context.Context) EndpointComputeTypePtrOutput {
	return EndpointComputeType(e).ToEndpointComputeTypeOutputWithContext(ctx).ToEndpointComputeTypePtrOutputWithContext(ctx)
}

func (e EndpointComputeType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EndpointComputeType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EndpointComputeType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EndpointComputeType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EndpointComputeTypeOutput struct{ *pulumi.OutputState }

func (EndpointComputeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointComputeType)(nil)).Elem()
}

func (o EndpointComputeTypeOutput) ToEndpointComputeTypeOutput() EndpointComputeTypeOutput {
	return o
}

func (o EndpointComputeTypeOutput) ToEndpointComputeTypeOutputWithContext(ctx context.Context) EndpointComputeTypeOutput {
	return o
}

func (o EndpointComputeTypeOutput) ToEndpointComputeTypePtrOutput() EndpointComputeTypePtrOutput {
	return o.ToEndpointComputeTypePtrOutputWithContext(context.Background())
}

func (o EndpointComputeTypeOutput) ToEndpointComputeTypePtrOutputWithContext(ctx context.Context) EndpointComputeTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointComputeType) *EndpointComputeType {
		return &v
	}).(EndpointComputeTypePtrOutput)
}

func (o EndpointComputeTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EndpointComputeTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EndpointComputeType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EndpointComputeTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EndpointComputeTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EndpointComputeType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EndpointComputeTypePtrOutput struct{ *pulumi.OutputState }

func (EndpointComputeTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointComputeType)(nil)).Elem()
}

func (o EndpointComputeTypePtrOutput) ToEndpointComputeTypePtrOutput() EndpointComputeTypePtrOutput {
	return o
}

func (o EndpointComputeTypePtrOutput) ToEndpointComputeTypePtrOutputWithContext(ctx context.Context) EndpointComputeTypePtrOutput {
	return o
}

func (o EndpointComputeTypePtrOutput) Elem() EndpointComputeTypeOutput {
	return o.ApplyT(func(v *EndpointComputeType) EndpointComputeType {
		if v != nil {
			return *v
		}
		var ret EndpointComputeType
		return ret
	}).(EndpointComputeTypeOutput)
}

func (o EndpointComputeTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EndpointComputeTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EndpointComputeType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EndpointComputeTypeInput interface {
	pulumi.Input

	ToEndpointComputeTypeOutput() EndpointComputeTypeOutput
	ToEndpointComputeTypeOutputWithContext(context.Context) EndpointComputeTypeOutput
}

var endpointComputeTypePtrType = reflect.TypeOf((**EndpointComputeType)(nil)).Elem()

type EndpointComputeTypePtrInput interface {
	pulumi.Input

	ToEndpointComputeTypePtrOutput() EndpointComputeTypePtrOutput
	ToEndpointComputeTypePtrOutputWithContext(context.Context) EndpointComputeTypePtrOutput
}

type endpointComputeTypePtr string

func EndpointComputeTypePtr(v string) EndpointComputeTypePtrInput {
	return (*endpointComputeTypePtr)(&v)
}

func (*endpointComputeTypePtr) ElementType() reflect.Type {
	return endpointComputeTypePtrType
}

func (in *endpointComputeTypePtr) ToEndpointComputeTypePtrOutput() EndpointComputeTypePtrOutput {
	return pulumi.ToOutput(in).(EndpointComputeTypePtrOutput)
}

func (in *endpointComputeTypePtr) ToEndpointComputeTypePtrOutputWithContext(ctx context.Context) EndpointComputeTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EndpointComputeTypePtrOutput)
}

type Goal string

const (
	GoalMinimize = Goal("Minimize")
	GoalMaximize = Goal("Maximize")
)

func (Goal) ElementType() reflect.Type {
	return reflect.TypeOf((*Goal)(nil)).Elem()
}

func (e Goal) ToGoalOutput() GoalOutput {
	return pulumi.ToOutput(e).(GoalOutput)
}

func (e Goal) ToGoalOutputWithContext(ctx context.Context) GoalOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GoalOutput)
}

func (e Goal) ToGoalPtrOutput() GoalPtrOutput {
	return e.ToGoalPtrOutputWithContext(context.Background())
}

func (e Goal) ToGoalPtrOutputWithContext(ctx context.Context) GoalPtrOutput {
	return Goal(e).ToGoalOutputWithContext(ctx).ToGoalPtrOutputWithContext(ctx)
}

func (e Goal) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Goal) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Goal) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Goal) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type GoalOutput struct{ *pulumi.OutputState }

func (GoalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Goal)(nil)).Elem()
}

func (o GoalOutput) ToGoalOutput() GoalOutput {
	return o
}

func (o GoalOutput) ToGoalOutputWithContext(ctx context.Context) GoalOutput {
	return o
}

func (o GoalOutput) ToGoalPtrOutput() GoalPtrOutput {
	return o.ToGoalPtrOutputWithContext(context.Background())
}

func (o GoalOutput) ToGoalPtrOutputWithContext(ctx context.Context) GoalPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Goal) *Goal {
		return &v
	}).(GoalPtrOutput)
}

func (o GoalOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GoalOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Goal) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GoalOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GoalOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Goal) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GoalPtrOutput struct{ *pulumi.OutputState }

func (GoalPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Goal)(nil)).Elem()
}

func (o GoalPtrOutput) ToGoalPtrOutput() GoalPtrOutput {
	return o
}

func (o GoalPtrOutput) ToGoalPtrOutputWithContext(ctx context.Context) GoalPtrOutput {
	return o
}

func (o GoalPtrOutput) Elem() GoalOutput {
	return o.ApplyT(func(v *Goal) Goal {
		if v != nil {
			return *v
		}
		var ret Goal
		return ret
	}).(GoalOutput)
}

func (o GoalPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GoalPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Goal) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type GoalInput interface {
	pulumi.Input

	ToGoalOutput() GoalOutput
	ToGoalOutputWithContext(context.Context) GoalOutput
}

var goalPtrType = reflect.TypeOf((**Goal)(nil)).Elem()

type GoalPtrInput interface {
	pulumi.Input

	ToGoalPtrOutput() GoalPtrOutput
	ToGoalPtrOutputWithContext(context.Context) GoalPtrOutput
}

type goalPtr string

func GoalPtr(v string) GoalPtrInput {
	return (*goalPtr)(&v)
}

func (*goalPtr) ElementType() reflect.Type {
	return goalPtrType
}

func (in *goalPtr) ToGoalPtrOutput() GoalPtrOutput {
	return pulumi.ToOutput(in).(GoalPtrOutput)
}

func (in *goalPtr) ToGoalPtrOutputWithContext(ctx context.Context) GoalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GoalPtrOutput)
}

type Header string

const (
	Header_All_files_have_same_headers = Header("all_files_have_same_headers")
	Header_Only_first_file_has_headers = Header("only_first_file_has_headers")
	Header_No_headers                  = Header("no_headers")
	Header_Combine_all_files_headers   = Header("combine_all_files_headers")
)

func (Header) ElementType() reflect.Type {
	return reflect.TypeOf((*Header)(nil)).Elem()
}

func (e Header) ToHeaderOutput() HeaderOutput {
	return pulumi.ToOutput(e).(HeaderOutput)
}

func (e Header) ToHeaderOutputWithContext(ctx context.Context) HeaderOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HeaderOutput)
}

func (e Header) ToHeaderPtrOutput() HeaderPtrOutput {
	return e.ToHeaderPtrOutputWithContext(context.Background())
}

func (e Header) ToHeaderPtrOutputWithContext(ctx context.Context) HeaderPtrOutput {
	return Header(e).ToHeaderOutputWithContext(ctx).ToHeaderPtrOutputWithContext(ctx)
}

func (e Header) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Header) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Header) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Header) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HeaderOutput struct{ *pulumi.OutputState }

func (HeaderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Header)(nil)).Elem()
}

func (o HeaderOutput) ToHeaderOutput() HeaderOutput {
	return o
}

func (o HeaderOutput) ToHeaderOutputWithContext(ctx context.Context) HeaderOutput {
	return o
}

func (o HeaderOutput) ToHeaderPtrOutput() HeaderPtrOutput {
	return o.ToHeaderPtrOutputWithContext(context.Background())
}

func (o HeaderOutput) ToHeaderPtrOutputWithContext(ctx context.Context) HeaderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Header) *Header {
		return &v
	}).(HeaderPtrOutput)
}

func (o HeaderOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HeaderOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Header) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HeaderOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HeaderOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Header) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HeaderPtrOutput struct{ *pulumi.OutputState }

func (HeaderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Header)(nil)).Elem()
}

func (o HeaderPtrOutput) ToHeaderPtrOutput() HeaderPtrOutput {
	return o
}

func (o HeaderPtrOutput) ToHeaderPtrOutputWithContext(ctx context.Context) HeaderPtrOutput {
	return o
}

func (o HeaderPtrOutput) Elem() HeaderOutput {
	return o.ApplyT(func(v *Header) Header {
		if v != nil {
			return *v
		}
		var ret Header
		return ret
	}).(HeaderOutput)
}

func (o HeaderPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HeaderPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Header) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HeaderInput interface {
	pulumi.Input

	ToHeaderOutput() HeaderOutput
	ToHeaderOutputWithContext(context.Context) HeaderOutput
}

var headerPtrType = reflect.TypeOf((**Header)(nil)).Elem()

type HeaderPtrInput interface {
	pulumi.Input

	ToHeaderPtrOutput() HeaderPtrOutput
	ToHeaderPtrOutputWithContext(context.Context) HeaderPtrOutput
}

type headerPtr string

func HeaderPtr(v string) HeaderPtrInput {
	return (*headerPtr)(&v)
}

func (*headerPtr) ElementType() reflect.Type {
	return headerPtrType
}

func (in *headerPtr) ToHeaderPtrOutput() HeaderPtrOutput {
	return pulumi.ToOutput(in).(HeaderPtrOutput)
}

func (in *headerPtr) ToHeaderPtrOutputWithContext(ctx context.Context) HeaderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HeaderPtrOutput)
}

type IdentityConfigurationType string

const (
	IdentityConfigurationTypeManaged  = IdentityConfigurationType("Managed")
	IdentityConfigurationTypeAMLToken = IdentityConfigurationType("AMLToken")
)

func (IdentityConfigurationType) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityConfigurationType)(nil)).Elem()
}

func (e IdentityConfigurationType) ToIdentityConfigurationTypeOutput() IdentityConfigurationTypeOutput {
	return pulumi.ToOutput(e).(IdentityConfigurationTypeOutput)
}

func (e IdentityConfigurationType) ToIdentityConfigurationTypeOutputWithContext(ctx context.Context) IdentityConfigurationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IdentityConfigurationTypeOutput)
}

func (e IdentityConfigurationType) ToIdentityConfigurationTypePtrOutput() IdentityConfigurationTypePtrOutput {
	return e.ToIdentityConfigurationTypePtrOutputWithContext(context.Background())
}

func (e IdentityConfigurationType) ToIdentityConfigurationTypePtrOutputWithContext(ctx context.Context) IdentityConfigurationTypePtrOutput {
	return IdentityConfigurationType(e).ToIdentityConfigurationTypeOutputWithContext(ctx).ToIdentityConfigurationTypePtrOutputWithContext(ctx)
}

func (e IdentityConfigurationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityConfigurationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityConfigurationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IdentityConfigurationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IdentityConfigurationTypeOutput struct{ *pulumi.OutputState }

func (IdentityConfigurationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityConfigurationType)(nil)).Elem()
}

func (o IdentityConfigurationTypeOutput) ToIdentityConfigurationTypeOutput() IdentityConfigurationTypeOutput {
	return o
}

func (o IdentityConfigurationTypeOutput) ToIdentityConfigurationTypeOutputWithContext(ctx context.Context) IdentityConfigurationTypeOutput {
	return o
}

func (o IdentityConfigurationTypeOutput) ToIdentityConfigurationTypePtrOutput() IdentityConfigurationTypePtrOutput {
	return o.ToIdentityConfigurationTypePtrOutputWithContext(context.Background())
}

func (o IdentityConfigurationTypeOutput) ToIdentityConfigurationTypePtrOutputWithContext(ctx context.Context) IdentityConfigurationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityConfigurationType) *IdentityConfigurationType {
		return &v
	}).(IdentityConfigurationTypePtrOutput)
}

func (o IdentityConfigurationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IdentityConfigurationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityConfigurationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IdentityConfigurationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityConfigurationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityConfigurationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IdentityConfigurationTypePtrOutput struct{ *pulumi.OutputState }

func (IdentityConfigurationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityConfigurationType)(nil)).Elem()
}

func (o IdentityConfigurationTypePtrOutput) ToIdentityConfigurationTypePtrOutput() IdentityConfigurationTypePtrOutput {
	return o
}

func (o IdentityConfigurationTypePtrOutput) ToIdentityConfigurationTypePtrOutputWithContext(ctx context.Context) IdentityConfigurationTypePtrOutput {
	return o
}

func (o IdentityConfigurationTypePtrOutput) Elem() IdentityConfigurationTypeOutput {
	return o.ApplyT(func(v *IdentityConfigurationType) IdentityConfigurationType {
		if v != nil {
			return *v
		}
		var ret IdentityConfigurationType
		return ret
	}).(IdentityConfigurationTypeOutput)
}

func (o IdentityConfigurationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityConfigurationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IdentityConfigurationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IdentityConfigurationTypeInput interface {
	pulumi.Input

	ToIdentityConfigurationTypeOutput() IdentityConfigurationTypeOutput
	ToIdentityConfigurationTypeOutputWithContext(context.Context) IdentityConfigurationTypeOutput
}

var identityConfigurationTypePtrType = reflect.TypeOf((**IdentityConfigurationType)(nil)).Elem()

type IdentityConfigurationTypePtrInput interface {
	pulumi.Input

	ToIdentityConfigurationTypePtrOutput() IdentityConfigurationTypePtrOutput
	ToIdentityConfigurationTypePtrOutputWithContext(context.Context) IdentityConfigurationTypePtrOutput
}

type identityConfigurationTypePtr string

func IdentityConfigurationTypePtr(v string) IdentityConfigurationTypePtrInput {
	return (*identityConfigurationTypePtr)(&v)
}

func (*identityConfigurationTypePtr) ElementType() reflect.Type {
	return identityConfigurationTypePtrType
}

func (in *identityConfigurationTypePtr) ToIdentityConfigurationTypePtrOutput() IdentityConfigurationTypePtrOutput {
	return pulumi.ToOutput(in).(IdentityConfigurationTypePtrOutput)
}

func (in *identityConfigurationTypePtr) ToIdentityConfigurationTypePtrOutputWithContext(ctx context.Context) IdentityConfigurationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IdentityConfigurationTypePtrOutput)
}

type ImageAnnotationType string

const (
	ImageAnnotationTypeClassification       = ImageAnnotationType("Classification")
	ImageAnnotationTypeBoundingBox          = ImageAnnotationType("BoundingBox")
	ImageAnnotationTypeInstanceSegmentation = ImageAnnotationType("InstanceSegmentation")
)

func (ImageAnnotationType) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageAnnotationType)(nil)).Elem()
}

func (e ImageAnnotationType) ToImageAnnotationTypeOutput() ImageAnnotationTypeOutput {
	return pulumi.ToOutput(e).(ImageAnnotationTypeOutput)
}

func (e ImageAnnotationType) ToImageAnnotationTypeOutputWithContext(ctx context.Context) ImageAnnotationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ImageAnnotationTypeOutput)
}

func (e ImageAnnotationType) ToImageAnnotationTypePtrOutput() ImageAnnotationTypePtrOutput {
	return e.ToImageAnnotationTypePtrOutputWithContext(context.Background())
}

func (e ImageAnnotationType) ToImageAnnotationTypePtrOutputWithContext(ctx context.Context) ImageAnnotationTypePtrOutput {
	return ImageAnnotationType(e).ToImageAnnotationTypeOutputWithContext(ctx).ToImageAnnotationTypePtrOutputWithContext(ctx)
}

func (e ImageAnnotationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ImageAnnotationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ImageAnnotationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ImageAnnotationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ImageAnnotationTypeOutput struct{ *pulumi.OutputState }

func (ImageAnnotationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageAnnotationType)(nil)).Elem()
}

func (o ImageAnnotationTypeOutput) ToImageAnnotationTypeOutput() ImageAnnotationTypeOutput {
	return o
}

func (o ImageAnnotationTypeOutput) ToImageAnnotationTypeOutputWithContext(ctx context.Context) ImageAnnotationTypeOutput {
	return o
}

func (o ImageAnnotationTypeOutput) ToImageAnnotationTypePtrOutput() ImageAnnotationTypePtrOutput {
	return o.ToImageAnnotationTypePtrOutputWithContext(context.Background())
}

func (o ImageAnnotationTypeOutput) ToImageAnnotationTypePtrOutputWithContext(ctx context.Context) ImageAnnotationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageAnnotationType) *ImageAnnotationType {
		return &v
	}).(ImageAnnotationTypePtrOutput)
}

func (o ImageAnnotationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ImageAnnotationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ImageAnnotationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ImageAnnotationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ImageAnnotationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ImageAnnotationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ImageAnnotationTypePtrOutput struct{ *pulumi.OutputState }

func (ImageAnnotationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageAnnotationType)(nil)).Elem()
}

func (o ImageAnnotationTypePtrOutput) ToImageAnnotationTypePtrOutput() ImageAnnotationTypePtrOutput {
	return o
}

func (o ImageAnnotationTypePtrOutput) ToImageAnnotationTypePtrOutputWithContext(ctx context.Context) ImageAnnotationTypePtrOutput {
	return o
}

func (o ImageAnnotationTypePtrOutput) Elem() ImageAnnotationTypeOutput {
	return o.ApplyT(func(v *ImageAnnotationType) ImageAnnotationType {
		if v != nil {
			return *v
		}
		var ret ImageAnnotationType
		return ret
	}).(ImageAnnotationTypeOutput)
}

func (o ImageAnnotationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ImageAnnotationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ImageAnnotationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ImageAnnotationTypeInput interface {
	pulumi.Input

	ToImageAnnotationTypeOutput() ImageAnnotationTypeOutput
	ToImageAnnotationTypeOutputWithContext(context.Context) ImageAnnotationTypeOutput
}

var imageAnnotationTypePtrType = reflect.TypeOf((**ImageAnnotationType)(nil)).Elem()

type ImageAnnotationTypePtrInput interface {
	pulumi.Input

	ToImageAnnotationTypePtrOutput() ImageAnnotationTypePtrOutput
	ToImageAnnotationTypePtrOutputWithContext(context.Context) ImageAnnotationTypePtrOutput
}

type imageAnnotationTypePtr string

func ImageAnnotationTypePtr(v string) ImageAnnotationTypePtrInput {
	return (*imageAnnotationTypePtr)(&v)
}

func (*imageAnnotationTypePtr) ElementType() reflect.Type {
	return imageAnnotationTypePtrType
}

func (in *imageAnnotationTypePtr) ToImageAnnotationTypePtrOutput() ImageAnnotationTypePtrOutput {
	return pulumi.ToOutput(in).(ImageAnnotationTypePtrOutput)
}

func (in *imageAnnotationTypePtr) ToImageAnnotationTypePtrOutputWithContext(ctx context.Context) ImageAnnotationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ImageAnnotationTypePtrOutput)
}

type JobType string

const (
	JobTypeCommand  = JobType("Command")
	JobTypeSweep    = JobType("Sweep")
	JobTypeLabeling = JobType("Labeling")
)

func (JobType) ElementType() reflect.Type {
	return reflect.TypeOf((*JobType)(nil)).Elem()
}

func (e JobType) ToJobTypeOutput() JobTypeOutput {
	return pulumi.ToOutput(e).(JobTypeOutput)
}

func (e JobType) ToJobTypeOutputWithContext(ctx context.Context) JobTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(JobTypeOutput)
}

func (e JobType) ToJobTypePtrOutput() JobTypePtrOutput {
	return e.ToJobTypePtrOutputWithContext(context.Background())
}

func (e JobType) ToJobTypePtrOutputWithContext(ctx context.Context) JobTypePtrOutput {
	return JobType(e).ToJobTypeOutputWithContext(ctx).ToJobTypePtrOutputWithContext(ctx)
}

func (e JobType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e JobType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e JobType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type JobTypeOutput struct{ *pulumi.OutputState }

func (JobTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobType)(nil)).Elem()
}

func (o JobTypeOutput) ToJobTypeOutput() JobTypeOutput {
	return o
}

func (o JobTypeOutput) ToJobTypeOutputWithContext(ctx context.Context) JobTypeOutput {
	return o
}

func (o JobTypeOutput) ToJobTypePtrOutput() JobTypePtrOutput {
	return o.ToJobTypePtrOutputWithContext(context.Background())
}

func (o JobTypeOutput) ToJobTypePtrOutputWithContext(ctx context.Context) JobTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobType) *JobType {
		return &v
	}).(JobTypePtrOutput)
}

func (o JobTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o JobTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e JobType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o JobTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o JobTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e JobType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type JobTypePtrOutput struct{ *pulumi.OutputState }

func (JobTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobType)(nil)).Elem()
}

func (o JobTypePtrOutput) ToJobTypePtrOutput() JobTypePtrOutput {
	return o
}

func (o JobTypePtrOutput) ToJobTypePtrOutputWithContext(ctx context.Context) JobTypePtrOutput {
	return o
}

func (o JobTypePtrOutput) Elem() JobTypeOutput {
	return o.ApplyT(func(v *JobType) JobType {
		if v != nil {
			return *v
		}
		var ret JobType
		return ret
	}).(JobTypeOutput)
}

func (o JobTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o JobTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *JobType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type JobTypeInput interface {
	pulumi.Input

	ToJobTypeOutput() JobTypeOutput
	ToJobTypeOutputWithContext(context.Context) JobTypeOutput
}

var jobTypePtrType = reflect.TypeOf((**JobType)(nil)).Elem()

type JobTypePtrInput interface {
	pulumi.Input

	ToJobTypePtrOutput() JobTypePtrOutput
	ToJobTypePtrOutputWithContext(context.Context) JobTypePtrOutput
}

type jobTypePtr string

func JobTypePtr(v string) JobTypePtrInput {
	return (*jobTypePtr)(&v)
}

func (*jobTypePtr) ElementType() reflect.Type {
	return jobTypePtrType
}

func (in *jobTypePtr) ToJobTypePtrOutput() JobTypePtrOutput {
	return pulumi.ToOutput(in).(JobTypePtrOutput)
}

func (in *jobTypePtr) ToJobTypePtrOutputWithContext(ctx context.Context) JobTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(JobTypePtrOutput)
}

type LinkedServiceLinkType string

const (
	LinkedServiceLinkTypeSynapse = LinkedServiceLinkType("Synapse")
)

func (LinkedServiceLinkType) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceLinkType)(nil)).Elem()
}

func (e LinkedServiceLinkType) ToLinkedServiceLinkTypeOutput() LinkedServiceLinkTypeOutput {
	return pulumi.ToOutput(e).(LinkedServiceLinkTypeOutput)
}

func (e LinkedServiceLinkType) ToLinkedServiceLinkTypeOutputWithContext(ctx context.Context) LinkedServiceLinkTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LinkedServiceLinkTypeOutput)
}

func (e LinkedServiceLinkType) ToLinkedServiceLinkTypePtrOutput() LinkedServiceLinkTypePtrOutput {
	return e.ToLinkedServiceLinkTypePtrOutputWithContext(context.Background())
}

func (e LinkedServiceLinkType) ToLinkedServiceLinkTypePtrOutputWithContext(ctx context.Context) LinkedServiceLinkTypePtrOutput {
	return LinkedServiceLinkType(e).ToLinkedServiceLinkTypeOutputWithContext(ctx).ToLinkedServiceLinkTypePtrOutputWithContext(ctx)
}

func (e LinkedServiceLinkType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LinkedServiceLinkType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LinkedServiceLinkType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LinkedServiceLinkType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LinkedServiceLinkTypeOutput struct{ *pulumi.OutputState }

func (LinkedServiceLinkTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceLinkType)(nil)).Elem()
}

func (o LinkedServiceLinkTypeOutput) ToLinkedServiceLinkTypeOutput() LinkedServiceLinkTypeOutput {
	return o
}

func (o LinkedServiceLinkTypeOutput) ToLinkedServiceLinkTypeOutputWithContext(ctx context.Context) LinkedServiceLinkTypeOutput {
	return o
}

func (o LinkedServiceLinkTypeOutput) ToLinkedServiceLinkTypePtrOutput() LinkedServiceLinkTypePtrOutput {
	return o.ToLinkedServiceLinkTypePtrOutputWithContext(context.Background())
}

func (o LinkedServiceLinkTypeOutput) ToLinkedServiceLinkTypePtrOutputWithContext(ctx context.Context) LinkedServiceLinkTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinkedServiceLinkType) *LinkedServiceLinkType {
		return &v
	}).(LinkedServiceLinkTypePtrOutput)
}

func (o LinkedServiceLinkTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LinkedServiceLinkTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LinkedServiceLinkType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LinkedServiceLinkTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LinkedServiceLinkTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LinkedServiceLinkType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LinkedServiceLinkTypePtrOutput struct{ *pulumi.OutputState }

func (LinkedServiceLinkTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceLinkType)(nil)).Elem()
}

func (o LinkedServiceLinkTypePtrOutput) ToLinkedServiceLinkTypePtrOutput() LinkedServiceLinkTypePtrOutput {
	return o
}

func (o LinkedServiceLinkTypePtrOutput) ToLinkedServiceLinkTypePtrOutputWithContext(ctx context.Context) LinkedServiceLinkTypePtrOutput {
	return o
}

func (o LinkedServiceLinkTypePtrOutput) Elem() LinkedServiceLinkTypeOutput {
	return o.ApplyT(func(v *LinkedServiceLinkType) LinkedServiceLinkType {
		if v != nil {
			return *v
		}
		var ret LinkedServiceLinkType
		return ret
	}).(LinkedServiceLinkTypeOutput)
}

func (o LinkedServiceLinkTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LinkedServiceLinkTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LinkedServiceLinkType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LinkedServiceLinkTypeInput interface {
	pulumi.Input

	ToLinkedServiceLinkTypeOutput() LinkedServiceLinkTypeOutput
	ToLinkedServiceLinkTypeOutputWithContext(context.Context) LinkedServiceLinkTypeOutput
}

var linkedServiceLinkTypePtrType = reflect.TypeOf((**LinkedServiceLinkType)(nil)).Elem()

type LinkedServiceLinkTypePtrInput interface {
	pulumi.Input

	ToLinkedServiceLinkTypePtrOutput() LinkedServiceLinkTypePtrOutput
	ToLinkedServiceLinkTypePtrOutputWithContext(context.Context) LinkedServiceLinkTypePtrOutput
}

type linkedServiceLinkTypePtr string

func LinkedServiceLinkTypePtr(v string) LinkedServiceLinkTypePtrInput {
	return (*linkedServiceLinkTypePtr)(&v)
}

func (*linkedServiceLinkTypePtr) ElementType() reflect.Type {
	return linkedServiceLinkTypePtrType
}

func (in *linkedServiceLinkTypePtr) ToLinkedServiceLinkTypePtrOutput() LinkedServiceLinkTypePtrOutput {
	return pulumi.ToOutput(in).(LinkedServiceLinkTypePtrOutput)
}

func (in *linkedServiceLinkTypePtr) ToLinkedServiceLinkTypePtrOutputWithContext(ctx context.Context) LinkedServiceLinkTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LinkedServiceLinkTypePtrOutput)
}

type MediaType string

const (
	MediaTypeImage = MediaType("Image")
	MediaTypeText  = MediaType("Text")
)

func (MediaType) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaType)(nil)).Elem()
}

func (e MediaType) ToMediaTypeOutput() MediaTypeOutput {
	return pulumi.ToOutput(e).(MediaTypeOutput)
}

func (e MediaType) ToMediaTypeOutputWithContext(ctx context.Context) MediaTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MediaTypeOutput)
}

func (e MediaType) ToMediaTypePtrOutput() MediaTypePtrOutput {
	return e.ToMediaTypePtrOutputWithContext(context.Background())
}

func (e MediaType) ToMediaTypePtrOutputWithContext(ctx context.Context) MediaTypePtrOutput {
	return MediaType(e).ToMediaTypeOutputWithContext(ctx).ToMediaTypePtrOutputWithContext(ctx)
}

func (e MediaType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MediaType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MediaType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MediaType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MediaTypeOutput struct{ *pulumi.OutputState }

func (MediaTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaType)(nil)).Elem()
}

func (o MediaTypeOutput) ToMediaTypeOutput() MediaTypeOutput {
	return o
}

func (o MediaTypeOutput) ToMediaTypeOutputWithContext(ctx context.Context) MediaTypeOutput {
	return o
}

func (o MediaTypeOutput) ToMediaTypePtrOutput() MediaTypePtrOutput {
	return o.ToMediaTypePtrOutputWithContext(context.Background())
}

func (o MediaTypeOutput) ToMediaTypePtrOutputWithContext(ctx context.Context) MediaTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MediaType) *MediaType {
		return &v
	}).(MediaTypePtrOutput)
}

func (o MediaTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MediaTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MediaType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MediaTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MediaTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MediaType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MediaTypePtrOutput struct{ *pulumi.OutputState }

func (MediaTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaType)(nil)).Elem()
}

func (o MediaTypePtrOutput) ToMediaTypePtrOutput() MediaTypePtrOutput {
	return o
}

func (o MediaTypePtrOutput) ToMediaTypePtrOutputWithContext(ctx context.Context) MediaTypePtrOutput {
	return o
}

func (o MediaTypePtrOutput) Elem() MediaTypeOutput {
	return o.ApplyT(func(v *MediaType) MediaType {
		if v != nil {
			return *v
		}
		var ret MediaType
		return ret
	}).(MediaTypeOutput)
}

func (o MediaTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MediaTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MediaType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MediaTypeInput interface {
	pulumi.Input

	ToMediaTypeOutput() MediaTypeOutput
	ToMediaTypeOutputWithContext(context.Context) MediaTypeOutput
}

var mediaTypePtrType = reflect.TypeOf((**MediaType)(nil)).Elem()

type MediaTypePtrInput interface {
	pulumi.Input

	ToMediaTypePtrOutput() MediaTypePtrOutput
	ToMediaTypePtrOutputWithContext(context.Context) MediaTypePtrOutput
}

type mediaTypePtr string

func MediaTypePtr(v string) MediaTypePtrInput {
	return (*mediaTypePtr)(&v)
}

func (*mediaTypePtr) ElementType() reflect.Type {
	return mediaTypePtrType
}

func (in *mediaTypePtr) ToMediaTypePtrOutput() MediaTypePtrOutput {
	return pulumi.ToOutput(in).(MediaTypePtrOutput)
}

func (in *mediaTypePtr) ToMediaTypePtrOutputWithContext(ctx context.Context) MediaTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MediaTypePtrOutput)
}

type OperatingSystemType string

const (
	OperatingSystemTypeLinux   = OperatingSystemType("Linux")
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

type OsType string

const (
	OsTypeLinux   = OsType("Linux")
	OsTypeWindows = OsType("Windows")
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

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusPending      = PrivateEndpointServiceConnectionStatus("Pending")
	PrivateEndpointServiceConnectionStatusApproved     = PrivateEndpointServiceConnectionStatus("Approved")
	PrivateEndpointServiceConnectionStatusRejected     = PrivateEndpointServiceConnectionStatus("Rejected")
	PrivateEndpointServiceConnectionStatusDisconnected = PrivateEndpointServiceConnectionStatus("Disconnected")
	PrivateEndpointServiceConnectionStatusTimeout      = PrivateEndpointServiceConnectionStatus("Timeout")
)

func (PrivateEndpointServiceConnectionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointServiceConnectionStatus)(nil)).Elem()
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusOutput() PrivateEndpointServiceConnectionStatusOutput {
	return pulumi.ToOutput(e).(PrivateEndpointServiceConnectionStatusOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrivateEndpointServiceConnectionStatusOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return e.ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(context.Background())
}

func (e PrivateEndpointServiceConnectionStatus) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return PrivateEndpointServiceConnectionStatus(e).ToPrivateEndpointServiceConnectionStatusOutputWithContext(ctx).ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx)
}

func (e PrivateEndpointServiceConnectionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateEndpointServiceConnectionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrivateEndpointServiceConnectionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrivateEndpointServiceConnectionStatusOutput struct{ *pulumi.OutputState }

func (PrivateEndpointServiceConnectionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointServiceConnectionStatus)(nil)).Elem()
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusOutput() PrivateEndpointServiceConnectionStatusOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return o.ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointServiceConnectionStatus) *PrivateEndpointServiceConnectionStatus {
		return &v
	}).(PrivateEndpointServiceConnectionStatusPtrOutput)
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateEndpointServiceConnectionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateEndpointServiceConnectionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointServiceConnectionStatusPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointServiceConnectionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointServiceConnectionStatus)(nil)).Elem()
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return o
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) Elem() PrivateEndpointServiceConnectionStatusOutput {
	return o.ApplyT(func(v *PrivateEndpointServiceConnectionStatus) PrivateEndpointServiceConnectionStatus {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointServiceConnectionStatus
		return ret
	}).(PrivateEndpointServiceConnectionStatusOutput)
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointServiceConnectionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrivateEndpointServiceConnectionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PrivateEndpointServiceConnectionStatusInput interface {
	pulumi.Input

	ToPrivateEndpointServiceConnectionStatusOutput() PrivateEndpointServiceConnectionStatusOutput
	ToPrivateEndpointServiceConnectionStatusOutputWithContext(context.Context) PrivateEndpointServiceConnectionStatusOutput
}

var privateEndpointServiceConnectionStatusPtrType = reflect.TypeOf((**PrivateEndpointServiceConnectionStatus)(nil)).Elem()

type PrivateEndpointServiceConnectionStatusPtrInput interface {
	pulumi.Input

	ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput
	ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(context.Context) PrivateEndpointServiceConnectionStatusPtrOutput
}

type privateEndpointServiceConnectionStatusPtr string

func PrivateEndpointServiceConnectionStatusPtr(v string) PrivateEndpointServiceConnectionStatusPtrInput {
	return (*privateEndpointServiceConnectionStatusPtr)(&v)
}

func (*privateEndpointServiceConnectionStatusPtr) ElementType() reflect.Type {
	return privateEndpointServiceConnectionStatusPtrType
}

func (in *privateEndpointServiceConnectionStatusPtr) ToPrivateEndpointServiceConnectionStatusPtrOutput() PrivateEndpointServiceConnectionStatusPtrOutput {
	return pulumi.ToOutput(in).(PrivateEndpointServiceConnectionStatusPtrOutput)
}

func (in *privateEndpointServiceConnectionStatusPtr) ToPrivateEndpointServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateEndpointServiceConnectionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrivateEndpointServiceConnectionStatusPtrOutput)
}

type ReferenceType string

const (
	ReferenceTypeId         = ReferenceType("Id")
	ReferenceTypeDataPath   = ReferenceType("DataPath")
	ReferenceTypeOutputPath = ReferenceType("OutputPath")
)

func (ReferenceType) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceType)(nil)).Elem()
}

func (e ReferenceType) ToReferenceTypeOutput() ReferenceTypeOutput {
	return pulumi.ToOutput(e).(ReferenceTypeOutput)
}

func (e ReferenceType) ToReferenceTypeOutputWithContext(ctx context.Context) ReferenceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ReferenceTypeOutput)
}

func (e ReferenceType) ToReferenceTypePtrOutput() ReferenceTypePtrOutput {
	return e.ToReferenceTypePtrOutputWithContext(context.Background())
}

func (e ReferenceType) ToReferenceTypePtrOutputWithContext(ctx context.Context) ReferenceTypePtrOutput {
	return ReferenceType(e).ToReferenceTypeOutputWithContext(ctx).ToReferenceTypePtrOutputWithContext(ctx)
}

func (e ReferenceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReferenceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReferenceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ReferenceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ReferenceTypeOutput struct{ *pulumi.OutputState }

func (ReferenceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceType)(nil)).Elem()
}

func (o ReferenceTypeOutput) ToReferenceTypeOutput() ReferenceTypeOutput {
	return o
}

func (o ReferenceTypeOutput) ToReferenceTypeOutputWithContext(ctx context.Context) ReferenceTypeOutput {
	return o
}

func (o ReferenceTypeOutput) ToReferenceTypePtrOutput() ReferenceTypePtrOutput {
	return o.ToReferenceTypePtrOutputWithContext(context.Background())
}

func (o ReferenceTypeOutput) ToReferenceTypePtrOutputWithContext(ctx context.Context) ReferenceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReferenceType) *ReferenceType {
		return &v
	}).(ReferenceTypePtrOutput)
}

func (o ReferenceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ReferenceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReferenceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ReferenceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReferenceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReferenceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ReferenceTypePtrOutput struct{ *pulumi.OutputState }

func (ReferenceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReferenceType)(nil)).Elem()
}

func (o ReferenceTypePtrOutput) ToReferenceTypePtrOutput() ReferenceTypePtrOutput {
	return o
}

func (o ReferenceTypePtrOutput) ToReferenceTypePtrOutputWithContext(ctx context.Context) ReferenceTypePtrOutput {
	return o
}

func (o ReferenceTypePtrOutput) Elem() ReferenceTypeOutput {
	return o.ApplyT(func(v *ReferenceType) ReferenceType {
		if v != nil {
			return *v
		}
		var ret ReferenceType
		return ret
	}).(ReferenceTypeOutput)
}

func (o ReferenceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReferenceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ReferenceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ReferenceTypeInput interface {
	pulumi.Input

	ToReferenceTypeOutput() ReferenceTypeOutput
	ToReferenceTypeOutputWithContext(context.Context) ReferenceTypeOutput
}

var referenceTypePtrType = reflect.TypeOf((**ReferenceType)(nil)).Elem()

type ReferenceTypePtrInput interface {
	pulumi.Input

	ToReferenceTypePtrOutput() ReferenceTypePtrOutput
	ToReferenceTypePtrOutputWithContext(context.Context) ReferenceTypePtrOutput
}

type referenceTypePtr string

func ReferenceTypePtr(v string) ReferenceTypePtrInput {
	return (*referenceTypePtr)(&v)
}

func (*referenceTypePtr) ElementType() reflect.Type {
	return referenceTypePtrType
}

func (in *referenceTypePtr) ToReferenceTypePtrOutput() ReferenceTypePtrOutput {
	return pulumi.ToOutput(in).(ReferenceTypePtrOutput)
}

func (in *referenceTypePtr) ToReferenceTypePtrOutputWithContext(ctx context.Context) ReferenceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ReferenceTypePtrOutput)
}

type RemoteLoginPortPublicAccess string

const (
	RemoteLoginPortPublicAccessEnabled      = RemoteLoginPortPublicAccess("Enabled")
	RemoteLoginPortPublicAccessDisabled     = RemoteLoginPortPublicAccess("Disabled")
	RemoteLoginPortPublicAccessNotSpecified = RemoteLoginPortPublicAccess("NotSpecified")
)

func (RemoteLoginPortPublicAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoteLoginPortPublicAccess)(nil)).Elem()
}

func (e RemoteLoginPortPublicAccess) ToRemoteLoginPortPublicAccessOutput() RemoteLoginPortPublicAccessOutput {
	return pulumi.ToOutput(e).(RemoteLoginPortPublicAccessOutput)
}

func (e RemoteLoginPortPublicAccess) ToRemoteLoginPortPublicAccessOutputWithContext(ctx context.Context) RemoteLoginPortPublicAccessOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RemoteLoginPortPublicAccessOutput)
}

func (e RemoteLoginPortPublicAccess) ToRemoteLoginPortPublicAccessPtrOutput() RemoteLoginPortPublicAccessPtrOutput {
	return e.ToRemoteLoginPortPublicAccessPtrOutputWithContext(context.Background())
}

func (e RemoteLoginPortPublicAccess) ToRemoteLoginPortPublicAccessPtrOutputWithContext(ctx context.Context) RemoteLoginPortPublicAccessPtrOutput {
	return RemoteLoginPortPublicAccess(e).ToRemoteLoginPortPublicAccessOutputWithContext(ctx).ToRemoteLoginPortPublicAccessPtrOutputWithContext(ctx)
}

func (e RemoteLoginPortPublicAccess) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RemoteLoginPortPublicAccess) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RemoteLoginPortPublicAccess) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RemoteLoginPortPublicAccess) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RemoteLoginPortPublicAccessOutput struct{ *pulumi.OutputState }

func (RemoteLoginPortPublicAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoteLoginPortPublicAccess)(nil)).Elem()
}

func (o RemoteLoginPortPublicAccessOutput) ToRemoteLoginPortPublicAccessOutput() RemoteLoginPortPublicAccessOutput {
	return o
}

func (o RemoteLoginPortPublicAccessOutput) ToRemoteLoginPortPublicAccessOutputWithContext(ctx context.Context) RemoteLoginPortPublicAccessOutput {
	return o
}

func (o RemoteLoginPortPublicAccessOutput) ToRemoteLoginPortPublicAccessPtrOutput() RemoteLoginPortPublicAccessPtrOutput {
	return o.ToRemoteLoginPortPublicAccessPtrOutputWithContext(context.Background())
}

func (o RemoteLoginPortPublicAccessOutput) ToRemoteLoginPortPublicAccessPtrOutputWithContext(ctx context.Context) RemoteLoginPortPublicAccessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RemoteLoginPortPublicAccess) *RemoteLoginPortPublicAccess {
		return &v
	}).(RemoteLoginPortPublicAccessPtrOutput)
}

func (o RemoteLoginPortPublicAccessOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RemoteLoginPortPublicAccessOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RemoteLoginPortPublicAccess) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RemoteLoginPortPublicAccessOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RemoteLoginPortPublicAccessOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RemoteLoginPortPublicAccess) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RemoteLoginPortPublicAccessPtrOutput struct{ *pulumi.OutputState }

func (RemoteLoginPortPublicAccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemoteLoginPortPublicAccess)(nil)).Elem()
}

func (o RemoteLoginPortPublicAccessPtrOutput) ToRemoteLoginPortPublicAccessPtrOutput() RemoteLoginPortPublicAccessPtrOutput {
	return o
}

func (o RemoteLoginPortPublicAccessPtrOutput) ToRemoteLoginPortPublicAccessPtrOutputWithContext(ctx context.Context) RemoteLoginPortPublicAccessPtrOutput {
	return o
}

func (o RemoteLoginPortPublicAccessPtrOutput) Elem() RemoteLoginPortPublicAccessOutput {
	return o.ApplyT(func(v *RemoteLoginPortPublicAccess) RemoteLoginPortPublicAccess {
		if v != nil {
			return *v
		}
		var ret RemoteLoginPortPublicAccess
		return ret
	}).(RemoteLoginPortPublicAccessOutput)
}

func (o RemoteLoginPortPublicAccessPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RemoteLoginPortPublicAccessPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RemoteLoginPortPublicAccess) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RemoteLoginPortPublicAccessInput interface {
	pulumi.Input

	ToRemoteLoginPortPublicAccessOutput() RemoteLoginPortPublicAccessOutput
	ToRemoteLoginPortPublicAccessOutputWithContext(context.Context) RemoteLoginPortPublicAccessOutput
}

var remoteLoginPortPublicAccessPtrType = reflect.TypeOf((**RemoteLoginPortPublicAccess)(nil)).Elem()

type RemoteLoginPortPublicAccessPtrInput interface {
	pulumi.Input

	ToRemoteLoginPortPublicAccessPtrOutput() RemoteLoginPortPublicAccessPtrOutput
	ToRemoteLoginPortPublicAccessPtrOutputWithContext(context.Context) RemoteLoginPortPublicAccessPtrOutput
}

type remoteLoginPortPublicAccessPtr string

func RemoteLoginPortPublicAccessPtr(v string) RemoteLoginPortPublicAccessPtrInput {
	return (*remoteLoginPortPublicAccessPtr)(&v)
}

func (*remoteLoginPortPublicAccessPtr) ElementType() reflect.Type {
	return remoteLoginPortPublicAccessPtrType
}

func (in *remoteLoginPortPublicAccessPtr) ToRemoteLoginPortPublicAccessPtrOutput() RemoteLoginPortPublicAccessPtrOutput {
	return pulumi.ToOutput(in).(RemoteLoginPortPublicAccessPtrOutput)
}

func (in *remoteLoginPortPublicAccessPtr) ToRemoteLoginPortPublicAccessPtrOutputWithContext(ctx context.Context) RemoteLoginPortPublicAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RemoteLoginPortPublicAccessPtrOutput)
}

type ResourceIdentityAssignment string

const (
	ResourceIdentityAssignmentSystemAssigned               = ResourceIdentityAssignment("SystemAssigned")
	ResourceIdentityAssignmentUserAssigned                 = ResourceIdentityAssignment("UserAssigned")
	ResourceIdentityAssignment_SystemAssigned_UserAssigned = ResourceIdentityAssignment("SystemAssigned,UserAssigned")
	ResourceIdentityAssignmentNone                         = ResourceIdentityAssignment("None")
)

func (ResourceIdentityAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityAssignment)(nil)).Elem()
}

func (e ResourceIdentityAssignment) ToResourceIdentityAssignmentOutput() ResourceIdentityAssignmentOutput {
	return pulumi.ToOutput(e).(ResourceIdentityAssignmentOutput)
}

func (e ResourceIdentityAssignment) ToResourceIdentityAssignmentOutputWithContext(ctx context.Context) ResourceIdentityAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityAssignmentOutput)
}

func (e ResourceIdentityAssignment) ToResourceIdentityAssignmentPtrOutput() ResourceIdentityAssignmentPtrOutput {
	return e.ToResourceIdentityAssignmentPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityAssignment) ToResourceIdentityAssignmentPtrOutputWithContext(ctx context.Context) ResourceIdentityAssignmentPtrOutput {
	return ResourceIdentityAssignment(e).ToResourceIdentityAssignmentOutputWithContext(ctx).ToResourceIdentityAssignmentPtrOutputWithContext(ctx)
}

func (e ResourceIdentityAssignment) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityAssignment) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityAssignment) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityAssignment) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityAssignmentOutput struct{ *pulumi.OutputState }

func (ResourceIdentityAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityAssignment)(nil)).Elem()
}

func (o ResourceIdentityAssignmentOutput) ToResourceIdentityAssignmentOutput() ResourceIdentityAssignmentOutput {
	return o
}

func (o ResourceIdentityAssignmentOutput) ToResourceIdentityAssignmentOutputWithContext(ctx context.Context) ResourceIdentityAssignmentOutput {
	return o
}

func (o ResourceIdentityAssignmentOutput) ToResourceIdentityAssignmentPtrOutput() ResourceIdentityAssignmentPtrOutput {
	return o.ToResourceIdentityAssignmentPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityAssignmentOutput) ToResourceIdentityAssignmentPtrOutputWithContext(ctx context.Context) ResourceIdentityAssignmentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityAssignment) *ResourceIdentityAssignment {
		return &v
	}).(ResourceIdentityAssignmentPtrOutput)
}

func (o ResourceIdentityAssignmentOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityAssignmentOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityAssignment) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityAssignmentOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityAssignmentOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityAssignment) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityAssignmentPtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityAssignmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityAssignment)(nil)).Elem()
}

func (o ResourceIdentityAssignmentPtrOutput) ToResourceIdentityAssignmentPtrOutput() ResourceIdentityAssignmentPtrOutput {
	return o
}

func (o ResourceIdentityAssignmentPtrOutput) ToResourceIdentityAssignmentPtrOutputWithContext(ctx context.Context) ResourceIdentityAssignmentPtrOutput {
	return o
}

func (o ResourceIdentityAssignmentPtrOutput) Elem() ResourceIdentityAssignmentOutput {
	return o.ApplyT(func(v *ResourceIdentityAssignment) ResourceIdentityAssignment {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityAssignment
		return ret
	}).(ResourceIdentityAssignmentOutput)
}

func (o ResourceIdentityAssignmentPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityAssignmentPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityAssignment) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ResourceIdentityAssignmentInput interface {
	pulumi.Input

	ToResourceIdentityAssignmentOutput() ResourceIdentityAssignmentOutput
	ToResourceIdentityAssignmentOutputWithContext(context.Context) ResourceIdentityAssignmentOutput
}

var resourceIdentityAssignmentPtrType = reflect.TypeOf((**ResourceIdentityAssignment)(nil)).Elem()

type ResourceIdentityAssignmentPtrInput interface {
	pulumi.Input

	ToResourceIdentityAssignmentPtrOutput() ResourceIdentityAssignmentPtrOutput
	ToResourceIdentityAssignmentPtrOutputWithContext(context.Context) ResourceIdentityAssignmentPtrOutput
}

type resourceIdentityAssignmentPtr string

func ResourceIdentityAssignmentPtr(v string) ResourceIdentityAssignmentPtrInput {
	return (*resourceIdentityAssignmentPtr)(&v)
}

func (*resourceIdentityAssignmentPtr) ElementType() reflect.Type {
	return resourceIdentityAssignmentPtrType
}

func (in *resourceIdentityAssignmentPtr) ToResourceIdentityAssignmentPtrOutput() ResourceIdentityAssignmentPtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityAssignmentPtrOutput)
}

func (in *resourceIdentityAssignmentPtr) ToResourceIdentityAssignmentPtrOutputWithContext(ctx context.Context) ResourceIdentityAssignmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityAssignmentPtrOutput)
}

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned,UserAssigned")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityTypeNone                         = ResourceIdentityType("None")
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

type SamplingAlgorithm string

const (
	SamplingAlgorithmGrid     = SamplingAlgorithm("Grid")
	SamplingAlgorithmRandom   = SamplingAlgorithm("Random")
	SamplingAlgorithmBayesian = SamplingAlgorithm("Bayesian")
)

func (SamplingAlgorithm) ElementType() reflect.Type {
	return reflect.TypeOf((*SamplingAlgorithm)(nil)).Elem()
}

func (e SamplingAlgorithm) ToSamplingAlgorithmOutput() SamplingAlgorithmOutput {
	return pulumi.ToOutput(e).(SamplingAlgorithmOutput)
}

func (e SamplingAlgorithm) ToSamplingAlgorithmOutputWithContext(ctx context.Context) SamplingAlgorithmOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SamplingAlgorithmOutput)
}

func (e SamplingAlgorithm) ToSamplingAlgorithmPtrOutput() SamplingAlgorithmPtrOutput {
	return e.ToSamplingAlgorithmPtrOutputWithContext(context.Background())
}

func (e SamplingAlgorithm) ToSamplingAlgorithmPtrOutputWithContext(ctx context.Context) SamplingAlgorithmPtrOutput {
	return SamplingAlgorithm(e).ToSamplingAlgorithmOutputWithContext(ctx).ToSamplingAlgorithmPtrOutputWithContext(ctx)
}

func (e SamplingAlgorithm) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SamplingAlgorithm) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SamplingAlgorithm) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SamplingAlgorithm) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SamplingAlgorithmOutput struct{ *pulumi.OutputState }

func (SamplingAlgorithmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SamplingAlgorithm)(nil)).Elem()
}

func (o SamplingAlgorithmOutput) ToSamplingAlgorithmOutput() SamplingAlgorithmOutput {
	return o
}

func (o SamplingAlgorithmOutput) ToSamplingAlgorithmOutputWithContext(ctx context.Context) SamplingAlgorithmOutput {
	return o
}

func (o SamplingAlgorithmOutput) ToSamplingAlgorithmPtrOutput() SamplingAlgorithmPtrOutput {
	return o.ToSamplingAlgorithmPtrOutputWithContext(context.Background())
}

func (o SamplingAlgorithmOutput) ToSamplingAlgorithmPtrOutputWithContext(ctx context.Context) SamplingAlgorithmPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SamplingAlgorithm) *SamplingAlgorithm {
		return &v
	}).(SamplingAlgorithmPtrOutput)
}

func (o SamplingAlgorithmOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SamplingAlgorithmOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SamplingAlgorithm) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SamplingAlgorithmOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SamplingAlgorithmOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SamplingAlgorithm) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SamplingAlgorithmPtrOutput struct{ *pulumi.OutputState }

func (SamplingAlgorithmPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SamplingAlgorithm)(nil)).Elem()
}

func (o SamplingAlgorithmPtrOutput) ToSamplingAlgorithmPtrOutput() SamplingAlgorithmPtrOutput {
	return o
}

func (o SamplingAlgorithmPtrOutput) ToSamplingAlgorithmPtrOutputWithContext(ctx context.Context) SamplingAlgorithmPtrOutput {
	return o
}

func (o SamplingAlgorithmPtrOutput) Elem() SamplingAlgorithmOutput {
	return o.ApplyT(func(v *SamplingAlgorithm) SamplingAlgorithm {
		if v != nil {
			return *v
		}
		var ret SamplingAlgorithm
		return ret
	}).(SamplingAlgorithmOutput)
}

func (o SamplingAlgorithmPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SamplingAlgorithmPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SamplingAlgorithm) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SamplingAlgorithmInput interface {
	pulumi.Input

	ToSamplingAlgorithmOutput() SamplingAlgorithmOutput
	ToSamplingAlgorithmOutputWithContext(context.Context) SamplingAlgorithmOutput
}

var samplingAlgorithmPtrType = reflect.TypeOf((**SamplingAlgorithm)(nil)).Elem()

type SamplingAlgorithmPtrInput interface {
	pulumi.Input

	ToSamplingAlgorithmPtrOutput() SamplingAlgorithmPtrOutput
	ToSamplingAlgorithmPtrOutputWithContext(context.Context) SamplingAlgorithmPtrOutput
}

type samplingAlgorithmPtr string

func SamplingAlgorithmPtr(v string) SamplingAlgorithmPtrInput {
	return (*samplingAlgorithmPtr)(&v)
}

func (*samplingAlgorithmPtr) ElementType() reflect.Type {
	return samplingAlgorithmPtrType
}

func (in *samplingAlgorithmPtr) ToSamplingAlgorithmPtrOutput() SamplingAlgorithmPtrOutput {
	return pulumi.ToOutput(in).(SamplingAlgorithmPtrOutput)
}

func (in *samplingAlgorithmPtr) ToSamplingAlgorithmPtrOutputWithContext(ctx context.Context) SamplingAlgorithmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SamplingAlgorithmPtrOutput)
}

type ScaleType string

const (
	ScaleTypeAuto   = ScaleType("Auto")
	ScaleTypeManual = ScaleType("Manual")
)

func (ScaleType) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleType)(nil)).Elem()
}

func (e ScaleType) ToScaleTypeOutput() ScaleTypeOutput {
	return pulumi.ToOutput(e).(ScaleTypeOutput)
}

func (e ScaleType) ToScaleTypeOutputWithContext(ctx context.Context) ScaleTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScaleTypeOutput)
}

func (e ScaleType) ToScaleTypePtrOutput() ScaleTypePtrOutput {
	return e.ToScaleTypePtrOutputWithContext(context.Background())
}

func (e ScaleType) ToScaleTypePtrOutputWithContext(ctx context.Context) ScaleTypePtrOutput {
	return ScaleType(e).ToScaleTypeOutputWithContext(ctx).ToScaleTypePtrOutputWithContext(ctx)
}

func (e ScaleType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScaleType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScaleType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScaleType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScaleTypeOutput struct{ *pulumi.OutputState }

func (ScaleTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleType)(nil)).Elem()
}

func (o ScaleTypeOutput) ToScaleTypeOutput() ScaleTypeOutput {
	return o
}

func (o ScaleTypeOutput) ToScaleTypeOutputWithContext(ctx context.Context) ScaleTypeOutput {
	return o
}

func (o ScaleTypeOutput) ToScaleTypePtrOutput() ScaleTypePtrOutput {
	return o.ToScaleTypePtrOutputWithContext(context.Background())
}

func (o ScaleTypeOutput) ToScaleTypePtrOutputWithContext(ctx context.Context) ScaleTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScaleType) *ScaleType {
		return &v
	}).(ScaleTypePtrOutput)
}

func (o ScaleTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScaleTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScaleType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScaleTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScaleTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScaleType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScaleTypePtrOutput struct{ *pulumi.OutputState }

func (ScaleTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleType)(nil)).Elem()
}

func (o ScaleTypePtrOutput) ToScaleTypePtrOutput() ScaleTypePtrOutput {
	return o
}

func (o ScaleTypePtrOutput) ToScaleTypePtrOutputWithContext(ctx context.Context) ScaleTypePtrOutput {
	return o
}

func (o ScaleTypePtrOutput) Elem() ScaleTypeOutput {
	return o.ApplyT(func(v *ScaleType) ScaleType {
		if v != nil {
			return *v
		}
		var ret ScaleType
		return ret
	}).(ScaleTypeOutput)
}

func (o ScaleTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScaleTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScaleType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ScaleTypeInput interface {
	pulumi.Input

	ToScaleTypeOutput() ScaleTypeOutput
	ToScaleTypeOutputWithContext(context.Context) ScaleTypeOutput
}

var scaleTypePtrType = reflect.TypeOf((**ScaleType)(nil)).Elem()

type ScaleTypePtrInput interface {
	pulumi.Input

	ToScaleTypePtrOutput() ScaleTypePtrOutput
	ToScaleTypePtrOutputWithContext(context.Context) ScaleTypePtrOutput
}

type scaleTypePtr string

func ScaleTypePtr(v string) ScaleTypePtrInput {
	return (*scaleTypePtr)(&v)
}

func (*scaleTypePtr) ElementType() reflect.Type {
	return scaleTypePtrType
}

func (in *scaleTypePtr) ToScaleTypePtrOutput() ScaleTypePtrOutput {
	return pulumi.ToOutput(in).(ScaleTypePtrOutput)
}

func (in *scaleTypePtr) ToScaleTypePtrOutputWithContext(ctx context.Context) ScaleTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScaleTypePtrOutput)
}

type SourceType string

const (
	SourceType_Delimited_files  = SourceType("delimited_files")
	SourceType_Json_lines_files = SourceType("json_lines_files")
	SourceType_Parquet_files    = SourceType("parquet_files")
)

func (SourceType) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceType)(nil)).Elem()
}

func (e SourceType) ToSourceTypeOutput() SourceTypeOutput {
	return pulumi.ToOutput(e).(SourceTypeOutput)
}

func (e SourceType) ToSourceTypeOutputWithContext(ctx context.Context) SourceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SourceTypeOutput)
}

func (e SourceType) ToSourceTypePtrOutput() SourceTypePtrOutput {
	return e.ToSourceTypePtrOutputWithContext(context.Background())
}

func (e SourceType) ToSourceTypePtrOutputWithContext(ctx context.Context) SourceTypePtrOutput {
	return SourceType(e).ToSourceTypeOutputWithContext(ctx).ToSourceTypePtrOutputWithContext(ctx)
}

func (e SourceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SourceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SourceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SourceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SourceTypeOutput struct{ *pulumi.OutputState }

func (SourceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceType)(nil)).Elem()
}

func (o SourceTypeOutput) ToSourceTypeOutput() SourceTypeOutput {
	return o
}

func (o SourceTypeOutput) ToSourceTypeOutputWithContext(ctx context.Context) SourceTypeOutput {
	return o
}

func (o SourceTypeOutput) ToSourceTypePtrOutput() SourceTypePtrOutput {
	return o.ToSourceTypePtrOutputWithContext(context.Background())
}

func (o SourceTypeOutput) ToSourceTypePtrOutputWithContext(ctx context.Context) SourceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceType) *SourceType {
		return &v
	}).(SourceTypePtrOutput)
}

func (o SourceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SourceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SourceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SourceTypePtrOutput struct{ *pulumi.OutputState }

func (SourceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceType)(nil)).Elem()
}

func (o SourceTypePtrOutput) ToSourceTypePtrOutput() SourceTypePtrOutput {
	return o
}

func (o SourceTypePtrOutput) ToSourceTypePtrOutputWithContext(ctx context.Context) SourceTypePtrOutput {
	return o
}

func (o SourceTypePtrOutput) Elem() SourceTypeOutput {
	return o.ApplyT(func(v *SourceType) SourceType {
		if v != nil {
			return *v
		}
		var ret SourceType
		return ret
	}).(SourceTypeOutput)
}

func (o SourceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SourceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SourceTypeInput interface {
	pulumi.Input

	ToSourceTypeOutput() SourceTypeOutput
	ToSourceTypeOutputWithContext(context.Context) SourceTypeOutput
}

var sourceTypePtrType = reflect.TypeOf((**SourceType)(nil)).Elem()

type SourceTypePtrInput interface {
	pulumi.Input

	ToSourceTypePtrOutput() SourceTypePtrOutput
	ToSourceTypePtrOutputWithContext(context.Context) SourceTypePtrOutput
}

type sourceTypePtr string

func SourceTypePtr(v string) SourceTypePtrInput {
	return (*sourceTypePtr)(&v)
}

func (*sourceTypePtr) ElementType() reflect.Type {
	return sourceTypePtrType
}

func (in *sourceTypePtr) ToSourceTypePtrOutput() SourceTypePtrOutput {
	return pulumi.ToOutput(in).(SourceTypePtrOutput)
}

func (in *sourceTypePtr) ToSourceTypePtrOutputWithContext(ctx context.Context) SourceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SourceTypePtrOutput)
}

type SshPublicAccess string

const (
	SshPublicAccessEnabled  = SshPublicAccess("Enabled")
	SshPublicAccessDisabled = SshPublicAccess("Disabled")
)

func (SshPublicAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*SshPublicAccess)(nil)).Elem()
}

func (e SshPublicAccess) ToSshPublicAccessOutput() SshPublicAccessOutput {
	return pulumi.ToOutput(e).(SshPublicAccessOutput)
}

func (e SshPublicAccess) ToSshPublicAccessOutputWithContext(ctx context.Context) SshPublicAccessOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SshPublicAccessOutput)
}

func (e SshPublicAccess) ToSshPublicAccessPtrOutput() SshPublicAccessPtrOutput {
	return e.ToSshPublicAccessPtrOutputWithContext(context.Background())
}

func (e SshPublicAccess) ToSshPublicAccessPtrOutputWithContext(ctx context.Context) SshPublicAccessPtrOutput {
	return SshPublicAccess(e).ToSshPublicAccessOutputWithContext(ctx).ToSshPublicAccessPtrOutputWithContext(ctx)
}

func (e SshPublicAccess) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SshPublicAccess) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SshPublicAccess) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SshPublicAccess) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SshPublicAccessOutput struct{ *pulumi.OutputState }

func (SshPublicAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshPublicAccess)(nil)).Elem()
}

func (o SshPublicAccessOutput) ToSshPublicAccessOutput() SshPublicAccessOutput {
	return o
}

func (o SshPublicAccessOutput) ToSshPublicAccessOutputWithContext(ctx context.Context) SshPublicAccessOutput {
	return o
}

func (o SshPublicAccessOutput) ToSshPublicAccessPtrOutput() SshPublicAccessPtrOutput {
	return o.ToSshPublicAccessPtrOutputWithContext(context.Background())
}

func (o SshPublicAccessOutput) ToSshPublicAccessPtrOutputWithContext(ctx context.Context) SshPublicAccessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SshPublicAccess) *SshPublicAccess {
		return &v
	}).(SshPublicAccessPtrOutput)
}

func (o SshPublicAccessOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SshPublicAccessOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SshPublicAccess) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SshPublicAccessOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SshPublicAccessOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SshPublicAccess) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SshPublicAccessPtrOutput struct{ *pulumi.OutputState }

func (SshPublicAccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SshPublicAccess)(nil)).Elem()
}

func (o SshPublicAccessPtrOutput) ToSshPublicAccessPtrOutput() SshPublicAccessPtrOutput {
	return o
}

func (o SshPublicAccessPtrOutput) ToSshPublicAccessPtrOutputWithContext(ctx context.Context) SshPublicAccessPtrOutput {
	return o
}

func (o SshPublicAccessPtrOutput) Elem() SshPublicAccessOutput {
	return o.ApplyT(func(v *SshPublicAccess) SshPublicAccess {
		if v != nil {
			return *v
		}
		var ret SshPublicAccess
		return ret
	}).(SshPublicAccessOutput)
}

func (o SshPublicAccessPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SshPublicAccessPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SshPublicAccess) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SshPublicAccessInput interface {
	pulumi.Input

	ToSshPublicAccessOutput() SshPublicAccessOutput
	ToSshPublicAccessOutputWithContext(context.Context) SshPublicAccessOutput
}

var sshPublicAccessPtrType = reflect.TypeOf((**SshPublicAccess)(nil)).Elem()

type SshPublicAccessPtrInput interface {
	pulumi.Input

	ToSshPublicAccessPtrOutput() SshPublicAccessPtrOutput
	ToSshPublicAccessPtrOutputWithContext(context.Context) SshPublicAccessPtrOutput
}

type sshPublicAccessPtr string

func SshPublicAccessPtr(v string) SshPublicAccessPtrInput {
	return (*sshPublicAccessPtr)(&v)
}

func (*sshPublicAccessPtr) ElementType() reflect.Type {
	return sshPublicAccessPtrType
}

func (in *sshPublicAccessPtr) ToSshPublicAccessPtrOutput() SshPublicAccessPtrOutput {
	return pulumi.ToOutput(in).(SshPublicAccessPtrOutput)
}

func (in *sshPublicAccessPtr) ToSshPublicAccessPtrOutputWithContext(ctx context.Context) SshPublicAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SshPublicAccessPtrOutput)
}

type ValueFormat string

const (
	ValueFormatJSON = ValueFormat("JSON")
)

func (ValueFormat) ElementType() reflect.Type {
	return reflect.TypeOf((*ValueFormat)(nil)).Elem()
}

func (e ValueFormat) ToValueFormatOutput() ValueFormatOutput {
	return pulumi.ToOutput(e).(ValueFormatOutput)
}

func (e ValueFormat) ToValueFormatOutputWithContext(ctx context.Context) ValueFormatOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ValueFormatOutput)
}

func (e ValueFormat) ToValueFormatPtrOutput() ValueFormatPtrOutput {
	return e.ToValueFormatPtrOutputWithContext(context.Background())
}

func (e ValueFormat) ToValueFormatPtrOutputWithContext(ctx context.Context) ValueFormatPtrOutput {
	return ValueFormat(e).ToValueFormatOutputWithContext(ctx).ToValueFormatPtrOutputWithContext(ctx)
}

func (e ValueFormat) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ValueFormat) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ValueFormat) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ValueFormat) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ValueFormatOutput struct{ *pulumi.OutputState }

func (ValueFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ValueFormat)(nil)).Elem()
}

func (o ValueFormatOutput) ToValueFormatOutput() ValueFormatOutput {
	return o
}

func (o ValueFormatOutput) ToValueFormatOutputWithContext(ctx context.Context) ValueFormatOutput {
	return o
}

func (o ValueFormatOutput) ToValueFormatPtrOutput() ValueFormatPtrOutput {
	return o.ToValueFormatPtrOutputWithContext(context.Background())
}

func (o ValueFormatOutput) ToValueFormatPtrOutputWithContext(ctx context.Context) ValueFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ValueFormat) *ValueFormat {
		return &v
	}).(ValueFormatPtrOutput)
}

func (o ValueFormatOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ValueFormatOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ValueFormat) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ValueFormatOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ValueFormatOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ValueFormat) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ValueFormatPtrOutput struct{ *pulumi.OutputState }

func (ValueFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ValueFormat)(nil)).Elem()
}

func (o ValueFormatPtrOutput) ToValueFormatPtrOutput() ValueFormatPtrOutput {
	return o
}

func (o ValueFormatPtrOutput) ToValueFormatPtrOutputWithContext(ctx context.Context) ValueFormatPtrOutput {
	return o
}

func (o ValueFormatPtrOutput) Elem() ValueFormatOutput {
	return o.ApplyT(func(v *ValueFormat) ValueFormat {
		if v != nil {
			return *v
		}
		var ret ValueFormat
		return ret
	}).(ValueFormatOutput)
}

func (o ValueFormatPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ValueFormatPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ValueFormat) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ValueFormatInput interface {
	pulumi.Input

	ToValueFormatOutput() ValueFormatOutput
	ToValueFormatOutputWithContext(context.Context) ValueFormatOutput
}

var valueFormatPtrType = reflect.TypeOf((**ValueFormat)(nil)).Elem()

type ValueFormatPtrInput interface {
	pulumi.Input

	ToValueFormatPtrOutput() ValueFormatPtrOutput
	ToValueFormatPtrOutputWithContext(context.Context) ValueFormatPtrOutput
}

type valueFormatPtr string

func ValueFormatPtr(v string) ValueFormatPtrInput {
	return (*valueFormatPtr)(&v)
}

func (*valueFormatPtr) ElementType() reflect.Type {
	return valueFormatPtrType
}

func (in *valueFormatPtr) ToValueFormatPtrOutput() ValueFormatPtrOutput {
	return pulumi.ToOutput(in).(ValueFormatPtrOutput)
}

func (in *valueFormatPtr) ToValueFormatPtrOutputWithContext(ctx context.Context) ValueFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ValueFormatPtrOutput)
}

type VariantType string

const (
	VariantTypeControl   = VariantType("Control")
	VariantTypeTreatment = VariantType("Treatment")
)

func (VariantType) ElementType() reflect.Type {
	return reflect.TypeOf((*VariantType)(nil)).Elem()
}

func (e VariantType) ToVariantTypeOutput() VariantTypeOutput {
	return pulumi.ToOutput(e).(VariantTypeOutput)
}

func (e VariantType) ToVariantTypeOutputWithContext(ctx context.Context) VariantTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VariantTypeOutput)
}

func (e VariantType) ToVariantTypePtrOutput() VariantTypePtrOutput {
	return e.ToVariantTypePtrOutputWithContext(context.Background())
}

func (e VariantType) ToVariantTypePtrOutputWithContext(ctx context.Context) VariantTypePtrOutput {
	return VariantType(e).ToVariantTypeOutputWithContext(ctx).ToVariantTypePtrOutputWithContext(ctx)
}

func (e VariantType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VariantType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VariantType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VariantType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VariantTypeOutput struct{ *pulumi.OutputState }

func (VariantTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VariantType)(nil)).Elem()
}

func (o VariantTypeOutput) ToVariantTypeOutput() VariantTypeOutput {
	return o
}

func (o VariantTypeOutput) ToVariantTypeOutputWithContext(ctx context.Context) VariantTypeOutput {
	return o
}

func (o VariantTypeOutput) ToVariantTypePtrOutput() VariantTypePtrOutput {
	return o.ToVariantTypePtrOutputWithContext(context.Background())
}

func (o VariantTypeOutput) ToVariantTypePtrOutputWithContext(ctx context.Context) VariantTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VariantType) *VariantType {
		return &v
	}).(VariantTypePtrOutput)
}

func (o VariantTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VariantTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VariantType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VariantTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VariantTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VariantType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VariantTypePtrOutput struct{ *pulumi.OutputState }

func (VariantTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VariantType)(nil)).Elem()
}

func (o VariantTypePtrOutput) ToVariantTypePtrOutput() VariantTypePtrOutput {
	return o
}

func (o VariantTypePtrOutput) ToVariantTypePtrOutputWithContext(ctx context.Context) VariantTypePtrOutput {
	return o
}

func (o VariantTypePtrOutput) Elem() VariantTypeOutput {
	return o.ApplyT(func(v *VariantType) VariantType {
		if v != nil {
			return *v
		}
		var ret VariantType
		return ret
	}).(VariantTypeOutput)
}

func (o VariantTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VariantTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VariantType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VariantTypeInput interface {
	pulumi.Input

	ToVariantTypeOutput() VariantTypeOutput
	ToVariantTypeOutputWithContext(context.Context) VariantTypeOutput
}

var variantTypePtrType = reflect.TypeOf((**VariantType)(nil)).Elem()

type VariantTypePtrInput interface {
	pulumi.Input

	ToVariantTypePtrOutput() VariantTypePtrOutput
	ToVariantTypePtrOutputWithContext(context.Context) VariantTypePtrOutput
}

type variantTypePtr string

func VariantTypePtr(v string) VariantTypePtrInput {
	return (*variantTypePtr)(&v)
}

func (*variantTypePtr) ElementType() reflect.Type {
	return variantTypePtrType
}

func (in *variantTypePtr) ToVariantTypePtrOutput() VariantTypePtrOutput {
	return pulumi.ToOutput(in).(VariantTypePtrOutput)
}

func (in *variantTypePtr) ToVariantTypePtrOutputWithContext(ctx context.Context) VariantTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VariantTypePtrOutput)
}

type VmPriority string

const (
	VmPriorityDedicated   = VmPriority("Dedicated")
	VmPriorityLowPriority = VmPriority("LowPriority")
)

func (VmPriority) ElementType() reflect.Type {
	return reflect.TypeOf((*VmPriority)(nil)).Elem()
}

func (e VmPriority) ToVmPriorityOutput() VmPriorityOutput {
	return pulumi.ToOutput(e).(VmPriorityOutput)
}

func (e VmPriority) ToVmPriorityOutputWithContext(ctx context.Context) VmPriorityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VmPriorityOutput)
}

func (e VmPriority) ToVmPriorityPtrOutput() VmPriorityPtrOutput {
	return e.ToVmPriorityPtrOutputWithContext(context.Background())
}

func (e VmPriority) ToVmPriorityPtrOutputWithContext(ctx context.Context) VmPriorityPtrOutput {
	return VmPriority(e).ToVmPriorityOutputWithContext(ctx).ToVmPriorityPtrOutputWithContext(ctx)
}

func (e VmPriority) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VmPriority) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VmPriority) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VmPriority) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VmPriorityOutput struct{ *pulumi.OutputState }

func (VmPriorityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmPriority)(nil)).Elem()
}

func (o VmPriorityOutput) ToVmPriorityOutput() VmPriorityOutput {
	return o
}

func (o VmPriorityOutput) ToVmPriorityOutputWithContext(ctx context.Context) VmPriorityOutput {
	return o
}

func (o VmPriorityOutput) ToVmPriorityPtrOutput() VmPriorityPtrOutput {
	return o.ToVmPriorityPtrOutputWithContext(context.Background())
}

func (o VmPriorityOutput) ToVmPriorityPtrOutputWithContext(ctx context.Context) VmPriorityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VmPriority) *VmPriority {
		return &v
	}).(VmPriorityPtrOutput)
}

func (o VmPriorityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VmPriorityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VmPriority) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VmPriorityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VmPriorityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VmPriority) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VmPriorityPtrOutput struct{ *pulumi.OutputState }

func (VmPriorityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VmPriority)(nil)).Elem()
}

func (o VmPriorityPtrOutput) ToVmPriorityPtrOutput() VmPriorityPtrOutput {
	return o
}

func (o VmPriorityPtrOutput) ToVmPriorityPtrOutputWithContext(ctx context.Context) VmPriorityPtrOutput {
	return o
}

func (o VmPriorityPtrOutput) Elem() VmPriorityOutput {
	return o.ApplyT(func(v *VmPriority) VmPriority {
		if v != nil {
			return *v
		}
		var ret VmPriority
		return ret
	}).(VmPriorityOutput)
}

func (o VmPriorityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VmPriorityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VmPriority) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VmPriorityInput interface {
	pulumi.Input

	ToVmPriorityOutput() VmPriorityOutput
	ToVmPriorityOutputWithContext(context.Context) VmPriorityOutput
}

var vmPriorityPtrType = reflect.TypeOf((**VmPriority)(nil)).Elem()

type VmPriorityPtrInput interface {
	pulumi.Input

	ToVmPriorityPtrOutput() VmPriorityPtrOutput
	ToVmPriorityPtrOutputWithContext(context.Context) VmPriorityPtrOutput
}

type vmPriorityPtr string

func VmPriorityPtr(v string) VmPriorityPtrInput {
	return (*vmPriorityPtr)(&v)
}

func (*vmPriorityPtr) ElementType() reflect.Type {
	return vmPriorityPtrType
}

func (in *vmPriorityPtr) ToVmPriorityPtrOutput() VmPriorityPtrOutput {
	return pulumi.ToOutput(in).(VmPriorityPtrOutput)
}

func (in *vmPriorityPtr) ToVmPriorityPtrOutputWithContext(ctx context.Context) VmPriorityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VmPriorityPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationSharingPolicyInput)(nil)).Elem(), ApplicationSharingPolicy("Personal"))
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationSharingPolicyPtrInput)(nil)).Elem(), ApplicationSharingPolicy("Personal"))
	pulumi.RegisterInputType(reflect.TypeOf((*BatchLoggingLevelInput)(nil)).Elem(), BatchLoggingLevel("Info"))
	pulumi.RegisterInputType(reflect.TypeOf((*BatchLoggingLevelPtrInput)(nil)).Elem(), BatchLoggingLevel("Info"))
	pulumi.RegisterInputType(reflect.TypeOf((*BatchOutputActionInput)(nil)).Elem(), BatchOutputAction("SummaryOnly"))
	pulumi.RegisterInputType(reflect.TypeOf((*BatchOutputActionPtrInput)(nil)).Elem(), BatchOutputAction("SummaryOnly"))
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterPurposeInput)(nil)).Elem(), ClusterPurpose("FastProd"))
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterPurposePtrInput)(nil)).Elem(), ClusterPurpose("FastProd"))
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeEnvironmentTypeInput)(nil)).Elem(), ComputeEnvironmentType("ACI"))
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeEnvironmentTypePtrInput)(nil)).Elem(), ComputeEnvironmentType("ACI"))
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeInstanceAuthorizationTypeInput)(nil)).Elem(), ComputeInstanceAuthorizationType("personal"))
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeInstanceAuthorizationTypePtrInput)(nil)).Elem(), ComputeInstanceAuthorizationType("personal"))
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeTypeInput)(nil)).Elem(), ComputeType("AKS"))
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeTypePtrInput)(nil)).Elem(), ComputeType("AKS"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerTypeInput)(nil)).Elem(), ContainerType("StorageInitializer"))
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerTypePtrInput)(nil)).Elem(), ContainerType("StorageInitializer"))
	pulumi.RegisterInputType(reflect.TypeOf((*DataBindingModeInput)(nil)).Elem(), DataBindingMode("Mount"))
	pulumi.RegisterInputType(reflect.TypeOf((*DataBindingModePtrInput)(nil)).Elem(), DataBindingMode("Mount"))
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetTypeInput)(nil)).Elem(), DatasetType("tabular"))
	pulumi.RegisterInputType(reflect.TypeOf((*DatasetTypePtrInput)(nil)).Elem(), DatasetType("tabular"))
	pulumi.RegisterInputType(reflect.TypeOf((*DatastoreTypeArmInput)(nil)).Elem(), DatastoreTypeArm("blob"))
	pulumi.RegisterInputType(reflect.TypeOf((*DatastoreTypeArmPtrInput)(nil)).Elem(), DatastoreTypeArm("blob"))
	pulumi.RegisterInputType(reflect.TypeOf((*DistributionTypeInput)(nil)).Elem(), DistributionType("PyTorch"))
	pulumi.RegisterInputType(reflect.TypeOf((*DistributionTypePtrInput)(nil)).Elem(), DistributionType("PyTorch"))
	pulumi.RegisterInputType(reflect.TypeOf((*DockerSpecificationTypeInput)(nil)).Elem(), DockerSpecificationType("Build"))
	pulumi.RegisterInputType(reflect.TypeOf((*DockerSpecificationTypePtrInput)(nil)).Elem(), DockerSpecificationType("Build"))
	pulumi.RegisterInputType(reflect.TypeOf((*EarlyTerminationPolicyTypeInput)(nil)).Elem(), EarlyTerminationPolicyType("Bandit"))
	pulumi.RegisterInputType(reflect.TypeOf((*EarlyTerminationPolicyTypePtrInput)(nil)).Elem(), EarlyTerminationPolicyType("Bandit"))
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionStatusInput)(nil)).Elem(), EncryptionStatus("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionStatusPtrInput)(nil)).Elem(), EncryptionStatus("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointAuthModeInput)(nil)).Elem(), EndpointAuthMode("AMLToken"))
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointAuthModePtrInput)(nil)).Elem(), EndpointAuthMode("AMLToken"))
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointComputeTypeInput)(nil)).Elem(), EndpointComputeType("Managed"))
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointComputeTypePtrInput)(nil)).Elem(), EndpointComputeType("Managed"))
	pulumi.RegisterInputType(reflect.TypeOf((*GoalInput)(nil)).Elem(), Goal("Minimize"))
	pulumi.RegisterInputType(reflect.TypeOf((*GoalPtrInput)(nil)).Elem(), Goal("Minimize"))
	pulumi.RegisterInputType(reflect.TypeOf((*HeaderInput)(nil)).Elem(), Header("all_files_have_same_headers"))
	pulumi.RegisterInputType(reflect.TypeOf((*HeaderPtrInput)(nil)).Elem(), Header("all_files_have_same_headers"))
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityConfigurationTypeInput)(nil)).Elem(), IdentityConfigurationType("Managed"))
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityConfigurationTypePtrInput)(nil)).Elem(), IdentityConfigurationType("Managed"))
	pulumi.RegisterInputType(reflect.TypeOf((*ImageAnnotationTypeInput)(nil)).Elem(), ImageAnnotationType("Classification"))
	pulumi.RegisterInputType(reflect.TypeOf((*ImageAnnotationTypePtrInput)(nil)).Elem(), ImageAnnotationType("Classification"))
	pulumi.RegisterInputType(reflect.TypeOf((*JobTypeInput)(nil)).Elem(), JobType("Command"))
	pulumi.RegisterInputType(reflect.TypeOf((*JobTypePtrInput)(nil)).Elem(), JobType("Command"))
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedServiceLinkTypeInput)(nil)).Elem(), LinkedServiceLinkType("Synapse"))
	pulumi.RegisterInputType(reflect.TypeOf((*LinkedServiceLinkTypePtrInput)(nil)).Elem(), LinkedServiceLinkType("Synapse"))
	pulumi.RegisterInputType(reflect.TypeOf((*MediaTypeInput)(nil)).Elem(), MediaType("Image"))
	pulumi.RegisterInputType(reflect.TypeOf((*MediaTypePtrInput)(nil)).Elem(), MediaType("Image"))
	pulumi.RegisterInputType(reflect.TypeOf((*OperatingSystemTypeInput)(nil)).Elem(), OperatingSystemType("Linux"))
	pulumi.RegisterInputType(reflect.TypeOf((*OperatingSystemTypePtrInput)(nil)).Elem(), OperatingSystemType("Linux"))
	pulumi.RegisterInputType(reflect.TypeOf((*OsTypeInput)(nil)).Elem(), OsType("Linux"))
	pulumi.RegisterInputType(reflect.TypeOf((*OsTypePtrInput)(nil)).Elem(), OsType("Linux"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointServiceConnectionStatusInput)(nil)).Elem(), PrivateEndpointServiceConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointServiceConnectionStatusPtrInput)(nil)).Elem(), PrivateEndpointServiceConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*ReferenceTypeInput)(nil)).Elem(), ReferenceType("Id"))
	pulumi.RegisterInputType(reflect.TypeOf((*ReferenceTypePtrInput)(nil)).Elem(), ReferenceType("Id"))
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteLoginPortPublicAccessInput)(nil)).Elem(), RemoteLoginPortPublicAccess("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteLoginPortPublicAccessPtrInput)(nil)).Elem(), RemoteLoginPortPublicAccess("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityAssignmentInput)(nil)).Elem(), ResourceIdentityAssignment("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityAssignmentPtrInput)(nil)).Elem(), ResourceIdentityAssignment("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypeInput)(nil)).Elem(), ResourceIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypePtrInput)(nil)).Elem(), ResourceIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*SamplingAlgorithmInput)(nil)).Elem(), SamplingAlgorithm("Grid"))
	pulumi.RegisterInputType(reflect.TypeOf((*SamplingAlgorithmPtrInput)(nil)).Elem(), SamplingAlgorithm("Grid"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScaleTypeInput)(nil)).Elem(), ScaleType("Auto"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScaleTypePtrInput)(nil)).Elem(), ScaleType("Auto"))
	pulumi.RegisterInputType(reflect.TypeOf((*SourceTypeInput)(nil)).Elem(), SourceType("delimited_files"))
	pulumi.RegisterInputType(reflect.TypeOf((*SourceTypePtrInput)(nil)).Elem(), SourceType("delimited_files"))
	pulumi.RegisterInputType(reflect.TypeOf((*SshPublicAccessInput)(nil)).Elem(), SshPublicAccess("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*SshPublicAccessPtrInput)(nil)).Elem(), SshPublicAccess("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*ValueFormatInput)(nil)).Elem(), ValueFormat("JSON"))
	pulumi.RegisterInputType(reflect.TypeOf((*ValueFormatPtrInput)(nil)).Elem(), ValueFormat("JSON"))
	pulumi.RegisterInputType(reflect.TypeOf((*VariantTypeInput)(nil)).Elem(), VariantType("Control"))
	pulumi.RegisterInputType(reflect.TypeOf((*VariantTypePtrInput)(nil)).Elem(), VariantType("Control"))
	pulumi.RegisterInputType(reflect.TypeOf((*VmPriorityInput)(nil)).Elem(), VmPriority("Dedicated"))
	pulumi.RegisterInputType(reflect.TypeOf((*VmPriorityPtrInput)(nil)).Elem(), VmPriority("Dedicated"))
	pulumi.RegisterOutputType(ApplicationSharingPolicyOutput{})
	pulumi.RegisterOutputType(ApplicationSharingPolicyPtrOutput{})
	pulumi.RegisterOutputType(BatchLoggingLevelOutput{})
	pulumi.RegisterOutputType(BatchLoggingLevelPtrOutput{})
	pulumi.RegisterOutputType(BatchOutputActionOutput{})
	pulumi.RegisterOutputType(BatchOutputActionPtrOutput{})
	pulumi.RegisterOutputType(ClusterPurposeOutput{})
	pulumi.RegisterOutputType(ClusterPurposePtrOutput{})
	pulumi.RegisterOutputType(ComputeEnvironmentTypeOutput{})
	pulumi.RegisterOutputType(ComputeEnvironmentTypePtrOutput{})
	pulumi.RegisterOutputType(ComputeInstanceAuthorizationTypeOutput{})
	pulumi.RegisterOutputType(ComputeInstanceAuthorizationTypePtrOutput{})
	pulumi.RegisterOutputType(ComputeTypeOutput{})
	pulumi.RegisterOutputType(ComputeTypePtrOutput{})
	pulumi.RegisterOutputType(ContainerTypeOutput{})
	pulumi.RegisterOutputType(ContainerTypePtrOutput{})
	pulumi.RegisterOutputType(DataBindingModeOutput{})
	pulumi.RegisterOutputType(DataBindingModePtrOutput{})
	pulumi.RegisterOutputType(DatasetTypeOutput{})
	pulumi.RegisterOutputType(DatasetTypePtrOutput{})
	pulumi.RegisterOutputType(DatastoreTypeArmOutput{})
	pulumi.RegisterOutputType(DatastoreTypeArmPtrOutput{})
	pulumi.RegisterOutputType(DistributionTypeOutput{})
	pulumi.RegisterOutputType(DistributionTypePtrOutput{})
	pulumi.RegisterOutputType(DockerSpecificationTypeOutput{})
	pulumi.RegisterOutputType(DockerSpecificationTypePtrOutput{})
	pulumi.RegisterOutputType(EarlyTerminationPolicyTypeOutput{})
	pulumi.RegisterOutputType(EarlyTerminationPolicyTypePtrOutput{})
	pulumi.RegisterOutputType(EncryptionStatusOutput{})
	pulumi.RegisterOutputType(EncryptionStatusPtrOutput{})
	pulumi.RegisterOutputType(EndpointAuthModeOutput{})
	pulumi.RegisterOutputType(EndpointAuthModePtrOutput{})
	pulumi.RegisterOutputType(EndpointComputeTypeOutput{})
	pulumi.RegisterOutputType(EndpointComputeTypePtrOutput{})
	pulumi.RegisterOutputType(GoalOutput{})
	pulumi.RegisterOutputType(GoalPtrOutput{})
	pulumi.RegisterOutputType(HeaderOutput{})
	pulumi.RegisterOutputType(HeaderPtrOutput{})
	pulumi.RegisterOutputType(IdentityConfigurationTypeOutput{})
	pulumi.RegisterOutputType(IdentityConfigurationTypePtrOutput{})
	pulumi.RegisterOutputType(ImageAnnotationTypeOutput{})
	pulumi.RegisterOutputType(ImageAnnotationTypePtrOutput{})
	pulumi.RegisterOutputType(JobTypeOutput{})
	pulumi.RegisterOutputType(JobTypePtrOutput{})
	pulumi.RegisterOutputType(LinkedServiceLinkTypeOutput{})
	pulumi.RegisterOutputType(LinkedServiceLinkTypePtrOutput{})
	pulumi.RegisterOutputType(MediaTypeOutput{})
	pulumi.RegisterOutputType(MediaTypePtrOutput{})
	pulumi.RegisterOutputType(OperatingSystemTypeOutput{})
	pulumi.RegisterOutputType(OperatingSystemTypePtrOutput{})
	pulumi.RegisterOutputType(OsTypeOutput{})
	pulumi.RegisterOutputType(OsTypePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(ReferenceTypeOutput{})
	pulumi.RegisterOutputType(ReferenceTypePtrOutput{})
	pulumi.RegisterOutputType(RemoteLoginPortPublicAccessOutput{})
	pulumi.RegisterOutputType(RemoteLoginPortPublicAccessPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityAssignmentOutput{})
	pulumi.RegisterOutputType(ResourceIdentityAssignmentPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(SamplingAlgorithmOutput{})
	pulumi.RegisterOutputType(SamplingAlgorithmPtrOutput{})
	pulumi.RegisterOutputType(ScaleTypeOutput{})
	pulumi.RegisterOutputType(ScaleTypePtrOutput{})
	pulumi.RegisterOutputType(SourceTypeOutput{})
	pulumi.RegisterOutputType(SourceTypePtrOutput{})
	pulumi.RegisterOutputType(SshPublicAccessOutput{})
	pulumi.RegisterOutputType(SshPublicAccessPtrOutput{})
	pulumi.RegisterOutputType(ValueFormatOutput{})
	pulumi.RegisterOutputType(ValueFormatPtrOutput{})
	pulumi.RegisterOutputType(VariantTypeOutput{})
	pulumi.RegisterOutputType(VariantTypePtrOutput{})
	pulumi.RegisterOutputType(VmPriorityOutput{})
	pulumi.RegisterOutputType(VmPriorityPtrOutput{})
}
