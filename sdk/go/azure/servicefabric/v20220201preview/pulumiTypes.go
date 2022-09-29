


package v20220201preview

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

type AddRemoveIncrementalNamedPartitionScalingMechanismResponse struct {
	Kind              string `pulumi:"kind"`
	MaxPartitionCount int    `pulumi:"maxPartitionCount"`
	MinPartitionCount int    `pulumi:"minPartitionCount"`
	ScaleIncrement    int    `pulumi:"scaleIncrement"`
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

type AveragePartitionLoadScalingTriggerResponse struct {
	Kind               string  `pulumi:"kind"`
	LowerLoadThreshold float64 `pulumi:"lowerLoadThreshold"`
	MetricName         string  `pulumi:"metricName"`
	ScaleInterval      string  `pulumi:"scaleInterval"`
	UpperLoadThreshold float64 `pulumi:"upperLoadThreshold"`
}

type AverageServiceLoadScalingTrigger struct {
	Kind               string  `pulumi:"kind"`
	LowerLoadThreshold float64 `pulumi:"lowerLoadThreshold"`
	MetricName         string  `pulumi:"metricName"`
	ScaleInterval      string  `pulumi:"scaleInterval"`
	UpperLoadThreshold float64 `pulumi:"upperLoadThreshold"`
	UseOnlyPrimaryLoad bool    `pulumi:"useOnlyPrimaryLoad"`
}

type AverageServiceLoadScalingTriggerResponse struct {
	Kind               string  `pulumi:"kind"`
	LowerLoadThreshold float64 `pulumi:"lowerLoadThreshold"`
	MetricName         string  `pulumi:"metricName"`
	ScaleInterval      string  `pulumi:"scaleInterval"`
	UpperLoadThreshold float64 `pulumi:"upperLoadThreshold"`
	UseOnlyPrimaryLoad bool    `pulumi:"useOnlyPrimaryLoad"`
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

type FrontendConfiguration struct {
	IpAddressType                    *string `pulumi:"ipAddressType"`
	LoadBalancerBackendAddressPoolId *string `pulumi:"loadBalancerBackendAddressPoolId"`
	LoadBalancerInboundNatPoolId     *string `pulumi:"loadBalancerInboundNatPoolId"`
}





type FrontendConfigurationInput interface {
	pulumi.Input

	ToFrontendConfigurationOutput() FrontendConfigurationOutput
	ToFrontendConfigurationOutputWithContext(context.Context) FrontendConfigurationOutput
}

type FrontendConfigurationArgs struct {
	IpAddressType                    pulumi.StringPtrInput `pulumi:"ipAddressType"`
	LoadBalancerBackendAddressPoolId pulumi.StringPtrInput `pulumi:"loadBalancerBackendAddressPoolId"`
	LoadBalancerInboundNatPoolId     pulumi.StringPtrInput `pulumi:"loadBalancerInboundNatPoolId"`
}

func (FrontendConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendConfiguration)(nil)).Elem()
}

func (i FrontendConfigurationArgs) ToFrontendConfigurationOutput() FrontendConfigurationOutput {
	return i.ToFrontendConfigurationOutputWithContext(context.Background())
}

func (i FrontendConfigurationArgs) ToFrontendConfigurationOutputWithContext(ctx context.Context) FrontendConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendConfigurationOutput)
}





type FrontendConfigurationArrayInput interface {
	pulumi.Input

	ToFrontendConfigurationArrayOutput() FrontendConfigurationArrayOutput
	ToFrontendConfigurationArrayOutputWithContext(context.Context) FrontendConfigurationArrayOutput
}

type FrontendConfigurationArray []FrontendConfigurationInput

func (FrontendConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontendConfiguration)(nil)).Elem()
}

func (i FrontendConfigurationArray) ToFrontendConfigurationArrayOutput() FrontendConfigurationArrayOutput {
	return i.ToFrontendConfigurationArrayOutputWithContext(context.Background())
}

func (i FrontendConfigurationArray) ToFrontendConfigurationArrayOutputWithContext(ctx context.Context) FrontendConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendConfigurationArrayOutput)
}

type FrontendConfigurationOutput struct{ *pulumi.OutputState }

func (FrontendConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendConfiguration)(nil)).Elem()
}

func (o FrontendConfigurationOutput) ToFrontendConfigurationOutput() FrontendConfigurationOutput {
	return o
}

func (o FrontendConfigurationOutput) ToFrontendConfigurationOutputWithContext(ctx context.Context) FrontendConfigurationOutput {
	return o
}

func (o FrontendConfigurationOutput) IpAddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendConfiguration) *string { return v.IpAddressType }).(pulumi.StringPtrOutput)
}

func (o FrontendConfigurationOutput) LoadBalancerBackendAddressPoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendConfiguration) *string { return v.LoadBalancerBackendAddressPoolId }).(pulumi.StringPtrOutput)
}

func (o FrontendConfigurationOutput) LoadBalancerInboundNatPoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendConfiguration) *string { return v.LoadBalancerInboundNatPoolId }).(pulumi.StringPtrOutput)
}

type FrontendConfigurationArrayOutput struct{ *pulumi.OutputState }

func (FrontendConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontendConfiguration)(nil)).Elem()
}

func (o FrontendConfigurationArrayOutput) ToFrontendConfigurationArrayOutput() FrontendConfigurationArrayOutput {
	return o
}

func (o FrontendConfigurationArrayOutput) ToFrontendConfigurationArrayOutputWithContext(ctx context.Context) FrontendConfigurationArrayOutput {
	return o
}

func (o FrontendConfigurationArrayOutput) Index(i pulumi.IntInput) FrontendConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontendConfiguration {
		return vs[0].([]FrontendConfiguration)[vs[1].(int)]
	}).(FrontendConfigurationOutput)
}

type FrontendConfigurationResponse struct {
	IpAddressType                    *string `pulumi:"ipAddressType"`
	LoadBalancerBackendAddressPoolId *string `pulumi:"loadBalancerBackendAddressPoolId"`
	LoadBalancerInboundNatPoolId     *string `pulumi:"loadBalancerInboundNatPoolId"`
}

type FrontendConfigurationResponseOutput struct{ *pulumi.OutputState }

