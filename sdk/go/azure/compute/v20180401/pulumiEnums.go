


package v20180401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CachingTypes string

const (
	CachingTypesNone      = CachingTypes("None")
	CachingTypesReadOnly  = CachingTypes("ReadOnly")
	CachingTypesReadWrite = CachingTypes("ReadWrite")
)

func (CachingTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*CachingTypes)(nil)).Elem()
}

func (e CachingTypes) ToCachingTypesOutput() CachingTypesOutput {
	return pulumi.ToOutput(e).(CachingTypesOutput)
}

func (e CachingTypes) ToCachingTypesOutputWithContext(ctx context.Context) CachingTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CachingTypesOutput)
}

func (e CachingTypes) ToCachingTypesPtrOutput() CachingTypesPtrOutput {
	return e.ToCachingTypesPtrOutputWithContext(context.Background())
}

func (e CachingTypes) ToCachingTypesPtrOutputWithContext(ctx context.Context) CachingTypesPtrOutput {
	return CachingTypes(e).ToCachingTypesOutputWithContext(ctx).ToCachingTypesPtrOutputWithContext(ctx)
}

func (e CachingTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CachingTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CachingTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CachingTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CachingTypesOutput struct{ *pulumi.OutputState }

func (CachingTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CachingTypes)(nil)).Elem()
}

func (o CachingTypesOutput) ToCachingTypesOutput() CachingTypesOutput {
	return o
}

func (o CachingTypesOutput) ToCachingTypesOutputWithContext(ctx context.Context) CachingTypesOutput {
	return o
}

func (o CachingTypesOutput) ToCachingTypesPtrOutput() CachingTypesPtrOutput {
	return o.ToCachingTypesPtrOutputWithContext(context.Background())
}

func (o CachingTypesOutput) ToCachingTypesPtrOutputWithContext(ctx context.Context) CachingTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CachingTypes) *CachingTypes {
		return &v
	}).(CachingTypesPtrOutput)
}

func (o CachingTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CachingTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CachingTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CachingTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CachingTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CachingTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CachingTypesPtrOutput struct{ *pulumi.OutputState }

func (CachingTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CachingTypes)(nil)).Elem()
}

func (o CachingTypesPtrOutput) ToCachingTypesPtrOutput() CachingTypesPtrOutput {
	return o
}

func (o CachingTypesPtrOutput) ToCachingTypesPtrOutputWithContext(ctx context.Context) CachingTypesPtrOutput {
	return o
}

func (o CachingTypesPtrOutput) Elem() CachingTypesOutput {
	return o.ApplyT(func(v *CachingTypes) CachingTypes {
		if v != nil {
			return *v
		}
		var ret CachingTypes
		return ret
	}).(CachingTypesOutput)
}

func (o CachingTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CachingTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CachingTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CachingTypesInput interface {
	pulumi.Input

	ToCachingTypesOutput() CachingTypesOutput
	ToCachingTypesOutputWithContext(context.Context) CachingTypesOutput
}

var cachingTypesPtrType = reflect.TypeOf((**CachingTypes)(nil)).Elem()

type CachingTypesPtrInput interface {
	pulumi.Input

	ToCachingTypesPtrOutput() CachingTypesPtrOutput
	ToCachingTypesPtrOutputWithContext(context.Context) CachingTypesPtrOutput
}

type cachingTypesPtr string

func CachingTypesPtr(v string) CachingTypesPtrInput {
	return (*cachingTypesPtr)(&v)
}

func (*cachingTypesPtr) ElementType() reflect.Type {
	return cachingTypesPtrType
}

func (in *cachingTypesPtr) ToCachingTypesPtrOutput() CachingTypesPtrOutput {
	return pulumi.ToOutput(in).(CachingTypesPtrOutput)
}

func (in *cachingTypesPtr) ToCachingTypesPtrOutputWithContext(ctx context.Context) CachingTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CachingTypesPtrOutput)
}

type ComponentNames string

const (
	ComponentNames_Microsoft_Windows_Shell_Setup = ComponentNames("Microsoft-Windows-Shell-Setup")
)

func (ComponentNames) ElementType() reflect.Type {
	return reflect.TypeOf((*ComponentNames)(nil)).Elem()
}

func (e ComponentNames) ToComponentNamesOutput() ComponentNamesOutput {
	return pulumi.ToOutput(e).(ComponentNamesOutput)
}

func (e ComponentNames) ToComponentNamesOutputWithContext(ctx context.Context) ComponentNamesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ComponentNamesOutput)
}

func (e ComponentNames) ToComponentNamesPtrOutput() ComponentNamesPtrOutput {
	return e.ToComponentNamesPtrOutputWithContext(context.Background())
}

func (e ComponentNames) ToComponentNamesPtrOutputWithContext(ctx context.Context) ComponentNamesPtrOutput {
	return ComponentNames(e).ToComponentNamesOutputWithContext(ctx).ToComponentNamesPtrOutputWithContext(ctx)
}

func (e ComponentNames) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComponentNames) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComponentNames) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ComponentNames) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ComponentNamesOutput struct{ *pulumi.OutputState }

func (ComponentNamesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComponentNames)(nil)).Elem()
}

func (o ComponentNamesOutput) ToComponentNamesOutput() ComponentNamesOutput {
	return o
}

func (o ComponentNamesOutput) ToComponentNamesOutputWithContext(ctx context.Context) ComponentNamesOutput {
	return o
}

func (o ComponentNamesOutput) ToComponentNamesPtrOutput() ComponentNamesPtrOutput {
	return o.ToComponentNamesPtrOutputWithContext(context.Background())
}

func (o ComponentNamesOutput) ToComponentNamesPtrOutputWithContext(ctx context.Context) ComponentNamesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComponentNames) *ComponentNames {
		return &v
	}).(ComponentNamesPtrOutput)
}

func (o ComponentNamesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ComponentNamesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComponentNames) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ComponentNamesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComponentNamesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComponentNames) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ComponentNamesPtrOutput struct{ *pulumi.OutputState }

func (ComponentNamesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComponentNames)(nil)).Elem()
}

func (o ComponentNamesPtrOutput) ToComponentNamesPtrOutput() ComponentNamesPtrOutput {
	return o
}

func (o ComponentNamesPtrOutput) ToComponentNamesPtrOutputWithContext(ctx context.Context) ComponentNamesPtrOutput {
	return o
}

func (o ComponentNamesPtrOutput) Elem() ComponentNamesOutput {
	return o.ApplyT(func(v *ComponentNames) ComponentNames {
		if v != nil {
			return *v
		}
		var ret ComponentNames
		return ret
	}).(ComponentNamesOutput)
}

func (o ComponentNamesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComponentNamesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ComponentNames) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ComponentNamesInput interface {
	pulumi.Input

	ToComponentNamesOutput() ComponentNamesOutput
	ToComponentNamesOutputWithContext(context.Context) ComponentNamesOutput
}

var componentNamesPtrType = reflect.TypeOf((**ComponentNames)(nil)).Elem()

type ComponentNamesPtrInput interface {
	pulumi.Input

	ToComponentNamesPtrOutput() ComponentNamesPtrOutput
	ToComponentNamesPtrOutputWithContext(context.Context) ComponentNamesPtrOutput
}

type componentNamesPtr string

func ComponentNamesPtr(v string) ComponentNamesPtrInput {
	return (*componentNamesPtr)(&v)
}

func (*componentNamesPtr) ElementType() reflect.Type {
	return componentNamesPtrType
}

func (in *componentNamesPtr) ToComponentNamesPtrOutput() ComponentNamesPtrOutput {
	return pulumi.ToOutput(in).(ComponentNamesPtrOutput)
}

func (in *componentNamesPtr) ToComponentNamesPtrOutputWithContext(ctx context.Context) ComponentNamesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ComponentNamesPtrOutput)
}

type DiskCreateOption string

const (
	DiskCreateOptionEmpty     = DiskCreateOption("Empty")
	DiskCreateOptionAttach    = DiskCreateOption("Attach")
	DiskCreateOptionFromImage = DiskCreateOption("FromImage")
	DiskCreateOptionImport    = DiskCreateOption("Import")
	DiskCreateOptionCopy      = DiskCreateOption("Copy")
	DiskCreateOptionRestore   = DiskCreateOption("Restore")
)

