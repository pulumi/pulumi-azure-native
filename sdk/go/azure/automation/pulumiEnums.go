


package automation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContentSourceType string

const (
	ContentSourceTypeEmbeddedContent = ContentSourceType("embeddedContent")
	ContentSourceTypeUri             = ContentSourceType("uri")
)

func (ContentSourceType) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentSourceType)(nil)).Elem()
}

func (e ContentSourceType) ToContentSourceTypeOutput() ContentSourceTypeOutput {
	return pulumi.ToOutput(e).(ContentSourceTypeOutput)
}

func (e ContentSourceType) ToContentSourceTypeOutputWithContext(ctx context.Context) ContentSourceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ContentSourceTypeOutput)
}

func (e ContentSourceType) ToContentSourceTypePtrOutput() ContentSourceTypePtrOutput {
	return e.ToContentSourceTypePtrOutputWithContext(context.Background())
}

func (e ContentSourceType) ToContentSourceTypePtrOutputWithContext(ctx context.Context) ContentSourceTypePtrOutput {
	return ContentSourceType(e).ToContentSourceTypeOutputWithContext(ctx).ToContentSourceTypePtrOutputWithContext(ctx)
}

func (e ContentSourceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContentSourceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ContentSourceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ContentSourceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ContentSourceTypeOutput struct{ *pulumi.OutputState }

func (ContentSourceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentSourceType)(nil)).Elem()
}

func (o ContentSourceTypeOutput) ToContentSourceTypeOutput() ContentSourceTypeOutput {
	return o
}

func (o ContentSourceTypeOutput) ToContentSourceTypeOutputWithContext(ctx context.Context) ContentSourceTypeOutput {
	return o
}

func (o ContentSourceTypeOutput) ToContentSourceTypePtrOutput() ContentSourceTypePtrOutput {
	return o.ToContentSourceTypePtrOutputWithContext(context.Background())
}

func (o ContentSourceTypeOutput) ToContentSourceTypePtrOutputWithContext(ctx context.Context) ContentSourceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentSourceType) *ContentSourceType {
		return &v
	}).(ContentSourceTypePtrOutput)
}

func (o ContentSourceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ContentSourceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContentSourceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ContentSourceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContentSourceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ContentSourceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ContentSourceTypePtrOutput struct{ *pulumi.OutputState }

func (ContentSourceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentSourceType)(nil)).Elem()
}

func (o ContentSourceTypePtrOutput) ToContentSourceTypePtrOutput() ContentSourceTypePtrOutput {
	return o
}

func (o ContentSourceTypePtrOutput) ToContentSourceTypePtrOutputWithContext(ctx context.Context) ContentSourceTypePtrOutput {
	return o
}

func (o ContentSourceTypePtrOutput) Elem() ContentSourceTypeOutput {
	return o.ApplyT(func(v *ContentSourceType) ContentSourceType {
		if v != nil {
			return *v
		}
		var ret ContentSourceType
		return ret
	}).(ContentSourceTypeOutput)
}

func (o ContentSourceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ContentSourceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ContentSourceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ContentSourceTypeInput interface {
	pulumi.Input

	ToContentSourceTypeOutput() ContentSourceTypeOutput
	ToContentSourceTypeOutputWithContext(context.Context) ContentSourceTypeOutput
}

var contentSourceTypePtrType = reflect.TypeOf((**ContentSourceType)(nil)).Elem()

type ContentSourceTypePtrInput interface {
	pulumi.Input

	ToContentSourceTypePtrOutput() ContentSourceTypePtrOutput
	ToContentSourceTypePtrOutputWithContext(context.Context) ContentSourceTypePtrOutput
}

type contentSourceTypePtr string

func ContentSourceTypePtr(v string) ContentSourceTypePtrInput {
	return (*contentSourceTypePtr)(&v)
}

func (*contentSourceTypePtr) ElementType() reflect.Type {
	return contentSourceTypePtrType
}

func (in *contentSourceTypePtr) ToContentSourceTypePtrOutput() ContentSourceTypePtrOutput {
	return pulumi.ToOutput(in).(ContentSourceTypePtrOutput)
}

func (in *contentSourceTypePtr) ToContentSourceTypePtrOutputWithContext(ctx context.Context) ContentSourceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ContentSourceTypePtrOutput)
}

type EncryptionKeySourceType string

const (
	EncryptionKeySourceType_Microsoft_Automation = EncryptionKeySourceType("Microsoft.Automation")
	EncryptionKeySourceType_Microsoft_Keyvault   = EncryptionKeySourceType("Microsoft.Keyvault")
)

func (EncryptionKeySourceType) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionKeySourceType)(nil)).Elem()
}

func (e EncryptionKeySourceType) ToEncryptionKeySourceTypeOutput() EncryptionKeySourceTypeOutput {
	return pulumi.ToOutput(e).(EncryptionKeySourceTypeOutput)
}

func (e EncryptionKeySourceType) ToEncryptionKeySourceTypeOutputWithContext(ctx context.Context) EncryptionKeySourceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EncryptionKeySourceTypeOutput)
}

func (e EncryptionKeySourceType) ToEncryptionKeySourceTypePtrOutput() EncryptionKeySourceTypePtrOutput {
	return e.ToEncryptionKeySourceTypePtrOutputWithContext(context.Background())
}

func (e EncryptionKeySourceType) ToEncryptionKeySourceTypePtrOutputWithContext(ctx context.Context) EncryptionKeySourceTypePtrOutput {
	return EncryptionKeySourceType(e).ToEncryptionKeySourceTypeOutputWithContext(ctx).ToEncryptionKeySourceTypePtrOutputWithContext(ctx)
}

func (e EncryptionKeySourceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncryptionKeySourceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncryptionKeySourceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EncryptionKeySourceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EncryptionKeySourceTypeOutput struct{ *pulumi.OutputState }

func (EncryptionKeySourceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionKeySourceType)(nil)).Elem()
}

func (o EncryptionKeySourceTypeOutput) ToEncryptionKeySourceTypeOutput() EncryptionKeySourceTypeOutput {
	return o
}

func (o EncryptionKeySourceTypeOutput) ToEncryptionKeySourceTypeOutputWithContext(ctx context.Context) EncryptionKeySourceTypeOutput {
	return o
}

func (o EncryptionKeySourceTypeOutput) ToEncryptionKeySourceTypePtrOutput() EncryptionKeySourceTypePtrOutput {
	return o.ToEncryptionKeySourceTypePtrOutputWithContext(context.Background())
}

func (o EncryptionKeySourceTypeOutput) ToEncryptionKeySourceTypePtrOutputWithContext(ctx context.Context) EncryptionKeySourceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionKeySourceType) *EncryptionKeySourceType {
		return &v
	}).(EncryptionKeySourceTypePtrOutput)
}

func (o EncryptionKeySourceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EncryptionKeySourceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncryptionKeySourceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EncryptionKeySourceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncryptionKeySourceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncryptionKeySourceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EncryptionKeySourceTypePtrOutput struct{ *pulumi.OutputState }

func (EncryptionKeySourceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionKeySourceType)(nil)).Elem()
}