func (FrontendConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendConfigurationResponse)(nil)).Elem()
}

func (o FrontendConfigurationResponseOutput) ToFrontendConfigurationResponseOutput() FrontendConfigurationResponseOutput {
	return o
}

func (o FrontendConfigurationResponseOutput) ToFrontendConfigurationResponseOutputWithContext(ctx context.Context) FrontendConfigurationResponseOutput {
	return o
}

func (o FrontendConfigurationResponseOutput) IpAddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendConfigurationResponse) *string { return v.IpAddressType }).(pulumi.StringPtrOutput)
}

func (o FrontendConfigurationResponseOutput) LoadBalancerBackendAddressPoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendConfigurationResponse) *string { return v.LoadBalancerBackendAddressPoolId }).(pulumi.StringPtrOutput)
}

func (o FrontendConfigurationResponseOutput) LoadBalancerInboundNatPoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendConfigurationResponse) *string { return v.LoadBalancerInboundNatPoolId }).(pulumi.StringPtrOutput)
}

type FrontendConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (FrontendConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontendConfigurationResponse)(nil)).Elem()
}

func (o FrontendConfigurationResponseArrayOutput) ToFrontendConfigurationResponseArrayOutput() FrontendConfigurationResponseArrayOutput {
	return o
}

func (o FrontendConfigurationResponseArrayOutput) ToFrontendConfigurationResponseArrayOutputWithContext(ctx context.Context) FrontendConfigurationResponseArrayOutput {
	return o
}

func (o FrontendConfigurationResponseArrayOutput) Index(i pulumi.IntInput) FrontendConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontendConfigurationResponse {
		return vs[0].([]FrontendConfigurationResponse)[vs[1].(int)]
	}).(FrontendConfigurationResponseOutput)
}

type IPTag struct {
	IpTagType string `pulumi:"ipTagType"`
	Tag       string `pulumi:"tag"`
}





type IPTagInput interface {
	pulumi.Input

	ToIPTagOutput() IPTagOutput
	ToIPTagOutputWithContext(context.Context) IPTagOutput
}

type IPTagArgs struct {
	IpTagType pulumi.StringInput `pulumi:"ipTagType"`
	Tag       pulumi.StringInput `pulumi:"tag"`
}

func (IPTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPTag)(nil)).Elem()
}

func (i IPTagArgs) ToIPTagOutput() IPTagOutput {
	return i.ToIPTagOutputWithContext(context.Background())
}

func (i IPTagArgs) ToIPTagOutputWithContext(ctx context.Context) IPTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPTagOutput)
}





type IPTagArrayInput interface {
	pulumi.Input

	ToIPTagArrayOutput() IPTagArrayOutput
	ToIPTagArrayOutputWithContext(context.Context) IPTagArrayOutput
}

type IPTagArray []IPTagInput

func (IPTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPTag)(nil)).Elem()
}

func (i IPTagArray) ToIPTagArrayOutput() IPTagArrayOutput {
	return i.ToIPTagArrayOutputWithContext(context.Background())
}

func (i IPTagArray) ToIPTagArrayOutputWithContext(ctx context.Context) IPTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPTagArrayOutput)
}

type IPTagOutput struct{ *pulumi.OutputState }

func (IPTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPTag)(nil)).Elem()
}

func (o IPTagOutput) ToIPTagOutput() IPTagOutput {
	return o
}

func (o IPTagOutput) ToIPTagOutputWithContext(ctx context.Context) IPTagOutput {
	return o
}

func (o IPTagOutput) IpTagType() pulumi.StringOutput {
	return o.ApplyT(func(v IPTag) string { return v.IpTagType }).(pulumi.StringOutput)
}

func (o IPTagOutput) Tag() pulumi.StringOutput {
	return o.ApplyT(func(v IPTag) string { return v.Tag }).(pulumi.StringOutput)
}

type IPTagArrayOutput struct{ *pulumi.OutputState }

func (IPTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPTag)(nil)).Elem()
}

func (o IPTagArrayOutput) ToIPTagArrayOutput() IPTagArrayOutput {
	return o
}

func (o IPTagArrayOutput) ToIPTagArrayOutputWithContext(ctx context.Context) IPTagArrayOutput {
	return o
}

func (o IPTagArrayOutput) Index(i pulumi.IntInput) IPTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPTag {
		return vs[0].([]IPTag)[vs[1].(int)]
	}).(IPTagOutput)
}

type IPTagResponse struct {
	IpTagType string `pulumi:"ipTagType"`
	Tag       string `pulumi:"tag"`
}

type IPTagResponseOutput struct{ *pulumi.OutputState }

func (IPTagResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPTagResponse)(nil)).Elem()
}

func (o IPTagResponseOutput) ToIPTagResponseOutput() IPTagResponseOutput {
	return o
}

func (o IPTagResponseOutput) ToIPTagResponseOutputWithContext(ctx context.Context) IPTagResponseOutput {
	return o
}

func (o IPTagResponseOutput) IpTagType() pulumi.StringOutput {
	return o.ApplyT(func(v IPTagResponse) string { return v.IpTagType }).(pulumi.StringOutput)
}

func (o IPTagResponseOutput) Tag() pulumi.StringOutput {
	return o.ApplyT(func(v IPTagResponse) string { return v.Tag }).(pulumi.StringOutput)
}

type IPTagResponseArrayOutput struct{ *pulumi.OutputState }

func (IPTagResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPTagResponse)(nil)).Elem()
}

func (o IPTagResponseArrayOutput) ToIPTagResponseArrayOutput() IPTagResponseArrayOutput {
	return o
}

func (o IPTagResponseArrayOutput) ToIPTagResponseArrayOutputWithContext(ctx context.Context) IPTagResponseArrayOutput {
	return o
}

func (o IPTagResponseArrayOutput) Index(i pulumi.IntInput) IPTagResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPTagResponse {
		return vs[0].([]IPTagResponse)[vs[1].(int)]
	}).(IPTagResponseOutput)
}