type DiskCreateOptionTypes string

const (
	DiskCreateOptionTypesFromImage = DiskCreateOptionTypes("FromImage")
	DiskCreateOptionTypesEmpty     = DiskCreateOptionTypes("Empty")
	DiskCreateOptionTypesAttach    = DiskCreateOptionTypes("Attach")
)

type IPVersion string

const (
	IPVersionIPv4 = IPVersion("IPv4")
	IPVersionIPv6 = IPVersion("IPv6")
)

type IntervalInMins string

const (
	IntervalInMinsThreeMins  = IntervalInMins("ThreeMins")
	IntervalInMinsFiveMins   = IntervalInMins("FiveMins")
	IntervalInMinsThirtyMins = IntervalInMins("ThirtyMins")
	IntervalInMinsSixtyMins  = IntervalInMins("SixtyMins")
)

func (IntervalInMins) ElementType() reflect.Type {
	return reflect.TypeOf((*IntervalInMins)(nil)).Elem()
}

func (e IntervalInMins) ToIntervalInMinsOutput() IntervalInMinsOutput {
	return pulumi.ToOutput(e).(IntervalInMinsOutput)
}

func (e IntervalInMins) ToIntervalInMinsOutputWithContext(ctx context.Context) IntervalInMinsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IntervalInMinsOutput)
}

func (e IntervalInMins) ToIntervalInMinsPtrOutput() IntervalInMinsPtrOutput {
	return e.ToIntervalInMinsPtrOutputWithContext(context.Background())
}

func (e IntervalInMins) ToIntervalInMinsPtrOutputWithContext(ctx context.Context) IntervalInMinsPtrOutput {
	return IntervalInMins(e).ToIntervalInMinsOutputWithContext(ctx).ToIntervalInMinsPtrOutputWithContext(ctx)
}

func (e IntervalInMins) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IntervalInMins) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IntervalInMins) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IntervalInMins) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IntervalInMinsOutput struct{ *pulumi.OutputState }

func (IntervalInMinsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntervalInMins)(nil)).Elem()
}

func (o IntervalInMinsOutput) ToIntervalInMinsOutput() IntervalInMinsOutput {
	return o
}

func (o IntervalInMinsOutput) ToIntervalInMinsOutputWithContext(ctx context.Context) IntervalInMinsOutput {
	return o
}

func (o IntervalInMinsOutput) ToIntervalInMinsPtrOutput() IntervalInMinsPtrOutput {
	return o.ToIntervalInMinsPtrOutputWithContext(context.Background())
}

func (o IntervalInMinsOutput) ToIntervalInMinsPtrOutputWithContext(ctx context.Context) IntervalInMinsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IntervalInMins) *IntervalInMins {
		return &v
	}).(IntervalInMinsPtrOutput)
}

func (o IntervalInMinsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IntervalInMinsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IntervalInMins) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IntervalInMinsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IntervalInMinsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IntervalInMins) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IntervalInMinsPtrOutput struct{ *pulumi.OutputState }

func (IntervalInMinsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntervalInMins)(nil)).Elem()
}

func (o IntervalInMinsPtrOutput) ToIntervalInMinsPtrOutput() IntervalInMinsPtrOutput {
	return o
}

func (o IntervalInMinsPtrOutput) ToIntervalInMinsPtrOutputWithContext(ctx context.Context) IntervalInMinsPtrOutput {
	return o
}

func (o IntervalInMinsPtrOutput) Elem() IntervalInMinsOutput {
	return o.ApplyT(func(v *IntervalInMins) IntervalInMins {
		if v != nil {
			return *v
		}
		var ret IntervalInMins
		return ret
	}).(IntervalInMinsOutput)
}

func (o IntervalInMinsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IntervalInMinsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IntervalInMins) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IntervalInMinsInput interface {
	pulumi.Input

	ToIntervalInMinsOutput() IntervalInMinsOutput
	ToIntervalInMinsOutputWithContext(context.Context) IntervalInMinsOutput
}

var intervalInMinsPtrType = reflect.TypeOf((**IntervalInMins)(nil)).Elem()

type IntervalInMinsPtrInput interface {
	pulumi.Input

	ToIntervalInMinsPtrOutput() IntervalInMinsPtrOutput
	ToIntervalInMinsPtrOutputWithContext(context.Context) IntervalInMinsPtrOutput
}

type intervalInMinsPtr string

func IntervalInMinsPtr(v string) IntervalInMinsPtrInput {
	return (*intervalInMinsPtr)(&v)
}

func (*intervalInMinsPtr) ElementType() reflect.Type {
	return intervalInMinsPtrType
}

func (in *intervalInMinsPtr) ToIntervalInMinsPtrOutput() IntervalInMinsPtrOutput {
	return pulumi.ToOutput(in).(IntervalInMinsPtrOutput)
}

func (in *intervalInMinsPtr) ToIntervalInMinsPtrOutputWithContext(ctx context.Context) IntervalInMinsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IntervalInMinsPtrOutput)
}

type OperatingSystemStateTypes string

const (
	OperatingSystemStateTypesGeneralized = OperatingSystemStateTypes("Generalized")
	OperatingSystemStateTypesSpecialized = OperatingSystemStateTypes("Specialized")
)

func (OperatingSystemStateTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatingSystemStateTypes)(nil)).Elem()
}

func (e OperatingSystemStateTypes) ToOperatingSystemStateTypesOutput() OperatingSystemStateTypesOutput {
	return pulumi.ToOutput(e).(OperatingSystemStateTypesOutput)
}

func (e OperatingSystemStateTypes) ToOperatingSystemStateTypesOutputWithContext(ctx context.Context) OperatingSystemStateTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OperatingSystemStateTypesOutput)
}

func (e OperatingSystemStateTypes) ToOperatingSystemStateTypesPtrOutput() OperatingSystemStateTypesPtrOutput {
	return e.ToOperatingSystemStateTypesPtrOutputWithContext(context.Background())
}

func (e OperatingSystemStateTypes) ToOperatingSystemStateTypesPtrOutputWithContext(ctx context.Context) OperatingSystemStateTypesPtrOutput {
	return OperatingSystemStateTypes(e).ToOperatingSystemStateTypesOutputWithContext(ctx).ToOperatingSystemStateTypesPtrOutputWithContext(ctx)
}

func (e OperatingSystemStateTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatingSystemStateTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatingSystemStateTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OperatingSystemStateTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OperatingSystemStateTypesOutput struct{ *pulumi.OutputState }

func (OperatingSystemStateTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatingSystemStateTypes)(nil)).Elem()
}

func (o OperatingSystemStateTypesOutput) ToOperatingSystemStateTypesOutput() OperatingSystemStateTypesOutput {
	return o
}

func (o OperatingSystemStateTypesOutput) ToOperatingSystemStateTypesOutputWithContext(ctx context.Context) OperatingSystemStateTypesOutput {
	return o
}

func (o OperatingSystemStateTypesOutput) ToOperatingSystemStateTypesPtrOutput() OperatingSystemStateTypesPtrOutput {
	return o.ToOperatingSystemStateTypesPtrOutputWithContext(context.Background())
}

func (o OperatingSystemStateTypesOutput) ToOperatingSystemStateTypesPtrOutputWithContext(ctx context.Context) OperatingSystemStateTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OperatingSystemStateTypes) *OperatingSystemStateTypes {
		return &v
	}).(OperatingSystemStateTypesPtrOutput)
}

func (o OperatingSystemStateTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OperatingSystemStateTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatingSystemStateTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OperatingSystemStateTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatingSystemStateTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatingSystemStateTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OperatingSystemStateTypesPtrOutput struct{ *pulumi.OutputState }

func (OperatingSystemStateTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OperatingSystemStateTypes)(nil)).Elem()
}

func (o OperatingSystemStateTypesPtrOutput) ToOperatingSystemStateTypesPtrOutput() OperatingSystemStateTypesPtrOutput {
	return o
}

func (o OperatingSystemStateTypesPtrOutput) ToOperatingSystemStateTypesPtrOutputWithContext(ctx context.Context) OperatingSystemStateTypesPtrOutput {
	return o
}

func (o OperatingSystemStateTypesPtrOutput) Elem() OperatingSystemStateTypesOutput {
	return o.ApplyT(func(v *OperatingSystemStateTypes) OperatingSystemStateTypes {
		if v != nil {
			return *v
		}
		var ret OperatingSystemStateTypes
		return ret
	}).(OperatingSystemStateTypesOutput)
}

