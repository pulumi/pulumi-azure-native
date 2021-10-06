


package desktopvirtualization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationGroupType string

const (
	ApplicationGroupTypeRemoteApp = ApplicationGroupType("RemoteApp")
	ApplicationGroupTypeDesktop   = ApplicationGroupType("Desktop")
)

func (ApplicationGroupType) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGroupType)(nil)).Elem()
}

func (e ApplicationGroupType) ToApplicationGroupTypeOutput() ApplicationGroupTypeOutput {
	return pulumi.ToOutput(e).(ApplicationGroupTypeOutput)
}

func (e ApplicationGroupType) ToApplicationGroupTypeOutputWithContext(ctx context.Context) ApplicationGroupTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ApplicationGroupTypeOutput)
}

func (e ApplicationGroupType) ToApplicationGroupTypePtrOutput() ApplicationGroupTypePtrOutput {
	return e.ToApplicationGroupTypePtrOutputWithContext(context.Background())
}

func (e ApplicationGroupType) ToApplicationGroupTypePtrOutputWithContext(ctx context.Context) ApplicationGroupTypePtrOutput {
	return ApplicationGroupType(e).ToApplicationGroupTypeOutputWithContext(ctx).ToApplicationGroupTypePtrOutputWithContext(ctx)
}

func (e ApplicationGroupType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGroupType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ApplicationGroupType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ApplicationGroupType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ApplicationGroupTypeOutput struct{ *pulumi.OutputState }

func (ApplicationGroupTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGroupType)(nil)).Elem()
}

func (o ApplicationGroupTypeOutput) ToApplicationGroupTypeOutput() ApplicationGroupTypeOutput {
	return o
}

func (o ApplicationGroupTypeOutput) ToApplicationGroupTypeOutputWithContext(ctx context.Context) ApplicationGroupTypeOutput {
	return o
}

func (o ApplicationGroupTypeOutput) ToApplicationGroupTypePtrOutput() ApplicationGroupTypePtrOutput {
	return o.ToApplicationGroupTypePtrOutputWithContext(context.Background())
}

func (o ApplicationGroupTypeOutput) ToApplicationGroupTypePtrOutputWithContext(ctx context.Context) ApplicationGroupTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGroupType) *ApplicationGroupType {
		return &v
	}).(ApplicationGroupTypePtrOutput)
}

func (o ApplicationGroupTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ApplicationGroupTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGroupType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ApplicationGroupTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGroupTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ApplicationGroupType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ApplicationGroupTypePtrOutput struct{ *pulumi.OutputState }

func (ApplicationGroupTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGroupType)(nil)).Elem()
}

func (o ApplicationGroupTypePtrOutput) ToApplicationGroupTypePtrOutput() ApplicationGroupTypePtrOutput {
	return o
}

func (o ApplicationGroupTypePtrOutput) ToApplicationGroupTypePtrOutputWithContext(ctx context.Context) ApplicationGroupTypePtrOutput {
	return o
}

func (o ApplicationGroupTypePtrOutput) Elem() ApplicationGroupTypeOutput {
	return o.ApplyT(func(v *ApplicationGroupType) ApplicationGroupType {
		if v != nil {
			return *v
		}
		var ret ApplicationGroupType
		return ret
	}).(ApplicationGroupTypeOutput)
}

func (o ApplicationGroupTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ApplicationGroupTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ApplicationGroupType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ApplicationGroupTypeInput interface {
	pulumi.Input

	ToApplicationGroupTypeOutput() ApplicationGroupTypeOutput
	ToApplicationGroupTypeOutputWithContext(context.Context) ApplicationGroupTypeOutput
}

var applicationGroupTypePtrType = reflect.TypeOf((**ApplicationGroupType)(nil)).Elem()

type ApplicationGroupTypePtrInput interface {
	pulumi.Input

	ToApplicationGroupTypePtrOutput() ApplicationGroupTypePtrOutput
	ToApplicationGroupTypePtrOutputWithContext(context.Context) ApplicationGroupTypePtrOutput
}

type applicationGroupTypePtr string

func ApplicationGroupTypePtr(v string) ApplicationGroupTypePtrInput {
	return (*applicationGroupTypePtr)(&v)
}

func (*applicationGroupTypePtr) ElementType() reflect.Type {
	return applicationGroupTypePtrType
}

func (in *applicationGroupTypePtr) ToApplicationGroupTypePtrOutput() ApplicationGroupTypePtrOutput {
	return pulumi.ToOutput(in).(ApplicationGroupTypePtrOutput)
}

func (in *applicationGroupTypePtr) ToApplicationGroupTypePtrOutputWithContext(ctx context.Context) ApplicationGroupTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ApplicationGroupTypePtrOutput)
}

type CommandLineSetting string

const (
	CommandLineSettingDoNotAllow = CommandLineSetting("DoNotAllow")
	CommandLineSettingAllow      = CommandLineSetting("Allow")
	CommandLineSettingRequire    = CommandLineSetting("Require")
)

func (CommandLineSetting) ElementType() reflect.Type {
	return reflect.TypeOf((*CommandLineSetting)(nil)).Elem()
}

func (e CommandLineSetting) ToCommandLineSettingOutput() CommandLineSettingOutput {
	return pulumi.ToOutput(e).(CommandLineSettingOutput)
}

func (e CommandLineSetting) ToCommandLineSettingOutputWithContext(ctx context.Context) CommandLineSettingOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CommandLineSettingOutput)
}

func (e CommandLineSetting) ToCommandLineSettingPtrOutput() CommandLineSettingPtrOutput {
	return e.ToCommandLineSettingPtrOutputWithContext(context.Background())
}

func (e CommandLineSetting) ToCommandLineSettingPtrOutputWithContext(ctx context.Context) CommandLineSettingPtrOutput {
	return CommandLineSetting(e).ToCommandLineSettingOutputWithContext(ctx).ToCommandLineSettingPtrOutputWithContext(ctx)
}

func (e CommandLineSetting) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CommandLineSetting) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CommandLineSetting) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CommandLineSetting) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CommandLineSettingOutput struct{ *pulumi.OutputState }

func (CommandLineSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommandLineSetting)(nil)).Elem()
}

func (o CommandLineSettingOutput) ToCommandLineSettingOutput() CommandLineSettingOutput {
	return o
}

func (o CommandLineSettingOutput) ToCommandLineSettingOutputWithContext(ctx context.Context) CommandLineSettingOutput {
	return o
}

func (o CommandLineSettingOutput) ToCommandLineSettingPtrOutput() CommandLineSettingPtrOutput {
	return o.ToCommandLineSettingPtrOutputWithContext(context.Background())
}

func (o CommandLineSettingOutput) ToCommandLineSettingPtrOutputWithContext(ctx context.Context) CommandLineSettingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CommandLineSetting) *CommandLineSetting {
		return &v
	}).(CommandLineSettingPtrOutput)
}

func (o CommandLineSettingOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CommandLineSettingOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CommandLineSetting) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CommandLineSettingOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CommandLineSettingOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CommandLineSetting) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CommandLineSettingPtrOutput struct{ *pulumi.OutputState }

func (CommandLineSettingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommandLineSetting)(nil)).Elem()
}

func (o CommandLineSettingPtrOutput) ToCommandLineSettingPtrOutput() CommandLineSettingPtrOutput {
	return o
}

func (o CommandLineSettingPtrOutput) ToCommandLineSettingPtrOutputWithContext(ctx context.Context) CommandLineSettingPtrOutput {
	return o
}

func (o CommandLineSettingPtrOutput) Elem() CommandLineSettingOutput {
	return o.ApplyT(func(v *CommandLineSetting) CommandLineSetting {
		if v != nil {
			return *v
		}
		var ret CommandLineSetting
		return ret
	}).(CommandLineSettingOutput)
}

