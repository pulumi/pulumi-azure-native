


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssociationType string

const (
	AssociationTypeSubnets = AssociationType("subnets")
)

func (AssociationType) ElementType() reflect.Type {
	return reflect.TypeOf((*AssociationType)(nil)).Elem()
}

func (e AssociationType) ToAssociationTypeOutput() AssociationTypeOutput {
	return pulumi.ToOutput(e).(AssociationTypeOutput)
}

func (e AssociationType) ToAssociationTypeOutputWithContext(ctx context.Context) AssociationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AssociationTypeOutput)
}

func (e AssociationType) ToAssociationTypePtrOutput() AssociationTypePtrOutput {
	return e.ToAssociationTypePtrOutputWithContext(context.Background())
}

func (e AssociationType) ToAssociationTypePtrOutputWithContext(ctx context.Context) AssociationTypePtrOutput {
	return AssociationType(e).ToAssociationTypeOutputWithContext(ctx).ToAssociationTypePtrOutputWithContext(ctx)
}

func (e AssociationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssociationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssociationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AssociationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AssociationTypeOutput struct{ *pulumi.OutputState }

func (AssociationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssociationType)(nil)).Elem()
}

func (o AssociationTypeOutput) ToAssociationTypeOutput() AssociationTypeOutput {
	return o
}

func (o AssociationTypeOutput) ToAssociationTypeOutputWithContext(ctx context.Context) AssociationTypeOutput {
	return o
}

func (o AssociationTypeOutput) ToAssociationTypePtrOutput() AssociationTypePtrOutput {
	return o.ToAssociationTypePtrOutputWithContext(context.Background())
}

func (o AssociationTypeOutput) ToAssociationTypePtrOutputWithContext(ctx context.Context) AssociationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssociationType) *AssociationType {
		return &v
	}).(AssociationTypePtrOutput)
}

func (o AssociationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AssociationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssociationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AssociationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssociationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssociationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AssociationTypePtrOutput struct{ *pulumi.OutputState }

func (AssociationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssociationType)(nil)).Elem()
}

func (o AssociationTypePtrOutput) ToAssociationTypePtrOutput() AssociationTypePtrOutput {
	return o
}

func (o AssociationTypePtrOutput) ToAssociationTypePtrOutputWithContext(ctx context.Context) AssociationTypePtrOutput {
	return o
}

func (o AssociationTypePtrOutput) Elem() AssociationTypeOutput {
	return o.ApplyT(func(v *AssociationType) AssociationType {
		if v != nil {
			return *v
		}
		var ret AssociationType
		return ret
	}).(AssociationTypeOutput)
}

func (o AssociationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssociationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AssociationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AssociationTypeInput interface {
	pulumi.Input

	ToAssociationTypeOutput() AssociationTypeOutput
	ToAssociationTypeOutputWithContext(context.Context) AssociationTypeOutput
}

var associationTypePtrType = reflect.TypeOf((**AssociationType)(nil)).Elem()

type AssociationTypePtrInput interface {
	pulumi.Input

	ToAssociationTypePtrOutput() AssociationTypePtrOutput
	ToAssociationTypePtrOutputWithContext(context.Context) AssociationTypePtrOutput
}

type associationTypePtr string

func AssociationTypePtr(v string) AssociationTypePtrInput {
	return (*associationTypePtr)(&v)
}

func (*associationTypePtr) ElementType() reflect.Type {
	return associationTypePtrType
}

func (in *associationTypePtr) ToAssociationTypePtrOutput() AssociationTypePtrOutput {
	return pulumi.ToOutput(in).(AssociationTypePtrOutput)
}

func (in *associationTypePtr) ToAssociationTypePtrOutputWithContext(ctx context.Context) AssociationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AssociationTypePtrOutput)
}

type FrontendIPAddressVersion string

const (
	FrontendIPAddressVersionIPv4 = FrontendIPAddressVersion("IPv4")
	FrontendIPAddressVersionIPv6 = FrontendIPAddressVersion("IPv6")
)

func (FrontendIPAddressVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendIPAddressVersion)(nil)).Elem()
}

func (e FrontendIPAddressVersion) ToFrontendIPAddressVersionOutput() FrontendIPAddressVersionOutput {
	return pulumi.ToOutput(e).(FrontendIPAddressVersionOutput)
}