func (o OperatingSystemStateTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatingSystemStateTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OperatingSystemStateTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OperatingSystemStateTypesInput interface {
	pulumi.Input

	ToOperatingSystemStateTypesOutput() OperatingSystemStateTypesOutput
	ToOperatingSystemStateTypesOutputWithContext(context.Context) OperatingSystemStateTypesOutput
}

var operatingSystemStateTypesPtrType = reflect.TypeOf((**OperatingSystemStateTypes)(nil)).Elem()

type OperatingSystemStateTypesPtrInput interface {
	pulumi.Input

	ToOperatingSystemStateTypesPtrOutput() OperatingSystemStateTypesPtrOutput
	ToOperatingSystemStateTypesPtrOutputWithContext(context.Context) OperatingSystemStateTypesPtrOutput
}

type operatingSystemStateTypesPtr string

func OperatingSystemStateTypesPtr(v string) OperatingSystemStateTypesPtrInput {
	return (*operatingSystemStateTypesPtr)(&v)
}

func (*operatingSystemStateTypesPtr) ElementType() reflect.Type {
	return operatingSystemStateTypesPtrType
}

func (in *operatingSystemStateTypesPtr) ToOperatingSystemStateTypesPtrOutput() OperatingSystemStateTypesPtrOutput {
	return pulumi.ToOutput(in).(OperatingSystemStateTypesPtrOutput)
}

func (in *operatingSystemStateTypesPtr) ToOperatingSystemStateTypesPtrOutputWithContext(ctx context.Context) OperatingSystemStateTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OperatingSystemStateTypesPtrOutput)
}

type OperatingSystemTypes string

const (
	OperatingSystemTypesWindows = OperatingSystemTypes("Windows")
	OperatingSystemTypesLinux   = OperatingSystemTypes("Linux")
)

func (OperatingSystemTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatingSystemTypes)(nil)).Elem()
}

func (e OperatingSystemTypes) ToOperatingSystemTypesOutput() OperatingSystemTypesOutput {
	return pulumi.ToOutput(e).(OperatingSystemTypesOutput)
}

func (e OperatingSystemTypes) ToOperatingSystemTypesOutputWithContext(ctx context.Context) OperatingSystemTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OperatingSystemTypesOutput)
}

func (e OperatingSystemTypes) ToOperatingSystemTypesPtrOutput() OperatingSystemTypesPtrOutput {
	return e.ToOperatingSystemTypesPtrOutputWithContext(context.Background())
}

func (e OperatingSystemTypes) ToOperatingSystemTypesPtrOutputWithContext(ctx context.Context) OperatingSystemTypesPtrOutput {
	return OperatingSystemTypes(e).ToOperatingSystemTypesOutputWithContext(ctx).ToOperatingSystemTypesPtrOutputWithContext(ctx)
}

func (e OperatingSystemTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatingSystemTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatingSystemTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OperatingSystemTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OperatingSystemTypesOutput struct{ *pulumi.OutputState }

func (OperatingSystemTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatingSystemTypes)(nil)).Elem()
}

func (o OperatingSystemTypesOutput) ToOperatingSystemTypesOutput() OperatingSystemTypesOutput {
	return o
}

func (o OperatingSystemTypesOutput) ToOperatingSystemTypesOutputWithContext(ctx context.Context) OperatingSystemTypesOutput {
	return o
}

func (o OperatingSystemTypesOutput) ToOperatingSystemTypesPtrOutput() OperatingSystemTypesPtrOutput {
	return o.ToOperatingSystemTypesPtrOutputWithContext(context.Background())
}

func (o OperatingSystemTypesOutput) ToOperatingSystemTypesPtrOutputWithContext(ctx context.Context) OperatingSystemTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OperatingSystemTypes) *OperatingSystemTypes {
		return &v
	}).(OperatingSystemTypesPtrOutput)
}

func (o OperatingSystemTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OperatingSystemTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatingSystemTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OperatingSystemTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatingSystemTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatingSystemTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OperatingSystemTypesPtrOutput struct{ *pulumi.OutputState }

func (OperatingSystemTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OperatingSystemTypes)(nil)).Elem()
}

func (o OperatingSystemTypesPtrOutput) ToOperatingSystemTypesPtrOutput() OperatingSystemTypesPtrOutput {
	return o
}

func (o OperatingSystemTypesPtrOutput) ToOperatingSystemTypesPtrOutputWithContext(ctx context.Context) OperatingSystemTypesPtrOutput {
	return o
}

func (o OperatingSystemTypesPtrOutput) Elem() OperatingSystemTypesOutput {
	return o.ApplyT(func(v *OperatingSystemTypes) OperatingSystemTypes {
		if v != nil {
			return *v
		}
		var ret OperatingSystemTypes
		return ret
	}).(OperatingSystemTypesOutput)
}

func (o OperatingSystemTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatingSystemTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OperatingSystemTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OperatingSystemTypesInput interface {
	pulumi.Input

	ToOperatingSystemTypesOutput() OperatingSystemTypesOutput
	ToOperatingSystemTypesOutputWithContext(context.Context) OperatingSystemTypesOutput
}

var operatingSystemTypesPtrType = reflect.TypeOf((**OperatingSystemTypes)(nil)).Elem()

type OperatingSystemTypesPtrInput interface {
	pulumi.Input

	ToOperatingSystemTypesPtrOutput() OperatingSystemTypesPtrOutput
	ToOperatingSystemTypesPtrOutputWithContext(context.Context) OperatingSystemTypesPtrOutput
}

type operatingSystemTypesPtr string

func OperatingSystemTypesPtr(v string) OperatingSystemTypesPtrInput {
	return (*operatingSystemTypesPtr)(&v)
}

func (*operatingSystemTypesPtr) ElementType() reflect.Type {
	return operatingSystemTypesPtrType
}

func (in *operatingSystemTypesPtr) ToOperatingSystemTypesPtrOutput() OperatingSystemTypesPtrOutput {
	return pulumi.ToOutput(in).(OperatingSystemTypesPtrOutput)
}

func (in *operatingSystemTypesPtr) ToOperatingSystemTypesPtrOutputWithContext(ctx context.Context) OperatingSystemTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OperatingSystemTypesPtrOutput)
}

type PassNames string

const (
	PassNamesOobeSystem = PassNames("OobeSystem")
)

func (PassNames) ElementType() reflect.Type {
	return reflect.TypeOf((*PassNames)(nil)).Elem()
}

func (e PassNames) ToPassNamesOutput() PassNamesOutput {
	return pulumi.ToOutput(e).(PassNamesOutput)
}

func (e PassNames) ToPassNamesOutputWithContext(ctx context.Context) PassNamesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PassNamesOutput)
}

func (e PassNames) ToPassNamesPtrOutput() PassNamesPtrOutput {
	return e.ToPassNamesPtrOutputWithContext(context.Background())
}

func (e PassNames) ToPassNamesPtrOutputWithContext(ctx context.Context) PassNamesPtrOutput {
	return PassNames(e).ToPassNamesOutputWithContext(ctx).ToPassNamesPtrOutputWithContext(ctx)
}

func (e PassNames) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PassNames) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PassNames) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PassNames) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PassNamesOutput struct{ *pulumi.OutputState }

func (PassNamesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PassNames)(nil)).Elem()
}

func (o PassNamesOutput) ToPassNamesOutput() PassNamesOutput {
	return o
}

func (o PassNamesOutput) ToPassNamesOutputWithContext(ctx context.Context) PassNamesOutput {
	return o
}

func (o PassNamesOutput) ToPassNamesPtrOutput() PassNamesPtrOutput {
	return o.ToPassNamesPtrOutputWithContext(context.Background())
}

func (o PassNamesOutput) ToPassNamesPtrOutputWithContext(ctx context.Context) PassNamesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PassNames) *PassNames {
		return &v
	}).(PassNamesPtrOutput)
}

func (o PassNamesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PassNamesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PassNames) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PassNamesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PassNamesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PassNames) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PassNamesPtrOutput struct{ *pulumi.OutputState }

func (PassNamesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PassNames)(nil)).Elem()
}

