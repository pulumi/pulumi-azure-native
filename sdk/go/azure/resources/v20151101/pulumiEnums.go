


package v20151101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeploymentModeInput)(nil)).Elem(), DeploymentMode("Incremental"))
	pulumi.RegisterInputType(reflect.TypeOf((*DeploymentModePtrInput)(nil)).Elem(), DeploymentMode("Incremental"))
	pulumi.RegisterOutputType(DeploymentModeOutput{})
	pulumi.RegisterOutputType(DeploymentModePtrOutput{})
}