func (o CommandLineSettingPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CommandLineSettingPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CommandLineSetting) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CommandLineSettingInput interface {
	pulumi.Input

	ToCommandLineSettingOutput() CommandLineSettingOutput
	ToCommandLineSettingOutputWithContext(context.Context) CommandLineSettingOutput
}

var commandLineSettingPtrType = reflect.TypeOf((**CommandLineSetting)(nil)).Elem()

type CommandLineSettingPtrInput interface {
	pulumi.Input

	ToCommandLineSettingPtrOutput() CommandLineSettingPtrOutput
	ToCommandLineSettingPtrOutputWithContext(context.Context) CommandLineSettingPtrOutput
}

type commandLineSettingPtr string

func CommandLineSettingPtr(v string) CommandLineSettingPtrInput {
	return (*commandLineSettingPtr)(&v)
}

func (*commandLineSettingPtr) ElementType() reflect.Type {
	return commandLineSettingPtrType
}

func (in *commandLineSettingPtr) ToCommandLineSettingPtrOutput() CommandLineSettingPtrOutput {
	return pulumi.ToOutput(in).(CommandLineSettingPtrOutput)
}

func (in *commandLineSettingPtr) ToCommandLineSettingPtrOutputWithContext(ctx context.Context) CommandLineSettingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CommandLineSettingPtrOutput)
}

type HostPoolType string

const (
	// Users will be assigned a SessionHost either by administrators (PersonalDesktopAssignmentType = Direct) or upon connecting to the pool (PersonalDesktopAssignmentType = Automatic). They will always be redirected to their assigned SessionHost.
	HostPoolTypePersonal = HostPoolType("Personal")
	// Users get a new (random) SessionHost every time it connects to the HostPool.
	HostPoolTypePooled = HostPoolType("Pooled")
	// Users assign their own machines, load balancing logic remains the same as Personal. PersonalDesktopAssignmentType must be Direct.
	HostPoolTypeBYODesktop = HostPoolType("BYODesktop")
)

func (HostPoolType) ElementType() reflect.Type {
	return reflect.TypeOf((*HostPoolType)(nil)).Elem()
}

func (e HostPoolType) ToHostPoolTypeOutput() HostPoolTypeOutput {
	return pulumi.ToOutput(e).(HostPoolTypeOutput)
}

func (e HostPoolType) ToHostPoolTypeOutputWithContext(ctx context.Context) HostPoolTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HostPoolTypeOutput)
}

func (e HostPoolType) ToHostPoolTypePtrOutput() HostPoolTypePtrOutput {
	return e.ToHostPoolTypePtrOutputWithContext(context.Background())
}

func (e HostPoolType) ToHostPoolTypePtrOutputWithContext(ctx context.Context) HostPoolTypePtrOutput {
	return HostPoolType(e).ToHostPoolTypeOutputWithContext(ctx).ToHostPoolTypePtrOutputWithContext(ctx)
}

func (e HostPoolType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostPoolType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostPoolType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HostPoolType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HostPoolTypeOutput struct{ *pulumi.OutputState }

func (HostPoolTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostPoolType)(nil)).Elem()
}

func (o HostPoolTypeOutput) ToHostPoolTypeOutput() HostPoolTypeOutput {
	return o
}

func (o HostPoolTypeOutput) ToHostPoolTypeOutputWithContext(ctx context.Context) HostPoolTypeOutput {
	return o
}

func (o HostPoolTypeOutput) ToHostPoolTypePtrOutput() HostPoolTypePtrOutput {
	return o.ToHostPoolTypePtrOutputWithContext(context.Background())
}

func (o HostPoolTypeOutput) ToHostPoolTypePtrOutputWithContext(ctx context.Context) HostPoolTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HostPoolType) *HostPoolType {
		return &v
	}).(HostPoolTypePtrOutput)
}

func (o HostPoolTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HostPoolTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostPoolType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HostPoolTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostPoolTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostPoolType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HostPoolTypePtrOutput struct{ *pulumi.OutputState }

func (HostPoolTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostPoolType)(nil)).Elem()
}

func (o HostPoolTypePtrOutput) ToHostPoolTypePtrOutput() HostPoolTypePtrOutput {
	return o
}

func (o HostPoolTypePtrOutput) ToHostPoolTypePtrOutputWithContext(ctx context.Context) HostPoolTypePtrOutput {
	return o
}

func (o HostPoolTypePtrOutput) Elem() HostPoolTypeOutput {
	return o.ApplyT(func(v *HostPoolType) HostPoolType {
		if v != nil {
			return *v
		}
		var ret HostPoolType
		return ret
	}).(HostPoolTypeOutput)
}

func (o HostPoolTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostPoolTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HostPoolType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HostPoolTypeInput interface {
	pulumi.Input

	ToHostPoolTypeOutput() HostPoolTypeOutput
	ToHostPoolTypeOutputWithContext(context.Context) HostPoolTypeOutput
}

var hostPoolTypePtrType = reflect.TypeOf((**HostPoolType)(nil)).Elem()

type HostPoolTypePtrInput interface {
	pulumi.Input

	ToHostPoolTypePtrOutput() HostPoolTypePtrOutput
	ToHostPoolTypePtrOutputWithContext(context.Context) HostPoolTypePtrOutput
}

type hostPoolTypePtr string

func HostPoolTypePtr(v string) HostPoolTypePtrInput {
	return (*hostPoolTypePtr)(&v)
}

func (*hostPoolTypePtr) ElementType() reflect.Type {
	return hostPoolTypePtrType
}

func (in *hostPoolTypePtr) ToHostPoolTypePtrOutput() HostPoolTypePtrOutput {
	return pulumi.ToOutput(in).(HostPoolTypePtrOutput)
}

func (in *hostPoolTypePtr) ToHostPoolTypePtrOutputWithContext(ctx context.Context) HostPoolTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HostPoolTypePtrOutput)
}

type LoadBalancerType string

const (
	LoadBalancerTypeBreadthFirst = LoadBalancerType("BreadthFirst")
	LoadBalancerTypeDepthFirst   = LoadBalancerType("DepthFirst")
	LoadBalancerTypePersistent   = LoadBalancerType("Persistent")
)

func (LoadBalancerType) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerType)(nil)).Elem()
}

func (e LoadBalancerType) ToLoadBalancerTypeOutput() LoadBalancerTypeOutput {
	return pulumi.ToOutput(e).(LoadBalancerTypeOutput)
}

func (e LoadBalancerType) ToLoadBalancerTypeOutputWithContext(ctx context.Context) LoadBalancerTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LoadBalancerTypeOutput)
}

func (e LoadBalancerType) ToLoadBalancerTypePtrOutput() LoadBalancerTypePtrOutput {
	return e.ToLoadBalancerTypePtrOutputWithContext(context.Background())
}

func (e LoadBalancerType) ToLoadBalancerTypePtrOutputWithContext(ctx context.Context) LoadBalancerTypePtrOutput {
	return LoadBalancerType(e).ToLoadBalancerTypeOutputWithContext(ctx).ToLoadBalancerTypePtrOutputWithContext(ctx)
}

func (e LoadBalancerType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LoadBalancerType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LoadBalancerType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LoadBalancerType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LoadBalancerTypeOutput struct{ *pulumi.OutputState }

func (LoadBalancerTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerType)(nil)).Elem()
}

func (o LoadBalancerTypeOutput) ToLoadBalancerTypeOutput() LoadBalancerTypeOutput {
	return o
}

func (o LoadBalancerTypeOutput) ToLoadBalancerTypeOutputWithContext(ctx context.Context) LoadBalancerTypeOutput {
	return o
}

func (o LoadBalancerTypeOutput) ToLoadBalancerTypePtrOutput() LoadBalancerTypePtrOutput {
	return o.ToLoadBalancerTypePtrOutputWithContext(context.Background())
}

func (o LoadBalancerTypeOutput) ToLoadBalancerTypePtrOutputWithContext(ctx context.Context) LoadBalancerTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoadBalancerType) *LoadBalancerType {
		return &v
	}).(LoadBalancerTypePtrOutput)
}

