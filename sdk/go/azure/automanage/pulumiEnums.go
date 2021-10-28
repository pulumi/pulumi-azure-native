


package automanage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigurationProfile string

const (
	ConfigurationProfile_Azure_virtual_machine_best_practices_Dev_Test   = ConfigurationProfile("Azure virtual machine best practices – Dev/Test")
	ConfigurationProfile_Azure_virtual_machine_best_practices_Production = ConfigurationProfile("Azure virtual machine best practices – Production")
)

func (ConfigurationProfile) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfile)(nil)).Elem()
}

func (e ConfigurationProfile) ToConfigurationProfileOutput() ConfigurationProfileOutput {
	return pulumi.ToOutput(e).(ConfigurationProfileOutput)
}

func (e ConfigurationProfile) ToConfigurationProfileOutputWithContext(ctx context.Context) ConfigurationProfileOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConfigurationProfileOutput)
}

func (e ConfigurationProfile) ToConfigurationProfilePtrOutput() ConfigurationProfilePtrOutput {
	return e.ToConfigurationProfilePtrOutputWithContext(context.Background())
}

func (e ConfigurationProfile) ToConfigurationProfilePtrOutputWithContext(ctx context.Context) ConfigurationProfilePtrOutput {
	return ConfigurationProfile(e).ToConfigurationProfileOutputWithContext(ctx).ToConfigurationProfilePtrOutputWithContext(ctx)
}

func (e ConfigurationProfile) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConfigurationProfile) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConfigurationProfile) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConfigurationProfile) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConfigurationProfileOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfile)(nil)).Elem()
}

func (o ConfigurationProfileOutput) ToConfigurationProfileOutput() ConfigurationProfileOutput {
	return o
}

func (o ConfigurationProfileOutput) ToConfigurationProfileOutputWithContext(ctx context.Context) ConfigurationProfileOutput {
	return o
}

func (o ConfigurationProfileOutput) ToConfigurationProfilePtrOutput() ConfigurationProfilePtrOutput {
	return o.ToConfigurationProfilePtrOutputWithContext(context.Background())
}

func (o ConfigurationProfileOutput) ToConfigurationProfilePtrOutputWithContext(ctx context.Context) ConfigurationProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationProfile) *ConfigurationProfile {
		return &v
	}).(ConfigurationProfilePtrOutput)
}

func (o ConfigurationProfileOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConfigurationProfileOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConfigurationProfile) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConfigurationProfileOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConfigurationProfileOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConfigurationProfile) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConfigurationProfilePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfile)(nil)).Elem()
}

func (o ConfigurationProfilePtrOutput) ToConfigurationProfilePtrOutput() ConfigurationProfilePtrOutput {
	return o
}

func (o ConfigurationProfilePtrOutput) ToConfigurationProfilePtrOutputWithContext(ctx context.Context) ConfigurationProfilePtrOutput {
	return o
}

func (o ConfigurationProfilePtrOutput) Elem() ConfigurationProfileOutput {
	return o.ApplyT(func(v *ConfigurationProfile) ConfigurationProfile {
		if v != nil {
			return *v
		}
		var ret ConfigurationProfile
		return ret
	}).(ConfigurationProfileOutput)
}

func (o ConfigurationProfilePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConfigurationProfilePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConfigurationProfile) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ConfigurationProfileInput interface {
	pulumi.Input

	ToConfigurationProfileOutput() ConfigurationProfileOutput
	ToConfigurationProfileOutputWithContext(context.Context) ConfigurationProfileOutput
}

var configurationProfilePtrType = reflect.TypeOf((**ConfigurationProfile)(nil)).Elem()

type ConfigurationProfilePtrInput interface {
	pulumi.Input

	ToConfigurationProfilePtrOutput() ConfigurationProfilePtrOutput
	ToConfigurationProfilePtrOutputWithContext(context.Context) ConfigurationProfilePtrOutput
}