func (o EncryptionKeySourceTypePtrOutput) ToEncryptionKeySourceTypePtrOutput() EncryptionKeySourceTypePtrOutput {
	return o
}

func (o EncryptionKeySourceTypePtrOutput) ToEncryptionKeySourceTypePtrOutputWithContext(ctx context.Context) EncryptionKeySourceTypePtrOutput {
	return o
}

func (o EncryptionKeySourceTypePtrOutput) Elem() EncryptionKeySourceTypeOutput {
	return o.ApplyT(func(v *EncryptionKeySourceType) EncryptionKeySourceType {
		if v != nil {
			return *v
		}
		var ret EncryptionKeySourceType
		return ret
	}).(EncryptionKeySourceTypeOutput)
}

func (o EncryptionKeySourceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncryptionKeySourceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EncryptionKeySourceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EncryptionKeySourceTypeInput interface {
	pulumi.Input

	ToEncryptionKeySourceTypeOutput() EncryptionKeySourceTypeOutput
	ToEncryptionKeySourceTypeOutputWithContext(context.Context) EncryptionKeySourceTypeOutput
}

var encryptionKeySourceTypePtrType = reflect.TypeOf((**EncryptionKeySourceType)(nil)).Elem()

type EncryptionKeySourceTypePtrInput interface {
	pulumi.Input

	ToEncryptionKeySourceTypePtrOutput() EncryptionKeySourceTypePtrOutput
	ToEncryptionKeySourceTypePtrOutputWithContext(context.Context) EncryptionKeySourceTypePtrOutput
}

type encryptionKeySourceTypePtr string

func EncryptionKeySourceTypePtr(v string) EncryptionKeySourceTypePtrInput {
	return (*encryptionKeySourceTypePtr)(&v)
}

func (*encryptionKeySourceTypePtr) ElementType() reflect.Type {
	return encryptionKeySourceTypePtrType
}

func (in *encryptionKeySourceTypePtr) ToEncryptionKeySourceTypePtrOutput() EncryptionKeySourceTypePtrOutput {
	return pulumi.ToOutput(in).(EncryptionKeySourceTypePtrOutput)
}

func (in *encryptionKeySourceTypePtr) ToEncryptionKeySourceTypePtrOutputWithContext(ctx context.Context) EncryptionKeySourceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EncryptionKeySourceTypePtrOutput)
}

type LinuxUpdateClasses string

const (
	LinuxUpdateClassesUnclassified = LinuxUpdateClasses("Unclassified")
	LinuxUpdateClassesCritical     = LinuxUpdateClasses("Critical")
	LinuxUpdateClassesSecurity     = LinuxUpdateClasses("Security")
	LinuxUpdateClassesOther        = LinuxUpdateClasses("Other")
)

func (LinuxUpdateClasses) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxUpdateClasses)(nil)).Elem()
}

func (e LinuxUpdateClasses) ToLinuxUpdateClassesOutput() LinuxUpdateClassesOutput {
	return pulumi.ToOutput(e).(LinuxUpdateClassesOutput)
}

func (e LinuxUpdateClasses) ToLinuxUpdateClassesOutputWithContext(ctx context.Context) LinuxUpdateClassesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LinuxUpdateClassesOutput)
}

func (e LinuxUpdateClasses) ToLinuxUpdateClassesPtrOutput() LinuxUpdateClassesPtrOutput {
	return e.ToLinuxUpdateClassesPtrOutputWithContext(context.Background())
}

func (e LinuxUpdateClasses) ToLinuxUpdateClassesPtrOutputWithContext(ctx context.Context) LinuxUpdateClassesPtrOutput {
	return LinuxUpdateClasses(e).ToLinuxUpdateClassesOutputWithContext(ctx).ToLinuxUpdateClassesPtrOutputWithContext(ctx)
}

func (e LinuxUpdateClasses) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LinuxUpdateClasses) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LinuxUpdateClasses) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LinuxUpdateClasses) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LinuxUpdateClassesOutput struct{ *pulumi.OutputState }

func (LinuxUpdateClassesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxUpdateClasses)(nil)).Elem()
}

func (o LinuxUpdateClassesOutput) ToLinuxUpdateClassesOutput() LinuxUpdateClassesOutput {
	return o
}

func (o LinuxUpdateClassesOutput) ToLinuxUpdateClassesOutputWithContext(ctx context.Context) LinuxUpdateClassesOutput {
	return o
}

func (o LinuxUpdateClassesOutput) ToLinuxUpdateClassesPtrOutput() LinuxUpdateClassesPtrOutput {
	return o.ToLinuxUpdateClassesPtrOutputWithContext(context.Background())
}

func (o LinuxUpdateClassesOutput) ToLinuxUpdateClassesPtrOutputWithContext(ctx context.Context) LinuxUpdateClassesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinuxUpdateClasses) *LinuxUpdateClasses {
		return &v
	}).(LinuxUpdateClassesPtrOutput)
}

func (o LinuxUpdateClassesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LinuxUpdateClassesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LinuxUpdateClasses) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LinuxUpdateClassesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LinuxUpdateClassesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LinuxUpdateClasses) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LinuxUpdateClassesPtrOutput struct{ *pulumi.OutputState }

func (LinuxUpdateClassesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxUpdateClasses)(nil)).Elem()
}

func (o LinuxUpdateClassesPtrOutput) ToLinuxUpdateClassesPtrOutput() LinuxUpdateClassesPtrOutput {
	return o
}

func (o LinuxUpdateClassesPtrOutput) ToLinuxUpdateClassesPtrOutputWithContext(ctx context.Context) LinuxUpdateClassesPtrOutput {
	return o
}

func (o LinuxUpdateClassesPtrOutput) Elem() LinuxUpdateClassesOutput {
	return o.ApplyT(func(v *LinuxUpdateClasses) LinuxUpdateClasses {
		if v != nil {
			return *v
		}
		var ret LinuxUpdateClasses
		return ret
	}).(LinuxUpdateClassesOutput)
}

func (o LinuxUpdateClassesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LinuxUpdateClassesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LinuxUpdateClasses) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LinuxUpdateClassesInput interface {
	pulumi.Input

	ToLinuxUpdateClassesOutput() LinuxUpdateClassesOutput
	ToLinuxUpdateClassesOutputWithContext(context.Context) LinuxUpdateClassesOutput
}

var linuxUpdateClassesPtrType = reflect.TypeOf((**LinuxUpdateClasses)(nil)).Elem()

type LinuxUpdateClassesPtrInput interface {
	pulumi.Input

	ToLinuxUpdateClassesPtrOutput() LinuxUpdateClassesPtrOutput
	ToLinuxUpdateClassesPtrOutputWithContext(context.Context) LinuxUpdateClassesPtrOutput
}

type linuxUpdateClassesPtr string

func LinuxUpdateClassesPtr(v string) LinuxUpdateClassesPtrInput {
	return (*linuxUpdateClassesPtr)(&v)
}

