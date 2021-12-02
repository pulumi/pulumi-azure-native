


package v20211001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceDiscoveryMode string

const (
	// Remediate resources that are already known to be non-compliant.
	ResourceDiscoveryModeExistingNonCompliant = ResourceDiscoveryMode("ExistingNonCompliant")
	// Re-evaluate the compliance state of resources and then remediate the resources found to be non-compliant.
	ResourceDiscoveryModeReEvaluateCompliance = ResourceDiscoveryMode("ReEvaluateCompliance")
)

func (ResourceDiscoveryMode) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceDiscoveryMode)(nil)).Elem()
}

func (e ResourceDiscoveryMode) ToResourceDiscoveryModeOutput() ResourceDiscoveryModeOutput {
	return pulumi.ToOutput(e).(ResourceDiscoveryModeOutput)
}

func (e ResourceDiscoveryMode) ToResourceDiscoveryModeOutputWithContext(ctx context.Context) ResourceDiscoveryModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceDiscoveryModeOutput)
}

func (e ResourceDiscoveryMode) ToResourceDiscoveryModePtrOutput() ResourceDiscoveryModePtrOutput {
	return e.ToResourceDiscoveryModePtrOutputWithContext(context.Background())
}

func (e ResourceDiscoveryMode) ToResourceDiscoveryModePtrOutputWithContext(ctx context.Context) ResourceDiscoveryModePtrOutput {
	return ResourceDiscoveryMode(e).ToResourceDiscoveryModeOutputWithContext(ctx).ToResourceDiscoveryModePtrOutputWithContext(ctx)
}

func (e ResourceDiscoveryMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceDiscoveryMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceDiscoveryMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceDiscoveryMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceDiscoveryModeOutput struct{ *pulumi.OutputState }

func (ResourceDiscoveryModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceDiscoveryMode)(nil)).Elem()
}

func (o ResourceDiscoveryModeOutput) ToResourceDiscoveryModeOutput() ResourceDiscoveryModeOutput {
	return o
}

func (o ResourceDiscoveryModeOutput) ToResourceDiscoveryModeOutputWithContext(ctx context.Context) ResourceDiscoveryModeOutput {
	return o
}

func (o ResourceDiscoveryModeOutput) ToResourceDiscoveryModePtrOutput() ResourceDiscoveryModePtrOutput {
	return o.ToResourceDiscoveryModePtrOutputWithContext(context.Background())
}

func (o ResourceDiscoveryModeOutput) ToResourceDiscoveryModePtrOutputWithContext(ctx context.Context) ResourceDiscoveryModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceDiscoveryMode) *ResourceDiscoveryMode {
		return &v
	}).(ResourceDiscoveryModePtrOutput)
}

func (o ResourceDiscoveryModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceDiscoveryModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceDiscoveryMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceDiscoveryModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceDiscoveryModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceDiscoveryMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceDiscoveryModePtrOutput struct{ *pulumi.OutputState }

func (ResourceDiscoveryModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceDiscoveryMode)(nil)).Elem()
}

func (o ResourceDiscoveryModePtrOutput) ToResourceDiscoveryModePtrOutput() ResourceDiscoveryModePtrOutput {
	return o
}

func (o ResourceDiscoveryModePtrOutput) ToResourceDiscoveryModePtrOutputWithContext(ctx context.Context) ResourceDiscoveryModePtrOutput {
	return o
}

func (o ResourceDiscoveryModePtrOutput) Elem() ResourceDiscoveryModeOutput {
	return o.ApplyT(func(v *ResourceDiscoveryMode) ResourceDiscoveryMode {
		if v != nil {
			return *v
		}
		var ret ResourceDiscoveryMode
		return ret
	}).(ResourceDiscoveryModeOutput)
}

func (o ResourceDiscoveryModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceDiscoveryModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceDiscoveryMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ResourceDiscoveryModeInput interface {
	pulumi.Input

	ToResourceDiscoveryModeOutput() ResourceDiscoveryModeOutput
	ToResourceDiscoveryModeOutputWithContext(context.Context) ResourceDiscoveryModeOutput
}

var resourceDiscoveryModePtrType = reflect.TypeOf((**ResourceDiscoveryMode)(nil)).Elem()

type ResourceDiscoveryModePtrInput interface {
	pulumi.Input

	ToResourceDiscoveryModePtrOutput() ResourceDiscoveryModePtrOutput
	ToResourceDiscoveryModePtrOutputWithContext(context.Context) ResourceDiscoveryModePtrOutput
}

type resourceDiscoveryModePtr string

func ResourceDiscoveryModePtr(v string) ResourceDiscoveryModePtrInput {
	return (*resourceDiscoveryModePtr)(&v)
}

func (*resourceDiscoveryModePtr) ElementType() reflect.Type {
	return resourceDiscoveryModePtrType
}

func (in *resourceDiscoveryModePtr) ToResourceDiscoveryModePtrOutput() ResourceDiscoveryModePtrOutput {
	return pulumi.ToOutput(in).(ResourceDiscoveryModePtrOutput)
}

func (in *resourceDiscoveryModePtr) ToResourceDiscoveryModePtrOutputWithContext(ctx context.Context) ResourceDiscoveryModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceDiscoveryModePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceDiscoveryModeOutput{})
	pulumi.RegisterOutputType(ResourceDiscoveryModePtrOutput{})
}
