


package hybridcompute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PublicNetworkAccessType string

const (
	// Allows Azure Arc agents to communicate with Azure Arc services over both public (internet) and private endpoints.
	PublicNetworkAccessTypeEnabled = PublicNetworkAccessType("Enabled")
	// Does not allow Azure Arc agents to communicate with Azure Arc services over public (internet) endpoints. The agents must use the private link.
	PublicNetworkAccessTypeDisabled = PublicNetworkAccessType("Disabled")
)

func (PublicNetworkAccessType) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccessType)(nil)).Elem()
}

func (e PublicNetworkAccessType) ToPublicNetworkAccessTypeOutput() PublicNetworkAccessTypeOutput {
	return pulumi.ToOutput(e).(PublicNetworkAccessTypeOutput)
}

func (e PublicNetworkAccessType) ToPublicNetworkAccessTypeOutputWithContext(ctx context.Context) PublicNetworkAccessTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PublicNetworkAccessTypeOutput)
}

func (e PublicNetworkAccessType) ToPublicNetworkAccessTypePtrOutput() PublicNetworkAccessTypePtrOutput {
	return e.ToPublicNetworkAccessTypePtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccessType) ToPublicNetworkAccessTypePtrOutputWithContext(ctx context.Context) PublicNetworkAccessTypePtrOutput {
	return PublicNetworkAccessType(e).ToPublicNetworkAccessTypeOutputWithContext(ctx).ToPublicNetworkAccessTypePtrOutputWithContext(ctx)
}

func (e PublicNetworkAccessType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccessType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccessType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccessType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PublicNetworkAccessTypeOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccessType)(nil)).Elem()
}

func (o PublicNetworkAccessTypeOutput) ToPublicNetworkAccessTypeOutput() PublicNetworkAccessTypeOutput {
	return o
}

func (o PublicNetworkAccessTypeOutput) ToPublicNetworkAccessTypeOutputWithContext(ctx context.Context) PublicNetworkAccessTypeOutput {
	return o
}

func (o PublicNetworkAccessTypeOutput) ToPublicNetworkAccessTypePtrOutput() PublicNetworkAccessTypePtrOutput {
	return o.ToPublicNetworkAccessTypePtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessTypeOutput) ToPublicNetworkAccessTypePtrOutputWithContext(ctx context.Context) PublicNetworkAccessTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicNetworkAccessType) *PublicNetworkAccessType {
		return &v
	}).(PublicNetworkAccessTypePtrOutput)
}

func (o PublicNetworkAccessTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PublicNetworkAccessTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccessType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PublicNetworkAccessTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccessType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PublicNetworkAccessTypePtrOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicNetworkAccessType)(nil)).Elem()
}

func (o PublicNetworkAccessTypePtrOutput) ToPublicNetworkAccessTypePtrOutput() PublicNetworkAccessTypePtrOutput {
	return o
}

func (o PublicNetworkAccessTypePtrOutput) ToPublicNetworkAccessTypePtrOutputWithContext(ctx context.Context) PublicNetworkAccessTypePtrOutput {
	return o
}

func (o PublicNetworkAccessTypePtrOutput) Elem() PublicNetworkAccessTypeOutput {
	return o.ApplyT(func(v *PublicNetworkAccessType) PublicNetworkAccessType {
		if v != nil {
			return *v
		}
		var ret PublicNetworkAccessType
		return ret
	}).(PublicNetworkAccessTypeOutput)
}

func (o PublicNetworkAccessTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PublicNetworkAccessType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PublicNetworkAccessTypeInput interface {
	pulumi.Input

	ToPublicNetworkAccessTypeOutput() PublicNetworkAccessTypeOutput
	ToPublicNetworkAccessTypeOutputWithContext(context.Context) PublicNetworkAccessTypeOutput
}

var publicNetworkAccessTypePtrType = reflect.TypeOf((**PublicNetworkAccessType)(nil)).Elem()

type PublicNetworkAccessTypePtrInput interface {
	pulumi.Input

	ToPublicNetworkAccessTypePtrOutput() PublicNetworkAccessTypePtrOutput
	ToPublicNetworkAccessTypePtrOutputWithContext(context.Context) PublicNetworkAccessTypePtrOutput
}

type publicNetworkAccessTypePtr string

func PublicNetworkAccessTypePtr(v string) PublicNetworkAccessTypePtrInput {
	return (*publicNetworkAccessTypePtr)(&v)
}

func (*publicNetworkAccessTypePtr) ElementType() reflect.Type {
	return publicNetworkAccessTypePtrType
}

func (in *publicNetworkAccessTypePtr) ToPublicNetworkAccessTypePtrOutput() PublicNetworkAccessTypePtrOutput {
	return pulumi.ToOutput(in).(PublicNetworkAccessTypePtrOutput)
}

func (in *publicNetworkAccessTypePtr) ToPublicNetworkAccessTypePtrOutputWithContext(ctx context.Context) PublicNetworkAccessTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PublicNetworkAccessTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PublicNetworkAccessTypeOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessTypePtrOutput{})
}