func (*linuxUpdateClassesPtr) ElementType() reflect.Type {
	return linuxUpdateClassesPtrType
}

func (in *linuxUpdateClassesPtr) ToLinuxUpdateClassesPtrOutput() LinuxUpdateClassesPtrOutput {
	return pulumi.ToOutput(in).(LinuxUpdateClassesPtrOutput)
}

func (in *linuxUpdateClassesPtr) ToLinuxUpdateClassesPtrOutputWithContext(ctx context.Context) LinuxUpdateClassesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LinuxUpdateClassesPtrOutput)
}

type OperatingSystemType string

const (
	OperatingSystemTypeWindows = OperatingSystemType("Windows")
	OperatingSystemTypeLinux   = OperatingSystemType("Linux")
)

func (OperatingSystemType) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatingSystemType)(nil)).Elem()
}

func (e OperatingSystemType) ToOperatingSystemTypeOutput() OperatingSystemTypeOutput {
	return pulumi.ToOutput(e).(OperatingSystemTypeOutput)
}

func (e OperatingSystemType) ToOperatingSystemTypeOutputWithContext(ctx context.Context) OperatingSystemTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OperatingSystemTypeOutput)
}

func (e OperatingSystemType) ToOperatingSystemTypePtrOutput() OperatingSystemTypePtrOutput {
	return e.ToOperatingSystemTypePtrOutputWithContext(context.Background())
}

func (e OperatingSystemType) ToOperatingSystemTypePtrOutputWithContext(ctx context.Context) OperatingSystemTypePtrOutput {
	return OperatingSystemType(e).ToOperatingSystemTypeOutputWithContext(ctx).ToOperatingSystemTypePtrOutputWithContext(ctx)
}

func (e OperatingSystemType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatingSystemType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatingSystemType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OperatingSystemType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OperatingSystemTypeOutput struct{ *pulumi.OutputState }

func (OperatingSystemTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatingSystemType)(nil)).Elem()
}

func (o OperatingSystemTypeOutput) ToOperatingSystemTypeOutput() OperatingSystemTypeOutput {
	return o
}

func (o OperatingSystemTypeOutput) ToOperatingSystemTypeOutputWithContext(ctx context.Context) OperatingSystemTypeOutput {
	return o
}

func (o OperatingSystemTypeOutput) ToOperatingSystemTypePtrOutput() OperatingSystemTypePtrOutput {
	return o.ToOperatingSystemTypePtrOutputWithContext(context.Background())
}

func (o OperatingSystemTypeOutput) ToOperatingSystemTypePtrOutputWithContext(ctx context.Context) OperatingSystemTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OperatingSystemType) *OperatingSystemType {
		return &v
	}).(OperatingSystemTypePtrOutput)
}

func (o OperatingSystemTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OperatingSystemTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatingSystemType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OperatingSystemTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatingSystemTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatingSystemType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OperatingSystemTypePtrOutput struct{ *pulumi.OutputState }

func (OperatingSystemTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OperatingSystemType)(nil)).Elem()
}

func (o OperatingSystemTypePtrOutput) ToOperatingSystemTypePtrOutput() OperatingSystemTypePtrOutput {
	return o
}

func (o OperatingSystemTypePtrOutput) ToOperatingSystemTypePtrOutputWithContext(ctx context.Context) OperatingSystemTypePtrOutput {
	return o
}

func (o OperatingSystemTypePtrOutput) Elem() OperatingSystemTypeOutput {
	return o.ApplyT(func(v *OperatingSystemType) OperatingSystemType {
		if v != nil {
			return *v
		}
		var ret OperatingSystemType
		return ret
	}).(OperatingSystemTypeOutput)
}

func (o OperatingSystemTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatingSystemTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OperatingSystemType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OperatingSystemTypeInput interface {
	pulumi.Input

	ToOperatingSystemTypeOutput() OperatingSystemTypeOutput
	ToOperatingSystemTypeOutputWithContext(context.Context) OperatingSystemTypeOutput
}

var operatingSystemTypePtrType = reflect.TypeOf((**OperatingSystemType)(nil)).Elem()

type OperatingSystemTypePtrInput interface {
	pulumi.Input

	ToOperatingSystemTypePtrOutput() OperatingSystemTypePtrOutput
	ToOperatingSystemTypePtrOutputWithContext(context.Context) OperatingSystemTypePtrOutput
}

type operatingSystemTypePtr string

func OperatingSystemTypePtr(v string) OperatingSystemTypePtrInput {
	return (*operatingSystemTypePtr)(&v)
}

func (*operatingSystemTypePtr) ElementType() reflect.Type {
	return operatingSystemTypePtrType
}

func (in *operatingSystemTypePtr) ToOperatingSystemTypePtrOutput() OperatingSystemTypePtrOutput {
	return pulumi.ToOutput(in).(OperatingSystemTypePtrOutput)
}

func (in *operatingSystemTypePtr) ToOperatingSystemTypePtrOutputWithContext(ctx context.Context) OperatingSystemTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OperatingSystemTypePtrOutput)
}

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned, UserAssigned")
	ResourceIdentityTypeNone                         = ResourceIdentityType("None")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

type RunbookTypeEnum string

const (
	RunbookTypeEnumScript                  = RunbookTypeEnum("Script")
	RunbookTypeEnumGraph                   = RunbookTypeEnum("Graph")
	RunbookTypeEnumPowerShellWorkflow      = RunbookTypeEnum("PowerShellWorkflow")
	RunbookTypeEnumPowerShell              = RunbookTypeEnum("PowerShell")
	RunbookTypeEnumGraphPowerShellWorkflow = RunbookTypeEnum("GraphPowerShellWorkflow")
	RunbookTypeEnumGraphPowerShell         = RunbookTypeEnum("GraphPowerShell")
)

func (RunbookTypeEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookTypeEnum)(nil)).Elem()
}

func (e RunbookTypeEnum) ToRunbookTypeEnumOutput() RunbookTypeEnumOutput {
	return pulumi.ToOutput(e).(RunbookTypeEnumOutput)
}

func (e RunbookTypeEnum) ToRunbookTypeEnumOutputWithContext(ctx context.Context) RunbookTypeEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RunbookTypeEnumOutput)
}

func (e RunbookTypeEnum) ToRunbookTypeEnumPtrOutput() RunbookTypeEnumPtrOutput {
	return e.ToRunbookTypeEnumPtrOutputWithContext(context.Background())
}

func (e RunbookTypeEnum) ToRunbookTypeEnumPtrOutputWithContext(ctx context.Context) RunbookTypeEnumPtrOutput {
	return RunbookTypeEnum(e).ToRunbookTypeEnumOutputWithContext(ctx).ToRunbookTypeEnumPtrOutputWithContext(ctx)
}

func (e RunbookTypeEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RunbookTypeEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RunbookTypeEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RunbookTypeEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RunbookTypeEnumOutput struct{ *pulumi.OutputState }

func (RunbookTypeEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookTypeEnum)(nil)).Elem()
}

