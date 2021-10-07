


package v20200901preview

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
	pulumi.RegisterOutputType(ApplicationSharingPolicyOutput{})
	pulumi.RegisterOutputType(ApplicationSharingPolicyPtrOutput{})
	pulumi.RegisterOutputType(ComputeEnvironmentTypeOutput{})
	pulumi.RegisterOutputType(ComputeEnvironmentTypePtrOutput{})
	pulumi.RegisterOutputType(ComputeInstanceAuthorizationTypeOutput{})
	pulumi.RegisterOutputType(ComputeInstanceAuthorizationTypePtrOutput{})
	pulumi.RegisterOutputType(ComputeTypeOutput{})
	pulumi.RegisterOutputType(ComputeTypePtrOutput{})
	pulumi.RegisterOutputType(EncryptionStatusOutput{})
	pulumi.RegisterOutputType(EncryptionStatusPtrOutput{})
	pulumi.RegisterOutputType(ImageAnnotationTypeOutput{})
	pulumi.RegisterOutputType(ImageAnnotationTypePtrOutput{})
	pulumi.RegisterOutputType(LinkedServiceLinkTypeOutput{})
	pulumi.RegisterOutputType(LinkedServiceLinkTypePtrOutput{})
	pulumi.RegisterOutputType(MediaTypeOutput{})
	pulumi.RegisterOutputType(MediaTypePtrOutput{})
	pulumi.RegisterOutputType(OsTypeOutput{})
	pulumi.RegisterOutputType(OsTypePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(RemoteLoginPortPublicAccessOutput{})
	pulumi.RegisterOutputType(RemoteLoginPortPublicAccessPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(SshPublicAccessOutput{})
	pulumi.RegisterOutputType(SshPublicAccessPtrOutput{})
	pulumi.RegisterOutputType(VariantTypeOutput{})
	pulumi.RegisterOutputType(VariantTypePtrOutput{})
	pulumi.RegisterOutputType(VmPriorityOutput{})
	pulumi.RegisterOutputType(VmPriorityPtrOutput{})
}
