


package v20201001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CleanupOptions string

const (
	CleanupOptionsAlways       = CleanupOptions("Always")
	CleanupOptionsOnSuccess    = CleanupOptions("OnSuccess")
	CleanupOptionsOnExpiration = CleanupOptions("OnExpiration")
)

func (CleanupOptions) ElementType() reflect.Type {
	return reflect.TypeOf((*CleanupOptions)(nil)).Elem()
}

func (e CleanupOptions) ToCleanupOptionsOutput() CleanupOptionsOutput {
	return pulumi.ToOutput(e).(CleanupOptionsOutput)
}

func (e CleanupOptions) ToCleanupOptionsOutputWithContext(ctx context.Context) CleanupOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CleanupOptionsOutput)
}

func (e CleanupOptions) ToCleanupOptionsPtrOutput() CleanupOptionsPtrOutput {
	return e.ToCleanupOptionsPtrOutputWithContext(context.Background())
}

func (e CleanupOptions) ToCleanupOptionsPtrOutputWithContext(ctx context.Context) CleanupOptionsPtrOutput {
	return CleanupOptions(e).ToCleanupOptionsOutputWithContext(ctx).ToCleanupOptionsPtrOutputWithContext(ctx)
}

func (e CleanupOptions) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CleanupOptions) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CleanupOptions) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CleanupOptions) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CleanupOptionsOutput struct{ *pulumi.OutputState }

func (CleanupOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CleanupOptions)(nil)).Elem()
}

func (o CleanupOptionsOutput) ToCleanupOptionsOutput() CleanupOptionsOutput {
	return o
}

func (o CleanupOptionsOutput) ToCleanupOptionsOutputWithContext(ctx context.Context) CleanupOptionsOutput {
	return o
}

func (o CleanupOptionsOutput) ToCleanupOptionsPtrOutput() CleanupOptionsPtrOutput {
	return o.ToCleanupOptionsPtrOutputWithContext(context.Background())
}

func (o CleanupOptionsOutput) ToCleanupOptionsPtrOutputWithContext(ctx context.Context) CleanupOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CleanupOptions) *CleanupOptions {
		return &v
	}).(CleanupOptionsPtrOutput)
}

func (o CleanupOptionsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CleanupOptionsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CleanupOptions) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CleanupOptionsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CleanupOptionsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CleanupOptions) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CleanupOptionsPtrOutput struct{ *pulumi.OutputState }

func (CleanupOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CleanupOptions)(nil)).Elem()
}

func (o CleanupOptionsPtrOutput) ToCleanupOptionsPtrOutput() CleanupOptionsPtrOutput {
	return o
}

func (o CleanupOptionsPtrOutput) ToCleanupOptionsPtrOutputWithContext(ctx context.Context) CleanupOptionsPtrOutput {
	return o
}

func (o CleanupOptionsPtrOutput) Elem() CleanupOptionsOutput {
	return o.ApplyT(func(v *CleanupOptions) CleanupOptions {
		if v != nil {
			return *v
		}
		var ret CleanupOptions
		return ret
	}).(CleanupOptionsOutput)
}

func (o CleanupOptionsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CleanupOptionsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CleanupOptions) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CleanupOptionsInput interface {
	pulumi.Input

	ToCleanupOptionsOutput() CleanupOptionsOutput
	ToCleanupOptionsOutputWithContext(context.Context) CleanupOptionsOutput
}

var cleanupOptionsPtrType = reflect.TypeOf((**CleanupOptions)(nil)).Elem()

type CleanupOptionsPtrInput interface {
	pulumi.Input

	ToCleanupOptionsPtrOutput() CleanupOptionsPtrOutput
	ToCleanupOptionsPtrOutputWithContext(context.Context) CleanupOptionsPtrOutput
}

type cleanupOptionsPtr string

func CleanupOptionsPtr(v string) CleanupOptionsPtrInput {
	return (*cleanupOptionsPtr)(&v)
}

func (*cleanupOptionsPtr) ElementType() reflect.Type {
	return cleanupOptionsPtrType
}

func (in *cleanupOptionsPtr) ToCleanupOptionsPtrOutput() CleanupOptionsPtrOutput {
	return pulumi.ToOutput(in).(CleanupOptionsPtrOutput)
}