type configurationProfilePtr string

func ConfigurationProfilePtr(v string) ConfigurationProfilePtrInput {
	return (*configurationProfilePtr)(&v)
}

func (*configurationProfilePtr) ElementType() reflect.Type {
	return configurationProfilePtrType
}

func (in *configurationProfilePtr) ToConfigurationProfilePtrOutput() ConfigurationProfilePtrOutput {
	return pulumi.ToOutput(in).(ConfigurationProfilePtrOutput)
}

func (in *configurationProfilePtr) ToConfigurationProfilePtrOutputWithContext(ctx context.Context) ConfigurationProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConfigurationProfilePtrOutput)
}

type EnableRealTimeProtection string

const (
	EnableRealTimeProtectionTrue  = EnableRealTimeProtection("True")
	EnableRealTimeProtectionFalse = EnableRealTimeProtection("False")
)

func (EnableRealTimeProtection) ElementType() reflect.Type {
	return reflect.TypeOf((*EnableRealTimeProtection)(nil)).Elem()
}

func (e EnableRealTimeProtection) ToEnableRealTimeProtectionOutput() EnableRealTimeProtectionOutput {
	return pulumi.ToOutput(e).(EnableRealTimeProtectionOutput)
}

func (e EnableRealTimeProtection) ToEnableRealTimeProtectionOutputWithContext(ctx context.Context) EnableRealTimeProtectionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EnableRealTimeProtectionOutput)
}

func (e EnableRealTimeProtection) ToEnableRealTimeProtectionPtrOutput() EnableRealTimeProtectionPtrOutput {
	return e.ToEnableRealTimeProtectionPtrOutputWithContext(context.Background())
}

func (e EnableRealTimeProtection) ToEnableRealTimeProtectionPtrOutputWithContext(ctx context.Context) EnableRealTimeProtectionPtrOutput {
	return EnableRealTimeProtection(e).ToEnableRealTimeProtectionOutputWithContext(ctx).ToEnableRealTimeProtectionPtrOutputWithContext(ctx)
}

func (e EnableRealTimeProtection) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnableRealTimeProtection) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnableRealTimeProtection) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EnableRealTimeProtection) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EnableRealTimeProtectionOutput struct{ *pulumi.OutputState }

func (EnableRealTimeProtectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnableRealTimeProtection)(nil)).Elem()
}

func (o EnableRealTimeProtectionOutput) ToEnableRealTimeProtectionOutput() EnableRealTimeProtectionOutput {
	return o
}

func (o EnableRealTimeProtectionOutput) ToEnableRealTimeProtectionOutputWithContext(ctx context.Context) EnableRealTimeProtectionOutput {
	return o
}

func (o EnableRealTimeProtectionOutput) ToEnableRealTimeProtectionPtrOutput() EnableRealTimeProtectionPtrOutput {
	return o.ToEnableRealTimeProtectionPtrOutputWithContext(context.Background())
}

func (o EnableRealTimeProtectionOutput) ToEnableRealTimeProtectionPtrOutputWithContext(ctx context.Context) EnableRealTimeProtectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnableRealTimeProtection) *EnableRealTimeProtection {
		return &v
	}).(EnableRealTimeProtectionPtrOutput)
}

func (o EnableRealTimeProtectionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EnableRealTimeProtectionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnableRealTimeProtection) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EnableRealTimeProtectionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnableRealTimeProtectionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnableRealTimeProtection) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EnableRealTimeProtectionPtrOutput struct{ *pulumi.OutputState }

func (EnableRealTimeProtectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnableRealTimeProtection)(nil)).Elem()
}

func (o EnableRealTimeProtectionPtrOutput) ToEnableRealTimeProtectionPtrOutput() EnableRealTimeProtectionPtrOutput {
	return o
}

func (o EnableRealTimeProtectionPtrOutput) ToEnableRealTimeProtectionPtrOutputWithContext(ctx context.Context) EnableRealTimeProtectionPtrOutput {
	return o
}