type LoadBalancingRule struct {
	BackendPort      int     `pulumi:"backendPort"`
	FrontendPort     int     `pulumi:"frontendPort"`
	LoadDistribution *string `pulumi:"loadDistribution"`
	ProbePort        *int    `pulumi:"probePort"`
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
	LoadDistribution pulumi.StringPtrInput `pulumi:"loadDistribution"`
	ProbePort        pulumi.IntPtrInput    `pulumi:"probePort"`
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

func (o LoadBalancingRuleOutput) LoadDistribution() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *string { return v.LoadDistribution }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleOutput) ProbePort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *int { return v.ProbePort }).(pulumi.IntPtrOutput)
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
	LoadDistribution *string `pulumi:"loadDistribution"`
	ProbePort        *int    `pulumi:"probePort"`
	ProbeProtocol    string  `pulumi:"probeProtocol"`
	ProbeRequestPath *string `pulumi:"probeRequestPath"`
	Protocol         string  `pulumi:"protocol"`
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

func (o LoadBalancingRuleResponseOutput) LoadDistribution() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *string { return v.LoadDistribution }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleResponseOutput) ProbePort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *int { return v.ProbePort }).(pulumi.IntPtrOutput)
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

type NamedPartitionSchemeResponse struct {
	Names           []string `pulumi:"names"`
	PartitionScheme string   `pulumi:"partitionScheme"`
}

type NetworkSecurityRule struct {
	Access                     string   `pulumi:"access"`
	Description                *string  `pulumi:"description"`
	DestinationAddressPrefix   *string  `pulumi:"destinationAddressPrefix"`
	DestinationAddressPrefixes []string `pulumi:"destinationAddressPrefixes"`
	DestinationPortRange       *string  `pulumi:"destinationPortRange"`
	DestinationPortRanges      []string `pulumi:"destinationPortRanges"`
	Direction                  string   `pulumi:"direction"`
	Name                       string   `pulumi:"name"`
	Priority                   int      `pulumi:"priority"`
	Protocol                   string   `pulumi:"protocol"`
	SourceAddressPrefix        *string  `pulumi:"sourceAddressPrefix"`
	SourceAddressPrefixes      []string `pulumi:"sourceAddressPrefixes"`
	SourcePortRange            *string  `pulumi:"sourcePortRange"`
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
	DestinationAddressPrefix   pulumi.StringPtrInput   `pulumi:"destinationAddressPrefix"`
	DestinationAddressPrefixes pulumi.StringArrayInput `pulumi:"destinationAddressPrefixes"`
	DestinationPortRange       pulumi.StringPtrInput   `pulumi:"destinationPortRange"`
	DestinationPortRanges      pulumi.StringArrayInput `pulumi:"destinationPortRanges"`
	Direction                  pulumi.StringInput      `pulumi:"direction"`
	Name                       pulumi.StringInput      `pulumi:"name"`
	Priority                   pulumi.IntInput         `pulumi:"priority"`
	Protocol                   pulumi.StringInput      `pulumi:"protocol"`
	SourceAddressPrefix        pulumi.StringPtrInput   `pulumi:"sourceAddressPrefix"`
	SourceAddressPrefixes      pulumi.StringArrayInput `pulumi:"sourceAddressPrefixes"`
	SourcePortRange            pulumi.StringPtrInput   `pulumi:"sourcePortRange"`
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

func (o NetworkSecurityRuleOutput) DestinationAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityRule) *string { return v.DestinationAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityRuleOutput) DestinationAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkSecurityRule) []string { return v.DestinationAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o NetworkSecurityRuleOutput) DestinationPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityRule) *string { return v.DestinationPortRange }).(pulumi.StringPtrOutput)
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

func (o NetworkSecurityRuleOutput) SourceAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityRule) *string { return v.SourceAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityRuleOutput) SourceAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkSecurityRule) []string { return v.SourceAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o NetworkSecurityRuleOutput) SourcePortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityRule) *string { return v.SourcePortRange }).(pulumi.StringPtrOutput)
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
	DestinationAddressPrefix   *string  `pulumi:"destinationAddressPrefix"`
	DestinationAddressPrefixes []string `pulumi:"destinationAddressPrefixes"`
	DestinationPortRange       *string  `pulumi:"destinationPortRange"`
	DestinationPortRanges      []string `pulumi:"destinationPortRanges"`
	Direction                  string   `pulumi:"direction"`
	Name                       string   `pulumi:"name"`
	Priority                   int      `pulumi:"priority"`
	Protocol                   string   `pulumi:"protocol"`
	SourceAddressPrefix        *string  `pulumi:"sourceAddressPrefix"`
	SourceAddressPrefixes      []string `pulumi:"sourceAddressPrefixes"`
	SourcePortRange            *string  `pulumi:"sourcePortRange"`
	SourcePortRanges           []string `pulumi:"sourcePortRanges"`
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

func (o NetworkSecurityRuleResponseOutput) DestinationAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityRuleResponse) *string { return v.DestinationAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityRuleResponseOutput) DestinationAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkSecurityRuleResponse) []string { return v.DestinationAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o NetworkSecurityRuleResponseOutput) DestinationPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityRuleResponse) *string { return v.DestinationPortRange }).(pulumi.StringPtrOutput)
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

func (o NetworkSecurityRuleResponseOutput) SourceAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityRuleResponse) *string { return v.SourceAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityRuleResponseOutput) SourceAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkSecurityRuleResponse) []string { return v.SourceAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o NetworkSecurityRuleResponseOutput) SourcePortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityRuleResponse) *string { return v.SourcePortRange }).(pulumi.StringPtrOutput)
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

type NodeTypeSku struct {
	Capacity int     `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}





type NodeTypeSkuInput interface {
	pulumi.Input

	ToNodeTypeSkuOutput() NodeTypeSkuOutput
	ToNodeTypeSkuOutputWithContext(context.Context) NodeTypeSkuOutput
}

type NodeTypeSkuArgs struct {
	Capacity pulumi.IntInput       `pulumi:"capacity"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (NodeTypeSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeTypeSku)(nil)).Elem()
}

func (i NodeTypeSkuArgs) ToNodeTypeSkuOutput() NodeTypeSkuOutput {
	return i.ToNodeTypeSkuOutputWithContext(context.Background())
}

func (i NodeTypeSkuArgs) ToNodeTypeSkuOutputWithContext(ctx context.Context) NodeTypeSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeTypeSkuOutput)
}

