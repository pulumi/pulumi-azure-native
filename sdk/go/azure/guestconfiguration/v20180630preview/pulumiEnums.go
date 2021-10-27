


package v20180630preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActionAfterReboot string

const (
	ActionAfterRebootContinueConfiguration = ActionAfterReboot("ContinueConfiguration")
	ActionAfterRebootStopConfiguration     = ActionAfterReboot("StopConfiguration")
)

func (ActionAfterReboot) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionAfterReboot)(nil)).Elem()
}

func (e ActionAfterReboot) ToActionAfterRebootOutput() ActionAfterRebootOutput {
	return pulumi.ToOutput(e).(ActionAfterRebootOutput)
}

func (e ActionAfterReboot) ToActionAfterRebootOutputWithContext(ctx context.Context) ActionAfterRebootOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ActionAfterRebootOutput)
}

func (e ActionAfterReboot) ToActionAfterRebootPtrOutput() ActionAfterRebootPtrOutput {
	return e.ToActionAfterRebootPtrOutputWithContext(context.Background())
}

func (e ActionAfterReboot) ToActionAfterRebootPtrOutputWithContext(ctx context.Context) ActionAfterRebootPtrOutput {
	return ActionAfterReboot(e).ToActionAfterRebootOutputWithContext(ctx).ToActionAfterRebootPtrOutputWithContext(ctx)
}

func (e ActionAfterReboot) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ActionAfterReboot) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ActionAfterReboot) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ActionAfterReboot) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ActionAfterRebootOutput struct{ *pulumi.OutputState }

func (ActionAfterRebootOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionAfterReboot)(nil)).Elem()
}

func (o ActionAfterRebootOutput) ToActionAfterRebootOutput() ActionAfterRebootOutput {
	return o
}

func (o ActionAfterRebootOutput) ToActionAfterRebootOutputWithContext(ctx context.Context) ActionAfterRebootOutput {
	return o
}

func (o ActionAfterRebootOutput) ToActionAfterRebootPtrOutput() ActionAfterRebootPtrOutput {
	return o.ToActionAfterRebootPtrOutputWithContext(context.Background())
}

func (o ActionAfterRebootOutput) ToActionAfterRebootPtrOutputWithContext(ctx context.Context) ActionAfterRebootPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActionAfterReboot) *ActionAfterReboot {
		return &v
	}).(ActionAfterRebootPtrOutput)
}

func (o ActionAfterRebootOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ActionAfterRebootOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ActionAfterReboot) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ActionAfterRebootOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActionAfterRebootOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ActionAfterReboot) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ActionAfterRebootPtrOutput struct{ *pulumi.OutputState }

func (ActionAfterRebootPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionAfterReboot)(nil)).Elem()
}

func (o ActionAfterRebootPtrOutput) ToActionAfterRebootPtrOutput() ActionAfterRebootPtrOutput {
	return o
}

func (o ActionAfterRebootPtrOutput) ToActionAfterRebootPtrOutputWithContext(ctx context.Context) ActionAfterRebootPtrOutput {
	return o
}

func (o ActionAfterRebootPtrOutput) Elem() ActionAfterRebootOutput {
	return o.ApplyT(func(v *ActionAfterReboot) ActionAfterReboot {
		if v != nil {
			return *v
		}
		var ret ActionAfterReboot
		return ret
	}).(ActionAfterRebootOutput)
}

func (o ActionAfterRebootPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActionAfterRebootPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ActionAfterReboot) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ActionAfterRebootInput interface {
	pulumi.Input

	ToActionAfterRebootOutput() ActionAfterRebootOutput
	ToActionAfterRebootOutputWithContext(context.Context) ActionAfterRebootOutput
}

var actionAfterRebootPtrType = reflect.TypeOf((**ActionAfterReboot)(nil)).Elem()

type ActionAfterRebootPtrInput interface {
	pulumi.Input

	ToActionAfterRebootPtrOutput() ActionAfterRebootPtrOutput
	ToActionAfterRebootPtrOutputWithContext(context.Context) ActionAfterRebootPtrOutput
}

type actionAfterRebootPtr string

func ActionAfterRebootPtr(v string) ActionAfterRebootPtrInput {
	return (*actionAfterRebootPtr)(&v)
}