func (in *cleanupOptionsPtr) ToCleanupOptionsPtrOutputWithContext(ctx context.Context) CleanupOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CleanupOptionsPtrOutput)
}

type DeploymentMode string

const (
	DeploymentModeIncremental = DeploymentMode("Incremental")
	DeploymentModeComplete    = DeploymentMode("Complete")
)

func (DeploymentMode) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentMode)(nil)).Elem()
}

func (e DeploymentMode) ToDeploymentModeOutput() DeploymentModeOutput {
	return pulumi.ToOutput(e).(DeploymentModeOutput)
}

func (e DeploymentMode) ToDeploymentModeOutputWithContext(ctx context.Context) DeploymentModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DeploymentModeOutput)
}

func (e DeploymentMode) ToDeploymentModePtrOutput() DeploymentModePtrOutput {
	return e.ToDeploymentModePtrOutputWithContext(context.Background())
}

func (e DeploymentMode) ToDeploymentModePtrOutputWithContext(ctx context.Context) DeploymentModePtrOutput {
	return DeploymentMode(e).ToDeploymentModeOutputWithContext(ctx).ToDeploymentModePtrOutputWithContext(ctx)
}

func (e DeploymentMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeploymentMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeploymentMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DeploymentMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DeploymentModeOutput struct{ *pulumi.OutputState }

func (DeploymentModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentMode)(nil)).Elem()
}

func (o DeploymentModeOutput) ToDeploymentModeOutput() DeploymentModeOutput {
	return o
}

func (o DeploymentModeOutput) ToDeploymentModeOutputWithContext(ctx context.Context) DeploymentModeOutput {
	return o
}

func (o DeploymentModeOutput) ToDeploymentModePtrOutput() DeploymentModePtrOutput {
	return o.ToDeploymentModePtrOutputWithContext(context.Background())
}

func (o DeploymentModeOutput) ToDeploymentModePtrOutputWithContext(ctx context.Context) DeploymentModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeploymentMode) *DeploymentMode {
		return &v
	}).(DeploymentModePtrOutput)
}

func (o DeploymentModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DeploymentModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeploymentMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DeploymentModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeploymentModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeploymentMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DeploymentModePtrOutput struct{ *pulumi.OutputState }

func (DeploymentModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentMode)(nil)).Elem()
}

func (o DeploymentModePtrOutput) ToDeploymentModePtrOutput() DeploymentModePtrOutput {
	return o
}

func (o DeploymentModePtrOutput) ToDeploymentModePtrOutputWithContext(ctx context.Context) DeploymentModePtrOutput {
	return o
}

func (o DeploymentModePtrOutput) Elem() DeploymentModeOutput {
	return o.ApplyT(func(v *DeploymentMode) DeploymentMode {
		if v != nil {
			return *v
		}
		var ret DeploymentMode
		return ret
	}).(DeploymentModeOutput)
}

func (o DeploymentModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeploymentModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DeploymentMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DeploymentModeInput interface {
	pulumi.Input

	ToDeploymentModeOutput() DeploymentModeOutput
	ToDeploymentModeOutputWithContext(context.Context) DeploymentModeOutput
}

var deploymentModePtrType = reflect.TypeOf((**DeploymentMode)(nil)).Elem()

type DeploymentModePtrInput interface {
	pulumi.Input

	ToDeploymentModePtrOutput() DeploymentModePtrOutput
	ToDeploymentModePtrOutputWithContext(context.Context) DeploymentModePtrOutput
}

type deploymentModePtr string

func DeploymentModePtr(v string) DeploymentModePtrInput {
	return (*deploymentModePtr)(&v)
}

func (*deploymentModePtr) ElementType() reflect.Type {
	return deploymentModePtrType
}

func (in *deploymentModePtr) ToDeploymentModePtrOutput() DeploymentModePtrOutput {
	return pulumi.ToOutput(in).(DeploymentModePtrOutput)
}

func (in *deploymentModePtr) ToDeploymentModePtrOutputWithContext(ctx context.Context) DeploymentModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DeploymentModePtrOutput)
}

type ExpressionEvaluationOptionsScopeType string

