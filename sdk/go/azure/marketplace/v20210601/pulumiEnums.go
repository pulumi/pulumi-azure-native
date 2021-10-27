


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Accessibility string

const (
	AccessibilityUnknown                    = Accessibility("Unknown")
	AccessibilityPublic                     = Accessibility("Public")
	AccessibilityPrivateTenantOnLevel       = Accessibility("PrivateTenantOnLevel")
	AccessibilityPrivateSubscriptionOnLevel = Accessibility("PrivateSubscriptionOnLevel")
)

func (Accessibility) ElementType() reflect.Type {
	return reflect.TypeOf((*Accessibility)(nil)).Elem()
}

func (e Accessibility) ToAccessibilityOutput() AccessibilityOutput {
	return pulumi.ToOutput(e).(AccessibilityOutput)
}

func (e Accessibility) ToAccessibilityOutputWithContext(ctx context.Context) AccessibilityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessibilityOutput)
}

func (e Accessibility) ToAccessibilityPtrOutput() AccessibilityPtrOutput {
	return e.ToAccessibilityPtrOutputWithContext(context.Background())
}

func (e Accessibility) ToAccessibilityPtrOutputWithContext(ctx context.Context) AccessibilityPtrOutput {
	return Accessibility(e).ToAccessibilityOutputWithContext(ctx).ToAccessibilityPtrOutputWithContext(ctx)
}

func (e Accessibility) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Accessibility) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Accessibility) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Accessibility) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessibilityOutput struct{ *pulumi.OutputState }

func (AccessibilityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Accessibility)(nil)).Elem()
}

func (o AccessibilityOutput) ToAccessibilityOutput() AccessibilityOutput {
	return o
}

func (o AccessibilityOutput) ToAccessibilityOutputWithContext(ctx context.Context) AccessibilityOutput {
	return o
}

func (o AccessibilityOutput) ToAccessibilityPtrOutput() AccessibilityPtrOutput {
	return o.ToAccessibilityPtrOutputWithContext(context.Background())
}

func (o AccessibilityOutput) ToAccessibilityPtrOutputWithContext(ctx context.Context) AccessibilityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Accessibility) *Accessibility {
		return &v
	}).(AccessibilityPtrOutput)
}

func (o AccessibilityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessibilityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Accessibility) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessibilityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessibilityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Accessibility) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessibilityPtrOutput struct{ *pulumi.OutputState }

func (AccessibilityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Accessibility)(nil)).Elem()
}

func (o AccessibilityPtrOutput) ToAccessibilityPtrOutput() AccessibilityPtrOutput {
	return o
}

func (o AccessibilityPtrOutput) ToAccessibilityPtrOutputWithContext(ctx context.Context) AccessibilityPtrOutput {
	return o
}

func (o AccessibilityPtrOutput) Elem() AccessibilityOutput {
	return o.ApplyT(func(v *Accessibility) Accessibility {
		if v != nil {
			return *v
		}
		var ret Accessibility
		return ret
	}).(AccessibilityOutput)
}

func (o AccessibilityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessibilityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Accessibility) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AccessibilityInput interface {
	pulumi.Input

	ToAccessibilityOutput() AccessibilityOutput
	ToAccessibilityOutputWithContext(context.Context) AccessibilityOutput
}

var accessibilityPtrType = reflect.TypeOf((**Accessibility)(nil)).Elem()

type AccessibilityPtrInput interface {
	pulumi.Input

	ToAccessibilityPtrOutput() AccessibilityPtrOutput
	ToAccessibilityPtrOutputWithContext(context.Context) AccessibilityPtrOutput
}

type accessibilityPtr string

func AccessibilityPtr(v string) AccessibilityPtrInput {
	return (*accessibilityPtr)(&v)
}

func (*accessibilityPtr) ElementType() reflect.Type {
	return accessibilityPtrType
}

func (in *accessibilityPtr) ToAccessibilityPtrOutput() AccessibilityPtrOutput {
	return pulumi.ToOutput(in).(AccessibilityPtrOutput)
}

func (in *accessibilityPtr) ToAccessibilityPtrOutputWithContext(ctx context.Context) AccessibilityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessibilityPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessibilityInput)(nil)).Elem(), Accessibility("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*AccessibilityPtrInput)(nil)).Elem(), Accessibility("Unknown"))
	pulumi.RegisterOutputType(AccessibilityOutput{})
	pulumi.RegisterOutputType(AccessibilityPtrOutput{})
}