func (*actionAfterRebootPtr) ElementType() reflect.Type {
	return actionAfterRebootPtrType
}

func (in *actionAfterRebootPtr) ToActionAfterRebootPtrOutput() ActionAfterRebootPtrOutput {
	return pulumi.ToOutput(in).(ActionAfterRebootPtrOutput)
}

func (in *actionAfterRebootPtr) ToActionAfterRebootPtrOutputWithContext(ctx context.Context) ActionAfterRebootPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ActionAfterRebootPtrOutput)
}

type ConfigurationMode string

const (
	ConfigurationModeApplyOnly           = ConfigurationMode("ApplyOnly")
	ConfigurationModeApplyAndMonitor     = ConfigurationMode("ApplyAndMonitor")
	ConfigurationModeApplyAndAutoCorrect = ConfigurationMode("ApplyAndAutoCorrect")
)

func (ConfigurationMode) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationMode)(nil)).Elem()
}

func (e ConfigurationMode) ToConfigurationModeOutput() ConfigurationModeOutput {
	return pulumi.ToOutput(e).(ConfigurationModeOutput)
}

func (e ConfigurationMode) ToConfigurationModeOutputWithContext(ctx context.Context) ConfigurationModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConfigurationModeOutput)
}

func (e ConfigurationMode) ToConfigurationModePtrOutput() ConfigurationModePtrOutput {
	return e.ToConfigurationModePtrOutputWithContext(context.Background())
}

func (e ConfigurationMode) ToConfigurationModePtrOutputWithContext(ctx context.Context) ConfigurationModePtrOutput {
	return ConfigurationMode(e).ToConfigurationModeOutputWithContext(ctx).ToConfigurationModePtrOutputWithContext(ctx)
}

func (e ConfigurationMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConfigurationMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConfigurationMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConfigurationMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConfigurationModeOutput struct{ *pulumi.OutputState }

func (ConfigurationModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationMode)(nil)).Elem()
}

func (o ConfigurationModeOutput) ToConfigurationModeOutput() ConfigurationModeOutput {
	return o
}

func (o ConfigurationModeOutput) ToConfigurationModeOutputWithContext(ctx context.Context) ConfigurationModeOutput {
	return o
}

func (o ConfigurationModeOutput) ToConfigurationModePtrOutput() ConfigurationModePtrOutput {
	return o.ToConfigurationModePtrOutputWithContext(context.Background())
}

func (o ConfigurationModeOutput) ToConfigurationModePtrOutputWithContext(ctx context.Context) ConfigurationModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationMode) *ConfigurationMode {
		return &v
	}).(ConfigurationModePtrOutput)
}

func (o ConfigurationModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConfigurationModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConfigurationMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConfigurationModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConfigurationModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConfigurationMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConfigurationModePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationMode)(nil)).Elem()
}

func (o ConfigurationModePtrOutput) ToConfigurationModePtrOutput() ConfigurationModePtrOutput {
	return o
}

func (o ConfigurationModePtrOutput) ToConfigurationModePtrOutputWithContext(ctx context.Context) ConfigurationModePtrOutput {
	return o
}

func (o ConfigurationModePtrOutput) Elem() ConfigurationModeOutput {
	return o.ApplyT(func(v *ConfigurationMode) ConfigurationMode {
		if v != nil {
			return *v
		}
		var ret ConfigurationMode
		return ret
	}).(ConfigurationModeOutput)
}

func (o ConfigurationModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConfigurationModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConfigurationMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ConfigurationModeInput interface {
	pulumi.Input

	ToConfigurationModeOutput() ConfigurationModeOutput
	ToConfigurationModeOutputWithContext(context.Context) ConfigurationModeOutput
}

var configurationModePtrType = reflect.TypeOf((**ConfigurationMode)(nil)).Elem()

type ConfigurationModePtrInput interface {
	pulumi.Input

	ToConfigurationModePtrOutput() ConfigurationModePtrOutput
	ToConfigurationModePtrOutputWithContext(context.Context) ConfigurationModePtrOutput
}

type configurationModePtr string

func ConfigurationModePtr(v string) ConfigurationModePtrInput {
	return (*configurationModePtr)(&v)
}

func (*configurationModePtr) ElementType() reflect.Type {
	return configurationModePtrType
}

func (in *configurationModePtr) ToConfigurationModePtrOutput() ConfigurationModePtrOutput {
	return pulumi.ToOutput(in).(ConfigurationModePtrOutput)
}

func (in *configurationModePtr) ToConfigurationModePtrOutputWithContext(ctx context.Context) ConfigurationModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConfigurationModePtrOutput)
}