const (
	ExpressionEvaluationOptionsScopeTypeNotSpecified = ExpressionEvaluationOptionsScopeType("NotSpecified")
	ExpressionEvaluationOptionsScopeTypeOuter        = ExpressionEvaluationOptionsScopeType("Outer")
	ExpressionEvaluationOptionsScopeTypeInner        = ExpressionEvaluationOptionsScopeType("Inner")
)

func (ExpressionEvaluationOptionsScopeType) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressionEvaluationOptionsScopeType)(nil)).Elem()
}

func (e ExpressionEvaluationOptionsScopeType) ToExpressionEvaluationOptionsScopeTypeOutput() ExpressionEvaluationOptionsScopeTypeOutput {
	return pulumi.ToOutput(e).(ExpressionEvaluationOptionsScopeTypeOutput)
}

func (e ExpressionEvaluationOptionsScopeType) ToExpressionEvaluationOptionsScopeTypeOutputWithContext(ctx context.Context) ExpressionEvaluationOptionsScopeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExpressionEvaluationOptionsScopeTypeOutput)
}

func (e ExpressionEvaluationOptionsScopeType) ToExpressionEvaluationOptionsScopeTypePtrOutput() ExpressionEvaluationOptionsScopeTypePtrOutput {
	return e.ToExpressionEvaluationOptionsScopeTypePtrOutputWithContext(context.Background())
}

func (e ExpressionEvaluationOptionsScopeType) ToExpressionEvaluationOptionsScopeTypePtrOutputWithContext(ctx context.Context) ExpressionEvaluationOptionsScopeTypePtrOutput {
	return ExpressionEvaluationOptionsScopeType(e).ToExpressionEvaluationOptionsScopeTypeOutputWithContext(ctx).ToExpressionEvaluationOptionsScopeTypePtrOutputWithContext(ctx)
}

func (e ExpressionEvaluationOptionsScopeType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressionEvaluationOptionsScopeType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExpressionEvaluationOptionsScopeType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExpressionEvaluationOptionsScopeType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExpressionEvaluationOptionsScopeTypeOutput struct{ *pulumi.OutputState }

func (ExpressionEvaluationOptionsScopeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressionEvaluationOptionsScopeType)(nil)).Elem()
}

func (o ExpressionEvaluationOptionsScopeTypeOutput) ToExpressionEvaluationOptionsScopeTypeOutput() ExpressionEvaluationOptionsScopeTypeOutput {
	return o
}

func (o ExpressionEvaluationOptionsScopeTypeOutput) ToExpressionEvaluationOptionsScopeTypeOutputWithContext(ctx context.Context) ExpressionEvaluationOptionsScopeTypeOutput {
	return o
}

func (o ExpressionEvaluationOptionsScopeTypeOutput) ToExpressionEvaluationOptionsScopeTypePtrOutput() ExpressionEvaluationOptionsScopeTypePtrOutput {
	return o.ToExpressionEvaluationOptionsScopeTypePtrOutputWithContext(context.Background())
}

func (o ExpressionEvaluationOptionsScopeTypeOutput) ToExpressionEvaluationOptionsScopeTypePtrOutputWithContext(ctx context.Context) ExpressionEvaluationOptionsScopeTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressionEvaluationOptionsScopeType) *ExpressionEvaluationOptionsScopeType {
		return &v
	}).(ExpressionEvaluationOptionsScopeTypePtrOutput)
}

func (o ExpressionEvaluationOptionsScopeTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExpressionEvaluationOptionsScopeTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressionEvaluationOptionsScopeType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExpressionEvaluationOptionsScopeTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressionEvaluationOptionsScopeTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExpressionEvaluationOptionsScopeType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExpressionEvaluationOptionsScopeTypePtrOutput struct{ *pulumi.OutputState }

func (ExpressionEvaluationOptionsScopeTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressionEvaluationOptionsScopeType)(nil)).Elem()
}

func (o ExpressionEvaluationOptionsScopeTypePtrOutput) ToExpressionEvaluationOptionsScopeTypePtrOutput() ExpressionEvaluationOptionsScopeTypePtrOutput {
	return o
}

func (o ExpressionEvaluationOptionsScopeTypePtrOutput) ToExpressionEvaluationOptionsScopeTypePtrOutputWithContext(ctx context.Context) ExpressionEvaluationOptionsScopeTypePtrOutput {
	return o
}

