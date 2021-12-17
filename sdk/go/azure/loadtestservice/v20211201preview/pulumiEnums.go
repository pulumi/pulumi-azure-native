


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemAssignedServiceIdentityType string

const (
	SystemAssignedServiceIdentityTypeNone           = SystemAssignedServiceIdentityType("None")
	SystemAssignedServiceIdentityTypeSystemAssigned = SystemAssignedServiceIdentityType("SystemAssigned")
)

func (SystemAssignedServiceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAssignedServiceIdentityType)(nil)).Elem()
}

func (e SystemAssignedServiceIdentityType) ToSystemAssignedServiceIdentityTypeOutput() SystemAssignedServiceIdentityTypeOutput {
	return pulumi.ToOutput(e).(SystemAssignedServiceIdentityTypeOutput)
}

func (e SystemAssignedServiceIdentityType) ToSystemAssignedServiceIdentityTypeOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SystemAssignedServiceIdentityTypeOutput)
}

func (e SystemAssignedServiceIdentityType) ToSystemAssignedServiceIdentityTypePtrOutput() SystemAssignedServiceIdentityTypePtrOutput {
	return e.ToSystemAssignedServiceIdentityTypePtrOutputWithContext(context.Background())
}

func (e SystemAssignedServiceIdentityType) ToSystemAssignedServiceIdentityTypePtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityTypePtrOutput {
	return SystemAssignedServiceIdentityType(e).ToSystemAssignedServiceIdentityTypeOutputWithContext(ctx).ToSystemAssignedServiceIdentityTypePtrOutputWithContext(ctx)
}

func (e SystemAssignedServiceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SystemAssignedServiceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SystemAssignedServiceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SystemAssignedServiceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SystemAssignedServiceIdentityTypeOutput struct{ *pulumi.OutputState }

func (SystemAssignedServiceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAssignedServiceIdentityType)(nil)).Elem()
}

func (o SystemAssignedServiceIdentityTypeOutput) ToSystemAssignedServiceIdentityTypeOutput() SystemAssignedServiceIdentityTypeOutput {
	return o
}

func (o SystemAssignedServiceIdentityTypeOutput) ToSystemAssignedServiceIdentityTypeOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityTypeOutput {
	return o
}

func (o SystemAssignedServiceIdentityTypeOutput) ToSystemAssignedServiceIdentityTypePtrOutput() SystemAssignedServiceIdentityTypePtrOutput {
	return o.ToSystemAssignedServiceIdentityTypePtrOutputWithContext(context.Background())
}

func (o SystemAssignedServiceIdentityTypeOutput) ToSystemAssignedServiceIdentityTypePtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemAssignedServiceIdentityType) *SystemAssignedServiceIdentityType {
		return &v
	}).(SystemAssignedServiceIdentityTypePtrOutput)
}

func (o SystemAssignedServiceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SystemAssignedServiceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SystemAssignedServiceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SystemAssignedServiceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SystemAssignedServiceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SystemAssignedServiceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SystemAssignedServiceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (SystemAssignedServiceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAssignedServiceIdentityType)(nil)).Elem()
}

func (o SystemAssignedServiceIdentityTypePtrOutput) ToSystemAssignedServiceIdentityTypePtrOutput() SystemAssignedServiceIdentityTypePtrOutput {
	return o
}

func (o SystemAssignedServiceIdentityTypePtrOutput) ToSystemAssignedServiceIdentityTypePtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityTypePtrOutput {
	return o
}

func (o SystemAssignedServiceIdentityTypePtrOutput) Elem() SystemAssignedServiceIdentityTypeOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentityType) SystemAssignedServiceIdentityType {
		if v != nil {
			return *v
		}
		var ret SystemAssignedServiceIdentityType
		return ret
	}).(SystemAssignedServiceIdentityTypeOutput)
}

func (o SystemAssignedServiceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SystemAssignedServiceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SystemAssignedServiceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SystemAssignedServiceIdentityTypeInput interface {
	pulumi.Input

	ToSystemAssignedServiceIdentityTypeOutput() SystemAssignedServiceIdentityTypeOutput
	ToSystemAssignedServiceIdentityTypeOutputWithContext(context.Context) SystemAssignedServiceIdentityTypeOutput
}

var systemAssignedServiceIdentityTypePtrType = reflect.TypeOf((**SystemAssignedServiceIdentityType)(nil)).Elem()

type SystemAssignedServiceIdentityTypePtrInput interface {
	pulumi.Input

	ToSystemAssignedServiceIdentityTypePtrOutput() SystemAssignedServiceIdentityTypePtrOutput
	ToSystemAssignedServiceIdentityTypePtrOutputWithContext(context.Context) SystemAssignedServiceIdentityTypePtrOutput
}

type systemAssignedServiceIdentityTypePtr string

func SystemAssignedServiceIdentityTypePtr(v string) SystemAssignedServiceIdentityTypePtrInput {
	return (*systemAssignedServiceIdentityTypePtr)(&v)
}

func (*systemAssignedServiceIdentityTypePtr) ElementType() reflect.Type {
	return systemAssignedServiceIdentityTypePtrType
}

func (in *systemAssignedServiceIdentityTypePtr) ToSystemAssignedServiceIdentityTypePtrOutput() SystemAssignedServiceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(SystemAssignedServiceIdentityTypePtrOutput)
}

func (in *systemAssignedServiceIdentityTypePtr) ToSystemAssignedServiceIdentityTypePtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SystemAssignedServiceIdentityTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityTypeOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityTypePtrOutput{})
}