type Kind string

const (
	KindDSC = Kind("DSC")
)

func (Kind) ElementType() reflect.Type {
	return reflect.TypeOf((*Kind)(nil)).Elem()
}

func (e Kind) ToKindOutput() KindOutput {
	return pulumi.ToOutput(e).(KindOutput)
}

func (e Kind) ToKindOutputWithContext(ctx context.Context) KindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KindOutput)
}

func (e Kind) ToKindPtrOutput() KindPtrOutput {
	return e.ToKindPtrOutputWithContext(context.Background())
}

func (e Kind) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return Kind(e).ToKindOutputWithContext(ctx).ToKindPtrOutputWithContext(ctx)
}

func (e Kind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Kind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Kind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Kind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KindOutput struct{ *pulumi.OutputState }

func (KindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Kind)(nil)).Elem()
}

func (o KindOutput) ToKindOutput() KindOutput {
	return o
}

func (o KindOutput) ToKindOutputWithContext(ctx context.Context) KindOutput {
	return o
}

func (o KindOutput) ToKindPtrOutput() KindPtrOutput {
	return o.ToKindPtrOutputWithContext(context.Background())
}

func (o KindOutput) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Kind) *Kind {
		return &v
	}).(KindPtrOutput)
}

func (o KindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Kind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Kind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KindPtrOutput struct{ *pulumi.OutputState }

func (KindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Kind)(nil)).Elem()
}

func (o KindPtrOutput) ToKindPtrOutput() KindPtrOutput {
	return o
}

func (o KindPtrOutput) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return o
}

func (o KindPtrOutput) Elem() KindOutput {
	return o.ApplyT(func(v *Kind) Kind {
		if v != nil {
			return *v
		}
		var ret Kind
		return ret
	}).(KindOutput)
}

func (o KindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Kind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type KindInput interface {
	pulumi.Input

	ToKindOutput() KindOutput
	ToKindOutputWithContext(context.Context) KindOutput
}

var kindPtrType = reflect.TypeOf((**Kind)(nil)).Elem()

type KindPtrInput interface {
	pulumi.Input

	ToKindPtrOutput() KindPtrOutput
	ToKindPtrOutputWithContext(context.Context) KindPtrOutput
}

type kindPtr string

func KindPtr(v string) KindPtrInput {
	return (*kindPtr)(&v)
}

func (*kindPtr) ElementType() reflect.Type {
	return kindPtrType
}

func (in *kindPtr) ToKindPtrOutput() KindPtrOutput {
	return pulumi.ToOutput(in).(KindPtrOutput)
}

func (in *kindPtr) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KindPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActionAfterRebootInput)(nil)).Elem(), ActionAfterReboot("ContinueConfiguration"))
	pulumi.RegisterInputType(reflect.TypeOf((*ActionAfterRebootPtrInput)(nil)).Elem(), ActionAfterReboot("ContinueConfiguration"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationModeInput)(nil)).Elem(), ConfigurationMode("ApplyOnly"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationModePtrInput)(nil)).Elem(), ConfigurationMode("ApplyOnly"))
	pulumi.RegisterInputType(reflect.TypeOf((*KindInput)(nil)).Elem(), Kind("DSC"))
	pulumi.RegisterInputType(reflect.TypeOf((*KindPtrInput)(nil)).Elem(), Kind("DSC"))
	pulumi.RegisterOutputType(ActionAfterRebootOutput{})
	pulumi.RegisterOutputType(ActionAfterRebootPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationModeOutput{})
	pulumi.RegisterOutputType(ConfigurationModePtrOutput{})
	pulumi.RegisterOutputType(KindOutput{})
	pulumi.RegisterOutputType(KindPtrOutput{})
}