func (o PassNamesPtrOutput) ToPassNamesPtrOutput() PassNamesPtrOutput {
	return o
}

func (o PassNamesPtrOutput) ToPassNamesPtrOutputWithContext(ctx context.Context) PassNamesPtrOutput {
	return o
}

func (o PassNamesPtrOutput) Elem() PassNamesOutput {
	return o.ApplyT(func(v *PassNames) PassNames {
		if v != nil {
			return *v
		}
		var ret PassNames
		return ret
	}).(PassNamesOutput)
}

func (o PassNamesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PassNamesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PassNames) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PassNamesInput interface {
	pulumi.Input

	ToPassNamesOutput() PassNamesOutput
	ToPassNamesOutputWithContext(context.Context) PassNamesOutput
}

var passNamesPtrType = reflect.TypeOf((**PassNames)(nil)).Elem()

type PassNamesPtrInput interface {
	pulumi.Input

	ToPassNamesPtrOutput() PassNamesPtrOutput
	ToPassNamesPtrOutputWithContext(context.Context) PassNamesPtrOutput
}

type passNamesPtr string

func PassNamesPtr(v string) PassNamesPtrInput {
	return (*passNamesPtr)(&v)
}

func (*passNamesPtr) ElementType() reflect.Type {
	return passNamesPtrType
}

func (in *passNamesPtr) ToPassNamesPtrOutput() PassNamesPtrOutput {
	return pulumi.ToOutput(in).(PassNamesPtrOutput)
}

func (in *passNamesPtr) ToPassNamesPtrOutputWithContext(ctx context.Context) PassNamesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PassNamesPtrOutput)
}

type ProtocolTypes string

const (
	ProtocolTypesHttp  = ProtocolTypes("Http")
	ProtocolTypesHttps = ProtocolTypes("Https")
)

func (ProtocolTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtocolTypes)(nil)).Elem()
}

func (e ProtocolTypes) ToProtocolTypesOutput() ProtocolTypesOutput {
	return pulumi.ToOutput(e).(ProtocolTypesOutput)
}

func (e ProtocolTypes) ToProtocolTypesOutputWithContext(ctx context.Context) ProtocolTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProtocolTypesOutput)
}

func (e ProtocolTypes) ToProtocolTypesPtrOutput() ProtocolTypesPtrOutput {
	return e.ToProtocolTypesPtrOutputWithContext(context.Background())
}

func (e ProtocolTypes) ToProtocolTypesPtrOutputWithContext(ctx context.Context) ProtocolTypesPtrOutput {
	return ProtocolTypes(e).ToProtocolTypesOutputWithContext(ctx).ToProtocolTypesPtrOutputWithContext(ctx)
}

func (e ProtocolTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProtocolTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProtocolTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProtocolTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProtocolTypesOutput struct{ *pulumi.OutputState }

func (ProtocolTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtocolTypes)(nil)).Elem()
}

func (o ProtocolTypesOutput) ToProtocolTypesOutput() ProtocolTypesOutput {
	return o
}

func (o ProtocolTypesOutput) ToProtocolTypesOutputWithContext(ctx context.Context) ProtocolTypesOutput {
	return o
}

func (o ProtocolTypesOutput) ToProtocolTypesPtrOutput() ProtocolTypesPtrOutput {
	return o.ToProtocolTypesPtrOutputWithContext(context.Background())
}

func (o ProtocolTypesOutput) ToProtocolTypesPtrOutputWithContext(ctx context.Context) ProtocolTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProtocolTypes) *ProtocolTypes {
		return &v
	}).(ProtocolTypesPtrOutput)
}

func (o ProtocolTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProtocolTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProtocolTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProtocolTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtocolTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProtocolTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProtocolTypesPtrOutput struct{ *pulumi.OutputState }

func (ProtocolTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtocolTypes)(nil)).Elem()
}

func (o ProtocolTypesPtrOutput) ToProtocolTypesPtrOutput() ProtocolTypesPtrOutput {
	return o
}

func (o ProtocolTypesPtrOutput) ToProtocolTypesPtrOutputWithContext(ctx context.Context) ProtocolTypesPtrOutput {
	return o
}

func (o ProtocolTypesPtrOutput) Elem() ProtocolTypesOutput {
	return o.ApplyT(func(v *ProtocolTypes) ProtocolTypes {
		if v != nil {
			return *v
		}
		var ret ProtocolTypes
		return ret
	}).(ProtocolTypesOutput)
}

func (o ProtocolTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtocolTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProtocolTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProtocolTypesInput interface {
	pulumi.Input

	ToProtocolTypesOutput() ProtocolTypesOutput
	ToProtocolTypesOutputWithContext(context.Context) ProtocolTypesOutput
}

var protocolTypesPtrType = reflect.TypeOf((**ProtocolTypes)(nil)).Elem()

type ProtocolTypesPtrInput interface {
	pulumi.Input

	ToProtocolTypesPtrOutput() ProtocolTypesPtrOutput
	ToProtocolTypesPtrOutputWithContext(context.Context) ProtocolTypesPtrOutput
}

type protocolTypesPtr string

func ProtocolTypesPtr(v string) ProtocolTypesPtrInput {
	return (*protocolTypesPtr)(&v)
}

func (*protocolTypesPtr) ElementType() reflect.Type {
	return protocolTypesPtrType
}

func (in *protocolTypesPtr) ToProtocolTypesPtrOutput() ProtocolTypesPtrOutput {
	return pulumi.ToOutput(in).(ProtocolTypesPtrOutput)
}

func (in *protocolTypesPtr) ToProtocolTypesPtrOutputWithContext(ctx context.Context) ProtocolTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProtocolTypesPtrOutput)
}

type ProximityPlacementGroupType string

const (
	ProximityPlacementGroupTypeStandard = ProximityPlacementGroupType("Standard")
	ProximityPlacementGroupTypeUltra    = ProximityPlacementGroupType("Ultra")
)

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

type SettingNames string

const (
	SettingNamesAutoLogon          = SettingNames("AutoLogon")
	SettingNamesFirstLogonCommands = SettingNames("FirstLogonCommands")
)

func (SettingNames) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingNames)(nil)).Elem()
}

func (e SettingNames) ToSettingNamesOutput() SettingNamesOutput {
	return pulumi.ToOutput(e).(SettingNamesOutput)
}

func (e SettingNames) ToSettingNamesOutputWithContext(ctx context.Context) SettingNamesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SettingNamesOutput)
}

func (e SettingNames) ToSettingNamesPtrOutput() SettingNamesPtrOutput {
	return e.ToSettingNamesPtrOutputWithContext(context.Background())
}

func (e SettingNames) ToSettingNamesPtrOutputWithContext(ctx context.Context) SettingNamesPtrOutput {
	return SettingNames(e).ToSettingNamesOutputWithContext(ctx).ToSettingNamesPtrOutputWithContext(ctx)
}

func (e SettingNames) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SettingNames) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SettingNames) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SettingNames) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SettingNamesOutput struct{ *pulumi.OutputState }

func (SettingNamesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingNames)(nil)).Elem()
}

func (o SettingNamesOutput) ToSettingNamesOutput() SettingNamesOutput {
	return o
}

func (o SettingNamesOutput) ToSettingNamesOutputWithContext(ctx context.Context) SettingNamesOutput {
	return o
}

func (o SettingNamesOutput) ToSettingNamesPtrOutput() SettingNamesPtrOutput {
	return o.ToSettingNamesPtrOutputWithContext(context.Background())
}

func (o SettingNamesOutput) ToSettingNamesPtrOutputWithContext(ctx context.Context) SettingNamesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SettingNames) *SettingNames {
		return &v
	}).(SettingNamesPtrOutput)
}

func (o SettingNamesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SettingNamesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SettingNames) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SettingNamesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SettingNamesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SettingNames) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SettingNamesPtrOutput struct{ *pulumi.OutputState }

func (SettingNamesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SettingNames)(nil)).Elem()
}

func (o SettingNamesPtrOutput) ToSettingNamesPtrOutput() SettingNamesPtrOutput {
	return o
}

func (o SettingNamesPtrOutput) ToSettingNamesPtrOutputWithContext(ctx context.Context) SettingNamesPtrOutput {
	return o
}

