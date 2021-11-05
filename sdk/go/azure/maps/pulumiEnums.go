


package maps

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SigningKey string

const (
	SigningKeyPrimaryKey   = SigningKey("primaryKey")
	SigningKeySecondaryKey = SigningKey("secondaryKey")
)

func (SigningKey) ElementType() reflect.Type {
	return reflect.TypeOf((*SigningKey)(nil)).Elem()
}

func (e SigningKey) ToSigningKeyOutput() SigningKeyOutput {
	return pulumi.ToOutput(e).(SigningKeyOutput)
}

func (e SigningKey) ToSigningKeyOutputWithContext(ctx context.Context) SigningKeyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SigningKeyOutput)
}

func (e SigningKey) ToSigningKeyPtrOutput() SigningKeyPtrOutput {
	return e.ToSigningKeyPtrOutputWithContext(context.Background())
}

func (e SigningKey) ToSigningKeyPtrOutputWithContext(ctx context.Context) SigningKeyPtrOutput {
	return SigningKey(e).ToSigningKeyOutputWithContext(ctx).ToSigningKeyPtrOutputWithContext(ctx)
}

func (e SigningKey) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SigningKey) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SigningKey) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SigningKey) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SigningKeyOutput struct{ *pulumi.OutputState }

func (SigningKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SigningKey)(nil)).Elem()
}

func (o SigningKeyOutput) ToSigningKeyOutput() SigningKeyOutput {
	return o
}

func (o SigningKeyOutput) ToSigningKeyOutputWithContext(ctx context.Context) SigningKeyOutput {
	return o
}

func (o SigningKeyOutput) ToSigningKeyPtrOutput() SigningKeyPtrOutput {
	return o.ToSigningKeyPtrOutputWithContext(context.Background())
}

func (o SigningKeyOutput) ToSigningKeyPtrOutputWithContext(ctx context.Context) SigningKeyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SigningKey) *SigningKey {
		return &v
	}).(SigningKeyPtrOutput)
}

func (o SigningKeyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SigningKeyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SigningKey) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SigningKeyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SigningKeyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SigningKey) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SigningKeyPtrOutput struct{ *pulumi.OutputState }

func (SigningKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SigningKey)(nil)).Elem()
}

func (o SigningKeyPtrOutput) ToSigningKeyPtrOutput() SigningKeyPtrOutput {
	return o
}

func (o SigningKeyPtrOutput) ToSigningKeyPtrOutputWithContext(ctx context.Context) SigningKeyPtrOutput {
	return o
}

func (o SigningKeyPtrOutput) Elem() SigningKeyOutput {
	return o.ApplyT(func(v *SigningKey) SigningKey {
		if v != nil {
			return *v
		}
		var ret SigningKey
		return ret
	}).(SigningKeyOutput)
}

func (o SigningKeyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SigningKeyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SigningKey) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SigningKeyInput interface {
	pulumi.Input

	ToSigningKeyOutput() SigningKeyOutput
	ToSigningKeyOutputWithContext(context.Context) SigningKeyOutput
}

var signingKeyPtrType = reflect.TypeOf((**SigningKey)(nil)).Elem()

type SigningKeyPtrInput interface {
	pulumi.Input

	ToSigningKeyPtrOutput() SigningKeyPtrOutput
	ToSigningKeyPtrOutputWithContext(context.Context) SigningKeyPtrOutput
}

type signingKeyPtr string

func SigningKeyPtr(v string) SigningKeyPtrInput {
	return (*signingKeyPtr)(&v)
}

func (*signingKeyPtr) ElementType() reflect.Type {
	return signingKeyPtrType
}

func (in *signingKeyPtr) ToSigningKeyPtrOutput() SigningKeyPtrOutput {
	return pulumi.ToOutput(in).(SigningKeyPtrOutput)
}

func (in *signingKeyPtr) ToSigningKeyPtrOutputWithContext(ctx context.Context) SigningKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SigningKeyPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SigningKeyOutput{})
	pulumi.RegisterOutputType(SigningKeyPtrOutput{})
}