func (o ExpressionEvaluationOptionsScopeTypePtrOutput) Elem() ExpressionEvaluationOptionsScopeTypeOutput {
	return o.ApplyT(func(v *ExpressionEvaluationOptionsScopeType) ExpressionEvaluationOptionsScopeType {
		if v != nil {
			return *v
		}
		var ret ExpressionEvaluationOptionsScopeType
		return ret
	}).(ExpressionEvaluationOptionsScopeTypeOutput)
}

func (o ExpressionEvaluationOptionsScopeTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExpressionEvaluationOptionsScopeTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExpressionEvaluationOptionsScopeType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExpressionEvaluationOptionsScopeTypeInput interface {
	pulumi.Input

	ToExpressionEvaluationOptionsScopeTypeOutput() ExpressionEvaluationOptionsScopeTypeOutput
	ToExpressionEvaluationOptionsScopeTypeOutputWithContext(context.Context) ExpressionEvaluationOptionsScopeTypeOutput
}

var expressionEvaluationOptionsScopeTypePtrType = reflect.TypeOf((**ExpressionEvaluationOptionsScopeType)(nil)).Elem()

type ExpressionEvaluationOptionsScopeTypePtrInput interface {
	pulumi.Input

	ToExpressionEvaluationOptionsScopeTypePtrOutput() ExpressionEvaluationOptionsScopeTypePtrOutput
	ToExpressionEvaluationOptionsScopeTypePtrOutputWithContext(context.Context) ExpressionEvaluationOptionsScopeTypePtrOutput
}

type expressionEvaluationOptionsScopeTypePtr string

func ExpressionEvaluationOptionsScopeTypePtr(v string) ExpressionEvaluationOptionsScopeTypePtrInput {
	return (*expressionEvaluationOptionsScopeTypePtr)(&v)
}

func (*expressionEvaluationOptionsScopeTypePtr) ElementType() reflect.Type {
	return expressionEvaluationOptionsScopeTypePtrType
}

func (in *expressionEvaluationOptionsScopeTypePtr) ToExpressionEvaluationOptionsScopeTypePtrOutput() ExpressionEvaluationOptionsScopeTypePtrOutput {
	return pulumi.ToOutput(in).(ExpressionEvaluationOptionsScopeTypePtrOutput)
}

func (in *expressionEvaluationOptionsScopeTypePtr) ToExpressionEvaluationOptionsScopeTypePtrOutputWithContext(ctx context.Context) ExpressionEvaluationOptionsScopeTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExpressionEvaluationOptionsScopeTypePtrOutput)
}

type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeUserAssigned = ManagedServiceIdentityType("UserAssigned")
)

func (ManagedServiceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityType)(nil)).Elem()
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypeOutput() ManagedServiceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ManagedServiceIdentityTypeOutput)
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypeOutputWithContext(ctx context.Context) ManagedServiceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedServiceIdentityTypeOutput)
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return e.ToManagedServiceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return ManagedServiceIdentityType(e).ToManagedServiceIdentityTypeOutputWithContext(ctx).ToManagedServiceIdentityTypePtrOutputWithContext(ctx)
}

func (e ManagedServiceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedServiceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedServiceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedServiceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedServiceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityType)(nil)).Elem()
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypeOutput() ManagedServiceIdentityTypeOutput {
	return o
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypeOutputWithContext(ctx context.Context) ManagedServiceIdentityTypeOutput {
	return o
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return o.ToManagedServiceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServiceIdentityType) *ManagedServiceIdentityType {
		return &v
	}).(ManagedServiceIdentityTypePtrOutput)
}

func (o ManagedServiceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedServiceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedServiceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedServiceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentityType)(nil)).Elem()
}

func (o ManagedServiceIdentityTypePtrOutput) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return o
}

func (o ManagedServiceIdentityTypePtrOutput) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return o
}

func (o ManagedServiceIdentityTypePtrOutput) Elem() ManagedServiceIdentityTypeOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityType) ManagedServiceIdentityType {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentityType
		return ret
	}).(ManagedServiceIdentityTypeOutput)
}