func (o SettingNamesPtrOutput) Elem() SettingNamesOutput {
	return o.ApplyT(func(v *SettingNames) SettingNames {
		if v != nil {
			return *v
		}
		var ret SettingNames
		return ret
	}).(SettingNamesOutput)
}

func (o SettingNamesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SettingNamesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SettingNames) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SettingNamesInput interface {
	pulumi.Input

	ToSettingNamesOutput() SettingNamesOutput
	ToSettingNamesOutputWithContext(context.Context) SettingNamesOutput
}

var settingNamesPtrType = reflect.TypeOf((**SettingNames)(nil)).Elem()

type SettingNamesPtrInput interface {
	pulumi.Input

	ToSettingNamesPtrOutput() SettingNamesPtrOutput
	ToSettingNamesPtrOutputWithContext(context.Context) SettingNamesPtrOutput
}

type settingNamesPtr string

func SettingNamesPtr(v string) SettingNamesPtrInput {
	return (*settingNamesPtr)(&v)
}

func (*settingNamesPtr) ElementType() reflect.Type {
	return settingNamesPtrType
}

func (in *settingNamesPtr) ToSettingNamesPtrOutput() SettingNamesPtrOutput {
	return pulumi.ToOutput(in).(SettingNamesPtrOutput)
}

func (in *settingNamesPtr) ToSettingNamesPtrOutputWithContext(ctx context.Context) SettingNamesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SettingNamesPtrOutput)
}

type SnapshotStorageAccountTypes string

const (
	SnapshotStorageAccountTypes_Standard_LRS = SnapshotStorageAccountTypes("Standard_LRS")
	SnapshotStorageAccountTypes_Premium_LRS  = SnapshotStorageAccountTypes("Premium_LRS")
	SnapshotStorageAccountTypes_Standard_ZRS = SnapshotStorageAccountTypes("Standard_ZRS")
)

type StatusLevelTypes string

const (
	StatusLevelTypesInfo    = StatusLevelTypes("Info")
	StatusLevelTypesWarning = StatusLevelTypes("Warning")
	StatusLevelTypesError   = StatusLevelTypes("Error")
)

func (StatusLevelTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusLevelTypes)(nil)).Elem()
}

func (e StatusLevelTypes) ToStatusLevelTypesOutput() StatusLevelTypesOutput {
	return pulumi.ToOutput(e).(StatusLevelTypesOutput)
}

func (e StatusLevelTypes) ToStatusLevelTypesOutputWithContext(ctx context.Context) StatusLevelTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StatusLevelTypesOutput)
}

func (e StatusLevelTypes) ToStatusLevelTypesPtrOutput() StatusLevelTypesPtrOutput {
	return e.ToStatusLevelTypesPtrOutputWithContext(context.Background())
}

func (e StatusLevelTypes) ToStatusLevelTypesPtrOutputWithContext(ctx context.Context) StatusLevelTypesPtrOutput {
	return StatusLevelTypes(e).ToStatusLevelTypesOutputWithContext(ctx).ToStatusLevelTypesPtrOutputWithContext(ctx)
}

func (e StatusLevelTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StatusLevelTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StatusLevelTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StatusLevelTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StatusLevelTypesOutput struct{ *pulumi.OutputState }

func (StatusLevelTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusLevelTypes)(nil)).Elem()
}

func (o StatusLevelTypesOutput) ToStatusLevelTypesOutput() StatusLevelTypesOutput {
	return o
}

func (o StatusLevelTypesOutput) ToStatusLevelTypesOutputWithContext(ctx context.Context) StatusLevelTypesOutput {
	return o
}

func (o StatusLevelTypesOutput) ToStatusLevelTypesPtrOutput() StatusLevelTypesPtrOutput {
	return o.ToStatusLevelTypesPtrOutputWithContext(context.Background())
}

func (o StatusLevelTypesOutput) ToStatusLevelTypesPtrOutputWithContext(ctx context.Context) StatusLevelTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StatusLevelTypes) *StatusLevelTypes {
		return &v
	}).(StatusLevelTypesPtrOutput)
}

func (o StatusLevelTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StatusLevelTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StatusLevelTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StatusLevelTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusLevelTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StatusLevelTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StatusLevelTypesPtrOutput struct{ *pulumi.OutputState }

func (StatusLevelTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StatusLevelTypes)(nil)).Elem()
}

func (o StatusLevelTypesPtrOutput) ToStatusLevelTypesPtrOutput() StatusLevelTypesPtrOutput {
	return o
}

func (o StatusLevelTypesPtrOutput) ToStatusLevelTypesPtrOutputWithContext(ctx context.Context) StatusLevelTypesPtrOutput {
	return o
}

func (o StatusLevelTypesPtrOutput) Elem() StatusLevelTypesOutput {
	return o.ApplyT(func(v *StatusLevelTypes) StatusLevelTypes {
		if v != nil {
			return *v
		}
		var ret StatusLevelTypes
		return ret
	}).(StatusLevelTypesOutput)
}

func (o StatusLevelTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusLevelTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StatusLevelTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StatusLevelTypesInput interface {
	pulumi.Input

	ToStatusLevelTypesOutput() StatusLevelTypesOutput
	ToStatusLevelTypesOutputWithContext(context.Context) StatusLevelTypesOutput
}

var statusLevelTypesPtrType = reflect.TypeOf((**StatusLevelTypes)(nil)).Elem()

type StatusLevelTypesPtrInput interface {
	pulumi.Input

	ToStatusLevelTypesPtrOutput() StatusLevelTypesPtrOutput
	ToStatusLevelTypesPtrOutputWithContext(context.Context) StatusLevelTypesPtrOutput
}

type statusLevelTypesPtr string

func StatusLevelTypesPtr(v string) StatusLevelTypesPtrInput {
	return (*statusLevelTypesPtr)(&v)
}

func (*statusLevelTypesPtr) ElementType() reflect.Type {
	return statusLevelTypesPtrType
}

func (in *statusLevelTypesPtr) ToStatusLevelTypesPtrOutput() StatusLevelTypesPtrOutput {
	return pulumi.ToOutput(in).(StatusLevelTypesPtrOutput)
}

func (in *statusLevelTypesPtr) ToStatusLevelTypesPtrOutputWithContext(ctx context.Context) StatusLevelTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StatusLevelTypesPtrOutput)
}

type StorageAccountTypes string

const (
	StorageAccountTypes_Standard_LRS    = StorageAccountTypes("Standard_LRS")
	StorageAccountTypes_Premium_LRS     = StorageAccountTypes("Premium_LRS")
	StorageAccountTypes_StandardSSD_LRS = StorageAccountTypes("StandardSSD_LRS")
)

type UpgradeMode string

const (
	UpgradeModeAutomatic = UpgradeMode("Automatic")
	UpgradeModeManual    = UpgradeMode("Manual")
	UpgradeModeRolling   = UpgradeMode("Rolling")
)

func (UpgradeMode) ElementType() reflect.Type {
	return reflect.TypeOf((*UpgradeMode)(nil)).Elem()
}

func (e UpgradeMode) ToUpgradeModeOutput() UpgradeModeOutput {
	return pulumi.ToOutput(e).(UpgradeModeOutput)
}

func (e UpgradeMode) ToUpgradeModeOutputWithContext(ctx context.Context) UpgradeModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(UpgradeModeOutput)
}

func (e UpgradeMode) ToUpgradeModePtrOutput() UpgradeModePtrOutput {
	return e.ToUpgradeModePtrOutputWithContext(context.Background())
}

func (e UpgradeMode) ToUpgradeModePtrOutputWithContext(ctx context.Context) UpgradeModePtrOutput {
	return UpgradeMode(e).ToUpgradeModeOutputWithContext(ctx).ToUpgradeModePtrOutputWithContext(ctx)
}

func (e UpgradeMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e UpgradeMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e UpgradeMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e UpgradeMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type UpgradeModeOutput struct{ *pulumi.OutputState }

func (UpgradeModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpgradeMode)(nil)).Elem()
}

func (o UpgradeModeOutput) ToUpgradeModeOutput() UpgradeModeOutput {
	return o
}

func (o UpgradeModeOutput) ToUpgradeModeOutputWithContext(ctx context.Context) UpgradeModeOutput {
	return o
}

func (o UpgradeModeOutput) ToUpgradeModePtrOutput() UpgradeModePtrOutput {
	return o.ToUpgradeModePtrOutputWithContext(context.Background())
}