func (e FrontendIPAddressVersion) ToFrontendIPAddressVersionOutputWithContext(ctx context.Context) FrontendIPAddressVersionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FrontendIPAddressVersionOutput)
}

func (e FrontendIPAddressVersion) ToFrontendIPAddressVersionPtrOutput() FrontendIPAddressVersionPtrOutput {
	return e.ToFrontendIPAddressVersionPtrOutputWithContext(context.Background())
}

func (e FrontendIPAddressVersion) ToFrontendIPAddressVersionPtrOutputWithContext(ctx context.Context) FrontendIPAddressVersionPtrOutput {
	return FrontendIPAddressVersion(e).ToFrontendIPAddressVersionOutputWithContext(ctx).ToFrontendIPAddressVersionPtrOutputWithContext(ctx)
}

func (e FrontendIPAddressVersion) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FrontendIPAddressVersion) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FrontendIPAddressVersion) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FrontendIPAddressVersion) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FrontendIPAddressVersionOutput struct{ *pulumi.OutputState }

func (FrontendIPAddressVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendIPAddressVersion)(nil)).Elem()
}

func (o FrontendIPAddressVersionOutput) ToFrontendIPAddressVersionOutput() FrontendIPAddressVersionOutput {
	return o
}

func (o FrontendIPAddressVersionOutput) ToFrontendIPAddressVersionOutputWithContext(ctx context.Context) FrontendIPAddressVersionOutput {
	return o
}

func (o FrontendIPAddressVersionOutput) ToFrontendIPAddressVersionPtrOutput() FrontendIPAddressVersionPtrOutput {
	return o.ToFrontendIPAddressVersionPtrOutputWithContext(context.Background())
}

func (o FrontendIPAddressVersionOutput) ToFrontendIPAddressVersionPtrOutputWithContext(ctx context.Context) FrontendIPAddressVersionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FrontendIPAddressVersion) *FrontendIPAddressVersion {
		return &v
	}).(FrontendIPAddressVersionPtrOutput)
}

func (o FrontendIPAddressVersionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FrontendIPAddressVersionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FrontendIPAddressVersion) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FrontendIPAddressVersionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FrontendIPAddressVersionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FrontendIPAddressVersion) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FrontendIPAddressVersionPtrOutput struct{ *pulumi.OutputState }

func (FrontendIPAddressVersionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontendIPAddressVersion)(nil)).Elem()
}

func (o FrontendIPAddressVersionPtrOutput) ToFrontendIPAddressVersionPtrOutput() FrontendIPAddressVersionPtrOutput {
	return o
}

func (o FrontendIPAddressVersionPtrOutput) ToFrontendIPAddressVersionPtrOutputWithContext(ctx context.Context) FrontendIPAddressVersionPtrOutput {
	return o
}

func (o FrontendIPAddressVersionPtrOutput) Elem() FrontendIPAddressVersionOutput {
	return o.ApplyT(func(v *FrontendIPAddressVersion) FrontendIPAddressVersion {
		if v != nil {
			return *v
		}
		var ret FrontendIPAddressVersion
		return ret
	}).(FrontendIPAddressVersionOutput)
}

func (o FrontendIPAddressVersionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FrontendIPAddressVersionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FrontendIPAddressVersion) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FrontendIPAddressVersionInput interface {
	pulumi.Input

	ToFrontendIPAddressVersionOutput() FrontendIPAddressVersionOutput
	ToFrontendIPAddressVersionOutputWithContext(context.Context) FrontendIPAddressVersionOutput
}

var frontendIPAddressVersionPtrType = reflect.TypeOf((**FrontendIPAddressVersion)(nil)).Elem()

type FrontendIPAddressVersionPtrInput interface {
	pulumi.Input

	ToFrontendIPAddressVersionPtrOutput() FrontendIPAddressVersionPtrOutput
	ToFrontendIPAddressVersionPtrOutputWithContext(context.Context) FrontendIPAddressVersionPtrOutput
}

type frontendIPAddressVersionPtr string

func FrontendIPAddressVersionPtr(v string) FrontendIPAddressVersionPtrInput {
	return (*frontendIPAddressVersionPtr)(&v)
}

func (*frontendIPAddressVersionPtr) ElementType() reflect.Type {
	return frontendIPAddressVersionPtrType
}

func (in *frontendIPAddressVersionPtr) ToFrontendIPAddressVersionPtrOutput() FrontendIPAddressVersionPtrOutput {
	return pulumi.ToOutput(in).(FrontendIPAddressVersionPtrOutput)
}