func (o ManagedServiceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedServiceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ManagedServiceIdentityTypeInput interface {
	pulumi.Input

	ToManagedServiceIdentityTypeOutput() ManagedServiceIdentityTypeOutput
	ToManagedServiceIdentityTypeOutputWithContext(context.Context) ManagedServiceIdentityTypeOutput
}

var managedServiceIdentityTypePtrType = reflect.TypeOf((**ManagedServiceIdentityType)(nil)).Elem()

type ManagedServiceIdentityTypePtrInput interface {
	pulumi.Input

	ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput
	ToManagedServiceIdentityTypePtrOutputWithContext(context.Context) ManagedServiceIdentityTypePtrOutput
}

type managedServiceIdentityTypePtr string

func ManagedServiceIdentityTypePtr(v string) ManagedServiceIdentityTypePtrInput {
	return (*managedServiceIdentityTypePtr)(&v)
}

func (*managedServiceIdentityTypePtr) ElementType() reflect.Type {
	return managedServiceIdentityTypePtrType
}

func (in *managedServiceIdentityTypePtr) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ManagedServiceIdentityTypePtrOutput)
}

func (in *managedServiceIdentityTypePtr) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedServiceIdentityTypePtrOutput)
}

type OnErrorDeploymentType string

const (
	OnErrorDeploymentTypeLastSuccessful     = OnErrorDeploymentType("LastSuccessful")
	OnErrorDeploymentTypeSpecificDeployment = OnErrorDeploymentType("SpecificDeployment")
)

func (OnErrorDeploymentType) ElementType() reflect.Type {
	return reflect.TypeOf((*OnErrorDeploymentType)(nil)).Elem()
}

func (e OnErrorDeploymentType) ToOnErrorDeploymentTypeOutput() OnErrorDeploymentTypeOutput {
	return pulumi.ToOutput(e).(OnErrorDeploymentTypeOutput)
}

func (e OnErrorDeploymentType) ToOnErrorDeploymentTypeOutputWithContext(ctx context.Context) OnErrorDeploymentTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OnErrorDeploymentTypeOutput)
}

func (e OnErrorDeploymentType) ToOnErrorDeploymentTypePtrOutput() OnErrorDeploymentTypePtrOutput {
	return e.ToOnErrorDeploymentTypePtrOutputWithContext(context.Background())
}

func (e OnErrorDeploymentType) ToOnErrorDeploymentTypePtrOutputWithContext(ctx context.Context) OnErrorDeploymentTypePtrOutput {
	return OnErrorDeploymentType(e).ToOnErrorDeploymentTypeOutputWithContext(ctx).ToOnErrorDeploymentTypePtrOutputWithContext(ctx)
}

func (e OnErrorDeploymentType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OnErrorDeploymentType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OnErrorDeploymentType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OnErrorDeploymentType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OnErrorDeploymentTypeOutput struct{ *pulumi.OutputState }

func (OnErrorDeploymentTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OnErrorDeploymentType)(nil)).Elem()
}

func (o OnErrorDeploymentTypeOutput) ToOnErrorDeploymentTypeOutput() OnErrorDeploymentTypeOutput {
	return o
}

func (o OnErrorDeploymentTypeOutput) ToOnErrorDeploymentTypeOutputWithContext(ctx context.Context) OnErrorDeploymentTypeOutput {
	return o
}

func (o OnErrorDeploymentTypeOutput) ToOnErrorDeploymentTypePtrOutput() OnErrorDeploymentTypePtrOutput {
	return o.ToOnErrorDeploymentTypePtrOutputWithContext(context.Background())
}

func (o OnErrorDeploymentTypeOutput) ToOnErrorDeploymentTypePtrOutputWithContext(ctx context.Context) OnErrorDeploymentTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OnErrorDeploymentType) *OnErrorDeploymentType {
		return &v
	}).(OnErrorDeploymentTypePtrOutput)
}

func (o OnErrorDeploymentTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OnErrorDeploymentTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OnErrorDeploymentType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OnErrorDeploymentTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OnErrorDeploymentTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OnErrorDeploymentType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OnErrorDeploymentTypePtrOutput struct{ *pulumi.OutputState }

func (OnErrorDeploymentTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OnErrorDeploymentType)(nil)).Elem()
}

func (o OnErrorDeploymentTypePtrOutput) ToOnErrorDeploymentTypePtrOutput() OnErrorDeploymentTypePtrOutput {
	return o
}