func (o LoadBalancerTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LoadBalancerTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LoadBalancerType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LoadBalancerTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LoadBalancerTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LoadBalancerType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LoadBalancerTypePtrOutput struct{ *pulumi.OutputState }

func (LoadBalancerTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerType)(nil)).Elem()
}

func (o LoadBalancerTypePtrOutput) ToLoadBalancerTypePtrOutput() LoadBalancerTypePtrOutput {
	return o
}

func (o LoadBalancerTypePtrOutput) ToLoadBalancerTypePtrOutputWithContext(ctx context.Context) LoadBalancerTypePtrOutput {
	return o
}

func (o LoadBalancerTypePtrOutput) Elem() LoadBalancerTypeOutput {
	return o.ApplyT(func(v *LoadBalancerType) LoadBalancerType {
		if v != nil {
			return *v
		}
		var ret LoadBalancerType
		return ret
	}).(LoadBalancerTypeOutput)
}

func (o LoadBalancerTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LoadBalancerTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LoadBalancerType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LoadBalancerTypeInput interface {
	pulumi.Input

	ToLoadBalancerTypeOutput() LoadBalancerTypeOutput
	ToLoadBalancerTypeOutputWithContext(context.Context) LoadBalancerTypeOutput
}

var loadBalancerTypePtrType = reflect.TypeOf((**LoadBalancerType)(nil)).Elem()

type LoadBalancerTypePtrInput interface {
	pulumi.Input

	ToLoadBalancerTypePtrOutput() LoadBalancerTypePtrOutput
	ToLoadBalancerTypePtrOutputWithContext(context.Context) LoadBalancerTypePtrOutput
}

type loadBalancerTypePtr string

func LoadBalancerTypePtr(v string) LoadBalancerTypePtrInput {
	return (*loadBalancerTypePtr)(&v)
}

func (*loadBalancerTypePtr) ElementType() reflect.Type {
	return loadBalancerTypePtrType
}

func (in *loadBalancerTypePtr) ToLoadBalancerTypePtrOutput() LoadBalancerTypePtrOutput {
	return pulumi.ToOutput(in).(LoadBalancerTypePtrOutput)
}

func (in *loadBalancerTypePtr) ToLoadBalancerTypePtrOutputWithContext(ctx context.Context) LoadBalancerTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LoadBalancerTypePtrOutput)
}

type Operation string

const (
	// Start the migration.
	OperationStart = Operation("Start")
	// Revoke the migration.
	OperationRevoke = Operation("Revoke")
	// Complete the migration.
	OperationComplete = Operation("Complete")
	// Hide the hostpool.
	OperationHide = Operation("Hide")
	// Unhide the hostpool.
	OperationUnhide = Operation("Unhide")
)

func (Operation) ElementType() reflect.Type {
	return reflect.TypeOf((*Operation)(nil)).Elem()
}

func (e Operation) ToOperationOutput() OperationOutput {
	return pulumi.ToOutput(e).(OperationOutput)
}

func (e Operation) ToOperationOutputWithContext(ctx context.Context) OperationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OperationOutput)
}

func (e Operation) ToOperationPtrOutput() OperationPtrOutput {
	return e.ToOperationPtrOutputWithContext(context.Background())
}

func (e Operation) ToOperationPtrOutputWithContext(ctx context.Context) OperationPtrOutput {
	return Operation(e).ToOperationOutputWithContext(ctx).ToOperationPtrOutputWithContext(ctx)
}

func (e Operation) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Operation) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Operation) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Operation) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OperationOutput struct{ *pulumi.OutputState }

func (OperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Operation)(nil)).Elem()
}

func (o OperationOutput) ToOperationOutput() OperationOutput {
	return o
}

func (o OperationOutput) ToOperationOutputWithContext(ctx context.Context) OperationOutput {
	return o
}

func (o OperationOutput) ToOperationPtrOutput() OperationPtrOutput {
	return o.ToOperationPtrOutputWithContext(context.Background())
}

func (o OperationOutput) ToOperationPtrOutputWithContext(ctx context.Context) OperationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Operation) *Operation {
		return &v
	}).(OperationPtrOutput)
}

func (o OperationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OperationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Operation) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OperationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Operation) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OperationPtrOutput struct{ *pulumi.OutputState }

func (OperationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Operation)(nil)).Elem()
}

func (o OperationPtrOutput) ToOperationPtrOutput() OperationPtrOutput {
	return o
}

func (o OperationPtrOutput) ToOperationPtrOutputWithContext(ctx context.Context) OperationPtrOutput {
	return o
}

func (o OperationPtrOutput) Elem() OperationOutput {
	return o.ApplyT(func(v *Operation) Operation {
		if v != nil {
			return *v
		}
		var ret Operation
		return ret
	}).(OperationOutput)
}

func (o OperationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Operation) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OperationInput interface {
	pulumi.Input

	ToOperationOutput() OperationOutput
	ToOperationOutputWithContext(context.Context) OperationOutput
}

var operationPtrType = reflect.TypeOf((**Operation)(nil)).Elem()

type OperationPtrInput interface {
	pulumi.Input

	ToOperationPtrOutput() OperationPtrOutput
	ToOperationPtrOutputWithContext(context.Context) OperationPtrOutput
}

type operationPtr string

func OperationPtr(v string) OperationPtrInput {
	return (*operationPtr)(&v)
}

func (*operationPtr) ElementType() reflect.Type {
	return operationPtrType
}

func (in *operationPtr) ToOperationPtrOutput() OperationPtrOutput {
	return pulumi.ToOutput(in).(OperationPtrOutput)
}

func (in *operationPtr) ToOperationPtrOutputWithContext(ctx context.Context) OperationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OperationPtrOutput)
}

type PersonalDesktopAssignmentType string

const (
	PersonalDesktopAssignmentTypeAutomatic = PersonalDesktopAssignmentType("Automatic")
	PersonalDesktopAssignmentTypeDirect    = PersonalDesktopAssignmentType("Direct")
)

func (PersonalDesktopAssignmentType) ElementType() reflect.Type {
	return reflect.TypeOf((*PersonalDesktopAssignmentType)(nil)).Elem()
}

func (e PersonalDesktopAssignmentType) ToPersonalDesktopAssignmentTypeOutput() PersonalDesktopAssignmentTypeOutput {
	return pulumi.ToOutput(e).(PersonalDesktopAssignmentTypeOutput)
}

func (e PersonalDesktopAssignmentType) ToPersonalDesktopAssignmentTypeOutputWithContext(ctx context.Context) PersonalDesktopAssignmentTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PersonalDesktopAssignmentTypeOutput)
}

func (e PersonalDesktopAssignmentType) ToPersonalDesktopAssignmentTypePtrOutput() PersonalDesktopAssignmentTypePtrOutput {
	return e.ToPersonalDesktopAssignmentTypePtrOutputWithContext(context.Background())
}

func (e PersonalDesktopAssignmentType) ToPersonalDesktopAssignmentTypePtrOutputWithContext(ctx context.Context) PersonalDesktopAssignmentTypePtrOutput {
	return PersonalDesktopAssignmentType(e).ToPersonalDesktopAssignmentTypeOutputWithContext(ctx).ToPersonalDesktopAssignmentTypePtrOutputWithContext(ctx)
}

func (e PersonalDesktopAssignmentType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PersonalDesktopAssignmentType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PersonalDesktopAssignmentType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PersonalDesktopAssignmentType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PersonalDesktopAssignmentTypeOutput struct{ *pulumi.OutputState }

func (PersonalDesktopAssignmentTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PersonalDesktopAssignmentType)(nil)).Elem()
}

func (o PersonalDesktopAssignmentTypeOutput) ToPersonalDesktopAssignmentTypeOutput() PersonalDesktopAssignmentTypeOutput {
	return o
}

