


package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrincipalType string

const (
	PrincipalTypeUser             = PrincipalType("User")
	PrincipalTypeGroup            = PrincipalType("Group")
	PrincipalTypeServicePrincipal = PrincipalType("ServicePrincipal")
	PrincipalTypeForeignGroup     = PrincipalType("ForeignGroup")
)

func (PrincipalType) ElementType() reflect.Type {
	return reflect.TypeOf((*PrincipalType)(nil)).Elem()
}

func (e PrincipalType) ToPrincipalTypeOutput() PrincipalTypeOutput {
	return pulumi.ToOutput(e).(PrincipalTypeOutput)
}

func (e PrincipalType) ToPrincipalTypeOutputWithContext(ctx context.Context) PrincipalTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrincipalTypeOutput)
}

func (e PrincipalType) ToPrincipalTypePtrOutput() PrincipalTypePtrOutput {
	return e.ToPrincipalTypePtrOutputWithContext(context.Background())
}

func (e PrincipalType) ToPrincipalTypePtrOutputWithContext(ctx context.Context) PrincipalTypePtrOutput {
	return PrincipalType(e).ToPrincipalTypeOutputWithContext(ctx).ToPrincipalTypePtrOutputWithContext(ctx)
}

func (e PrincipalType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrincipalType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrincipalType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrincipalType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrincipalTypeOutput struct{ *pulumi.OutputState }

func (PrincipalTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrincipalType)(nil)).Elem()
}

func (o PrincipalTypeOutput) ToPrincipalTypeOutput() PrincipalTypeOutput {
	return o
}

func (o PrincipalTypeOutput) ToPrincipalTypeOutputWithContext(ctx context.Context) PrincipalTypeOutput {
	return o
}

func (o PrincipalTypeOutput) ToPrincipalTypePtrOutput() PrincipalTypePtrOutput {
	return o.ToPrincipalTypePtrOutputWithContext(context.Background())
}

func (o PrincipalTypeOutput) ToPrincipalTypePtrOutputWithContext(ctx context.Context) PrincipalTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrincipalType) *PrincipalType {
		return &v
	}).(PrincipalTypePtrOutput)
}

func (o PrincipalTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrincipalTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrincipalType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrincipalTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrincipalTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrincipalType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrincipalTypePtrOutput struct{ *pulumi.OutputState }

func (PrincipalTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrincipalType)(nil)).Elem()
}

func (o PrincipalTypePtrOutput) ToPrincipalTypePtrOutput() PrincipalTypePtrOutput {
	return o
}

func (o PrincipalTypePtrOutput) ToPrincipalTypePtrOutputWithContext(ctx context.Context) PrincipalTypePtrOutput {
	return o
}

func (o PrincipalTypePtrOutput) Elem() PrincipalTypeOutput {
	return o.ApplyT(func(v *PrincipalType) PrincipalType {
		if v != nil {
			return *v
		}
		var ret PrincipalType
		return ret
	}).(PrincipalTypeOutput)
}

func (o PrincipalTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrincipalTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrincipalType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PrincipalTypeInput interface {
	pulumi.Input

	ToPrincipalTypeOutput() PrincipalTypeOutput
	ToPrincipalTypeOutputWithContext(context.Context) PrincipalTypeOutput
}

var principalTypePtrType = reflect.TypeOf((**PrincipalType)(nil)).Elem()

type PrincipalTypePtrInput interface {
	pulumi.Input

	ToPrincipalTypePtrOutput() PrincipalTypePtrOutput
	ToPrincipalTypePtrOutputWithContext(context.Context) PrincipalTypePtrOutput
}

type principalTypePtr string

func PrincipalTypePtr(v string) PrincipalTypePtrInput {
	return (*principalTypePtr)(&v)
}

func (*principalTypePtr) ElementType() reflect.Type {
	return principalTypePtrType
}

func (in *principalTypePtr) ToPrincipalTypePtrOutput() PrincipalTypePtrOutput {
	return pulumi.ToOutput(in).(PrincipalTypePtrOutput)
}

func (in *principalTypePtr) ToPrincipalTypePtrOutputWithContext(ctx context.Context) PrincipalTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrincipalTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrincipalTypeInput)(nil)).Elem(), PrincipalType("User"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrincipalTypePtrInput)(nil)).Elem(), PrincipalType("User"))
	pulumi.RegisterOutputType(PrincipalTypeOutput{})
	pulumi.RegisterOutputType(PrincipalTypePtrOutput{})
}