func (o EnableRealTimeProtectionPtrOutput) Elem() EnableRealTimeProtectionOutput {
	return o.ApplyT(func(v *EnableRealTimeProtection) EnableRealTimeProtection {
		if v != nil {
			return *v
		}
		var ret EnableRealTimeProtection
		return ret
	}).(EnableRealTimeProtectionOutput)
}

func (o EnableRealTimeProtectionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnableRealTimeProtectionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EnableRealTimeProtection) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EnableRealTimeProtectionInput interface {
	pulumi.Input

	ToEnableRealTimeProtectionOutput() EnableRealTimeProtectionOutput
	ToEnableRealTimeProtectionOutputWithContext(context.Context) EnableRealTimeProtectionOutput
}

var enableRealTimeProtectionPtrType = reflect.TypeOf((**EnableRealTimeProtection)(nil)).Elem()

type EnableRealTimeProtectionPtrInput interface {
	pulumi.Input

	ToEnableRealTimeProtectionPtrOutput() EnableRealTimeProtectionPtrOutput
	ToEnableRealTimeProtectionPtrOutputWithContext(context.Context) EnableRealTimeProtectionPtrOutput
}

type enableRealTimeProtectionPtr string

func EnableRealTimeProtectionPtr(v string) EnableRealTimeProtectionPtrInput {
	return (*enableRealTimeProtectionPtr)(&v)
}

func (*enableRealTimeProtectionPtr) ElementType() reflect.Type {
	return enableRealTimeProtectionPtrType
}

func (in *enableRealTimeProtectionPtr) ToEnableRealTimeProtectionPtrOutput() EnableRealTimeProtectionPtrOutput {
	return pulumi.ToOutput(in).(EnableRealTimeProtectionPtrOutput)
}

func (in *enableRealTimeProtectionPtr) ToEnableRealTimeProtectionPtrOutputWithContext(ctx context.Context) EnableRealTimeProtectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EnableRealTimeProtectionPtrOutput)
}

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeNone           = ResourceIdentityType("None")
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

type RunScheduledScan string

const (
	RunScheduledScanTrue  = RunScheduledScan("True")
	RunScheduledScanFalse = RunScheduledScan("False")
)

func (RunScheduledScan) ElementType() reflect.Type {
	return reflect.TypeOf((*RunScheduledScan)(nil)).Elem()
}

func (e RunScheduledScan) ToRunScheduledScanOutput() RunScheduledScanOutput {
	return pulumi.ToOutput(e).(RunScheduledScanOutput)
}

func (e RunScheduledScan) ToRunScheduledScanOutputWithContext(ctx context.Context) RunScheduledScanOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RunScheduledScanOutput)
}

func (e RunScheduledScan) ToRunScheduledScanPtrOutput() RunScheduledScanPtrOutput {
	return e.ToRunScheduledScanPtrOutputWithContext(context.Background())
}

func (e RunScheduledScan) ToRunScheduledScanPtrOutputWithContext(ctx context.Context) RunScheduledScanPtrOutput {
	return RunScheduledScan(e).ToRunScheduledScanOutputWithContext(ctx).ToRunScheduledScanPtrOutputWithContext(ctx)
}

func (e RunScheduledScan) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RunScheduledScan) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RunScheduledScan) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RunScheduledScan) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RunScheduledScanOutput struct{ *pulumi.OutputState }

func (RunScheduledScanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunScheduledScan)(nil)).Elem()
}

func (o RunScheduledScanOutput) ToRunScheduledScanOutput() RunScheduledScanOutput {
	return o
}

func (o RunScheduledScanOutput) ToRunScheduledScanOutputWithContext(ctx context.Context) RunScheduledScanOutput {
	return o
}

func (o RunScheduledScanOutput) ToRunScheduledScanPtrOutput() RunScheduledScanPtrOutput {
	return o.ToRunScheduledScanPtrOutputWithContext(context.Background())
}