func (o UpgradeModeOutput) ToUpgradeModePtrOutputWithContext(ctx context.Context) UpgradeModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UpgradeMode) *UpgradeMode {
		return &v
	}).(UpgradeModePtrOutput)
}

func (o UpgradeModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o UpgradeModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UpgradeMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o UpgradeModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UpgradeModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UpgradeMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type UpgradeModePtrOutput struct{ *pulumi.OutputState }

func (UpgradeModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpgradeMode)(nil)).Elem()
}

func (o UpgradeModePtrOutput) ToUpgradeModePtrOutput() UpgradeModePtrOutput {
	return o
}

func (o UpgradeModePtrOutput) ToUpgradeModePtrOutputWithContext(ctx context.Context) UpgradeModePtrOutput {
	return o
}

func (o UpgradeModePtrOutput) Elem() UpgradeModeOutput {
	return o.ApplyT(func(v *UpgradeMode) UpgradeMode {
		if v != nil {
			return *v
		}
		var ret UpgradeMode
		return ret
	}).(UpgradeModeOutput)
}

func (o UpgradeModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UpgradeModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *UpgradeMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type UpgradeModeInput interface {
	pulumi.Input

	ToUpgradeModeOutput() UpgradeModeOutput
	ToUpgradeModeOutputWithContext(context.Context) UpgradeModeOutput
}

var upgradeModePtrType = reflect.TypeOf((**UpgradeMode)(nil)).Elem()

type UpgradeModePtrInput interface {
	pulumi.Input

	ToUpgradeModePtrOutput() UpgradeModePtrOutput
	ToUpgradeModePtrOutputWithContext(context.Context) UpgradeModePtrOutput
}

type upgradeModePtr string

func UpgradeModePtr(v string) UpgradeModePtrInput {
	return (*upgradeModePtr)(&v)
}

func (*upgradeModePtr) ElementType() reflect.Type {
	return upgradeModePtrType
}

func (in *upgradeModePtr) ToUpgradeModePtrOutput() UpgradeModePtrOutput {
	return pulumi.ToOutput(in).(UpgradeModePtrOutput)
}

func (in *upgradeModePtr) ToUpgradeModePtrOutputWithContext(ctx context.Context) UpgradeModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(UpgradeModePtrOutput)
}

type VirtualMachineEvictionPolicyTypes string

const (
	VirtualMachineEvictionPolicyTypesDeallocate = VirtualMachineEvictionPolicyTypes("Deallocate")
	VirtualMachineEvictionPolicyTypesDelete     = VirtualMachineEvictionPolicyTypes("Delete")
)

type VirtualMachinePriorityTypes string

const (
	VirtualMachinePriorityTypesRegular = VirtualMachinePriorityTypes("Regular")
	VirtualMachinePriorityTypesLow     = VirtualMachinePriorityTypes("Low")
)

type VirtualMachineSizeTypes string