func (o OnErrorDeploymentTypePtrOutput) ToOnErrorDeploymentTypePtrOutputWithContext(ctx context.Context) OnErrorDeploymentTypePtrOutput {
	return o
}

func (o OnErrorDeploymentTypePtrOutput) Elem() OnErrorDeploymentTypeOutput {
	return o.ApplyT(func(v *OnErrorDeploymentType) OnErrorDeploymentType {
		if v != nil {
			return *v
		}
		var ret OnErrorDeploymentType
		return ret
	}).(OnErrorDeploymentTypeOutput)
}

func (o OnErrorDeploymentTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OnErrorDeploymentTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OnErrorDeploymentType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OnErrorDeploymentTypeInput interface {
	pulumi.Input

	ToOnErrorDeploymentTypeOutput() OnErrorDeploymentTypeOutput
	ToOnErrorDeploymentTypeOutputWithContext(context.Context) OnErrorDeploymentTypeOutput
}

var onErrorDeploymentTypePtrType = reflect.TypeOf((**OnErrorDeploymentType)(nil)).Elem()

type OnErrorDeploymentTypePtrInput interface {
	pulumi.Input

	ToOnErrorDeploymentTypePtrOutput() OnErrorDeploymentTypePtrOutput
	ToOnErrorDeploymentTypePtrOutputWithContext(context.Context) OnErrorDeploymentTypePtrOutput
}

type onErrorDeploymentTypePtr string

func OnErrorDeploymentTypePtr(v string) OnErrorDeploymentTypePtrInput {
	return (*onErrorDeploymentTypePtr)(&v)
}

func (*onErrorDeploymentTypePtr) ElementType() reflect.Type {
	return onErrorDeploymentTypePtrType
}

func (in *onErrorDeploymentTypePtr) ToOnErrorDeploymentTypePtrOutput() OnErrorDeploymentTypePtrOutput {
	return pulumi.ToOutput(in).(OnErrorDeploymentTypePtrOutput)
}

func (in *onErrorDeploymentTypePtr) ToOnErrorDeploymentTypePtrOutputWithContext(ctx context.Context) OnErrorDeploymentTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OnErrorDeploymentTypePtrOutput)
}

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned, UserAssigned")
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

type ScriptType string

const (
	ScriptTypeAzurePowerShell = ScriptType("AzurePowerShell")
	ScriptTypeAzureCLI        = ScriptType("AzureCLI")
)

func (ScriptType) ElementType() reflect.Type {
	return reflect.TypeOf((*ScriptType)(nil)).Elem()
}

func (e ScriptType) ToScriptTypeOutput() ScriptTypeOutput {
	return pulumi.ToOutput(e).(ScriptTypeOutput)
}

func (e ScriptType) ToScriptTypeOutputWithContext(ctx context.Context) ScriptTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScriptTypeOutput)
}

func (e ScriptType) ToScriptTypePtrOutput() ScriptTypePtrOutput {
	return e.ToScriptTypePtrOutputWithContext(context.Background())
}

func (e ScriptType) ToScriptTypePtrOutputWithContext(ctx context.Context) ScriptTypePtrOutput {
	return ScriptType(e).ToScriptTypeOutputWithContext(ctx).ToScriptTypePtrOutputWithContext(ctx)
}

func (e ScriptType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScriptType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScriptType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScriptType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScriptTypeOutput struct{ *pulumi.OutputState }

func (ScriptTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScriptType)(nil)).Elem()
}

func (o ScriptTypeOutput) ToScriptTypeOutput() ScriptTypeOutput {
	return o
}

func (o ScriptTypeOutput) ToScriptTypeOutputWithContext(ctx context.Context) ScriptTypeOutput {
	return o
}

func (o ScriptTypeOutput) ToScriptTypePtrOutput() ScriptTypePtrOutput {
	return o.ToScriptTypePtrOutputWithContext(context.Background())
}

func (o ScriptTypeOutput) ToScriptTypePtrOutputWithContext(ctx context.Context) ScriptTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScriptType) *ScriptType {
		return &v
	}).(ScriptTypePtrOutput)
}

func (o ScriptTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScriptTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScriptType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScriptTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScriptTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScriptType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScriptTypePtrOutput struct{ *pulumi.OutputState }

func (ScriptTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScriptType)(nil)).Elem()
}