func (o RunScheduledScanOutput) ToRunScheduledScanPtrOutputWithContext(ctx context.Context) RunScheduledScanPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RunScheduledScan) *RunScheduledScan {
		return &v
	}).(RunScheduledScanPtrOutput)
}

func (o RunScheduledScanOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RunScheduledScanOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RunScheduledScan) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RunScheduledScanOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RunScheduledScanOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RunScheduledScan) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RunScheduledScanPtrOutput struct{ *pulumi.OutputState }

func (RunScheduledScanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunScheduledScan)(nil)).Elem()
}

func (o RunScheduledScanPtrOutput) ToRunScheduledScanPtrOutput() RunScheduledScanPtrOutput {
	return o
}

func (o RunScheduledScanPtrOutput) ToRunScheduledScanPtrOutputWithContext(ctx context.Context) RunScheduledScanPtrOutput {
	return o
}

func (o RunScheduledScanPtrOutput) Elem() RunScheduledScanOutput {
	return o.ApplyT(func(v *RunScheduledScan) RunScheduledScan {
		if v != nil {
			return *v
		}
		var ret RunScheduledScan
		return ret
	}).(RunScheduledScanOutput)
}

func (o RunScheduledScanPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RunScheduledScanPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RunScheduledScan) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RunScheduledScanInput interface {
	pulumi.Input

	ToRunScheduledScanOutput() RunScheduledScanOutput
	ToRunScheduledScanOutputWithContext(context.Context) RunScheduledScanOutput
}

var runScheduledScanPtrType = reflect.TypeOf((**RunScheduledScan)(nil)).Elem()

type RunScheduledScanPtrInput interface {
	pulumi.Input

	ToRunScheduledScanPtrOutput() RunScheduledScanPtrOutput
	ToRunScheduledScanPtrOutputWithContext(context.Context) RunScheduledScanPtrOutput
}

type runScheduledScanPtr string

func RunScheduledScanPtr(v string) RunScheduledScanPtrInput {
	return (*runScheduledScanPtr)(&v)
}

func (*runScheduledScanPtr) ElementType() reflect.Type {
	return runScheduledScanPtrType
}

func (in *runScheduledScanPtr) ToRunScheduledScanPtrOutput() RunScheduledScanPtrOutput {
	return pulumi.ToOutput(in).(RunScheduledScanPtrOutput)
}

func (in *runScheduledScanPtr) ToRunScheduledScanPtrOutputWithContext(ctx context.Context) RunScheduledScanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RunScheduledScanPtrOutput)
}

type ScanType string

const (
	ScanTypeQuick = ScanType("Quick")
	ScanTypeFull  = ScanType("Full")
)

func (ScanType) ElementType() reflect.Type {
	return reflect.TypeOf((*ScanType)(nil)).Elem()
}

func (e ScanType) ToScanTypeOutput() ScanTypeOutput {
	return pulumi.ToOutput(e).(ScanTypeOutput)
}

func (e ScanType) ToScanTypeOutputWithContext(ctx context.Context) ScanTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScanTypeOutput)
}

func (e ScanType) ToScanTypePtrOutput() ScanTypePtrOutput {
	return e.ToScanTypePtrOutputWithContext(context.Background())
}

func (e ScanType) ToScanTypePtrOutputWithContext(ctx context.Context) ScanTypePtrOutput {
	return ScanType(e).ToScanTypeOutputWithContext(ctx).ToScanTypePtrOutputWithContext(ctx)
}

func (e ScanType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScanType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScanType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScanType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScanTypeOutput struct{ *pulumi.OutputState }

func (ScanTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScanType)(nil)).Elem()
}

func (o ScanTypeOutput) ToScanTypeOutput() ScanTypeOutput {
	return o
}

func (o ScanTypeOutput) ToScanTypeOutputWithContext(ctx context.Context) ScanTypeOutput {
	return o
}