func (o RunbookTypeEnumOutput) ToRunbookTypeEnumOutput() RunbookTypeEnumOutput {
	return o
}

func (o RunbookTypeEnumOutput) ToRunbookTypeEnumOutputWithContext(ctx context.Context) RunbookTypeEnumOutput {
	return o
}

func (o RunbookTypeEnumOutput) ToRunbookTypeEnumPtrOutput() RunbookTypeEnumPtrOutput {
	return o.ToRunbookTypeEnumPtrOutputWithContext(context.Background())
}

func (o RunbookTypeEnumOutput) ToRunbookTypeEnumPtrOutputWithContext(ctx context.Context) RunbookTypeEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RunbookTypeEnum) *RunbookTypeEnum {
		return &v
	}).(RunbookTypeEnumPtrOutput)
}

func (o RunbookTypeEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RunbookTypeEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RunbookTypeEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RunbookTypeEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RunbookTypeEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RunbookTypeEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RunbookTypeEnumPtrOutput struct{ *pulumi.OutputState }

func (RunbookTypeEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunbookTypeEnum)(nil)).Elem()
}

func (o RunbookTypeEnumPtrOutput) ToRunbookTypeEnumPtrOutput() RunbookTypeEnumPtrOutput {
	return o
}

func (o RunbookTypeEnumPtrOutput) ToRunbookTypeEnumPtrOutputWithContext(ctx context.Context) RunbookTypeEnumPtrOutput {
	return o
}

func (o RunbookTypeEnumPtrOutput) Elem() RunbookTypeEnumOutput {
	return o.ApplyT(func(v *RunbookTypeEnum) RunbookTypeEnum {
		if v != nil {
			return *v
		}
		var ret RunbookTypeEnum
		return ret
	}).(RunbookTypeEnumOutput)
}

func (o RunbookTypeEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RunbookTypeEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RunbookTypeEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RunbookTypeEnumInput interface {
	pulumi.Input

	ToRunbookTypeEnumOutput() RunbookTypeEnumOutput
	ToRunbookTypeEnumOutputWithContext(context.Context) RunbookTypeEnumOutput
}

var runbookTypeEnumPtrType = reflect.TypeOf((**RunbookTypeEnum)(nil)).Elem()

type RunbookTypeEnumPtrInput interface {
	pulumi.Input

	ToRunbookTypeEnumPtrOutput() RunbookTypeEnumPtrOutput
	ToRunbookTypeEnumPtrOutputWithContext(context.Context) RunbookTypeEnumPtrOutput
}

type runbookTypeEnumPtr string

func RunbookTypeEnumPtr(v string) RunbookTypeEnumPtrInput {
	return (*runbookTypeEnumPtr)(&v)
}

func (*runbookTypeEnumPtr) ElementType() reflect.Type {
	return runbookTypeEnumPtrType
}

func (in *runbookTypeEnumPtr) ToRunbookTypeEnumPtrOutput() RunbookTypeEnumPtrOutput {
	return pulumi.ToOutput(in).(RunbookTypeEnumPtrOutput)
}

func (in *runbookTypeEnumPtr) ToRunbookTypeEnumPtrOutputWithContext(ctx context.Context) RunbookTypeEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RunbookTypeEnumPtrOutput)
}

type ScheduleDay string

const (
	ScheduleDayMonday    = ScheduleDay("Monday")
	ScheduleDayTuesday   = ScheduleDay("Tuesday")
	ScheduleDayWednesday = ScheduleDay("Wednesday")
	ScheduleDayThursday  = ScheduleDay("Thursday")
	ScheduleDayFriday    = ScheduleDay("Friday")
	ScheduleDaySaturday  = ScheduleDay("Saturday")
	ScheduleDaySunday    = ScheduleDay("Sunday")
)

func (ScheduleDay) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleDay)(nil)).Elem()
}

func (e ScheduleDay) ToScheduleDayOutput() ScheduleDayOutput {
	return pulumi.ToOutput(e).(ScheduleDayOutput)
}

func (e ScheduleDay) ToScheduleDayOutputWithContext(ctx context.Context) ScheduleDayOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScheduleDayOutput)
}

func (e ScheduleDay) ToScheduleDayPtrOutput() ScheduleDayPtrOutput {
	return e.ToScheduleDayPtrOutputWithContext(context.Background())
}

func (e ScheduleDay) ToScheduleDayPtrOutputWithContext(ctx context.Context) ScheduleDayPtrOutput {
	return ScheduleDay(e).ToScheduleDayOutputWithContext(ctx).ToScheduleDayPtrOutputWithContext(ctx)
}

func (e ScheduleDay) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScheduleDay) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScheduleDay) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScheduleDay) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScheduleDayOutput struct{ *pulumi.OutputState }

func (ScheduleDayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleDay)(nil)).Elem()
}

func (o ScheduleDayOutput) ToScheduleDayOutput() ScheduleDayOutput {
	return o
}

func (o ScheduleDayOutput) ToScheduleDayOutputWithContext(ctx context.Context) ScheduleDayOutput {
	return o
}

func (o ScheduleDayOutput) ToScheduleDayPtrOutput() ScheduleDayPtrOutput {
	return o.ToScheduleDayPtrOutputWithContext(context.Background())
}

func (o ScheduleDayOutput) ToScheduleDayPtrOutputWithContext(ctx context.Context) ScheduleDayPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScheduleDay) *ScheduleDay {
		return &v
	}).(ScheduleDayPtrOutput)
}

func (o ScheduleDayOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScheduleDayOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScheduleDay) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScheduleDayOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScheduleDayOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScheduleDay) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScheduleDayPtrOutput struct{ *pulumi.OutputState }

func (ScheduleDayPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleDay)(nil)).Elem()
}

func (o ScheduleDayPtrOutput) ToScheduleDayPtrOutput() ScheduleDayPtrOutput {
	return o
}

func (o ScheduleDayPtrOutput) ToScheduleDayPtrOutputWithContext(ctx context.Context) ScheduleDayPtrOutput {
	return o
}

func (o ScheduleDayPtrOutput) Elem() ScheduleDayOutput {
	return o.ApplyT(func(v *ScheduleDay) ScheduleDay {
		if v != nil {
			return *v
		}
		var ret ScheduleDay
		return ret
	}).(ScheduleDayOutput)
}

func (o ScheduleDayPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScheduleDayPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScheduleDay) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ScheduleDayInput interface {
	pulumi.Input

	ToScheduleDayOutput() ScheduleDayOutput
	ToScheduleDayOutputWithContext(context.Context) ScheduleDayOutput
}

var scheduleDayPtrType = reflect.TypeOf((**ScheduleDay)(nil)).Elem()

type ScheduleDayPtrInput interface {
	pulumi.Input

	ToScheduleDayPtrOutput() ScheduleDayPtrOutput
	ToScheduleDayPtrOutputWithContext(context.Context) ScheduleDayPtrOutput
}

type scheduleDayPtr string