func (i NodeTypeSkuArgs) ToNodeTypeSkuPtrOutput() NodeTypeSkuPtrOutput {
	return i.ToNodeTypeSkuPtrOutputWithContext(context.Background())
}

func (i NodeTypeSkuArgs) ToNodeTypeSkuPtrOutputWithContext(ctx context.Context) NodeTypeSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeTypeSkuOutput).ToNodeTypeSkuPtrOutputWithContext(ctx)
}









type NodeTypeSkuPtrInput interface {
	pulumi.Input

	ToNodeTypeSkuPtrOutput() NodeTypeSkuPtrOutput
	ToNodeTypeSkuPtrOutputWithContext(context.Context) NodeTypeSkuPtrOutput
}

type nodeTypeSkuPtrType NodeTypeSkuArgs

func NodeTypeSkuPtr(v *NodeTypeSkuArgs) NodeTypeSkuPtrInput {
	return (*nodeTypeSkuPtrType)(v)
}

func (*nodeTypeSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeTypeSku)(nil)).Elem()
}

func (i *nodeTypeSkuPtrType) ToNodeTypeSkuPtrOutput() NodeTypeSkuPtrOutput {
	return i.ToNodeTypeSkuPtrOutputWithContext(context.Background())
}

func (i *nodeTypeSkuPtrType) ToNodeTypeSkuPtrOutputWithContext(ctx context.Context) NodeTypeSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeTypeSkuPtrOutput)
}

type NodeTypeSkuOutput struct{ *pulumi.OutputState }

func (NodeTypeSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeTypeSku)(nil)).Elem()
}

func (o NodeTypeSkuOutput) ToNodeTypeSkuOutput() NodeTypeSkuOutput {
	return o
}

func (o NodeTypeSkuOutput) ToNodeTypeSkuOutputWithContext(ctx context.Context) NodeTypeSkuOutput {
	return o
}

func (o NodeTypeSkuOutput) ToNodeTypeSkuPtrOutput() NodeTypeSkuPtrOutput {
	return o.ToNodeTypeSkuPtrOutputWithContext(context.Background())
}

func (o NodeTypeSkuOutput) ToNodeTypeSkuPtrOutputWithContext(ctx context.Context) NodeTypeSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NodeTypeSku) *NodeTypeSku {
		return &v
	}).(NodeTypeSkuPtrOutput)
}

func (o NodeTypeSkuOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v NodeTypeSku) int { return v.Capacity }).(pulumi.IntOutput)
}

func (o NodeTypeSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodeTypeSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NodeTypeSkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodeTypeSku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type NodeTypeSkuPtrOutput struct{ *pulumi.OutputState }

func (NodeTypeSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeTypeSku)(nil)).Elem()
}

func (o NodeTypeSkuPtrOutput) ToNodeTypeSkuPtrOutput() NodeTypeSkuPtrOutput {
	return o
}

func (o NodeTypeSkuPtrOutput) ToNodeTypeSkuPtrOutputWithContext(ctx context.Context) NodeTypeSkuPtrOutput {
	return o
}

func (o NodeTypeSkuPtrOutput) Elem() NodeTypeSkuOutput {
	return o.ApplyT(func(v *NodeTypeSku) NodeTypeSku {
		if v != nil {
			return *v
		}
		var ret NodeTypeSku
		return ret
	}).(NodeTypeSkuOutput)
}