func (o ScanTypeOutput) ToScanTypePtrOutput() ScanTypePtrOutput {
	return o.ToScanTypePtrOutputWithContext(context.Background())
}

func (o ScanTypeOutput) ToScanTypePtrOutputWithContext(ctx context.Context) ScanTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScanType) *ScanType {
		return &v
	}).(ScanTypePtrOutput)
}

func (o ScanTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScanTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScanType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScanTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScanTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScanType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScanTypePtrOutput struct{ *pulumi.OutputState }

func (ScanTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScanType)(nil)).Elem()
}

func (o ScanTypePtrOutput) ToScanTypePtrOutput() ScanTypePtrOutput {
	return o
}

func (o ScanTypePtrOutput) ToScanTypePtrOutputWithContext(ctx context.Context) ScanTypePtrOutput {
	return o
}

func (o ScanTypePtrOutput) Elem() ScanTypeOutput {
	return o.ApplyT(func(v *ScanType) ScanType {
		if v != nil {
			return *v
		}
		var ret ScanType
		return ret
	}).(ScanTypeOutput)
}

func (o ScanTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScanTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScanType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ScanTypeInput interface {
	pulumi.Input

	ToScanTypeOutput() ScanTypeOutput
	ToScanTypeOutputWithContext(context.Context) ScanTypeOutput
}

var scanTypePtrType = reflect.TypeOf((**ScanType)(nil)).Elem()

type ScanTypePtrInput interface {
	pulumi.Input

	ToScanTypePtrOutput() ScanTypePtrOutput
	ToScanTypePtrOutputWithContext(context.Context) ScanTypePtrOutput
}

type scanTypePtr string

func ScanTypePtr(v string) ScanTypePtrInput {
	return (*scanTypePtr)(&v)
}

func (*scanTypePtr) ElementType() reflect.Type {
	return scanTypePtrType
}

func (in *scanTypePtr) ToScanTypePtrOutput() ScanTypePtrOutput {
	return pulumi.ToOutput(in).(ScanTypePtrOutput)
}

func (in *scanTypePtr) ToScanTypePtrOutputWithContext(ctx context.Context) ScanTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScanTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationProfileInput)(nil)).Elem(), ConfigurationProfile("Azure virtual machine best practices – Dev/Test"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationProfilePtrInput)(nil)).Elem(), ConfigurationProfile("Azure virtual machine best practices – Dev/Test"))
	pulumi.RegisterInputType(reflect.TypeOf((*EnableRealTimeProtectionInput)(nil)).Elem(), EnableRealTimeProtection("True"))
	pulumi.RegisterInputType(reflect.TypeOf((*EnableRealTimeProtectionPtrInput)(nil)).Elem(), EnableRealTimeProtection("True"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypeInput)(nil)).Elem(), ResourceIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypePtrInput)(nil)).Elem(), ResourceIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*RunScheduledScanInput)(nil)).Elem(), RunScheduledScan("True"))
	pulumi.RegisterInputType(reflect.TypeOf((*RunScheduledScanPtrInput)(nil)).Elem(), RunScheduledScan("True"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScanTypeInput)(nil)).Elem(), ScanType("Quick"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScanTypePtrInput)(nil)).Elem(), ScanType("Quick"))
	pulumi.RegisterOutputType(ConfigurationProfileOutput{})
	pulumi.RegisterOutputType(ConfigurationProfilePtrOutput{})
	pulumi.RegisterOutputType(EnableRealTimeProtectionOutput{})
	pulumi.RegisterOutputType(EnableRealTimeProtectionPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(RunScheduledScanOutput{})
	pulumi.RegisterOutputType(RunScheduledScanPtrOutput{})
	pulumi.RegisterOutputType(ScanTypeOutput{})
	pulumi.RegisterOutputType(ScanTypePtrOutput{})
}