func ScheduleDayPtr(v string) ScheduleDayPtrInput {
	return (*scheduleDayPtr)(&v)
}

func (*scheduleDayPtr) ElementType() reflect.Type {
	return scheduleDayPtrType
}

func (in *scheduleDayPtr) ToScheduleDayPtrOutput() ScheduleDayPtrOutput {
	return pulumi.ToOutput(in).(ScheduleDayPtrOutput)
}

func (in *scheduleDayPtr) ToScheduleDayPtrOutputWithContext(ctx context.Context) ScheduleDayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScheduleDayPtrOutput)
}

type ScheduleFrequency string

const (
	ScheduleFrequencyOneTime = ScheduleFrequency("OneTime")
	ScheduleFrequencyDay     = ScheduleFrequency("Day")
	ScheduleFrequencyHour    = ScheduleFrequency("Hour")
	ScheduleFrequencyWeek    = ScheduleFrequency("Week")
	ScheduleFrequencyMonth   = ScheduleFrequency("Month")
	// The minimum allowed interval for Minute schedules is 15 minutes.
	ScheduleFrequencyMinute = ScheduleFrequency("Minute")
)

func (ScheduleFrequency) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleFrequency)(nil)).Elem()
}

func (e ScheduleFrequency) ToScheduleFrequencyOutput() ScheduleFrequencyOutput {
	return pulumi.ToOutput(e).(ScheduleFrequencyOutput)
}

func (e ScheduleFrequency) ToScheduleFrequencyOutputWithContext(ctx context.Context) ScheduleFrequencyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScheduleFrequencyOutput)
}

func (e ScheduleFrequency) ToScheduleFrequencyPtrOutput() ScheduleFrequencyPtrOutput {
	return e.ToScheduleFrequencyPtrOutputWithContext(context.Background())
}

func (e ScheduleFrequency) ToScheduleFrequencyPtrOutputWithContext(ctx context.Context) ScheduleFrequencyPtrOutput {
	return ScheduleFrequency(e).ToScheduleFrequencyOutputWithContext(ctx).ToScheduleFrequencyPtrOutputWithContext(ctx)
}

func (e ScheduleFrequency) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScheduleFrequency) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScheduleFrequency) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScheduleFrequency) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScheduleFrequencyOutput struct{ *pulumi.OutputState }

func (ScheduleFrequencyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleFrequency)(nil)).Elem()
}

func (o ScheduleFrequencyOutput) ToScheduleFrequencyOutput() ScheduleFrequencyOutput {
	return o
}

func (o ScheduleFrequencyOutput) ToScheduleFrequencyOutputWithContext(ctx context.Context) ScheduleFrequencyOutput {
	return o
}

func (o ScheduleFrequencyOutput) ToScheduleFrequencyPtrOutput() ScheduleFrequencyPtrOutput {
	return o.ToScheduleFrequencyPtrOutputWithContext(context.Background())
}

func (o ScheduleFrequencyOutput) ToScheduleFrequencyPtrOutputWithContext(ctx context.Context) ScheduleFrequencyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScheduleFrequency) *ScheduleFrequency {
		return &v
	}).(ScheduleFrequencyPtrOutput)
}

func (o ScheduleFrequencyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScheduleFrequencyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScheduleFrequency) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScheduleFrequencyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScheduleFrequencyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScheduleFrequency) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScheduleFrequencyPtrOutput struct{ *pulumi.OutputState }

func (ScheduleFrequencyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleFrequency)(nil)).Elem()
}

func (o ScheduleFrequencyPtrOutput) ToScheduleFrequencyPtrOutput() ScheduleFrequencyPtrOutput {
	return o
}

func (o ScheduleFrequencyPtrOutput) ToScheduleFrequencyPtrOutputWithContext(ctx context.Context) ScheduleFrequencyPtrOutput {
	return o
}

func (o ScheduleFrequencyPtrOutput) Elem() ScheduleFrequencyOutput {
	return o.ApplyT(func(v *ScheduleFrequency) ScheduleFrequency {
		if v != nil {
			return *v
		}
		var ret ScheduleFrequency
		return ret
	}).(ScheduleFrequencyOutput)
}

func (o ScheduleFrequencyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScheduleFrequencyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScheduleFrequency) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ScheduleFrequencyInput interface {
	pulumi.Input

	ToScheduleFrequencyOutput() ScheduleFrequencyOutput
	ToScheduleFrequencyOutputWithContext(context.Context) ScheduleFrequencyOutput
}

var scheduleFrequencyPtrType = reflect.TypeOf((**ScheduleFrequency)(nil)).Elem()

type ScheduleFrequencyPtrInput interface {
	pulumi.Input

	ToScheduleFrequencyPtrOutput() ScheduleFrequencyPtrOutput
	ToScheduleFrequencyPtrOutputWithContext(context.Context) ScheduleFrequencyPtrOutput
}

type scheduleFrequencyPtr string

func ScheduleFrequencyPtr(v string) ScheduleFrequencyPtrInput {
	return (*scheduleFrequencyPtr)(&v)
}

func (*scheduleFrequencyPtr) ElementType() reflect.Type {
	return scheduleFrequencyPtrType
}

func (in *scheduleFrequencyPtr) ToScheduleFrequencyPtrOutput() ScheduleFrequencyPtrOutput {
	return pulumi.ToOutput(in).(ScheduleFrequencyPtrOutput)
}

func (in *scheduleFrequencyPtr) ToScheduleFrequencyPtrOutputWithContext(ctx context.Context) ScheduleFrequencyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScheduleFrequencyPtrOutput)
}

type SkuNameEnum string

const (
	SkuNameEnumFree  = SkuNameEnum("Free")
	SkuNameEnumBasic = SkuNameEnum("Basic")
)

func (SkuNameEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuNameEnum)(nil)).Elem()
}

func (e SkuNameEnum) ToSkuNameEnumOutput() SkuNameEnumOutput {
	return pulumi.ToOutput(e).(SkuNameEnumOutput)
}

func (e SkuNameEnum) ToSkuNameEnumOutputWithContext(ctx context.Context) SkuNameEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuNameEnumOutput)
}

func (e SkuNameEnum) ToSkuNameEnumPtrOutput() SkuNameEnumPtrOutput {
	return e.ToSkuNameEnumPtrOutputWithContext(context.Background())
}

func (e SkuNameEnum) ToSkuNameEnumPtrOutputWithContext(ctx context.Context) SkuNameEnumPtrOutput {
	return SkuNameEnum(e).ToSkuNameEnumOutputWithContext(ctx).ToSkuNameEnumPtrOutputWithContext(ctx)
}

func (e SkuNameEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuNameEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuNameEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuNameEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuNameEnumOutput struct{ *pulumi.OutputState }

func (SkuNameEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuNameEnum)(nil)).Elem()
}

func (o SkuNameEnumOutput) ToSkuNameEnumOutput() SkuNameEnumOutput {
	return o
}

func (o SkuNameEnumOutput) ToSkuNameEnumOutputWithContext(ctx context.Context) SkuNameEnumOutput {
	return o
}