func (o NodeTypeSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NodeTypeSku) *int {
		if v == nil {
			return nil
		}
		return &v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o NodeTypeSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodeTypeSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o NodeTypeSkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodeTypeSku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type NodeTypeSkuResponse struct {
	Capacity int     `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}

type NodeTypeSkuResponseOutput struct{ *pulumi.OutputState }

func (NodeTypeSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeTypeSkuResponse)(nil)).Elem()
}

func (o NodeTypeSkuResponseOutput) ToNodeTypeSkuResponseOutput() NodeTypeSkuResponseOutput {
	return o
}

func (o NodeTypeSkuResponseOutput) ToNodeTypeSkuResponseOutputWithContext(ctx context.Context) NodeTypeSkuResponseOutput {
	return o
}

func (o NodeTypeSkuResponseOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v NodeTypeSkuResponse) int { return v.Capacity }).(pulumi.IntOutput)
}

func (o NodeTypeSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodeTypeSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NodeTypeSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodeTypeSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type NodeTypeSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (NodeTypeSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeTypeSkuResponse)(nil)).Elem()
}

func (o NodeTypeSkuResponsePtrOutput) ToNodeTypeSkuResponsePtrOutput() NodeTypeSkuResponsePtrOutput {
	return o
}

func (o NodeTypeSkuResponsePtrOutput) ToNodeTypeSkuResponsePtrOutputWithContext(ctx context.Context) NodeTypeSkuResponsePtrOutput {
	return o
}

func (o NodeTypeSkuResponsePtrOutput) Elem() NodeTypeSkuResponseOutput {
	return o.ApplyT(func(v *NodeTypeSkuResponse) NodeTypeSkuResponse {
		if v != nil {
			return *v
		}
		var ret NodeTypeSkuResponse
		return ret
	}).(NodeTypeSkuResponseOutput)
}

func (o NodeTypeSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NodeTypeSkuResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o NodeTypeSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodeTypeSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o NodeTypeSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NodeTypeSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type PartitionInstanceCountScaleMechanism struct {
	Kind             string `pulumi:"kind"`
	MaxInstanceCount int    `pulumi:"maxInstanceCount"`
	MinInstanceCount int    `pulumi:"minInstanceCount"`
	ScaleIncrement   int    `pulumi:"scaleIncrement"`
}

type PartitionInstanceCountScaleMechanismResponse struct {
	Kind             string `pulumi:"kind"`
	MaxInstanceCount int    `pulumi:"maxInstanceCount"`
	MinInstanceCount int    `pulumi:"minInstanceCount"`
	ScaleIncrement   int    `pulumi:"scaleIncrement"`
}

type ResourceAzStatusResponse struct {
	IsZoneResilient bool   `pulumi:"isZoneResilient"`
	ResourceName    string `pulumi:"resourceName"`
	ResourceType    string `pulumi:"resourceType"`
}

type ResourceAzStatusResponseOutput struct{ *pulumi.OutputState }

func (ResourceAzStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceAzStatusResponse)(nil)).Elem()
}

func (o ResourceAzStatusResponseOutput) ToResourceAzStatusResponseOutput() ResourceAzStatusResponseOutput {
	return o
}

func (o ResourceAzStatusResponseOutput) ToResourceAzStatusResponseOutputWithContext(ctx context.Context) ResourceAzStatusResponseOutput {
	return o
}

func (o ResourceAzStatusResponseOutput) IsZoneResilient() pulumi.BoolOutput {
	return o.ApplyT(func(v ResourceAzStatusResponse) bool { return v.IsZoneResilient }).(pulumi.BoolOutput)
}

func (o ResourceAzStatusResponseOutput) ResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceAzStatusResponse) string { return v.ResourceName }).(pulumi.StringOutput)
}

func (o ResourceAzStatusResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceAzStatusResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

type ResourceAzStatusResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceAzStatusResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceAzStatusResponse)(nil)).Elem()
}

func (o ResourceAzStatusResponseArrayOutput) ToResourceAzStatusResponseArrayOutput() ResourceAzStatusResponseArrayOutput {
	return o
}

func (o ResourceAzStatusResponseArrayOutput) ToResourceAzStatusResponseArrayOutputWithContext(ctx context.Context) ResourceAzStatusResponseArrayOutput {
	return o
}

func (o ResourceAzStatusResponseArrayOutput) Index(i pulumi.IntInput) ResourceAzStatusResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceAzStatusResponse {
		return vs[0].([]ResourceAzStatusResponse)[vs[1].(int)]
	}).(ResourceAzStatusResponseOutput)
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

type ScalingPolicyResponse struct {
	ScalingMechanism interface{} `pulumi:"scalingMechanism"`
	ScalingTrigger   interface{} `pulumi:"scalingTrigger"`
}

type ServiceCorrelation struct {
	Scheme      string `pulumi:"scheme"`
	ServiceName string `pulumi:"serviceName"`
}

type ServiceCorrelationResponse struct {
	Scheme      string `pulumi:"scheme"`
	ServiceName string `pulumi:"serviceName"`
}

type ServiceEndpoint struct {
	Locations []string `pulumi:"locations"`
	Service   string   `pulumi:"service"`
}





type ServiceEndpointInput interface {
	pulumi.Input

	ToServiceEndpointOutput() ServiceEndpointOutput
	ToServiceEndpointOutputWithContext(context.Context) ServiceEndpointOutput
}

type ServiceEndpointArgs struct {
	Locations pulumi.StringArrayInput `pulumi:"locations"`
	Service   pulumi.StringInput      `pulumi:"service"`
}

func (ServiceEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpoint)(nil)).Elem()
}

func (i ServiceEndpointArgs) ToServiceEndpointOutput() ServiceEndpointOutput {
	return i.ToServiceEndpointOutputWithContext(context.Background())
}

func (i ServiceEndpointArgs) ToServiceEndpointOutputWithContext(ctx context.Context) ServiceEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointOutput)
}





type ServiceEndpointArrayInput interface {
	pulumi.Input

	ToServiceEndpointArrayOutput() ServiceEndpointArrayOutput
	ToServiceEndpointArrayOutputWithContext(context.Context) ServiceEndpointArrayOutput
}

type ServiceEndpointArray []ServiceEndpointInput

func (ServiceEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpoint)(nil)).Elem()
}

func (i ServiceEndpointArray) ToServiceEndpointArrayOutput() ServiceEndpointArrayOutput {
	return i.ToServiceEndpointArrayOutputWithContext(context.Background())
}

func (i ServiceEndpointArray) ToServiceEndpointArrayOutputWithContext(ctx context.Context) ServiceEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointArrayOutput)
}

type ServiceEndpointOutput struct{ *pulumi.OutputState }

func (ServiceEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpoint)(nil)).Elem()
}

func (o ServiceEndpointOutput) ToServiceEndpointOutput() ServiceEndpointOutput {
	return o
}

func (o ServiceEndpointOutput) ToServiceEndpointOutputWithContext(ctx context.Context) ServiceEndpointOutput {
	return o
}

func (o ServiceEndpointOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceEndpoint) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o ServiceEndpointOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceEndpoint) string { return v.Service }).(pulumi.StringOutput)
}

type ServiceEndpointArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpoint)(nil)).Elem()
}

func (o ServiceEndpointArrayOutput) ToServiceEndpointArrayOutput() ServiceEndpointArrayOutput {
	return o
}

func (o ServiceEndpointArrayOutput) ToServiceEndpointArrayOutputWithContext(ctx context.Context) ServiceEndpointArrayOutput {
	return o
}

func (o ServiceEndpointArrayOutput) Index(i pulumi.IntInput) ServiceEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceEndpoint {
		return vs[0].([]ServiceEndpoint)[vs[1].(int)]
	}).(ServiceEndpointOutput)
}

type ServiceEndpointResponse struct {
	Locations []string `pulumi:"locations"`
	Service   string   `pulumi:"service"`
}

type ServiceEndpointResponseOutput struct{ *pulumi.OutputState }

func (ServiceEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceEndpointResponse)(nil)).Elem()
}

func (o ServiceEndpointResponseOutput) ToServiceEndpointResponseOutput() ServiceEndpointResponseOutput {
	return o
}

func (o ServiceEndpointResponseOutput) ToServiceEndpointResponseOutputWithContext(ctx context.Context) ServiceEndpointResponseOutput {
	return o
}

func (o ServiceEndpointResponseOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceEndpointResponse) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o ServiceEndpointResponseOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceEndpointResponse) string { return v.Service }).(pulumi.StringOutput)
}

type ServiceEndpointResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceEndpointResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceEndpointResponse)(nil)).Elem()
}

func (o ServiceEndpointResponseArrayOutput) ToServiceEndpointResponseArrayOutput() ServiceEndpointResponseArrayOutput {
	return o
}

func (o ServiceEndpointResponseArrayOutput) ToServiceEndpointResponseArrayOutputWithContext(ctx context.Context) ServiceEndpointResponseArrayOutput {
	return o
}

func (o ServiceEndpointResponseArrayOutput) Index(i pulumi.IntInput) ServiceEndpointResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceEndpointResponse {
		return vs[0].([]ServiceEndpointResponse)[vs[1].(int)]
	}).(ServiceEndpointResponseOutput)
}

type ServiceLoadMetric struct {
	DefaultLoad          *int    `pulumi:"defaultLoad"`
	Name                 string  `pulumi:"name"`
	PrimaryDefaultLoad   *int    `pulumi:"primaryDefaultLoad"`
	SecondaryDefaultLoad *int    `pulumi:"secondaryDefaultLoad"`
	Weight               *string `pulumi:"weight"`
}

type ServiceLoadMetricResponse struct {
	DefaultLoad          *int    `pulumi:"defaultLoad"`
	Name                 string  `pulumi:"name"`
	PrimaryDefaultLoad   *int    `pulumi:"primaryDefaultLoad"`
	SecondaryDefaultLoad *int    `pulumi:"secondaryDefaultLoad"`
	Weight               *string `pulumi:"weight"`
}

type ServicePlacementInvalidDomainPolicy struct {
	DomainName string `pulumi:"domainName"`
	Type       string `pulumi:"type"`
}

type ServicePlacementInvalidDomainPolicyResponse struct {
	DomainName string `pulumi:"domainName"`
	Type       string `pulumi:"type"`
}

type ServicePlacementNonPartiallyPlaceServicePolicy struct {
	Type string `pulumi:"type"`
}

type ServicePlacementNonPartiallyPlaceServicePolicyResponse struct {
	Type string `pulumi:"type"`
}

type ServicePlacementPreferPrimaryDomainPolicy struct {
	DomainName string `pulumi:"domainName"`
	Type       string `pulumi:"type"`
}

type ServicePlacementPreferPrimaryDomainPolicyResponse struct {
	DomainName string `pulumi:"domainName"`
	Type       string `pulumi:"type"`
}

type ServicePlacementRequireDomainDistributionPolicy struct {
	DomainName string `pulumi:"domainName"`
	Type       string `pulumi:"type"`
}

type ServicePlacementRequireDomainDistributionPolicyResponse struct {
	DomainName string `pulumi:"domainName"`
	Type       string `pulumi:"type"`
}

type ServicePlacementRequiredDomainPolicy struct {
	DomainName string `pulumi:"domainName"`
	Type       string `pulumi:"type"`
}

type ServicePlacementRequiredDomainPolicyResponse struct {
	DomainName string `pulumi:"domainName"`
	Type       string `pulumi:"type"`
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

type SingletonPartitionSchemeResponse struct {
	PartitionScheme string `pulumi:"partitionScheme"`
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

type Subnet struct {
	EnableIpv6                        *bool   `pulumi:"enableIpv6"`
	Name                              string  `pulumi:"name"`
	NetworkSecurityGroupId            *string `pulumi:"networkSecurityGroupId"`
	PrivateEndpointNetworkPolicies    *string `pulumi:"privateEndpointNetworkPolicies"`
	PrivateLinkServiceNetworkPolicies *string `pulumi:"privateLinkServiceNetworkPolicies"`
}





type SubnetInput interface {
	pulumi.Input

	ToSubnetOutput() SubnetOutput
	ToSubnetOutputWithContext(context.Context) SubnetOutput
}

type SubnetArgs struct {
	EnableIpv6                        pulumi.BoolPtrInput   `pulumi:"enableIpv6"`
	Name                              pulumi.StringInput    `pulumi:"name"`
	NetworkSecurityGroupId            pulumi.StringPtrInput `pulumi:"networkSecurityGroupId"`
	PrivateEndpointNetworkPolicies    pulumi.StringPtrInput `pulumi:"privateEndpointNetworkPolicies"`
	PrivateLinkServiceNetworkPolicies pulumi.StringPtrInput `pulumi:"privateLinkServiceNetworkPolicies"`
}

func (SubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Subnet)(nil)).Elem()
}

func (i SubnetArgs) ToSubnetOutput() SubnetOutput {
	return i.ToSubnetOutputWithContext(context.Background())
}

func (i SubnetArgs) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetOutput)
}





type SubnetArrayInput interface {
	pulumi.Input

	ToSubnetArrayOutput() SubnetArrayOutput
	ToSubnetArrayOutputWithContext(context.Context) SubnetArrayOutput
}

type SubnetArray []SubnetInput

func (SubnetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Subnet)(nil)).Elem()
}

func (i SubnetArray) ToSubnetArrayOutput() SubnetArrayOutput {
	return i.ToSubnetArrayOutputWithContext(context.Background())
}

func (i SubnetArray) ToSubnetArrayOutputWithContext(ctx context.Context) SubnetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetArrayOutput)
}

type SubnetOutput struct{ *pulumi.OutputState }

func (SubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Subnet)(nil)).Elem()
}

func (o SubnetOutput) ToSubnetOutput() SubnetOutput {
	return o
}

func (o SubnetOutput) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return o
}

func (o SubnetOutput) EnableIpv6() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Subnet) *bool { return v.EnableIpv6 }).(pulumi.BoolPtrOutput)
}

func (o SubnetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Subnet) string { return v.Name }).(pulumi.StringOutput)
}

func (o SubnetOutput) NetworkSecurityGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Subnet) *string { return v.NetworkSecurityGroupId }).(pulumi.StringPtrOutput)
}

func (o SubnetOutput) PrivateEndpointNetworkPolicies() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Subnet) *string { return v.PrivateEndpointNetworkPolicies }).(pulumi.StringPtrOutput)
}

func (o SubnetOutput) PrivateLinkServiceNetworkPolicies() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Subnet) *string { return v.PrivateLinkServiceNetworkPolicies }).(pulumi.StringPtrOutput)
}

type SubnetArrayOutput struct{ *pulumi.OutputState }

func (SubnetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Subnet)(nil)).Elem()
}

func (o SubnetArrayOutput) ToSubnetArrayOutput() SubnetArrayOutput {
	return o
}

func (o SubnetArrayOutput) ToSubnetArrayOutputWithContext(ctx context.Context) SubnetArrayOutput {
	return o
}

func (o SubnetArrayOutput) Index(i pulumi.IntInput) SubnetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Subnet {
		return vs[0].([]Subnet)[vs[1].(int)]
	}).(SubnetOutput)
}

type SubnetResponse struct {
	EnableIpv6                        *bool   `pulumi:"enableIpv6"`
	Name                              string  `pulumi:"name"`
	NetworkSecurityGroupId            *string `pulumi:"networkSecurityGroupId"`
	PrivateEndpointNetworkPolicies    *string `pulumi:"privateEndpointNetworkPolicies"`
	PrivateLinkServiceNetworkPolicies *string `pulumi:"privateLinkServiceNetworkPolicies"`
}

type SubnetResponseOutput struct{ *pulumi.OutputState }

func (SubnetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetResponse)(nil)).Elem()
}

func (o SubnetResponseOutput) ToSubnetResponseOutput() SubnetResponseOutput {
	return o
}

func (o SubnetResponseOutput) ToSubnetResponseOutputWithContext(ctx context.Context) SubnetResponseOutput {
	return o
}

func (o SubnetResponseOutput) EnableIpv6() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *bool { return v.EnableIpv6 }).(pulumi.BoolPtrOutput)
}

func (o SubnetResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SubnetResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SubnetResponseOutput) NetworkSecurityGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.NetworkSecurityGroupId }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) PrivateEndpointNetworkPolicies() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.PrivateEndpointNetworkPolicies }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) PrivateLinkServiceNetworkPolicies() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.PrivateLinkServiceNetworkPolicies }).(pulumi.StringPtrOutput)
}

type SubnetResponseArrayOutput struct{ *pulumi.OutputState }

func (SubnetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetResponse)(nil)).Elem()
}

func (o SubnetResponseArrayOutput) ToSubnetResponseArrayOutput() SubnetResponseArrayOutput {
	return o
}

func (o SubnetResponseArrayOutput) ToSubnetResponseArrayOutputWithContext(ctx context.Context) SubnetResponseArrayOutput {
	return o
}

func (o SubnetResponseArrayOutput) Index(i pulumi.IntInput) SubnetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubnetResponse {
		return vs[0].([]SubnetResponse)[vs[1].(int)]
	}).(SubnetResponseOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
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

type UniformInt64RangePartitionScheme struct {
	Count           int     `pulumi:"count"`
	HighKey         float64 `pulumi:"highKey"`
	LowKey          float64 `pulumi:"lowKey"`
	PartitionScheme string  `pulumi:"partitionScheme"`
}

type UniformInt64RangePartitionSchemeResponse struct {
	Count           int     `pulumi:"count"`
	HighKey         float64 `pulumi:"highKey"`
	LowKey          float64 `pulumi:"lowKey"`
	PartitionScheme string  `pulumi:"partitionScheme"`
}

type UserAssignedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
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
	EnableAutomaticUpgrade   *bool       `pulumi:"enableAutomaticUpgrade"`
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
	EnableAutomaticUpgrade   pulumi.BoolPtrInput     `pulumi:"enableAutomaticUpgrade"`
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

func (o VMSSExtensionOutput) EnableAutomaticUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VMSSExtension) *bool { return v.EnableAutomaticUpgrade }).(pulumi.BoolPtrOutput)
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
	EnableAutomaticUpgrade   *bool       `pulumi:"enableAutomaticUpgrade"`
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

func (o VMSSExtensionResponseOutput) EnableAutomaticUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) *bool { return v.EnableAutomaticUpgrade }).(pulumi.BoolPtrOutput)
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

type VmssDataDisk struct {
	DiskLetter string `pulumi:"diskLetter"`
	DiskSizeGB int    `pulumi:"diskSizeGB"`
	DiskType   string `pulumi:"diskType"`
	Lun        int    `pulumi:"lun"`
}





type VmssDataDiskInput interface {
	pulumi.Input

	ToVmssDataDiskOutput() VmssDataDiskOutput
	ToVmssDataDiskOutputWithContext(context.Context) VmssDataDiskOutput
}

type VmssDataDiskArgs struct {
	DiskLetter pulumi.StringInput `pulumi:"diskLetter"`
	DiskSizeGB pulumi.IntInput    `pulumi:"diskSizeGB"`
	DiskType   pulumi.StringInput `pulumi:"diskType"`
	Lun        pulumi.IntInput    `pulumi:"lun"`
}

func (VmssDataDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmssDataDisk)(nil)).Elem()
}

func (i VmssDataDiskArgs) ToVmssDataDiskOutput() VmssDataDiskOutput {
	return i.ToVmssDataDiskOutputWithContext(context.Background())
}

func (i VmssDataDiskArgs) ToVmssDataDiskOutputWithContext(ctx context.Context) VmssDataDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmssDataDiskOutput)
}





type VmssDataDiskArrayInput interface {
	pulumi.Input

	ToVmssDataDiskArrayOutput() VmssDataDiskArrayOutput
	ToVmssDataDiskArrayOutputWithContext(context.Context) VmssDataDiskArrayOutput
}

type VmssDataDiskArray []VmssDataDiskInput

func (VmssDataDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmssDataDisk)(nil)).Elem()
}

func (i VmssDataDiskArray) ToVmssDataDiskArrayOutput() VmssDataDiskArrayOutput {
	return i.ToVmssDataDiskArrayOutputWithContext(context.Background())
}

func (i VmssDataDiskArray) ToVmssDataDiskArrayOutputWithContext(ctx context.Context) VmssDataDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmssDataDiskArrayOutput)
}

type VmssDataDiskOutput struct{ *pulumi.OutputState }

func (VmssDataDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmssDataDisk)(nil)).Elem()
}

func (o VmssDataDiskOutput) ToVmssDataDiskOutput() VmssDataDiskOutput {
	return o
}

func (o VmssDataDiskOutput) ToVmssDataDiskOutputWithContext(ctx context.Context) VmssDataDiskOutput {
	return o
}

func (o VmssDataDiskOutput) DiskLetter() pulumi.StringOutput {
	return o.ApplyT(func(v VmssDataDisk) string { return v.DiskLetter }).(pulumi.StringOutput)
}

func (o VmssDataDiskOutput) DiskSizeGB() pulumi.IntOutput {
	return o.ApplyT(func(v VmssDataDisk) int { return v.DiskSizeGB }).(pulumi.IntOutput)
}

func (o VmssDataDiskOutput) DiskType() pulumi.StringOutput {
	return o.ApplyT(func(v VmssDataDisk) string { return v.DiskType }).(pulumi.StringOutput)
}

func (o VmssDataDiskOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v VmssDataDisk) int { return v.Lun }).(pulumi.IntOutput)
}

type VmssDataDiskArrayOutput struct{ *pulumi.OutputState }

func (VmssDataDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmssDataDisk)(nil)).Elem()
}

func (o VmssDataDiskArrayOutput) ToVmssDataDiskArrayOutput() VmssDataDiskArrayOutput {
	return o
}

func (o VmssDataDiskArrayOutput) ToVmssDataDiskArrayOutputWithContext(ctx context.Context) VmssDataDiskArrayOutput {
	return o
}

func (o VmssDataDiskArrayOutput) Index(i pulumi.IntInput) VmssDataDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VmssDataDisk {
		return vs[0].([]VmssDataDisk)[vs[1].(int)]
	}).(VmssDataDiskOutput)
}

type VmssDataDiskResponse struct {
	DiskLetter string `pulumi:"diskLetter"`
	DiskSizeGB int    `pulumi:"diskSizeGB"`
	DiskType   string `pulumi:"diskType"`
	Lun        int    `pulumi:"lun"`
}

type VmssDataDiskResponseOutput struct{ *pulumi.OutputState }

func (VmssDataDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmssDataDiskResponse)(nil)).Elem()
}

func (o VmssDataDiskResponseOutput) ToVmssDataDiskResponseOutput() VmssDataDiskResponseOutput {
	return o
}

func (o VmssDataDiskResponseOutput) ToVmssDataDiskResponseOutputWithContext(ctx context.Context) VmssDataDiskResponseOutput {
	return o
}

func (o VmssDataDiskResponseOutput) DiskLetter() pulumi.StringOutput {
	return o.ApplyT(func(v VmssDataDiskResponse) string { return v.DiskLetter }).(pulumi.StringOutput)
}

func (o VmssDataDiskResponseOutput) DiskSizeGB() pulumi.IntOutput {
	return o.ApplyT(func(v VmssDataDiskResponse) int { return v.DiskSizeGB }).(pulumi.IntOutput)
}

func (o VmssDataDiskResponseOutput) DiskType() pulumi.StringOutput {
	return o.ApplyT(func(v VmssDataDiskResponse) string { return v.DiskType }).(pulumi.StringOutput)
}

func (o VmssDataDiskResponseOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v VmssDataDiskResponse) int { return v.Lun }).(pulumi.IntOutput)
}

type VmssDataDiskResponseArrayOutput struct{ *pulumi.OutputState }

func (VmssDataDiskResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmssDataDiskResponse)(nil)).Elem()
}

func (o VmssDataDiskResponseArrayOutput) ToVmssDataDiskResponseArrayOutput() VmssDataDiskResponseArrayOutput {
	return o
}

func (o VmssDataDiskResponseArrayOutput) ToVmssDataDiskResponseArrayOutputWithContext(ctx context.Context) VmssDataDiskResponseArrayOutput {
	return o
}

func (o VmssDataDiskResponseArrayOutput) Index(i pulumi.IntInput) VmssDataDiskResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VmssDataDiskResponse {
		return vs[0].([]VmssDataDiskResponse)[vs[1].(int)]
	}).(VmssDataDiskResponseOutput)
}

func init() {
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
	pulumi.RegisterOutputType(FrontendConfigurationOutput{})
	pulumi.RegisterOutputType(FrontendConfigurationArrayOutput{})
	pulumi.RegisterOutputType(FrontendConfigurationResponseOutput{})
	pulumi.RegisterOutputType(FrontendConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(IPTagOutput{})
	pulumi.RegisterOutputType(IPTagArrayOutput{})
	pulumi.RegisterOutputType(IPTagResponseOutput{})
	pulumi.RegisterOutputType(IPTagResponseArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleResponseOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedIdentityOutput{})
	pulumi.RegisterOutputType(ManagedIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkSecurityRuleOutput{})
	pulumi.RegisterOutputType(NetworkSecurityRuleArrayOutput{})
	pulumi.RegisterOutputType(NetworkSecurityRuleResponseOutput{})
	pulumi.RegisterOutputType(NetworkSecurityRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(NodeTypeSkuOutput{})
	pulumi.RegisterOutputType(NodeTypeSkuPtrOutput{})
	pulumi.RegisterOutputType(NodeTypeSkuResponseOutput{})
	pulumi.RegisterOutputType(NodeTypeSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceAzStatusResponseOutput{})
	pulumi.RegisterOutputType(ResourceAzStatusResponseArrayOutput{})
	pulumi.RegisterOutputType(RollingUpgradeMonitoringPolicyOutput{})
	pulumi.RegisterOutputType(RollingUpgradeMonitoringPolicyPtrOutput{})
	pulumi.RegisterOutputType(RollingUpgradeMonitoringPolicyResponseOutput{})
	pulumi.RegisterOutputType(RollingUpgradeMonitoringPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceEndpointOutput{})
	pulumi.RegisterOutputType(ServiceEndpointArrayOutput{})
	pulumi.RegisterOutputType(ServiceEndpointResponseOutput{})
	pulumi.RegisterOutputType(ServiceEndpointResponseArrayOutput{})
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
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SubResourceOutput{})
	pulumi.RegisterOutputType(SubResourceResponseOutput{})
	pulumi.RegisterOutputType(SubnetOutput{})
	pulumi.RegisterOutputType(SubnetArrayOutput{})
	pulumi.RegisterOutputType(SubnetResponseOutput{})
	pulumi.RegisterOutputType(SubnetResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
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
	pulumi.RegisterOutputType(VmssDataDiskOutput{})
	pulumi.RegisterOutputType(VmssDataDiskArrayOutput{})
	pulumi.RegisterOutputType(VmssDataDiskResponseOutput{})
	pulumi.RegisterOutputType(VmssDataDiskResponseArrayOutput{})
}
