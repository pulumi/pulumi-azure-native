


package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TemplateSpecArtifactKind string

const (
	// The artifact represents an embedded Azure Resource Manager template.
	TemplateSpecArtifactKindTemplate = TemplateSpecArtifactKind("template")
)

func (TemplateSpecArtifactKind) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateSpecArtifactKind)(nil)).Elem()
}

func (e TemplateSpecArtifactKind) ToTemplateSpecArtifactKindOutput() TemplateSpecArtifactKindOutput {
	return pulumi.ToOutput(e).(TemplateSpecArtifactKindOutput)
}

func (e TemplateSpecArtifactKind) ToTemplateSpecArtifactKindOutputWithContext(ctx context.Context) TemplateSpecArtifactKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TemplateSpecArtifactKindOutput)
}

func (e TemplateSpecArtifactKind) ToTemplateSpecArtifactKindPtrOutput() TemplateSpecArtifactKindPtrOutput {
	return e.ToTemplateSpecArtifactKindPtrOutputWithContext(context.Background())
}

func (e TemplateSpecArtifactKind) ToTemplateSpecArtifactKindPtrOutputWithContext(ctx context.Context) TemplateSpecArtifactKindPtrOutput {
	return TemplateSpecArtifactKind(e).ToTemplateSpecArtifactKindOutputWithContext(ctx).ToTemplateSpecArtifactKindPtrOutputWithContext(ctx)
}

func (e TemplateSpecArtifactKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TemplateSpecArtifactKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TemplateSpecArtifactKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TemplateSpecArtifactKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TemplateSpecArtifactKindOutput struct{ *pulumi.OutputState }

func (TemplateSpecArtifactKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateSpecArtifactKind)(nil)).Elem()
}

func (o TemplateSpecArtifactKindOutput) ToTemplateSpecArtifactKindOutput() TemplateSpecArtifactKindOutput {
	return o
}

func (o TemplateSpecArtifactKindOutput) ToTemplateSpecArtifactKindOutputWithContext(ctx context.Context) TemplateSpecArtifactKindOutput {
	return o
}

func (o TemplateSpecArtifactKindOutput) ToTemplateSpecArtifactKindPtrOutput() TemplateSpecArtifactKindPtrOutput {
	return o.ToTemplateSpecArtifactKindPtrOutputWithContext(context.Background())
}

func (o TemplateSpecArtifactKindOutput) ToTemplateSpecArtifactKindPtrOutputWithContext(ctx context.Context) TemplateSpecArtifactKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TemplateSpecArtifactKind) *TemplateSpecArtifactKind {
		return &v
	}).(TemplateSpecArtifactKindPtrOutput)
}

func (o TemplateSpecArtifactKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TemplateSpecArtifactKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TemplateSpecArtifactKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TemplateSpecArtifactKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TemplateSpecArtifactKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TemplateSpecArtifactKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TemplateSpecArtifactKindPtrOutput struct{ *pulumi.OutputState }

func (TemplateSpecArtifactKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateSpecArtifactKind)(nil)).Elem()
}

func (o TemplateSpecArtifactKindPtrOutput) ToTemplateSpecArtifactKindPtrOutput() TemplateSpecArtifactKindPtrOutput {
	return o
}

func (o TemplateSpecArtifactKindPtrOutput) ToTemplateSpecArtifactKindPtrOutputWithContext(ctx context.Context) TemplateSpecArtifactKindPtrOutput {
	return o
}

func (o TemplateSpecArtifactKindPtrOutput) Elem() TemplateSpecArtifactKindOutput {
	return o.ApplyT(func(v *TemplateSpecArtifactKind) TemplateSpecArtifactKind {
		if v != nil {
			return *v
		}
		var ret TemplateSpecArtifactKind
		return ret
	}).(TemplateSpecArtifactKindOutput)
}

func (o TemplateSpecArtifactKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TemplateSpecArtifactKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TemplateSpecArtifactKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TemplateSpecArtifactKindInput interface {
	pulumi.Input

	ToTemplateSpecArtifactKindOutput() TemplateSpecArtifactKindOutput
	ToTemplateSpecArtifactKindOutputWithContext(context.Context) TemplateSpecArtifactKindOutput
}

var templateSpecArtifactKindPtrType = reflect.TypeOf((**TemplateSpecArtifactKind)(nil)).Elem()

type TemplateSpecArtifactKindPtrInput interface {
	pulumi.Input

	ToTemplateSpecArtifactKindPtrOutput() TemplateSpecArtifactKindPtrOutput
	ToTemplateSpecArtifactKindPtrOutputWithContext(context.Context) TemplateSpecArtifactKindPtrOutput
}

type templateSpecArtifactKindPtr string

func TemplateSpecArtifactKindPtr(v string) TemplateSpecArtifactKindPtrInput {
	return (*templateSpecArtifactKindPtr)(&v)
}

func (*templateSpecArtifactKindPtr) ElementType() reflect.Type {
	return templateSpecArtifactKindPtrType
}

func (in *templateSpecArtifactKindPtr) ToTemplateSpecArtifactKindPtrOutput() TemplateSpecArtifactKindPtrOutput {
	return pulumi.ToOutput(in).(TemplateSpecArtifactKindPtrOutput)
}

func (in *templateSpecArtifactKindPtr) ToTemplateSpecArtifactKindPtrOutputWithContext(ctx context.Context) TemplateSpecArtifactKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TemplateSpecArtifactKindPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(TemplateSpecArtifactKindOutput{})
	pulumi.RegisterOutputType(TemplateSpecArtifactKindPtrOutput{})
}