func (o SkuNameEnumOutput) ToSkuNameEnumPtrOutput() SkuNameEnumPtrOutput {
	return o.ToSkuNameEnumPtrOutputWithContext(context.Background())
}

func (o SkuNameEnumOutput) ToSkuNameEnumPtrOutputWithContext(ctx context.Context) SkuNameEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuNameEnum) *SkuNameEnum {
		return &v
	}).(SkuNameEnumPtrOutput)
}

func (o SkuNameEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuNameEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuNameEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuNameEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNameEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuNameEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuNameEnumPtrOutput struct{ *pulumi.OutputState }

func (SkuNameEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuNameEnum)(nil)).Elem()
}

func (o SkuNameEnumPtrOutput) ToSkuNameEnumPtrOutput() SkuNameEnumPtrOutput {
	return o
}

func (o SkuNameEnumPtrOutput) ToSkuNameEnumPtrOutputWithContext(ctx context.Context) SkuNameEnumPtrOutput {
	return o
}

func (o SkuNameEnumPtrOutput) Elem() SkuNameEnumOutput {
	return o.ApplyT(func(v *SkuNameEnum) SkuNameEnum {
		if v != nil {
			return *v
		}
		var ret SkuNameEnum
		return ret
	}).(SkuNameEnumOutput)
}

func (o SkuNameEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNameEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuNameEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SkuNameEnumInput interface {
	pulumi.Input

	ToSkuNameEnumOutput() SkuNameEnumOutput
	ToSkuNameEnumOutputWithContext(context.Context) SkuNameEnumOutput
}

var skuNameEnumPtrType = reflect.TypeOf((**SkuNameEnum)(nil)).Elem()

type SkuNameEnumPtrInput interface {
	pulumi.Input

	ToSkuNameEnumPtrOutput() SkuNameEnumPtrOutput
	ToSkuNameEnumPtrOutputWithContext(context.Context) SkuNameEnumPtrOutput
}

type skuNameEnumPtr string

func SkuNameEnumPtr(v string) SkuNameEnumPtrInput {
	return (*skuNameEnumPtr)(&v)
}

func (*skuNameEnumPtr) ElementType() reflect.Type {
	return skuNameEnumPtrType
}

func (in *skuNameEnumPtr) ToSkuNameEnumPtrOutput() SkuNameEnumPtrOutput {
	return pulumi.ToOutput(in).(SkuNameEnumPtrOutput)
}

func (in *skuNameEnumPtr) ToSkuNameEnumPtrOutputWithContext(ctx context.Context) SkuNameEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuNameEnumPtrOutput)
}

type SourceType string

const (
	SourceTypeVsoGit  = SourceType("VsoGit")
	SourceTypeVsoTfvc = SourceType("VsoTfvc")
	SourceTypeGitHub  = SourceType("GitHub")
)

func (SourceType) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceType)(nil)).Elem()
}

func (e SourceType) ToSourceTypeOutput() SourceTypeOutput {
	return pulumi.ToOutput(e).(SourceTypeOutput)
}

func (e SourceType) ToSourceTypeOutputWithContext(ctx context.Context) SourceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SourceTypeOutput)
}

func (e SourceType) ToSourceTypePtrOutput() SourceTypePtrOutput {
	return e.ToSourceTypePtrOutputWithContext(context.Background())
}

func (e SourceType) ToSourceTypePtrOutputWithContext(ctx context.Context) SourceTypePtrOutput {
	return SourceType(e).ToSourceTypeOutputWithContext(ctx).ToSourceTypePtrOutputWithContext(ctx)
}

func (e SourceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SourceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SourceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SourceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SourceTypeOutput struct{ *pulumi.OutputState }

func (SourceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceType)(nil)).Elem()
}

func (o SourceTypeOutput) ToSourceTypeOutput() SourceTypeOutput {
	return o
}

func (o SourceTypeOutput) ToSourceTypeOutputWithContext(ctx context.Context) SourceTypeOutput {
	return o
}

func (o SourceTypeOutput) ToSourceTypePtrOutput() SourceTypePtrOutput {
	return o.ToSourceTypePtrOutputWithContext(context.Background())
}

func (o SourceTypeOutput) ToSourceTypePtrOutputWithContext(ctx context.Context) SourceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceType) *SourceType {
		return &v
	}).(SourceTypePtrOutput)
}

func (o SourceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SourceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SourceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SourceTypePtrOutput struct{ *pulumi.OutputState }

func (SourceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceType)(nil)).Elem()
}

func (o SourceTypePtrOutput) ToSourceTypePtrOutput() SourceTypePtrOutput {
	return o
}

func (o SourceTypePtrOutput) ToSourceTypePtrOutputWithContext(ctx context.Context) SourceTypePtrOutput {
	return o
}

func (o SourceTypePtrOutput) Elem() SourceTypeOutput {
	return o.ApplyT(func(v *SourceType) SourceType {
		if v != nil {
			return *v
		}
		var ret SourceType
		return ret
	}).(SourceTypeOutput)
}

func (o SourceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SourceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SourceTypeInput interface {
	pulumi.Input

	ToSourceTypeOutput() SourceTypeOutput
	ToSourceTypeOutputWithContext(context.Context) SourceTypeOutput
}

var sourceTypePtrType = reflect.TypeOf((**SourceType)(nil)).Elem()

type SourceTypePtrInput interface {
	pulumi.Input

	ToSourceTypePtrOutput() SourceTypePtrOutput
	ToSourceTypePtrOutputWithContext(context.Context) SourceTypePtrOutput
}

type sourceTypePtr string

func SourceTypePtr(v string) SourceTypePtrInput {
	return (*sourceTypePtr)(&v)
}

func (*sourceTypePtr) ElementType() reflect.Type {
	return sourceTypePtrType
}

func (in *sourceTypePtr) ToSourceTypePtrOutput() SourceTypePtrOutput {
	return pulumi.ToOutput(in).(SourceTypePtrOutput)
}

func (in *sourceTypePtr) ToSourceTypePtrOutputWithContext(ctx context.Context) SourceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SourceTypePtrOutput)
}

type TagOperators string

const (
	TagOperatorsAll = TagOperators("All")
	TagOperatorsAny = TagOperators("Any")
)

func (TagOperators) ElementType() reflect.Type {
	return reflect.TypeOf((*TagOperators)(nil)).Elem()
}

func (e TagOperators) ToTagOperatorsOutput() TagOperatorsOutput {
	return pulumi.ToOutput(e).(TagOperatorsOutput)
}

func (e TagOperators) ToTagOperatorsOutputWithContext(ctx context.Context) TagOperatorsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TagOperatorsOutput)
}

func (e TagOperators) ToTagOperatorsPtrOutput() TagOperatorsPtrOutput {
	return e.ToTagOperatorsPtrOutputWithContext(context.Background())
}

func (e TagOperators) ToTagOperatorsPtrOutputWithContext(ctx context.Context) TagOperatorsPtrOutput {
	return TagOperators(e).ToTagOperatorsOutputWithContext(ctx).ToTagOperatorsPtrOutputWithContext(ctx)
}