func (o PersonalDesktopAssignmentTypeOutput) ToPersonalDesktopAssignmentTypeOutputWithContext(ctx context.Context) PersonalDesktopAssignmentTypeOutput {
	return o
}

func (o PersonalDesktopAssignmentTypeOutput) ToPersonalDesktopAssignmentTypePtrOutput() PersonalDesktopAssignmentTypePtrOutput {
	return o.ToPersonalDesktopAssignmentTypePtrOutputWithContext(context.Background())
}

func (o PersonalDesktopAssignmentTypeOutput) ToPersonalDesktopAssignmentTypePtrOutputWithContext(ctx context.Context) PersonalDesktopAssignmentTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PersonalDesktopAssignmentType) *PersonalDesktopAssignmentType {
		return &v
	}).(PersonalDesktopAssignmentTypePtrOutput)
}

func (o PersonalDesktopAssignmentTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PersonalDesktopAssignmentTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PersonalDesktopAssignmentType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PersonalDesktopAssignmentTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PersonalDesktopAssignmentTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PersonalDesktopAssignmentType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PersonalDesktopAssignmentTypePtrOutput struct{ *pulumi.OutputState }

func (PersonalDesktopAssignmentTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PersonalDesktopAssignmentType)(nil)).Elem()
}

func (o PersonalDesktopAssignmentTypePtrOutput) ToPersonalDesktopAssignmentTypePtrOutput() PersonalDesktopAssignmentTypePtrOutput {
	return o
}

func (o PersonalDesktopAssignmentTypePtrOutput) ToPersonalDesktopAssignmentTypePtrOutputWithContext(ctx context.Context) PersonalDesktopAssignmentTypePtrOutput {
	return o
}

func (o PersonalDesktopAssignmentTypePtrOutput) Elem() PersonalDesktopAssignmentTypeOutput {
	return o.ApplyT(func(v *PersonalDesktopAssignmentType) PersonalDesktopAssignmentType {
		if v != nil {
			return *v
		}
		var ret PersonalDesktopAssignmentType
		return ret
	}).(PersonalDesktopAssignmentTypeOutput)
}

func (o PersonalDesktopAssignmentTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PersonalDesktopAssignmentTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PersonalDesktopAssignmentType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PersonalDesktopAssignmentTypeInput interface {
	pulumi.Input

	ToPersonalDesktopAssignmentTypeOutput() PersonalDesktopAssignmentTypeOutput
	ToPersonalDesktopAssignmentTypeOutputWithContext(context.Context) PersonalDesktopAssignmentTypeOutput
}

var personalDesktopAssignmentTypePtrType = reflect.TypeOf((**PersonalDesktopAssignmentType)(nil)).Elem()

type PersonalDesktopAssignmentTypePtrInput interface {
	pulumi.Input

	ToPersonalDesktopAssignmentTypePtrOutput() PersonalDesktopAssignmentTypePtrOutput
	ToPersonalDesktopAssignmentTypePtrOutputWithContext(context.Context) PersonalDesktopAssignmentTypePtrOutput
}

type personalDesktopAssignmentTypePtr string

func PersonalDesktopAssignmentTypePtr(v string) PersonalDesktopAssignmentTypePtrInput {
	return (*personalDesktopAssignmentTypePtr)(&v)
}

func (*personalDesktopAssignmentTypePtr) ElementType() reflect.Type {
	return personalDesktopAssignmentTypePtrType
}

func (in *personalDesktopAssignmentTypePtr) ToPersonalDesktopAssignmentTypePtrOutput() PersonalDesktopAssignmentTypePtrOutput {
	return pulumi.ToOutput(in).(PersonalDesktopAssignmentTypePtrOutput)
}

func (in *personalDesktopAssignmentTypePtr) ToPersonalDesktopAssignmentTypePtrOutputWithContext(ctx context.Context) PersonalDesktopAssignmentTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PersonalDesktopAssignmentTypePtrOutput)
}

type PreferredAppGroupType string

const (
	PreferredAppGroupTypeNone             = PreferredAppGroupType("None")
	PreferredAppGroupTypeDesktop          = PreferredAppGroupType("Desktop")
	PreferredAppGroupTypeRailApplications = PreferredAppGroupType("RailApplications")
)

func (PreferredAppGroupType) ElementType() reflect.Type {
	return reflect.TypeOf((*PreferredAppGroupType)(nil)).Elem()
}

func (e PreferredAppGroupType) ToPreferredAppGroupTypeOutput() PreferredAppGroupTypeOutput {
	return pulumi.ToOutput(e).(PreferredAppGroupTypeOutput)
}

func (e PreferredAppGroupType) ToPreferredAppGroupTypeOutputWithContext(ctx context.Context) PreferredAppGroupTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PreferredAppGroupTypeOutput)
}

func (e PreferredAppGroupType) ToPreferredAppGroupTypePtrOutput() PreferredAppGroupTypePtrOutput {
	return e.ToPreferredAppGroupTypePtrOutputWithContext(context.Background())
}

func (e PreferredAppGroupType) ToPreferredAppGroupTypePtrOutputWithContext(ctx context.Context) PreferredAppGroupTypePtrOutput {
	return PreferredAppGroupType(e).ToPreferredAppGroupTypeOutputWithContext(ctx).ToPreferredAppGroupTypePtrOutputWithContext(ctx)
}

func (e PreferredAppGroupType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PreferredAppGroupType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PreferredAppGroupType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PreferredAppGroupType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PreferredAppGroupTypeOutput struct{ *pulumi.OutputState }

func (PreferredAppGroupTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PreferredAppGroupType)(nil)).Elem()
}

func (o PreferredAppGroupTypeOutput) ToPreferredAppGroupTypeOutput() PreferredAppGroupTypeOutput {
	return o
}

func (o PreferredAppGroupTypeOutput) ToPreferredAppGroupTypeOutputWithContext(ctx context.Context) PreferredAppGroupTypeOutput {
	return o
}

func (o PreferredAppGroupTypeOutput) ToPreferredAppGroupTypePtrOutput() PreferredAppGroupTypePtrOutput {
	return o.ToPreferredAppGroupTypePtrOutputWithContext(context.Background())
}

func (o PreferredAppGroupTypeOutput) ToPreferredAppGroupTypePtrOutputWithContext(ctx context.Context) PreferredAppGroupTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PreferredAppGroupType) *PreferredAppGroupType {
		return &v
	}).(PreferredAppGroupTypePtrOutput)
}

func (o PreferredAppGroupTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PreferredAppGroupTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PreferredAppGroupType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PreferredAppGroupTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PreferredAppGroupTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PreferredAppGroupType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PreferredAppGroupTypePtrOutput struct{ *pulumi.OutputState }

func (PreferredAppGroupTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PreferredAppGroupType)(nil)).Elem()
}

func (o PreferredAppGroupTypePtrOutput) ToPreferredAppGroupTypePtrOutput() PreferredAppGroupTypePtrOutput {
	return o
}

func (o PreferredAppGroupTypePtrOutput) ToPreferredAppGroupTypePtrOutputWithContext(ctx context.Context) PreferredAppGroupTypePtrOutput {
	return o
}

func (o PreferredAppGroupTypePtrOutput) Elem() PreferredAppGroupTypeOutput {
	return o.ApplyT(func(v *PreferredAppGroupType) PreferredAppGroupType {
		if v != nil {
			return *v
		}
		var ret PreferredAppGroupType
		return ret
	}).(PreferredAppGroupTypeOutput)
}