const (
	VirtualMachineSizeTypes_Basic_A0            = VirtualMachineSizeTypes("Basic_A0")
	VirtualMachineSizeTypes_Basic_A1            = VirtualMachineSizeTypes("Basic_A1")
	VirtualMachineSizeTypes_Basic_A2            = VirtualMachineSizeTypes("Basic_A2")
	VirtualMachineSizeTypes_Basic_A3            = VirtualMachineSizeTypes("Basic_A3")
	VirtualMachineSizeTypes_Basic_A4            = VirtualMachineSizeTypes("Basic_A4")
	VirtualMachineSizeTypes_Standard_A0         = VirtualMachineSizeTypes("Standard_A0")
	VirtualMachineSizeTypes_Standard_A1         = VirtualMachineSizeTypes("Standard_A1")
	VirtualMachineSizeTypes_Standard_A2         = VirtualMachineSizeTypes("Standard_A2")
	VirtualMachineSizeTypes_Standard_A3         = VirtualMachineSizeTypes("Standard_A3")
	VirtualMachineSizeTypes_Standard_A4         = VirtualMachineSizeTypes("Standard_A4")
	VirtualMachineSizeTypes_Standard_A5         = VirtualMachineSizeTypes("Standard_A5")
	VirtualMachineSizeTypes_Standard_A6         = VirtualMachineSizeTypes("Standard_A6")
	VirtualMachineSizeTypes_Standard_A7         = VirtualMachineSizeTypes("Standard_A7")
	VirtualMachineSizeTypes_Standard_A8         = VirtualMachineSizeTypes("Standard_A8")
	VirtualMachineSizeTypes_Standard_A9         = VirtualMachineSizeTypes("Standard_A9")
	VirtualMachineSizeTypes_Standard_A10        = VirtualMachineSizeTypes("Standard_A10")
	VirtualMachineSizeTypes_Standard_A11        = VirtualMachineSizeTypes("Standard_A11")
	VirtualMachineSizeTypes_Standard_A1_v2      = VirtualMachineSizeTypes("Standard_A1_v2")
	VirtualMachineSizeTypes_Standard_A2_v2      = VirtualMachineSizeTypes("Standard_A2_v2")
	VirtualMachineSizeTypes_Standard_A4_v2      = VirtualMachineSizeTypes("Standard_A4_v2")
	VirtualMachineSizeTypes_Standard_A8_v2      = VirtualMachineSizeTypes("Standard_A8_v2")
	VirtualMachineSizeTypes_Standard_A2m_v2     = VirtualMachineSizeTypes("Standard_A2m_v2")
	VirtualMachineSizeTypes_Standard_A4m_v2     = VirtualMachineSizeTypes("Standard_A4m_v2")
	VirtualMachineSizeTypes_Standard_A8m_v2     = VirtualMachineSizeTypes("Standard_A8m_v2")
	VirtualMachineSizeTypes_Standard_B1s        = VirtualMachineSizeTypes("Standard_B1s")
	VirtualMachineSizeTypes_Standard_B1ms       = VirtualMachineSizeTypes("Standard_B1ms")
	VirtualMachineSizeTypes_Standard_B2s        = VirtualMachineSizeTypes("Standard_B2s")
	VirtualMachineSizeTypes_Standard_B2ms       = VirtualMachineSizeTypes("Standard_B2ms")
	VirtualMachineSizeTypes_Standard_B4ms       = VirtualMachineSizeTypes("Standard_B4ms")
	VirtualMachineSizeTypes_Standard_B8ms       = VirtualMachineSizeTypes("Standard_B8ms")
	VirtualMachineSizeTypes_Standard_D1         = VirtualMachineSizeTypes("Standard_D1")
	VirtualMachineSizeTypes_Standard_D2         = VirtualMachineSizeTypes("Standard_D2")
	VirtualMachineSizeTypes_Standard_D3         = VirtualMachineSizeTypes("Standard_D3")
	VirtualMachineSizeTypes_Standard_D4         = VirtualMachineSizeTypes("Standard_D4")
	VirtualMachineSizeTypes_Standard_D11        = VirtualMachineSizeTypes("Standard_D11")
	VirtualMachineSizeTypes_Standard_D12        = VirtualMachineSizeTypes("Standard_D12")
	VirtualMachineSizeTypes_Standard_D13        = VirtualMachineSizeTypes("Standard_D13")
	VirtualMachineSizeTypes_Standard_D14        = VirtualMachineSizeTypes("Standard_D14")
	VirtualMachineSizeTypes_Standard_D1_v2      = VirtualMachineSizeTypes("Standard_D1_v2")
	VirtualMachineSizeTypes_Standard_D2_v2      = VirtualMachineSizeTypes("Standard_D2_v2")
	VirtualMachineSizeTypes_Standard_D3_v2      = VirtualMachineSizeTypes("Standard_D3_v2")
	VirtualMachineSizeTypes_Standard_D4_v2      = VirtualMachineSizeTypes("Standard_D4_v2")
	VirtualMachineSizeTypes_Standard_D5_v2      = VirtualMachineSizeTypes("Standard_D5_v2")
	VirtualMachineSizeTypes_Standard_D2_v3      = VirtualMachineSizeTypes("Standard_D2_v3")
	VirtualMachineSizeTypes_Standard_D4_v3      = VirtualMachineSizeTypes("Standard_D4_v3")
	VirtualMachineSizeTypes_Standard_D8_v3      = VirtualMachineSizeTypes("Standard_D8_v3")
	VirtualMachineSizeTypes_Standard_D16_v3     = VirtualMachineSizeTypes("Standard_D16_v3")
	VirtualMachineSizeTypes_Standard_D32_v3     = VirtualMachineSizeTypes("Standard_D32_v3")
	VirtualMachineSizeTypes_Standard_D64_v3     = VirtualMachineSizeTypes("Standard_D64_v3")
	VirtualMachineSizeTypes_Standard_D2s_v3     = VirtualMachineSizeTypes("Standard_D2s_v3")
	VirtualMachineSizeTypes_Standard_D4s_v3     = VirtualMachineSizeTypes("Standard_D4s_v3")
	VirtualMachineSizeTypes_Standard_D8s_v3     = VirtualMachineSizeTypes("Standard_D8s_v3")
	VirtualMachineSizeTypes_Standard_D16s_v3    = VirtualMachineSizeTypes("Standard_D16s_v3")
	VirtualMachineSizeTypes_Standard_D32s_v3    = VirtualMachineSizeTypes("Standard_D32s_v3")
	VirtualMachineSizeTypes_Standard_D64s_v3    = VirtualMachineSizeTypes("Standard_D64s_v3")
	VirtualMachineSizeTypes_Standard_D11_v2     = VirtualMachineSizeTypes("Standard_D11_v2")
	VirtualMachineSizeTypes_Standard_D12_v2     = VirtualMachineSizeTypes("Standard_D12_v2")
	VirtualMachineSizeTypes_Standard_D13_v2     = VirtualMachineSizeTypes("Standard_D13_v2")
	VirtualMachineSizeTypes_Standard_D14_v2     = VirtualMachineSizeTypes("Standard_D14_v2")
	VirtualMachineSizeTypes_Standard_D15_v2     = VirtualMachineSizeTypes("Standard_D15_v2")
	VirtualMachineSizeTypes_Standard_DS1        = VirtualMachineSizeTypes("Standard_DS1")
	VirtualMachineSizeTypes_Standard_DS2        = VirtualMachineSizeTypes("Standard_DS2")
	VirtualMachineSizeTypes_Standard_DS3        = VirtualMachineSizeTypes("Standard_DS3")
	VirtualMachineSizeTypes_Standard_DS4        = VirtualMachineSizeTypes("Standard_DS4")
	VirtualMachineSizeTypes_Standard_DS11       = VirtualMachineSizeTypes("Standard_DS11")
	VirtualMachineSizeTypes_Standard_DS12       = VirtualMachineSizeTypes("Standard_DS12")
	VirtualMachineSizeTypes_Standard_DS13       = VirtualMachineSizeTypes("Standard_DS13")
	VirtualMachineSizeTypes_Standard_DS14       = VirtualMachineSizeTypes("Standard_DS14")
	VirtualMachineSizeTypes_Standard_DS1_v2     = VirtualMachineSizeTypes("Standard_DS1_v2")
	VirtualMachineSizeTypes_Standard_DS2_v2     = VirtualMachineSizeTypes("Standard_DS2_v2")
	VirtualMachineSizeTypes_Standard_DS3_v2     = VirtualMachineSizeTypes("Standard_DS3_v2")
	VirtualMachineSizeTypes_Standard_DS4_v2     = VirtualMachineSizeTypes("Standard_DS4_v2")
	VirtualMachineSizeTypes_Standard_DS5_v2     = VirtualMachineSizeTypes("Standard_DS5_v2")
	VirtualMachineSizeTypes_Standard_DS11_v2    = VirtualMachineSizeTypes("Standard_DS11_v2")
	VirtualMachineSizeTypes_Standard_DS12_v2    = VirtualMachineSizeTypes("Standard_DS12_v2")
	VirtualMachineSizeTypes_Standard_DS13_v2    = VirtualMachineSizeTypes("Standard_DS13_v2")
	VirtualMachineSizeTypes_Standard_DS14_v2    = VirtualMachineSizeTypes("Standard_DS14_v2")
	VirtualMachineSizeTypes_Standard_DS15_v2    = VirtualMachineSizeTypes("Standard_DS15_v2")
	VirtualMachineSizeTypes_Standard_DS13_4_v2  = VirtualMachineSizeTypes("Standard_DS13-4_v2")
	VirtualMachineSizeTypes_Standard_DS13_2_v2  = VirtualMachineSizeTypes("Standard_DS13-2_v2")
	VirtualMachineSizeTypes_Standard_DS14_8_v2  = VirtualMachineSizeTypes("Standard_DS14-8_v2")
	VirtualMachineSizeTypes_Standard_DS14_4_v2  = VirtualMachineSizeTypes("Standard_DS14-4_v2")
	VirtualMachineSizeTypes_Standard_E2_v3      = VirtualMachineSizeTypes("Standard_E2_v3")
	VirtualMachineSizeTypes_Standard_E4_v3      = VirtualMachineSizeTypes("Standard_E4_v3")
	VirtualMachineSizeTypes_Standard_E8_v3      = VirtualMachineSizeTypes("Standard_E8_v3")
	VirtualMachineSizeTypes_Standard_E16_v3     = VirtualMachineSizeTypes("Standard_E16_v3")
	VirtualMachineSizeTypes_Standard_E32_v3     = VirtualMachineSizeTypes("Standard_E32_v3")
	VirtualMachineSizeTypes_Standard_E64_v3     = VirtualMachineSizeTypes("Standard_E64_v3")
	VirtualMachineSizeTypes_Standard_E2s_v3     = VirtualMachineSizeTypes("Standard_E2s_v3")
	VirtualMachineSizeTypes_Standard_E4s_v3     = VirtualMachineSizeTypes("Standard_E4s_v3")
	VirtualMachineSizeTypes_Standard_E8s_v3     = VirtualMachineSizeTypes("Standard_E8s_v3")
	VirtualMachineSizeTypes_Standard_E16s_v3    = VirtualMachineSizeTypes("Standard_E16s_v3")
	VirtualMachineSizeTypes_Standard_E32s_v3    = VirtualMachineSizeTypes("Standard_E32s_v3")
	VirtualMachineSizeTypes_Standard_E64s_v3    = VirtualMachineSizeTypes("Standard_E64s_v3")
	VirtualMachineSizeTypes_Standard_E32_16_v3  = VirtualMachineSizeTypes("Standard_E32-16_v3")
	VirtualMachineSizeTypes_Standard_E32_8s_v3  = VirtualMachineSizeTypes("Standard_E32-8s_v3")
	VirtualMachineSizeTypes_Standard_E64_32s_v3 = VirtualMachineSizeTypes("Standard_E64-32s_v3")
	VirtualMachineSizeTypes_Standard_E64_16s_v3 = VirtualMachineSizeTypes("Standard_E64-16s_v3")
	VirtualMachineSizeTypes_Standard_F1         = VirtualMachineSizeTypes("Standard_F1")
	VirtualMachineSizeTypes_Standard_F2         = VirtualMachineSizeTypes("Standard_F2")
	VirtualMachineSizeTypes_Standard_F4         = VirtualMachineSizeTypes("Standard_F4")
	VirtualMachineSizeTypes_Standard_F8         = VirtualMachineSizeTypes("Standard_F8")
	VirtualMachineSizeTypes_Standard_F16        = VirtualMachineSizeTypes("Standard_F16")
	VirtualMachineSizeTypes_Standard_F1s        = VirtualMachineSizeTypes("Standard_F1s")
	VirtualMachineSizeTypes_Standard_F2s        = VirtualMachineSizeTypes("Standard_F2s")
	VirtualMachineSizeTypes_Standard_F4s        = VirtualMachineSizeTypes("Standard_F4s")
	VirtualMachineSizeTypes_Standard_F8s        = VirtualMachineSizeTypes("Standard_F8s")
	VirtualMachineSizeTypes_Standard_F16s       = VirtualMachineSizeTypes("Standard_F16s")
	VirtualMachineSizeTypes_Standard_F2s_v2     = VirtualMachineSizeTypes("Standard_F2s_v2")
	VirtualMachineSizeTypes_Standard_F4s_v2     = VirtualMachineSizeTypes("Standard_F4s_v2")
	VirtualMachineSizeTypes_Standard_F8s_v2     = VirtualMachineSizeTypes("Standard_F8s_v2")
	VirtualMachineSizeTypes_Standard_F16s_v2    = VirtualMachineSizeTypes("Standard_F16s_v2")
	VirtualMachineSizeTypes_Standard_F32s_v2    = VirtualMachineSizeTypes("Standard_F32s_v2")
	VirtualMachineSizeTypes_Standard_F64s_v2    = VirtualMachineSizeTypes("Standard_F64s_v2")
	VirtualMachineSizeTypes_Standard_F72s_v2    = VirtualMachineSizeTypes("Standard_F72s_v2")
	VirtualMachineSizeTypes_Standard_G1         = VirtualMachineSizeTypes("Standard_G1")
	VirtualMachineSizeTypes_Standard_G2         = VirtualMachineSizeTypes("Standard_G2")
	VirtualMachineSizeTypes_Standard_G3         = VirtualMachineSizeTypes("Standard_G3")
	VirtualMachineSizeTypes_Standard_G4         = VirtualMachineSizeTypes("Standard_G4")
	VirtualMachineSizeTypes_Standard_G5         = VirtualMachineSizeTypes("Standard_G5")
	VirtualMachineSizeTypes_Standard_GS1        = VirtualMachineSizeTypes("Standard_GS1")
	VirtualMachineSizeTypes_Standard_GS2        = VirtualMachineSizeTypes("Standard_GS2")
	VirtualMachineSizeTypes_Standard_GS3        = VirtualMachineSizeTypes("Standard_GS3")
	VirtualMachineSizeTypes_Standard_GS4        = VirtualMachineSizeTypes("Standard_GS4")
	VirtualMachineSizeTypes_Standard_GS5        = VirtualMachineSizeTypes("Standard_GS5")
	VirtualMachineSizeTypes_Standard_GS4_8      = VirtualMachineSizeTypes("Standard_GS4-8")
	VirtualMachineSizeTypes_Standard_GS4_4      = VirtualMachineSizeTypes("Standard_GS4-4")
	VirtualMachineSizeTypes_Standard_GS5_16     = VirtualMachineSizeTypes("Standard_GS5-16")
	VirtualMachineSizeTypes_Standard_GS5_8      = VirtualMachineSizeTypes("Standard_GS5-8")
	VirtualMachineSizeTypes_Standard_H8         = VirtualMachineSizeTypes("Standard_H8")
	VirtualMachineSizeTypes_Standard_H16        = VirtualMachineSizeTypes("Standard_H16")
	VirtualMachineSizeTypes_Standard_H8m        = VirtualMachineSizeTypes("Standard_H8m")
	VirtualMachineSizeTypes_Standard_H16m       = VirtualMachineSizeTypes("Standard_H16m")
	VirtualMachineSizeTypes_Standard_H16r       = VirtualMachineSizeTypes("Standard_H16r")
	VirtualMachineSizeTypes_Standard_H16mr      = VirtualMachineSizeTypes("Standard_H16mr")
	VirtualMachineSizeTypes_Standard_L4s        = VirtualMachineSizeTypes("Standard_L4s")
	VirtualMachineSizeTypes_Standard_L8s        = VirtualMachineSizeTypes("Standard_L8s")
	VirtualMachineSizeTypes_Standard_L16s       = VirtualMachineSizeTypes("Standard_L16s")
	VirtualMachineSizeTypes_Standard_L32s       = VirtualMachineSizeTypes("Standard_L32s")
	VirtualMachineSizeTypes_Standard_M64s       = VirtualMachineSizeTypes("Standard_M64s")
	VirtualMachineSizeTypes_Standard_M64ms      = VirtualMachineSizeTypes("Standard_M64ms")
	VirtualMachineSizeTypes_Standard_M128s      = VirtualMachineSizeTypes("Standard_M128s")
	VirtualMachineSizeTypes_Standard_M128ms     = VirtualMachineSizeTypes("Standard_M128ms")
	VirtualMachineSizeTypes_Standard_M64_32ms   = VirtualMachineSizeTypes("Standard_M64-32ms")
	VirtualMachineSizeTypes_Standard_M64_16ms   = VirtualMachineSizeTypes("Standard_M64-16ms")
	VirtualMachineSizeTypes_Standard_M128_64ms  = VirtualMachineSizeTypes("Standard_M128-64ms")
	VirtualMachineSizeTypes_Standard_M128_32ms  = VirtualMachineSizeTypes("Standard_M128-32ms")
	VirtualMachineSizeTypes_Standard_NC6        = VirtualMachineSizeTypes("Standard_NC6")
	VirtualMachineSizeTypes_Standard_NC12       = VirtualMachineSizeTypes("Standard_NC12")
	VirtualMachineSizeTypes_Standard_NC24       = VirtualMachineSizeTypes("Standard_NC24")
	VirtualMachineSizeTypes_Standard_NC24r      = VirtualMachineSizeTypes("Standard_NC24r")
	VirtualMachineSizeTypes_Standard_NC6s_v2    = VirtualMachineSizeTypes("Standard_NC6s_v2")
	VirtualMachineSizeTypes_Standard_NC12s_v2   = VirtualMachineSizeTypes("Standard_NC12s_v2")
	VirtualMachineSizeTypes_Standard_NC24s_v2   = VirtualMachineSizeTypes("Standard_NC24s_v2")
	VirtualMachineSizeTypes_Standard_NC24rs_v2  = VirtualMachineSizeTypes("Standard_NC24rs_v2")
	VirtualMachineSizeTypes_Standard_NC6s_v3    = VirtualMachineSizeTypes("Standard_NC6s_v3")
	VirtualMachineSizeTypes_Standard_NC12s_v3   = VirtualMachineSizeTypes("Standard_NC12s_v3")
	VirtualMachineSizeTypes_Standard_NC24s_v3   = VirtualMachineSizeTypes("Standard_NC24s_v3")
	VirtualMachineSizeTypes_Standard_NC24rs_v3  = VirtualMachineSizeTypes("Standard_NC24rs_v3")
	VirtualMachineSizeTypes_Standard_ND6s       = VirtualMachineSizeTypes("Standard_ND6s")
	VirtualMachineSizeTypes_Standard_ND12s      = VirtualMachineSizeTypes("Standard_ND12s")
	VirtualMachineSizeTypes_Standard_ND24s      = VirtualMachineSizeTypes("Standard_ND24s")
	VirtualMachineSizeTypes_Standard_ND24rs     = VirtualMachineSizeTypes("Standard_ND24rs")
	VirtualMachineSizeTypes_Standard_NV6        = VirtualMachineSizeTypes("Standard_NV6")
	VirtualMachineSizeTypes_Standard_NV12       = VirtualMachineSizeTypes("Standard_NV12")
	VirtualMachineSizeTypes_Standard_NV24       = VirtualMachineSizeTypes("Standard_NV24")
)

