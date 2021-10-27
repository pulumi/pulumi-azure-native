


package v20180617preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SharedTypeKind string

const (
	SharedTypeKindUser   = SharedTypeKind("user")
	SharedTypeKindShared = SharedTypeKind("shared")
)

func (SharedTypeKind) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedTypeKind)(nil)).Elem()
}

func (e SharedTypeKind) ToSharedTypeKindOutput() SharedTypeKindOutput {
	return pulumi.ToOutput(e).(SharedTypeKindOutput)
}

func (e SharedTypeKind) ToSharedTypeKindOutputWithContext(ctx context.Context) SharedTypeKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SharedTypeKindOutput)
}

func (e SharedTypeKind) ToSharedTypeKindPtrOutput() SharedTypeKindPtrOutput {
	return e.ToSharedTypeKindPtrOutputWithContext(context.Background())
}

func (e SharedTypeKind) ToSharedTypeKindPtrOutputWithContext(ctx context.Context) SharedTypeKindPtrOutput {
	return SharedTypeKind(e).ToSharedTypeKindOutputWithContext(ctx).ToSharedTypeKindPtrOutputWithContext(ctx)
}

func (e SharedTypeKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SharedTypeKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SharedTypeKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SharedTypeKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SharedTypeKindOutput struct{ *pulumi.OutputState }

func (SharedTypeKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedTypeKind)(nil)).Elem()
}

func (o SharedTypeKindOutput) ToSharedTypeKindOutput() SharedTypeKindOutput {
	return o
}

func (o SharedTypeKindOutput) ToSharedTypeKindOutputWithContext(ctx context.Context) SharedTypeKindOutput {
	return o
}

func (o SharedTypeKindOutput) ToSharedTypeKindPtrOutput() SharedTypeKindPtrOutput {
	return o.ToSharedTypeKindPtrOutputWithContext(context.Background())
}

func (o SharedTypeKindOutput) ToSharedTypeKindPtrOutputWithContext(ctx context.Context) SharedTypeKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SharedTypeKind) *SharedTypeKind {
		return &v
	}).(SharedTypeKindPtrOutput)
}

func (o SharedTypeKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SharedTypeKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SharedTypeKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SharedTypeKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SharedTypeKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SharedTypeKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SharedTypeKindPtrOutput struct{ *pulumi.OutputState }

func (SharedTypeKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SharedTypeKind)(nil)).Elem()
}

func (o SharedTypeKindPtrOutput) ToSharedTypeKindPtrOutput() SharedTypeKindPtrOutput {
	return o
}

func (o SharedTypeKindPtrOutput) ToSharedTypeKindPtrOutputWithContext(ctx context.Context) SharedTypeKindPtrOutput {
	return o
}

func (o SharedTypeKindPtrOutput) Elem() SharedTypeKindOutput {
	return o.ApplyT(func(v *SharedTypeKind) SharedTypeKind {
		if v != nil {
			return *v
		}
		var ret SharedTypeKind
		return ret
	}).(SharedTypeKindOutput)
}

func (o SharedTypeKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SharedTypeKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SharedTypeKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SharedTypeKindInput interface {
	pulumi.Input

	ToSharedTypeKindOutput() SharedTypeKindOutput
	ToSharedTypeKindOutputWithContext(context.Context) SharedTypeKindOutput
}

var sharedTypeKindPtrType = reflect.TypeOf((**SharedTypeKind)(nil)).Elem()

type SharedTypeKindPtrInput interface {
	pulumi.Input

	ToSharedTypeKindPtrOutput() SharedTypeKindPtrOutput
	ToSharedTypeKindPtrOutputWithContext(context.Context) SharedTypeKindPtrOutput
}

type sharedTypeKindPtr string

func SharedTypeKindPtr(v string) SharedTypeKindPtrInput {
	return (*sharedTypeKindPtr)(&v)
}

func (*sharedTypeKindPtr) ElementType() reflect.Type {
	return sharedTypeKindPtrType
}

func (in *sharedTypeKindPtr) ToSharedTypeKindPtrOutput() SharedTypeKindPtrOutput {
	return pulumi.ToOutput(in).(SharedTypeKindPtrOutput)
}

func (in *sharedTypeKindPtr) ToSharedTypeKindPtrOutputWithContext(ctx context.Context) SharedTypeKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SharedTypeKindPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SharedTypeKindInput)(nil)).Elem(), SharedTypeKind("user"))
	pulumi.RegisterInputType(reflect.TypeOf((*SharedTypeKindPtrInput)(nil)).Elem(), SharedTypeKind("user"))
	pulumi.RegisterOutputType(SharedTypeKindOutput{})
	pulumi.RegisterOutputType(SharedTypeKindPtrOutput{})
}