func (e TagOperators) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TagOperators) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TagOperators) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TagOperators) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TagOperatorsOutput struct{ *pulumi.OutputState }

func (TagOperatorsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagOperators)(nil)).Elem()
}

func (o TagOperatorsOutput) ToTagOperatorsOutput() TagOperatorsOutput {
	return o
}

func (o TagOperatorsOutput) ToTagOperatorsOutputWithContext(ctx context.Context) TagOperatorsOutput {
	return o
}

func (o TagOperatorsOutput) ToTagOperatorsPtrOutput() TagOperatorsPtrOutput {
	return o.ToTagOperatorsPtrOutputWithContext(context.Background())
}

func (o TagOperatorsOutput) ToTagOperatorsPtrOutputWithContext(ctx context.Context) TagOperatorsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TagOperators) *TagOperators {
		return &v
	}).(TagOperatorsPtrOutput)
}

func (o TagOperatorsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TagOperatorsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TagOperators) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TagOperatorsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TagOperatorsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TagOperators) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TagOperatorsPtrOutput struct{ *pulumi.OutputState }

func (TagOperatorsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagOperators)(nil)).Elem()
}

func (o TagOperatorsPtrOutput) ToTagOperatorsPtrOutput() TagOperatorsPtrOutput {
	return o
}

func (o TagOperatorsPtrOutput) ToTagOperatorsPtrOutputWithContext(ctx context.Context) TagOperatorsPtrOutput {
	return o
}

func (o TagOperatorsPtrOutput) Elem() TagOperatorsOutput {
	return o.ApplyT(func(v *TagOperators) TagOperators {
		if v != nil {
			return *v
		}
		var ret TagOperators
		return ret
	}).(TagOperatorsOutput)
}

func (o TagOperatorsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TagOperatorsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TagOperators) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TagOperatorsInput interface {
	pulumi.Input

	ToTagOperatorsOutput() TagOperatorsOutput
	ToTagOperatorsOutputWithContext(context.Context) TagOperatorsOutput
}

var tagOperatorsPtrType = reflect.TypeOf((**TagOperators)(nil)).Elem()

type TagOperatorsPtrInput interface {
	pulumi.Input

	ToTagOperatorsPtrOutput() TagOperatorsPtrOutput
	ToTagOperatorsPtrOutputWithContext(context.Context) TagOperatorsPtrOutput
}

type tagOperatorsPtr string

func TagOperatorsPtr(v string) TagOperatorsPtrInput {
	return (*tagOperatorsPtr)(&v)
}

func (*tagOperatorsPtr) ElementType() reflect.Type {
	return tagOperatorsPtrType
}

func (in *tagOperatorsPtr) ToTagOperatorsPtrOutput() TagOperatorsPtrOutput {
	return pulumi.ToOutput(in).(TagOperatorsPtrOutput)
}

func (in *tagOperatorsPtr) ToTagOperatorsPtrOutputWithContext(ctx context.Context) TagOperatorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TagOperatorsPtrOutput)
}

type TokenType string

const (
	TokenTypePersonalAccessToken = TokenType("PersonalAccessToken")
	TokenTypeOauth               = TokenType("Oauth")
)

func (TokenType) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenType)(nil)).Elem()
}

func (e TokenType) ToTokenTypeOutput() TokenTypeOutput {
	return pulumi.ToOutput(e).(TokenTypeOutput)
}

func (e TokenType) ToTokenTypeOutputWithContext(ctx context.Context) TokenTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TokenTypeOutput)
}

func (e TokenType) ToTokenTypePtrOutput() TokenTypePtrOutput {
	return e.ToTokenTypePtrOutputWithContext(context.Background())
}

func (e TokenType) ToTokenTypePtrOutputWithContext(ctx context.Context) TokenTypePtrOutput {
	return TokenType(e).ToTokenTypeOutputWithContext(ctx).ToTokenTypePtrOutputWithContext(ctx)
}

func (e TokenType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TokenType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TokenType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TokenType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TokenTypeOutput struct{ *pulumi.OutputState }

func (TokenTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenType)(nil)).Elem()
}

func (o TokenTypeOutput) ToTokenTypeOutput() TokenTypeOutput {
	return o
}

func (o TokenTypeOutput) ToTokenTypeOutputWithContext(ctx context.Context) TokenTypeOutput {
	return o
}

func (o TokenTypeOutput) ToTokenTypePtrOutput() TokenTypePtrOutput {
	return o.ToTokenTypePtrOutputWithContext(context.Background())
}

func (o TokenTypeOutput) ToTokenTypePtrOutputWithContext(ctx context.Context) TokenTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TokenType) *TokenType {
		return &v
	}).(TokenTypePtrOutput)
}

func (o TokenTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TokenTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TokenType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TokenTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TokenTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TokenType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TokenTypePtrOutput struct{ *pulumi.OutputState }

func (TokenTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenType)(nil)).Elem()
}

func (o TokenTypePtrOutput) ToTokenTypePtrOutput() TokenTypePtrOutput {
	return o
}

func (o TokenTypePtrOutput) ToTokenTypePtrOutputWithContext(ctx context.Context) TokenTypePtrOutput {
	return o
}

func (o TokenTypePtrOutput) Elem() TokenTypeOutput {
	return o.ApplyT(func(v *TokenType) TokenType {
		if v != nil {
			return *v
		}
		var ret TokenType
		return ret
	}).(TokenTypeOutput)
}

func (o TokenTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TokenTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TokenType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TokenTypeInput interface {
	pulumi.Input

	ToTokenTypeOutput() TokenTypeOutput
	ToTokenTypeOutputWithContext(context.Context) TokenTypeOutput
}

var tokenTypePtrType = reflect.TypeOf((**TokenType)(nil)).Elem()

type TokenTypePtrInput interface {
	pulumi.Input

	ToTokenTypePtrOutput() TokenTypePtrOutput
	ToTokenTypePtrOutputWithContext(context.Context) TokenTypePtrOutput
}

type tokenTypePtr string

func TokenTypePtr(v string) TokenTypePtrInput {
	return (*tokenTypePtr)(&v)
}

func (*tokenTypePtr) ElementType() reflect.Type {
	return tokenTypePtrType
}

func (in *tokenTypePtr) ToTokenTypePtrOutput() TokenTypePtrOutput {
	return pulumi.ToOutput(in).(TokenTypePtrOutput)
}

func (in *tokenTypePtr) ToTokenTypePtrOutputWithContext(ctx context.Context) TokenTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TokenTypePtrOutput)
}

type WindowsUpdateClasses string