func init() {
	pulumi.RegisterOutputType(CachingTypesOutput{})
	pulumi.RegisterOutputType(CachingTypesPtrOutput{})
	pulumi.RegisterOutputType(ComponentNamesOutput{})
	pulumi.RegisterOutputType(ComponentNamesPtrOutput{})
	pulumi.RegisterOutputType(IntervalInMinsOutput{})
	pulumi.RegisterOutputType(IntervalInMinsPtrOutput{})
	pulumi.RegisterOutputType(OperatingSystemStateTypesOutput{})
	pulumi.RegisterOutputType(OperatingSystemStateTypesPtrOutput{})
	pulumi.RegisterOutputType(OperatingSystemTypesOutput{})
	pulumi.RegisterOutputType(OperatingSystemTypesPtrOutput{})
	pulumi.RegisterOutputType(PassNamesOutput{})
	pulumi.RegisterOutputType(PassNamesPtrOutput{})
	pulumi.RegisterOutputType(ProtocolTypesOutput{})
	pulumi.RegisterOutputType(ProtocolTypesPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(SettingNamesOutput{})
	pulumi.RegisterOutputType(SettingNamesPtrOutput{})
	pulumi.RegisterOutputType(StatusLevelTypesOutput{})
	pulumi.RegisterOutputType(StatusLevelTypesPtrOutput{})
	pulumi.RegisterOutputType(UpgradeModeOutput{})
	pulumi.RegisterOutputType(UpgradeModePtrOutput{})
}