func (o PreferredAppGroupTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PreferredAppGroupTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PreferredAppGroupType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PreferredAppGroupTypeInput interface {
	pulumi.Input

	ToPreferredAppGroupTypeOutput() PreferredAppGroupTypeOutput
	ToPreferredAppGroupTypeOutputWithContext(context.Context) PreferredAppGroupTypeOutput
}

var preferredAppGroupTypePtrType = reflect.TypeOf((**PreferredAppGroupType)(nil)).Elem()

type PreferredAppGroupTypePtrInput interface {
	pulumi.Input

	ToPreferredAppGroupTypePtrOutput() PreferredAppGroupTypePtrOutput
	ToPreferredAppGroupTypePtrOutputWithContext(context.Context) PreferredAppGroupTypePtrOutput
}

type preferredAppGroupTypePtr string

func PreferredAppGroupTypePtr(v string) PreferredAppGroupTypePtrInput {
	return (*preferredAppGroupTypePtr)(&v)
}

func (*preferredAppGroupTypePtr) ElementType() reflect.Type {
	return preferredAppGroupTypePtrType
}

func (in *preferredAppGroupTypePtr) ToPreferredAppGroupTypePtrOutput() PreferredAppGroupTypePtrOutput {
	return pulumi.ToOutput(in).(PreferredAppGroupTypePtrOutput)
}

func (in *preferredAppGroupTypePtr) ToPreferredAppGroupTypePtrOutputWithContext(ctx context.Context) PreferredAppGroupTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PreferredAppGroupTypePtrOutput)
}

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusPending  = PrivateEndpointServiceConnectionStatus("Pending")
	PrivateEndpointServiceConnectionStatusApproved = PrivateEndpointServiceConnectionStatus("Approved")
	PrivateEndpointServiceConnectionStatusRejected = PrivateEndpointServiceConnectionStatus("Rejected")
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

type RegistrationTokenOperation string

const (
	RegistrationTokenOperationDelete = RegistrationTokenOperation("Delete")
	RegistrationTokenOperationNone   = RegistrationTokenOperation("None")
	RegistrationTokenOperationUpdate = RegistrationTokenOperation("Update")
)

func (RegistrationTokenOperation) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationTokenOperation)(nil)).Elem()
}

func (e RegistrationTokenOperation) ToRegistrationTokenOperationOutput() RegistrationTokenOperationOutput {
	return pulumi.ToOutput(e).(RegistrationTokenOperationOutput)
}

func (e RegistrationTokenOperation) ToRegistrationTokenOperationOutputWithContext(ctx context.Context) RegistrationTokenOperationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RegistrationTokenOperationOutput)
}

func (e RegistrationTokenOperation) ToRegistrationTokenOperationPtrOutput() RegistrationTokenOperationPtrOutput {
	return e.ToRegistrationTokenOperationPtrOutputWithContext(context.Background())
}

func (e RegistrationTokenOperation) ToRegistrationTokenOperationPtrOutputWithContext(ctx context.Context) RegistrationTokenOperationPtrOutput {
	return RegistrationTokenOperation(e).ToRegistrationTokenOperationOutputWithContext(ctx).ToRegistrationTokenOperationPtrOutputWithContext(ctx)
}

func (e RegistrationTokenOperation) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RegistrationTokenOperation) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RegistrationTokenOperation) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RegistrationTokenOperation) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RegistrationTokenOperationOutput struct{ *pulumi.OutputState }

func (RegistrationTokenOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationTokenOperation)(nil)).Elem()
}

func (o RegistrationTokenOperationOutput) ToRegistrationTokenOperationOutput() RegistrationTokenOperationOutput {
	return o
}

func (o RegistrationTokenOperationOutput) ToRegistrationTokenOperationOutputWithContext(ctx context.Context) RegistrationTokenOperationOutput {
	return o
}

func (o RegistrationTokenOperationOutput) ToRegistrationTokenOperationPtrOutput() RegistrationTokenOperationPtrOutput {
	return o.ToRegistrationTokenOperationPtrOutputWithContext(context.Background())
}

func (o RegistrationTokenOperationOutput) ToRegistrationTokenOperationPtrOutputWithContext(ctx context.Context) RegistrationTokenOperationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RegistrationTokenOperation) *RegistrationTokenOperation {
		return &v
	}).(RegistrationTokenOperationPtrOutput)
}

func (o RegistrationTokenOperationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RegistrationTokenOperationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RegistrationTokenOperation) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RegistrationTokenOperationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RegistrationTokenOperationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RegistrationTokenOperation) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RegistrationTokenOperationPtrOutput struct{ *pulumi.OutputState }

func (RegistrationTokenOperationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationTokenOperation)(nil)).Elem()
}

func (o RegistrationTokenOperationPtrOutput) ToRegistrationTokenOperationPtrOutput() RegistrationTokenOperationPtrOutput {
	return o
}

func (o RegistrationTokenOperationPtrOutput) ToRegistrationTokenOperationPtrOutputWithContext(ctx context.Context) RegistrationTokenOperationPtrOutput {
	return o
}

func (o RegistrationTokenOperationPtrOutput) Elem() RegistrationTokenOperationOutput {
	return o.ApplyT(func(v *RegistrationTokenOperation) RegistrationTokenOperation {
		if v != nil {
			return *v
		}
		var ret RegistrationTokenOperation
		return ret
	}).(RegistrationTokenOperationOutput)
}

func (o RegistrationTokenOperationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RegistrationTokenOperationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RegistrationTokenOperation) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RegistrationTokenOperationInput interface {
	pulumi.Input

	ToRegistrationTokenOperationOutput() RegistrationTokenOperationOutput
	ToRegistrationTokenOperationOutputWithContext(context.Context) RegistrationTokenOperationOutput
}

var registrationTokenOperationPtrType = reflect.TypeOf((**RegistrationTokenOperation)(nil)).Elem()

type RegistrationTokenOperationPtrInput interface {
	pulumi.Input

	ToRegistrationTokenOperationPtrOutput() RegistrationTokenOperationPtrOutput
	ToRegistrationTokenOperationPtrOutputWithContext(context.Context) RegistrationTokenOperationPtrOutput
}

type registrationTokenOperationPtr string

func RegistrationTokenOperationPtr(v string) RegistrationTokenOperationPtrInput {
	return (*registrationTokenOperationPtr)(&v)
}

func (*registrationTokenOperationPtr) ElementType() reflect.Type {
	return registrationTokenOperationPtrType
}

func (in *registrationTokenOperationPtr) ToRegistrationTokenOperationPtrOutput() RegistrationTokenOperationPtrOutput {
	return pulumi.ToOutput(in).(RegistrationTokenOperationPtrOutput)
}

func (in *registrationTokenOperationPtr) ToRegistrationTokenOperationPtrOutputWithContext(ctx context.Context) RegistrationTokenOperationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RegistrationTokenOperationPtrOutput)
}

type RemoteApplicationType string

const (
	RemoteApplicationTypeInBuilt         = RemoteApplicationType("InBuilt")
	RemoteApplicationTypeMsixApplication = RemoteApplicationType("MsixApplication")
)

func (RemoteApplicationType) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoteApplicationType)(nil)).Elem()
}

func (e RemoteApplicationType) ToRemoteApplicationTypeOutput() RemoteApplicationTypeOutput {
	return pulumi.ToOutput(e).(RemoteApplicationTypeOutput)
}

func (e RemoteApplicationType) ToRemoteApplicationTypeOutputWithContext(ctx context.Context) RemoteApplicationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RemoteApplicationTypeOutput)
}

func (e RemoteApplicationType) ToRemoteApplicationTypePtrOutput() RemoteApplicationTypePtrOutput {
	return e.ToRemoteApplicationTypePtrOutputWithContext(context.Background())
}

func (e RemoteApplicationType) ToRemoteApplicationTypePtrOutputWithContext(ctx context.Context) RemoteApplicationTypePtrOutput {
	return RemoteApplicationType(e).ToRemoteApplicationTypeOutputWithContext(ctx).ToRemoteApplicationTypePtrOutputWithContext(ctx)
}

func (e RemoteApplicationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RemoteApplicationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RemoteApplicationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RemoteApplicationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RemoteApplicationTypeOutput struct{ *pulumi.OutputState }

func (RemoteApplicationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoteApplicationType)(nil)).Elem()
}

