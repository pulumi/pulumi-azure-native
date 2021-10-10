


package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AddRemoveIncrementalNamedPartitionScalingMechanism struct {
	Kind              string `pulumi:"kind"`
	MaxPartitionCount int    `pulumi:"maxPartitionCount"`
	MinPartitionCount int    `pulumi:"minPartitionCount"`
	ScaleIncrement    int    `pulumi:"scaleIncrement"`
}





type AddRemoveIncrementalNamedPartitionScalingMechanismInput interface {
	pulumi.Input

	ToAddRemoveIncrementalNamedPartitionScalingMechanismOutput() AddRemoveIncrementalNamedPartitionScalingMechanismOutput
	ToAddRemoveIncrementalNamedPartitionScalingMechanismOutputWithContext(context.Context) AddRemoveIncrementalNamedPartitionScalingMechanismOutput
}

type AddRemoveIncrementalNamedPartitionScalingMechanismArgs struct {
	Kind              pulumi.StringInput `pulumi:"kind"`
	MaxPartitionCount pulumi.IntInput    `pulumi:"maxPartitionCount"`
	MinPartitionCount pulumi.IntInput    `pulumi:"minPartitionCount"`
	ScaleIncrement    pulumi.IntInput    `pulumi:"scaleIncrement"`
}

func (AddRemoveIncrementalNamedPartitionScalingMechanismArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddRemoveIncrementalNamedPartitionScalingMechanism)(nil)).Elem()
}

func (i AddRemoveIncrementalNamedPartitionScalingMechanismArgs) ToAddRemoveIncrementalNamedPartitionScalingMechanismOutput() AddRemoveIncrementalNamedPartitionScalingMechanismOutput {
	return i.ToAddRemoveIncrementalNamedPartitionScalingMechanismOutputWithContext(context.Background())
}

func (i AddRemoveIncrementalNamedPartitionScalingMechanismArgs) ToAddRemoveIncrementalNamedPartitionScalingMechanismOutputWithContext(ctx context.Context) AddRemoveIncrementalNamedPartitionScalingMechanismOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddRemoveIncrementalNamedPartitionScalingMechanismOutput)
}

type AddRemoveIncrementalNamedPartitionScalingMechanismOutput struct{ *pulumi.OutputState }

func (AddRemoveIncrementalNamedPartitionScalingMechanismOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddRemoveIncrementalNamedPartitionScalingMechanism)(nil)).Elem()
}

func (o AddRemoveIncrementalNamedPartitionScalingMechanismOutput) ToAddRemoveIncrementalNamedPartitionScalingMechanismOutput() AddRemoveIncrementalNamedPartitionScalingMechanismOutput {
	return o
}

func (o AddRemoveIncrementalNamedPartitionScalingMechanismOutput) ToAddRemoveIncrementalNamedPartitionScalingMechanismOutputWithContext(ctx context.Context) AddRemoveIncrementalNamedPartitionScalingMechanismOutput {
	return o
}

func (o AddRemoveIncrementalNamedPartitionScalingMechanismOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v AddRemoveIncrementalNamedPartitionScalingMechanism) string { return v.Kind }).(pulumi.StringOutput)
}

func (o AddRemoveIncrementalNamedPartitionScalingMechanismOutput) MaxPartitionCount() pulumi.IntOutput {
	return o.ApplyT(func(v AddRemoveIncrementalNamedPartitionScalingMechanism) int { return v.MaxPartitionCount }).(pulumi.IntOutput)
}

func (o AddRemoveIncrementalNamedPartitionScalingMechanismOutput) MinPartitionCount() pulumi.IntOutput {
	return o.ApplyT(func(v AddRemoveIncrementalNamedPartitionScalingMechanism) int { return v.MinPartitionCount }).(pulumi.IntOutput)
}

func (o AddRemoveIncrementalNamedPartitionScalingMechanismOutput) ScaleIncrement() pulumi.IntOutput {
	return o.ApplyT(func(v AddRemoveIncrementalNamedPartitionScalingMechanism) int { return v.ScaleIncrement }).(pulumi.IntOutput)
}

type AddRemoveIncrementalNamedPartitionScalingMechanismResponse struct {
	Kind              string `pulumi:"kind"`
	MaxPartitionCount int    `pulumi:"maxPartitionCount"`
	MinPartitionCount int    `pulumi:"minPartitionCount"`
	ScaleIncrement    int    `pulumi:"scaleIncrement"`
}





type AddRemoveIncrementalNamedPartitionScalingMechanismResponseInput interface {
	pulumi.Input

	ToAddRemoveIncrementalNamedPartitionScalingMechanismResponseOutput() AddRemoveIncrementalNamedPartitionScalingMechanismResponseOutput
	ToAddRemoveIncrementalNamedPartitionScalingMechanismResponseOutputWithContext(context.Context) AddRemoveIncrementalNamedPartitionScalingMechanismResponseOutput
}

type AddRemoveIncrementalNamedPartitionScalingMechanismResponseArgs struct {
	Kind              pulumi.StringInput `pulumi:"kind"`
	MaxPartitionCount pulumi.IntInput    `pulumi:"maxPartitionCount"`
	MinPartitionCount pulumi.IntInput    `pulumi:"minPartitionCount"`
	ScaleIncrement    pulumi.IntInput    `pulumi:"scaleIncrement"`
}

func (AddRemoveIncrementalNamedPartitionScalingMechanismResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddRemoveIncrementalNamedPartitionScalingMechanismResponse)(nil)).Elem()
}

func (i AddRemoveIncrementalNamedPartitionScalingMechanismResponseArgs) ToAddRemoveIncrementalNamedPartitionScalingMechanismResponseOutput() AddRemoveIncrementalNamedPartitionScalingMechanismResponseOutput {
	return i.ToAddRemoveIncrementalNamedPartitionScalingMechanismResponseOutputWithContext(context.Background())
}

func (i AddRemoveIncrementalNamedPartitionScalingMechanismResponseArgs) ToAddRemoveIncrementalNamedPartitionScalingMechanismResponseOutputWithContext(ctx context.Context) AddRemoveIncrementalNamedPartitionScalingMechanismResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddRemoveIncrementalNamedPartitionScalingMechanismResponseOutput)
}

type AddRemoveIncrementalNamedPartitionScalingMechanismResponseOutput struct{ *pulumi.OutputState }

func (AddRemoveIncrementalNamedPartitionScalingMechanismResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddRemoveIncrementalNamedPartitionScalingMechanismResponse)(nil)).Elem()
}

func (o AddRemoveIncrementalNamedPartitionScalingMechanismResponseOutput) ToAddRemoveIncrementalNamedPartitionScalingMechanismResponseOutput() AddRemoveIncrementalNamedPartitionScalingMechanismResponseOutput {
	return o
}

func (o AddRemoveIncrementalNamedPartitionScalingMechanismResponseOutput) ToAddRemoveIncrementalNamedPartitionScalingMechanismResponseOutputWithContext(ctx context.Context) AddRemoveIncrementalNamedPartitionScalingMechanismResponseOutput {
	return o
}

func (o AddRemoveIncrementalNamedPartitionScalingMechanismResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v AddRemoveIncrementalNamedPartitionScalingMechanismResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o AddRemoveIncrementalNamedPartitionScalingMechanismResponseOutput) MaxPartitionCount() pulumi.IntOutput {
	return o.ApplyT(func(v AddRemoveIncrementalNamedPartitionScalingMechanismResponse) int { return v.MaxPartitionCount }).(pulumi.IntOutput)
}

func (o AddRemoveIncrementalNamedPartitionScalingMechanismResponseOutput) MinPartitionCount() pulumi.IntOutput {
	return o.ApplyT(func(v AddRemoveIncrementalNamedPartitionScalingMechanismResponse) int { return v.MinPartitionCount }).(pulumi.IntOutput)
}

func (o AddRemoveIncrementalNamedPartitionScalingMechanismResponseOutput) ScaleIncrement() pulumi.IntOutput {
	return o.ApplyT(func(v AddRemoveIncrementalNamedPartitionScalingMechanismResponse) int { return v.ScaleIncrement }).(pulumi.IntOutput)
}

type ApplicationHealthPolicy struct {
	ConsiderWarningAsError                  bool                               `pulumi:"considerWarningAsError"`
	DefaultServiceTypeHealthPolicy          *ServiceTypeHealthPolicy           `pulumi:"defaultServiceTypeHealthPolicy"`
	MaxPercentUnhealthyDeployedApplications int                                `pulumi:"maxPercentUnhealthyDeployedApplications"`
	ServiceTypeHealthPolicyMap              map[string]ServiceTypeHealthPolicy `pulumi:"serviceTypeHealthPolicyMap"`
}





type ApplicationHealthPolicyInput interface {
	pulumi.Input

	ToApplicationHealthPolicyOutput() ApplicationHealthPolicyOutput
	ToApplicationHealthPolicyOutputWithContext(context.Context) ApplicationHealthPolicyOutput
}

type ApplicationHealthPolicyArgs struct {
	ConsiderWarningAsError                  pulumi.BoolInput                `pulumi:"considerWarningAsError"`
	DefaultServiceTypeHealthPolicy          ServiceTypeHealthPolicyPtrInput `pulumi:"defaultServiceTypeHealthPolicy"`
	MaxPercentUnhealthyDeployedApplications pulumi.IntInput                 `pulumi:"maxPercentUnhealthyDeployedApplications"`
	ServiceTypeHealthPolicyMap              ServiceTypeHealthPolicyMapInput `pulumi:"serviceTypeHealthPolicyMap"`
}

func (ApplicationHealthPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationHealthPolicy)(nil)).Elem()
}

func (i ApplicationHealthPolicyArgs) ToApplicationHealthPolicyOutput() ApplicationHealthPolicyOutput {
	return i.ToApplicationHealthPolicyOutputWithContext(context.Background())
}

func (i ApplicationHealthPolicyArgs) ToApplicationHealthPolicyOutputWithContext(ctx context.Context) ApplicationHealthPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationHealthPolicyOutput)
}

func (i ApplicationHealthPolicyArgs) ToApplicationHealthPolicyPtrOutput() ApplicationHealthPolicyPtrOutput {
	return i.ToApplicationHealthPolicyPtrOutputWithContext(context.Background())
}

func (i ApplicationHealthPolicyArgs) ToApplicationHealthPolicyPtrOutputWithContext(ctx context.Context) ApplicationHealthPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationHealthPolicyOutput).ToApplicationHealthPolicyPtrOutputWithContext(ctx)
}









type ApplicationHealthPolicyPtrInput interface {
	pulumi.Input

	ToApplicationHealthPolicyPtrOutput() ApplicationHealthPolicyPtrOutput
	ToApplicationHealthPolicyPtrOutputWithContext(context.Context) ApplicationHealthPolicyPtrOutput
}

type applicationHealthPolicyPtrType ApplicationHealthPolicyArgs

func ApplicationHealthPolicyPtr(v *ApplicationHealthPolicyArgs) ApplicationHealthPolicyPtrInput {
	return (*applicationHealthPolicyPtrType)(v)
}

func (*applicationHealthPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationHealthPolicy)(nil)).Elem()
}

func (i *applicationHealthPolicyPtrType) ToApplicationHealthPolicyPtrOutput() ApplicationHealthPolicyPtrOutput {
	return i.ToApplicationHealthPolicyPtrOutputWithContext(context.Background())
}

func (i *applicationHealthPolicyPtrType) ToApplicationHealthPolicyPtrOutputWithContext(ctx context.Context) ApplicationHealthPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationHealthPolicyPtrOutput)
}

type ApplicationHealthPolicyOutput struct{ *pulumi.OutputState }

func (ApplicationHealthPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationHealthPolicy)(nil)).Elem()
}

func (o ApplicationHealthPolicyOutput) ToApplicationHealthPolicyOutput() ApplicationHealthPolicyOutput {
	return o
}

func (o ApplicationHealthPolicyOutput) ToApplicationHealthPolicyOutputWithContext(ctx context.Context) ApplicationHealthPolicyOutput {
	return o
}

func (o ApplicationHealthPolicyOutput) ToApplicationHealthPolicyPtrOutput() ApplicationHealthPolicyPtrOutput {
	return o.ToApplicationHealthPolicyPtrOutputWithContext(context.Background())
}

func (o ApplicationHealthPolicyOutput) ToApplicationHealthPolicyPtrOutputWithContext(ctx context.Context) ApplicationHealthPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationHealthPolicy) *ApplicationHealthPolicy {
		return &v
	}).(ApplicationHealthPolicyPtrOutput)
}

func (o ApplicationHealthPolicyOutput) ConsiderWarningAsError() pulumi.BoolOutput {
	return o.ApplyT(func(v ApplicationHealthPolicy) bool { return v.ConsiderWarningAsError }).(pulumi.BoolOutput)
}

func (o ApplicationHealthPolicyOutput) DefaultServiceTypeHealthPolicy() ServiceTypeHealthPolicyPtrOutput {
	return o.ApplyT(func(v ApplicationHealthPolicy) *ServiceTypeHealthPolicy { return v.DefaultServiceTypeHealthPolicy }).(ServiceTypeHealthPolicyPtrOutput)
}

func (o ApplicationHealthPolicyOutput) MaxPercentUnhealthyDeployedApplications() pulumi.IntOutput {
	return o.ApplyT(func(v ApplicationHealthPolicy) int { return v.MaxPercentUnhealthyDeployedApplications }).(pulumi.IntOutput)
}

func (o ApplicationHealthPolicyOutput) ServiceTypeHealthPolicyMap() ServiceTypeHealthPolicyMapOutput {
	return o.ApplyT(func(v ApplicationHealthPolicy) map[string]ServiceTypeHealthPolicy {
		return v.ServiceTypeHealthPolicyMap
	}).(ServiceTypeHealthPolicyMapOutput)
}

type ApplicationHealthPolicyPtrOutput struct{ *pulumi.OutputState }

func (ApplicationHealthPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationHealthPolicy)(nil)).Elem()
}

func (o ApplicationHealthPolicyPtrOutput) ToApplicationHealthPolicyPtrOutput() ApplicationHealthPolicyPtrOutput {
	return o
}

func (o ApplicationHealthPolicyPtrOutput) ToApplicationHealthPolicyPtrOutputWithContext(ctx context.Context) ApplicationHealthPolicyPtrOutput {
	return o
}

func (o ApplicationHealthPolicyPtrOutput) Elem() ApplicationHealthPolicyOutput {
	return o.ApplyT(func(v *ApplicationHealthPolicy) ApplicationHealthPolicy {
		if v != nil {
			return *v
		}
		var ret ApplicationHealthPolicy
		return ret
	}).(ApplicationHealthPolicyOutput)
}

func (o ApplicationHealthPolicyPtrOutput) ConsiderWarningAsError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationHealthPolicy) *bool {
		if v == nil {
			return nil
		}
		return &v.ConsiderWarningAsError
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationHealthPolicyPtrOutput) DefaultServiceTypeHealthPolicy() ServiceTypeHealthPolicyPtrOutput {
	return o.ApplyT(func(v *ApplicationHealthPolicy) *ServiceTypeHealthPolicy {
		if v == nil {
			return nil
		}
		return v.DefaultServiceTypeHealthPolicy
	}).(ServiceTypeHealthPolicyPtrOutput)
}

func (o ApplicationHealthPolicyPtrOutput) MaxPercentUnhealthyDeployedApplications() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApplicationHealthPolicy) *int {
		if v == nil {
			return nil
		}
		return &v.MaxPercentUnhealthyDeployedApplications
	}).(pulumi.IntPtrOutput)
}

func (o ApplicationHealthPolicyPtrOutput) ServiceTypeHealthPolicyMap() ServiceTypeHealthPolicyMapOutput {
	return o.ApplyT(func(v *ApplicationHealthPolicy) map[string]ServiceTypeHealthPolicy {
		if v == nil {
			return nil
		}
		return v.ServiceTypeHealthPolicyMap
	}).(ServiceTypeHealthPolicyMapOutput)
}

type ApplicationHealthPolicyResponse struct {
	ConsiderWarningAsError                  bool                                       `pulumi:"considerWarningAsError"`
	DefaultServiceTypeHealthPolicy          *ServiceTypeHealthPolicyResponse           `pulumi:"defaultServiceTypeHealthPolicy"`
	MaxPercentUnhealthyDeployedApplications int                                        `pulumi:"maxPercentUnhealthyDeployedApplications"`
	ServiceTypeHealthPolicyMap              map[string]ServiceTypeHealthPolicyResponse `pulumi:"serviceTypeHealthPolicyMap"`
}





type ApplicationHealthPolicyResponseInput interface {
	pulumi.Input

	ToApplicationHealthPolicyResponseOutput() ApplicationHealthPolicyResponseOutput
	ToApplicationHealthPolicyResponseOutputWithContext(context.Context) ApplicationHealthPolicyResponseOutput
}

type ApplicationHealthPolicyResponseArgs struct {
	ConsiderWarningAsError                  pulumi.BoolInput                        `pulumi:"considerWarningAsError"`
	DefaultServiceTypeHealthPolicy          ServiceTypeHealthPolicyResponsePtrInput `pulumi:"defaultServiceTypeHealthPolicy"`
	MaxPercentUnhealthyDeployedApplications pulumi.IntInput                         `pulumi:"maxPercentUnhealthyDeployedApplications"`
	ServiceTypeHealthPolicyMap              ServiceTypeHealthPolicyResponseMapInput `pulumi:"serviceTypeHealthPolicyMap"`
}

func (ApplicationHealthPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationHealthPolicyResponse)(nil)).Elem()
}

func (i ApplicationHealthPolicyResponseArgs) ToApplicationHealthPolicyResponseOutput() ApplicationHealthPolicyResponseOutput {
	return i.ToApplicationHealthPolicyResponseOutputWithContext(context.Background())
}

func (i ApplicationHealthPolicyResponseArgs) ToApplicationHealthPolicyResponseOutputWithContext(ctx context.Context) ApplicationHealthPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationHealthPolicyResponseOutput)
}

func (i ApplicationHealthPolicyResponseArgs) ToApplicationHealthPolicyResponsePtrOutput() ApplicationHealthPolicyResponsePtrOutput {
	return i.ToApplicationHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationHealthPolicyResponseArgs) ToApplicationHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationHealthPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationHealthPolicyResponseOutput).ToApplicationHealthPolicyResponsePtrOutputWithContext(ctx)
}









type ApplicationHealthPolicyResponsePtrInput interface {
	pulumi.Input

	ToApplicationHealthPolicyResponsePtrOutput() ApplicationHealthPolicyResponsePtrOutput
	ToApplicationHealthPolicyResponsePtrOutputWithContext(context.Context) ApplicationHealthPolicyResponsePtrOutput
}

type applicationHealthPolicyResponsePtrType ApplicationHealthPolicyResponseArgs

func ApplicationHealthPolicyResponsePtr(v *ApplicationHealthPolicyResponseArgs) ApplicationHealthPolicyResponsePtrInput {
	return (*applicationHealthPolicyResponsePtrType)(v)
}

func (*applicationHealthPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationHealthPolicyResponse)(nil)).Elem()
}

func (i *applicationHealthPolicyResponsePtrType) ToApplicationHealthPolicyResponsePtrOutput() ApplicationHealthPolicyResponsePtrOutput {
	return i.ToApplicationHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *applicationHealthPolicyResponsePtrType) ToApplicationHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationHealthPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationHealthPolicyResponsePtrOutput)
}

type ApplicationHealthPolicyResponseOutput struct{ *pulumi.OutputState }

func (ApplicationHealthPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationHealthPolicyResponse)(nil)).Elem()
}

func (o ApplicationHealthPolicyResponseOutput) ToApplicationHealthPolicyResponseOutput() ApplicationHealthPolicyResponseOutput {
	return o
}

func (o ApplicationHealthPolicyResponseOutput) ToApplicationHealthPolicyResponseOutputWithContext(ctx context.Context) ApplicationHealthPolicyResponseOutput {
	return o
}

func (o ApplicationHealthPolicyResponseOutput) ToApplicationHealthPolicyResponsePtrOutput() ApplicationHealthPolicyResponsePtrOutput {
	return o.ToApplicationHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationHealthPolicyResponseOutput) ToApplicationHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationHealthPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationHealthPolicyResponse) *ApplicationHealthPolicyResponse {
		return &v
	}).(ApplicationHealthPolicyResponsePtrOutput)
}

func (o ApplicationHealthPolicyResponseOutput) ConsiderWarningAsError() pulumi.BoolOutput {
	return o.ApplyT(func(v ApplicationHealthPolicyResponse) bool { return v.ConsiderWarningAsError }).(pulumi.BoolOutput)
}

func (o ApplicationHealthPolicyResponseOutput) DefaultServiceTypeHealthPolicy() ServiceTypeHealthPolicyResponsePtrOutput {
	return o.ApplyT(func(v ApplicationHealthPolicyResponse) *ServiceTypeHealthPolicyResponse {
		return v.DefaultServiceTypeHealthPolicy
	}).(ServiceTypeHealthPolicyResponsePtrOutput)
}

func (o ApplicationHealthPolicyResponseOutput) MaxPercentUnhealthyDeployedApplications() pulumi.IntOutput {
	return o.ApplyT(func(v ApplicationHealthPolicyResponse) int { return v.MaxPercentUnhealthyDeployedApplications }).(pulumi.IntOutput)
}

func (o ApplicationHealthPolicyResponseOutput) ServiceTypeHealthPolicyMap() ServiceTypeHealthPolicyResponseMapOutput {
	return o.ApplyT(func(v ApplicationHealthPolicyResponse) map[string]ServiceTypeHealthPolicyResponse {
		return v.ServiceTypeHealthPolicyMap
	}).(ServiceTypeHealthPolicyResponseMapOutput)
}

type ApplicationHealthPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationHealthPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationHealthPolicyResponse)(nil)).Elem()
}

func (o ApplicationHealthPolicyResponsePtrOutput) ToApplicationHealthPolicyResponsePtrOutput() ApplicationHealthPolicyResponsePtrOutput {
	return o
}

func (o ApplicationHealthPolicyResponsePtrOutput) ToApplicationHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationHealthPolicyResponsePtrOutput {
	return o
}

func (o ApplicationHealthPolicyResponsePtrOutput) Elem() ApplicationHealthPolicyResponseOutput {
	return o.ApplyT(func(v *ApplicationHealthPolicyResponse) ApplicationHealthPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationHealthPolicyResponse
		return ret
	}).(ApplicationHealthPolicyResponseOutput)
}

func (o ApplicationHealthPolicyResponsePtrOutput) ConsiderWarningAsError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationHealthPolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.ConsiderWarningAsError
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationHealthPolicyResponsePtrOutput) DefaultServiceTypeHealthPolicy() ServiceTypeHealthPolicyResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationHealthPolicyResponse) *ServiceTypeHealthPolicyResponse {
		if v == nil {
			return nil
		}
		return v.DefaultServiceTypeHealthPolicy
	}).(ServiceTypeHealthPolicyResponsePtrOutput)
}

func (o ApplicationHealthPolicyResponsePtrOutput) MaxPercentUnhealthyDeployedApplications() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApplicationHealthPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return &v.MaxPercentUnhealthyDeployedApplications
	}).(pulumi.IntPtrOutput)
}

func (o ApplicationHealthPolicyResponsePtrOutput) ServiceTypeHealthPolicyMap() ServiceTypeHealthPolicyResponseMapOutput {
	return o.ApplyT(func(v *ApplicationHealthPolicyResponse) map[string]ServiceTypeHealthPolicyResponse {
		if v == nil {
			return nil
		}
		return v.ServiceTypeHealthPolicyMap
	}).(ServiceTypeHealthPolicyResponseMapOutput)
}

type ApplicationTypeVersionsCleanupPolicy struct {
	MaxUnusedVersionsToKeep int `pulumi:"maxUnusedVersionsToKeep"`
}





type ApplicationTypeVersionsCleanupPolicyInput interface {
	pulumi.Input

	ToApplicationTypeVersionsCleanupPolicyOutput() ApplicationTypeVersionsCleanupPolicyOutput
	ToApplicationTypeVersionsCleanupPolicyOutputWithContext(context.Context) ApplicationTypeVersionsCleanupPolicyOutput
}

type ApplicationTypeVersionsCleanupPolicyArgs struct {
	MaxUnusedVersionsToKeep pulumi.IntInput `pulumi:"maxUnusedVersionsToKeep"`
}

func (ApplicationTypeVersionsCleanupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationTypeVersionsCleanupPolicy)(nil)).Elem()
}

func (i ApplicationTypeVersionsCleanupPolicyArgs) ToApplicationTypeVersionsCleanupPolicyOutput() ApplicationTypeVersionsCleanupPolicyOutput {
	return i.ToApplicationTypeVersionsCleanupPolicyOutputWithContext(context.Background())
}

func (i ApplicationTypeVersionsCleanupPolicyArgs) ToApplicationTypeVersionsCleanupPolicyOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationTypeVersionsCleanupPolicyOutput)
}

func (i ApplicationTypeVersionsCleanupPolicyArgs) ToApplicationTypeVersionsCleanupPolicyPtrOutput() ApplicationTypeVersionsCleanupPolicyPtrOutput {
	return i.ToApplicationTypeVersionsCleanupPolicyPtrOutputWithContext(context.Background())
}

func (i ApplicationTypeVersionsCleanupPolicyArgs) ToApplicationTypeVersionsCleanupPolicyPtrOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationTypeVersionsCleanupPolicyOutput).ToApplicationTypeVersionsCleanupPolicyPtrOutputWithContext(ctx)
}









type ApplicationTypeVersionsCleanupPolicyPtrInput interface {
	pulumi.Input

	ToApplicationTypeVersionsCleanupPolicyPtrOutput() ApplicationTypeVersionsCleanupPolicyPtrOutput
	ToApplicationTypeVersionsCleanupPolicyPtrOutputWithContext(context.Context) ApplicationTypeVersionsCleanupPolicyPtrOutput
}

type applicationTypeVersionsCleanupPolicyPtrType ApplicationTypeVersionsCleanupPolicyArgs

func ApplicationTypeVersionsCleanupPolicyPtr(v *ApplicationTypeVersionsCleanupPolicyArgs) ApplicationTypeVersionsCleanupPolicyPtrInput {
	return (*applicationTypeVersionsCleanupPolicyPtrType)(v)
}

func (*applicationTypeVersionsCleanupPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationTypeVersionsCleanupPolicy)(nil)).Elem()
}

func (i *applicationTypeVersionsCleanupPolicyPtrType) ToApplicationTypeVersionsCleanupPolicyPtrOutput() ApplicationTypeVersionsCleanupPolicyPtrOutput {
	return i.ToApplicationTypeVersionsCleanupPolicyPtrOutputWithContext(context.Background())
}

func (i *applicationTypeVersionsCleanupPolicyPtrType) ToApplicationTypeVersionsCleanupPolicyPtrOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationTypeVersionsCleanupPolicyPtrOutput)
}

type ApplicationTypeVersionsCleanupPolicyOutput struct{ *pulumi.OutputState }

func (ApplicationTypeVersionsCleanupPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationTypeVersionsCleanupPolicy)(nil)).Elem()
}

func (o ApplicationTypeVersionsCleanupPolicyOutput) ToApplicationTypeVersionsCleanupPolicyOutput() ApplicationTypeVersionsCleanupPolicyOutput {
	return o
}

func (o ApplicationTypeVersionsCleanupPolicyOutput) ToApplicationTypeVersionsCleanupPolicyOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyOutput {
	return o
}

func (o ApplicationTypeVersionsCleanupPolicyOutput) ToApplicationTypeVersionsCleanupPolicyPtrOutput() ApplicationTypeVersionsCleanupPolicyPtrOutput {
	return o.ToApplicationTypeVersionsCleanupPolicyPtrOutputWithContext(context.Background())
}

func (o ApplicationTypeVersionsCleanupPolicyOutput) ToApplicationTypeVersionsCleanupPolicyPtrOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationTypeVersionsCleanupPolicy) *ApplicationTypeVersionsCleanupPolicy {
		return &v
	}).(ApplicationTypeVersionsCleanupPolicyPtrOutput)
}

func (o ApplicationTypeVersionsCleanupPolicyOutput) MaxUnusedVersionsToKeep() pulumi.IntOutput {
	return o.ApplyT(func(v ApplicationTypeVersionsCleanupPolicy) int { return v.MaxUnusedVersionsToKeep }).(pulumi.IntOutput)
}

type ApplicationTypeVersionsCleanupPolicyPtrOutput struct{ *pulumi.OutputState }

func (ApplicationTypeVersionsCleanupPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationTypeVersionsCleanupPolicy)(nil)).Elem()
}

func (o ApplicationTypeVersionsCleanupPolicyPtrOutput) ToApplicationTypeVersionsCleanupPolicyPtrOutput() ApplicationTypeVersionsCleanupPolicyPtrOutput {
	return o
}

func (o ApplicationTypeVersionsCleanupPolicyPtrOutput) ToApplicationTypeVersionsCleanupPolicyPtrOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyPtrOutput {
	return o
}

func (o ApplicationTypeVersionsCleanupPolicyPtrOutput) Elem() ApplicationTypeVersionsCleanupPolicyOutput {
	return o.ApplyT(func(v *ApplicationTypeVersionsCleanupPolicy) ApplicationTypeVersionsCleanupPolicy {
		if v != nil {
			return *v
		}
		var ret ApplicationTypeVersionsCleanupPolicy
		return ret
	}).(ApplicationTypeVersionsCleanupPolicyOutput)
}

func (o ApplicationTypeVersionsCleanupPolicyPtrOutput) MaxUnusedVersionsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApplicationTypeVersionsCleanupPolicy) *int {
		if v == nil {
			return nil
		}
		return &v.MaxUnusedVersionsToKeep
	}).(pulumi.IntPtrOutput)
}

type ApplicationTypeVersionsCleanupPolicyResponse struct {
	MaxUnusedVersionsToKeep int `pulumi:"maxUnusedVersionsToKeep"`
}





type ApplicationTypeVersionsCleanupPolicyResponseInput interface {
	pulumi.Input

	ToApplicationTypeVersionsCleanupPolicyResponseOutput() ApplicationTypeVersionsCleanupPolicyResponseOutput
	ToApplicationTypeVersionsCleanupPolicyResponseOutputWithContext(context.Context) ApplicationTypeVersionsCleanupPolicyResponseOutput
}

type ApplicationTypeVersionsCleanupPolicyResponseArgs struct {
	MaxUnusedVersionsToKeep pulumi.IntInput `pulumi:"maxUnusedVersionsToKeep"`
}

func (ApplicationTypeVersionsCleanupPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationTypeVersionsCleanupPolicyResponse)(nil)).Elem()
}

func (i ApplicationTypeVersionsCleanupPolicyResponseArgs) ToApplicationTypeVersionsCleanupPolicyResponseOutput() ApplicationTypeVersionsCleanupPolicyResponseOutput {
	return i.ToApplicationTypeVersionsCleanupPolicyResponseOutputWithContext(context.Background())
}

func (i ApplicationTypeVersionsCleanupPolicyResponseArgs) ToApplicationTypeVersionsCleanupPolicyResponseOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationTypeVersionsCleanupPolicyResponseOutput)
}

func (i ApplicationTypeVersionsCleanupPolicyResponseArgs) ToApplicationTypeVersionsCleanupPolicyResponsePtrOutput() ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
	return i.ToApplicationTypeVersionsCleanupPolicyResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationTypeVersionsCleanupPolicyResponseArgs) ToApplicationTypeVersionsCleanupPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationTypeVersionsCleanupPolicyResponseOutput).ToApplicationTypeVersionsCleanupPolicyResponsePtrOutputWithContext(ctx)
}









type ApplicationTypeVersionsCleanupPolicyResponsePtrInput interface {
	pulumi.Input

	ToApplicationTypeVersionsCleanupPolicyResponsePtrOutput() ApplicationTypeVersionsCleanupPolicyResponsePtrOutput
	ToApplicationTypeVersionsCleanupPolicyResponsePtrOutputWithContext(context.Context) ApplicationTypeVersionsCleanupPolicyResponsePtrOutput
}

type applicationTypeVersionsCleanupPolicyResponsePtrType ApplicationTypeVersionsCleanupPolicyResponseArgs

func ApplicationTypeVersionsCleanupPolicyResponsePtr(v *ApplicationTypeVersionsCleanupPolicyResponseArgs) ApplicationTypeVersionsCleanupPolicyResponsePtrInput {
	return (*applicationTypeVersionsCleanupPolicyResponsePtrType)(v)
}

func (*applicationTypeVersionsCleanupPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationTypeVersionsCleanupPolicyResponse)(nil)).Elem()
}

func (i *applicationTypeVersionsCleanupPolicyResponsePtrType) ToApplicationTypeVersionsCleanupPolicyResponsePtrOutput() ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
	return i.ToApplicationTypeVersionsCleanupPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *applicationTypeVersionsCleanupPolicyResponsePtrType) ToApplicationTypeVersionsCleanupPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationTypeVersionsCleanupPolicyResponsePtrOutput)
}

type ApplicationTypeVersionsCleanupPolicyResponseOutput struct{ *pulumi.OutputState }

func (ApplicationTypeVersionsCleanupPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationTypeVersionsCleanupPolicyResponse)(nil)).Elem()
}

func (o ApplicationTypeVersionsCleanupPolicyResponseOutput) ToApplicationTypeVersionsCleanupPolicyResponseOutput() ApplicationTypeVersionsCleanupPolicyResponseOutput {
	return o
}

func (o ApplicationTypeVersionsCleanupPolicyResponseOutput) ToApplicationTypeVersionsCleanupPolicyResponseOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyResponseOutput {
	return o
}

func (o ApplicationTypeVersionsCleanupPolicyResponseOutput) ToApplicationTypeVersionsCleanupPolicyResponsePtrOutput() ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
	return o.ToApplicationTypeVersionsCleanupPolicyResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationTypeVersionsCleanupPolicyResponseOutput) ToApplicationTypeVersionsCleanupPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationTypeVersionsCleanupPolicyResponse) *ApplicationTypeVersionsCleanupPolicyResponse {
		return &v
	}).(ApplicationTypeVersionsCleanupPolicyResponsePtrOutput)
}

func (o ApplicationTypeVersionsCleanupPolicyResponseOutput) MaxUnusedVersionsToKeep() pulumi.IntOutput {
	return o.ApplyT(func(v ApplicationTypeVersionsCleanupPolicyResponse) int { return v.MaxUnusedVersionsToKeep }).(pulumi.IntOutput)
}

type ApplicationTypeVersionsCleanupPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationTypeVersionsCleanupPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationTypeVersionsCleanupPolicyResponse)(nil)).Elem()
}

func (o ApplicationTypeVersionsCleanupPolicyResponsePtrOutput) ToApplicationTypeVersionsCleanupPolicyResponsePtrOutput() ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
	return o
}

func (o ApplicationTypeVersionsCleanupPolicyResponsePtrOutput) ToApplicationTypeVersionsCleanupPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationTypeVersionsCleanupPolicyResponsePtrOutput {
	return o
}

func (o ApplicationTypeVersionsCleanupPolicyResponsePtrOutput) Elem() ApplicationTypeVersionsCleanupPolicyResponseOutput {
	return o.ApplyT(func(v *ApplicationTypeVersionsCleanupPolicyResponse) ApplicationTypeVersionsCleanupPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationTypeVersionsCleanupPolicyResponse
		return ret
	}).(ApplicationTypeVersionsCleanupPolicyResponseOutput)
}

func (o ApplicationTypeVersionsCleanupPolicyResponsePtrOutput) MaxUnusedVersionsToKeep() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApplicationTypeVersionsCleanupPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return &v.MaxUnusedVersionsToKeep
	}).(pulumi.IntPtrOutput)
}

type ApplicationUpgradePolicy struct {
	ApplicationHealthPolicy        *ApplicationHealthPolicy        `pulumi:"applicationHealthPolicy"`
	ForceRestart                   *bool                           `pulumi:"forceRestart"`
	InstanceCloseDelayDuration     *float64                        `pulumi:"instanceCloseDelayDuration"`
	RecreateApplication            *bool                           `pulumi:"recreateApplication"`
	RollingUpgradeMonitoringPolicy *RollingUpgradeMonitoringPolicy `pulumi:"rollingUpgradeMonitoringPolicy"`
	UpgradeMode                    *string                         `pulumi:"upgradeMode"`
	UpgradeReplicaSetCheckTimeout  *float64                        `pulumi:"upgradeReplicaSetCheckTimeout"`
}





type ApplicationUpgradePolicyInput interface {
	pulumi.Input

	ToApplicationUpgradePolicyOutput() ApplicationUpgradePolicyOutput
	ToApplicationUpgradePolicyOutputWithContext(context.Context) ApplicationUpgradePolicyOutput
}

type ApplicationUpgradePolicyArgs struct {
	ApplicationHealthPolicy        ApplicationHealthPolicyPtrInput        `pulumi:"applicationHealthPolicy"`
	ForceRestart                   pulumi.BoolPtrInput                    `pulumi:"forceRestart"`
	InstanceCloseDelayDuration     pulumi.Float64PtrInput                 `pulumi:"instanceCloseDelayDuration"`
	RecreateApplication            pulumi.BoolPtrInput                    `pulumi:"recreateApplication"`
	RollingUpgradeMonitoringPolicy RollingUpgradeMonitoringPolicyPtrInput `pulumi:"rollingUpgradeMonitoringPolicy"`
	UpgradeMode                    pulumi.StringPtrInput                  `pulumi:"upgradeMode"`
	UpgradeReplicaSetCheckTimeout  pulumi.Float64PtrInput                 `pulumi:"upgradeReplicaSetCheckTimeout"`
}

func (ApplicationUpgradePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationUpgradePolicy)(nil)).Elem()
}

func (i ApplicationUpgradePolicyArgs) ToApplicationUpgradePolicyOutput() ApplicationUpgradePolicyOutput {
	return i.ToApplicationUpgradePolicyOutputWithContext(context.Background())
}

func (i ApplicationUpgradePolicyArgs) ToApplicationUpgradePolicyOutputWithContext(ctx context.Context) ApplicationUpgradePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUpgradePolicyOutput)
}

func (i ApplicationUpgradePolicyArgs) ToApplicationUpgradePolicyPtrOutput() ApplicationUpgradePolicyPtrOutput {
	return i.ToApplicationUpgradePolicyPtrOutputWithContext(context.Background())
}

func (i ApplicationUpgradePolicyArgs) ToApplicationUpgradePolicyPtrOutputWithContext(ctx context.Context) ApplicationUpgradePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUpgradePolicyOutput).ToApplicationUpgradePolicyPtrOutputWithContext(ctx)
}









type ApplicationUpgradePolicyPtrInput interface {
	pulumi.Input

	ToApplicationUpgradePolicyPtrOutput() ApplicationUpgradePolicyPtrOutput
	ToApplicationUpgradePolicyPtrOutputWithContext(context.Context) ApplicationUpgradePolicyPtrOutput
}

type applicationUpgradePolicyPtrType ApplicationUpgradePolicyArgs

func ApplicationUpgradePolicyPtr(v *ApplicationUpgradePolicyArgs) ApplicationUpgradePolicyPtrInput {
	return (*applicationUpgradePolicyPtrType)(v)
}

func (*applicationUpgradePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationUpgradePolicy)(nil)).Elem()
}

func (i *applicationUpgradePolicyPtrType) ToApplicationUpgradePolicyPtrOutput() ApplicationUpgradePolicyPtrOutput {
	return i.ToApplicationUpgradePolicyPtrOutputWithContext(context.Background())
}

func (i *applicationUpgradePolicyPtrType) ToApplicationUpgradePolicyPtrOutputWithContext(ctx context.Context) ApplicationUpgradePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUpgradePolicyPtrOutput)
}

type ApplicationUpgradePolicyOutput struct{ *pulumi.OutputState }

func (ApplicationUpgradePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationUpgradePolicy)(nil)).Elem()
}

func (o ApplicationUpgradePolicyOutput) ToApplicationUpgradePolicyOutput() ApplicationUpgradePolicyOutput {
	return o
}

func (o ApplicationUpgradePolicyOutput) ToApplicationUpgradePolicyOutputWithContext(ctx context.Context) ApplicationUpgradePolicyOutput {
	return o
}

func (o ApplicationUpgradePolicyOutput) ToApplicationUpgradePolicyPtrOutput() ApplicationUpgradePolicyPtrOutput {
	return o.ToApplicationUpgradePolicyPtrOutputWithContext(context.Background())
}

func (o ApplicationUpgradePolicyOutput) ToApplicationUpgradePolicyPtrOutputWithContext(ctx context.Context) ApplicationUpgradePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationUpgradePolicy) *ApplicationUpgradePolicy {
		return &v
	}).(ApplicationUpgradePolicyPtrOutput)
}

func (o ApplicationUpgradePolicyOutput) ApplicationHealthPolicy() ApplicationHealthPolicyPtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicy) *ApplicationHealthPolicy { return v.ApplicationHealthPolicy }).(ApplicationHealthPolicyPtrOutput)
}

func (o ApplicationUpgradePolicyOutput) ForceRestart() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicy) *bool { return v.ForceRestart }).(pulumi.BoolPtrOutput)
}

func (o ApplicationUpgradePolicyOutput) InstanceCloseDelayDuration() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicy) *float64 { return v.InstanceCloseDelayDuration }).(pulumi.Float64PtrOutput)
}

func (o ApplicationUpgradePolicyOutput) RecreateApplication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicy) *bool { return v.RecreateApplication }).(pulumi.BoolPtrOutput)
}

func (o ApplicationUpgradePolicyOutput) RollingUpgradeMonitoringPolicy() RollingUpgradeMonitoringPolicyPtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicy) *RollingUpgradeMonitoringPolicy {
		return v.RollingUpgradeMonitoringPolicy
	}).(RollingUpgradeMonitoringPolicyPtrOutput)
}

func (o ApplicationUpgradePolicyOutput) UpgradeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicy) *string { return v.UpgradeMode }).(pulumi.StringPtrOutput)
}

func (o ApplicationUpgradePolicyOutput) UpgradeReplicaSetCheckTimeout() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicy) *float64 { return v.UpgradeReplicaSetCheckTimeout }).(pulumi.Float64PtrOutput)
}

type ApplicationUpgradePolicyPtrOutput struct{ *pulumi.OutputState }

func (ApplicationUpgradePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationUpgradePolicy)(nil)).Elem()
}

func (o ApplicationUpgradePolicyPtrOutput) ToApplicationUpgradePolicyPtrOutput() ApplicationUpgradePolicyPtrOutput {
	return o
}

func (o ApplicationUpgradePolicyPtrOutput) ToApplicationUpgradePolicyPtrOutputWithContext(ctx context.Context) ApplicationUpgradePolicyPtrOutput {
	return o
}

func (o ApplicationUpgradePolicyPtrOutput) Elem() ApplicationUpgradePolicyOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicy) ApplicationUpgradePolicy {
		if v != nil {
			return *v
		}
		var ret ApplicationUpgradePolicy
		return ret
	}).(ApplicationUpgradePolicyOutput)
}

func (o ApplicationUpgradePolicyPtrOutput) ApplicationHealthPolicy() ApplicationHealthPolicyPtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicy) *ApplicationHealthPolicy {
		if v == nil {
			return nil
		}
		return v.ApplicationHealthPolicy
	}).(ApplicationHealthPolicyPtrOutput)
}

func (o ApplicationUpgradePolicyPtrOutput) ForceRestart() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicy) *bool {
		if v == nil {
			return nil
		}
		return v.ForceRestart
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationUpgradePolicyPtrOutput) InstanceCloseDelayDuration() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicy) *float64 {
		if v == nil {
			return nil
		}
		return v.InstanceCloseDelayDuration
	}).(pulumi.Float64PtrOutput)
}

func (o ApplicationUpgradePolicyPtrOutput) RecreateApplication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicy) *bool {
		if v == nil {
			return nil
		}
		return v.RecreateApplication
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationUpgradePolicyPtrOutput) RollingUpgradeMonitoringPolicy() RollingUpgradeMonitoringPolicyPtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicy) *RollingUpgradeMonitoringPolicy {
		if v == nil {
			return nil
		}
		return v.RollingUpgradeMonitoringPolicy
	}).(RollingUpgradeMonitoringPolicyPtrOutput)
}

func (o ApplicationUpgradePolicyPtrOutput) UpgradeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicy) *string {
		if v == nil {
			return nil
		}
		return v.UpgradeMode
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationUpgradePolicyPtrOutput) UpgradeReplicaSetCheckTimeout() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicy) *float64 {
		if v == nil {
			return nil
		}
		return v.UpgradeReplicaSetCheckTimeout
	}).(pulumi.Float64PtrOutput)
}

type ApplicationUpgradePolicyResponse struct {
	ApplicationHealthPolicy        *ApplicationHealthPolicyResponse        `pulumi:"applicationHealthPolicy"`
	ForceRestart                   *bool                                   `pulumi:"forceRestart"`
	InstanceCloseDelayDuration     *float64                                `pulumi:"instanceCloseDelayDuration"`
	RecreateApplication            *bool                                   `pulumi:"recreateApplication"`
	RollingUpgradeMonitoringPolicy *RollingUpgradeMonitoringPolicyResponse `pulumi:"rollingUpgradeMonitoringPolicy"`
	UpgradeMode                    *string                                 `pulumi:"upgradeMode"`
	UpgradeReplicaSetCheckTimeout  *float64                                `pulumi:"upgradeReplicaSetCheckTimeout"`
}





type ApplicationUpgradePolicyResponseInput interface {
	pulumi.Input

	ToApplicationUpgradePolicyResponseOutput() ApplicationUpgradePolicyResponseOutput
	ToApplicationUpgradePolicyResponseOutputWithContext(context.Context) ApplicationUpgradePolicyResponseOutput
}

type ApplicationUpgradePolicyResponseArgs struct {
	ApplicationHealthPolicy        ApplicationHealthPolicyResponsePtrInput        `pulumi:"applicationHealthPolicy"`
	ForceRestart                   pulumi.BoolPtrInput                            `pulumi:"forceRestart"`
	InstanceCloseDelayDuration     pulumi.Float64PtrInput                         `pulumi:"instanceCloseDelayDuration"`
	RecreateApplication            pulumi.BoolPtrInput                            `pulumi:"recreateApplication"`
	RollingUpgradeMonitoringPolicy RollingUpgradeMonitoringPolicyResponsePtrInput `pulumi:"rollingUpgradeMonitoringPolicy"`
	UpgradeMode                    pulumi.StringPtrInput                          `pulumi:"upgradeMode"`
	UpgradeReplicaSetCheckTimeout  pulumi.Float64PtrInput                         `pulumi:"upgradeReplicaSetCheckTimeout"`
}

func (ApplicationUpgradePolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationUpgradePolicyResponse)(nil)).Elem()
}

func (i ApplicationUpgradePolicyResponseArgs) ToApplicationUpgradePolicyResponseOutput() ApplicationUpgradePolicyResponseOutput {
	return i.ToApplicationUpgradePolicyResponseOutputWithContext(context.Background())
}

func (i ApplicationUpgradePolicyResponseArgs) ToApplicationUpgradePolicyResponseOutputWithContext(ctx context.Context) ApplicationUpgradePolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUpgradePolicyResponseOutput)
}

func (i ApplicationUpgradePolicyResponseArgs) ToApplicationUpgradePolicyResponsePtrOutput() ApplicationUpgradePolicyResponsePtrOutput {
	return i.ToApplicationUpgradePolicyResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationUpgradePolicyResponseArgs) ToApplicationUpgradePolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationUpgradePolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUpgradePolicyResponseOutput).ToApplicationUpgradePolicyResponsePtrOutputWithContext(ctx)
}









type ApplicationUpgradePolicyResponsePtrInput interface {
	pulumi.Input

	ToApplicationUpgradePolicyResponsePtrOutput() ApplicationUpgradePolicyResponsePtrOutput
	ToApplicationUpgradePolicyResponsePtrOutputWithContext(context.Context) ApplicationUpgradePolicyResponsePtrOutput
}

type applicationUpgradePolicyResponsePtrType ApplicationUpgradePolicyResponseArgs

func ApplicationUpgradePolicyResponsePtr(v *ApplicationUpgradePolicyResponseArgs) ApplicationUpgradePolicyResponsePtrInput {
	return (*applicationUpgradePolicyResponsePtrType)(v)
}

func (*applicationUpgradePolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationUpgradePolicyResponse)(nil)).Elem()
}

func (i *applicationUpgradePolicyResponsePtrType) ToApplicationUpgradePolicyResponsePtrOutput() ApplicationUpgradePolicyResponsePtrOutput {
	return i.ToApplicationUpgradePolicyResponsePtrOutputWithContext(context.Background())
}

func (i *applicationUpgradePolicyResponsePtrType) ToApplicationUpgradePolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationUpgradePolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUpgradePolicyResponsePtrOutput)
}

type ApplicationUpgradePolicyResponseOutput struct{ *pulumi.OutputState }

func (ApplicationUpgradePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationUpgradePolicyResponse)(nil)).Elem()
}

func (o ApplicationUpgradePolicyResponseOutput) ToApplicationUpgradePolicyResponseOutput() ApplicationUpgradePolicyResponseOutput {
	return o
}

func (o ApplicationUpgradePolicyResponseOutput) ToApplicationUpgradePolicyResponseOutputWithContext(ctx context.Context) ApplicationUpgradePolicyResponseOutput {
	return o
}

func (o ApplicationUpgradePolicyResponseOutput) ToApplicationUpgradePolicyResponsePtrOutput() ApplicationUpgradePolicyResponsePtrOutput {
	return o.ToApplicationUpgradePolicyResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationUpgradePolicyResponseOutput) ToApplicationUpgradePolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationUpgradePolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationUpgradePolicyResponse) *ApplicationUpgradePolicyResponse {
		return &v
	}).(ApplicationUpgradePolicyResponsePtrOutput)
}

func (o ApplicationUpgradePolicyResponseOutput) ApplicationHealthPolicy() ApplicationHealthPolicyResponsePtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicyResponse) *ApplicationHealthPolicyResponse {
		return v.ApplicationHealthPolicy
	}).(ApplicationHealthPolicyResponsePtrOutput)
}

func (o ApplicationUpgradePolicyResponseOutput) ForceRestart() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicyResponse) *bool { return v.ForceRestart }).(pulumi.BoolPtrOutput)
}

func (o ApplicationUpgradePolicyResponseOutput) InstanceCloseDelayDuration() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicyResponse) *float64 { return v.InstanceCloseDelayDuration }).(pulumi.Float64PtrOutput)
}

func (o ApplicationUpgradePolicyResponseOutput) RecreateApplication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicyResponse) *bool { return v.RecreateApplication }).(pulumi.BoolPtrOutput)
}

func (o ApplicationUpgradePolicyResponseOutput) RollingUpgradeMonitoringPolicy() RollingUpgradeMonitoringPolicyResponsePtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicyResponse) *RollingUpgradeMonitoringPolicyResponse {
		return v.RollingUpgradeMonitoringPolicy
	}).(RollingUpgradeMonitoringPolicyResponsePtrOutput)
}

func (o ApplicationUpgradePolicyResponseOutput) UpgradeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicyResponse) *string { return v.UpgradeMode }).(pulumi.StringPtrOutput)
}

func (o ApplicationUpgradePolicyResponseOutput) UpgradeReplicaSetCheckTimeout() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ApplicationUpgradePolicyResponse) *float64 { return v.UpgradeReplicaSetCheckTimeout }).(pulumi.Float64PtrOutput)
}

type ApplicationUpgradePolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationUpgradePolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationUpgradePolicyResponse)(nil)).Elem()
}

func (o ApplicationUpgradePolicyResponsePtrOutput) ToApplicationUpgradePolicyResponsePtrOutput() ApplicationUpgradePolicyResponsePtrOutput {
	return o
}

func (o ApplicationUpgradePolicyResponsePtrOutput) ToApplicationUpgradePolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationUpgradePolicyResponsePtrOutput {
	return o
}

func (o ApplicationUpgradePolicyResponsePtrOutput) Elem() ApplicationUpgradePolicyResponseOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicyResponse) ApplicationUpgradePolicyResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationUpgradePolicyResponse
		return ret
	}).(ApplicationUpgradePolicyResponseOutput)
}

func (o ApplicationUpgradePolicyResponsePtrOutput) ApplicationHealthPolicy() ApplicationHealthPolicyResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicyResponse) *ApplicationHealthPolicyResponse {
		if v == nil {
			return nil
		}
		return v.ApplicationHealthPolicy
	}).(ApplicationHealthPolicyResponsePtrOutput)
}

func (o ApplicationUpgradePolicyResponsePtrOutput) ForceRestart() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ForceRestart
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationUpgradePolicyResponsePtrOutput) InstanceCloseDelayDuration() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicyResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.InstanceCloseDelayDuration
	}).(pulumi.Float64PtrOutput)
}

func (o ApplicationUpgradePolicyResponsePtrOutput) RecreateApplication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RecreateApplication
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationUpgradePolicyResponsePtrOutput) RollingUpgradeMonitoringPolicy() RollingUpgradeMonitoringPolicyResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicyResponse) *RollingUpgradeMonitoringPolicyResponse {
		if v == nil {
			return nil
		}
		return v.RollingUpgradeMonitoringPolicy
	}).(RollingUpgradeMonitoringPolicyResponsePtrOutput)
}

func (o ApplicationUpgradePolicyResponsePtrOutput) UpgradeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.UpgradeMode
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationUpgradePolicyResponsePtrOutput) UpgradeReplicaSetCheckTimeout() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ApplicationUpgradePolicyResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.UpgradeReplicaSetCheckTimeout
	}).(pulumi.Float64PtrOutput)
}

type ApplicationUserAssignedIdentity struct {
	Name        string `pulumi:"name"`
	PrincipalId string `pulumi:"principalId"`
}





type ApplicationUserAssignedIdentityInput interface {
	pulumi.Input

	ToApplicationUserAssignedIdentityOutput() ApplicationUserAssignedIdentityOutput
	ToApplicationUserAssignedIdentityOutputWithContext(context.Context) ApplicationUserAssignedIdentityOutput
}

type ApplicationUserAssignedIdentityArgs struct {
	Name        pulumi.StringInput `pulumi:"name"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
}

func (ApplicationUserAssignedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationUserAssignedIdentity)(nil)).Elem()
}

func (i ApplicationUserAssignedIdentityArgs) ToApplicationUserAssignedIdentityOutput() ApplicationUserAssignedIdentityOutput {
	return i.ToApplicationUserAssignedIdentityOutputWithContext(context.Background())
}

func (i ApplicationUserAssignedIdentityArgs) ToApplicationUserAssignedIdentityOutputWithContext(ctx context.Context) ApplicationUserAssignedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUserAssignedIdentityOutput)
}





type ApplicationUserAssignedIdentityArrayInput interface {
	pulumi.Input

	ToApplicationUserAssignedIdentityArrayOutput() ApplicationUserAssignedIdentityArrayOutput
	ToApplicationUserAssignedIdentityArrayOutputWithContext(context.Context) ApplicationUserAssignedIdentityArrayOutput
}

type ApplicationUserAssignedIdentityArray []ApplicationUserAssignedIdentityInput

func (ApplicationUserAssignedIdentityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationUserAssignedIdentity)(nil)).Elem()
}

func (i ApplicationUserAssignedIdentityArray) ToApplicationUserAssignedIdentityArrayOutput() ApplicationUserAssignedIdentityArrayOutput {
	return i.ToApplicationUserAssignedIdentityArrayOutputWithContext(context.Background())
}

func (i ApplicationUserAssignedIdentityArray) ToApplicationUserAssignedIdentityArrayOutputWithContext(ctx context.Context) ApplicationUserAssignedIdentityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUserAssignedIdentityArrayOutput)
}

type ApplicationUserAssignedIdentityOutput struct{ *pulumi.OutputState }

func (ApplicationUserAssignedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationUserAssignedIdentity)(nil)).Elem()
}

func (o ApplicationUserAssignedIdentityOutput) ToApplicationUserAssignedIdentityOutput() ApplicationUserAssignedIdentityOutput {
	return o
}

func (o ApplicationUserAssignedIdentityOutput) ToApplicationUserAssignedIdentityOutputWithContext(ctx context.Context) ApplicationUserAssignedIdentityOutput {
	return o
}

func (o ApplicationUserAssignedIdentityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationUserAssignedIdentity) string { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationUserAssignedIdentityOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationUserAssignedIdentity) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type ApplicationUserAssignedIdentityArrayOutput struct{ *pulumi.OutputState }

func (ApplicationUserAssignedIdentityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationUserAssignedIdentity)(nil)).Elem()
}

func (o ApplicationUserAssignedIdentityArrayOutput) ToApplicationUserAssignedIdentityArrayOutput() ApplicationUserAssignedIdentityArrayOutput {
	return o
}

func (o ApplicationUserAssignedIdentityArrayOutput) ToApplicationUserAssignedIdentityArrayOutputWithContext(ctx context.Context) ApplicationUserAssignedIdentityArrayOutput {
	return o
}

func (o ApplicationUserAssignedIdentityArrayOutput) Index(i pulumi.IntInput) ApplicationUserAssignedIdentityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationUserAssignedIdentity {
		return vs[0].([]ApplicationUserAssignedIdentity)[vs[1].(int)]
	}).(ApplicationUserAssignedIdentityOutput)
}

type ApplicationUserAssignedIdentityResponse struct {
	Name        string `pulumi:"name"`
	PrincipalId string `pulumi:"principalId"`
}





type ApplicationUserAssignedIdentityResponseInput interface {
	pulumi.Input

	ToApplicationUserAssignedIdentityResponseOutput() ApplicationUserAssignedIdentityResponseOutput
	ToApplicationUserAssignedIdentityResponseOutputWithContext(context.Context) ApplicationUserAssignedIdentityResponseOutput
}

type ApplicationUserAssignedIdentityResponseArgs struct {
	Name        pulumi.StringInput `pulumi:"name"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
}

func (ApplicationUserAssignedIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationUserAssignedIdentityResponse)(nil)).Elem()
}

func (i ApplicationUserAssignedIdentityResponseArgs) ToApplicationUserAssignedIdentityResponseOutput() ApplicationUserAssignedIdentityResponseOutput {
	return i.ToApplicationUserAssignedIdentityResponseOutputWithContext(context.Background())
}

func (i ApplicationUserAssignedIdentityResponseArgs) ToApplicationUserAssignedIdentityResponseOutputWithContext(ctx context.Context) ApplicationUserAssignedIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUserAssignedIdentityResponseOutput)
}





type ApplicationUserAssignedIdentityResponseArrayInput interface {
	pulumi.Input

	ToApplicationUserAssignedIdentityResponseArrayOutput() ApplicationUserAssignedIdentityResponseArrayOutput
	ToApplicationUserAssignedIdentityResponseArrayOutputWithContext(context.Context) ApplicationUserAssignedIdentityResponseArrayOutput
}

type ApplicationUserAssignedIdentityResponseArray []ApplicationUserAssignedIdentityResponseInput

func (ApplicationUserAssignedIdentityResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationUserAssignedIdentityResponse)(nil)).Elem()
}

func (i ApplicationUserAssignedIdentityResponseArray) ToApplicationUserAssignedIdentityResponseArrayOutput() ApplicationUserAssignedIdentityResponseArrayOutput {
	return i.ToApplicationUserAssignedIdentityResponseArrayOutputWithContext(context.Background())
}

func (i ApplicationUserAssignedIdentityResponseArray) ToApplicationUserAssignedIdentityResponseArrayOutputWithContext(ctx context.Context) ApplicationUserAssignedIdentityResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationUserAssignedIdentityResponseArrayOutput)
}

type ApplicationUserAssignedIdentityResponseOutput struct{ *pulumi.OutputState }

func (ApplicationUserAssignedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationUserAssignedIdentityResponse)(nil)).Elem()
}

func (o ApplicationUserAssignedIdentityResponseOutput) ToApplicationUserAssignedIdentityResponseOutput() ApplicationUserAssignedIdentityResponseOutput {
	return o
}

func (o ApplicationUserAssignedIdentityResponseOutput) ToApplicationUserAssignedIdentityResponseOutputWithContext(ctx context.Context) ApplicationUserAssignedIdentityResponseOutput {
	return o
}

func (o ApplicationUserAssignedIdentityResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationUserAssignedIdentityResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationUserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationUserAssignedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type ApplicationUserAssignedIdentityResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationUserAssignedIdentityResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationUserAssignedIdentityResponse)(nil)).Elem()
}

func (o ApplicationUserAssignedIdentityResponseArrayOutput) ToApplicationUserAssignedIdentityResponseArrayOutput() ApplicationUserAssignedIdentityResponseArrayOutput {
	return o
}

func (o ApplicationUserAssignedIdentityResponseArrayOutput) ToApplicationUserAssignedIdentityResponseArrayOutputWithContext(ctx context.Context) ApplicationUserAssignedIdentityResponseArrayOutput {
	return o
}

func (o ApplicationUserAssignedIdentityResponseArrayOutput) Index(i pulumi.IntInput) ApplicationUserAssignedIdentityResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationUserAssignedIdentityResponse {
		return vs[0].([]ApplicationUserAssignedIdentityResponse)[vs[1].(int)]
	}).(ApplicationUserAssignedIdentityResponseOutput)
}

type AveragePartitionLoadScalingTrigger struct {
	Kind               string  `pulumi:"kind"`
	LowerLoadThreshold float64 `pulumi:"lowerLoadThreshold"`
	MetricName         string  `pulumi:"metricName"`
	ScaleInterval      string  `pulumi:"scaleInterval"`
	UpperLoadThreshold float64 `pulumi:"upperLoadThreshold"`
}





type AveragePartitionLoadScalingTriggerInput interface {
	pulumi.Input

	ToAveragePartitionLoadScalingTriggerOutput() AveragePartitionLoadScalingTriggerOutput
	ToAveragePartitionLoadScalingTriggerOutputWithContext(context.Context) AveragePartitionLoadScalingTriggerOutput
}

type AveragePartitionLoadScalingTriggerArgs struct {
	Kind               pulumi.StringInput  `pulumi:"kind"`
	LowerLoadThreshold pulumi.Float64Input `pulumi:"lowerLoadThreshold"`
	MetricName         pulumi.StringInput  `pulumi:"metricName"`
	ScaleInterval      pulumi.StringInput  `pulumi:"scaleInterval"`
	UpperLoadThreshold pulumi.Float64Input `pulumi:"upperLoadThreshold"`
}

func (AveragePartitionLoadScalingTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AveragePartitionLoadScalingTrigger)(nil)).Elem()
}

func (i AveragePartitionLoadScalingTriggerArgs) ToAveragePartitionLoadScalingTriggerOutput() AveragePartitionLoadScalingTriggerOutput {
	return i.ToAveragePartitionLoadScalingTriggerOutputWithContext(context.Background())
}

func (i AveragePartitionLoadScalingTriggerArgs) ToAveragePartitionLoadScalingTriggerOutputWithContext(ctx context.Context) AveragePartitionLoadScalingTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AveragePartitionLoadScalingTriggerOutput)
}

type AveragePartitionLoadScalingTriggerOutput struct{ *pulumi.OutputState }

func (AveragePartitionLoadScalingTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AveragePartitionLoadScalingTrigger)(nil)).Elem()
}

func (o AveragePartitionLoadScalingTriggerOutput) ToAveragePartitionLoadScalingTriggerOutput() AveragePartitionLoadScalingTriggerOutput {
	return o
}

func (o AveragePartitionLoadScalingTriggerOutput) ToAveragePartitionLoadScalingTriggerOutputWithContext(ctx context.Context) AveragePartitionLoadScalingTriggerOutput {
	return o
}

func (o AveragePartitionLoadScalingTriggerOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v AveragePartitionLoadScalingTrigger) string { return v.Kind }).(pulumi.StringOutput)
}

func (o AveragePartitionLoadScalingTriggerOutput) LowerLoadThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v AveragePartitionLoadScalingTrigger) float64 { return v.LowerLoadThreshold }).(pulumi.Float64Output)
}

func (o AveragePartitionLoadScalingTriggerOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v AveragePartitionLoadScalingTrigger) string { return v.MetricName }).(pulumi.StringOutput)
}

func (o AveragePartitionLoadScalingTriggerOutput) ScaleInterval() pulumi.StringOutput {
	return o.ApplyT(func(v AveragePartitionLoadScalingTrigger) string { return v.ScaleInterval }).(pulumi.StringOutput)
}

func (o AveragePartitionLoadScalingTriggerOutput) UpperLoadThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v AveragePartitionLoadScalingTrigger) float64 { return v.UpperLoadThreshold }).(pulumi.Float64Output)
}

type AveragePartitionLoadScalingTriggerResponse struct {
	Kind               string  `pulumi:"kind"`
	LowerLoadThreshold float64 `pulumi:"lowerLoadThreshold"`
	MetricName         string  `pulumi:"metricName"`
	ScaleInterval      string  `pulumi:"scaleInterval"`
	UpperLoadThreshold float64 `pulumi:"upperLoadThreshold"`
}





type AveragePartitionLoadScalingTriggerResponseInput interface {
	pulumi.Input

	ToAveragePartitionLoadScalingTriggerResponseOutput() AveragePartitionLoadScalingTriggerResponseOutput
	ToAveragePartitionLoadScalingTriggerResponseOutputWithContext(context.Context) AveragePartitionLoadScalingTriggerResponseOutput
}

type AveragePartitionLoadScalingTriggerResponseArgs struct {
	Kind               pulumi.StringInput  `pulumi:"kind"`
	LowerLoadThreshold pulumi.Float64Input `pulumi:"lowerLoadThreshold"`
	MetricName         pulumi.StringInput  `pulumi:"metricName"`
	ScaleInterval      pulumi.StringInput  `pulumi:"scaleInterval"`
	UpperLoadThreshold pulumi.Float64Input `pulumi:"upperLoadThreshold"`
}

func (AveragePartitionLoadScalingTriggerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AveragePartitionLoadScalingTriggerResponse)(nil)).Elem()
}

func (i AveragePartitionLoadScalingTriggerResponseArgs) ToAveragePartitionLoadScalingTriggerResponseOutput() AveragePartitionLoadScalingTriggerResponseOutput {
	return i.ToAveragePartitionLoadScalingTriggerResponseOutputWithContext(context.Background())
}

func (i AveragePartitionLoadScalingTriggerResponseArgs) ToAveragePartitionLoadScalingTriggerResponseOutputWithContext(ctx context.Context) AveragePartitionLoadScalingTriggerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AveragePartitionLoadScalingTriggerResponseOutput)
}

type AveragePartitionLoadScalingTriggerResponseOutput struct{ *pulumi.OutputState }

func (AveragePartitionLoadScalingTriggerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AveragePartitionLoadScalingTriggerResponse)(nil)).Elem()
}

func (o AveragePartitionLoadScalingTriggerResponseOutput) ToAveragePartitionLoadScalingTriggerResponseOutput() AveragePartitionLoadScalingTriggerResponseOutput {
	return o
}

func (o AveragePartitionLoadScalingTriggerResponseOutput) ToAveragePartitionLoadScalingTriggerResponseOutputWithContext(ctx context.Context) AveragePartitionLoadScalingTriggerResponseOutput {
	return o
}

func (o AveragePartitionLoadScalingTriggerResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v AveragePartitionLoadScalingTriggerResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o AveragePartitionLoadScalingTriggerResponseOutput) LowerLoadThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v AveragePartitionLoadScalingTriggerResponse) float64 { return v.LowerLoadThreshold }).(pulumi.Float64Output)
}

func (o AveragePartitionLoadScalingTriggerResponseOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v AveragePartitionLoadScalingTriggerResponse) string { return v.MetricName }).(pulumi.StringOutput)
}

func (o AveragePartitionLoadScalingTriggerResponseOutput) ScaleInterval() pulumi.StringOutput {
	return o.ApplyT(func(v AveragePartitionLoadScalingTriggerResponse) string { return v.ScaleInterval }).(pulumi.StringOutput)
}

func (o AveragePartitionLoadScalingTriggerResponseOutput) UpperLoadThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v AveragePartitionLoadScalingTriggerResponse) float64 { return v.UpperLoadThreshold }).(pulumi.Float64Output)
}

type AverageServiceLoadScalingTrigger struct {
	Kind               string  `pulumi:"kind"`
	LowerLoadThreshold float64 `pulumi:"lowerLoadThreshold"`
	MetricName         string  `pulumi:"metricName"`
	ScaleInterval      string  `pulumi:"scaleInterval"`
	UpperLoadThreshold float64 `pulumi:"upperLoadThreshold"`
	UseOnlyPrimaryLoad bool    `pulumi:"useOnlyPrimaryLoad"`
}





type AverageServiceLoadScalingTriggerInput interface {
	pulumi.Input

	ToAverageServiceLoadScalingTriggerOutput() AverageServiceLoadScalingTriggerOutput
	ToAverageServiceLoadScalingTriggerOutputWithContext(context.Context) AverageServiceLoadScalingTriggerOutput
}

type AverageServiceLoadScalingTriggerArgs struct {
	Kind               pulumi.StringInput  `pulumi:"kind"`
	LowerLoadThreshold pulumi.Float64Input `pulumi:"lowerLoadThreshold"`
	MetricName         pulumi.StringInput  `pulumi:"metricName"`
	ScaleInterval      pulumi.StringInput  `pulumi:"scaleInterval"`
	UpperLoadThreshold pulumi.Float64Input `pulumi:"upperLoadThreshold"`
	UseOnlyPrimaryLoad pulumi.BoolInput    `pulumi:"useOnlyPrimaryLoad"`
}

func (AverageServiceLoadScalingTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AverageServiceLoadScalingTrigger)(nil)).Elem()
}

func (i AverageServiceLoadScalingTriggerArgs) ToAverageServiceLoadScalingTriggerOutput() AverageServiceLoadScalingTriggerOutput {
	return i.ToAverageServiceLoadScalingTriggerOutputWithContext(context.Background())
}

func (i AverageServiceLoadScalingTriggerArgs) ToAverageServiceLoadScalingTriggerOutputWithContext(ctx context.Context) AverageServiceLoadScalingTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AverageServiceLoadScalingTriggerOutput)
}

type AverageServiceLoadScalingTriggerOutput struct{ *pulumi.OutputState }

func (AverageServiceLoadScalingTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AverageServiceLoadScalingTrigger)(nil)).Elem()
}

func (o AverageServiceLoadScalingTriggerOutput) ToAverageServiceLoadScalingTriggerOutput() AverageServiceLoadScalingTriggerOutput {
	return o
}

func (o AverageServiceLoadScalingTriggerOutput) ToAverageServiceLoadScalingTriggerOutputWithContext(ctx context.Context) AverageServiceLoadScalingTriggerOutput {
	return o
}

func (o AverageServiceLoadScalingTriggerOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v AverageServiceLoadScalingTrigger) string { return v.Kind }).(pulumi.StringOutput)
}

func (o AverageServiceLoadScalingTriggerOutput) LowerLoadThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v AverageServiceLoadScalingTrigger) float64 { return v.LowerLoadThreshold }).(pulumi.Float64Output)
}

func (o AverageServiceLoadScalingTriggerOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v AverageServiceLoadScalingTrigger) string { return v.MetricName }).(pulumi.StringOutput)
}

func (o AverageServiceLoadScalingTriggerOutput) ScaleInterval() pulumi.StringOutput {
	return o.ApplyT(func(v AverageServiceLoadScalingTrigger) string { return v.ScaleInterval }).(pulumi.StringOutput)
}

func (o AverageServiceLoadScalingTriggerOutput) UpperLoadThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v AverageServiceLoadScalingTrigger) float64 { return v.UpperLoadThreshold }).(pulumi.Float64Output)
}

func (o AverageServiceLoadScalingTriggerOutput) UseOnlyPrimaryLoad() pulumi.BoolOutput {
	return o.ApplyT(func(v AverageServiceLoadScalingTrigger) bool { return v.UseOnlyPrimaryLoad }).(pulumi.BoolOutput)
}

type AverageServiceLoadScalingTriggerResponse struct {
	Kind               string  `pulumi:"kind"`
	LowerLoadThreshold float64 `pulumi:"lowerLoadThreshold"`
	MetricName         string  `pulumi:"metricName"`
	ScaleInterval      string  `pulumi:"scaleInterval"`
	UpperLoadThreshold float64 `pulumi:"upperLoadThreshold"`
	UseOnlyPrimaryLoad bool    `pulumi:"useOnlyPrimaryLoad"`
}





type AverageServiceLoadScalingTriggerResponseInput interface {
	pulumi.Input

	ToAverageServiceLoadScalingTriggerResponseOutput() AverageServiceLoadScalingTriggerResponseOutput
	ToAverageServiceLoadScalingTriggerResponseOutputWithContext(context.Context) AverageServiceLoadScalingTriggerResponseOutput
}

type AverageServiceLoadScalingTriggerResponseArgs struct {
	Kind               pulumi.StringInput  `pulumi:"kind"`
	LowerLoadThreshold pulumi.Float64Input `pulumi:"lowerLoadThreshold"`
	MetricName         pulumi.StringInput  `pulumi:"metricName"`
	ScaleInterval      pulumi.StringInput  `pulumi:"scaleInterval"`
	UpperLoadThreshold pulumi.Float64Input `pulumi:"upperLoadThreshold"`
	UseOnlyPrimaryLoad pulumi.BoolInput    `pulumi:"useOnlyPrimaryLoad"`
}

func (AverageServiceLoadScalingTriggerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AverageServiceLoadScalingTriggerResponse)(nil)).Elem()
}

func (i AverageServiceLoadScalingTriggerResponseArgs) ToAverageServiceLoadScalingTriggerResponseOutput() AverageServiceLoadScalingTriggerResponseOutput {
	return i.ToAverageServiceLoadScalingTriggerResponseOutputWithContext(context.Background())
}

func (i AverageServiceLoadScalingTriggerResponseArgs) ToAverageServiceLoadScalingTriggerResponseOutputWithContext(ctx context.Context) AverageServiceLoadScalingTriggerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AverageServiceLoadScalingTriggerResponseOutput)
}

type AverageServiceLoadScalingTriggerResponseOutput struct{ *pulumi.OutputState }

func (AverageServiceLoadScalingTriggerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AverageServiceLoadScalingTriggerResponse)(nil)).Elem()
}

func (o AverageServiceLoadScalingTriggerResponseOutput) ToAverageServiceLoadScalingTriggerResponseOutput() AverageServiceLoadScalingTriggerResponseOutput {
	return o
}

func (o AverageServiceLoadScalingTriggerResponseOutput) ToAverageServiceLoadScalingTriggerResponseOutputWithContext(ctx context.Context) AverageServiceLoadScalingTriggerResponseOutput {
	return o
}

func (o AverageServiceLoadScalingTriggerResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v AverageServiceLoadScalingTriggerResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o AverageServiceLoadScalingTriggerResponseOutput) LowerLoadThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v AverageServiceLoadScalingTriggerResponse) float64 { return v.LowerLoadThreshold }).(pulumi.Float64Output)
}

func (o AverageServiceLoadScalingTriggerResponseOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v AverageServiceLoadScalingTriggerResponse) string { return v.MetricName }).(pulumi.StringOutput)
}

func (o AverageServiceLoadScalingTriggerResponseOutput) ScaleInterval() pulumi.StringOutput {
	return o.ApplyT(func(v AverageServiceLoadScalingTriggerResponse) string { return v.ScaleInterval }).(pulumi.StringOutput)
}

func (o AverageServiceLoadScalingTriggerResponseOutput) UpperLoadThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v AverageServiceLoadScalingTriggerResponse) float64 { return v.UpperLoadThreshold }).(pulumi.Float64Output)
}

func (o AverageServiceLoadScalingTriggerResponseOutput) UseOnlyPrimaryLoad() pulumi.BoolOutput {
	return o.ApplyT(func(v AverageServiceLoadScalingTriggerResponse) bool { return v.UseOnlyPrimaryLoad }).(pulumi.BoolOutput)
}

type AzureActiveDirectory struct {
	ClientApplication  *string `pulumi:"clientApplication"`
	ClusterApplication *string `pulumi:"clusterApplication"`
	TenantId           *string `pulumi:"tenantId"`
}





type AzureActiveDirectoryInput interface {
	pulumi.Input

	ToAzureActiveDirectoryOutput() AzureActiveDirectoryOutput
	ToAzureActiveDirectoryOutputWithContext(context.Context) AzureActiveDirectoryOutput
}

type AzureActiveDirectoryArgs struct {
	ClientApplication  pulumi.StringPtrInput `pulumi:"clientApplication"`
	ClusterApplication pulumi.StringPtrInput `pulumi:"clusterApplication"`
	TenantId           pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (AzureActiveDirectoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectory)(nil)).Elem()
}

func (i AzureActiveDirectoryArgs) ToAzureActiveDirectoryOutput() AzureActiveDirectoryOutput {
	return i.ToAzureActiveDirectoryOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryArgs) ToAzureActiveDirectoryOutputWithContext(ctx context.Context) AzureActiveDirectoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryOutput)
}

func (i AzureActiveDirectoryArgs) ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput {
	return i.ToAzureActiveDirectoryPtrOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryArgs) ToAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryOutput).ToAzureActiveDirectoryPtrOutputWithContext(ctx)
}









type AzureActiveDirectoryPtrInput interface {
	pulumi.Input

	ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput
	ToAzureActiveDirectoryPtrOutputWithContext(context.Context) AzureActiveDirectoryPtrOutput
}

type azureActiveDirectoryPtrType AzureActiveDirectoryArgs

func AzureActiveDirectoryPtr(v *AzureActiveDirectoryArgs) AzureActiveDirectoryPtrInput {
	return (*azureActiveDirectoryPtrType)(v)
}

func (*azureActiveDirectoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectory)(nil)).Elem()
}

func (i *azureActiveDirectoryPtrType) ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput {
	return i.ToAzureActiveDirectoryPtrOutputWithContext(context.Background())
}

func (i *azureActiveDirectoryPtrType) ToAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryPtrOutput)
}

type AzureActiveDirectoryOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectory)(nil)).Elem()
}

func (o AzureActiveDirectoryOutput) ToAzureActiveDirectoryOutput() AzureActiveDirectoryOutput {
	return o
}

func (o AzureActiveDirectoryOutput) ToAzureActiveDirectoryOutputWithContext(ctx context.Context) AzureActiveDirectoryOutput {
	return o
}

func (o AzureActiveDirectoryOutput) ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput {
	return o.ToAzureActiveDirectoryPtrOutputWithContext(context.Background())
}

func (o AzureActiveDirectoryOutput) ToAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureActiveDirectory) *AzureActiveDirectory {
		return &v
	}).(AzureActiveDirectoryPtrOutput)
}

func (o AzureActiveDirectoryOutput) ClientApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectory) *string { return v.ClientApplication }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryOutput) ClusterApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectory) *string { return v.ClusterApplication }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectory) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type AzureActiveDirectoryPtrOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectory)(nil)).Elem()
}

func (o AzureActiveDirectoryPtrOutput) ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput {
	return o
}

func (o AzureActiveDirectoryPtrOutput) ToAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryPtrOutput {
	return o
}

func (o AzureActiveDirectoryPtrOutput) Elem() AzureActiveDirectoryOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) AzureActiveDirectory {
		if v != nil {
			return *v
		}
		var ret AzureActiveDirectory
		return ret
	}).(AzureActiveDirectoryOutput)
}

func (o AzureActiveDirectoryPtrOutput) ClientApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) *string {
		if v == nil {
			return nil
		}
		return v.ClientApplication
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryPtrOutput) ClusterApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) *string {
		if v == nil {
			return nil
		}
		return v.ClusterApplication
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type AzureActiveDirectoryResponse struct {
	ClientApplication  *string `pulumi:"clientApplication"`
	ClusterApplication *string `pulumi:"clusterApplication"`
	TenantId           *string `pulumi:"tenantId"`
}





type AzureActiveDirectoryResponseInput interface {
	pulumi.Input

	ToAzureActiveDirectoryResponseOutput() AzureActiveDirectoryResponseOutput
	ToAzureActiveDirectoryResponseOutputWithContext(context.Context) AzureActiveDirectoryResponseOutput
}

type AzureActiveDirectoryResponseArgs struct {
	ClientApplication  pulumi.StringPtrInput `pulumi:"clientApplication"`
	ClusterApplication pulumi.StringPtrInput `pulumi:"clusterApplication"`
	TenantId           pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (AzureActiveDirectoryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryResponse)(nil)).Elem()
}

func (i AzureActiveDirectoryResponseArgs) ToAzureActiveDirectoryResponseOutput() AzureActiveDirectoryResponseOutput {
	return i.ToAzureActiveDirectoryResponseOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryResponseArgs) ToAzureActiveDirectoryResponseOutputWithContext(ctx context.Context) AzureActiveDirectoryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryResponseOutput)
}

func (i AzureActiveDirectoryResponseArgs) ToAzureActiveDirectoryResponsePtrOutput() AzureActiveDirectoryResponsePtrOutput {
	return i.ToAzureActiveDirectoryResponsePtrOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryResponseArgs) ToAzureActiveDirectoryResponsePtrOutputWithContext(ctx context.Context) AzureActiveDirectoryResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryResponseOutput).ToAzureActiveDirectoryResponsePtrOutputWithContext(ctx)
}









type AzureActiveDirectoryResponsePtrInput interface {
	pulumi.Input

	ToAzureActiveDirectoryResponsePtrOutput() AzureActiveDirectoryResponsePtrOutput
	ToAzureActiveDirectoryResponsePtrOutputWithContext(context.Context) AzureActiveDirectoryResponsePtrOutput
}

type azureActiveDirectoryResponsePtrType AzureActiveDirectoryResponseArgs

func AzureActiveDirectoryResponsePtr(v *AzureActiveDirectoryResponseArgs) AzureActiveDirectoryResponsePtrInput {
	return (*azureActiveDirectoryResponsePtrType)(v)
}

func (*azureActiveDirectoryResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryResponse)(nil)).Elem()
}

func (i *azureActiveDirectoryResponsePtrType) ToAzureActiveDirectoryResponsePtrOutput() AzureActiveDirectoryResponsePtrOutput {
	return i.ToAzureActiveDirectoryResponsePtrOutputWithContext(context.Background())
}

func (i *azureActiveDirectoryResponsePtrType) ToAzureActiveDirectoryResponsePtrOutputWithContext(ctx context.Context) AzureActiveDirectoryResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryResponsePtrOutput)
}

type AzureActiveDirectoryResponseOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryResponse)(nil)).Elem()
}

func (o AzureActiveDirectoryResponseOutput) ToAzureActiveDirectoryResponseOutput() AzureActiveDirectoryResponseOutput {
	return o
}

func (o AzureActiveDirectoryResponseOutput) ToAzureActiveDirectoryResponseOutputWithContext(ctx context.Context) AzureActiveDirectoryResponseOutput {
	return o
}

func (o AzureActiveDirectoryResponseOutput) ToAzureActiveDirectoryResponsePtrOutput() AzureActiveDirectoryResponsePtrOutput {
	return o.ToAzureActiveDirectoryResponsePtrOutputWithContext(context.Background())
}

func (o AzureActiveDirectoryResponseOutput) ToAzureActiveDirectoryResponsePtrOutputWithContext(ctx context.Context) AzureActiveDirectoryResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureActiveDirectoryResponse) *AzureActiveDirectoryResponse {
		return &v
	}).(AzureActiveDirectoryResponsePtrOutput)
}

func (o AzureActiveDirectoryResponseOutput) ClientApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryResponse) *string { return v.ClientApplication }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryResponseOutput) ClusterApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryResponse) *string { return v.ClusterApplication }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type AzureActiveDirectoryResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryResponse)(nil)).Elem()
}

func (o AzureActiveDirectoryResponsePtrOutput) ToAzureActiveDirectoryResponsePtrOutput() AzureActiveDirectoryResponsePtrOutput {
	return o
}

func (o AzureActiveDirectoryResponsePtrOutput) ToAzureActiveDirectoryResponsePtrOutputWithContext(ctx context.Context) AzureActiveDirectoryResponsePtrOutput {
	return o
}

func (o AzureActiveDirectoryResponsePtrOutput) Elem() AzureActiveDirectoryResponseOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) AzureActiveDirectoryResponse {
		if v != nil {
			return *v
		}
		var ret AzureActiveDirectoryResponse
		return ret
	}).(AzureActiveDirectoryResponseOutput)
}

func (o AzureActiveDirectoryResponsePtrOutput) ClientApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientApplication
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryResponsePtrOutput) ClusterApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClusterApplication
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type ClientCertificate struct {
	CommonName       *string `pulumi:"commonName"`
	IsAdmin          bool    `pulumi:"isAdmin"`
	IssuerThumbprint *string `pulumi:"issuerThumbprint"`
	Thumbprint       *string `pulumi:"thumbprint"`
}





type ClientCertificateInput interface {
	pulumi.Input

	ToClientCertificateOutput() ClientCertificateOutput
	ToClientCertificateOutputWithContext(context.Context) ClientCertificateOutput
}

type ClientCertificateArgs struct {
	CommonName       pulumi.StringPtrInput `pulumi:"commonName"`
	IsAdmin          pulumi.BoolInput      `pulumi:"isAdmin"`
	IssuerThumbprint pulumi.StringPtrInput `pulumi:"issuerThumbprint"`
	Thumbprint       pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (ClientCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificate)(nil)).Elem()
}

func (i ClientCertificateArgs) ToClientCertificateOutput() ClientCertificateOutput {
	return i.ToClientCertificateOutputWithContext(context.Background())
}

func (i ClientCertificateArgs) ToClientCertificateOutputWithContext(ctx context.Context) ClientCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateOutput)
}





type ClientCertificateArrayInput interface {
	pulumi.Input

	ToClientCertificateArrayOutput() ClientCertificateArrayOutput
	ToClientCertificateArrayOutputWithContext(context.Context) ClientCertificateArrayOutput
}

type ClientCertificateArray []ClientCertificateInput

func (ClientCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificate)(nil)).Elem()
}

func (i ClientCertificateArray) ToClientCertificateArrayOutput() ClientCertificateArrayOutput {
	return i.ToClientCertificateArrayOutputWithContext(context.Background())
}

func (i ClientCertificateArray) ToClientCertificateArrayOutputWithContext(ctx context.Context) ClientCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateArrayOutput)
}

type ClientCertificateOutput struct{ *pulumi.OutputState }

func (ClientCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificate)(nil)).Elem()
}

func (o ClientCertificateOutput) ToClientCertificateOutput() ClientCertificateOutput {
	return o
}

func (o ClientCertificateOutput) ToClientCertificateOutputWithContext(ctx context.Context) ClientCertificateOutput {
	return o
}

func (o ClientCertificateOutput) CommonName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCertificate) *string { return v.CommonName }).(pulumi.StringPtrOutput)
}

func (o ClientCertificateOutput) IsAdmin() pulumi.BoolOutput {
	return o.ApplyT(func(v ClientCertificate) bool { return v.IsAdmin }).(pulumi.BoolOutput)
}

func (o ClientCertificateOutput) IssuerThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCertificate) *string { return v.IssuerThumbprint }).(pulumi.StringPtrOutput)
}

func (o ClientCertificateOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCertificate) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type ClientCertificateArrayOutput struct{ *pulumi.OutputState }

func (ClientCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificate)(nil)).Elem()
}

func (o ClientCertificateArrayOutput) ToClientCertificateArrayOutput() ClientCertificateArrayOutput {
	return o
}

func (o ClientCertificateArrayOutput) ToClientCertificateArrayOutputWithContext(ctx context.Context) ClientCertificateArrayOutput {
	return o
}

func (o ClientCertificateArrayOutput) Index(i pulumi.IntInput) ClientCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientCertificate {
		return vs[0].([]ClientCertificate)[vs[1].(int)]
	}).(ClientCertificateOutput)
}

type ClientCertificateResponse struct {
	CommonName       *string `pulumi:"commonName"`
	IsAdmin          bool    `pulumi:"isAdmin"`
	IssuerThumbprint *string `pulumi:"issuerThumbprint"`
	Thumbprint       *string `pulumi:"thumbprint"`
}





type ClientCertificateResponseInput interface {
	pulumi.Input

	ToClientCertificateResponseOutput() ClientCertificateResponseOutput
	ToClientCertificateResponseOutputWithContext(context.Context) ClientCertificateResponseOutput
}

type ClientCertificateResponseArgs struct {
	CommonName       pulumi.StringPtrInput `pulumi:"commonName"`
	IsAdmin          pulumi.BoolInput      `pulumi:"isAdmin"`
	IssuerThumbprint pulumi.StringPtrInput `pulumi:"issuerThumbprint"`
	Thumbprint       pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (ClientCertificateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificateResponse)(nil)).Elem()
}

func (i ClientCertificateResponseArgs) ToClientCertificateResponseOutput() ClientCertificateResponseOutput {
	return i.ToClientCertificateResponseOutputWithContext(context.Background())
}

func (i ClientCertificateResponseArgs) ToClientCertificateResponseOutputWithContext(ctx context.Context) ClientCertificateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateResponseOutput)
}





type ClientCertificateResponseArrayInput interface {
	pulumi.Input

	ToClientCertificateResponseArrayOutput() ClientCertificateResponseArrayOutput
	ToClientCertificateResponseArrayOutputWithContext(context.Context) ClientCertificateResponseArrayOutput
}

type ClientCertificateResponseArray []ClientCertificateResponseInput

func (ClientCertificateResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificateResponse)(nil)).Elem()
}

func (i ClientCertificateResponseArray) ToClientCertificateResponseArrayOutput() ClientCertificateResponseArrayOutput {
	return i.ToClientCertificateResponseArrayOutputWithContext(context.Background())
}

func (i ClientCertificateResponseArray) ToClientCertificateResponseArrayOutputWithContext(ctx context.Context) ClientCertificateResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateResponseArrayOutput)
}

type ClientCertificateResponseOutput struct{ *pulumi.OutputState }

func (ClientCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificateResponse)(nil)).Elem()
}

func (o ClientCertificateResponseOutput) ToClientCertificateResponseOutput() ClientCertificateResponseOutput {
	return o
}

func (o ClientCertificateResponseOutput) ToClientCertificateResponseOutputWithContext(ctx context.Context) ClientCertificateResponseOutput {
	return o
}

func (o ClientCertificateResponseOutput) CommonName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCertificateResponse) *string { return v.CommonName }).(pulumi.StringPtrOutput)
}

func (o ClientCertificateResponseOutput) IsAdmin() pulumi.BoolOutput {
	return o.ApplyT(func(v ClientCertificateResponse) bool { return v.IsAdmin }).(pulumi.BoolOutput)
}

func (o ClientCertificateResponseOutput) IssuerThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCertificateResponse) *string { return v.IssuerThumbprint }).(pulumi.StringPtrOutput)
}

func (o ClientCertificateResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCertificateResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type ClientCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (ClientCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificateResponse)(nil)).Elem()
}

func (o ClientCertificateResponseArrayOutput) ToClientCertificateResponseArrayOutput() ClientCertificateResponseArrayOutput {
	return o
}

func (o ClientCertificateResponseArrayOutput) ToClientCertificateResponseArrayOutputWithContext(ctx context.Context) ClientCertificateResponseArrayOutput {
	return o
}

func (o ClientCertificateResponseArrayOutput) Index(i pulumi.IntInput) ClientCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientCertificateResponse {
		return vs[0].([]ClientCertificateResponse)[vs[1].(int)]
	}).(ClientCertificateResponseOutput)
}

type EndpointRangeDescription struct {
	EndPort   int `pulumi:"endPort"`
	StartPort int `pulumi:"startPort"`
}





type EndpointRangeDescriptionInput interface {
	pulumi.Input

	ToEndpointRangeDescriptionOutput() EndpointRangeDescriptionOutput
	ToEndpointRangeDescriptionOutputWithContext(context.Context) EndpointRangeDescriptionOutput
}

type EndpointRangeDescriptionArgs struct {
	EndPort   pulumi.IntInput `pulumi:"endPort"`
	StartPort pulumi.IntInput `pulumi:"startPort"`
}

func (EndpointRangeDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointRangeDescription)(nil)).Elem()
}

func (i EndpointRangeDescriptionArgs) ToEndpointRangeDescriptionOutput() EndpointRangeDescriptionOutput {
	return i.ToEndpointRangeDescriptionOutputWithContext(context.Background())
}

func (i EndpointRangeDescriptionArgs) ToEndpointRangeDescriptionOutputWithContext(ctx context.Context) EndpointRangeDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointRangeDescriptionOutput)
}

func (i EndpointRangeDescriptionArgs) ToEndpointRangeDescriptionPtrOutput() EndpointRangeDescriptionPtrOutput {
	return i.ToEndpointRangeDescriptionPtrOutputWithContext(context.Background())
}

func (i EndpointRangeDescriptionArgs) ToEndpointRangeDescriptionPtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointRangeDescriptionOutput).ToEndpointRangeDescriptionPtrOutputWithContext(ctx)
}









type EndpointRangeDescriptionPtrInput interface {
	pulumi.Input

	ToEndpointRangeDescriptionPtrOutput() EndpointRangeDescriptionPtrOutput
	ToEndpointRangeDescriptionPtrOutputWithContext(context.Context) EndpointRangeDescriptionPtrOutput
}

type endpointRangeDescriptionPtrType EndpointRangeDescriptionArgs

func EndpointRangeDescriptionPtr(v *EndpointRangeDescriptionArgs) EndpointRangeDescriptionPtrInput {
	return (*endpointRangeDescriptionPtrType)(v)
}

func (*endpointRangeDescriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointRangeDescription)(nil)).Elem()
}

func (i *endpointRangeDescriptionPtrType) ToEndpointRangeDescriptionPtrOutput() EndpointRangeDescriptionPtrOutput {
	return i.ToEndpointRangeDescriptionPtrOutputWithContext(context.Background())
}

func (i *endpointRangeDescriptionPtrType) ToEndpointRangeDescriptionPtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointRangeDescriptionPtrOutput)
}

type EndpointRangeDescriptionOutput struct{ *pulumi.OutputState }

func (EndpointRangeDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointRangeDescription)(nil)).Elem()
}

func (o EndpointRangeDescriptionOutput) ToEndpointRangeDescriptionOutput() EndpointRangeDescriptionOutput {
	return o
}

func (o EndpointRangeDescriptionOutput) ToEndpointRangeDescriptionOutputWithContext(ctx context.Context) EndpointRangeDescriptionOutput {
	return o
}

func (o EndpointRangeDescriptionOutput) ToEndpointRangeDescriptionPtrOutput() EndpointRangeDescriptionPtrOutput {
	return o.ToEndpointRangeDescriptionPtrOutputWithContext(context.Background())
}

func (o EndpointRangeDescriptionOutput) ToEndpointRangeDescriptionPtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointRangeDescription) *EndpointRangeDescription {
		return &v
	}).(EndpointRangeDescriptionPtrOutput)
}

func (o EndpointRangeDescriptionOutput) EndPort() pulumi.IntOutput {
	return o.ApplyT(func(v EndpointRangeDescription) int { return v.EndPort }).(pulumi.IntOutput)
}

func (o EndpointRangeDescriptionOutput) StartPort() pulumi.IntOutput {
	return o.ApplyT(func(v EndpointRangeDescription) int { return v.StartPort }).(pulumi.IntOutput)
}

type EndpointRangeDescriptionPtrOutput struct{ *pulumi.OutputState }

func (EndpointRangeDescriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointRangeDescription)(nil)).Elem()
}

func (o EndpointRangeDescriptionPtrOutput) ToEndpointRangeDescriptionPtrOutput() EndpointRangeDescriptionPtrOutput {
	return o
}

func (o EndpointRangeDescriptionPtrOutput) ToEndpointRangeDescriptionPtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionPtrOutput {
	return o
}

func (o EndpointRangeDescriptionPtrOutput) Elem() EndpointRangeDescriptionOutput {
	return o.ApplyT(func(v *EndpointRangeDescription) EndpointRangeDescription {
		if v != nil {
			return *v
		}
		var ret EndpointRangeDescription
		return ret
	}).(EndpointRangeDescriptionOutput)
}

func (o EndpointRangeDescriptionPtrOutput) EndPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EndpointRangeDescription) *int {
		if v == nil {
			return nil
		}
		return &v.EndPort
	}).(pulumi.IntPtrOutput)
}

func (o EndpointRangeDescriptionPtrOutput) StartPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EndpointRangeDescription) *int {
		if v == nil {
			return nil
		}
		return &v.StartPort
	}).(pulumi.IntPtrOutput)
}

type EndpointRangeDescriptionResponse struct {
	EndPort   int `pulumi:"endPort"`
	StartPort int `pulumi:"startPort"`
}





type EndpointRangeDescriptionResponseInput interface {
	pulumi.Input

	ToEndpointRangeDescriptionResponseOutput() EndpointRangeDescriptionResponseOutput
	ToEndpointRangeDescriptionResponseOutputWithContext(context.Context) EndpointRangeDescriptionResponseOutput
}

type EndpointRangeDescriptionResponseArgs struct {
	EndPort   pulumi.IntInput `pulumi:"endPort"`
	StartPort pulumi.IntInput `pulumi:"startPort"`
}

func (EndpointRangeDescriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointRangeDescriptionResponse)(nil)).Elem()
}

func (i EndpointRangeDescriptionResponseArgs) ToEndpointRangeDescriptionResponseOutput() EndpointRangeDescriptionResponseOutput {
	return i.ToEndpointRangeDescriptionResponseOutputWithContext(context.Background())
}

func (i EndpointRangeDescriptionResponseArgs) ToEndpointRangeDescriptionResponseOutputWithContext(ctx context.Context) EndpointRangeDescriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointRangeDescriptionResponseOutput)
}

func (i EndpointRangeDescriptionResponseArgs) ToEndpointRangeDescriptionResponsePtrOutput() EndpointRangeDescriptionResponsePtrOutput {
	return i.ToEndpointRangeDescriptionResponsePtrOutputWithContext(context.Background())
}

func (i EndpointRangeDescriptionResponseArgs) ToEndpointRangeDescriptionResponsePtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointRangeDescriptionResponseOutput).ToEndpointRangeDescriptionResponsePtrOutputWithContext(ctx)
}









type EndpointRangeDescriptionResponsePtrInput interface {
	pulumi.Input

	ToEndpointRangeDescriptionResponsePtrOutput() EndpointRangeDescriptionResponsePtrOutput
	ToEndpointRangeDescriptionResponsePtrOutputWithContext(context.Context) EndpointRangeDescriptionResponsePtrOutput
}

type endpointRangeDescriptionResponsePtrType EndpointRangeDescriptionResponseArgs

func EndpointRangeDescriptionResponsePtr(v *EndpointRangeDescriptionResponseArgs) EndpointRangeDescriptionResponsePtrInput {
	return (*endpointRangeDescriptionResponsePtrType)(v)
}

func (*endpointRangeDescriptionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointRangeDescriptionResponse)(nil)).Elem()
}

func (i *endpointRangeDescriptionResponsePtrType) ToEndpointRangeDescriptionResponsePtrOutput() EndpointRangeDescriptionResponsePtrOutput {
	return i.ToEndpointRangeDescriptionResponsePtrOutputWithContext(context.Background())
}

func (i *endpointRangeDescriptionResponsePtrType) ToEndpointRangeDescriptionResponsePtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointRangeDescriptionResponsePtrOutput)
}

type EndpointRangeDescriptionResponseOutput struct{ *pulumi.OutputState }

func (EndpointRangeDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointRangeDescriptionResponse)(nil)).Elem()
}

func (o EndpointRangeDescriptionResponseOutput) ToEndpointRangeDescriptionResponseOutput() EndpointRangeDescriptionResponseOutput {
	return o
}

func (o EndpointRangeDescriptionResponseOutput) ToEndpointRangeDescriptionResponseOutputWithContext(ctx context.Context) EndpointRangeDescriptionResponseOutput {
	return o
}

func (o EndpointRangeDescriptionResponseOutput) ToEndpointRangeDescriptionResponsePtrOutput() EndpointRangeDescriptionResponsePtrOutput {
	return o.ToEndpointRangeDescriptionResponsePtrOutputWithContext(context.Background())
}

func (o EndpointRangeDescriptionResponseOutput) ToEndpointRangeDescriptionResponsePtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointRangeDescriptionResponse) *EndpointRangeDescriptionResponse {
		return &v
	}).(EndpointRangeDescriptionResponsePtrOutput)
}

func (o EndpointRangeDescriptionResponseOutput) EndPort() pulumi.IntOutput {
	return o.ApplyT(func(v EndpointRangeDescriptionResponse) int { return v.EndPort }).(pulumi.IntOutput)
}

func (o EndpointRangeDescriptionResponseOutput) StartPort() pulumi.IntOutput {
	return o.ApplyT(func(v EndpointRangeDescriptionResponse) int { return v.StartPort }).(pulumi.IntOutput)
}

type EndpointRangeDescriptionResponsePtrOutput struct{ *pulumi.OutputState }

func (EndpointRangeDescriptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointRangeDescriptionResponse)(nil)).Elem()
}

func (o EndpointRangeDescriptionResponsePtrOutput) ToEndpointRangeDescriptionResponsePtrOutput() EndpointRangeDescriptionResponsePtrOutput {
	return o
}

func (o EndpointRangeDescriptionResponsePtrOutput) ToEndpointRangeDescriptionResponsePtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionResponsePtrOutput {
	return o
}

func (o EndpointRangeDescriptionResponsePtrOutput) Elem() EndpointRangeDescriptionResponseOutput {
	return o.ApplyT(func(v *EndpointRangeDescriptionResponse) EndpointRangeDescriptionResponse {
		if v != nil {
			return *v
		}
		var ret EndpointRangeDescriptionResponse
		return ret
	}).(EndpointRangeDescriptionResponseOutput)
}

func (o EndpointRangeDescriptionResponsePtrOutput) EndPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EndpointRangeDescriptionResponse) *int {
		if v == nil {
			return nil
		}
		return &v.EndPort
	}).(pulumi.IntPtrOutput)
}

func (o EndpointRangeDescriptionResponsePtrOutput) StartPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EndpointRangeDescriptionResponse) *int {
		if v == nil {
			return nil
		}
		return &v.StartPort
	}).(pulumi.IntPtrOutput)
}

type LoadBalancingRule struct {
	BackendPort      int     `pulumi:"backendPort"`
	FrontendPort     int     `pulumi:"frontendPort"`
	ProbeProtocol    string  `pulumi:"probeProtocol"`
	ProbeRequestPath *string `pulumi:"probeRequestPath"`
	Protocol         string  `pulumi:"protocol"`
}





type LoadBalancingRuleInput interface {
	pulumi.Input

	ToLoadBalancingRuleOutput() LoadBalancingRuleOutput
	ToLoadBalancingRuleOutputWithContext(context.Context) LoadBalancingRuleOutput
}

type LoadBalancingRuleArgs struct {
	BackendPort      pulumi.IntInput       `pulumi:"backendPort"`
	FrontendPort     pulumi.IntInput       `pulumi:"frontendPort"`
	ProbeProtocol    pulumi.StringInput    `pulumi:"probeProtocol"`
	ProbeRequestPath pulumi.StringPtrInput `pulumi:"probeRequestPath"`
	Protocol         pulumi.StringInput    `pulumi:"protocol"`
}

func (LoadBalancingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingRule)(nil)).Elem()
}

func (i LoadBalancingRuleArgs) ToLoadBalancingRuleOutput() LoadBalancingRuleOutput {
	return i.ToLoadBalancingRuleOutputWithContext(context.Background())
}

func (i LoadBalancingRuleArgs) ToLoadBalancingRuleOutputWithContext(ctx context.Context) LoadBalancingRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancingRuleOutput)
}





type LoadBalancingRuleArrayInput interface {
	pulumi.Input

	ToLoadBalancingRuleArrayOutput() LoadBalancingRuleArrayOutput
	ToLoadBalancingRuleArrayOutputWithContext(context.Context) LoadBalancingRuleArrayOutput
}

type LoadBalancingRuleArray []LoadBalancingRuleInput

func (LoadBalancingRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingRule)(nil)).Elem()
}

func (i LoadBalancingRuleArray) ToLoadBalancingRuleArrayOutput() LoadBalancingRuleArrayOutput {
	return i.ToLoadBalancingRuleArrayOutputWithContext(context.Background())
}

func (i LoadBalancingRuleArray) ToLoadBalancingRuleArrayOutputWithContext(ctx context.Context) LoadBalancingRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancingRuleArrayOutput)
}

type LoadBalancingRuleOutput struct{ *pulumi.OutputState }

func (LoadBalancingRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingRule)(nil)).Elem()
}

func (o LoadBalancingRuleOutput) ToLoadBalancingRuleOutput() LoadBalancingRuleOutput {
	return o
}

func (o LoadBalancingRuleOutput) ToLoadBalancingRuleOutputWithContext(ctx context.Context) LoadBalancingRuleOutput {
	return o
}

func (o LoadBalancingRuleOutput) BackendPort() pulumi.IntOutput {
	return o.ApplyT(func(v LoadBalancingRule) int { return v.BackendPort }).(pulumi.IntOutput)
}

func (o LoadBalancingRuleOutput) FrontendPort() pulumi.IntOutput {
	return o.ApplyT(func(v LoadBalancingRule) int { return v.FrontendPort }).(pulumi.IntOutput)
}

func (o LoadBalancingRuleOutput) ProbeProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancingRule) string { return v.ProbeProtocol }).(pulumi.StringOutput)
}

func (o LoadBalancingRuleOutput) ProbeRequestPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *string { return v.ProbeRequestPath }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancingRule) string { return v.Protocol }).(pulumi.StringOutput)
}

type LoadBalancingRuleArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancingRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingRule)(nil)).Elem()
}

func (o LoadBalancingRuleArrayOutput) ToLoadBalancingRuleArrayOutput() LoadBalancingRuleArrayOutput {
	return o
}

func (o LoadBalancingRuleArrayOutput) ToLoadBalancingRuleArrayOutputWithContext(ctx context.Context) LoadBalancingRuleArrayOutput {
	return o
}

func (o LoadBalancingRuleArrayOutput) Index(i pulumi.IntInput) LoadBalancingRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancingRule {
		return vs[0].([]LoadBalancingRule)[vs[1].(int)]
	}).(LoadBalancingRuleOutput)
}

type LoadBalancingRuleResponse struct {
	BackendPort      int     `pulumi:"backendPort"`
	FrontendPort     int     `pulumi:"frontendPort"`
	ProbeProtocol    string  `pulumi:"probeProtocol"`
	ProbeRequestPath *string `pulumi:"probeRequestPath"`
	Protocol         string  `pulumi:"protocol"`
}





type LoadBalancingRuleResponseInput interface {
	pulumi.Input

	ToLoadBalancingRuleResponseOutput() LoadBalancingRuleResponseOutput
	ToLoadBalancingRuleResponseOutputWithContext(context.Context) LoadBalancingRuleResponseOutput
}

type LoadBalancingRuleResponseArgs struct {
	BackendPort      pulumi.IntInput       `pulumi:"backendPort"`
	FrontendPort     pulumi.IntInput       `pulumi:"frontendPort"`
	ProbeProtocol    pulumi.StringInput    `pulumi:"probeProtocol"`
	ProbeRequestPath pulumi.StringPtrInput `pulumi:"probeRequestPath"`
	Protocol         pulumi.StringInput    `pulumi:"protocol"`
}

func (LoadBalancingRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingRuleResponse)(nil)).Elem()
}

func (i LoadBalancingRuleResponseArgs) ToLoadBalancingRuleResponseOutput() LoadBalancingRuleResponseOutput {
	return i.ToLoadBalancingRuleResponseOutputWithContext(context.Background())
}

func (i LoadBalancingRuleResponseArgs) ToLoadBalancingRuleResponseOutputWithContext(ctx context.Context) LoadBalancingRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancingRuleResponseOutput)
}





type LoadBalancingRuleResponseArrayInput interface {
	pulumi.Input

	ToLoadBalancingRuleResponseArrayOutput() LoadBalancingRuleResponseArrayOutput
	ToLoadBalancingRuleResponseArrayOutputWithContext(context.Context) LoadBalancingRuleResponseArrayOutput
}

type LoadBalancingRuleResponseArray []LoadBalancingRuleResponseInput

func (LoadBalancingRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingRuleResponse)(nil)).Elem()
}

func (i LoadBalancingRuleResponseArray) ToLoadBalancingRuleResponseArrayOutput() LoadBalancingRuleResponseArrayOutput {
	return i.ToLoadBalancingRuleResponseArrayOutputWithContext(context.Background())
}

func (i LoadBalancingRuleResponseArray) ToLoadBalancingRuleResponseArrayOutputWithContext(ctx context.Context) LoadBalancingRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancingRuleResponseArrayOutput)
}

type LoadBalancingRuleResponseOutput struct{ *pulumi.OutputState }

func (LoadBalancingRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingRuleResponse)(nil)).Elem()
}

func (o LoadBalancingRuleResponseOutput) ToLoadBalancingRuleResponseOutput() LoadBalancingRuleResponseOutput {
	return o
}

func (o LoadBalancingRuleResponseOutput) ToLoadBalancingRuleResponseOutputWithContext(ctx context.Context) LoadBalancingRuleResponseOutput {
	return o
}

func (o LoadBalancingRuleResponseOutput) BackendPort() pulumi.IntOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) int { return v.BackendPort }).(pulumi.IntOutput)
}

func (o LoadBalancingRuleResponseOutput) FrontendPort() pulumi.IntOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) int { return v.FrontendPort }).(pulumi.IntOutput)
}

func (o LoadBalancingRuleResponseOutput) ProbeProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) string { return v.ProbeProtocol }).(pulumi.StringOutput)
}

func (o LoadBalancingRuleResponseOutput) ProbeRequestPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *string { return v.ProbeRequestPath }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

type LoadBalancingRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancingRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingRuleResponse)(nil)).Elem()
}

func (o LoadBalancingRuleResponseArrayOutput) ToLoadBalancingRuleResponseArrayOutput() LoadBalancingRuleResponseArrayOutput {
	return o
}

func (o LoadBalancingRuleResponseArrayOutput) ToLoadBalancingRuleResponseArrayOutputWithContext(ctx context.Context) LoadBalancingRuleResponseArrayOutput {
	return o
}

func (o LoadBalancingRuleResponseArrayOutput) Index(i pulumi.IntInput) LoadBalancingRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancingRuleResponse {
		return vs[0].([]LoadBalancingRuleResponse)[vs[1].(int)]
	}).(LoadBalancingRuleResponseOutput)
}

type ManagedIdentity struct {
	Type                   *ManagedIdentityType   `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ManagedIdentityInput interface {
	pulumi.Input

	ToManagedIdentityOutput() ManagedIdentityOutput
	ToManagedIdentityOutputWithContext(context.Context) ManagedIdentityOutput
}

type ManagedIdentityArgs struct {
	Type                   ManagedIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput             `pulumi:"userAssignedIdentities"`
}

func (ManagedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentity)(nil)).Elem()
}

func (i ManagedIdentityArgs) ToManagedIdentityOutput() ManagedIdentityOutput {
	return i.ToManagedIdentityOutputWithContext(context.Background())
}

func (i ManagedIdentityArgs) ToManagedIdentityOutputWithContext(ctx context.Context) ManagedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityOutput)
}

func (i ManagedIdentityArgs) ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput {
	return i.ToManagedIdentityPtrOutputWithContext(context.Background())
}

func (i ManagedIdentityArgs) ToManagedIdentityPtrOutputWithContext(ctx context.Context) ManagedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityOutput).ToManagedIdentityPtrOutputWithContext(ctx)
}









type ManagedIdentityPtrInput interface {
	pulumi.Input

	ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput
	ToManagedIdentityPtrOutputWithContext(context.Context) ManagedIdentityPtrOutput
}

type managedIdentityPtrType ManagedIdentityArgs

func ManagedIdentityPtr(v *ManagedIdentityArgs) ManagedIdentityPtrInput {
	return (*managedIdentityPtrType)(v)
}

func (*managedIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentity)(nil)).Elem()
}

func (i *managedIdentityPtrType) ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput {
	return i.ToManagedIdentityPtrOutputWithContext(context.Background())
}

func (i *managedIdentityPtrType) ToManagedIdentityPtrOutputWithContext(ctx context.Context) ManagedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityPtrOutput)
}

type ManagedIdentityOutput struct{ *pulumi.OutputState }

func (ManagedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentity)(nil)).Elem()
}

func (o ManagedIdentityOutput) ToManagedIdentityOutput() ManagedIdentityOutput {
	return o
}

func (o ManagedIdentityOutput) ToManagedIdentityOutputWithContext(ctx context.Context) ManagedIdentityOutput {
	return o
}

func (o ManagedIdentityOutput) ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput {
	return o.ToManagedIdentityPtrOutputWithContext(context.Background())
}

func (o ManagedIdentityOutput) ToManagedIdentityPtrOutputWithContext(ctx context.Context) ManagedIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedIdentity) *ManagedIdentity {
		return &v
	}).(ManagedIdentityPtrOutput)
}

func (o ManagedIdentityOutput) Type() ManagedIdentityTypePtrOutput {
	return o.ApplyT(func(v ManagedIdentity) *ManagedIdentityType { return v.Type }).(ManagedIdentityTypePtrOutput)
}

func (o ManagedIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ManagedIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ManagedIdentityPtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentity)(nil)).Elem()
}

func (o ManagedIdentityPtrOutput) ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput {
	return o
}

func (o ManagedIdentityPtrOutput) ToManagedIdentityPtrOutputWithContext(ctx context.Context) ManagedIdentityPtrOutput {
	return o
}

func (o ManagedIdentityPtrOutput) Elem() ManagedIdentityOutput {
	return o.ApplyT(func(v *ManagedIdentity) ManagedIdentity {
		if v != nil {
			return *v
		}
		var ret ManagedIdentity
		return ret
	}).(ManagedIdentityOutput)
}

func (o ManagedIdentityPtrOutput) Type() ManagedIdentityTypePtrOutput {
	return o.ApplyT(func(v *ManagedIdentity) *ManagedIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ManagedIdentityTypePtrOutput)
}

func (o ManagedIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ManagedIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ManagedIdentityResponse struct {
	PrincipalId            string                                  `pulumi:"principalId"`
	TenantId               string                                  `pulumi:"tenantId"`
	Type                   *string                                 `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}





type ManagedIdentityResponseInput interface {
	pulumi.Input

	ToManagedIdentityResponseOutput() ManagedIdentityResponseOutput
	ToManagedIdentityResponseOutputWithContext(context.Context) ManagedIdentityResponseOutput
}

type ManagedIdentityResponseArgs struct {
	PrincipalId            pulumi.StringInput                   `pulumi:"principalId"`
	TenantId               pulumi.StringInput                   `pulumi:"tenantId"`
	Type                   pulumi.StringPtrInput                `pulumi:"type"`
	UserAssignedIdentities UserAssignedIdentityResponseMapInput `pulumi:"userAssignedIdentities"`
}

func (ManagedIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityResponse)(nil)).Elem()
}

func (i ManagedIdentityResponseArgs) ToManagedIdentityResponseOutput() ManagedIdentityResponseOutput {
	return i.ToManagedIdentityResponseOutputWithContext(context.Background())
}

func (i ManagedIdentityResponseArgs) ToManagedIdentityResponseOutputWithContext(ctx context.Context) ManagedIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityResponseOutput)
}

func (i ManagedIdentityResponseArgs) ToManagedIdentityResponsePtrOutput() ManagedIdentityResponsePtrOutput {
	return i.ToManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ManagedIdentityResponseArgs) ToManagedIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityResponseOutput).ToManagedIdentityResponsePtrOutputWithContext(ctx)
}









type ManagedIdentityResponsePtrInput interface {
	pulumi.Input

	ToManagedIdentityResponsePtrOutput() ManagedIdentityResponsePtrOutput
	ToManagedIdentityResponsePtrOutputWithContext(context.Context) ManagedIdentityResponsePtrOutput
}

type managedIdentityResponsePtrType ManagedIdentityResponseArgs

func ManagedIdentityResponsePtr(v *ManagedIdentityResponseArgs) ManagedIdentityResponsePtrInput {
	return (*managedIdentityResponsePtrType)(v)
}

func (*managedIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityResponse)(nil)).Elem()
}

func (i *managedIdentityResponsePtrType) ToManagedIdentityResponsePtrOutput() ManagedIdentityResponsePtrOutput {
	return i.ToManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *managedIdentityResponsePtrType) ToManagedIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityResponsePtrOutput)
}

type ManagedIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityResponse)(nil)).Elem()
}

func (o ManagedIdentityResponseOutput) ToManagedIdentityResponseOutput() ManagedIdentityResponseOutput {
	return o
}

func (o ManagedIdentityResponseOutput) ToManagedIdentityResponseOutputWithContext(ctx context.Context) ManagedIdentityResponseOutput {
	return o
}

func (o ManagedIdentityResponseOutput) ToManagedIdentityResponsePtrOutput() ManagedIdentityResponsePtrOutput {
	return o.ToManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ManagedIdentityResponseOutput) ToManagedIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedIdentityResponse) *ManagedIdentityResponse {
		return &v
	}).(ManagedIdentityResponsePtrOutput)
}

func (o ManagedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ManagedIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v ManagedIdentityResponse) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type ManagedIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityResponse)(nil)).Elem()
}

func (o ManagedIdentityResponsePtrOutput) ToManagedIdentityResponsePtrOutput() ManagedIdentityResponsePtrOutput {
	return o
}

func (o ManagedIdentityResponsePtrOutput) ToManagedIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityResponsePtrOutput {
	return o
}

func (o ManagedIdentityResponsePtrOutput) Elem() ManagedIdentityResponseOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) ManagedIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ManagedIdentityResponse
		return ret
	}).(ManagedIdentityResponseOutput)
}

func (o ManagedIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type NamedPartitionScheme struct {
	Names           []string `pulumi:"names"`
	PartitionScheme string   `pulumi:"partitionScheme"`
}





type NamedPartitionSchemeInput interface {
	pulumi.Input

	ToNamedPartitionSchemeOutput() NamedPartitionSchemeOutput
	ToNamedPartitionSchemeOutputWithContext(context.Context) NamedPartitionSchemeOutput
}

type NamedPartitionSchemeArgs struct {
	Names           pulumi.StringArrayInput `pulumi:"names"`
	PartitionScheme pulumi.StringInput      `pulumi:"partitionScheme"`
}

func (NamedPartitionSchemeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NamedPartitionScheme)(nil)).Elem()
}

func (i NamedPartitionSchemeArgs) ToNamedPartitionSchemeOutput() NamedPartitionSchemeOutput {
	return i.ToNamedPartitionSchemeOutputWithContext(context.Background())
}

func (i NamedPartitionSchemeArgs) ToNamedPartitionSchemeOutputWithContext(ctx context.Context) NamedPartitionSchemeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedPartitionSchemeOutput)
}

type NamedPartitionSchemeOutput struct{ *pulumi.OutputState }

func (NamedPartitionSchemeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamedPartitionScheme)(nil)).Elem()
}

func (o NamedPartitionSchemeOutput) ToNamedPartitionSchemeOutput() NamedPartitionSchemeOutput {
	return o
}

func (o NamedPartitionSchemeOutput) ToNamedPartitionSchemeOutputWithContext(ctx context.Context) NamedPartitionSchemeOutput {
	return o
}

func (o NamedPartitionSchemeOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NamedPartitionScheme) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o NamedPartitionSchemeOutput) PartitionScheme() pulumi.StringOutput {
	return o.ApplyT(func(v NamedPartitionScheme) string { return v.PartitionScheme }).(pulumi.StringOutput)
}

type NamedPartitionSchemeResponse struct {
	Names           []string `pulumi:"names"`
	PartitionScheme string   `pulumi:"partitionScheme"`
}





type NamedPartitionSchemeResponseInput interface {
	pulumi.Input

	ToNamedPartitionSchemeResponseOutput() NamedPartitionSchemeResponseOutput
	ToNamedPartitionSchemeResponseOutputWithContext(context.Context) NamedPartitionSchemeResponseOutput
}

type NamedPartitionSchemeResponseArgs struct {
	Names           pulumi.StringArrayInput `pulumi:"names"`
	PartitionScheme pulumi.StringInput      `pulumi:"partitionScheme"`
}

func (NamedPartitionSchemeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NamedPartitionSchemeResponse)(nil)).Elem()
}

func (i NamedPartitionSchemeResponseArgs) ToNamedPartitionSchemeResponseOutput() NamedPartitionSchemeResponseOutput {
	return i.ToNamedPartitionSchemeResponseOutputWithContext(context.Background())
}

func (i NamedPartitionSchemeResponseArgs) ToNamedPartitionSchemeResponseOutputWithContext(ctx context.Context) NamedPartitionSchemeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamedPartitionSchemeResponseOutput)
}

type NamedPartitionSchemeResponseOutput struct{ *pulumi.OutputState }

func (NamedPartitionSchemeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamedPartitionSchemeResponse)(nil)).Elem()
}

func (o NamedPartitionSchemeResponseOutput) ToNamedPartitionSchemeResponseOutput() NamedPartitionSchemeResponseOutput {
	return o
}

func (o NamedPartitionSchemeResponseOutput) ToNamedPartitionSchemeResponseOutputWithContext(ctx context.Context) NamedPartitionSchemeResponseOutput {
	return o
}

func (o NamedPartitionSchemeResponseOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NamedPartitionSchemeResponse) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o NamedPartitionSchemeResponseOutput) PartitionScheme() pulumi.StringOutput {
	return o.ApplyT(func(v NamedPartitionSchemeResponse) string { return v.PartitionScheme }).(pulumi.StringOutput)
}

type NetworkSecurityRule struct {
	Access                     string   `pulumi:"access"`
	Description                *string  `pulumi:"description"`
	DestinationAddressPrefixes []string `pulumi:"destinationAddressPrefixes"`
	DestinationPortRanges      []string `pulumi:"destinationPortRanges"`
	Direction                  string   `pulumi:"direction"`
	Name                       string   `pulumi:"name"`
	Priority                   int      `pulumi:"priority"`
	Protocol                   string   `pulumi:"protocol"`
	SourceAddressPrefixes      []string `pulumi:"sourceAddressPrefixes"`
	SourcePortRanges           []string `pulumi:"sourcePortRanges"`
}





type NetworkSecurityRuleInput interface {
	pulumi.Input

	ToNetworkSecurityRuleOutput() NetworkSecurityRuleOutput
	ToNetworkSecurityRuleOutputWithContext(context.Context) NetworkSecurityRuleOutput
}

type NetworkSecurityRuleArgs struct {
	Access                     pulumi.StringInput      `pulumi:"access"`
	Description                pulumi.StringPtrInput   `pulumi:"description"`
	DestinationAddressPrefixes pulumi.StringArrayInput `pulumi:"destinationAddressPrefixes"`
	DestinationPortRanges      pulumi.StringArrayInput `pulumi:"destinationPortRanges"`
	Direction                  pulumi.StringInput      `pulumi:"direction"`
	Name                       pulumi.StringInput      `pulumi:"name"`
	Priority                   pulumi.IntInput         `pulumi:"priority"`
	Protocol                   pulumi.StringInput      `pulumi:"protocol"`
	SourceAddressPrefixes      pulumi.StringArrayInput `pulumi:"sourceAddressPrefixes"`
	SourcePortRanges           pulumi.StringArrayInput `pulumi:"sourcePortRanges"`
}

func (NetworkSecurityRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityRule)(nil)).Elem()
}

func (i NetworkSecurityRuleArgs) ToNetworkSecurityRuleOutput() NetworkSecurityRuleOutput {
	return i.ToNetworkSecurityRuleOutputWithContext(context.Background())
}

func (i NetworkSecurityRuleArgs) ToNetworkSecurityRuleOutputWithContext(ctx context.Context) NetworkSecurityRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityRuleOutput)
}





type NetworkSecurityRuleArrayInput interface {
	pulumi.Input

	ToNetworkSecurityRuleArrayOutput() NetworkSecurityRuleArrayOutput
	ToNetworkSecurityRuleArrayOutputWithContext(context.Context) NetworkSecurityRuleArrayOutput
}

type NetworkSecurityRuleArray []NetworkSecurityRuleInput

func (NetworkSecurityRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkSecurityRule)(nil)).Elem()
}

func (i NetworkSecurityRuleArray) ToNetworkSecurityRuleArrayOutput() NetworkSecurityRuleArrayOutput {
	return i.ToNetworkSecurityRuleArrayOutputWithContext(context.Background())
}

func (i NetworkSecurityRuleArray) ToNetworkSecurityRuleArrayOutputWithContext(ctx context.Context) NetworkSecurityRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityRuleArrayOutput)
}

type NetworkSecurityRuleOutput struct{ *pulumi.OutputState }

func (NetworkSecurityRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityRule)(nil)).Elem()
}

func (o NetworkSecurityRuleOutput) ToNetworkSecurityRuleOutput() NetworkSecurityRuleOutput {
	return o
}

func (o NetworkSecurityRuleOutput) ToNetworkSecurityRuleOutputWithContext(ctx context.Context) NetworkSecurityRuleOutput {
	return o
}

func (o NetworkSecurityRuleOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkSecurityRule) string { return v.Access }).(pulumi.StringOutput)
}

func (o NetworkSecurityRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityRule) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityRuleOutput) DestinationAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkSecurityRule) []string { return v.DestinationAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o NetworkSecurityRuleOutput) DestinationPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkSecurityRule) []string { return v.DestinationPortRanges }).(pulumi.StringArrayOutput)
}

func (o NetworkSecurityRuleOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkSecurityRule) string { return v.Direction }).(pulumi.StringOutput)
}

func (o NetworkSecurityRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkSecurityRule) string { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkSecurityRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v NetworkSecurityRule) int { return v.Priority }).(pulumi.IntOutput)
}

func (o NetworkSecurityRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkSecurityRule) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o NetworkSecurityRuleOutput) SourceAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkSecurityRule) []string { return v.SourceAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o NetworkSecurityRuleOutput) SourcePortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkSecurityRule) []string { return v.SourcePortRanges }).(pulumi.StringArrayOutput)
}

type NetworkSecurityRuleArrayOutput struct{ *pulumi.OutputState }

func (NetworkSecurityRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkSecurityRule)(nil)).Elem()
}

func (o NetworkSecurityRuleArrayOutput) ToNetworkSecurityRuleArrayOutput() NetworkSecurityRuleArrayOutput {
	return o
}

func (o NetworkSecurityRuleArrayOutput) ToNetworkSecurityRuleArrayOutputWithContext(ctx context.Context) NetworkSecurityRuleArrayOutput {
	return o
}

func (o NetworkSecurityRuleArrayOutput) Index(i pulumi.IntInput) NetworkSecurityRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkSecurityRule {
		return vs[0].([]NetworkSecurityRule)[vs[1].(int)]
	}).(NetworkSecurityRuleOutput)
}

type NetworkSecurityRuleResponse struct {
	Access                     string   `pulumi:"access"`
	Description                *string  `pulumi:"description"`
	DestinationAddressPrefixes []string `pulumi:"destinationAddressPrefixes"`
	DestinationPortRanges      []string `pulumi:"destinationPortRanges"`
	Direction                  string   `pulumi:"direction"`
	Name                       string   `pulumi:"name"`
	Priority                   int      `pulumi:"priority"`
	Protocol                   string   `pulumi:"protocol"`
	SourceAddressPrefixes      []string `pulumi:"sourceAddressPrefixes"`
	SourcePortRanges           []string `pulumi:"sourcePortRanges"`
}





type NetworkSecurityRuleResponseInput interface {
	pulumi.Input

	ToNetworkSecurityRuleResponseOutput() NetworkSecurityRuleResponseOutput
	ToNetworkSecurityRuleResponseOutputWithContext(context.Context) NetworkSecurityRuleResponseOutput
}

type NetworkSecurityRuleResponseArgs struct {
	Access                     pulumi.StringInput      `pulumi:"access"`
	Description                pulumi.StringPtrInput   `pulumi:"description"`
	DestinationAddressPrefixes pulumi.StringArrayInput `pulumi:"destinationAddressPrefixes"`
	DestinationPortRanges      pulumi.StringArrayInput `pulumi:"destinationPortRanges"`
	Direction                  pulumi.StringInput      `pulumi:"direction"`
	Name                       pulumi.StringInput      `pulumi:"name"`
	Priority                   pulumi.IntInput         `pulumi:"priority"`
	Protocol                   pulumi.StringInput      `pulumi:"protocol"`
	SourceAddressPrefixes      pulumi.StringArrayInput `pulumi:"sourceAddressPrefixes"`
	SourcePortRanges           pulumi.StringArrayInput `pulumi:"sourcePortRanges"`
}

func (NetworkSecurityRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityRuleResponse)(nil)).Elem()
}

func (i NetworkSecurityRuleResponseArgs) ToNetworkSecurityRuleResponseOutput() NetworkSecurityRuleResponseOutput {
	return i.ToNetworkSecurityRuleResponseOutputWithContext(context.Background())
}

func (i NetworkSecurityRuleResponseArgs) ToNetworkSecurityRuleResponseOutputWithContext(ctx context.Context) NetworkSecurityRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityRuleResponseOutput)
}





type NetworkSecurityRuleResponseArrayInput interface {
	pulumi.Input

	ToNetworkSecurityRuleResponseArrayOutput() NetworkSecurityRuleResponseArrayOutput
	ToNetworkSecurityRuleResponseArrayOutputWithContext(context.Context) NetworkSecurityRuleResponseArrayOutput
}

type NetworkSecurityRuleResponseArray []NetworkSecurityRuleResponseInput

func (NetworkSecurityRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkSecurityRuleResponse)(nil)).Elem()
}

func (i NetworkSecurityRuleResponseArray) ToNetworkSecurityRuleResponseArrayOutput() NetworkSecurityRuleResponseArrayOutput {
	return i.ToNetworkSecurityRuleResponseArrayOutputWithContext(context.Background())
}

func (i NetworkSecurityRuleResponseArray) ToNetworkSecurityRuleResponseArrayOutputWithContext(ctx context.Context) NetworkSecurityRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityRuleResponseArrayOutput)
}

type NetworkSecurityRuleResponseOutput struct{ *pulumi.OutputState }

func (NetworkSecurityRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityRuleResponse)(nil)).Elem()
}

func (o NetworkSecurityRuleResponseOutput) ToNetworkSecurityRuleResponseOutput() NetworkSecurityRuleResponseOutput {
	return o
}

func (o NetworkSecurityRuleResponseOutput) ToNetworkSecurityRuleResponseOutputWithContext(ctx context.Context) NetworkSecurityRuleResponseOutput {
	return o
}

func (o NetworkSecurityRuleResponseOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkSecurityRuleResponse) string { return v.Access }).(pulumi.StringOutput)
}

func (o NetworkSecurityRuleResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityRuleResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityRuleResponseOutput) DestinationAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkSecurityRuleResponse) []string { return v.DestinationAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o NetworkSecurityRuleResponseOutput) DestinationPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkSecurityRuleResponse) []string { return v.DestinationPortRanges }).(pulumi.StringArrayOutput)
}

func (o NetworkSecurityRuleResponseOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkSecurityRuleResponse) string { return v.Direction }).(pulumi.StringOutput)
}

func (o NetworkSecurityRuleResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkSecurityRuleResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkSecurityRuleResponseOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v NetworkSecurityRuleResponse) int { return v.Priority }).(pulumi.IntOutput)
}

func (o NetworkSecurityRuleResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkSecurityRuleResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o NetworkSecurityRuleResponseOutput) SourceAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkSecurityRuleResponse) []string { return v.SourceAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o NetworkSecurityRuleResponseOutput) SourcePortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkSecurityRuleResponse) []string { return v.SourcePortRanges }).(pulumi.StringArrayOutput)
}

type NetworkSecurityRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkSecurityRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkSecurityRuleResponse)(nil)).Elem()
}

func (o NetworkSecurityRuleResponseArrayOutput) ToNetworkSecurityRuleResponseArrayOutput() NetworkSecurityRuleResponseArrayOutput {
	return o
}

func (o NetworkSecurityRuleResponseArrayOutput) ToNetworkSecurityRuleResponseArrayOutputWithContext(ctx context.Context) NetworkSecurityRuleResponseArrayOutput {
	return o
}

func (o NetworkSecurityRuleResponseArrayOutput) Index(i pulumi.IntInput) NetworkSecurityRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkSecurityRuleResponse {
		return vs[0].([]NetworkSecurityRuleResponse)[vs[1].(int)]
	}).(NetworkSecurityRuleResponseOutput)
}

type PartitionInstanceCountScaleMechanism struct {
	Kind             string `pulumi:"kind"`
	MaxInstanceCount int    `pulumi:"maxInstanceCount"`
	MinInstanceCount int    `pulumi:"minInstanceCount"`
	ScaleIncrement   int    `pulumi:"scaleIncrement"`
}





type PartitionInstanceCountScaleMechanismInput interface {
	pulumi.Input

	ToPartitionInstanceCountScaleMechanismOutput() PartitionInstanceCountScaleMechanismOutput
	ToPartitionInstanceCountScaleMechanismOutputWithContext(context.Context) PartitionInstanceCountScaleMechanismOutput
}

type PartitionInstanceCountScaleMechanismArgs struct {
	Kind             pulumi.StringInput `pulumi:"kind"`
	MaxInstanceCount pulumi.IntInput    `pulumi:"maxInstanceCount"`
	MinInstanceCount pulumi.IntInput    `pulumi:"minInstanceCount"`
	ScaleIncrement   pulumi.IntInput    `pulumi:"scaleIncrement"`
}

func (PartitionInstanceCountScaleMechanismArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PartitionInstanceCountScaleMechanism)(nil)).Elem()
}

func (i PartitionInstanceCountScaleMechanismArgs) ToPartitionInstanceCountScaleMechanismOutput() PartitionInstanceCountScaleMechanismOutput {
	return i.ToPartitionInstanceCountScaleMechanismOutputWithContext(context.Background())
}

func (i PartitionInstanceCountScaleMechanismArgs) ToPartitionInstanceCountScaleMechanismOutputWithContext(ctx context.Context) PartitionInstanceCountScaleMechanismOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartitionInstanceCountScaleMechanismOutput)
}

type PartitionInstanceCountScaleMechanismOutput struct{ *pulumi.OutputState }

func (PartitionInstanceCountScaleMechanismOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartitionInstanceCountScaleMechanism)(nil)).Elem()
}

func (o PartitionInstanceCountScaleMechanismOutput) ToPartitionInstanceCountScaleMechanismOutput() PartitionInstanceCountScaleMechanismOutput {
	return o
}

func (o PartitionInstanceCountScaleMechanismOutput) ToPartitionInstanceCountScaleMechanismOutputWithContext(ctx context.Context) PartitionInstanceCountScaleMechanismOutput {
	return o
}

func (o PartitionInstanceCountScaleMechanismOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v PartitionInstanceCountScaleMechanism) string { return v.Kind }).(pulumi.StringOutput)
}

func (o PartitionInstanceCountScaleMechanismOutput) MaxInstanceCount() pulumi.IntOutput {
	return o.ApplyT(func(v PartitionInstanceCountScaleMechanism) int { return v.MaxInstanceCount }).(pulumi.IntOutput)
}

func (o PartitionInstanceCountScaleMechanismOutput) MinInstanceCount() pulumi.IntOutput {
	return o.ApplyT(func(v PartitionInstanceCountScaleMechanism) int { return v.MinInstanceCount }).(pulumi.IntOutput)
}

func (o PartitionInstanceCountScaleMechanismOutput) ScaleIncrement() pulumi.IntOutput {
	return o.ApplyT(func(v PartitionInstanceCountScaleMechanism) int { return v.ScaleIncrement }).(pulumi.IntOutput)
}

type PartitionInstanceCountScaleMechanismResponse struct {
	Kind             string `pulumi:"kind"`
	MaxInstanceCount int    `pulumi:"maxInstanceCount"`
	MinInstanceCount int    `pulumi:"minInstanceCount"`
	ScaleIncrement   int    `pulumi:"scaleIncrement"`
}





type PartitionInstanceCountScaleMechanismResponseInput interface {
	pulumi.Input

	ToPartitionInstanceCountScaleMechanismResponseOutput() PartitionInstanceCountScaleMechanismResponseOutput
	ToPartitionInstanceCountScaleMechanismResponseOutputWithContext(context.Context) PartitionInstanceCountScaleMechanismResponseOutput
}

type PartitionInstanceCountScaleMechanismResponseArgs struct {
	Kind             pulumi.StringInput `pulumi:"kind"`
	MaxInstanceCount pulumi.IntInput    `pulumi:"maxInstanceCount"`
	MinInstanceCount pulumi.IntInput    `pulumi:"minInstanceCount"`
	ScaleIncrement   pulumi.IntInput    `pulumi:"scaleIncrement"`
}

func (PartitionInstanceCountScaleMechanismResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PartitionInstanceCountScaleMechanismResponse)(nil)).Elem()
}

func (i PartitionInstanceCountScaleMechanismResponseArgs) ToPartitionInstanceCountScaleMechanismResponseOutput() PartitionInstanceCountScaleMechanismResponseOutput {
	return i.ToPartitionInstanceCountScaleMechanismResponseOutputWithContext(context.Background())
}

func (i PartitionInstanceCountScaleMechanismResponseArgs) ToPartitionInstanceCountScaleMechanismResponseOutputWithContext(ctx context.Context) PartitionInstanceCountScaleMechanismResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartitionInstanceCountScaleMechanismResponseOutput)
}

type PartitionInstanceCountScaleMechanismResponseOutput struct{ *pulumi.OutputState }

func (PartitionInstanceCountScaleMechanismResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartitionInstanceCountScaleMechanismResponse)(nil)).Elem()
}

func (o PartitionInstanceCountScaleMechanismResponseOutput) ToPartitionInstanceCountScaleMechanismResponseOutput() PartitionInstanceCountScaleMechanismResponseOutput {
	return o
}

func (o PartitionInstanceCountScaleMechanismResponseOutput) ToPartitionInstanceCountScaleMechanismResponseOutputWithContext(ctx context.Context) PartitionInstanceCountScaleMechanismResponseOutput {
	return o
}

func (o PartitionInstanceCountScaleMechanismResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v PartitionInstanceCountScaleMechanismResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o PartitionInstanceCountScaleMechanismResponseOutput) MaxInstanceCount() pulumi.IntOutput {
	return o.ApplyT(func(v PartitionInstanceCountScaleMechanismResponse) int { return v.MaxInstanceCount }).(pulumi.IntOutput)
}

func (o PartitionInstanceCountScaleMechanismResponseOutput) MinInstanceCount() pulumi.IntOutput {
	return o.ApplyT(func(v PartitionInstanceCountScaleMechanismResponse) int { return v.MinInstanceCount }).(pulumi.IntOutput)
}

func (o PartitionInstanceCountScaleMechanismResponseOutput) ScaleIncrement() pulumi.IntOutput {
	return o.ApplyT(func(v PartitionInstanceCountScaleMechanismResponse) int { return v.ScaleIncrement }).(pulumi.IntOutput)
}

type RollingUpgradeMonitoringPolicy struct {
	FailureAction             string `pulumi:"failureAction"`
	HealthCheckRetryTimeout   string `pulumi:"healthCheckRetryTimeout"`
	HealthCheckStableDuration string `pulumi:"healthCheckStableDuration"`
	HealthCheckWaitDuration   string `pulumi:"healthCheckWaitDuration"`
	UpgradeDomainTimeout      string `pulumi:"upgradeDomainTimeout"`
	UpgradeTimeout            string `pulumi:"upgradeTimeout"`
}





type RollingUpgradeMonitoringPolicyInput interface {
	pulumi.Input

	ToRollingUpgradeMonitoringPolicyOutput() RollingUpgradeMonitoringPolicyOutput
	ToRollingUpgradeMonitoringPolicyOutputWithContext(context.Context) RollingUpgradeMonitoringPolicyOutput
}

type RollingUpgradeMonitoringPolicyArgs struct {
	FailureAction             pulumi.StringInput `pulumi:"failureAction"`
	HealthCheckRetryTimeout   pulumi.StringInput `pulumi:"healthCheckRetryTimeout"`
	HealthCheckStableDuration pulumi.StringInput `pulumi:"healthCheckStableDuration"`
	HealthCheckWaitDuration   pulumi.StringInput `pulumi:"healthCheckWaitDuration"`
	UpgradeDomainTimeout      pulumi.StringInput `pulumi:"upgradeDomainTimeout"`
	UpgradeTimeout            pulumi.StringInput `pulumi:"upgradeTimeout"`
}

func (RollingUpgradeMonitoringPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RollingUpgradeMonitoringPolicy)(nil)).Elem()
}

func (i RollingUpgradeMonitoringPolicyArgs) ToRollingUpgradeMonitoringPolicyOutput() RollingUpgradeMonitoringPolicyOutput {
	return i.ToRollingUpgradeMonitoringPolicyOutputWithContext(context.Background())
}

func (i RollingUpgradeMonitoringPolicyArgs) ToRollingUpgradeMonitoringPolicyOutputWithContext(ctx context.Context) RollingUpgradeMonitoringPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RollingUpgradeMonitoringPolicyOutput)
}

func (i RollingUpgradeMonitoringPolicyArgs) ToRollingUpgradeMonitoringPolicyPtrOutput() RollingUpgradeMonitoringPolicyPtrOutput {
	return i.ToRollingUpgradeMonitoringPolicyPtrOutputWithContext(context.Background())
}

func (i RollingUpgradeMonitoringPolicyArgs) ToRollingUpgradeMonitoringPolicyPtrOutputWithContext(ctx context.Context) RollingUpgradeMonitoringPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RollingUpgradeMonitoringPolicyOutput).ToRollingUpgradeMonitoringPolicyPtrOutputWithContext(ctx)
}









type RollingUpgradeMonitoringPolicyPtrInput interface {
	pulumi.Input

	ToRollingUpgradeMonitoringPolicyPtrOutput() RollingUpgradeMonitoringPolicyPtrOutput
	ToRollingUpgradeMonitoringPolicyPtrOutputWithContext(context.Context) RollingUpgradeMonitoringPolicyPtrOutput
}

type rollingUpgradeMonitoringPolicyPtrType RollingUpgradeMonitoringPolicyArgs

func RollingUpgradeMonitoringPolicyPtr(v *RollingUpgradeMonitoringPolicyArgs) RollingUpgradeMonitoringPolicyPtrInput {
	return (*rollingUpgradeMonitoringPolicyPtrType)(v)
}

func (*rollingUpgradeMonitoringPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RollingUpgradeMonitoringPolicy)(nil)).Elem()
}

func (i *rollingUpgradeMonitoringPolicyPtrType) ToRollingUpgradeMonitoringPolicyPtrOutput() RollingUpgradeMonitoringPolicyPtrOutput {
	return i.ToRollingUpgradeMonitoringPolicyPtrOutputWithContext(context.Background())
}

func (i *rollingUpgradeMonitoringPolicyPtrType) ToRollingUpgradeMonitoringPolicyPtrOutputWithContext(ctx context.Context) RollingUpgradeMonitoringPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RollingUpgradeMonitoringPolicyPtrOutput)
}

type RollingUpgradeMonitoringPolicyOutput struct{ *pulumi.OutputState }

func (RollingUpgradeMonitoringPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RollingUpgradeMonitoringPolicy)(nil)).Elem()
}

func (o RollingUpgradeMonitoringPolicyOutput) ToRollingUpgradeMonitoringPolicyOutput() RollingUpgradeMonitoringPolicyOutput {
	return o
}

func (o RollingUpgradeMonitoringPolicyOutput) ToRollingUpgradeMonitoringPolicyOutputWithContext(ctx context.Context) RollingUpgradeMonitoringPolicyOutput {
	return o
}

func (o RollingUpgradeMonitoringPolicyOutput) ToRollingUpgradeMonitoringPolicyPtrOutput() RollingUpgradeMonitoringPolicyPtrOutput {
	return o.ToRollingUpgradeMonitoringPolicyPtrOutputWithContext(context.Background())
}

func (o RollingUpgradeMonitoringPolicyOutput) ToRollingUpgradeMonitoringPolicyPtrOutputWithContext(ctx context.Context) RollingUpgradeMonitoringPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RollingUpgradeMonitoringPolicy) *RollingUpgradeMonitoringPolicy {
		return &v
	}).(RollingUpgradeMonitoringPolicyPtrOutput)
}

func (o RollingUpgradeMonitoringPolicyOutput) FailureAction() pulumi.StringOutput {
	return o.ApplyT(func(v RollingUpgradeMonitoringPolicy) string { return v.FailureAction }).(pulumi.StringOutput)
}

func (o RollingUpgradeMonitoringPolicyOutput) HealthCheckRetryTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v RollingUpgradeMonitoringPolicy) string { return v.HealthCheckRetryTimeout }).(pulumi.StringOutput)
}

func (o RollingUpgradeMonitoringPolicyOutput) HealthCheckStableDuration() pulumi.StringOutput {
	return o.ApplyT(func(v RollingUpgradeMonitoringPolicy) string { return v.HealthCheckStableDuration }).(pulumi.StringOutput)
}

func (o RollingUpgradeMonitoringPolicyOutput) HealthCheckWaitDuration() pulumi.StringOutput {
	return o.ApplyT(func(v RollingUpgradeMonitoringPolicy) string { return v.HealthCheckWaitDuration }).(pulumi.StringOutput)
}

func (o RollingUpgradeMonitoringPolicyOutput) UpgradeDomainTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v RollingUpgradeMonitoringPolicy) string { return v.UpgradeDomainTimeout }).(pulumi.StringOutput)
}

func (o RollingUpgradeMonitoringPolicyOutput) UpgradeTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v RollingUpgradeMonitoringPolicy) string { return v.UpgradeTimeout }).(pulumi.StringOutput)
}

type RollingUpgradeMonitoringPolicyPtrOutput struct{ *pulumi.OutputState }

func (RollingUpgradeMonitoringPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RollingUpgradeMonitoringPolicy)(nil)).Elem()
}

func (o RollingUpgradeMonitoringPolicyPtrOutput) ToRollingUpgradeMonitoringPolicyPtrOutput() RollingUpgradeMonitoringPolicyPtrOutput {
	return o
}

func (o RollingUpgradeMonitoringPolicyPtrOutput) ToRollingUpgradeMonitoringPolicyPtrOutputWithContext(ctx context.Context) RollingUpgradeMonitoringPolicyPtrOutput {
	return o
}

func (o RollingUpgradeMonitoringPolicyPtrOutput) Elem() RollingUpgradeMonitoringPolicyOutput {
	return o.ApplyT(func(v *RollingUpgradeMonitoringPolicy) RollingUpgradeMonitoringPolicy {
		if v != nil {
			return *v
		}
		var ret RollingUpgradeMonitoringPolicy
		return ret
	}).(RollingUpgradeMonitoringPolicyOutput)
}

func (o RollingUpgradeMonitoringPolicyPtrOutput) FailureAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RollingUpgradeMonitoringPolicy) *string {
		if v == nil {
			return nil
		}
		return &v.FailureAction
	}).(pulumi.StringPtrOutput)
}

func (o RollingUpgradeMonitoringPolicyPtrOutput) HealthCheckRetryTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RollingUpgradeMonitoringPolicy) *string {
		if v == nil {
			return nil
		}
		return &v.HealthCheckRetryTimeout
	}).(pulumi.StringPtrOutput)
}

func (o RollingUpgradeMonitoringPolicyPtrOutput) HealthCheckStableDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RollingUpgradeMonitoringPolicy) *string {
		if v == nil {
			return nil
		}
		return &v.HealthCheckStableDuration
	}).(pulumi.StringPtrOutput)
}

func (o RollingUpgradeMonitoringPolicyPtrOutput) HealthCheckWaitDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RollingUpgradeMonitoringPolicy) *string {
		if v == nil {
			return nil
		}
		return &v.HealthCheckWaitDuration
	}).(pulumi.StringPtrOutput)
}

func (o RollingUpgradeMonitoringPolicyPtrOutput) UpgradeDomainTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RollingUpgradeMonitoringPolicy) *string {
		if v == nil {
			return nil
		}
		return &v.UpgradeDomainTimeout
	}).(pulumi.StringPtrOutput)
}

func (o RollingUpgradeMonitoringPolicyPtrOutput) UpgradeTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RollingUpgradeMonitoringPolicy) *string {
		if v == nil {
			return nil
		}
		return &v.UpgradeTimeout
	}).(pulumi.StringPtrOutput)
}

type RollingUpgradeMonitoringPolicyResponse struct {
	FailureAction             string `pulumi:"failureAction"`
	HealthCheckRetryTimeout   string `pulumi:"healthCheckRetryTimeout"`
	HealthCheckStableDuration string `pulumi:"healthCheckStableDuration"`
	HealthCheckWaitDuration   string `pulumi:"healthCheckWaitDuration"`
	UpgradeDomainTimeout      string `pulumi:"upgradeDomainTimeout"`
	UpgradeTimeout            string `pulumi:"upgradeTimeout"`
}





type RollingUpgradeMonitoringPolicyResponseInput interface {
	pulumi.Input

	ToRollingUpgradeMonitoringPolicyResponseOutput() RollingUpgradeMonitoringPolicyResponseOutput
	ToRollingUpgradeMonitoringPolicyResponseOutputWithContext(context.Context) RollingUpgradeMonitoringPolicyResponseOutput
}

type RollingUpgradeMonitoringPolicyResponseArgs struct {
	FailureAction             pulumi.StringInput `pulumi:"failureAction"`
	HealthCheckRetryTimeout   pulumi.StringInput `pulumi:"healthCheckRetryTimeout"`
	HealthCheckStableDuration pulumi.StringInput `pulumi:"healthCheckStableDuration"`
	HealthCheckWaitDuration   pulumi.StringInput `pulumi:"healthCheckWaitDuration"`
	UpgradeDomainTimeout      pulumi.StringInput `pulumi:"upgradeDomainTimeout"`
	UpgradeTimeout            pulumi.StringInput `pulumi:"upgradeTimeout"`
}

func (RollingUpgradeMonitoringPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RollingUpgradeMonitoringPolicyResponse)(nil)).Elem()
}

func (i RollingUpgradeMonitoringPolicyResponseArgs) ToRollingUpgradeMonitoringPolicyResponseOutput() RollingUpgradeMonitoringPolicyResponseOutput {
	return i.ToRollingUpgradeMonitoringPolicyResponseOutputWithContext(context.Background())
}

func (i RollingUpgradeMonitoringPolicyResponseArgs) ToRollingUpgradeMonitoringPolicyResponseOutputWithContext(ctx context.Context) RollingUpgradeMonitoringPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RollingUpgradeMonitoringPolicyResponseOutput)
}

func (i RollingUpgradeMonitoringPolicyResponseArgs) ToRollingUpgradeMonitoringPolicyResponsePtrOutput() RollingUpgradeMonitoringPolicyResponsePtrOutput {
	return i.ToRollingUpgradeMonitoringPolicyResponsePtrOutputWithContext(context.Background())
}

func (i RollingUpgradeMonitoringPolicyResponseArgs) ToRollingUpgradeMonitoringPolicyResponsePtrOutputWithContext(ctx context.Context) RollingUpgradeMonitoringPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RollingUpgradeMonitoringPolicyResponseOutput).ToRollingUpgradeMonitoringPolicyResponsePtrOutputWithContext(ctx)
}









type RollingUpgradeMonitoringPolicyResponsePtrInput interface {
	pulumi.Input

	ToRollingUpgradeMonitoringPolicyResponsePtrOutput() RollingUpgradeMonitoringPolicyResponsePtrOutput
	ToRollingUpgradeMonitoringPolicyResponsePtrOutputWithContext(context.Context) RollingUpgradeMonitoringPolicyResponsePtrOutput
}

type rollingUpgradeMonitoringPolicyResponsePtrType RollingUpgradeMonitoringPolicyResponseArgs

func RollingUpgradeMonitoringPolicyResponsePtr(v *RollingUpgradeMonitoringPolicyResponseArgs) RollingUpgradeMonitoringPolicyResponsePtrInput {
	return (*rollingUpgradeMonitoringPolicyResponsePtrType)(v)
}

func (*rollingUpgradeMonitoringPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RollingUpgradeMonitoringPolicyResponse)(nil)).Elem()
}

func (i *rollingUpgradeMonitoringPolicyResponsePtrType) ToRollingUpgradeMonitoringPolicyResponsePtrOutput() RollingUpgradeMonitoringPolicyResponsePtrOutput {
	return i.ToRollingUpgradeMonitoringPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *rollingUpgradeMonitoringPolicyResponsePtrType) ToRollingUpgradeMonitoringPolicyResponsePtrOutputWithContext(ctx context.Context) RollingUpgradeMonitoringPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RollingUpgradeMonitoringPolicyResponsePtrOutput)
}

type RollingUpgradeMonitoringPolicyResponseOutput struct{ *pulumi.OutputState }

func (RollingUpgradeMonitoringPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RollingUpgradeMonitoringPolicyResponse)(nil)).Elem()
}

func (o RollingUpgradeMonitoringPolicyResponseOutput) ToRollingUpgradeMonitoringPolicyResponseOutput() RollingUpgradeMonitoringPolicyResponseOutput {
	return o
}

func (o RollingUpgradeMonitoringPolicyResponseOutput) ToRollingUpgradeMonitoringPolicyResponseOutputWithContext(ctx context.Context) RollingUpgradeMonitoringPolicyResponseOutput {
	return o
}

func (o RollingUpgradeMonitoringPolicyResponseOutput) ToRollingUpgradeMonitoringPolicyResponsePtrOutput() RollingUpgradeMonitoringPolicyResponsePtrOutput {
	return o.ToRollingUpgradeMonitoringPolicyResponsePtrOutputWithContext(context.Background())
}

func (o RollingUpgradeMonitoringPolicyResponseOutput) ToRollingUpgradeMonitoringPolicyResponsePtrOutputWithContext(ctx context.Context) RollingUpgradeMonitoringPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RollingUpgradeMonitoringPolicyResponse) *RollingUpgradeMonitoringPolicyResponse {
		return &v
	}).(RollingUpgradeMonitoringPolicyResponsePtrOutput)
}

func (o RollingUpgradeMonitoringPolicyResponseOutput) FailureAction() pulumi.StringOutput {
	return o.ApplyT(func(v RollingUpgradeMonitoringPolicyResponse) string { return v.FailureAction }).(pulumi.StringOutput)
}

func (o RollingUpgradeMonitoringPolicyResponseOutput) HealthCheckRetryTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v RollingUpgradeMonitoringPolicyResponse) string { return v.HealthCheckRetryTimeout }).(pulumi.StringOutput)
}

func (o RollingUpgradeMonitoringPolicyResponseOutput) HealthCheckStableDuration() pulumi.StringOutput {
	return o.ApplyT(func(v RollingUpgradeMonitoringPolicyResponse) string { return v.HealthCheckStableDuration }).(pulumi.StringOutput)
}

func (o RollingUpgradeMonitoringPolicyResponseOutput) HealthCheckWaitDuration() pulumi.StringOutput {
	return o.ApplyT(func(v RollingUpgradeMonitoringPolicyResponse) string { return v.HealthCheckWaitDuration }).(pulumi.StringOutput)
}

func (o RollingUpgradeMonitoringPolicyResponseOutput) UpgradeDomainTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v RollingUpgradeMonitoringPolicyResponse) string { return v.UpgradeDomainTimeout }).(pulumi.StringOutput)
}

func (o RollingUpgradeMonitoringPolicyResponseOutput) UpgradeTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v RollingUpgradeMonitoringPolicyResponse) string { return v.UpgradeTimeout }).(pulumi.StringOutput)
}

type RollingUpgradeMonitoringPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (RollingUpgradeMonitoringPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RollingUpgradeMonitoringPolicyResponse)(nil)).Elem()
}

func (o RollingUpgradeMonitoringPolicyResponsePtrOutput) ToRollingUpgradeMonitoringPolicyResponsePtrOutput() RollingUpgradeMonitoringPolicyResponsePtrOutput {
	return o
}

func (o RollingUpgradeMonitoringPolicyResponsePtrOutput) ToRollingUpgradeMonitoringPolicyResponsePtrOutputWithContext(ctx context.Context) RollingUpgradeMonitoringPolicyResponsePtrOutput {
	return o
}

func (o RollingUpgradeMonitoringPolicyResponsePtrOutput) Elem() RollingUpgradeMonitoringPolicyResponseOutput {
	return o.ApplyT(func(v *RollingUpgradeMonitoringPolicyResponse) RollingUpgradeMonitoringPolicyResponse {
		if v != nil {
			return *v
		}
		var ret RollingUpgradeMonitoringPolicyResponse
		return ret
	}).(RollingUpgradeMonitoringPolicyResponseOutput)
}

func (o RollingUpgradeMonitoringPolicyResponsePtrOutput) FailureAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RollingUpgradeMonitoringPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.FailureAction
	}).(pulumi.StringPtrOutput)
}

func (o RollingUpgradeMonitoringPolicyResponsePtrOutput) HealthCheckRetryTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RollingUpgradeMonitoringPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.HealthCheckRetryTimeout
	}).(pulumi.StringPtrOutput)
}

func (o RollingUpgradeMonitoringPolicyResponsePtrOutput) HealthCheckStableDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RollingUpgradeMonitoringPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.HealthCheckStableDuration
	}).(pulumi.StringPtrOutput)
}

func (o RollingUpgradeMonitoringPolicyResponsePtrOutput) HealthCheckWaitDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RollingUpgradeMonitoringPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.HealthCheckWaitDuration
	}).(pulumi.StringPtrOutput)
}

func (o RollingUpgradeMonitoringPolicyResponsePtrOutput) UpgradeDomainTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RollingUpgradeMonitoringPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UpgradeDomainTimeout
	}).(pulumi.StringPtrOutput)
}

func (o RollingUpgradeMonitoringPolicyResponsePtrOutput) UpgradeTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RollingUpgradeMonitoringPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UpgradeTimeout
	}).(pulumi.StringPtrOutput)
}

type ScalingPolicy struct {
	ScalingMechanism interface{} `pulumi:"scalingMechanism"`
	ScalingTrigger   interface{} `pulumi:"scalingTrigger"`
}





type ScalingPolicyInput interface {
	pulumi.Input

	ToScalingPolicyOutput() ScalingPolicyOutput
	ToScalingPolicyOutputWithContext(context.Context) ScalingPolicyOutput
}

type ScalingPolicyArgs struct {
	ScalingMechanism pulumi.Input `pulumi:"scalingMechanism"`
	ScalingTrigger   pulumi.Input `pulumi:"scalingTrigger"`
}

func (ScalingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingPolicy)(nil)).Elem()
}

func (i ScalingPolicyArgs) ToScalingPolicyOutput() ScalingPolicyOutput {
	return i.ToScalingPolicyOutputWithContext(context.Background())
}

func (i ScalingPolicyArgs) ToScalingPolicyOutputWithContext(ctx context.Context) ScalingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingPolicyOutput)
}





type ScalingPolicyArrayInput interface {
	pulumi.Input

	ToScalingPolicyArrayOutput() ScalingPolicyArrayOutput
	ToScalingPolicyArrayOutputWithContext(context.Context) ScalingPolicyArrayOutput
}

type ScalingPolicyArray []ScalingPolicyInput

func (ScalingPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalingPolicy)(nil)).Elem()
}

func (i ScalingPolicyArray) ToScalingPolicyArrayOutput() ScalingPolicyArrayOutput {
	return i.ToScalingPolicyArrayOutputWithContext(context.Background())
}

func (i ScalingPolicyArray) ToScalingPolicyArrayOutputWithContext(ctx context.Context) ScalingPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingPolicyArrayOutput)
}

type ScalingPolicyOutput struct{ *pulumi.OutputState }

func (ScalingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingPolicy)(nil)).Elem()
}

func (o ScalingPolicyOutput) ToScalingPolicyOutput() ScalingPolicyOutput {
	return o
}

func (o ScalingPolicyOutput) ToScalingPolicyOutputWithContext(ctx context.Context) ScalingPolicyOutput {
	return o
}

func (o ScalingPolicyOutput) ScalingMechanism() pulumi.AnyOutput {
	return o.ApplyT(func(v ScalingPolicy) interface{} { return v.ScalingMechanism }).(pulumi.AnyOutput)
}

func (o ScalingPolicyOutput) ScalingTrigger() pulumi.AnyOutput {
	return o.ApplyT(func(v ScalingPolicy) interface{} { return v.ScalingTrigger }).(pulumi.AnyOutput)
}

type ScalingPolicyArrayOutput struct{ *pulumi.OutputState }

func (ScalingPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalingPolicy)(nil)).Elem()
}

func (o ScalingPolicyArrayOutput) ToScalingPolicyArrayOutput() ScalingPolicyArrayOutput {
	return o
}

func (o ScalingPolicyArrayOutput) ToScalingPolicyArrayOutputWithContext(ctx context.Context) ScalingPolicyArrayOutput {
	return o
}

func (o ScalingPolicyArrayOutput) Index(i pulumi.IntInput) ScalingPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScalingPolicy {
		return vs[0].([]ScalingPolicy)[vs[1].(int)]
	}).(ScalingPolicyOutput)
}

type ScalingPolicyResponse struct {
	ScalingMechanism interface{} `pulumi:"scalingMechanism"`
	ScalingTrigger   interface{} `pulumi:"scalingTrigger"`
}





type ScalingPolicyResponseInput interface {
	pulumi.Input

	ToScalingPolicyResponseOutput() ScalingPolicyResponseOutput
	ToScalingPolicyResponseOutputWithContext(context.Context) ScalingPolicyResponseOutput
}

type ScalingPolicyResponseArgs struct {
	ScalingMechanism pulumi.Input `pulumi:"scalingMechanism"`
	ScalingTrigger   pulumi.Input `pulumi:"scalingTrigger"`
}

func (ScalingPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingPolicyResponse)(nil)).Elem()
}

func (i ScalingPolicyResponseArgs) ToScalingPolicyResponseOutput() ScalingPolicyResponseOutput {
	return i.ToScalingPolicyResponseOutputWithContext(context.Background())
}

func (i ScalingPolicyResponseArgs) ToScalingPolicyResponseOutputWithContext(ctx context.Context) ScalingPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingPolicyResponseOutput)
}





type ScalingPolicyResponseArrayInput interface {
	pulumi.Input

	ToScalingPolicyResponseArrayOutput() ScalingPolicyResponseArrayOutput
	ToScalingPolicyResponseArrayOutputWithContext(context.Context) ScalingPolicyResponseArrayOutput
}

type ScalingPolicyResponseArray []ScalingPolicyResponseInput

func (ScalingPolicyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalingPolicyResponse)(nil)).Elem()
}

func (i ScalingPolicyResponseArray) ToScalingPolicyResponseArrayOutput() ScalingPolicyResponseArrayOutput {
	return i.ToScalingPolicyResponseArrayOutputWithContext(context.Background())
}

func (i ScalingPolicyResponseArray) ToScalingPolicyResponseArrayOutputWithContext(ctx context.Context) ScalingPolicyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingPolicyResponseArrayOutput)
}

type ScalingPolicyResponseOutput struct{ *pulumi.OutputState }

func (ScalingPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingPolicyResponse)(nil)).Elem()
}

func (o ScalingPolicyResponseOutput) ToScalingPolicyResponseOutput() ScalingPolicyResponseOutput {
	return o
}

func (o ScalingPolicyResponseOutput) ToScalingPolicyResponseOutputWithContext(ctx context.Context) ScalingPolicyResponseOutput {
	return o
}

func (o ScalingPolicyResponseOutput) ScalingMechanism() pulumi.AnyOutput {
	return o.ApplyT(func(v ScalingPolicyResponse) interface{} { return v.ScalingMechanism }).(pulumi.AnyOutput)
}

func (o ScalingPolicyResponseOutput) ScalingTrigger() pulumi.AnyOutput {
	return o.ApplyT(func(v ScalingPolicyResponse) interface{} { return v.ScalingTrigger }).(pulumi.AnyOutput)
}

type ScalingPolicyResponseArrayOutput struct{ *pulumi.OutputState }

func (ScalingPolicyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalingPolicyResponse)(nil)).Elem()
}

func (o ScalingPolicyResponseArrayOutput) ToScalingPolicyResponseArrayOutput() ScalingPolicyResponseArrayOutput {
	return o
}

func (o ScalingPolicyResponseArrayOutput) ToScalingPolicyResponseArrayOutputWithContext(ctx context.Context) ScalingPolicyResponseArrayOutput {
	return o
}

func (o ScalingPolicyResponseArrayOutput) Index(i pulumi.IntInput) ScalingPolicyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScalingPolicyResponse {
		return vs[0].([]ScalingPolicyResponse)[vs[1].(int)]
	}).(ScalingPolicyResponseOutput)
}

type ServiceCorrelation struct {
	Scheme      string `pulumi:"scheme"`
	ServiceName string `pulumi:"serviceName"`
}





type ServiceCorrelationInput interface {
	pulumi.Input

	ToServiceCorrelationOutput() ServiceCorrelationOutput
	ToServiceCorrelationOutputWithContext(context.Context) ServiceCorrelationOutput
}

type ServiceCorrelationArgs struct {
	Scheme      pulumi.StringInput `pulumi:"scheme"`
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (ServiceCorrelationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCorrelation)(nil)).Elem()
}

func (i ServiceCorrelationArgs) ToServiceCorrelationOutput() ServiceCorrelationOutput {
	return i.ToServiceCorrelationOutputWithContext(context.Background())
}

func (i ServiceCorrelationArgs) ToServiceCorrelationOutputWithContext(ctx context.Context) ServiceCorrelationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorrelationOutput)
}





type ServiceCorrelationArrayInput interface {
	pulumi.Input

	ToServiceCorrelationArrayOutput() ServiceCorrelationArrayOutput
	ToServiceCorrelationArrayOutputWithContext(context.Context) ServiceCorrelationArrayOutput
}

type ServiceCorrelationArray []ServiceCorrelationInput

func (ServiceCorrelationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceCorrelation)(nil)).Elem()
}

func (i ServiceCorrelationArray) ToServiceCorrelationArrayOutput() ServiceCorrelationArrayOutput {
	return i.ToServiceCorrelationArrayOutputWithContext(context.Background())
}

func (i ServiceCorrelationArray) ToServiceCorrelationArrayOutputWithContext(ctx context.Context) ServiceCorrelationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorrelationArrayOutput)
}

type ServiceCorrelationOutput struct{ *pulumi.OutputState }

func (ServiceCorrelationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCorrelation)(nil)).Elem()
}

func (o ServiceCorrelationOutput) ToServiceCorrelationOutput() ServiceCorrelationOutput {
	return o
}

func (o ServiceCorrelationOutput) ToServiceCorrelationOutputWithContext(ctx context.Context) ServiceCorrelationOutput {
	return o
}

func (o ServiceCorrelationOutput) Scheme() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceCorrelation) string { return v.Scheme }).(pulumi.StringOutput)
}

func (o ServiceCorrelationOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceCorrelation) string { return v.ServiceName }).(pulumi.StringOutput)
}

type ServiceCorrelationArrayOutput struct{ *pulumi.OutputState }

func (ServiceCorrelationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceCorrelation)(nil)).Elem()
}

func (o ServiceCorrelationArrayOutput) ToServiceCorrelationArrayOutput() ServiceCorrelationArrayOutput {
	return o
}

func (o ServiceCorrelationArrayOutput) ToServiceCorrelationArrayOutputWithContext(ctx context.Context) ServiceCorrelationArrayOutput {
	return o
}

func (o ServiceCorrelationArrayOutput) Index(i pulumi.IntInput) ServiceCorrelationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceCorrelation {
		return vs[0].([]ServiceCorrelation)[vs[1].(int)]
	}).(ServiceCorrelationOutput)
}

type ServiceCorrelationResponse struct {
	Scheme      string `pulumi:"scheme"`
	ServiceName string `pulumi:"serviceName"`
}





type ServiceCorrelationResponseInput interface {
	pulumi.Input

	ToServiceCorrelationResponseOutput() ServiceCorrelationResponseOutput
	ToServiceCorrelationResponseOutputWithContext(context.Context) ServiceCorrelationResponseOutput
}

type ServiceCorrelationResponseArgs struct {
	Scheme      pulumi.StringInput `pulumi:"scheme"`
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (ServiceCorrelationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCorrelationResponse)(nil)).Elem()
}

func (i ServiceCorrelationResponseArgs) ToServiceCorrelationResponseOutput() ServiceCorrelationResponseOutput {
	return i.ToServiceCorrelationResponseOutputWithContext(context.Background())
}

func (i ServiceCorrelationResponseArgs) ToServiceCorrelationResponseOutputWithContext(ctx context.Context) ServiceCorrelationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorrelationResponseOutput)
}





type ServiceCorrelationResponseArrayInput interface {
	pulumi.Input

	ToServiceCorrelationResponseArrayOutput() ServiceCorrelationResponseArrayOutput
	ToServiceCorrelationResponseArrayOutputWithContext(context.Context) ServiceCorrelationResponseArrayOutput
}

type ServiceCorrelationResponseArray []ServiceCorrelationResponseInput

func (ServiceCorrelationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceCorrelationResponse)(nil)).Elem()
}

func (i ServiceCorrelationResponseArray) ToServiceCorrelationResponseArrayOutput() ServiceCorrelationResponseArrayOutput {
	return i.ToServiceCorrelationResponseArrayOutputWithContext(context.Background())
}

func (i ServiceCorrelationResponseArray) ToServiceCorrelationResponseArrayOutputWithContext(ctx context.Context) ServiceCorrelationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorrelationResponseArrayOutput)
}

type ServiceCorrelationResponseOutput struct{ *pulumi.OutputState }

func (ServiceCorrelationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCorrelationResponse)(nil)).Elem()
}

func (o ServiceCorrelationResponseOutput) ToServiceCorrelationResponseOutput() ServiceCorrelationResponseOutput {
	return o
}

func (o ServiceCorrelationResponseOutput) ToServiceCorrelationResponseOutputWithContext(ctx context.Context) ServiceCorrelationResponseOutput {
	return o
}

func (o ServiceCorrelationResponseOutput) Scheme() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceCorrelationResponse) string { return v.Scheme }).(pulumi.StringOutput)
}

func (o ServiceCorrelationResponseOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceCorrelationResponse) string { return v.ServiceName }).(pulumi.StringOutput)
}

type ServiceCorrelationResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceCorrelationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceCorrelationResponse)(nil)).Elem()
}

func (o ServiceCorrelationResponseArrayOutput) ToServiceCorrelationResponseArrayOutput() ServiceCorrelationResponseArrayOutput {
	return o
}

func (o ServiceCorrelationResponseArrayOutput) ToServiceCorrelationResponseArrayOutputWithContext(ctx context.Context) ServiceCorrelationResponseArrayOutput {
	return o
}

func (o ServiceCorrelationResponseArrayOutput) Index(i pulumi.IntInput) ServiceCorrelationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceCorrelationResponse {
		return vs[0].([]ServiceCorrelationResponse)[vs[1].(int)]
	}).(ServiceCorrelationResponseOutput)
}

type ServiceLoadMetric struct {
	DefaultLoad          *int    `pulumi:"defaultLoad"`
	Name                 string  `pulumi:"name"`
	PrimaryDefaultLoad   *int    `pulumi:"primaryDefaultLoad"`
	SecondaryDefaultLoad *int    `pulumi:"secondaryDefaultLoad"`
	Weight               *string `pulumi:"weight"`
}





type ServiceLoadMetricInput interface {
	pulumi.Input

	ToServiceLoadMetricOutput() ServiceLoadMetricOutput
	ToServiceLoadMetricOutputWithContext(context.Context) ServiceLoadMetricOutput
}

type ServiceLoadMetricArgs struct {
	DefaultLoad          pulumi.IntPtrInput    `pulumi:"defaultLoad"`
	Name                 pulumi.StringInput    `pulumi:"name"`
	PrimaryDefaultLoad   pulumi.IntPtrInput    `pulumi:"primaryDefaultLoad"`
	SecondaryDefaultLoad pulumi.IntPtrInput    `pulumi:"secondaryDefaultLoad"`
	Weight               pulumi.StringPtrInput `pulumi:"weight"`
}

func (ServiceLoadMetricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceLoadMetric)(nil)).Elem()
}

func (i ServiceLoadMetricArgs) ToServiceLoadMetricOutput() ServiceLoadMetricOutput {
	return i.ToServiceLoadMetricOutputWithContext(context.Background())
}

func (i ServiceLoadMetricArgs) ToServiceLoadMetricOutputWithContext(ctx context.Context) ServiceLoadMetricOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceLoadMetricOutput)
}





type ServiceLoadMetricArrayInput interface {
	pulumi.Input

	ToServiceLoadMetricArrayOutput() ServiceLoadMetricArrayOutput
	ToServiceLoadMetricArrayOutputWithContext(context.Context) ServiceLoadMetricArrayOutput
}

type ServiceLoadMetricArray []ServiceLoadMetricInput

func (ServiceLoadMetricArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceLoadMetric)(nil)).Elem()
}

func (i ServiceLoadMetricArray) ToServiceLoadMetricArrayOutput() ServiceLoadMetricArrayOutput {
	return i.ToServiceLoadMetricArrayOutputWithContext(context.Background())
}

func (i ServiceLoadMetricArray) ToServiceLoadMetricArrayOutputWithContext(ctx context.Context) ServiceLoadMetricArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceLoadMetricArrayOutput)
}

type ServiceLoadMetricOutput struct{ *pulumi.OutputState }

func (ServiceLoadMetricOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceLoadMetric)(nil)).Elem()
}

func (o ServiceLoadMetricOutput) ToServiceLoadMetricOutput() ServiceLoadMetricOutput {
	return o
}

func (o ServiceLoadMetricOutput) ToServiceLoadMetricOutputWithContext(ctx context.Context) ServiceLoadMetricOutput {
	return o
}

func (o ServiceLoadMetricOutput) DefaultLoad() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceLoadMetric) *int { return v.DefaultLoad }).(pulumi.IntPtrOutput)
}

func (o ServiceLoadMetricOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceLoadMetric) string { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceLoadMetricOutput) PrimaryDefaultLoad() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceLoadMetric) *int { return v.PrimaryDefaultLoad }).(pulumi.IntPtrOutput)
}

func (o ServiceLoadMetricOutput) SecondaryDefaultLoad() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceLoadMetric) *int { return v.SecondaryDefaultLoad }).(pulumi.IntPtrOutput)
}

func (o ServiceLoadMetricOutput) Weight() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceLoadMetric) *string { return v.Weight }).(pulumi.StringPtrOutput)
}

type ServiceLoadMetricArrayOutput struct{ *pulumi.OutputState }

func (ServiceLoadMetricArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceLoadMetric)(nil)).Elem()
}

func (o ServiceLoadMetricArrayOutput) ToServiceLoadMetricArrayOutput() ServiceLoadMetricArrayOutput {
	return o
}

func (o ServiceLoadMetricArrayOutput) ToServiceLoadMetricArrayOutputWithContext(ctx context.Context) ServiceLoadMetricArrayOutput {
	return o
}

func (o ServiceLoadMetricArrayOutput) Index(i pulumi.IntInput) ServiceLoadMetricOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceLoadMetric {
		return vs[0].([]ServiceLoadMetric)[vs[1].(int)]
	}).(ServiceLoadMetricOutput)
}

type ServiceLoadMetricResponse struct {
	DefaultLoad          *int    `pulumi:"defaultLoad"`
	Name                 string  `pulumi:"name"`
	PrimaryDefaultLoad   *int    `pulumi:"primaryDefaultLoad"`
	SecondaryDefaultLoad *int    `pulumi:"secondaryDefaultLoad"`
	Weight               *string `pulumi:"weight"`
}





type ServiceLoadMetricResponseInput interface {
	pulumi.Input

	ToServiceLoadMetricResponseOutput() ServiceLoadMetricResponseOutput
	ToServiceLoadMetricResponseOutputWithContext(context.Context) ServiceLoadMetricResponseOutput
}

type ServiceLoadMetricResponseArgs struct {
	DefaultLoad          pulumi.IntPtrInput    `pulumi:"defaultLoad"`
	Name                 pulumi.StringInput    `pulumi:"name"`
	PrimaryDefaultLoad   pulumi.IntPtrInput    `pulumi:"primaryDefaultLoad"`
	SecondaryDefaultLoad pulumi.IntPtrInput    `pulumi:"secondaryDefaultLoad"`
	Weight               pulumi.StringPtrInput `pulumi:"weight"`
}

func (ServiceLoadMetricResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceLoadMetricResponse)(nil)).Elem()
}

func (i ServiceLoadMetricResponseArgs) ToServiceLoadMetricResponseOutput() ServiceLoadMetricResponseOutput {
	return i.ToServiceLoadMetricResponseOutputWithContext(context.Background())
}

func (i ServiceLoadMetricResponseArgs) ToServiceLoadMetricResponseOutputWithContext(ctx context.Context) ServiceLoadMetricResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceLoadMetricResponseOutput)
}





type ServiceLoadMetricResponseArrayInput interface {
	pulumi.Input

	ToServiceLoadMetricResponseArrayOutput() ServiceLoadMetricResponseArrayOutput
	ToServiceLoadMetricResponseArrayOutputWithContext(context.Context) ServiceLoadMetricResponseArrayOutput
}

type ServiceLoadMetricResponseArray []ServiceLoadMetricResponseInput

func (ServiceLoadMetricResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceLoadMetricResponse)(nil)).Elem()
}

func (i ServiceLoadMetricResponseArray) ToServiceLoadMetricResponseArrayOutput() ServiceLoadMetricResponseArrayOutput {
	return i.ToServiceLoadMetricResponseArrayOutputWithContext(context.Background())
}

func (i ServiceLoadMetricResponseArray) ToServiceLoadMetricResponseArrayOutputWithContext(ctx context.Context) ServiceLoadMetricResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceLoadMetricResponseArrayOutput)
}

type ServiceLoadMetricResponseOutput struct{ *pulumi.OutputState }

func (ServiceLoadMetricResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceLoadMetricResponse)(nil)).Elem()
}

func (o ServiceLoadMetricResponseOutput) ToServiceLoadMetricResponseOutput() ServiceLoadMetricResponseOutput {
	return o
}

func (o ServiceLoadMetricResponseOutput) ToServiceLoadMetricResponseOutputWithContext(ctx context.Context) ServiceLoadMetricResponseOutput {
	return o
}

func (o ServiceLoadMetricResponseOutput) DefaultLoad() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceLoadMetricResponse) *int { return v.DefaultLoad }).(pulumi.IntPtrOutput)
}

func (o ServiceLoadMetricResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceLoadMetricResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceLoadMetricResponseOutput) PrimaryDefaultLoad() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceLoadMetricResponse) *int { return v.PrimaryDefaultLoad }).(pulumi.IntPtrOutput)
}

func (o ServiceLoadMetricResponseOutput) SecondaryDefaultLoad() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceLoadMetricResponse) *int { return v.SecondaryDefaultLoad }).(pulumi.IntPtrOutput)
}

func (o ServiceLoadMetricResponseOutput) Weight() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceLoadMetricResponse) *string { return v.Weight }).(pulumi.StringPtrOutput)
}

type ServiceLoadMetricResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceLoadMetricResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceLoadMetricResponse)(nil)).Elem()
}

func (o ServiceLoadMetricResponseArrayOutput) ToServiceLoadMetricResponseArrayOutput() ServiceLoadMetricResponseArrayOutput {
	return o
}

func (o ServiceLoadMetricResponseArrayOutput) ToServiceLoadMetricResponseArrayOutputWithContext(ctx context.Context) ServiceLoadMetricResponseArrayOutput {
	return o
}

func (o ServiceLoadMetricResponseArrayOutput) Index(i pulumi.IntInput) ServiceLoadMetricResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceLoadMetricResponse {
		return vs[0].([]ServiceLoadMetricResponse)[vs[1].(int)]
	}).(ServiceLoadMetricResponseOutput)
}

type ServicePlacementInvalidDomainPolicy struct {
	DomainName string `pulumi:"domainName"`
	Type       string `pulumi:"type"`
}





type ServicePlacementInvalidDomainPolicyInput interface {
	pulumi.Input

	ToServicePlacementInvalidDomainPolicyOutput() ServicePlacementInvalidDomainPolicyOutput
	ToServicePlacementInvalidDomainPolicyOutputWithContext(context.Context) ServicePlacementInvalidDomainPolicyOutput
}

type ServicePlacementInvalidDomainPolicyArgs struct {
	DomainName pulumi.StringInput `pulumi:"domainName"`
	Type       pulumi.StringInput `pulumi:"type"`
}

func (ServicePlacementInvalidDomainPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementInvalidDomainPolicy)(nil)).Elem()
}

func (i ServicePlacementInvalidDomainPolicyArgs) ToServicePlacementInvalidDomainPolicyOutput() ServicePlacementInvalidDomainPolicyOutput {
	return i.ToServicePlacementInvalidDomainPolicyOutputWithContext(context.Background())
}

func (i ServicePlacementInvalidDomainPolicyArgs) ToServicePlacementInvalidDomainPolicyOutputWithContext(ctx context.Context) ServicePlacementInvalidDomainPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePlacementInvalidDomainPolicyOutput)
}

type ServicePlacementInvalidDomainPolicyOutput struct{ *pulumi.OutputState }

func (ServicePlacementInvalidDomainPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementInvalidDomainPolicy)(nil)).Elem()
}

func (o ServicePlacementInvalidDomainPolicyOutput) ToServicePlacementInvalidDomainPolicyOutput() ServicePlacementInvalidDomainPolicyOutput {
	return o
}

func (o ServicePlacementInvalidDomainPolicyOutput) ToServicePlacementInvalidDomainPolicyOutputWithContext(ctx context.Context) ServicePlacementInvalidDomainPolicyOutput {
	return o
}

func (o ServicePlacementInvalidDomainPolicyOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePlacementInvalidDomainPolicy) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o ServicePlacementInvalidDomainPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePlacementInvalidDomainPolicy) string { return v.Type }).(pulumi.StringOutput)
}

type ServicePlacementInvalidDomainPolicyResponse struct {
	DomainName string `pulumi:"domainName"`
	Type       string `pulumi:"type"`
}





type ServicePlacementInvalidDomainPolicyResponseInput interface {
	pulumi.Input

	ToServicePlacementInvalidDomainPolicyResponseOutput() ServicePlacementInvalidDomainPolicyResponseOutput
	ToServicePlacementInvalidDomainPolicyResponseOutputWithContext(context.Context) ServicePlacementInvalidDomainPolicyResponseOutput
}

type ServicePlacementInvalidDomainPolicyResponseArgs struct {
	DomainName pulumi.StringInput `pulumi:"domainName"`
	Type       pulumi.StringInput `pulumi:"type"`
}

func (ServicePlacementInvalidDomainPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementInvalidDomainPolicyResponse)(nil)).Elem()
}

func (i ServicePlacementInvalidDomainPolicyResponseArgs) ToServicePlacementInvalidDomainPolicyResponseOutput() ServicePlacementInvalidDomainPolicyResponseOutput {
	return i.ToServicePlacementInvalidDomainPolicyResponseOutputWithContext(context.Background())
}

func (i ServicePlacementInvalidDomainPolicyResponseArgs) ToServicePlacementInvalidDomainPolicyResponseOutputWithContext(ctx context.Context) ServicePlacementInvalidDomainPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePlacementInvalidDomainPolicyResponseOutput)
}

type ServicePlacementInvalidDomainPolicyResponseOutput struct{ *pulumi.OutputState }

func (ServicePlacementInvalidDomainPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementInvalidDomainPolicyResponse)(nil)).Elem()
}

func (o ServicePlacementInvalidDomainPolicyResponseOutput) ToServicePlacementInvalidDomainPolicyResponseOutput() ServicePlacementInvalidDomainPolicyResponseOutput {
	return o
}

func (o ServicePlacementInvalidDomainPolicyResponseOutput) ToServicePlacementInvalidDomainPolicyResponseOutputWithContext(ctx context.Context) ServicePlacementInvalidDomainPolicyResponseOutput {
	return o
}

func (o ServicePlacementInvalidDomainPolicyResponseOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePlacementInvalidDomainPolicyResponse) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o ServicePlacementInvalidDomainPolicyResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePlacementInvalidDomainPolicyResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ServicePlacementNonPartiallyPlaceServicePolicy struct {
	Type string `pulumi:"type"`
}





type ServicePlacementNonPartiallyPlaceServicePolicyInput interface {
	pulumi.Input

	ToServicePlacementNonPartiallyPlaceServicePolicyOutput() ServicePlacementNonPartiallyPlaceServicePolicyOutput
	ToServicePlacementNonPartiallyPlaceServicePolicyOutputWithContext(context.Context) ServicePlacementNonPartiallyPlaceServicePolicyOutput
}

type ServicePlacementNonPartiallyPlaceServicePolicyArgs struct {
	Type pulumi.StringInput `pulumi:"type"`
}

func (ServicePlacementNonPartiallyPlaceServicePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementNonPartiallyPlaceServicePolicy)(nil)).Elem()
}

func (i ServicePlacementNonPartiallyPlaceServicePolicyArgs) ToServicePlacementNonPartiallyPlaceServicePolicyOutput() ServicePlacementNonPartiallyPlaceServicePolicyOutput {
	return i.ToServicePlacementNonPartiallyPlaceServicePolicyOutputWithContext(context.Background())
}

func (i ServicePlacementNonPartiallyPlaceServicePolicyArgs) ToServicePlacementNonPartiallyPlaceServicePolicyOutputWithContext(ctx context.Context) ServicePlacementNonPartiallyPlaceServicePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePlacementNonPartiallyPlaceServicePolicyOutput)
}

type ServicePlacementNonPartiallyPlaceServicePolicyOutput struct{ *pulumi.OutputState }

func (ServicePlacementNonPartiallyPlaceServicePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementNonPartiallyPlaceServicePolicy)(nil)).Elem()
}

func (o ServicePlacementNonPartiallyPlaceServicePolicyOutput) ToServicePlacementNonPartiallyPlaceServicePolicyOutput() ServicePlacementNonPartiallyPlaceServicePolicyOutput {
	return o
}

func (o ServicePlacementNonPartiallyPlaceServicePolicyOutput) ToServicePlacementNonPartiallyPlaceServicePolicyOutputWithContext(ctx context.Context) ServicePlacementNonPartiallyPlaceServicePolicyOutput {
	return o
}

func (o ServicePlacementNonPartiallyPlaceServicePolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePlacementNonPartiallyPlaceServicePolicy) string { return v.Type }).(pulumi.StringOutput)
}

type ServicePlacementNonPartiallyPlaceServicePolicyResponse struct {
	Type string `pulumi:"type"`
}





type ServicePlacementNonPartiallyPlaceServicePolicyResponseInput interface {
	pulumi.Input

	ToServicePlacementNonPartiallyPlaceServicePolicyResponseOutput() ServicePlacementNonPartiallyPlaceServicePolicyResponseOutput
	ToServicePlacementNonPartiallyPlaceServicePolicyResponseOutputWithContext(context.Context) ServicePlacementNonPartiallyPlaceServicePolicyResponseOutput
}

type ServicePlacementNonPartiallyPlaceServicePolicyResponseArgs struct {
	Type pulumi.StringInput `pulumi:"type"`
}

func (ServicePlacementNonPartiallyPlaceServicePolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementNonPartiallyPlaceServicePolicyResponse)(nil)).Elem()
}

func (i ServicePlacementNonPartiallyPlaceServicePolicyResponseArgs) ToServicePlacementNonPartiallyPlaceServicePolicyResponseOutput() ServicePlacementNonPartiallyPlaceServicePolicyResponseOutput {
	return i.ToServicePlacementNonPartiallyPlaceServicePolicyResponseOutputWithContext(context.Background())
}

func (i ServicePlacementNonPartiallyPlaceServicePolicyResponseArgs) ToServicePlacementNonPartiallyPlaceServicePolicyResponseOutputWithContext(ctx context.Context) ServicePlacementNonPartiallyPlaceServicePolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePlacementNonPartiallyPlaceServicePolicyResponseOutput)
}

type ServicePlacementNonPartiallyPlaceServicePolicyResponseOutput struct{ *pulumi.OutputState }

func (ServicePlacementNonPartiallyPlaceServicePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementNonPartiallyPlaceServicePolicyResponse)(nil)).Elem()
}

func (o ServicePlacementNonPartiallyPlaceServicePolicyResponseOutput) ToServicePlacementNonPartiallyPlaceServicePolicyResponseOutput() ServicePlacementNonPartiallyPlaceServicePolicyResponseOutput {
	return o
}

func (o ServicePlacementNonPartiallyPlaceServicePolicyResponseOutput) ToServicePlacementNonPartiallyPlaceServicePolicyResponseOutputWithContext(ctx context.Context) ServicePlacementNonPartiallyPlaceServicePolicyResponseOutput {
	return o
}

func (o ServicePlacementNonPartiallyPlaceServicePolicyResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePlacementNonPartiallyPlaceServicePolicyResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ServicePlacementPreferPrimaryDomainPolicy struct {
	DomainName string `pulumi:"domainName"`
	Type       string `pulumi:"type"`
}





type ServicePlacementPreferPrimaryDomainPolicyInput interface {
	pulumi.Input

	ToServicePlacementPreferPrimaryDomainPolicyOutput() ServicePlacementPreferPrimaryDomainPolicyOutput
	ToServicePlacementPreferPrimaryDomainPolicyOutputWithContext(context.Context) ServicePlacementPreferPrimaryDomainPolicyOutput
}

type ServicePlacementPreferPrimaryDomainPolicyArgs struct {
	DomainName pulumi.StringInput `pulumi:"domainName"`
	Type       pulumi.StringInput `pulumi:"type"`
}

func (ServicePlacementPreferPrimaryDomainPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementPreferPrimaryDomainPolicy)(nil)).Elem()
}

func (i ServicePlacementPreferPrimaryDomainPolicyArgs) ToServicePlacementPreferPrimaryDomainPolicyOutput() ServicePlacementPreferPrimaryDomainPolicyOutput {
	return i.ToServicePlacementPreferPrimaryDomainPolicyOutputWithContext(context.Background())
}

func (i ServicePlacementPreferPrimaryDomainPolicyArgs) ToServicePlacementPreferPrimaryDomainPolicyOutputWithContext(ctx context.Context) ServicePlacementPreferPrimaryDomainPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePlacementPreferPrimaryDomainPolicyOutput)
}

type ServicePlacementPreferPrimaryDomainPolicyOutput struct{ *pulumi.OutputState }

func (ServicePlacementPreferPrimaryDomainPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementPreferPrimaryDomainPolicy)(nil)).Elem()
}

func (o ServicePlacementPreferPrimaryDomainPolicyOutput) ToServicePlacementPreferPrimaryDomainPolicyOutput() ServicePlacementPreferPrimaryDomainPolicyOutput {
	return o
}

func (o ServicePlacementPreferPrimaryDomainPolicyOutput) ToServicePlacementPreferPrimaryDomainPolicyOutputWithContext(ctx context.Context) ServicePlacementPreferPrimaryDomainPolicyOutput {
	return o
}

func (o ServicePlacementPreferPrimaryDomainPolicyOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePlacementPreferPrimaryDomainPolicy) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o ServicePlacementPreferPrimaryDomainPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePlacementPreferPrimaryDomainPolicy) string { return v.Type }).(pulumi.StringOutput)
}

type ServicePlacementPreferPrimaryDomainPolicyResponse struct {
	DomainName string `pulumi:"domainName"`
	Type       string `pulumi:"type"`
}





type ServicePlacementPreferPrimaryDomainPolicyResponseInput interface {
	pulumi.Input

	ToServicePlacementPreferPrimaryDomainPolicyResponseOutput() ServicePlacementPreferPrimaryDomainPolicyResponseOutput
	ToServicePlacementPreferPrimaryDomainPolicyResponseOutputWithContext(context.Context) ServicePlacementPreferPrimaryDomainPolicyResponseOutput
}

type ServicePlacementPreferPrimaryDomainPolicyResponseArgs struct {
	DomainName pulumi.StringInput `pulumi:"domainName"`
	Type       pulumi.StringInput `pulumi:"type"`
}

func (ServicePlacementPreferPrimaryDomainPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementPreferPrimaryDomainPolicyResponse)(nil)).Elem()
}

func (i ServicePlacementPreferPrimaryDomainPolicyResponseArgs) ToServicePlacementPreferPrimaryDomainPolicyResponseOutput() ServicePlacementPreferPrimaryDomainPolicyResponseOutput {
	return i.ToServicePlacementPreferPrimaryDomainPolicyResponseOutputWithContext(context.Background())
}

func (i ServicePlacementPreferPrimaryDomainPolicyResponseArgs) ToServicePlacementPreferPrimaryDomainPolicyResponseOutputWithContext(ctx context.Context) ServicePlacementPreferPrimaryDomainPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePlacementPreferPrimaryDomainPolicyResponseOutput)
}

type ServicePlacementPreferPrimaryDomainPolicyResponseOutput struct{ *pulumi.OutputState }

func (ServicePlacementPreferPrimaryDomainPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementPreferPrimaryDomainPolicyResponse)(nil)).Elem()
}

func (o ServicePlacementPreferPrimaryDomainPolicyResponseOutput) ToServicePlacementPreferPrimaryDomainPolicyResponseOutput() ServicePlacementPreferPrimaryDomainPolicyResponseOutput {
	return o
}

func (o ServicePlacementPreferPrimaryDomainPolicyResponseOutput) ToServicePlacementPreferPrimaryDomainPolicyResponseOutputWithContext(ctx context.Context) ServicePlacementPreferPrimaryDomainPolicyResponseOutput {
	return o
}

func (o ServicePlacementPreferPrimaryDomainPolicyResponseOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePlacementPreferPrimaryDomainPolicyResponse) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o ServicePlacementPreferPrimaryDomainPolicyResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePlacementPreferPrimaryDomainPolicyResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ServicePlacementRequireDomainDistributionPolicy struct {
	DomainName string `pulumi:"domainName"`
	Type       string `pulumi:"type"`
}





type ServicePlacementRequireDomainDistributionPolicyInput interface {
	pulumi.Input

	ToServicePlacementRequireDomainDistributionPolicyOutput() ServicePlacementRequireDomainDistributionPolicyOutput
	ToServicePlacementRequireDomainDistributionPolicyOutputWithContext(context.Context) ServicePlacementRequireDomainDistributionPolicyOutput
}

type ServicePlacementRequireDomainDistributionPolicyArgs struct {
	DomainName pulumi.StringInput `pulumi:"domainName"`
	Type       pulumi.StringInput `pulumi:"type"`
}

func (ServicePlacementRequireDomainDistributionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementRequireDomainDistributionPolicy)(nil)).Elem()
}

func (i ServicePlacementRequireDomainDistributionPolicyArgs) ToServicePlacementRequireDomainDistributionPolicyOutput() ServicePlacementRequireDomainDistributionPolicyOutput {
	return i.ToServicePlacementRequireDomainDistributionPolicyOutputWithContext(context.Background())
}

func (i ServicePlacementRequireDomainDistributionPolicyArgs) ToServicePlacementRequireDomainDistributionPolicyOutputWithContext(ctx context.Context) ServicePlacementRequireDomainDistributionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePlacementRequireDomainDistributionPolicyOutput)
}

type ServicePlacementRequireDomainDistributionPolicyOutput struct{ *pulumi.OutputState }

func (ServicePlacementRequireDomainDistributionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementRequireDomainDistributionPolicy)(nil)).Elem()
}

func (o ServicePlacementRequireDomainDistributionPolicyOutput) ToServicePlacementRequireDomainDistributionPolicyOutput() ServicePlacementRequireDomainDistributionPolicyOutput {
	return o
}

func (o ServicePlacementRequireDomainDistributionPolicyOutput) ToServicePlacementRequireDomainDistributionPolicyOutputWithContext(ctx context.Context) ServicePlacementRequireDomainDistributionPolicyOutput {
	return o
}

func (o ServicePlacementRequireDomainDistributionPolicyOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePlacementRequireDomainDistributionPolicy) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o ServicePlacementRequireDomainDistributionPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePlacementRequireDomainDistributionPolicy) string { return v.Type }).(pulumi.StringOutput)
}

type ServicePlacementRequireDomainDistributionPolicyResponse struct {
	DomainName string `pulumi:"domainName"`
	Type       string `pulumi:"type"`
}





type ServicePlacementRequireDomainDistributionPolicyResponseInput interface {
	pulumi.Input

	ToServicePlacementRequireDomainDistributionPolicyResponseOutput() ServicePlacementRequireDomainDistributionPolicyResponseOutput
	ToServicePlacementRequireDomainDistributionPolicyResponseOutputWithContext(context.Context) ServicePlacementRequireDomainDistributionPolicyResponseOutput
}

type ServicePlacementRequireDomainDistributionPolicyResponseArgs struct {
	DomainName pulumi.StringInput `pulumi:"domainName"`
	Type       pulumi.StringInput `pulumi:"type"`
}

func (ServicePlacementRequireDomainDistributionPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementRequireDomainDistributionPolicyResponse)(nil)).Elem()
}

func (i ServicePlacementRequireDomainDistributionPolicyResponseArgs) ToServicePlacementRequireDomainDistributionPolicyResponseOutput() ServicePlacementRequireDomainDistributionPolicyResponseOutput {
	return i.ToServicePlacementRequireDomainDistributionPolicyResponseOutputWithContext(context.Background())
}

func (i ServicePlacementRequireDomainDistributionPolicyResponseArgs) ToServicePlacementRequireDomainDistributionPolicyResponseOutputWithContext(ctx context.Context) ServicePlacementRequireDomainDistributionPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePlacementRequireDomainDistributionPolicyResponseOutput)
}

type ServicePlacementRequireDomainDistributionPolicyResponseOutput struct{ *pulumi.OutputState }

func (ServicePlacementRequireDomainDistributionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementRequireDomainDistributionPolicyResponse)(nil)).Elem()
}

func (o ServicePlacementRequireDomainDistributionPolicyResponseOutput) ToServicePlacementRequireDomainDistributionPolicyResponseOutput() ServicePlacementRequireDomainDistributionPolicyResponseOutput {
	return o
}

func (o ServicePlacementRequireDomainDistributionPolicyResponseOutput) ToServicePlacementRequireDomainDistributionPolicyResponseOutputWithContext(ctx context.Context) ServicePlacementRequireDomainDistributionPolicyResponseOutput {
	return o
}

func (o ServicePlacementRequireDomainDistributionPolicyResponseOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePlacementRequireDomainDistributionPolicyResponse) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o ServicePlacementRequireDomainDistributionPolicyResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePlacementRequireDomainDistributionPolicyResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ServicePlacementRequiredDomainPolicy struct {
	DomainName string `pulumi:"domainName"`
	Type       string `pulumi:"type"`
}





type ServicePlacementRequiredDomainPolicyInput interface {
	pulumi.Input

	ToServicePlacementRequiredDomainPolicyOutput() ServicePlacementRequiredDomainPolicyOutput
	ToServicePlacementRequiredDomainPolicyOutputWithContext(context.Context) ServicePlacementRequiredDomainPolicyOutput
}

type ServicePlacementRequiredDomainPolicyArgs struct {
	DomainName pulumi.StringInput `pulumi:"domainName"`
	Type       pulumi.StringInput `pulumi:"type"`
}

func (ServicePlacementRequiredDomainPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementRequiredDomainPolicy)(nil)).Elem()
}

func (i ServicePlacementRequiredDomainPolicyArgs) ToServicePlacementRequiredDomainPolicyOutput() ServicePlacementRequiredDomainPolicyOutput {
	return i.ToServicePlacementRequiredDomainPolicyOutputWithContext(context.Background())
}

func (i ServicePlacementRequiredDomainPolicyArgs) ToServicePlacementRequiredDomainPolicyOutputWithContext(ctx context.Context) ServicePlacementRequiredDomainPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePlacementRequiredDomainPolicyOutput)
}

type ServicePlacementRequiredDomainPolicyOutput struct{ *pulumi.OutputState }

func (ServicePlacementRequiredDomainPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementRequiredDomainPolicy)(nil)).Elem()
}

func (o ServicePlacementRequiredDomainPolicyOutput) ToServicePlacementRequiredDomainPolicyOutput() ServicePlacementRequiredDomainPolicyOutput {
	return o
}

func (o ServicePlacementRequiredDomainPolicyOutput) ToServicePlacementRequiredDomainPolicyOutputWithContext(ctx context.Context) ServicePlacementRequiredDomainPolicyOutput {
	return o
}

func (o ServicePlacementRequiredDomainPolicyOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePlacementRequiredDomainPolicy) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o ServicePlacementRequiredDomainPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePlacementRequiredDomainPolicy) string { return v.Type }).(pulumi.StringOutput)
}

type ServicePlacementRequiredDomainPolicyResponse struct {
	DomainName string `pulumi:"domainName"`
	Type       string `pulumi:"type"`
}





type ServicePlacementRequiredDomainPolicyResponseInput interface {
	pulumi.Input

	ToServicePlacementRequiredDomainPolicyResponseOutput() ServicePlacementRequiredDomainPolicyResponseOutput
	ToServicePlacementRequiredDomainPolicyResponseOutputWithContext(context.Context) ServicePlacementRequiredDomainPolicyResponseOutput
}

type ServicePlacementRequiredDomainPolicyResponseArgs struct {
	DomainName pulumi.StringInput `pulumi:"domainName"`
	Type       pulumi.StringInput `pulumi:"type"`
}

func (ServicePlacementRequiredDomainPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementRequiredDomainPolicyResponse)(nil)).Elem()
}

func (i ServicePlacementRequiredDomainPolicyResponseArgs) ToServicePlacementRequiredDomainPolicyResponseOutput() ServicePlacementRequiredDomainPolicyResponseOutput {
	return i.ToServicePlacementRequiredDomainPolicyResponseOutputWithContext(context.Background())
}

func (i ServicePlacementRequiredDomainPolicyResponseArgs) ToServicePlacementRequiredDomainPolicyResponseOutputWithContext(ctx context.Context) ServicePlacementRequiredDomainPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePlacementRequiredDomainPolicyResponseOutput)
}

type ServicePlacementRequiredDomainPolicyResponseOutput struct{ *pulumi.OutputState }

func (ServicePlacementRequiredDomainPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementRequiredDomainPolicyResponse)(nil)).Elem()
}

func (o ServicePlacementRequiredDomainPolicyResponseOutput) ToServicePlacementRequiredDomainPolicyResponseOutput() ServicePlacementRequiredDomainPolicyResponseOutput {
	return o
}

func (o ServicePlacementRequiredDomainPolicyResponseOutput) ToServicePlacementRequiredDomainPolicyResponseOutputWithContext(ctx context.Context) ServicePlacementRequiredDomainPolicyResponseOutput {
	return o
}

func (o ServicePlacementRequiredDomainPolicyResponseOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePlacementRequiredDomainPolicyResponse) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o ServicePlacementRequiredDomainPolicyResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServicePlacementRequiredDomainPolicyResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ServiceTypeHealthPolicy struct {
	MaxPercentUnhealthyPartitionsPerService int `pulumi:"maxPercentUnhealthyPartitionsPerService"`
	MaxPercentUnhealthyReplicasPerPartition int `pulumi:"maxPercentUnhealthyReplicasPerPartition"`
	MaxPercentUnhealthyServices             int `pulumi:"maxPercentUnhealthyServices"`
}





type ServiceTypeHealthPolicyInput interface {
	pulumi.Input

	ToServiceTypeHealthPolicyOutput() ServiceTypeHealthPolicyOutput
	ToServiceTypeHealthPolicyOutputWithContext(context.Context) ServiceTypeHealthPolicyOutput
}

type ServiceTypeHealthPolicyArgs struct {
	MaxPercentUnhealthyPartitionsPerService pulumi.IntInput `pulumi:"maxPercentUnhealthyPartitionsPerService"`
	MaxPercentUnhealthyReplicasPerPartition pulumi.IntInput `pulumi:"maxPercentUnhealthyReplicasPerPartition"`
	MaxPercentUnhealthyServices             pulumi.IntInput `pulumi:"maxPercentUnhealthyServices"`
}

func (ServiceTypeHealthPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTypeHealthPolicy)(nil)).Elem()
}

func (i ServiceTypeHealthPolicyArgs) ToServiceTypeHealthPolicyOutput() ServiceTypeHealthPolicyOutput {
	return i.ToServiceTypeHealthPolicyOutputWithContext(context.Background())
}

func (i ServiceTypeHealthPolicyArgs) ToServiceTypeHealthPolicyOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeHealthPolicyOutput)
}

func (i ServiceTypeHealthPolicyArgs) ToServiceTypeHealthPolicyPtrOutput() ServiceTypeHealthPolicyPtrOutput {
	return i.ToServiceTypeHealthPolicyPtrOutputWithContext(context.Background())
}

func (i ServiceTypeHealthPolicyArgs) ToServiceTypeHealthPolicyPtrOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeHealthPolicyOutput).ToServiceTypeHealthPolicyPtrOutputWithContext(ctx)
}









type ServiceTypeHealthPolicyPtrInput interface {
	pulumi.Input

	ToServiceTypeHealthPolicyPtrOutput() ServiceTypeHealthPolicyPtrOutput
	ToServiceTypeHealthPolicyPtrOutputWithContext(context.Context) ServiceTypeHealthPolicyPtrOutput
}

type serviceTypeHealthPolicyPtrType ServiceTypeHealthPolicyArgs

func ServiceTypeHealthPolicyPtr(v *ServiceTypeHealthPolicyArgs) ServiceTypeHealthPolicyPtrInput {
	return (*serviceTypeHealthPolicyPtrType)(v)
}

func (*serviceTypeHealthPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceTypeHealthPolicy)(nil)).Elem()
}

func (i *serviceTypeHealthPolicyPtrType) ToServiceTypeHealthPolicyPtrOutput() ServiceTypeHealthPolicyPtrOutput {
	return i.ToServiceTypeHealthPolicyPtrOutputWithContext(context.Background())
}

func (i *serviceTypeHealthPolicyPtrType) ToServiceTypeHealthPolicyPtrOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeHealthPolicyPtrOutput)
}





type ServiceTypeHealthPolicyMapInput interface {
	pulumi.Input

	ToServiceTypeHealthPolicyMapOutput() ServiceTypeHealthPolicyMapOutput
	ToServiceTypeHealthPolicyMapOutputWithContext(context.Context) ServiceTypeHealthPolicyMapOutput
}

type ServiceTypeHealthPolicyMap map[string]ServiceTypeHealthPolicyInput

func (ServiceTypeHealthPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceTypeHealthPolicy)(nil)).Elem()
}

func (i ServiceTypeHealthPolicyMap) ToServiceTypeHealthPolicyMapOutput() ServiceTypeHealthPolicyMapOutput {
	return i.ToServiceTypeHealthPolicyMapOutputWithContext(context.Background())
}

func (i ServiceTypeHealthPolicyMap) ToServiceTypeHealthPolicyMapOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeHealthPolicyMapOutput)
}

type ServiceTypeHealthPolicyOutput struct{ *pulumi.OutputState }

func (ServiceTypeHealthPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTypeHealthPolicy)(nil)).Elem()
}

func (o ServiceTypeHealthPolicyOutput) ToServiceTypeHealthPolicyOutput() ServiceTypeHealthPolicyOutput {
	return o
}

func (o ServiceTypeHealthPolicyOutput) ToServiceTypeHealthPolicyOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyOutput {
	return o
}

func (o ServiceTypeHealthPolicyOutput) ToServiceTypeHealthPolicyPtrOutput() ServiceTypeHealthPolicyPtrOutput {
	return o.ToServiceTypeHealthPolicyPtrOutputWithContext(context.Background())
}

func (o ServiceTypeHealthPolicyOutput) ToServiceTypeHealthPolicyPtrOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceTypeHealthPolicy) *ServiceTypeHealthPolicy {
		return &v
	}).(ServiceTypeHealthPolicyPtrOutput)
}

func (o ServiceTypeHealthPolicyOutput) MaxPercentUnhealthyPartitionsPerService() pulumi.IntOutput {
	return o.ApplyT(func(v ServiceTypeHealthPolicy) int { return v.MaxPercentUnhealthyPartitionsPerService }).(pulumi.IntOutput)
}

func (o ServiceTypeHealthPolicyOutput) MaxPercentUnhealthyReplicasPerPartition() pulumi.IntOutput {
	return o.ApplyT(func(v ServiceTypeHealthPolicy) int { return v.MaxPercentUnhealthyReplicasPerPartition }).(pulumi.IntOutput)
}

func (o ServiceTypeHealthPolicyOutput) MaxPercentUnhealthyServices() pulumi.IntOutput {
	return o.ApplyT(func(v ServiceTypeHealthPolicy) int { return v.MaxPercentUnhealthyServices }).(pulumi.IntOutput)
}

type ServiceTypeHealthPolicyPtrOutput struct{ *pulumi.OutputState }

func (ServiceTypeHealthPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceTypeHealthPolicy)(nil)).Elem()
}

func (o ServiceTypeHealthPolicyPtrOutput) ToServiceTypeHealthPolicyPtrOutput() ServiceTypeHealthPolicyPtrOutput {
	return o
}

func (o ServiceTypeHealthPolicyPtrOutput) ToServiceTypeHealthPolicyPtrOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyPtrOutput {
	return o
}

func (o ServiceTypeHealthPolicyPtrOutput) Elem() ServiceTypeHealthPolicyOutput {
	return o.ApplyT(func(v *ServiceTypeHealthPolicy) ServiceTypeHealthPolicy {
		if v != nil {
			return *v
		}
		var ret ServiceTypeHealthPolicy
		return ret
	}).(ServiceTypeHealthPolicyOutput)
}

func (o ServiceTypeHealthPolicyPtrOutput) MaxPercentUnhealthyPartitionsPerService() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceTypeHealthPolicy) *int {
		if v == nil {
			return nil
		}
		return &v.MaxPercentUnhealthyPartitionsPerService
	}).(pulumi.IntPtrOutput)
}

func (o ServiceTypeHealthPolicyPtrOutput) MaxPercentUnhealthyReplicasPerPartition() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceTypeHealthPolicy) *int {
		if v == nil {
			return nil
		}
		return &v.MaxPercentUnhealthyReplicasPerPartition
	}).(pulumi.IntPtrOutput)
}

func (o ServiceTypeHealthPolicyPtrOutput) MaxPercentUnhealthyServices() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceTypeHealthPolicy) *int {
		if v == nil {
			return nil
		}
		return &v.MaxPercentUnhealthyServices
	}).(pulumi.IntPtrOutput)
}

type ServiceTypeHealthPolicyMapOutput struct{ *pulumi.OutputState }

func (ServiceTypeHealthPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceTypeHealthPolicy)(nil)).Elem()
}

func (o ServiceTypeHealthPolicyMapOutput) ToServiceTypeHealthPolicyMapOutput() ServiceTypeHealthPolicyMapOutput {
	return o
}

func (o ServiceTypeHealthPolicyMapOutput) ToServiceTypeHealthPolicyMapOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyMapOutput {
	return o
}

func (o ServiceTypeHealthPolicyMapOutput) MapIndex(k pulumi.StringInput) ServiceTypeHealthPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServiceTypeHealthPolicy {
		return vs[0].(map[string]ServiceTypeHealthPolicy)[vs[1].(string)]
	}).(ServiceTypeHealthPolicyOutput)
}

type ServiceTypeHealthPolicyResponse struct {
	MaxPercentUnhealthyPartitionsPerService int `pulumi:"maxPercentUnhealthyPartitionsPerService"`
	MaxPercentUnhealthyReplicasPerPartition int `pulumi:"maxPercentUnhealthyReplicasPerPartition"`
	MaxPercentUnhealthyServices             int `pulumi:"maxPercentUnhealthyServices"`
}





type ServiceTypeHealthPolicyResponseInput interface {
	pulumi.Input

	ToServiceTypeHealthPolicyResponseOutput() ServiceTypeHealthPolicyResponseOutput
	ToServiceTypeHealthPolicyResponseOutputWithContext(context.Context) ServiceTypeHealthPolicyResponseOutput
}

type ServiceTypeHealthPolicyResponseArgs struct {
	MaxPercentUnhealthyPartitionsPerService pulumi.IntInput `pulumi:"maxPercentUnhealthyPartitionsPerService"`
	MaxPercentUnhealthyReplicasPerPartition pulumi.IntInput `pulumi:"maxPercentUnhealthyReplicasPerPartition"`
	MaxPercentUnhealthyServices             pulumi.IntInput `pulumi:"maxPercentUnhealthyServices"`
}

func (ServiceTypeHealthPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTypeHealthPolicyResponse)(nil)).Elem()
}

func (i ServiceTypeHealthPolicyResponseArgs) ToServiceTypeHealthPolicyResponseOutput() ServiceTypeHealthPolicyResponseOutput {
	return i.ToServiceTypeHealthPolicyResponseOutputWithContext(context.Background())
}

func (i ServiceTypeHealthPolicyResponseArgs) ToServiceTypeHealthPolicyResponseOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeHealthPolicyResponseOutput)
}

func (i ServiceTypeHealthPolicyResponseArgs) ToServiceTypeHealthPolicyResponsePtrOutput() ServiceTypeHealthPolicyResponsePtrOutput {
	return i.ToServiceTypeHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (i ServiceTypeHealthPolicyResponseArgs) ToServiceTypeHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeHealthPolicyResponseOutput).ToServiceTypeHealthPolicyResponsePtrOutputWithContext(ctx)
}









type ServiceTypeHealthPolicyResponsePtrInput interface {
	pulumi.Input

	ToServiceTypeHealthPolicyResponsePtrOutput() ServiceTypeHealthPolicyResponsePtrOutput
	ToServiceTypeHealthPolicyResponsePtrOutputWithContext(context.Context) ServiceTypeHealthPolicyResponsePtrOutput
}

type serviceTypeHealthPolicyResponsePtrType ServiceTypeHealthPolicyResponseArgs

func ServiceTypeHealthPolicyResponsePtr(v *ServiceTypeHealthPolicyResponseArgs) ServiceTypeHealthPolicyResponsePtrInput {
	return (*serviceTypeHealthPolicyResponsePtrType)(v)
}

func (*serviceTypeHealthPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceTypeHealthPolicyResponse)(nil)).Elem()
}

func (i *serviceTypeHealthPolicyResponsePtrType) ToServiceTypeHealthPolicyResponsePtrOutput() ServiceTypeHealthPolicyResponsePtrOutput {
	return i.ToServiceTypeHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *serviceTypeHealthPolicyResponsePtrType) ToServiceTypeHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeHealthPolicyResponsePtrOutput)
}





type ServiceTypeHealthPolicyResponseMapInput interface {
	pulumi.Input

	ToServiceTypeHealthPolicyResponseMapOutput() ServiceTypeHealthPolicyResponseMapOutput
	ToServiceTypeHealthPolicyResponseMapOutputWithContext(context.Context) ServiceTypeHealthPolicyResponseMapOutput
}

type ServiceTypeHealthPolicyResponseMap map[string]ServiceTypeHealthPolicyResponseInput

func (ServiceTypeHealthPolicyResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceTypeHealthPolicyResponse)(nil)).Elem()
}

func (i ServiceTypeHealthPolicyResponseMap) ToServiceTypeHealthPolicyResponseMapOutput() ServiceTypeHealthPolicyResponseMapOutput {
	return i.ToServiceTypeHealthPolicyResponseMapOutputWithContext(context.Background())
}

func (i ServiceTypeHealthPolicyResponseMap) ToServiceTypeHealthPolicyResponseMapOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTypeHealthPolicyResponseMapOutput)
}

type ServiceTypeHealthPolicyResponseOutput struct{ *pulumi.OutputState }

func (ServiceTypeHealthPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTypeHealthPolicyResponse)(nil)).Elem()
}

func (o ServiceTypeHealthPolicyResponseOutput) ToServiceTypeHealthPolicyResponseOutput() ServiceTypeHealthPolicyResponseOutput {
	return o
}

func (o ServiceTypeHealthPolicyResponseOutput) ToServiceTypeHealthPolicyResponseOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyResponseOutput {
	return o
}

func (o ServiceTypeHealthPolicyResponseOutput) ToServiceTypeHealthPolicyResponsePtrOutput() ServiceTypeHealthPolicyResponsePtrOutput {
	return o.ToServiceTypeHealthPolicyResponsePtrOutputWithContext(context.Background())
}

func (o ServiceTypeHealthPolicyResponseOutput) ToServiceTypeHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceTypeHealthPolicyResponse) *ServiceTypeHealthPolicyResponse {
		return &v
	}).(ServiceTypeHealthPolicyResponsePtrOutput)
}

func (o ServiceTypeHealthPolicyResponseOutput) MaxPercentUnhealthyPartitionsPerService() pulumi.IntOutput {
	return o.ApplyT(func(v ServiceTypeHealthPolicyResponse) int { return v.MaxPercentUnhealthyPartitionsPerService }).(pulumi.IntOutput)
}

func (o ServiceTypeHealthPolicyResponseOutput) MaxPercentUnhealthyReplicasPerPartition() pulumi.IntOutput {
	return o.ApplyT(func(v ServiceTypeHealthPolicyResponse) int { return v.MaxPercentUnhealthyReplicasPerPartition }).(pulumi.IntOutput)
}

func (o ServiceTypeHealthPolicyResponseOutput) MaxPercentUnhealthyServices() pulumi.IntOutput {
	return o.ApplyT(func(v ServiceTypeHealthPolicyResponse) int { return v.MaxPercentUnhealthyServices }).(pulumi.IntOutput)
}

type ServiceTypeHealthPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceTypeHealthPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceTypeHealthPolicyResponse)(nil)).Elem()
}

func (o ServiceTypeHealthPolicyResponsePtrOutput) ToServiceTypeHealthPolicyResponsePtrOutput() ServiceTypeHealthPolicyResponsePtrOutput {
	return o
}

func (o ServiceTypeHealthPolicyResponsePtrOutput) ToServiceTypeHealthPolicyResponsePtrOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyResponsePtrOutput {
	return o
}

func (o ServiceTypeHealthPolicyResponsePtrOutput) Elem() ServiceTypeHealthPolicyResponseOutput {
	return o.ApplyT(func(v *ServiceTypeHealthPolicyResponse) ServiceTypeHealthPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ServiceTypeHealthPolicyResponse
		return ret
	}).(ServiceTypeHealthPolicyResponseOutput)
}

func (o ServiceTypeHealthPolicyResponsePtrOutput) MaxPercentUnhealthyPartitionsPerService() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceTypeHealthPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return &v.MaxPercentUnhealthyPartitionsPerService
	}).(pulumi.IntPtrOutput)
}

func (o ServiceTypeHealthPolicyResponsePtrOutput) MaxPercentUnhealthyReplicasPerPartition() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceTypeHealthPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return &v.MaxPercentUnhealthyReplicasPerPartition
	}).(pulumi.IntPtrOutput)
}

func (o ServiceTypeHealthPolicyResponsePtrOutput) MaxPercentUnhealthyServices() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServiceTypeHealthPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return &v.MaxPercentUnhealthyServices
	}).(pulumi.IntPtrOutput)
}

type ServiceTypeHealthPolicyResponseMapOutput struct{ *pulumi.OutputState }

func (ServiceTypeHealthPolicyResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServiceTypeHealthPolicyResponse)(nil)).Elem()
}

func (o ServiceTypeHealthPolicyResponseMapOutput) ToServiceTypeHealthPolicyResponseMapOutput() ServiceTypeHealthPolicyResponseMapOutput {
	return o
}

func (o ServiceTypeHealthPolicyResponseMapOutput) ToServiceTypeHealthPolicyResponseMapOutputWithContext(ctx context.Context) ServiceTypeHealthPolicyResponseMapOutput {
	return o
}

func (o ServiceTypeHealthPolicyResponseMapOutput) MapIndex(k pulumi.StringInput) ServiceTypeHealthPolicyResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServiceTypeHealthPolicyResponse {
		return vs[0].(map[string]ServiceTypeHealthPolicyResponse)[vs[1].(string)]
	}).(ServiceTypeHealthPolicyResponseOutput)
}

type SettingsParameterDescription struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type SettingsParameterDescriptionInput interface {
	pulumi.Input

	ToSettingsParameterDescriptionOutput() SettingsParameterDescriptionOutput
	ToSettingsParameterDescriptionOutputWithContext(context.Context) SettingsParameterDescriptionOutput
}

type SettingsParameterDescriptionArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (SettingsParameterDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsParameterDescription)(nil)).Elem()
}

func (i SettingsParameterDescriptionArgs) ToSettingsParameterDescriptionOutput() SettingsParameterDescriptionOutput {
	return i.ToSettingsParameterDescriptionOutputWithContext(context.Background())
}

func (i SettingsParameterDescriptionArgs) ToSettingsParameterDescriptionOutputWithContext(ctx context.Context) SettingsParameterDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsParameterDescriptionOutput)
}





type SettingsParameterDescriptionArrayInput interface {
	pulumi.Input

	ToSettingsParameterDescriptionArrayOutput() SettingsParameterDescriptionArrayOutput
	ToSettingsParameterDescriptionArrayOutputWithContext(context.Context) SettingsParameterDescriptionArrayOutput
}

type SettingsParameterDescriptionArray []SettingsParameterDescriptionInput

func (SettingsParameterDescriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsParameterDescription)(nil)).Elem()
}

func (i SettingsParameterDescriptionArray) ToSettingsParameterDescriptionArrayOutput() SettingsParameterDescriptionArrayOutput {
	return i.ToSettingsParameterDescriptionArrayOutputWithContext(context.Background())
}

func (i SettingsParameterDescriptionArray) ToSettingsParameterDescriptionArrayOutputWithContext(ctx context.Context) SettingsParameterDescriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsParameterDescriptionArrayOutput)
}

type SettingsParameterDescriptionOutput struct{ *pulumi.OutputState }

func (SettingsParameterDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsParameterDescription)(nil)).Elem()
}

func (o SettingsParameterDescriptionOutput) ToSettingsParameterDescriptionOutput() SettingsParameterDescriptionOutput {
	return o
}

func (o SettingsParameterDescriptionOutput) ToSettingsParameterDescriptionOutputWithContext(ctx context.Context) SettingsParameterDescriptionOutput {
	return o
}

func (o SettingsParameterDescriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsParameterDescription) string { return v.Name }).(pulumi.StringOutput)
}

func (o SettingsParameterDescriptionOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsParameterDescription) string { return v.Value }).(pulumi.StringOutput)
}

type SettingsParameterDescriptionArrayOutput struct{ *pulumi.OutputState }

func (SettingsParameterDescriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsParameterDescription)(nil)).Elem()
}

func (o SettingsParameterDescriptionArrayOutput) ToSettingsParameterDescriptionArrayOutput() SettingsParameterDescriptionArrayOutput {
	return o
}

func (o SettingsParameterDescriptionArrayOutput) ToSettingsParameterDescriptionArrayOutputWithContext(ctx context.Context) SettingsParameterDescriptionArrayOutput {
	return o
}

func (o SettingsParameterDescriptionArrayOutput) Index(i pulumi.IntInput) SettingsParameterDescriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SettingsParameterDescription {
		return vs[0].([]SettingsParameterDescription)[vs[1].(int)]
	}).(SettingsParameterDescriptionOutput)
}

type SettingsParameterDescriptionResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type SettingsParameterDescriptionResponseInput interface {
	pulumi.Input

	ToSettingsParameterDescriptionResponseOutput() SettingsParameterDescriptionResponseOutput
	ToSettingsParameterDescriptionResponseOutputWithContext(context.Context) SettingsParameterDescriptionResponseOutput
}

type SettingsParameterDescriptionResponseArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (SettingsParameterDescriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsParameterDescriptionResponse)(nil)).Elem()
}

func (i SettingsParameterDescriptionResponseArgs) ToSettingsParameterDescriptionResponseOutput() SettingsParameterDescriptionResponseOutput {
	return i.ToSettingsParameterDescriptionResponseOutputWithContext(context.Background())
}

func (i SettingsParameterDescriptionResponseArgs) ToSettingsParameterDescriptionResponseOutputWithContext(ctx context.Context) SettingsParameterDescriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsParameterDescriptionResponseOutput)
}





type SettingsParameterDescriptionResponseArrayInput interface {
	pulumi.Input

	ToSettingsParameterDescriptionResponseArrayOutput() SettingsParameterDescriptionResponseArrayOutput
	ToSettingsParameterDescriptionResponseArrayOutputWithContext(context.Context) SettingsParameterDescriptionResponseArrayOutput
}

type SettingsParameterDescriptionResponseArray []SettingsParameterDescriptionResponseInput

func (SettingsParameterDescriptionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsParameterDescriptionResponse)(nil)).Elem()
}

func (i SettingsParameterDescriptionResponseArray) ToSettingsParameterDescriptionResponseArrayOutput() SettingsParameterDescriptionResponseArrayOutput {
	return i.ToSettingsParameterDescriptionResponseArrayOutputWithContext(context.Background())
}

func (i SettingsParameterDescriptionResponseArray) ToSettingsParameterDescriptionResponseArrayOutputWithContext(ctx context.Context) SettingsParameterDescriptionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsParameterDescriptionResponseArrayOutput)
}

type SettingsParameterDescriptionResponseOutput struct{ *pulumi.OutputState }

func (SettingsParameterDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsParameterDescriptionResponse)(nil)).Elem()
}

func (o SettingsParameterDescriptionResponseOutput) ToSettingsParameterDescriptionResponseOutput() SettingsParameterDescriptionResponseOutput {
	return o
}

func (o SettingsParameterDescriptionResponseOutput) ToSettingsParameterDescriptionResponseOutputWithContext(ctx context.Context) SettingsParameterDescriptionResponseOutput {
	return o
}

func (o SettingsParameterDescriptionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsParameterDescriptionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SettingsParameterDescriptionResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsParameterDescriptionResponse) string { return v.Value }).(pulumi.StringOutput)
}

type SettingsParameterDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (SettingsParameterDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsParameterDescriptionResponse)(nil)).Elem()
}

func (o SettingsParameterDescriptionResponseArrayOutput) ToSettingsParameterDescriptionResponseArrayOutput() SettingsParameterDescriptionResponseArrayOutput {
	return o
}

func (o SettingsParameterDescriptionResponseArrayOutput) ToSettingsParameterDescriptionResponseArrayOutputWithContext(ctx context.Context) SettingsParameterDescriptionResponseArrayOutput {
	return o
}

func (o SettingsParameterDescriptionResponseArrayOutput) Index(i pulumi.IntInput) SettingsParameterDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SettingsParameterDescriptionResponse {
		return vs[0].([]SettingsParameterDescriptionResponse)[vs[1].(int)]
	}).(SettingsParameterDescriptionResponseOutput)
}

type SettingsSectionDescription struct {
	Name       string                         `pulumi:"name"`
	Parameters []SettingsParameterDescription `pulumi:"parameters"`
}





type SettingsSectionDescriptionInput interface {
	pulumi.Input

	ToSettingsSectionDescriptionOutput() SettingsSectionDescriptionOutput
	ToSettingsSectionDescriptionOutputWithContext(context.Context) SettingsSectionDescriptionOutput
}

type SettingsSectionDescriptionArgs struct {
	Name       pulumi.StringInput                     `pulumi:"name"`
	Parameters SettingsParameterDescriptionArrayInput `pulumi:"parameters"`
}

func (SettingsSectionDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsSectionDescription)(nil)).Elem()
}

func (i SettingsSectionDescriptionArgs) ToSettingsSectionDescriptionOutput() SettingsSectionDescriptionOutput {
	return i.ToSettingsSectionDescriptionOutputWithContext(context.Background())
}

func (i SettingsSectionDescriptionArgs) ToSettingsSectionDescriptionOutputWithContext(ctx context.Context) SettingsSectionDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsSectionDescriptionOutput)
}





type SettingsSectionDescriptionArrayInput interface {
	pulumi.Input

	ToSettingsSectionDescriptionArrayOutput() SettingsSectionDescriptionArrayOutput
	ToSettingsSectionDescriptionArrayOutputWithContext(context.Context) SettingsSectionDescriptionArrayOutput
}

type SettingsSectionDescriptionArray []SettingsSectionDescriptionInput

func (SettingsSectionDescriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsSectionDescription)(nil)).Elem()
}

func (i SettingsSectionDescriptionArray) ToSettingsSectionDescriptionArrayOutput() SettingsSectionDescriptionArrayOutput {
	return i.ToSettingsSectionDescriptionArrayOutputWithContext(context.Background())
}

func (i SettingsSectionDescriptionArray) ToSettingsSectionDescriptionArrayOutputWithContext(ctx context.Context) SettingsSectionDescriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsSectionDescriptionArrayOutput)
}

type SettingsSectionDescriptionOutput struct{ *pulumi.OutputState }

func (SettingsSectionDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsSectionDescription)(nil)).Elem()
}

func (o SettingsSectionDescriptionOutput) ToSettingsSectionDescriptionOutput() SettingsSectionDescriptionOutput {
	return o
}

func (o SettingsSectionDescriptionOutput) ToSettingsSectionDescriptionOutputWithContext(ctx context.Context) SettingsSectionDescriptionOutput {
	return o
}

func (o SettingsSectionDescriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsSectionDescription) string { return v.Name }).(pulumi.StringOutput)
}

func (o SettingsSectionDescriptionOutput) Parameters() SettingsParameterDescriptionArrayOutput {
	return o.ApplyT(func(v SettingsSectionDescription) []SettingsParameterDescription { return v.Parameters }).(SettingsParameterDescriptionArrayOutput)
}

type SettingsSectionDescriptionArrayOutput struct{ *pulumi.OutputState }

func (SettingsSectionDescriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsSectionDescription)(nil)).Elem()
}

func (o SettingsSectionDescriptionArrayOutput) ToSettingsSectionDescriptionArrayOutput() SettingsSectionDescriptionArrayOutput {
	return o
}

func (o SettingsSectionDescriptionArrayOutput) ToSettingsSectionDescriptionArrayOutputWithContext(ctx context.Context) SettingsSectionDescriptionArrayOutput {
	return o
}

func (o SettingsSectionDescriptionArrayOutput) Index(i pulumi.IntInput) SettingsSectionDescriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SettingsSectionDescription {
		return vs[0].([]SettingsSectionDescription)[vs[1].(int)]
	}).(SettingsSectionDescriptionOutput)
}

type SettingsSectionDescriptionResponse struct {
	Name       string                                 `pulumi:"name"`
	Parameters []SettingsParameterDescriptionResponse `pulumi:"parameters"`
}





type SettingsSectionDescriptionResponseInput interface {
	pulumi.Input

	ToSettingsSectionDescriptionResponseOutput() SettingsSectionDescriptionResponseOutput
	ToSettingsSectionDescriptionResponseOutputWithContext(context.Context) SettingsSectionDescriptionResponseOutput
}

type SettingsSectionDescriptionResponseArgs struct {
	Name       pulumi.StringInput                             `pulumi:"name"`
	Parameters SettingsParameterDescriptionResponseArrayInput `pulumi:"parameters"`
}

func (SettingsSectionDescriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsSectionDescriptionResponse)(nil)).Elem()
}

func (i SettingsSectionDescriptionResponseArgs) ToSettingsSectionDescriptionResponseOutput() SettingsSectionDescriptionResponseOutput {
	return i.ToSettingsSectionDescriptionResponseOutputWithContext(context.Background())
}

func (i SettingsSectionDescriptionResponseArgs) ToSettingsSectionDescriptionResponseOutputWithContext(ctx context.Context) SettingsSectionDescriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsSectionDescriptionResponseOutput)
}





type SettingsSectionDescriptionResponseArrayInput interface {
	pulumi.Input

	ToSettingsSectionDescriptionResponseArrayOutput() SettingsSectionDescriptionResponseArrayOutput
	ToSettingsSectionDescriptionResponseArrayOutputWithContext(context.Context) SettingsSectionDescriptionResponseArrayOutput
}

type SettingsSectionDescriptionResponseArray []SettingsSectionDescriptionResponseInput

func (SettingsSectionDescriptionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsSectionDescriptionResponse)(nil)).Elem()
}

func (i SettingsSectionDescriptionResponseArray) ToSettingsSectionDescriptionResponseArrayOutput() SettingsSectionDescriptionResponseArrayOutput {
	return i.ToSettingsSectionDescriptionResponseArrayOutputWithContext(context.Background())
}

func (i SettingsSectionDescriptionResponseArray) ToSettingsSectionDescriptionResponseArrayOutputWithContext(ctx context.Context) SettingsSectionDescriptionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsSectionDescriptionResponseArrayOutput)
}

type SettingsSectionDescriptionResponseOutput struct{ *pulumi.OutputState }

func (SettingsSectionDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsSectionDescriptionResponse)(nil)).Elem()
}

func (o SettingsSectionDescriptionResponseOutput) ToSettingsSectionDescriptionResponseOutput() SettingsSectionDescriptionResponseOutput {
	return o
}

func (o SettingsSectionDescriptionResponseOutput) ToSettingsSectionDescriptionResponseOutputWithContext(ctx context.Context) SettingsSectionDescriptionResponseOutput {
	return o
}

func (o SettingsSectionDescriptionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsSectionDescriptionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SettingsSectionDescriptionResponseOutput) Parameters() SettingsParameterDescriptionResponseArrayOutput {
	return o.ApplyT(func(v SettingsSectionDescriptionResponse) []SettingsParameterDescriptionResponse { return v.Parameters }).(SettingsParameterDescriptionResponseArrayOutput)
}

type SettingsSectionDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (SettingsSectionDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsSectionDescriptionResponse)(nil)).Elem()
}

func (o SettingsSectionDescriptionResponseArrayOutput) ToSettingsSectionDescriptionResponseArrayOutput() SettingsSectionDescriptionResponseArrayOutput {
	return o
}

func (o SettingsSectionDescriptionResponseArrayOutput) ToSettingsSectionDescriptionResponseArrayOutputWithContext(ctx context.Context) SettingsSectionDescriptionResponseArrayOutput {
	return o
}

func (o SettingsSectionDescriptionResponseArrayOutput) Index(i pulumi.IntInput) SettingsSectionDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SettingsSectionDescriptionResponse {
		return vs[0].([]SettingsSectionDescriptionResponse)[vs[1].(int)]
	}).(SettingsSectionDescriptionResponseOutput)
}

type SingletonPartitionScheme struct {
	PartitionScheme string `pulumi:"partitionScheme"`
}





type SingletonPartitionSchemeInput interface {
	pulumi.Input

	ToSingletonPartitionSchemeOutput() SingletonPartitionSchemeOutput
	ToSingletonPartitionSchemeOutputWithContext(context.Context) SingletonPartitionSchemeOutput
}

type SingletonPartitionSchemeArgs struct {
	PartitionScheme pulumi.StringInput `pulumi:"partitionScheme"`
}

func (SingletonPartitionSchemeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SingletonPartitionScheme)(nil)).Elem()
}

func (i SingletonPartitionSchemeArgs) ToSingletonPartitionSchemeOutput() SingletonPartitionSchemeOutput {
	return i.ToSingletonPartitionSchemeOutputWithContext(context.Background())
}

func (i SingletonPartitionSchemeArgs) ToSingletonPartitionSchemeOutputWithContext(ctx context.Context) SingletonPartitionSchemeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SingletonPartitionSchemeOutput)
}

type SingletonPartitionSchemeOutput struct{ *pulumi.OutputState }

func (SingletonPartitionSchemeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SingletonPartitionScheme)(nil)).Elem()
}

func (o SingletonPartitionSchemeOutput) ToSingletonPartitionSchemeOutput() SingletonPartitionSchemeOutput {
	return o
}

func (o SingletonPartitionSchemeOutput) ToSingletonPartitionSchemeOutputWithContext(ctx context.Context) SingletonPartitionSchemeOutput {
	return o
}

func (o SingletonPartitionSchemeOutput) PartitionScheme() pulumi.StringOutput {
	return o.ApplyT(func(v SingletonPartitionScheme) string { return v.PartitionScheme }).(pulumi.StringOutput)
}

type SingletonPartitionSchemeResponse struct {
	PartitionScheme string `pulumi:"partitionScheme"`
}





type SingletonPartitionSchemeResponseInput interface {
	pulumi.Input

	ToSingletonPartitionSchemeResponseOutput() SingletonPartitionSchemeResponseOutput
	ToSingletonPartitionSchemeResponseOutputWithContext(context.Context) SingletonPartitionSchemeResponseOutput
}

type SingletonPartitionSchemeResponseArgs struct {
	PartitionScheme pulumi.StringInput `pulumi:"partitionScheme"`
}

func (SingletonPartitionSchemeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SingletonPartitionSchemeResponse)(nil)).Elem()
}

func (i SingletonPartitionSchemeResponseArgs) ToSingletonPartitionSchemeResponseOutput() SingletonPartitionSchemeResponseOutput {
	return i.ToSingletonPartitionSchemeResponseOutputWithContext(context.Background())
}

func (i SingletonPartitionSchemeResponseArgs) ToSingletonPartitionSchemeResponseOutputWithContext(ctx context.Context) SingletonPartitionSchemeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SingletonPartitionSchemeResponseOutput)
}

type SingletonPartitionSchemeResponseOutput struct{ *pulumi.OutputState }

func (SingletonPartitionSchemeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SingletonPartitionSchemeResponse)(nil)).Elem()
}

func (o SingletonPartitionSchemeResponseOutput) ToSingletonPartitionSchemeResponseOutput() SingletonPartitionSchemeResponseOutput {
	return o
}

func (o SingletonPartitionSchemeResponseOutput) ToSingletonPartitionSchemeResponseOutputWithContext(ctx context.Context) SingletonPartitionSchemeResponseOutput {
	return o
}

func (o SingletonPartitionSchemeResponseOutput) PartitionScheme() pulumi.StringOutput {
	return o.ApplyT(func(v SingletonPartitionSchemeResponse) string { return v.PartitionScheme }).(pulumi.StringOutput)
}

type Sku struct {
	Name string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type StatefulServiceProperties struct {
	CorrelationScheme            []ServiceCorrelation `pulumi:"correlationScheme"`
	DefaultMoveCost              *string              `pulumi:"defaultMoveCost"`
	HasPersistedState            *bool                `pulumi:"hasPersistedState"`
	MinReplicaSetSize            *int                 `pulumi:"minReplicaSetSize"`
	PartitionDescription         interface{}          `pulumi:"partitionDescription"`
	PlacementConstraints         *string              `pulumi:"placementConstraints"`
	QuorumLossWaitDuration       *string              `pulumi:"quorumLossWaitDuration"`
	ReplicaRestartWaitDuration   *string              `pulumi:"replicaRestartWaitDuration"`
	ScalingPolicies              []ScalingPolicy      `pulumi:"scalingPolicies"`
	ServiceKind                  string               `pulumi:"serviceKind"`
	ServiceLoadMetrics           []ServiceLoadMetric  `pulumi:"serviceLoadMetrics"`
	ServicePackageActivationMode *string              `pulumi:"servicePackageActivationMode"`
	ServicePlacementPolicies     []interface{}        `pulumi:"servicePlacementPolicies"`
	ServicePlacementTimeLimit    *string              `pulumi:"servicePlacementTimeLimit"`
	ServiceTypeName              string               `pulumi:"serviceTypeName"`
	StandByReplicaKeepDuration   *string              `pulumi:"standByReplicaKeepDuration"`
	TargetReplicaSetSize         *int                 `pulumi:"targetReplicaSetSize"`
}





type StatefulServicePropertiesInput interface {
	pulumi.Input

	ToStatefulServicePropertiesOutput() StatefulServicePropertiesOutput
	ToStatefulServicePropertiesOutputWithContext(context.Context) StatefulServicePropertiesOutput
}

type StatefulServicePropertiesArgs struct {
	CorrelationScheme            ServiceCorrelationArrayInput `pulumi:"correlationScheme"`
	DefaultMoveCost              pulumi.StringPtrInput        `pulumi:"defaultMoveCost"`
	HasPersistedState            pulumi.BoolPtrInput          `pulumi:"hasPersistedState"`
	MinReplicaSetSize            pulumi.IntPtrInput           `pulumi:"minReplicaSetSize"`
	PartitionDescription         pulumi.Input                 `pulumi:"partitionDescription"`
	PlacementConstraints         pulumi.StringPtrInput        `pulumi:"placementConstraints"`
	QuorumLossWaitDuration       pulumi.StringPtrInput        `pulumi:"quorumLossWaitDuration"`
	ReplicaRestartWaitDuration   pulumi.StringPtrInput        `pulumi:"replicaRestartWaitDuration"`
	ScalingPolicies              ScalingPolicyArrayInput      `pulumi:"scalingPolicies"`
	ServiceKind                  pulumi.StringInput           `pulumi:"serviceKind"`
	ServiceLoadMetrics           ServiceLoadMetricArrayInput  `pulumi:"serviceLoadMetrics"`
	ServicePackageActivationMode pulumi.StringPtrInput        `pulumi:"servicePackageActivationMode"`
	ServicePlacementPolicies     pulumi.ArrayInput            `pulumi:"servicePlacementPolicies"`
	ServicePlacementTimeLimit    pulumi.StringPtrInput        `pulumi:"servicePlacementTimeLimit"`
	ServiceTypeName              pulumi.StringInput           `pulumi:"serviceTypeName"`
	StandByReplicaKeepDuration   pulumi.StringPtrInput        `pulumi:"standByReplicaKeepDuration"`
	TargetReplicaSetSize         pulumi.IntPtrInput           `pulumi:"targetReplicaSetSize"`
}

func (StatefulServicePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StatefulServiceProperties)(nil)).Elem()
}

func (i StatefulServicePropertiesArgs) ToStatefulServicePropertiesOutput() StatefulServicePropertiesOutput {
	return i.ToStatefulServicePropertiesOutputWithContext(context.Background())
}

func (i StatefulServicePropertiesArgs) ToStatefulServicePropertiesOutputWithContext(ctx context.Context) StatefulServicePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatefulServicePropertiesOutput)
}

type StatefulServicePropertiesOutput struct{ *pulumi.OutputState }

func (StatefulServicePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatefulServiceProperties)(nil)).Elem()
}

func (o StatefulServicePropertiesOutput) ToStatefulServicePropertiesOutput() StatefulServicePropertiesOutput {
	return o
}

func (o StatefulServicePropertiesOutput) ToStatefulServicePropertiesOutputWithContext(ctx context.Context) StatefulServicePropertiesOutput {
	return o
}

func (o StatefulServicePropertiesOutput) CorrelationScheme() ServiceCorrelationArrayOutput {
	return o.ApplyT(func(v StatefulServiceProperties) []ServiceCorrelation { return v.CorrelationScheme }).(ServiceCorrelationArrayOutput)
}

func (o StatefulServicePropertiesOutput) DefaultMoveCost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatefulServiceProperties) *string { return v.DefaultMoveCost }).(pulumi.StringPtrOutput)
}

func (o StatefulServicePropertiesOutput) HasPersistedState() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StatefulServiceProperties) *bool { return v.HasPersistedState }).(pulumi.BoolPtrOutput)
}

func (o StatefulServicePropertiesOutput) MinReplicaSetSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatefulServiceProperties) *int { return v.MinReplicaSetSize }).(pulumi.IntPtrOutput)
}

func (o StatefulServicePropertiesOutput) PartitionDescription() pulumi.AnyOutput {
	return o.ApplyT(func(v StatefulServiceProperties) interface{} { return v.PartitionDescription }).(pulumi.AnyOutput)
}

func (o StatefulServicePropertiesOutput) PlacementConstraints() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatefulServiceProperties) *string { return v.PlacementConstraints }).(pulumi.StringPtrOutput)
}

func (o StatefulServicePropertiesOutput) QuorumLossWaitDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatefulServiceProperties) *string { return v.QuorumLossWaitDuration }).(pulumi.StringPtrOutput)
}

func (o StatefulServicePropertiesOutput) ReplicaRestartWaitDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatefulServiceProperties) *string { return v.ReplicaRestartWaitDuration }).(pulumi.StringPtrOutput)
}

func (o StatefulServicePropertiesOutput) ScalingPolicies() ScalingPolicyArrayOutput {
	return o.ApplyT(func(v StatefulServiceProperties) []ScalingPolicy { return v.ScalingPolicies }).(ScalingPolicyArrayOutput)
}

func (o StatefulServicePropertiesOutput) ServiceKind() pulumi.StringOutput {
	return o.ApplyT(func(v StatefulServiceProperties) string { return v.ServiceKind }).(pulumi.StringOutput)
}

func (o StatefulServicePropertiesOutput) ServiceLoadMetrics() ServiceLoadMetricArrayOutput {
	return o.ApplyT(func(v StatefulServiceProperties) []ServiceLoadMetric { return v.ServiceLoadMetrics }).(ServiceLoadMetricArrayOutput)
}

func (o StatefulServicePropertiesOutput) ServicePackageActivationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatefulServiceProperties) *string { return v.ServicePackageActivationMode }).(pulumi.StringPtrOutput)
}

func (o StatefulServicePropertiesOutput) ServicePlacementPolicies() pulumi.ArrayOutput {
	return o.ApplyT(func(v StatefulServiceProperties) []interface{} { return v.ServicePlacementPolicies }).(pulumi.ArrayOutput)
}

func (o StatefulServicePropertiesOutput) ServicePlacementTimeLimit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatefulServiceProperties) *string { return v.ServicePlacementTimeLimit }).(pulumi.StringPtrOutput)
}

func (o StatefulServicePropertiesOutput) ServiceTypeName() pulumi.StringOutput {
	return o.ApplyT(func(v StatefulServiceProperties) string { return v.ServiceTypeName }).(pulumi.StringOutput)
}

func (o StatefulServicePropertiesOutput) StandByReplicaKeepDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatefulServiceProperties) *string { return v.StandByReplicaKeepDuration }).(pulumi.StringPtrOutput)
}

func (o StatefulServicePropertiesOutput) TargetReplicaSetSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatefulServiceProperties) *int { return v.TargetReplicaSetSize }).(pulumi.IntPtrOutput)
}

type StatefulServicePropertiesResponse struct {
	CorrelationScheme            []ServiceCorrelationResponse `pulumi:"correlationScheme"`
	DefaultMoveCost              *string                      `pulumi:"defaultMoveCost"`
	HasPersistedState            *bool                        `pulumi:"hasPersistedState"`
	MinReplicaSetSize            *int                         `pulumi:"minReplicaSetSize"`
	PartitionDescription         interface{}                  `pulumi:"partitionDescription"`
	PlacementConstraints         *string                      `pulumi:"placementConstraints"`
	ProvisioningState            string                       `pulumi:"provisioningState"`
	QuorumLossWaitDuration       *string                      `pulumi:"quorumLossWaitDuration"`
	ReplicaRestartWaitDuration   *string                      `pulumi:"replicaRestartWaitDuration"`
	ScalingPolicies              []ScalingPolicyResponse      `pulumi:"scalingPolicies"`
	ServiceKind                  string                       `pulumi:"serviceKind"`
	ServiceLoadMetrics           []ServiceLoadMetricResponse  `pulumi:"serviceLoadMetrics"`
	ServicePackageActivationMode *string                      `pulumi:"servicePackageActivationMode"`
	ServicePlacementPolicies     []interface{}                `pulumi:"servicePlacementPolicies"`
	ServicePlacementTimeLimit    *string                      `pulumi:"servicePlacementTimeLimit"`
	ServiceTypeName              string                       `pulumi:"serviceTypeName"`
	StandByReplicaKeepDuration   *string                      `pulumi:"standByReplicaKeepDuration"`
	TargetReplicaSetSize         *int                         `pulumi:"targetReplicaSetSize"`
}





type StatefulServicePropertiesResponseInput interface {
	pulumi.Input

	ToStatefulServicePropertiesResponseOutput() StatefulServicePropertiesResponseOutput
	ToStatefulServicePropertiesResponseOutputWithContext(context.Context) StatefulServicePropertiesResponseOutput
}

type StatefulServicePropertiesResponseArgs struct {
	CorrelationScheme            ServiceCorrelationResponseArrayInput `pulumi:"correlationScheme"`
	DefaultMoveCost              pulumi.StringPtrInput                `pulumi:"defaultMoveCost"`
	HasPersistedState            pulumi.BoolPtrInput                  `pulumi:"hasPersistedState"`
	MinReplicaSetSize            pulumi.IntPtrInput                   `pulumi:"minReplicaSetSize"`
	PartitionDescription         pulumi.Input                         `pulumi:"partitionDescription"`
	PlacementConstraints         pulumi.StringPtrInput                `pulumi:"placementConstraints"`
	ProvisioningState            pulumi.StringInput                   `pulumi:"provisioningState"`
	QuorumLossWaitDuration       pulumi.StringPtrInput                `pulumi:"quorumLossWaitDuration"`
	ReplicaRestartWaitDuration   pulumi.StringPtrInput                `pulumi:"replicaRestartWaitDuration"`
	ScalingPolicies              ScalingPolicyResponseArrayInput      `pulumi:"scalingPolicies"`
	ServiceKind                  pulumi.StringInput                   `pulumi:"serviceKind"`
	ServiceLoadMetrics           ServiceLoadMetricResponseArrayInput  `pulumi:"serviceLoadMetrics"`
	ServicePackageActivationMode pulumi.StringPtrInput                `pulumi:"servicePackageActivationMode"`
	ServicePlacementPolicies     pulumi.ArrayInput                    `pulumi:"servicePlacementPolicies"`
	ServicePlacementTimeLimit    pulumi.StringPtrInput                `pulumi:"servicePlacementTimeLimit"`
	ServiceTypeName              pulumi.StringInput                   `pulumi:"serviceTypeName"`
	StandByReplicaKeepDuration   pulumi.StringPtrInput                `pulumi:"standByReplicaKeepDuration"`
	TargetReplicaSetSize         pulumi.IntPtrInput                   `pulumi:"targetReplicaSetSize"`
}

func (StatefulServicePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StatefulServicePropertiesResponse)(nil)).Elem()
}

func (i StatefulServicePropertiesResponseArgs) ToStatefulServicePropertiesResponseOutput() StatefulServicePropertiesResponseOutput {
	return i.ToStatefulServicePropertiesResponseOutputWithContext(context.Background())
}

func (i StatefulServicePropertiesResponseArgs) ToStatefulServicePropertiesResponseOutputWithContext(ctx context.Context) StatefulServicePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatefulServicePropertiesResponseOutput)
}

type StatefulServicePropertiesResponseOutput struct{ *pulumi.OutputState }

func (StatefulServicePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatefulServicePropertiesResponse)(nil)).Elem()
}

func (o StatefulServicePropertiesResponseOutput) ToStatefulServicePropertiesResponseOutput() StatefulServicePropertiesResponseOutput {
	return o
}

func (o StatefulServicePropertiesResponseOutput) ToStatefulServicePropertiesResponseOutputWithContext(ctx context.Context) StatefulServicePropertiesResponseOutput {
	return o
}

func (o StatefulServicePropertiesResponseOutput) CorrelationScheme() ServiceCorrelationResponseArrayOutput {
	return o.ApplyT(func(v StatefulServicePropertiesResponse) []ServiceCorrelationResponse { return v.CorrelationScheme }).(ServiceCorrelationResponseArrayOutput)
}

func (o StatefulServicePropertiesResponseOutput) DefaultMoveCost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatefulServicePropertiesResponse) *string { return v.DefaultMoveCost }).(pulumi.StringPtrOutput)
}

func (o StatefulServicePropertiesResponseOutput) HasPersistedState() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StatefulServicePropertiesResponse) *bool { return v.HasPersistedState }).(pulumi.BoolPtrOutput)
}

func (o StatefulServicePropertiesResponseOutput) MinReplicaSetSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatefulServicePropertiesResponse) *int { return v.MinReplicaSetSize }).(pulumi.IntPtrOutput)
}

func (o StatefulServicePropertiesResponseOutput) PartitionDescription() pulumi.AnyOutput {
	return o.ApplyT(func(v StatefulServicePropertiesResponse) interface{} { return v.PartitionDescription }).(pulumi.AnyOutput)
}

func (o StatefulServicePropertiesResponseOutput) PlacementConstraints() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatefulServicePropertiesResponse) *string { return v.PlacementConstraints }).(pulumi.StringPtrOutput)
}

func (o StatefulServicePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v StatefulServicePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o StatefulServicePropertiesResponseOutput) QuorumLossWaitDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatefulServicePropertiesResponse) *string { return v.QuorumLossWaitDuration }).(pulumi.StringPtrOutput)
}

func (o StatefulServicePropertiesResponseOutput) ReplicaRestartWaitDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatefulServicePropertiesResponse) *string { return v.ReplicaRestartWaitDuration }).(pulumi.StringPtrOutput)
}

func (o StatefulServicePropertiesResponseOutput) ScalingPolicies() ScalingPolicyResponseArrayOutput {
	return o.ApplyT(func(v StatefulServicePropertiesResponse) []ScalingPolicyResponse { return v.ScalingPolicies }).(ScalingPolicyResponseArrayOutput)
}

func (o StatefulServicePropertiesResponseOutput) ServiceKind() pulumi.StringOutput {
	return o.ApplyT(func(v StatefulServicePropertiesResponse) string { return v.ServiceKind }).(pulumi.StringOutput)
}

func (o StatefulServicePropertiesResponseOutput) ServiceLoadMetrics() ServiceLoadMetricResponseArrayOutput {
	return o.ApplyT(func(v StatefulServicePropertiesResponse) []ServiceLoadMetricResponse { return v.ServiceLoadMetrics }).(ServiceLoadMetricResponseArrayOutput)
}

func (o StatefulServicePropertiesResponseOutput) ServicePackageActivationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatefulServicePropertiesResponse) *string { return v.ServicePackageActivationMode }).(pulumi.StringPtrOutput)
}

func (o StatefulServicePropertiesResponseOutput) ServicePlacementPolicies() pulumi.ArrayOutput {
	return o.ApplyT(func(v StatefulServicePropertiesResponse) []interface{} { return v.ServicePlacementPolicies }).(pulumi.ArrayOutput)
}

func (o StatefulServicePropertiesResponseOutput) ServicePlacementTimeLimit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatefulServicePropertiesResponse) *string { return v.ServicePlacementTimeLimit }).(pulumi.StringPtrOutput)
}

func (o StatefulServicePropertiesResponseOutput) ServiceTypeName() pulumi.StringOutput {
	return o.ApplyT(func(v StatefulServicePropertiesResponse) string { return v.ServiceTypeName }).(pulumi.StringOutput)
}

func (o StatefulServicePropertiesResponseOutput) StandByReplicaKeepDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatefulServicePropertiesResponse) *string { return v.StandByReplicaKeepDuration }).(pulumi.StringPtrOutput)
}

func (o StatefulServicePropertiesResponseOutput) TargetReplicaSetSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatefulServicePropertiesResponse) *int { return v.TargetReplicaSetSize }).(pulumi.IntPtrOutput)
}

type StatelessServiceProperties struct {
	CorrelationScheme            []ServiceCorrelation `pulumi:"correlationScheme"`
	DefaultMoveCost              *string              `pulumi:"defaultMoveCost"`
	InstanceCount                int                  `pulumi:"instanceCount"`
	MinInstanceCount             *int                 `pulumi:"minInstanceCount"`
	MinInstancePercentage        *int                 `pulumi:"minInstancePercentage"`
	PartitionDescription         interface{}          `pulumi:"partitionDescription"`
	PlacementConstraints         *string              `pulumi:"placementConstraints"`
	ScalingPolicies              []ScalingPolicy      `pulumi:"scalingPolicies"`
	ServiceKind                  string               `pulumi:"serviceKind"`
	ServiceLoadMetrics           []ServiceLoadMetric  `pulumi:"serviceLoadMetrics"`
	ServicePackageActivationMode *string              `pulumi:"servicePackageActivationMode"`
	ServicePlacementPolicies     []interface{}        `pulumi:"servicePlacementPolicies"`
	ServiceTypeName              string               `pulumi:"serviceTypeName"`
}





type StatelessServicePropertiesInput interface {
	pulumi.Input

	ToStatelessServicePropertiesOutput() StatelessServicePropertiesOutput
	ToStatelessServicePropertiesOutputWithContext(context.Context) StatelessServicePropertiesOutput
}

type StatelessServicePropertiesArgs struct {
	CorrelationScheme            ServiceCorrelationArrayInput `pulumi:"correlationScheme"`
	DefaultMoveCost              pulumi.StringPtrInput        `pulumi:"defaultMoveCost"`
	InstanceCount                pulumi.IntInput              `pulumi:"instanceCount"`
	MinInstanceCount             pulumi.IntPtrInput           `pulumi:"minInstanceCount"`
	MinInstancePercentage        pulumi.IntPtrInput           `pulumi:"minInstancePercentage"`
	PartitionDescription         pulumi.Input                 `pulumi:"partitionDescription"`
	PlacementConstraints         pulumi.StringPtrInput        `pulumi:"placementConstraints"`
	ScalingPolicies              ScalingPolicyArrayInput      `pulumi:"scalingPolicies"`
	ServiceKind                  pulumi.StringInput           `pulumi:"serviceKind"`
	ServiceLoadMetrics           ServiceLoadMetricArrayInput  `pulumi:"serviceLoadMetrics"`
	ServicePackageActivationMode pulumi.StringPtrInput        `pulumi:"servicePackageActivationMode"`
	ServicePlacementPolicies     pulumi.ArrayInput            `pulumi:"servicePlacementPolicies"`
	ServiceTypeName              pulumi.StringInput           `pulumi:"serviceTypeName"`
}

func (StatelessServicePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StatelessServiceProperties)(nil)).Elem()
}

func (i StatelessServicePropertiesArgs) ToStatelessServicePropertiesOutput() StatelessServicePropertiesOutput {
	return i.ToStatelessServicePropertiesOutputWithContext(context.Background())
}

func (i StatelessServicePropertiesArgs) ToStatelessServicePropertiesOutputWithContext(ctx context.Context) StatelessServicePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatelessServicePropertiesOutput)
}

type StatelessServicePropertiesOutput struct{ *pulumi.OutputState }

func (StatelessServicePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatelessServiceProperties)(nil)).Elem()
}

func (o StatelessServicePropertiesOutput) ToStatelessServicePropertiesOutput() StatelessServicePropertiesOutput {
	return o
}

func (o StatelessServicePropertiesOutput) ToStatelessServicePropertiesOutputWithContext(ctx context.Context) StatelessServicePropertiesOutput {
	return o
}

func (o StatelessServicePropertiesOutput) CorrelationScheme() ServiceCorrelationArrayOutput {
	return o.ApplyT(func(v StatelessServiceProperties) []ServiceCorrelation { return v.CorrelationScheme }).(ServiceCorrelationArrayOutput)
}

func (o StatelessServicePropertiesOutput) DefaultMoveCost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatelessServiceProperties) *string { return v.DefaultMoveCost }).(pulumi.StringPtrOutput)
}

func (o StatelessServicePropertiesOutput) InstanceCount() pulumi.IntOutput {
	return o.ApplyT(func(v StatelessServiceProperties) int { return v.InstanceCount }).(pulumi.IntOutput)
}

func (o StatelessServicePropertiesOutput) MinInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatelessServiceProperties) *int { return v.MinInstanceCount }).(pulumi.IntPtrOutput)
}

func (o StatelessServicePropertiesOutput) MinInstancePercentage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatelessServiceProperties) *int { return v.MinInstancePercentage }).(pulumi.IntPtrOutput)
}

func (o StatelessServicePropertiesOutput) PartitionDescription() pulumi.AnyOutput {
	return o.ApplyT(func(v StatelessServiceProperties) interface{} { return v.PartitionDescription }).(pulumi.AnyOutput)
}

func (o StatelessServicePropertiesOutput) PlacementConstraints() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatelessServiceProperties) *string { return v.PlacementConstraints }).(pulumi.StringPtrOutput)
}

func (o StatelessServicePropertiesOutput) ScalingPolicies() ScalingPolicyArrayOutput {
	return o.ApplyT(func(v StatelessServiceProperties) []ScalingPolicy { return v.ScalingPolicies }).(ScalingPolicyArrayOutput)
}

func (o StatelessServicePropertiesOutput) ServiceKind() pulumi.StringOutput {
	return o.ApplyT(func(v StatelessServiceProperties) string { return v.ServiceKind }).(pulumi.StringOutput)
}

func (o StatelessServicePropertiesOutput) ServiceLoadMetrics() ServiceLoadMetricArrayOutput {
	return o.ApplyT(func(v StatelessServiceProperties) []ServiceLoadMetric { return v.ServiceLoadMetrics }).(ServiceLoadMetricArrayOutput)
}

func (o StatelessServicePropertiesOutput) ServicePackageActivationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatelessServiceProperties) *string { return v.ServicePackageActivationMode }).(pulumi.StringPtrOutput)
}

func (o StatelessServicePropertiesOutput) ServicePlacementPolicies() pulumi.ArrayOutput {
	return o.ApplyT(func(v StatelessServiceProperties) []interface{} { return v.ServicePlacementPolicies }).(pulumi.ArrayOutput)
}

func (o StatelessServicePropertiesOutput) ServiceTypeName() pulumi.StringOutput {
	return o.ApplyT(func(v StatelessServiceProperties) string { return v.ServiceTypeName }).(pulumi.StringOutput)
}

type StatelessServicePropertiesResponse struct {
	CorrelationScheme            []ServiceCorrelationResponse `pulumi:"correlationScheme"`
	DefaultMoveCost              *string                      `pulumi:"defaultMoveCost"`
	InstanceCount                int                          `pulumi:"instanceCount"`
	MinInstanceCount             *int                         `pulumi:"minInstanceCount"`
	MinInstancePercentage        *int                         `pulumi:"minInstancePercentage"`
	PartitionDescription         interface{}                  `pulumi:"partitionDescription"`
	PlacementConstraints         *string                      `pulumi:"placementConstraints"`
	ProvisioningState            string                       `pulumi:"provisioningState"`
	ScalingPolicies              []ScalingPolicyResponse      `pulumi:"scalingPolicies"`
	ServiceKind                  string                       `pulumi:"serviceKind"`
	ServiceLoadMetrics           []ServiceLoadMetricResponse  `pulumi:"serviceLoadMetrics"`
	ServicePackageActivationMode *string                      `pulumi:"servicePackageActivationMode"`
	ServicePlacementPolicies     []interface{}                `pulumi:"servicePlacementPolicies"`
	ServiceTypeName              string                       `pulumi:"serviceTypeName"`
}





type StatelessServicePropertiesResponseInput interface {
	pulumi.Input

	ToStatelessServicePropertiesResponseOutput() StatelessServicePropertiesResponseOutput
	ToStatelessServicePropertiesResponseOutputWithContext(context.Context) StatelessServicePropertiesResponseOutput
}

type StatelessServicePropertiesResponseArgs struct {
	CorrelationScheme            ServiceCorrelationResponseArrayInput `pulumi:"correlationScheme"`
	DefaultMoveCost              pulumi.StringPtrInput                `pulumi:"defaultMoveCost"`
	InstanceCount                pulumi.IntInput                      `pulumi:"instanceCount"`
	MinInstanceCount             pulumi.IntPtrInput                   `pulumi:"minInstanceCount"`
	MinInstancePercentage        pulumi.IntPtrInput                   `pulumi:"minInstancePercentage"`
	PartitionDescription         pulumi.Input                         `pulumi:"partitionDescription"`
	PlacementConstraints         pulumi.StringPtrInput                `pulumi:"placementConstraints"`
	ProvisioningState            pulumi.StringInput                   `pulumi:"provisioningState"`
	ScalingPolicies              ScalingPolicyResponseArrayInput      `pulumi:"scalingPolicies"`
	ServiceKind                  pulumi.StringInput                   `pulumi:"serviceKind"`
	ServiceLoadMetrics           ServiceLoadMetricResponseArrayInput  `pulumi:"serviceLoadMetrics"`
	ServicePackageActivationMode pulumi.StringPtrInput                `pulumi:"servicePackageActivationMode"`
	ServicePlacementPolicies     pulumi.ArrayInput                    `pulumi:"servicePlacementPolicies"`
	ServiceTypeName              pulumi.StringInput                   `pulumi:"serviceTypeName"`
}

func (StatelessServicePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StatelessServicePropertiesResponse)(nil)).Elem()
}

func (i StatelessServicePropertiesResponseArgs) ToStatelessServicePropertiesResponseOutput() StatelessServicePropertiesResponseOutput {
	return i.ToStatelessServicePropertiesResponseOutputWithContext(context.Background())
}

func (i StatelessServicePropertiesResponseArgs) ToStatelessServicePropertiesResponseOutputWithContext(ctx context.Context) StatelessServicePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatelessServicePropertiesResponseOutput)
}

type StatelessServicePropertiesResponseOutput struct{ *pulumi.OutputState }

func (StatelessServicePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatelessServicePropertiesResponse)(nil)).Elem()
}

func (o StatelessServicePropertiesResponseOutput) ToStatelessServicePropertiesResponseOutput() StatelessServicePropertiesResponseOutput {
	return o
}

func (o StatelessServicePropertiesResponseOutput) ToStatelessServicePropertiesResponseOutputWithContext(ctx context.Context) StatelessServicePropertiesResponseOutput {
	return o
}

func (o StatelessServicePropertiesResponseOutput) CorrelationScheme() ServiceCorrelationResponseArrayOutput {
	return o.ApplyT(func(v StatelessServicePropertiesResponse) []ServiceCorrelationResponse { return v.CorrelationScheme }).(ServiceCorrelationResponseArrayOutput)
}

func (o StatelessServicePropertiesResponseOutput) DefaultMoveCost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatelessServicePropertiesResponse) *string { return v.DefaultMoveCost }).(pulumi.StringPtrOutput)
}

func (o StatelessServicePropertiesResponseOutput) InstanceCount() pulumi.IntOutput {
	return o.ApplyT(func(v StatelessServicePropertiesResponse) int { return v.InstanceCount }).(pulumi.IntOutput)
}

func (o StatelessServicePropertiesResponseOutput) MinInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatelessServicePropertiesResponse) *int { return v.MinInstanceCount }).(pulumi.IntPtrOutput)
}

func (o StatelessServicePropertiesResponseOutput) MinInstancePercentage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatelessServicePropertiesResponse) *int { return v.MinInstancePercentage }).(pulumi.IntPtrOutput)
}

func (o StatelessServicePropertiesResponseOutput) PartitionDescription() pulumi.AnyOutput {
	return o.ApplyT(func(v StatelessServicePropertiesResponse) interface{} { return v.PartitionDescription }).(pulumi.AnyOutput)
}

func (o StatelessServicePropertiesResponseOutput) PlacementConstraints() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatelessServicePropertiesResponse) *string { return v.PlacementConstraints }).(pulumi.StringPtrOutput)
}

func (o StatelessServicePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v StatelessServicePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o StatelessServicePropertiesResponseOutput) ScalingPolicies() ScalingPolicyResponseArrayOutput {
	return o.ApplyT(func(v StatelessServicePropertiesResponse) []ScalingPolicyResponse { return v.ScalingPolicies }).(ScalingPolicyResponseArrayOutput)
}

func (o StatelessServicePropertiesResponseOutput) ServiceKind() pulumi.StringOutput {
	return o.ApplyT(func(v StatelessServicePropertiesResponse) string { return v.ServiceKind }).(pulumi.StringOutput)
}

func (o StatelessServicePropertiesResponseOutput) ServiceLoadMetrics() ServiceLoadMetricResponseArrayOutput {
	return o.ApplyT(func(v StatelessServicePropertiesResponse) []ServiceLoadMetricResponse { return v.ServiceLoadMetrics }).(ServiceLoadMetricResponseArrayOutput)
}

func (o StatelessServicePropertiesResponseOutput) ServicePackageActivationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatelessServicePropertiesResponse) *string { return v.ServicePackageActivationMode }).(pulumi.StringPtrOutput)
}

func (o StatelessServicePropertiesResponseOutput) ServicePlacementPolicies() pulumi.ArrayOutput {
	return o.ApplyT(func(v StatelessServicePropertiesResponse) []interface{} { return v.ServicePlacementPolicies }).(pulumi.ArrayOutput)
}

func (o StatelessServicePropertiesResponseOutput) ServiceTypeName() pulumi.StringOutput {
	return o.ApplyT(func(v StatelessServicePropertiesResponse) string { return v.ServiceTypeName }).(pulumi.StringOutput)
}

type SubResource struct {
	Id *string `pulumi:"id"`
}





type SubResourceInput interface {
	pulumi.Input

	ToSubResourceOutput() SubResourceOutput
	ToSubResourceOutputWithContext(context.Context) SubResourceOutput
}

type SubResourceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SubResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (i SubResourceArgs) ToSubResourceOutput() SubResourceOutput {
	return i.ToSubResourceOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput)
}

type SubResourceOutput struct{ *pulumi.OutputState }

func (SubResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (o SubResourceOutput) ToSubResourceOutput() SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return o
}

func (o SubResourceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResource) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourceResponse struct {
	Id *string `pulumi:"id"`
}





type SubResourceResponseInput interface {
	pulumi.Input

	ToSubResourceResponseOutput() SubResourceResponseOutput
	ToSubResourceResponseOutputWithContext(context.Context) SubResourceResponseOutput
}

type SubResourceResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SubResourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceResponse)(nil)).Elem()
}

func (i SubResourceResponseArgs) ToSubResourceResponseOutput() SubResourceResponseOutput {
	return i.ToSubResourceResponseOutputWithContext(context.Background())
}

func (i SubResourceResponseArgs) ToSubResourceResponseOutputWithContext(ctx context.Context) SubResourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceResponseOutput)
}

type SubResourceResponseOutput struct{ *pulumi.OutputState }

func (SubResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutput() SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutputWithContext(ctx context.Context) SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResourceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

type UniformInt64RangePartitionScheme struct {
	Count           int     `pulumi:"count"`
	HighKey         float64 `pulumi:"highKey"`
	LowKey          float64 `pulumi:"lowKey"`
	PartitionScheme string  `pulumi:"partitionScheme"`
}





type UniformInt64RangePartitionSchemeInput interface {
	pulumi.Input

	ToUniformInt64RangePartitionSchemeOutput() UniformInt64RangePartitionSchemeOutput
	ToUniformInt64RangePartitionSchemeOutputWithContext(context.Context) UniformInt64RangePartitionSchemeOutput
}

type UniformInt64RangePartitionSchemeArgs struct {
	Count           pulumi.IntInput     `pulumi:"count"`
	HighKey         pulumi.Float64Input `pulumi:"highKey"`
	LowKey          pulumi.Float64Input `pulumi:"lowKey"`
	PartitionScheme pulumi.StringInput  `pulumi:"partitionScheme"`
}

func (UniformInt64RangePartitionSchemeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UniformInt64RangePartitionScheme)(nil)).Elem()
}

func (i UniformInt64RangePartitionSchemeArgs) ToUniformInt64RangePartitionSchemeOutput() UniformInt64RangePartitionSchemeOutput {
	return i.ToUniformInt64RangePartitionSchemeOutputWithContext(context.Background())
}

func (i UniformInt64RangePartitionSchemeArgs) ToUniformInt64RangePartitionSchemeOutputWithContext(ctx context.Context) UniformInt64RangePartitionSchemeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniformInt64RangePartitionSchemeOutput)
}

type UniformInt64RangePartitionSchemeOutput struct{ *pulumi.OutputState }

func (UniformInt64RangePartitionSchemeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UniformInt64RangePartitionScheme)(nil)).Elem()
}

func (o UniformInt64RangePartitionSchemeOutput) ToUniformInt64RangePartitionSchemeOutput() UniformInt64RangePartitionSchemeOutput {
	return o
}

func (o UniformInt64RangePartitionSchemeOutput) ToUniformInt64RangePartitionSchemeOutputWithContext(ctx context.Context) UniformInt64RangePartitionSchemeOutput {
	return o
}

func (o UniformInt64RangePartitionSchemeOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v UniformInt64RangePartitionScheme) int { return v.Count }).(pulumi.IntOutput)
}

func (o UniformInt64RangePartitionSchemeOutput) HighKey() pulumi.Float64Output {
	return o.ApplyT(func(v UniformInt64RangePartitionScheme) float64 { return v.HighKey }).(pulumi.Float64Output)
}

func (o UniformInt64RangePartitionSchemeOutput) LowKey() pulumi.Float64Output {
	return o.ApplyT(func(v UniformInt64RangePartitionScheme) float64 { return v.LowKey }).(pulumi.Float64Output)
}

func (o UniformInt64RangePartitionSchemeOutput) PartitionScheme() pulumi.StringOutput {
	return o.ApplyT(func(v UniformInt64RangePartitionScheme) string { return v.PartitionScheme }).(pulumi.StringOutput)
}

type UniformInt64RangePartitionSchemeResponse struct {
	Count           int     `pulumi:"count"`
	HighKey         float64 `pulumi:"highKey"`
	LowKey          float64 `pulumi:"lowKey"`
	PartitionScheme string  `pulumi:"partitionScheme"`
}





type UniformInt64RangePartitionSchemeResponseInput interface {
	pulumi.Input

	ToUniformInt64RangePartitionSchemeResponseOutput() UniformInt64RangePartitionSchemeResponseOutput
	ToUniformInt64RangePartitionSchemeResponseOutputWithContext(context.Context) UniformInt64RangePartitionSchemeResponseOutput
}

type UniformInt64RangePartitionSchemeResponseArgs struct {
	Count           pulumi.IntInput     `pulumi:"count"`
	HighKey         pulumi.Float64Input `pulumi:"highKey"`
	LowKey          pulumi.Float64Input `pulumi:"lowKey"`
	PartitionScheme pulumi.StringInput  `pulumi:"partitionScheme"`
}

func (UniformInt64RangePartitionSchemeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UniformInt64RangePartitionSchemeResponse)(nil)).Elem()
}

func (i UniformInt64RangePartitionSchemeResponseArgs) ToUniformInt64RangePartitionSchemeResponseOutput() UniformInt64RangePartitionSchemeResponseOutput {
	return i.ToUniformInt64RangePartitionSchemeResponseOutputWithContext(context.Background())
}

func (i UniformInt64RangePartitionSchemeResponseArgs) ToUniformInt64RangePartitionSchemeResponseOutputWithContext(ctx context.Context) UniformInt64RangePartitionSchemeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniformInt64RangePartitionSchemeResponseOutput)
}

type UniformInt64RangePartitionSchemeResponseOutput struct{ *pulumi.OutputState }

func (UniformInt64RangePartitionSchemeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UniformInt64RangePartitionSchemeResponse)(nil)).Elem()
}

func (o UniformInt64RangePartitionSchemeResponseOutput) ToUniformInt64RangePartitionSchemeResponseOutput() UniformInt64RangePartitionSchemeResponseOutput {
	return o
}

func (o UniformInt64RangePartitionSchemeResponseOutput) ToUniformInt64RangePartitionSchemeResponseOutputWithContext(ctx context.Context) UniformInt64RangePartitionSchemeResponseOutput {
	return o
}

func (o UniformInt64RangePartitionSchemeResponseOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v UniformInt64RangePartitionSchemeResponse) int { return v.Count }).(pulumi.IntOutput)
}

func (o UniformInt64RangePartitionSchemeResponseOutput) HighKey() pulumi.Float64Output {
	return o.ApplyT(func(v UniformInt64RangePartitionSchemeResponse) float64 { return v.HighKey }).(pulumi.Float64Output)
}

func (o UniformInt64RangePartitionSchemeResponseOutput) LowKey() pulumi.Float64Output {
	return o.ApplyT(func(v UniformInt64RangePartitionSchemeResponse) float64 { return v.LowKey }).(pulumi.Float64Output)
}

func (o UniformInt64RangePartitionSchemeResponseOutput) PartitionScheme() pulumi.StringOutput {
	return o.ApplyT(func(v UniformInt64RangePartitionSchemeResponse) string { return v.PartitionScheme }).(pulumi.StringOutput)
}

type UserAssignedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}





type UserAssignedIdentityResponseInput interface {
	pulumi.Input

	ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput
	ToUserAssignedIdentityResponseOutputWithContext(context.Context) UserAssignedIdentityResponseOutput
}

type UserAssignedIdentityResponseArgs struct {
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
}

func (UserAssignedIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (i UserAssignedIdentityResponseArgs) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return i.ToUserAssignedIdentityResponseOutputWithContext(context.Background())
}

func (i UserAssignedIdentityResponseArgs) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityResponseOutput)
}





type UserAssignedIdentityResponseMapInput interface {
	pulumi.Input

	ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput
	ToUserAssignedIdentityResponseMapOutputWithContext(context.Context) UserAssignedIdentityResponseMapOutput
}

type UserAssignedIdentityResponseMap map[string]UserAssignedIdentityResponseInput

func (UserAssignedIdentityResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (i UserAssignedIdentityResponseMap) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return i.ToUserAssignedIdentityResponseMapOutputWithContext(context.Background())
}

func (i UserAssignedIdentityResponseMap) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityResponseMapOutput)
}

type UserAssignedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserAssignedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityResponse {
		return vs[0].(map[string]UserAssignedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedIdentityResponseOutput)
}

type VMSSExtension struct {
	AutoUpgradeMinorVersion  *bool       `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag           *string     `pulumi:"forceUpdateTag"`
	Name                     string      `pulumi:"name"`
	ProtectedSettings        interface{} `pulumi:"protectedSettings"`
	ProvisionAfterExtensions []string    `pulumi:"provisionAfterExtensions"`
	Publisher                string      `pulumi:"publisher"`
	Settings                 interface{} `pulumi:"settings"`
	Type                     string      `pulumi:"type"`
	TypeHandlerVersion       string      `pulumi:"typeHandlerVersion"`
}





type VMSSExtensionInput interface {
	pulumi.Input

	ToVMSSExtensionOutput() VMSSExtensionOutput
	ToVMSSExtensionOutputWithContext(context.Context) VMSSExtensionOutput
}

type VMSSExtensionArgs struct {
	AutoUpgradeMinorVersion  pulumi.BoolPtrInput     `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag           pulumi.StringPtrInput   `pulumi:"forceUpdateTag"`
	Name                     pulumi.StringInput      `pulumi:"name"`
	ProtectedSettings        pulumi.Input            `pulumi:"protectedSettings"`
	ProvisionAfterExtensions pulumi.StringArrayInput `pulumi:"provisionAfterExtensions"`
	Publisher                pulumi.StringInput      `pulumi:"publisher"`
	Settings                 pulumi.Input            `pulumi:"settings"`
	Type                     pulumi.StringInput      `pulumi:"type"`
	TypeHandlerVersion       pulumi.StringInput      `pulumi:"typeHandlerVersion"`
}

func (VMSSExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMSSExtension)(nil)).Elem()
}

func (i VMSSExtensionArgs) ToVMSSExtensionOutput() VMSSExtensionOutput {
	return i.ToVMSSExtensionOutputWithContext(context.Background())
}

func (i VMSSExtensionArgs) ToVMSSExtensionOutputWithContext(ctx context.Context) VMSSExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMSSExtensionOutput)
}





type VMSSExtensionArrayInput interface {
	pulumi.Input

	ToVMSSExtensionArrayOutput() VMSSExtensionArrayOutput
	ToVMSSExtensionArrayOutputWithContext(context.Context) VMSSExtensionArrayOutput
}

type VMSSExtensionArray []VMSSExtensionInput

func (VMSSExtensionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMSSExtension)(nil)).Elem()
}

func (i VMSSExtensionArray) ToVMSSExtensionArrayOutput() VMSSExtensionArrayOutput {
	return i.ToVMSSExtensionArrayOutputWithContext(context.Background())
}

func (i VMSSExtensionArray) ToVMSSExtensionArrayOutputWithContext(ctx context.Context) VMSSExtensionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMSSExtensionArrayOutput)
}

type VMSSExtensionOutput struct{ *pulumi.OutputState }

func (VMSSExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMSSExtension)(nil)).Elem()
}

func (o VMSSExtensionOutput) ToVMSSExtensionOutput() VMSSExtensionOutput {
	return o
}

func (o VMSSExtensionOutput) ToVMSSExtensionOutputWithContext(ctx context.Context) VMSSExtensionOutput {
	return o
}

func (o VMSSExtensionOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VMSSExtension) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o VMSSExtensionOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMSSExtension) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o VMSSExtensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtension) string { return v.Name }).(pulumi.StringOutput)
}

func (o VMSSExtensionOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v VMSSExtension) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o VMSSExtensionOutput) ProvisionAfterExtensions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VMSSExtension) []string { return v.ProvisionAfterExtensions }).(pulumi.StringArrayOutput)
}

func (o VMSSExtensionOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtension) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o VMSSExtensionOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v VMSSExtension) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o VMSSExtensionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtension) string { return v.Type }).(pulumi.StringOutput)
}

func (o VMSSExtensionOutput) TypeHandlerVersion() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtension) string { return v.TypeHandlerVersion }).(pulumi.StringOutput)
}

type VMSSExtensionArrayOutput struct{ *pulumi.OutputState }

func (VMSSExtensionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMSSExtension)(nil)).Elem()
}

func (o VMSSExtensionArrayOutput) ToVMSSExtensionArrayOutput() VMSSExtensionArrayOutput {
	return o
}

func (o VMSSExtensionArrayOutput) ToVMSSExtensionArrayOutputWithContext(ctx context.Context) VMSSExtensionArrayOutput {
	return o
}

func (o VMSSExtensionArrayOutput) Index(i pulumi.IntInput) VMSSExtensionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VMSSExtension {
		return vs[0].([]VMSSExtension)[vs[1].(int)]
	}).(VMSSExtensionOutput)
}

type VMSSExtensionResponse struct {
	AutoUpgradeMinorVersion  *bool       `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag           *string     `pulumi:"forceUpdateTag"`
	Name                     string      `pulumi:"name"`
	ProtectedSettings        interface{} `pulumi:"protectedSettings"`
	ProvisionAfterExtensions []string    `pulumi:"provisionAfterExtensions"`
	ProvisioningState        string      `pulumi:"provisioningState"`
	Publisher                string      `pulumi:"publisher"`
	Settings                 interface{} `pulumi:"settings"`
	Type                     string      `pulumi:"type"`
	TypeHandlerVersion       string      `pulumi:"typeHandlerVersion"`
}





type VMSSExtensionResponseInput interface {
	pulumi.Input

	ToVMSSExtensionResponseOutput() VMSSExtensionResponseOutput
	ToVMSSExtensionResponseOutputWithContext(context.Context) VMSSExtensionResponseOutput
}

type VMSSExtensionResponseArgs struct {
	AutoUpgradeMinorVersion  pulumi.BoolPtrInput     `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag           pulumi.StringPtrInput   `pulumi:"forceUpdateTag"`
	Name                     pulumi.StringInput      `pulumi:"name"`
	ProtectedSettings        pulumi.Input            `pulumi:"protectedSettings"`
	ProvisionAfterExtensions pulumi.StringArrayInput `pulumi:"provisionAfterExtensions"`
	ProvisioningState        pulumi.StringInput      `pulumi:"provisioningState"`
	Publisher                pulumi.StringInput      `pulumi:"publisher"`
	Settings                 pulumi.Input            `pulumi:"settings"`
	Type                     pulumi.StringInput      `pulumi:"type"`
	TypeHandlerVersion       pulumi.StringInput      `pulumi:"typeHandlerVersion"`
}

func (VMSSExtensionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMSSExtensionResponse)(nil)).Elem()
}

func (i VMSSExtensionResponseArgs) ToVMSSExtensionResponseOutput() VMSSExtensionResponseOutput {
	return i.ToVMSSExtensionResponseOutputWithContext(context.Background())
}

func (i VMSSExtensionResponseArgs) ToVMSSExtensionResponseOutputWithContext(ctx context.Context) VMSSExtensionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMSSExtensionResponseOutput)
}





type VMSSExtensionResponseArrayInput interface {
	pulumi.Input

	ToVMSSExtensionResponseArrayOutput() VMSSExtensionResponseArrayOutput
	ToVMSSExtensionResponseArrayOutputWithContext(context.Context) VMSSExtensionResponseArrayOutput
}

type VMSSExtensionResponseArray []VMSSExtensionResponseInput

func (VMSSExtensionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMSSExtensionResponse)(nil)).Elem()
}

func (i VMSSExtensionResponseArray) ToVMSSExtensionResponseArrayOutput() VMSSExtensionResponseArrayOutput {
	return i.ToVMSSExtensionResponseArrayOutputWithContext(context.Background())
}

func (i VMSSExtensionResponseArray) ToVMSSExtensionResponseArrayOutputWithContext(ctx context.Context) VMSSExtensionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMSSExtensionResponseArrayOutput)
}

type VMSSExtensionResponseOutput struct{ *pulumi.OutputState }

func (VMSSExtensionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMSSExtensionResponse)(nil)).Elem()
}

func (o VMSSExtensionResponseOutput) ToVMSSExtensionResponseOutput() VMSSExtensionResponseOutput {
	return o
}

func (o VMSSExtensionResponseOutput) ToVMSSExtensionResponseOutputWithContext(ctx context.Context) VMSSExtensionResponseOutput {
	return o
}

func (o VMSSExtensionResponseOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o VMSSExtensionResponseOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o VMSSExtensionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VMSSExtensionResponseOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o VMSSExtensionResponseOutput) ProvisionAfterExtensions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) []string { return v.ProvisionAfterExtensions }).(pulumi.StringArrayOutput)
}

func (o VMSSExtensionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VMSSExtensionResponseOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o VMSSExtensionResponseOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o VMSSExtensionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o VMSSExtensionResponseOutput) TypeHandlerVersion() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) string { return v.TypeHandlerVersion }).(pulumi.StringOutput)
}

type VMSSExtensionResponseArrayOutput struct{ *pulumi.OutputState }

func (VMSSExtensionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMSSExtensionResponse)(nil)).Elem()
}

func (o VMSSExtensionResponseArrayOutput) ToVMSSExtensionResponseArrayOutput() VMSSExtensionResponseArrayOutput {
	return o
}

func (o VMSSExtensionResponseArrayOutput) ToVMSSExtensionResponseArrayOutputWithContext(ctx context.Context) VMSSExtensionResponseArrayOutput {
	return o
}

func (o VMSSExtensionResponseArrayOutput) Index(i pulumi.IntInput) VMSSExtensionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VMSSExtensionResponse {
		return vs[0].([]VMSSExtensionResponse)[vs[1].(int)]
	}).(VMSSExtensionResponseOutput)
}

type VaultCertificate struct {
	CertificateStore string `pulumi:"certificateStore"`
	CertificateUrl   string `pulumi:"certificateUrl"`
}





type VaultCertificateInput interface {
	pulumi.Input

	ToVaultCertificateOutput() VaultCertificateOutput
	ToVaultCertificateOutputWithContext(context.Context) VaultCertificateOutput
}

type VaultCertificateArgs struct {
	CertificateStore pulumi.StringInput `pulumi:"certificateStore"`
	CertificateUrl   pulumi.StringInput `pulumi:"certificateUrl"`
}

func (VaultCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultCertificate)(nil)).Elem()
}

func (i VaultCertificateArgs) ToVaultCertificateOutput() VaultCertificateOutput {
	return i.ToVaultCertificateOutputWithContext(context.Background())
}

func (i VaultCertificateArgs) ToVaultCertificateOutputWithContext(ctx context.Context) VaultCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultCertificateOutput)
}





type VaultCertificateArrayInput interface {
	pulumi.Input

	ToVaultCertificateArrayOutput() VaultCertificateArrayOutput
	ToVaultCertificateArrayOutputWithContext(context.Context) VaultCertificateArrayOutput
}

type VaultCertificateArray []VaultCertificateInput

func (VaultCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultCertificate)(nil)).Elem()
}

func (i VaultCertificateArray) ToVaultCertificateArrayOutput() VaultCertificateArrayOutput {
	return i.ToVaultCertificateArrayOutputWithContext(context.Background())
}

func (i VaultCertificateArray) ToVaultCertificateArrayOutputWithContext(ctx context.Context) VaultCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultCertificateArrayOutput)
}

type VaultCertificateOutput struct{ *pulumi.OutputState }

func (VaultCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultCertificate)(nil)).Elem()
}

func (o VaultCertificateOutput) ToVaultCertificateOutput() VaultCertificateOutput {
	return o
}

func (o VaultCertificateOutput) ToVaultCertificateOutputWithContext(ctx context.Context) VaultCertificateOutput {
	return o
}

func (o VaultCertificateOutput) CertificateStore() pulumi.StringOutput {
	return o.ApplyT(func(v VaultCertificate) string { return v.CertificateStore }).(pulumi.StringOutput)
}

func (o VaultCertificateOutput) CertificateUrl() pulumi.StringOutput {
	return o.ApplyT(func(v VaultCertificate) string { return v.CertificateUrl }).(pulumi.StringOutput)
}

type VaultCertificateArrayOutput struct{ *pulumi.OutputState }

func (VaultCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultCertificate)(nil)).Elem()
}

func (o VaultCertificateArrayOutput) ToVaultCertificateArrayOutput() VaultCertificateArrayOutput {
	return o
}

func (o VaultCertificateArrayOutput) ToVaultCertificateArrayOutputWithContext(ctx context.Context) VaultCertificateArrayOutput {
	return o
}

func (o VaultCertificateArrayOutput) Index(i pulumi.IntInput) VaultCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VaultCertificate {
		return vs[0].([]VaultCertificate)[vs[1].(int)]
	}).(VaultCertificateOutput)
}

type VaultCertificateResponse struct {
	CertificateStore string `pulumi:"certificateStore"`
	CertificateUrl   string `pulumi:"certificateUrl"`
}





type VaultCertificateResponseInput interface {
	pulumi.Input

	ToVaultCertificateResponseOutput() VaultCertificateResponseOutput
	ToVaultCertificateResponseOutputWithContext(context.Context) VaultCertificateResponseOutput
}

type VaultCertificateResponseArgs struct {
	CertificateStore pulumi.StringInput `pulumi:"certificateStore"`
	CertificateUrl   pulumi.StringInput `pulumi:"certificateUrl"`
}

func (VaultCertificateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultCertificateResponse)(nil)).Elem()
}

func (i VaultCertificateResponseArgs) ToVaultCertificateResponseOutput() VaultCertificateResponseOutput {
	return i.ToVaultCertificateResponseOutputWithContext(context.Background())
}

func (i VaultCertificateResponseArgs) ToVaultCertificateResponseOutputWithContext(ctx context.Context) VaultCertificateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultCertificateResponseOutput)
}





type VaultCertificateResponseArrayInput interface {
	pulumi.Input

	ToVaultCertificateResponseArrayOutput() VaultCertificateResponseArrayOutput
	ToVaultCertificateResponseArrayOutputWithContext(context.Context) VaultCertificateResponseArrayOutput
}

type VaultCertificateResponseArray []VaultCertificateResponseInput

func (VaultCertificateResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultCertificateResponse)(nil)).Elem()
}

func (i VaultCertificateResponseArray) ToVaultCertificateResponseArrayOutput() VaultCertificateResponseArrayOutput {
	return i.ToVaultCertificateResponseArrayOutputWithContext(context.Background())
}

func (i VaultCertificateResponseArray) ToVaultCertificateResponseArrayOutputWithContext(ctx context.Context) VaultCertificateResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultCertificateResponseArrayOutput)
}

type VaultCertificateResponseOutput struct{ *pulumi.OutputState }

func (VaultCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultCertificateResponse)(nil)).Elem()
}

func (o VaultCertificateResponseOutput) ToVaultCertificateResponseOutput() VaultCertificateResponseOutput {
	return o
}

func (o VaultCertificateResponseOutput) ToVaultCertificateResponseOutputWithContext(ctx context.Context) VaultCertificateResponseOutput {
	return o
}

func (o VaultCertificateResponseOutput) CertificateStore() pulumi.StringOutput {
	return o.ApplyT(func(v VaultCertificateResponse) string { return v.CertificateStore }).(pulumi.StringOutput)
}

func (o VaultCertificateResponseOutput) CertificateUrl() pulumi.StringOutput {
	return o.ApplyT(func(v VaultCertificateResponse) string { return v.CertificateUrl }).(pulumi.StringOutput)
}

type VaultCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VaultCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultCertificateResponse)(nil)).Elem()
}

func (o VaultCertificateResponseArrayOutput) ToVaultCertificateResponseArrayOutput() VaultCertificateResponseArrayOutput {
	return o
}

func (o VaultCertificateResponseArrayOutput) ToVaultCertificateResponseArrayOutputWithContext(ctx context.Context) VaultCertificateResponseArrayOutput {
	return o
}

func (o VaultCertificateResponseArrayOutput) Index(i pulumi.IntInput) VaultCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VaultCertificateResponse {
		return vs[0].([]VaultCertificateResponse)[vs[1].(int)]
	}).(VaultCertificateResponseOutput)
}

type VaultSecretGroup struct {
	SourceVault       SubResource        `pulumi:"sourceVault"`
	VaultCertificates []VaultCertificate `pulumi:"vaultCertificates"`
}





type VaultSecretGroupInput interface {
	pulumi.Input

	ToVaultSecretGroupOutput() VaultSecretGroupOutput
	ToVaultSecretGroupOutputWithContext(context.Context) VaultSecretGroupOutput
}

type VaultSecretGroupArgs struct {
	SourceVault       SubResourceInput           `pulumi:"sourceVault"`
	VaultCertificates VaultCertificateArrayInput `pulumi:"vaultCertificates"`
}

func (VaultSecretGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultSecretGroup)(nil)).Elem()
}

func (i VaultSecretGroupArgs) ToVaultSecretGroupOutput() VaultSecretGroupOutput {
	return i.ToVaultSecretGroupOutputWithContext(context.Background())
}

func (i VaultSecretGroupArgs) ToVaultSecretGroupOutputWithContext(ctx context.Context) VaultSecretGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultSecretGroupOutput)
}





type VaultSecretGroupArrayInput interface {
	pulumi.Input

	ToVaultSecretGroupArrayOutput() VaultSecretGroupArrayOutput
	ToVaultSecretGroupArrayOutputWithContext(context.Context) VaultSecretGroupArrayOutput
}

type VaultSecretGroupArray []VaultSecretGroupInput

func (VaultSecretGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultSecretGroup)(nil)).Elem()
}

func (i VaultSecretGroupArray) ToVaultSecretGroupArrayOutput() VaultSecretGroupArrayOutput {
	return i.ToVaultSecretGroupArrayOutputWithContext(context.Background())
}

func (i VaultSecretGroupArray) ToVaultSecretGroupArrayOutputWithContext(ctx context.Context) VaultSecretGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultSecretGroupArrayOutput)
}

type VaultSecretGroupOutput struct{ *pulumi.OutputState }

func (VaultSecretGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultSecretGroup)(nil)).Elem()
}

func (o VaultSecretGroupOutput) ToVaultSecretGroupOutput() VaultSecretGroupOutput {
	return o
}

func (o VaultSecretGroupOutput) ToVaultSecretGroupOutputWithContext(ctx context.Context) VaultSecretGroupOutput {
	return o
}

func (o VaultSecretGroupOutput) SourceVault() SubResourceOutput {
	return o.ApplyT(func(v VaultSecretGroup) SubResource { return v.SourceVault }).(SubResourceOutput)
}

func (o VaultSecretGroupOutput) VaultCertificates() VaultCertificateArrayOutput {
	return o.ApplyT(func(v VaultSecretGroup) []VaultCertificate { return v.VaultCertificates }).(VaultCertificateArrayOutput)
}

type VaultSecretGroupArrayOutput struct{ *pulumi.OutputState }

func (VaultSecretGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultSecretGroup)(nil)).Elem()
}

func (o VaultSecretGroupArrayOutput) ToVaultSecretGroupArrayOutput() VaultSecretGroupArrayOutput {
	return o
}

func (o VaultSecretGroupArrayOutput) ToVaultSecretGroupArrayOutputWithContext(ctx context.Context) VaultSecretGroupArrayOutput {
	return o
}

func (o VaultSecretGroupArrayOutput) Index(i pulumi.IntInput) VaultSecretGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VaultSecretGroup {
		return vs[0].([]VaultSecretGroup)[vs[1].(int)]
	}).(VaultSecretGroupOutput)
}

type VaultSecretGroupResponse struct {
	SourceVault       SubResourceResponse        `pulumi:"sourceVault"`
	VaultCertificates []VaultCertificateResponse `pulumi:"vaultCertificates"`
}





type VaultSecretGroupResponseInput interface {
	pulumi.Input

	ToVaultSecretGroupResponseOutput() VaultSecretGroupResponseOutput
	ToVaultSecretGroupResponseOutputWithContext(context.Context) VaultSecretGroupResponseOutput
}

type VaultSecretGroupResponseArgs struct {
	SourceVault       SubResourceResponseInput           `pulumi:"sourceVault"`
	VaultCertificates VaultCertificateResponseArrayInput `pulumi:"vaultCertificates"`
}

func (VaultSecretGroupResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultSecretGroupResponse)(nil)).Elem()
}

func (i VaultSecretGroupResponseArgs) ToVaultSecretGroupResponseOutput() VaultSecretGroupResponseOutput {
	return i.ToVaultSecretGroupResponseOutputWithContext(context.Background())
}

func (i VaultSecretGroupResponseArgs) ToVaultSecretGroupResponseOutputWithContext(ctx context.Context) VaultSecretGroupResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultSecretGroupResponseOutput)
}





type VaultSecretGroupResponseArrayInput interface {
	pulumi.Input

	ToVaultSecretGroupResponseArrayOutput() VaultSecretGroupResponseArrayOutput
	ToVaultSecretGroupResponseArrayOutputWithContext(context.Context) VaultSecretGroupResponseArrayOutput
}

type VaultSecretGroupResponseArray []VaultSecretGroupResponseInput

func (VaultSecretGroupResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultSecretGroupResponse)(nil)).Elem()
}

func (i VaultSecretGroupResponseArray) ToVaultSecretGroupResponseArrayOutput() VaultSecretGroupResponseArrayOutput {
	return i.ToVaultSecretGroupResponseArrayOutputWithContext(context.Background())
}

func (i VaultSecretGroupResponseArray) ToVaultSecretGroupResponseArrayOutputWithContext(ctx context.Context) VaultSecretGroupResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultSecretGroupResponseArrayOutput)
}

type VaultSecretGroupResponseOutput struct{ *pulumi.OutputState }

func (VaultSecretGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultSecretGroupResponse)(nil)).Elem()
}

func (o VaultSecretGroupResponseOutput) ToVaultSecretGroupResponseOutput() VaultSecretGroupResponseOutput {
	return o
}

func (o VaultSecretGroupResponseOutput) ToVaultSecretGroupResponseOutputWithContext(ctx context.Context) VaultSecretGroupResponseOutput {
	return o
}

func (o VaultSecretGroupResponseOutput) SourceVault() SubResourceResponseOutput {
	return o.ApplyT(func(v VaultSecretGroupResponse) SubResourceResponse { return v.SourceVault }).(SubResourceResponseOutput)
}

func (o VaultSecretGroupResponseOutput) VaultCertificates() VaultCertificateResponseArrayOutput {
	return o.ApplyT(func(v VaultSecretGroupResponse) []VaultCertificateResponse { return v.VaultCertificates }).(VaultCertificateResponseArrayOutput)
}

type VaultSecretGroupResponseArrayOutput struct{ *pulumi.OutputState }

func (VaultSecretGroupResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultSecretGroupResponse)(nil)).Elem()
}

func (o VaultSecretGroupResponseArrayOutput) ToVaultSecretGroupResponseArrayOutput() VaultSecretGroupResponseArrayOutput {
	return o
}

func (o VaultSecretGroupResponseArrayOutput) ToVaultSecretGroupResponseArrayOutputWithContext(ctx context.Context) VaultSecretGroupResponseArrayOutput {
	return o
}

func (o VaultSecretGroupResponseArrayOutput) Index(i pulumi.IntInput) VaultSecretGroupResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VaultSecretGroupResponse {
		return vs[0].([]VaultSecretGroupResponse)[vs[1].(int)]
	}).(VaultSecretGroupResponseOutput)
}

type VmManagedIdentity struct {
	UserAssignedIdentities []string `pulumi:"userAssignedIdentities"`
}





type VmManagedIdentityInput interface {
	pulumi.Input

	ToVmManagedIdentityOutput() VmManagedIdentityOutput
	ToVmManagedIdentityOutputWithContext(context.Context) VmManagedIdentityOutput
}

type VmManagedIdentityArgs struct {
	UserAssignedIdentities pulumi.StringArrayInput `pulumi:"userAssignedIdentities"`
}

func (VmManagedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmManagedIdentity)(nil)).Elem()
}

func (i VmManagedIdentityArgs) ToVmManagedIdentityOutput() VmManagedIdentityOutput {
	return i.ToVmManagedIdentityOutputWithContext(context.Background())
}

func (i VmManagedIdentityArgs) ToVmManagedIdentityOutputWithContext(ctx context.Context) VmManagedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmManagedIdentityOutput)
}

func (i VmManagedIdentityArgs) ToVmManagedIdentityPtrOutput() VmManagedIdentityPtrOutput {
	return i.ToVmManagedIdentityPtrOutputWithContext(context.Background())
}

func (i VmManagedIdentityArgs) ToVmManagedIdentityPtrOutputWithContext(ctx context.Context) VmManagedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmManagedIdentityOutput).ToVmManagedIdentityPtrOutputWithContext(ctx)
}









type VmManagedIdentityPtrInput interface {
	pulumi.Input

	ToVmManagedIdentityPtrOutput() VmManagedIdentityPtrOutput
	ToVmManagedIdentityPtrOutputWithContext(context.Context) VmManagedIdentityPtrOutput
}

type vmManagedIdentityPtrType VmManagedIdentityArgs

func VmManagedIdentityPtr(v *VmManagedIdentityArgs) VmManagedIdentityPtrInput {
	return (*vmManagedIdentityPtrType)(v)
}

func (*vmManagedIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VmManagedIdentity)(nil)).Elem()
}

func (i *vmManagedIdentityPtrType) ToVmManagedIdentityPtrOutput() VmManagedIdentityPtrOutput {
	return i.ToVmManagedIdentityPtrOutputWithContext(context.Background())
}

func (i *vmManagedIdentityPtrType) ToVmManagedIdentityPtrOutputWithContext(ctx context.Context) VmManagedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmManagedIdentityPtrOutput)
}

type VmManagedIdentityOutput struct{ *pulumi.OutputState }

func (VmManagedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmManagedIdentity)(nil)).Elem()
}

func (o VmManagedIdentityOutput) ToVmManagedIdentityOutput() VmManagedIdentityOutput {
	return o
}

func (o VmManagedIdentityOutput) ToVmManagedIdentityOutputWithContext(ctx context.Context) VmManagedIdentityOutput {
	return o
}

func (o VmManagedIdentityOutput) ToVmManagedIdentityPtrOutput() VmManagedIdentityPtrOutput {
	return o.ToVmManagedIdentityPtrOutputWithContext(context.Background())
}

func (o VmManagedIdentityOutput) ToVmManagedIdentityPtrOutputWithContext(ctx context.Context) VmManagedIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VmManagedIdentity) *VmManagedIdentity {
		return &v
	}).(VmManagedIdentityPtrOutput)
}

func (o VmManagedIdentityOutput) UserAssignedIdentities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VmManagedIdentity) []string { return v.UserAssignedIdentities }).(pulumi.StringArrayOutput)
}

type VmManagedIdentityPtrOutput struct{ *pulumi.OutputState }

func (VmManagedIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VmManagedIdentity)(nil)).Elem()
}

func (o VmManagedIdentityPtrOutput) ToVmManagedIdentityPtrOutput() VmManagedIdentityPtrOutput {
	return o
}

func (o VmManagedIdentityPtrOutput) ToVmManagedIdentityPtrOutputWithContext(ctx context.Context) VmManagedIdentityPtrOutput {
	return o
}

func (o VmManagedIdentityPtrOutput) Elem() VmManagedIdentityOutput {
	return o.ApplyT(func(v *VmManagedIdentity) VmManagedIdentity {
		if v != nil {
			return *v
		}
		var ret VmManagedIdentity
		return ret
	}).(VmManagedIdentityOutput)
}

func (o VmManagedIdentityPtrOutput) UserAssignedIdentities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VmManagedIdentity) []string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.StringArrayOutput)
}

type VmManagedIdentityResponse struct {
	UserAssignedIdentities []string `pulumi:"userAssignedIdentities"`
}





type VmManagedIdentityResponseInput interface {
	pulumi.Input

	ToVmManagedIdentityResponseOutput() VmManagedIdentityResponseOutput
	ToVmManagedIdentityResponseOutputWithContext(context.Context) VmManagedIdentityResponseOutput
}

type VmManagedIdentityResponseArgs struct {
	UserAssignedIdentities pulumi.StringArrayInput `pulumi:"userAssignedIdentities"`
}

func (VmManagedIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmManagedIdentityResponse)(nil)).Elem()
}

func (i VmManagedIdentityResponseArgs) ToVmManagedIdentityResponseOutput() VmManagedIdentityResponseOutput {
	return i.ToVmManagedIdentityResponseOutputWithContext(context.Background())
}

func (i VmManagedIdentityResponseArgs) ToVmManagedIdentityResponseOutputWithContext(ctx context.Context) VmManagedIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmManagedIdentityResponseOutput)
}

func (i VmManagedIdentityResponseArgs) ToVmManagedIdentityResponsePtrOutput() VmManagedIdentityResponsePtrOutput {
	return i.ToVmManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (i VmManagedIdentityResponseArgs) ToVmManagedIdentityResponsePtrOutputWithContext(ctx context.Context) VmManagedIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmManagedIdentityResponseOutput).ToVmManagedIdentityResponsePtrOutputWithContext(ctx)
}









type VmManagedIdentityResponsePtrInput interface {
	pulumi.Input

	ToVmManagedIdentityResponsePtrOutput() VmManagedIdentityResponsePtrOutput
	ToVmManagedIdentityResponsePtrOutputWithContext(context.Context) VmManagedIdentityResponsePtrOutput
}

type vmManagedIdentityResponsePtrType VmManagedIdentityResponseArgs

func VmManagedIdentityResponsePtr(v *VmManagedIdentityResponseArgs) VmManagedIdentityResponsePtrInput {
	return (*vmManagedIdentityResponsePtrType)(v)
}

func (*vmManagedIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VmManagedIdentityResponse)(nil)).Elem()
}

func (i *vmManagedIdentityResponsePtrType) ToVmManagedIdentityResponsePtrOutput() VmManagedIdentityResponsePtrOutput {
	return i.ToVmManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *vmManagedIdentityResponsePtrType) ToVmManagedIdentityResponsePtrOutputWithContext(ctx context.Context) VmManagedIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmManagedIdentityResponsePtrOutput)
}

type VmManagedIdentityResponseOutput struct{ *pulumi.OutputState }

func (VmManagedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmManagedIdentityResponse)(nil)).Elem()
}

func (o VmManagedIdentityResponseOutput) ToVmManagedIdentityResponseOutput() VmManagedIdentityResponseOutput {
	return o
}

func (o VmManagedIdentityResponseOutput) ToVmManagedIdentityResponseOutputWithContext(ctx context.Context) VmManagedIdentityResponseOutput {
	return o
}

func (o VmManagedIdentityResponseOutput) ToVmManagedIdentityResponsePtrOutput() VmManagedIdentityResponsePtrOutput {
	return o.ToVmManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (o VmManagedIdentityResponseOutput) ToVmManagedIdentityResponsePtrOutputWithContext(ctx context.Context) VmManagedIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VmManagedIdentityResponse) *VmManagedIdentityResponse {
		return &v
	}).(VmManagedIdentityResponsePtrOutput)
}

func (o VmManagedIdentityResponseOutput) UserAssignedIdentities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VmManagedIdentityResponse) []string { return v.UserAssignedIdentities }).(pulumi.StringArrayOutput)
}

type VmManagedIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (VmManagedIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VmManagedIdentityResponse)(nil)).Elem()
}

func (o VmManagedIdentityResponsePtrOutput) ToVmManagedIdentityResponsePtrOutput() VmManagedIdentityResponsePtrOutput {
	return o
}

func (o VmManagedIdentityResponsePtrOutput) ToVmManagedIdentityResponsePtrOutputWithContext(ctx context.Context) VmManagedIdentityResponsePtrOutput {
	return o
}

func (o VmManagedIdentityResponsePtrOutput) Elem() VmManagedIdentityResponseOutput {
	return o.ApplyT(func(v *VmManagedIdentityResponse) VmManagedIdentityResponse {
		if v != nil {
			return *v
		}
		var ret VmManagedIdentityResponse
		return ret
	}).(VmManagedIdentityResponseOutput)
}

func (o VmManagedIdentityResponsePtrOutput) UserAssignedIdentities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VmManagedIdentityResponse) []string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AddRemoveIncrementalNamedPartitionScalingMechanismInput)(nil)).Elem(), AddRemoveIncrementalNamedPartitionScalingMechanismArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddRemoveIncrementalNamedPartitionScalingMechanismResponseInput)(nil)).Elem(), AddRemoveIncrementalNamedPartitionScalingMechanismResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationHealthPolicyInput)(nil)).Elem(), ApplicationHealthPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationHealthPolicyPtrInput)(nil)).Elem(), ApplicationHealthPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationHealthPolicyResponseInput)(nil)).Elem(), ApplicationHealthPolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationHealthPolicyResponsePtrInput)(nil)).Elem(), ApplicationHealthPolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationTypeVersionsCleanupPolicyInput)(nil)).Elem(), ApplicationTypeVersionsCleanupPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationTypeVersionsCleanupPolicyPtrInput)(nil)).Elem(), ApplicationTypeVersionsCleanupPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationTypeVersionsCleanupPolicyResponseInput)(nil)).Elem(), ApplicationTypeVersionsCleanupPolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationTypeVersionsCleanupPolicyResponsePtrInput)(nil)).Elem(), ApplicationTypeVersionsCleanupPolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationUpgradePolicyInput)(nil)).Elem(), ApplicationUpgradePolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationUpgradePolicyPtrInput)(nil)).Elem(), ApplicationUpgradePolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationUpgradePolicyResponseInput)(nil)).Elem(), ApplicationUpgradePolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationUpgradePolicyResponsePtrInput)(nil)).Elem(), ApplicationUpgradePolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationUserAssignedIdentityInput)(nil)).Elem(), ApplicationUserAssignedIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationUserAssignedIdentityArrayInput)(nil)).Elem(), ApplicationUserAssignedIdentityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationUserAssignedIdentityResponseInput)(nil)).Elem(), ApplicationUserAssignedIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationUserAssignedIdentityResponseArrayInput)(nil)).Elem(), ApplicationUserAssignedIdentityResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AveragePartitionLoadScalingTriggerInput)(nil)).Elem(), AveragePartitionLoadScalingTriggerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AveragePartitionLoadScalingTriggerResponseInput)(nil)).Elem(), AveragePartitionLoadScalingTriggerResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AverageServiceLoadScalingTriggerInput)(nil)).Elem(), AverageServiceLoadScalingTriggerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AverageServiceLoadScalingTriggerResponseInput)(nil)).Elem(), AverageServiceLoadScalingTriggerResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureActiveDirectoryInput)(nil)).Elem(), AzureActiveDirectoryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureActiveDirectoryPtrInput)(nil)).Elem(), AzureActiveDirectoryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureActiveDirectoryResponseInput)(nil)).Elem(), AzureActiveDirectoryResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureActiveDirectoryResponsePtrInput)(nil)).Elem(), AzureActiveDirectoryResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientCertificateInput)(nil)).Elem(), ClientCertificateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientCertificateArrayInput)(nil)).Elem(), ClientCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientCertificateResponseInput)(nil)).Elem(), ClientCertificateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientCertificateResponseArrayInput)(nil)).Elem(), ClientCertificateResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointRangeDescriptionInput)(nil)).Elem(), EndpointRangeDescriptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointRangeDescriptionPtrInput)(nil)).Elem(), EndpointRangeDescriptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointRangeDescriptionResponseInput)(nil)).Elem(), EndpointRangeDescriptionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointRangeDescriptionResponsePtrInput)(nil)).Elem(), EndpointRangeDescriptionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancingRuleInput)(nil)).Elem(), LoadBalancingRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancingRuleArrayInput)(nil)).Elem(), LoadBalancingRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancingRuleResponseInput)(nil)).Elem(), LoadBalancingRuleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancingRuleResponseArrayInput)(nil)).Elem(), LoadBalancingRuleResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedIdentityInput)(nil)).Elem(), ManagedIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedIdentityPtrInput)(nil)).Elem(), ManagedIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedIdentityResponseInput)(nil)).Elem(), ManagedIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedIdentityResponsePtrInput)(nil)).Elem(), ManagedIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamedPartitionSchemeInput)(nil)).Elem(), NamedPartitionSchemeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamedPartitionSchemeResponseInput)(nil)).Elem(), NamedPartitionSchemeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkSecurityRuleInput)(nil)).Elem(), NetworkSecurityRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkSecurityRuleArrayInput)(nil)).Elem(), NetworkSecurityRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkSecurityRuleResponseInput)(nil)).Elem(), NetworkSecurityRuleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkSecurityRuleResponseArrayInput)(nil)).Elem(), NetworkSecurityRuleResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartitionInstanceCountScaleMechanismInput)(nil)).Elem(), PartitionInstanceCountScaleMechanismArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartitionInstanceCountScaleMechanismResponseInput)(nil)).Elem(), PartitionInstanceCountScaleMechanismResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RollingUpgradeMonitoringPolicyInput)(nil)).Elem(), RollingUpgradeMonitoringPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RollingUpgradeMonitoringPolicyPtrInput)(nil)).Elem(), RollingUpgradeMonitoringPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RollingUpgradeMonitoringPolicyResponseInput)(nil)).Elem(), RollingUpgradeMonitoringPolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RollingUpgradeMonitoringPolicyResponsePtrInput)(nil)).Elem(), RollingUpgradeMonitoringPolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingPolicyInput)(nil)).Elem(), ScalingPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingPolicyArrayInput)(nil)).Elem(), ScalingPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingPolicyResponseInput)(nil)).Elem(), ScalingPolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingPolicyResponseArrayInput)(nil)).Elem(), ScalingPolicyResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceCorrelationInput)(nil)).Elem(), ServiceCorrelationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceCorrelationArrayInput)(nil)).Elem(), ServiceCorrelationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceCorrelationResponseInput)(nil)).Elem(), ServiceCorrelationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceCorrelationResponseArrayInput)(nil)).Elem(), ServiceCorrelationResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceLoadMetricInput)(nil)).Elem(), ServiceLoadMetricArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceLoadMetricArrayInput)(nil)).Elem(), ServiceLoadMetricArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceLoadMetricResponseInput)(nil)).Elem(), ServiceLoadMetricResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceLoadMetricResponseArrayInput)(nil)).Elem(), ServiceLoadMetricResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePlacementInvalidDomainPolicyInput)(nil)).Elem(), ServicePlacementInvalidDomainPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePlacementInvalidDomainPolicyResponseInput)(nil)).Elem(), ServicePlacementInvalidDomainPolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePlacementNonPartiallyPlaceServicePolicyInput)(nil)).Elem(), ServicePlacementNonPartiallyPlaceServicePolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePlacementNonPartiallyPlaceServicePolicyResponseInput)(nil)).Elem(), ServicePlacementNonPartiallyPlaceServicePolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePlacementPreferPrimaryDomainPolicyInput)(nil)).Elem(), ServicePlacementPreferPrimaryDomainPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePlacementPreferPrimaryDomainPolicyResponseInput)(nil)).Elem(), ServicePlacementPreferPrimaryDomainPolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePlacementRequireDomainDistributionPolicyInput)(nil)).Elem(), ServicePlacementRequireDomainDistributionPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePlacementRequireDomainDistributionPolicyResponseInput)(nil)).Elem(), ServicePlacementRequireDomainDistributionPolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePlacementRequiredDomainPolicyInput)(nil)).Elem(), ServicePlacementRequiredDomainPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePlacementRequiredDomainPolicyResponseInput)(nil)).Elem(), ServicePlacementRequiredDomainPolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceTypeHealthPolicyInput)(nil)).Elem(), ServiceTypeHealthPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceTypeHealthPolicyPtrInput)(nil)).Elem(), ServiceTypeHealthPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceTypeHealthPolicyMapInput)(nil)).Elem(), ServiceTypeHealthPolicyMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceTypeHealthPolicyResponseInput)(nil)).Elem(), ServiceTypeHealthPolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceTypeHealthPolicyResponsePtrInput)(nil)).Elem(), ServiceTypeHealthPolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceTypeHealthPolicyResponseMapInput)(nil)).Elem(), ServiceTypeHealthPolicyResponseMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingsParameterDescriptionInput)(nil)).Elem(), SettingsParameterDescriptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingsParameterDescriptionArrayInput)(nil)).Elem(), SettingsParameterDescriptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingsParameterDescriptionResponseInput)(nil)).Elem(), SettingsParameterDescriptionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingsParameterDescriptionResponseArrayInput)(nil)).Elem(), SettingsParameterDescriptionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingsSectionDescriptionInput)(nil)).Elem(), SettingsSectionDescriptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingsSectionDescriptionArrayInput)(nil)).Elem(), SettingsSectionDescriptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingsSectionDescriptionResponseInput)(nil)).Elem(), SettingsSectionDescriptionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SettingsSectionDescriptionResponseArrayInput)(nil)).Elem(), SettingsSectionDescriptionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SingletonPartitionSchemeInput)(nil)).Elem(), SingletonPartitionSchemeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SingletonPartitionSchemeResponseInput)(nil)).Elem(), SingletonPartitionSchemeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuPtrInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponseInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponsePtrInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StatefulServicePropertiesInput)(nil)).Elem(), StatefulServicePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StatefulServicePropertiesResponseInput)(nil)).Elem(), StatefulServicePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StatelessServicePropertiesInput)(nil)).Elem(), StatelessServicePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StatelessServicePropertiesResponseInput)(nil)).Elem(), StatelessServicePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubResourceInput)(nil)).Elem(), SubResourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubResourceResponseInput)(nil)).Elem(), SubResourceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UniformInt64RangePartitionSchemeInput)(nil)).Elem(), UniformInt64RangePartitionSchemeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UniformInt64RangePartitionSchemeResponseInput)(nil)).Elem(), UniformInt64RangePartitionSchemeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAssignedIdentityResponseInput)(nil)).Elem(), UserAssignedIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAssignedIdentityResponseMapInput)(nil)).Elem(), UserAssignedIdentityResponseMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMSSExtensionInput)(nil)).Elem(), VMSSExtensionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMSSExtensionArrayInput)(nil)).Elem(), VMSSExtensionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMSSExtensionResponseInput)(nil)).Elem(), VMSSExtensionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMSSExtensionResponseArrayInput)(nil)).Elem(), VMSSExtensionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VaultCertificateInput)(nil)).Elem(), VaultCertificateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VaultCertificateArrayInput)(nil)).Elem(), VaultCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VaultCertificateResponseInput)(nil)).Elem(), VaultCertificateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VaultCertificateResponseArrayInput)(nil)).Elem(), VaultCertificateResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VaultSecretGroupInput)(nil)).Elem(), VaultSecretGroupArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VaultSecretGroupArrayInput)(nil)).Elem(), VaultSecretGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VaultSecretGroupResponseInput)(nil)).Elem(), VaultSecretGroupResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VaultSecretGroupResponseArrayInput)(nil)).Elem(), VaultSecretGroupResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmManagedIdentityInput)(nil)).Elem(), VmManagedIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmManagedIdentityPtrInput)(nil)).Elem(), VmManagedIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmManagedIdentityResponseInput)(nil)).Elem(), VmManagedIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmManagedIdentityResponsePtrInput)(nil)).Elem(), VmManagedIdentityResponseArgs{})
	pulumi.RegisterOutputType(AddRemoveIncrementalNamedPartitionScalingMechanismOutput{})
	pulumi.RegisterOutputType(AddRemoveIncrementalNamedPartitionScalingMechanismResponseOutput{})
	pulumi.RegisterOutputType(ApplicationHealthPolicyOutput{})
	pulumi.RegisterOutputType(ApplicationHealthPolicyPtrOutput{})
	pulumi.RegisterOutputType(ApplicationHealthPolicyResponseOutput{})
	pulumi.RegisterOutputType(ApplicationHealthPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationTypeVersionsCleanupPolicyOutput{})
	pulumi.RegisterOutputType(ApplicationTypeVersionsCleanupPolicyPtrOutput{})
	pulumi.RegisterOutputType(ApplicationTypeVersionsCleanupPolicyResponseOutput{})
	pulumi.RegisterOutputType(ApplicationTypeVersionsCleanupPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationUpgradePolicyOutput{})
	pulumi.RegisterOutputType(ApplicationUpgradePolicyPtrOutput{})
	pulumi.RegisterOutputType(ApplicationUpgradePolicyResponseOutput{})
	pulumi.RegisterOutputType(ApplicationUpgradePolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationUserAssignedIdentityOutput{})
	pulumi.RegisterOutputType(ApplicationUserAssignedIdentityArrayOutput{})
	pulumi.RegisterOutputType(ApplicationUserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(ApplicationUserAssignedIdentityResponseArrayOutput{})
	pulumi.RegisterOutputType(AveragePartitionLoadScalingTriggerOutput{})
	pulumi.RegisterOutputType(AveragePartitionLoadScalingTriggerResponseOutput{})
	pulumi.RegisterOutputType(AverageServiceLoadScalingTriggerOutput{})
	pulumi.RegisterOutputType(AverageServiceLoadScalingTriggerResponseOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryPtrOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryResponseOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryResponsePtrOutput{})
	pulumi.RegisterOutputType(ClientCertificateOutput{})
	pulumi.RegisterOutputType(ClientCertificateArrayOutput{})
	pulumi.RegisterOutputType(ClientCertificateResponseOutput{})
	pulumi.RegisterOutputType(ClientCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(EndpointRangeDescriptionOutput{})
	pulumi.RegisterOutputType(EndpointRangeDescriptionPtrOutput{})
	pulumi.RegisterOutputType(EndpointRangeDescriptionResponseOutput{})
	pulumi.RegisterOutputType(EndpointRangeDescriptionResponsePtrOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleResponseOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedIdentityOutput{})
	pulumi.RegisterOutputType(ManagedIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(NamedPartitionSchemeOutput{})
	pulumi.RegisterOutputType(NamedPartitionSchemeResponseOutput{})
	pulumi.RegisterOutputType(NetworkSecurityRuleOutput{})
	pulumi.RegisterOutputType(NetworkSecurityRuleArrayOutput{})
	pulumi.RegisterOutputType(NetworkSecurityRuleResponseOutput{})
	pulumi.RegisterOutputType(NetworkSecurityRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(PartitionInstanceCountScaleMechanismOutput{})
	pulumi.RegisterOutputType(PartitionInstanceCountScaleMechanismResponseOutput{})
	pulumi.RegisterOutputType(RollingUpgradeMonitoringPolicyOutput{})
	pulumi.RegisterOutputType(RollingUpgradeMonitoringPolicyPtrOutput{})
	pulumi.RegisterOutputType(RollingUpgradeMonitoringPolicyResponseOutput{})
	pulumi.RegisterOutputType(RollingUpgradeMonitoringPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ScalingPolicyOutput{})
	pulumi.RegisterOutputType(ScalingPolicyArrayOutput{})
	pulumi.RegisterOutputType(ScalingPolicyResponseOutput{})
	pulumi.RegisterOutputType(ScalingPolicyResponseArrayOutput{})
	pulumi.RegisterOutputType(ServiceCorrelationOutput{})
	pulumi.RegisterOutputType(ServiceCorrelationArrayOutput{})
	pulumi.RegisterOutputType(ServiceCorrelationResponseOutput{})
	pulumi.RegisterOutputType(ServiceCorrelationResponseArrayOutput{})
	pulumi.RegisterOutputType(ServiceLoadMetricOutput{})
	pulumi.RegisterOutputType(ServiceLoadMetricArrayOutput{})
	pulumi.RegisterOutputType(ServiceLoadMetricResponseOutput{})
	pulumi.RegisterOutputType(ServiceLoadMetricResponseArrayOutput{})
	pulumi.RegisterOutputType(ServicePlacementInvalidDomainPolicyOutput{})
	pulumi.RegisterOutputType(ServicePlacementInvalidDomainPolicyResponseOutput{})
	pulumi.RegisterOutputType(ServicePlacementNonPartiallyPlaceServicePolicyOutput{})
	pulumi.RegisterOutputType(ServicePlacementNonPartiallyPlaceServicePolicyResponseOutput{})
	pulumi.RegisterOutputType(ServicePlacementPreferPrimaryDomainPolicyOutput{})
	pulumi.RegisterOutputType(ServicePlacementPreferPrimaryDomainPolicyResponseOutput{})
	pulumi.RegisterOutputType(ServicePlacementRequireDomainDistributionPolicyOutput{})
	pulumi.RegisterOutputType(ServicePlacementRequireDomainDistributionPolicyResponseOutput{})
	pulumi.RegisterOutputType(ServicePlacementRequiredDomainPolicyOutput{})
	pulumi.RegisterOutputType(ServicePlacementRequiredDomainPolicyResponseOutput{})
	pulumi.RegisterOutputType(ServiceTypeHealthPolicyOutput{})
	pulumi.RegisterOutputType(ServiceTypeHealthPolicyPtrOutput{})
	pulumi.RegisterOutputType(ServiceTypeHealthPolicyMapOutput{})
	pulumi.RegisterOutputType(ServiceTypeHealthPolicyResponseOutput{})
	pulumi.RegisterOutputType(ServiceTypeHealthPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceTypeHealthPolicyResponseMapOutput{})
	pulumi.RegisterOutputType(SettingsParameterDescriptionOutput{})
	pulumi.RegisterOutputType(SettingsParameterDescriptionArrayOutput{})
	pulumi.RegisterOutputType(SettingsParameterDescriptionResponseOutput{})
	pulumi.RegisterOutputType(SettingsParameterDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(SettingsSectionDescriptionOutput{})
	pulumi.RegisterOutputType(SettingsSectionDescriptionArrayOutput{})
	pulumi.RegisterOutputType(SettingsSectionDescriptionResponseOutput{})
	pulumi.RegisterOutputType(SettingsSectionDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(SingletonPartitionSchemeOutput{})
	pulumi.RegisterOutputType(SingletonPartitionSchemeResponseOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(StatefulServicePropertiesOutput{})
	pulumi.RegisterOutputType(StatefulServicePropertiesResponseOutput{})
	pulumi.RegisterOutputType(StatelessServicePropertiesOutput{})
	pulumi.RegisterOutputType(StatelessServicePropertiesResponseOutput{})
	pulumi.RegisterOutputType(SubResourceOutput{})
	pulumi.RegisterOutputType(SubResourceResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(UniformInt64RangePartitionSchemeOutput{})
	pulumi.RegisterOutputType(UniformInt64RangePartitionSchemeResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(VMSSExtensionOutput{})
	pulumi.RegisterOutputType(VMSSExtensionArrayOutput{})
	pulumi.RegisterOutputType(VMSSExtensionResponseOutput{})
	pulumi.RegisterOutputType(VMSSExtensionResponseArrayOutput{})
	pulumi.RegisterOutputType(VaultCertificateOutput{})
	pulumi.RegisterOutputType(VaultCertificateArrayOutput{})
	pulumi.RegisterOutputType(VaultCertificateResponseOutput{})
	pulumi.RegisterOutputType(VaultCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupArrayOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupResponseOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupResponseArrayOutput{})
	pulumi.RegisterOutputType(VmManagedIdentityOutput{})
	pulumi.RegisterOutputType(VmManagedIdentityPtrOutput{})
	pulumi.RegisterOutputType(VmManagedIdentityResponseOutput{})
	pulumi.RegisterOutputType(VmManagedIdentityResponsePtrOutput{})
}