const (
	WindowsUpdateClassesUnclassified = WindowsUpdateClasses("Unclassified")
	WindowsUpdateClassesCritical     = WindowsUpdateClasses("Critical")
	WindowsUpdateClassesSecurity     = WindowsUpdateClasses("Security")
	WindowsUpdateClassesUpdateRollup = WindowsUpdateClasses("UpdateRollup")
	WindowsUpdateClassesFeaturePack  = WindowsUpdateClasses("FeaturePack")
	WindowsUpdateClassesServicePack  = WindowsUpdateClasses("ServicePack")
	WindowsUpdateClassesDefinition   = WindowsUpdateClasses("Definition")
	WindowsUpdateClassesTools        = WindowsUpdateClasses("Tools")
	WindowsUpdateClassesUpdates      = WindowsUpdateClasses("Updates")
)

func (WindowsUpdateClasses) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsUpdateClasses)(nil)).Elem()
}

func (e WindowsUpdateClasses) ToWindowsUpdateClassesOutput() WindowsUpdateClassesOutput {
	return pulumi.ToOutput(e).(WindowsUpdateClassesOutput)
}

func (e WindowsUpdateClasses) ToWindowsUpdateClassesOutputWithContext(ctx context.Context) WindowsUpdateClassesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WindowsUpdateClassesOutput)
}

func (e WindowsUpdateClasses) ToWindowsUpdateClassesPtrOutput() WindowsUpdateClassesPtrOutput {
	return e.ToWindowsUpdateClassesPtrOutputWithContext(context.Background())
}

func (e WindowsUpdateClasses) ToWindowsUpdateClassesPtrOutputWithContext(ctx context.Context) WindowsUpdateClassesPtrOutput {
	return WindowsUpdateClasses(e).ToWindowsUpdateClassesOutputWithContext(ctx).ToWindowsUpdateClassesPtrOutputWithContext(ctx)
}

func (e WindowsUpdateClasses) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WindowsUpdateClasses) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WindowsUpdateClasses) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WindowsUpdateClasses) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WindowsUpdateClassesOutput struct{ *pulumi.OutputState }

func (WindowsUpdateClassesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsUpdateClasses)(nil)).Elem()
}

func (o WindowsUpdateClassesOutput) ToWindowsUpdateClassesOutput() WindowsUpdateClassesOutput {
	return o
}

func (o WindowsUpdateClassesOutput) ToWindowsUpdateClassesOutputWithContext(ctx context.Context) WindowsUpdateClassesOutput {
	return o
}

func (o WindowsUpdateClassesOutput) ToWindowsUpdateClassesPtrOutput() WindowsUpdateClassesPtrOutput {
	return o.ToWindowsUpdateClassesPtrOutputWithContext(context.Background())
}

func (o WindowsUpdateClassesOutput) ToWindowsUpdateClassesPtrOutputWithContext(ctx context.Context) WindowsUpdateClassesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WindowsUpdateClasses) *WindowsUpdateClasses {
		return &v
	}).(WindowsUpdateClassesPtrOutput)
}

func (o WindowsUpdateClassesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WindowsUpdateClassesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WindowsUpdateClasses) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WindowsUpdateClassesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WindowsUpdateClassesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WindowsUpdateClasses) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WindowsUpdateClassesPtrOutput struct{ *pulumi.OutputState }

func (WindowsUpdateClassesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsUpdateClasses)(nil)).Elem()
}

func (o WindowsUpdateClassesPtrOutput) ToWindowsUpdateClassesPtrOutput() WindowsUpdateClassesPtrOutput {
	return o
}

func (o WindowsUpdateClassesPtrOutput) ToWindowsUpdateClassesPtrOutputWithContext(ctx context.Context) WindowsUpdateClassesPtrOutput {
	return o
}

func (o WindowsUpdateClassesPtrOutput) Elem() WindowsUpdateClassesOutput {
	return o.ApplyT(func(v *WindowsUpdateClasses) WindowsUpdateClasses {
		if v != nil {
			return *v
		}
		var ret WindowsUpdateClasses
		return ret
	}).(WindowsUpdateClassesOutput)
}

func (o WindowsUpdateClassesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WindowsUpdateClassesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WindowsUpdateClasses) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WindowsUpdateClassesInput interface {
	pulumi.Input

	ToWindowsUpdateClassesOutput() WindowsUpdateClassesOutput
	ToWindowsUpdateClassesOutputWithContext(context.Context) WindowsUpdateClassesOutput
}

var windowsUpdateClassesPtrType = reflect.TypeOf((**WindowsUpdateClasses)(nil)).Elem()

type WindowsUpdateClassesPtrInput interface {
	pulumi.Input

	ToWindowsUpdateClassesPtrOutput() WindowsUpdateClassesPtrOutput
	ToWindowsUpdateClassesPtrOutputWithContext(context.Context) WindowsUpdateClassesPtrOutput
}

type windowsUpdateClassesPtr string

func WindowsUpdateClassesPtr(v string) WindowsUpdateClassesPtrInput {
	return (*windowsUpdateClassesPtr)(&v)
}

func (*windowsUpdateClassesPtr) ElementType() reflect.Type {
	return windowsUpdateClassesPtrType
}

func (in *windowsUpdateClassesPtr) ToWindowsUpdateClassesPtrOutput() WindowsUpdateClassesPtrOutput {
	return pulumi.ToOutput(in).(WindowsUpdateClassesPtrOutput)
}

func (in *windowsUpdateClassesPtr) ToWindowsUpdateClassesPtrOutputWithContext(ctx context.Context) WindowsUpdateClassesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WindowsUpdateClassesPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ContentSourceTypeOutput{})
	pulumi.RegisterOutputType(ContentSourceTypePtrOutput{})
	pulumi.RegisterOutputType(EncryptionKeySourceTypeOutput{})
	pulumi.RegisterOutputType(EncryptionKeySourceTypePtrOutput{})
	pulumi.RegisterOutputType(LinuxUpdateClassesOutput{})
	pulumi.RegisterOutputType(LinuxUpdateClassesPtrOutput{})
	pulumi.RegisterOutputType(OperatingSystemTypeOutput{})
	pulumi.RegisterOutputType(OperatingSystemTypePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(RunbookTypeEnumOutput{})
	pulumi.RegisterOutputType(RunbookTypeEnumPtrOutput{})
	pulumi.RegisterOutputType(ScheduleDayOutput{})
	pulumi.RegisterOutputType(ScheduleDayPtrOutput{})
	pulumi.RegisterOutputType(ScheduleFrequencyOutput{})
	pulumi.RegisterOutputType(ScheduleFrequencyPtrOutput{})
	pulumi.RegisterOutputType(SkuNameEnumOutput{})
	pulumi.RegisterOutputType(SkuNameEnumPtrOutput{})
	pulumi.RegisterOutputType(SourceTypeOutput{})
	pulumi.RegisterOutputType(SourceTypePtrOutput{})
	pulumi.RegisterOutputType(TagOperatorsOutput{})
	pulumi.RegisterOutputType(TagOperatorsPtrOutput{})
	pulumi.RegisterOutputType(TokenTypeOutput{})
	pulumi.RegisterOutputType(TokenTypePtrOutput{})
	pulumi.RegisterOutputType(WindowsUpdateClassesOutput{})
	pulumi.RegisterOutputType(WindowsUpdateClassesPtrOutput{})
}