func (o RemoteApplicationTypeOutput) ToRemoteApplicationTypeOutput() RemoteApplicationTypeOutput {
	return o
}

func (o RemoteApplicationTypeOutput) ToRemoteApplicationTypeOutputWithContext(ctx context.Context) RemoteApplicationTypeOutput {
	return o
}

func (o RemoteApplicationTypeOutput) ToRemoteApplicationTypePtrOutput() RemoteApplicationTypePtrOutput {
	return o.ToRemoteApplicationTypePtrOutputWithContext(context.Background())
}

func (o RemoteApplicationTypeOutput) ToRemoteApplicationTypePtrOutputWithContext(ctx context.Context) RemoteApplicationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RemoteApplicationType) *RemoteApplicationType {
		return &v
	}).(RemoteApplicationTypePtrOutput)
}

func (o RemoteApplicationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RemoteApplicationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RemoteApplicationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RemoteApplicationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RemoteApplicationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RemoteApplicationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RemoteApplicationTypePtrOutput struct{ *pulumi.OutputState }

func (RemoteApplicationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemoteApplicationType)(nil)).Elem()
}

func (o RemoteApplicationTypePtrOutput) ToRemoteApplicationTypePtrOutput() RemoteApplicationTypePtrOutput {
	return o
}

func (o RemoteApplicationTypePtrOutput) ToRemoteApplicationTypePtrOutputWithContext(ctx context.Context) RemoteApplicationTypePtrOutput {
	return o
}

func (o RemoteApplicationTypePtrOutput) Elem() RemoteApplicationTypeOutput {
	return o.ApplyT(func(v *RemoteApplicationType) RemoteApplicationType {
		if v != nil {
			return *v
		}
		var ret RemoteApplicationType
		return ret
	}).(RemoteApplicationTypeOutput)
}

func (o RemoteApplicationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RemoteApplicationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RemoteApplicationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RemoteApplicationTypeInput interface {
	pulumi.Input

	ToRemoteApplicationTypeOutput() RemoteApplicationTypeOutput
	ToRemoteApplicationTypeOutputWithContext(context.Context) RemoteApplicationTypeOutput
}

var remoteApplicationTypePtrType = reflect.TypeOf((**RemoteApplicationType)(nil)).Elem()

type RemoteApplicationTypePtrInput interface {
	pulumi.Input

	ToRemoteApplicationTypePtrOutput() RemoteApplicationTypePtrOutput
	ToRemoteApplicationTypePtrOutputWithContext(context.Context) RemoteApplicationTypePtrOutput
}

type remoteApplicationTypePtr string

func RemoteApplicationTypePtr(v string) RemoteApplicationTypePtrInput {
	return (*remoteApplicationTypePtr)(&v)
}

func (*remoteApplicationTypePtr) ElementType() reflect.Type {
	return remoteApplicationTypePtrType
}

func (in *remoteApplicationTypePtr) ToRemoteApplicationTypePtrOutput() RemoteApplicationTypePtrOutput {
	return pulumi.ToOutput(in).(RemoteApplicationTypePtrOutput)
}

func (in *remoteApplicationTypePtr) ToRemoteApplicationTypePtrOutputWithContext(ctx context.Context) RemoteApplicationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RemoteApplicationTypePtrOutput)
}

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
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

type SSOSecretType string

const (
	SSOSecretTypeSharedKey             = SSOSecretType("SharedKey")
	SSOSecretTypeCertificate           = SSOSecretType("Certificate")
	SSOSecretTypeSharedKeyInKeyVault   = SSOSecretType("SharedKeyInKeyVault")
	SSOSecretTypeCertificateInKeyVault = SSOSecretType("CertificateInKeyVault")
)

func (SSOSecretType) ElementType() reflect.Type {
	return reflect.TypeOf((*SSOSecretType)(nil)).Elem()
}

func (e SSOSecretType) ToSSOSecretTypeOutput() SSOSecretTypeOutput {
	return pulumi.ToOutput(e).(SSOSecretTypeOutput)
}

func (e SSOSecretType) ToSSOSecretTypeOutputWithContext(ctx context.Context) SSOSecretTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SSOSecretTypeOutput)
}

func (e SSOSecretType) ToSSOSecretTypePtrOutput() SSOSecretTypePtrOutput {
	return e.ToSSOSecretTypePtrOutputWithContext(context.Background())
}

func (e SSOSecretType) ToSSOSecretTypePtrOutputWithContext(ctx context.Context) SSOSecretTypePtrOutput {
	return SSOSecretType(e).ToSSOSecretTypeOutputWithContext(ctx).ToSSOSecretTypePtrOutputWithContext(ctx)
}

func (e SSOSecretType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SSOSecretType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SSOSecretType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SSOSecretType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SSOSecretTypeOutput struct{ *pulumi.OutputState }

func (SSOSecretTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SSOSecretType)(nil)).Elem()
}

func (o SSOSecretTypeOutput) ToSSOSecretTypeOutput() SSOSecretTypeOutput {
	return o
}

func (o SSOSecretTypeOutput) ToSSOSecretTypeOutputWithContext(ctx context.Context) SSOSecretTypeOutput {
	return o
}

func (o SSOSecretTypeOutput) ToSSOSecretTypePtrOutput() SSOSecretTypePtrOutput {
	return o.ToSSOSecretTypePtrOutputWithContext(context.Background())
}

func (o SSOSecretTypeOutput) ToSSOSecretTypePtrOutputWithContext(ctx context.Context) SSOSecretTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SSOSecretType) *SSOSecretType {
		return &v
	}).(SSOSecretTypePtrOutput)
}

func (o SSOSecretTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SSOSecretTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SSOSecretType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SSOSecretTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SSOSecretTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SSOSecretType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SSOSecretTypePtrOutput struct{ *pulumi.OutputState }

func (SSOSecretTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SSOSecretType)(nil)).Elem()
}

func (o SSOSecretTypePtrOutput) ToSSOSecretTypePtrOutput() SSOSecretTypePtrOutput {
	return o
}

func (o SSOSecretTypePtrOutput) ToSSOSecretTypePtrOutputWithContext(ctx context.Context) SSOSecretTypePtrOutput {
	return o
}

func (o SSOSecretTypePtrOutput) Elem() SSOSecretTypeOutput {
	return o.ApplyT(func(v *SSOSecretType) SSOSecretType {
		if v != nil {
			return *v
		}
		var ret SSOSecretType
		return ret
	}).(SSOSecretTypeOutput)
}

func (o SSOSecretTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SSOSecretTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SSOSecretType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SSOSecretTypeInput interface {
	pulumi.Input

	ToSSOSecretTypeOutput() SSOSecretTypeOutput
	ToSSOSecretTypeOutputWithContext(context.Context) SSOSecretTypeOutput
}

var ssosecretTypePtrType = reflect.TypeOf((**SSOSecretType)(nil)).Elem()

type SSOSecretTypePtrInput interface {
	pulumi.Input

	ToSSOSecretTypePtrOutput() SSOSecretTypePtrOutput
	ToSSOSecretTypePtrOutputWithContext(context.Context) SSOSecretTypePtrOutput
}

type ssosecretTypePtr string

func SSOSecretTypePtr(v string) SSOSecretTypePtrInput {
	return (*ssosecretTypePtr)(&v)
}

func (*ssosecretTypePtr) ElementType() reflect.Type {
	return ssosecretTypePtrType
}

func (in *ssosecretTypePtr) ToSSOSecretTypePtrOutput() SSOSecretTypePtrOutput {
	return pulumi.ToOutput(in).(SSOSecretTypePtrOutput)
}

func (in *ssosecretTypePtr) ToSSOSecretTypePtrOutputWithContext(ctx context.Context) SSOSecretTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SSOSecretTypePtrOutput)
}

type SessionHostLoadBalancingAlgorithm string