func (o ScriptTypePtrOutput) ToScriptTypePtrOutput() ScriptTypePtrOutput {
	return o
}

func (o ScriptTypePtrOutput) ToScriptTypePtrOutputWithContext(ctx context.Context) ScriptTypePtrOutput {
	return o
}

func (o ScriptTypePtrOutput) Elem() ScriptTypeOutput {
	return o.ApplyT(func(v *ScriptType) ScriptType {
		if v != nil {
			return *v
		}
		var ret ScriptType
		return ret
	}).(ScriptTypeOutput)
}

func (o ScriptTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScriptTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScriptType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ScriptTypeInput interface {
	pulumi.Input

	ToScriptTypeOutput() ScriptTypeOutput
	ToScriptTypeOutputWithContext(context.Context) ScriptTypeOutput
}

var scriptTypePtrType = reflect.TypeOf((**ScriptType)(nil)).Elem()

type ScriptTypePtrInput interface {
	pulumi.Input

	ToScriptTypePtrOutput() ScriptTypePtrOutput
	ToScriptTypePtrOutputWithContext(context.Context) ScriptTypePtrOutput
}

type scriptTypePtr string

func ScriptTypePtr(v string) ScriptTypePtrInput {
	return (*scriptTypePtr)(&v)
}

func (*scriptTypePtr) ElementType() reflect.Type {
	return scriptTypePtrType
}

func (in *scriptTypePtr) ToScriptTypePtrOutput() ScriptTypePtrOutput {
	return pulumi.ToOutput(in).(ScriptTypePtrOutput)
}

func (in *scriptTypePtr) ToScriptTypePtrOutputWithContext(ctx context.Context) ScriptTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScriptTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CleanupOptionsInput)(nil)).Elem(), CleanupOptions("Always"))
	pulumi.RegisterInputType(reflect.TypeOf((*CleanupOptionsPtrInput)(nil)).Elem(), CleanupOptions("Always"))
	pulumi.RegisterInputType(reflect.TypeOf((*DeploymentModeInput)(nil)).Elem(), DeploymentMode("Incremental"))
	pulumi.RegisterInputType(reflect.TypeOf((*DeploymentModePtrInput)(nil)).Elem(), DeploymentMode("Incremental"))
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressionEvaluationOptionsScopeTypeInput)(nil)).Elem(), ExpressionEvaluationOptionsScopeType("NotSpecified"))
	pulumi.RegisterInputType(reflect.TypeOf((*ExpressionEvaluationOptionsScopeTypePtrInput)(nil)).Elem(), ExpressionEvaluationOptionsScopeType("NotSpecified"))
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedServiceIdentityTypeInput)(nil)).Elem(), ManagedServiceIdentityType("UserAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedServiceIdentityTypePtrInput)(nil)).Elem(), ManagedServiceIdentityType("UserAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*OnErrorDeploymentTypeInput)(nil)).Elem(), OnErrorDeploymentType("LastSuccessful"))
	pulumi.RegisterInputType(reflect.TypeOf((*OnErrorDeploymentTypePtrInput)(nil)).Elem(), OnErrorDeploymentType("LastSuccessful"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypeInput)(nil)).Elem(), ResourceIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypePtrInput)(nil)).Elem(), ResourceIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScriptTypeInput)(nil)).Elem(), ScriptType("AzurePowerShell"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScriptTypePtrInput)(nil)).Elem(), ScriptType("AzurePowerShell"))
	pulumi.RegisterOutputType(CleanupOptionsOutput{})
	pulumi.RegisterOutputType(CleanupOptionsPtrOutput{})
	pulumi.RegisterOutputType(DeploymentModeOutput{})
	pulumi.RegisterOutputType(DeploymentModePtrOutput{})
	pulumi.RegisterOutputType(ExpressionEvaluationOptionsScopeTypeOutput{})
	pulumi.RegisterOutputType(ExpressionEvaluationOptionsScopeTypePtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(OnErrorDeploymentTypeOutput{})
	pulumi.RegisterOutputType(OnErrorDeploymentTypePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(ScriptTypeOutput{})
	pulumi.RegisterOutputType(ScriptTypePtrOutput{})
}