func (in *frontendIPAddressVersionPtr) ToFrontendIPAddressVersionPtrOutputWithContext(ctx context.Context) FrontendIPAddressVersionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FrontendIPAddressVersionPtrOutput)
}

type FrontendMode string

const (
	FrontendModePublic = FrontendMode("public")
)

func (FrontendMode) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendMode)(nil)).Elem()
}

func (e FrontendMode) ToFrontendModeOutput() FrontendModeOutput {
	return pulumi.ToOutput(e).(FrontendModeOutput)
}

func (e FrontendMode) ToFrontendModeOutputWithContext(ctx context.Context) FrontendModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FrontendModeOutput)
}

func (e FrontendMode) ToFrontendModePtrOutput() FrontendModePtrOutput {
	return e.ToFrontendModePtrOutputWithContext(context.Background())
}

func (e FrontendMode) ToFrontendModePtrOutputWithContext(ctx context.Context) FrontendModePtrOutput {
	return FrontendMode(e).ToFrontendModeOutputWithContext(ctx).ToFrontendModePtrOutputWithContext(ctx)
}

func (e FrontendMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FrontendMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FrontendMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FrontendMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FrontendModeOutput struct{ *pulumi.OutputState }

func (FrontendModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendMode)(nil)).Elem()
}

func (o FrontendModeOutput) ToFrontendModeOutput() FrontendModeOutput {
	return o
}

func (o FrontendModeOutput) ToFrontendModeOutputWithContext(ctx context.Context) FrontendModeOutput {
	return o
}

func (o FrontendModeOutput) ToFrontendModePtrOutput() FrontendModePtrOutput {
	return o.ToFrontendModePtrOutputWithContext(context.Background())
}

func (o FrontendModeOutput) ToFrontendModePtrOutputWithContext(ctx context.Context) FrontendModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FrontendMode) *FrontendMode {
		return &v
	}).(FrontendModePtrOutput)
}

func (o FrontendModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FrontendModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FrontendMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FrontendModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FrontendModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FrontendMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FrontendModePtrOutput struct{ *pulumi.OutputState }

func (FrontendModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontendMode)(nil)).Elem()
}

func (o FrontendModePtrOutput) ToFrontendModePtrOutput() FrontendModePtrOutput {
	return o
}

func (o FrontendModePtrOutput) ToFrontendModePtrOutputWithContext(ctx context.Context) FrontendModePtrOutput {
	return o
}

func (o FrontendModePtrOutput) Elem() FrontendModeOutput {
	return o.ApplyT(func(v *FrontendMode) FrontendMode {
		if v != nil {
			return *v
		}
		var ret FrontendMode
		return ret
	}).(FrontendModeOutput)
}

func (o FrontendModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FrontendModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FrontendMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FrontendModeInput interface {
	pulumi.Input

	ToFrontendModeOutput() FrontendModeOutput
	ToFrontendModeOutputWithContext(context.Context) FrontendModeOutput
}

var frontendModePtrType = reflect.TypeOf((**FrontendMode)(nil)).Elem()

type FrontendModePtrInput interface {
	pulumi.Input

	ToFrontendModePtrOutput() FrontendModePtrOutput
	ToFrontendModePtrOutputWithContext(context.Context) FrontendModePtrOutput
}

type frontendModePtr string

func FrontendModePtr(v string) FrontendModePtrInput {
	return (*frontendModePtr)(&v)
}

func (*frontendModePtr) ElementType() reflect.Type {
	return frontendModePtrType
}

func (in *frontendModePtr) ToFrontendModePtrOutput() FrontendModePtrOutput {
	return pulumi.ToOutput(in).(FrontendModePtrOutput)
}

func (in *frontendModePtr) ToFrontendModePtrOutputWithContext(ctx context.Context) FrontendModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FrontendModePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AssociationTypeOutput{})
	pulumi.RegisterOutputType(AssociationTypePtrOutput{})
	pulumi.RegisterOutputType(FrontendIPAddressVersionOutput{})
	pulumi.RegisterOutputType(FrontendIPAddressVersionPtrOutput{})
	pulumi.RegisterOutputType(FrontendModeOutput{})
	pulumi.RegisterOutputType(FrontendModePtrOutput{})
}