const (
	SessionHostLoadBalancingAlgorithmBreadthFirst = SessionHostLoadBalancingAlgorithm("BreadthFirst")
	SessionHostLoadBalancingAlgorithmDepthFirst   = SessionHostLoadBalancingAlgorithm("DepthFirst")
)

func (SessionHostLoadBalancingAlgorithm) ElementType() reflect.Type {
	return reflect.TypeOf((*SessionHostLoadBalancingAlgorithm)(nil)).Elem()
}

func (e SessionHostLoadBalancingAlgorithm) ToSessionHostLoadBalancingAlgorithmOutput() SessionHostLoadBalancingAlgorithmOutput {
	return pulumi.ToOutput(e).(SessionHostLoadBalancingAlgorithmOutput)
}

func (e SessionHostLoadBalancingAlgorithm) ToSessionHostLoadBalancingAlgorithmOutputWithContext(ctx context.Context) SessionHostLoadBalancingAlgorithmOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SessionHostLoadBalancingAlgorithmOutput)
}

func (e SessionHostLoadBalancingAlgorithm) ToSessionHostLoadBalancingAlgorithmPtrOutput() SessionHostLoadBalancingAlgorithmPtrOutput {
	return e.ToSessionHostLoadBalancingAlgorithmPtrOutputWithContext(context.Background())
}

func (e SessionHostLoadBalancingAlgorithm) ToSessionHostLoadBalancingAlgorithmPtrOutputWithContext(ctx context.Context) SessionHostLoadBalancingAlgorithmPtrOutput {
	return SessionHostLoadBalancingAlgorithm(e).ToSessionHostLoadBalancingAlgorithmOutputWithContext(ctx).ToSessionHostLoadBalancingAlgorithmPtrOutputWithContext(ctx)
}

func (e SessionHostLoadBalancingAlgorithm) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SessionHostLoadBalancingAlgorithm) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SessionHostLoadBalancingAlgorithm) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SessionHostLoadBalancingAlgorithm) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SessionHostLoadBalancingAlgorithmOutput struct{ *pulumi.OutputState }

func (SessionHostLoadBalancingAlgorithmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SessionHostLoadBalancingAlgorithm)(nil)).Elem()
}

func (o SessionHostLoadBalancingAlgorithmOutput) ToSessionHostLoadBalancingAlgorithmOutput() SessionHostLoadBalancingAlgorithmOutput {
	return o
}

func (o SessionHostLoadBalancingAlgorithmOutput) ToSessionHostLoadBalancingAlgorithmOutputWithContext(ctx context.Context) SessionHostLoadBalancingAlgorithmOutput {
	return o
}

func (o SessionHostLoadBalancingAlgorithmOutput) ToSessionHostLoadBalancingAlgorithmPtrOutput() SessionHostLoadBalancingAlgorithmPtrOutput {
	return o.ToSessionHostLoadBalancingAlgorithmPtrOutputWithContext(context.Background())
}

func (o SessionHostLoadBalancingAlgorithmOutput) ToSessionHostLoadBalancingAlgorithmPtrOutputWithContext(ctx context.Context) SessionHostLoadBalancingAlgorithmPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SessionHostLoadBalancingAlgorithm) *SessionHostLoadBalancingAlgorithm {
		return &v
	}).(SessionHostLoadBalancingAlgorithmPtrOutput)
}

func (o SessionHostLoadBalancingAlgorithmOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SessionHostLoadBalancingAlgorithmOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SessionHostLoadBalancingAlgorithm) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SessionHostLoadBalancingAlgorithmOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SessionHostLoadBalancingAlgorithmOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SessionHostLoadBalancingAlgorithm) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SessionHostLoadBalancingAlgorithmPtrOutput struct{ *pulumi.OutputState }

func (SessionHostLoadBalancingAlgorithmPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SessionHostLoadBalancingAlgorithm)(nil)).Elem()
}

func (o SessionHostLoadBalancingAlgorithmPtrOutput) ToSessionHostLoadBalancingAlgorithmPtrOutput() SessionHostLoadBalancingAlgorithmPtrOutput {
	return o
}

func (o SessionHostLoadBalancingAlgorithmPtrOutput) ToSessionHostLoadBalancingAlgorithmPtrOutputWithContext(ctx context.Context) SessionHostLoadBalancingAlgorithmPtrOutput {
	return o
}

func (o SessionHostLoadBalancingAlgorithmPtrOutput) Elem() SessionHostLoadBalancingAlgorithmOutput {
	return o.ApplyT(func(v *SessionHostLoadBalancingAlgorithm) SessionHostLoadBalancingAlgorithm {
		if v != nil {
			return *v
		}
		var ret SessionHostLoadBalancingAlgorithm
		return ret
	}).(SessionHostLoadBalancingAlgorithmOutput)
}

func (o SessionHostLoadBalancingAlgorithmPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SessionHostLoadBalancingAlgorithmPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SessionHostLoadBalancingAlgorithm) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SessionHostLoadBalancingAlgorithmInput interface {
	pulumi.Input

	ToSessionHostLoadBalancingAlgorithmOutput() SessionHostLoadBalancingAlgorithmOutput
	ToSessionHostLoadBalancingAlgorithmOutputWithContext(context.Context) SessionHostLoadBalancingAlgorithmOutput
}

var sessionHostLoadBalancingAlgorithmPtrType = reflect.TypeOf((**SessionHostLoadBalancingAlgorithm)(nil)).Elem()

type SessionHostLoadBalancingAlgorithmPtrInput interface {
	pulumi.Input

	ToSessionHostLoadBalancingAlgorithmPtrOutput() SessionHostLoadBalancingAlgorithmPtrOutput
	ToSessionHostLoadBalancingAlgorithmPtrOutputWithContext(context.Context) SessionHostLoadBalancingAlgorithmPtrOutput
}

type sessionHostLoadBalancingAlgorithmPtr string

func SessionHostLoadBalancingAlgorithmPtr(v string) SessionHostLoadBalancingAlgorithmPtrInput {
	return (*sessionHostLoadBalancingAlgorithmPtr)(&v)
}

func (*sessionHostLoadBalancingAlgorithmPtr) ElementType() reflect.Type {
	return sessionHostLoadBalancingAlgorithmPtrType
}

func (in *sessionHostLoadBalancingAlgorithmPtr) ToSessionHostLoadBalancingAlgorithmPtrOutput() SessionHostLoadBalancingAlgorithmPtrOutput {
	return pulumi.ToOutput(in).(SessionHostLoadBalancingAlgorithmPtrOutput)
}

func (in *sessionHostLoadBalancingAlgorithmPtr) ToSessionHostLoadBalancingAlgorithmPtrOutputWithContext(ctx context.Context) SessionHostLoadBalancingAlgorithmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SessionHostLoadBalancingAlgorithmPtrOutput)
}

type SkuTier string

const (
	SkuTierFree     = SkuTier("Free")
	SkuTierBasic    = SkuTier("Basic")
	SkuTierStandard = SkuTier("Standard")
	SkuTierPremium  = SkuTier("Premium")
)

func (SkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuTier)(nil)).Elem()
}

func (e SkuTier) ToSkuTierOutput() SkuTierOutput {
	return pulumi.ToOutput(e).(SkuTierOutput)
}

func (e SkuTier) ToSkuTierOutputWithContext(ctx context.Context) SkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuTierOutput)
}

func (e SkuTier) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return e.ToSkuTierPtrOutputWithContext(context.Background())
}

func (e SkuTier) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return SkuTier(e).ToSkuTierOutputWithContext(ctx).ToSkuTierPtrOutputWithContext(ctx)
}

func (e SkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuTierOutput struct{ *pulumi.OutputState }

func (SkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuTier)(nil)).Elem()
}

func (o SkuTierOutput) ToSkuTierOutput() SkuTierOutput {
	return o
}

func (o SkuTierOutput) ToSkuTierOutputWithContext(ctx context.Context) SkuTierOutput {
	return o
}

