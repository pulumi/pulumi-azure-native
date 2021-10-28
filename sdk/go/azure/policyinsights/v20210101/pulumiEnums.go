


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComplianceState string

const (
	// The resource is in compliance with the policy.
	ComplianceStateCompliant = ComplianceState("Compliant")
	// The resource is not in compliance with the policy.
	ComplianceStateNonCompliant = ComplianceState("NonCompliant")
	// The compliance state of the resource is not known.
	ComplianceStateUnknown = ComplianceState("Unknown")
)

func (ComplianceState) ElementType() reflect.Type {
	return reflect.TypeOf((*ComplianceState)(nil)).Elem()
}

func (e ComplianceState) ToComplianceStateOutput() ComplianceStateOutput {
	return pulumi.ToOutput(e).(ComplianceStateOutput)
}

func (e ComplianceState) ToComplianceStateOutputWithContext(ctx context.Context) ComplianceStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ComplianceStateOutput)
}

func (e ComplianceState) ToComplianceStatePtrOutput() ComplianceStatePtrOutput {
	return e.ToComplianceStatePtrOutputWithContext(context.Background())
}

func (e ComplianceState) ToComplianceStatePtrOutputWithContext(ctx context.Context) ComplianceStatePtrOutput {
	return ComplianceState(e).ToComplianceStateOutputWithContext(ctx).ToComplianceStatePtrOutputWithContext(ctx)
}

func (e ComplianceState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComplianceState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComplianceState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ComplianceState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ComplianceStateOutput struct{ *pulumi.OutputState }

func (ComplianceStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComplianceState)(nil)).Elem()
}

func (o ComplianceStateOutput) ToComplianceStateOutput() ComplianceStateOutput {
	return o
}

func (o ComplianceStateOutput) ToComplianceStateOutputWithContext(ctx context.Context) ComplianceStateOutput {
	return o
}

func (o ComplianceStateOutput) ToComplianceStatePtrOutput() ComplianceStatePtrOutput {
	return o.ToComplianceStatePtrOutputWithContext(context.Background())
}

func (o ComplianceStateOutput) ToComplianceStatePtrOutputWithContext(ctx context.Context) ComplianceStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComplianceState) *ComplianceState {
		return &v
	}).(ComplianceStatePtrOutput)
}

func (o ComplianceStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ComplianceStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComplianceState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ComplianceStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComplianceStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComplianceState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ComplianceStatePtrOutput struct{ *pulumi.OutputState }

func (ComplianceStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComplianceState)(nil)).Elem()
}

func (o ComplianceStatePtrOutput) ToComplianceStatePtrOutput() ComplianceStatePtrOutput {
	return o
}

func (o ComplianceStatePtrOutput) ToComplianceStatePtrOutputWithContext(ctx context.Context) ComplianceStatePtrOutput {
	return o
}

func (o ComplianceStatePtrOutput) Elem() ComplianceStateOutput {
	return o.ApplyT(func(v *ComplianceState) ComplianceState {
		if v != nil {
			return *v
		}
		var ret ComplianceState
		return ret
	}).(ComplianceStateOutput)
}

func (o ComplianceStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComplianceStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ComplianceState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ComplianceStateInput interface {
	pulumi.Input

	ToComplianceStateOutput() ComplianceStateOutput
	ToComplianceStateOutputWithContext(context.Context) ComplianceStateOutput
}

var complianceStatePtrType = reflect.TypeOf((**ComplianceState)(nil)).Elem()

type ComplianceStatePtrInput interface {
	pulumi.Input

	ToComplianceStatePtrOutput() ComplianceStatePtrOutput
	ToComplianceStatePtrOutputWithContext(context.Context) ComplianceStatePtrOutput
}

type complianceStatePtr string

func ComplianceStatePtr(v string) ComplianceStatePtrInput {
	return (*complianceStatePtr)(&v)
}

func (*complianceStatePtr) ElementType() reflect.Type {
	return complianceStatePtrType
}

func (in *complianceStatePtr) ToComplianceStatePtrOutput() ComplianceStatePtrOutput {
	return pulumi.ToOutput(in).(ComplianceStatePtrOutput)
}

func (in *complianceStatePtr) ToComplianceStatePtrOutputWithContext(ctx context.Context) ComplianceStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ComplianceStatePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ComplianceStateOutput{})
	pulumi.RegisterOutputType(ComplianceStatePtrOutput{})
}