func (o SkuTierOutput) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return o.ToSkuTierPtrOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuTier) *SkuTier {
		return &v
	}).(SkuTierPtrOutput)
}

func (o SkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuTierPtrOutput struct{ *pulumi.OutputState }

func (SkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuTier)(nil)).Elem()
}

func (o SkuTierPtrOutput) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return o
}

func (o SkuTierPtrOutput) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return o
}

func (o SkuTierPtrOutput) Elem() SkuTierOutput {
	return o.ApplyT(func(v *SkuTier) SkuTier {
		if v != nil {
			return *v
		}
		var ret SkuTier
		return ret
	}).(SkuTierOutput)
}

func (o SkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SkuTierInput interface {
	pulumi.Input

	ToSkuTierOutput() SkuTierOutput
	ToSkuTierOutputWithContext(context.Context) SkuTierOutput
}

var skuTierPtrType = reflect.TypeOf((**SkuTier)(nil)).Elem()

type SkuTierPtrInput interface {
	pulumi.Input

	ToSkuTierPtrOutput() SkuTierPtrOutput
	ToSkuTierPtrOutputWithContext(context.Context) SkuTierPtrOutput
}

type skuTierPtr string

func SkuTierPtr(v string) SkuTierPtrInput {
	return (*skuTierPtr)(&v)
}

func (*skuTierPtr) ElementType() reflect.Type {
	return skuTierPtrType
}

func (in *skuTierPtr) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return pulumi.ToOutput(in).(SkuTierPtrOutput)
}

func (in *skuTierPtr) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuTierPtrOutput)
}

type StopHostsWhen string

const (
	StopHostsWhenZeroSessions       = StopHostsWhen("ZeroSessions")
	StopHostsWhenZeroActiveSessions = StopHostsWhen("ZeroActiveSessions")
)

func (StopHostsWhen) ElementType() reflect.Type {
	return reflect.TypeOf((*StopHostsWhen)(nil)).Elem()
}

func (e StopHostsWhen) ToStopHostsWhenOutput() StopHostsWhenOutput {
	return pulumi.ToOutput(e).(StopHostsWhenOutput)
}

func (e StopHostsWhen) ToStopHostsWhenOutputWithContext(ctx context.Context) StopHostsWhenOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StopHostsWhenOutput)
}

func (e StopHostsWhen) ToStopHostsWhenPtrOutput() StopHostsWhenPtrOutput {
	return e.ToStopHostsWhenPtrOutputWithContext(context.Background())
}

func (e StopHostsWhen) ToStopHostsWhenPtrOutputWithContext(ctx context.Context) StopHostsWhenPtrOutput {
	return StopHostsWhen(e).ToStopHostsWhenOutputWithContext(ctx).ToStopHostsWhenPtrOutputWithContext(ctx)
}

func (e StopHostsWhen) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StopHostsWhen) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StopHostsWhen) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StopHostsWhen) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StopHostsWhenOutput struct{ *pulumi.OutputState }

func (StopHostsWhenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StopHostsWhen)(nil)).Elem()
}

func (o StopHostsWhenOutput) ToStopHostsWhenOutput() StopHostsWhenOutput {
	return o
}

func (o StopHostsWhenOutput) ToStopHostsWhenOutputWithContext(ctx context.Context) StopHostsWhenOutput {
	return o
}

func (o StopHostsWhenOutput) ToStopHostsWhenPtrOutput() StopHostsWhenPtrOutput {
	return o.ToStopHostsWhenPtrOutputWithContext(context.Background())
}

func (o StopHostsWhenOutput) ToStopHostsWhenPtrOutputWithContext(ctx context.Context) StopHostsWhenPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StopHostsWhen) *StopHostsWhen {
		return &v
	}).(StopHostsWhenPtrOutput)
}

func (o StopHostsWhenOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StopHostsWhenOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StopHostsWhen) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StopHostsWhenOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StopHostsWhenOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StopHostsWhen) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StopHostsWhenPtrOutput struct{ *pulumi.OutputState }

func (StopHostsWhenPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StopHostsWhen)(nil)).Elem()
}

func (o StopHostsWhenPtrOutput) ToStopHostsWhenPtrOutput() StopHostsWhenPtrOutput {
	return o
}

func (o StopHostsWhenPtrOutput) ToStopHostsWhenPtrOutputWithContext(ctx context.Context) StopHostsWhenPtrOutput {
	return o
}

func (o StopHostsWhenPtrOutput) Elem() StopHostsWhenOutput {
	return o.ApplyT(func(v *StopHostsWhen) StopHostsWhen {
		if v != nil {
			return *v
		}
		var ret StopHostsWhen
		return ret
	}).(StopHostsWhenOutput)
}

func (o StopHostsWhenPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StopHostsWhenPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StopHostsWhen) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StopHostsWhenInput interface {
	pulumi.Input

	ToStopHostsWhenOutput() StopHostsWhenOutput
	ToStopHostsWhenOutputWithContext(context.Context) StopHostsWhenOutput
}

var stopHostsWhenPtrType = reflect.TypeOf((**StopHostsWhen)(nil)).Elem()

type StopHostsWhenPtrInput interface {
	pulumi.Input

	ToStopHostsWhenPtrOutput() StopHostsWhenPtrOutput
	ToStopHostsWhenPtrOutputWithContext(context.Context) StopHostsWhenPtrOutput
}

type stopHostsWhenPtr string

func StopHostsWhenPtr(v string) StopHostsWhenPtrInput {
	return (*stopHostsWhenPtr)(&v)
}

func (*stopHostsWhenPtr) ElementType() reflect.Type {
	return stopHostsWhenPtrType
}

func (in *stopHostsWhenPtr) ToStopHostsWhenPtrOutput() StopHostsWhenPtrOutput {
	return pulumi.ToOutput(in).(StopHostsWhenPtrOutput)
}

func (in *stopHostsWhenPtr) ToStopHostsWhenPtrOutputWithContext(ctx context.Context) StopHostsWhenPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StopHostsWhenPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationGroupTypeOutput{})
	pulumi.RegisterOutputType(ApplicationGroupTypePtrOutput{})
	pulumi.RegisterOutputType(CommandLineSettingOutput{})
	pulumi.RegisterOutputType(CommandLineSettingPtrOutput{})
	pulumi.RegisterOutputType(HostPoolTypeOutput{})
	pulumi.RegisterOutputType(HostPoolTypePtrOutput{})
	pulumi.RegisterOutputType(LoadBalancerTypeOutput{})
	pulumi.RegisterOutputType(LoadBalancerTypePtrOutput{})
	pulumi.RegisterOutputType(OperationOutput{})
	pulumi.RegisterOutputType(OperationPtrOutput{})
	pulumi.RegisterOutputType(PersonalDesktopAssignmentTypeOutput{})
	pulumi.RegisterOutputType(PersonalDesktopAssignmentTypePtrOutput{})
	pulumi.RegisterOutputType(PreferredAppGroupTypeOutput{})
	pulumi.RegisterOutputType(PreferredAppGroupTypePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateEndpointServiceConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(RegistrationTokenOperationOutput{})
	pulumi.RegisterOutputType(RegistrationTokenOperationPtrOutput{})
	pulumi.RegisterOutputType(RemoteApplicationTypeOutput{})
	pulumi.RegisterOutputType(RemoteApplicationTypePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(SSOSecretTypeOutput{})
	pulumi.RegisterOutputType(SSOSecretTypePtrOutput{})
	pulumi.RegisterOutputType(SessionHostLoadBalancingAlgorithmOutput{})
	pulumi.RegisterOutputType(SessionHostLoadBalancingAlgorithmPtrOutput{})
	pulumi.RegisterOutputType(SkuTierOutput{})
	pulumi.RegisterOutputType(SkuTierPtrOutput{})
	pulumi.RegisterOutputType(StopHostsWhenOutput{})
	pulumi.RegisterOutputType(StopHostsWhenPtrOutput{})
}
