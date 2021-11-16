


package v20200601

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

type DedicatedHostLicenseTypes string

const (
	DedicatedHostLicenseTypesNone                      = DedicatedHostLicenseTypes("None")
	DedicatedHostLicenseTypes_Windows_Server_Hybrid    = DedicatedHostLicenseTypes("Windows_Server_Hybrid")
	DedicatedHostLicenseTypes_Windows_Server_Perpetual = DedicatedHostLicenseTypes("Windows_Server_Perpetual")
)

func (DedicatedHostLicenseTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedHostLicenseTypes)(nil)).Elem()
}

func (e DedicatedHostLicenseTypes) ToDedicatedHostLicenseTypesOutput() DedicatedHostLicenseTypesOutput {
	return pulumi.ToOutput(e).(DedicatedHostLicenseTypesOutput)
}

func (e DedicatedHostLicenseTypes) ToDedicatedHostLicenseTypesOutputWithContext(ctx context.Context) DedicatedHostLicenseTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DedicatedHostLicenseTypesOutput)
}

func (e DedicatedHostLicenseTypes) ToDedicatedHostLicenseTypesPtrOutput() DedicatedHostLicenseTypesPtrOutput {
	return e.ToDedicatedHostLicenseTypesPtrOutputWithContext(context.Background())
}

func (e DedicatedHostLicenseTypes) ToDedicatedHostLicenseTypesPtrOutputWithContext(ctx context.Context) DedicatedHostLicenseTypesPtrOutput {
	return DedicatedHostLicenseTypes(e).ToDedicatedHostLicenseTypesOutputWithContext(ctx).ToDedicatedHostLicenseTypesPtrOutputWithContext(ctx)
}

func (e DedicatedHostLicenseTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DedicatedHostLicenseTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DedicatedHostLicenseTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DedicatedHostLicenseTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DedicatedHostLicenseTypesOutput struct{ *pulumi.OutputState }

func (DedicatedHostLicenseTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedHostLicenseTypes)(nil)).Elem()
}

func (o DedicatedHostLicenseTypesOutput) ToDedicatedHostLicenseTypesOutput() DedicatedHostLicenseTypesOutput {
	return o
}

func (o DedicatedHostLicenseTypesOutput) ToDedicatedHostLicenseTypesOutputWithContext(ctx context.Context) DedicatedHostLicenseTypesOutput {
	return o
}

func (o DedicatedHostLicenseTypesOutput) ToDedicatedHostLicenseTypesPtrOutput() DedicatedHostLicenseTypesPtrOutput {
	return o.ToDedicatedHostLicenseTypesPtrOutputWithContext(context.Background())
}

func (o DedicatedHostLicenseTypesOutput) ToDedicatedHostLicenseTypesPtrOutputWithContext(ctx context.Context) DedicatedHostLicenseTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DedicatedHostLicenseTypes) *DedicatedHostLicenseTypes {
		return &v
	}).(DedicatedHostLicenseTypesPtrOutput)
}

func (o DedicatedHostLicenseTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DedicatedHostLicenseTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DedicatedHostLicenseTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DedicatedHostLicenseTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DedicatedHostLicenseTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DedicatedHostLicenseTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DedicatedHostLicenseTypesPtrOutput struct{ *pulumi.OutputState }

func (DedicatedHostLicenseTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHostLicenseTypes)(nil)).Elem()
}

func (o DedicatedHostLicenseTypesPtrOutput) ToDedicatedHostLicenseTypesPtrOutput() DedicatedHostLicenseTypesPtrOutput {
	return o
}

func (o DedicatedHostLicenseTypesPtrOutput) ToDedicatedHostLicenseTypesPtrOutputWithContext(ctx context.Context) DedicatedHostLicenseTypesPtrOutput {
	return o
}

func (o DedicatedHostLicenseTypesPtrOutput) Elem() DedicatedHostLicenseTypesOutput {
	return o.ApplyT(func(v *DedicatedHostLicenseTypes) DedicatedHostLicenseTypes {
		if v != nil {
			return *v
		}
		var ret DedicatedHostLicenseTypes
		return ret
	}).(DedicatedHostLicenseTypesOutput)
}

func (o DedicatedHostLicenseTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DedicatedHostLicenseTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DedicatedHostLicenseTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DedicatedHostLicenseTypesInput interface {
	pulumi.Input

	ToDedicatedHostLicenseTypesOutput() DedicatedHostLicenseTypesOutput
	ToDedicatedHostLicenseTypesOutputWithContext(context.Context) DedicatedHostLicenseTypesOutput
}

var dedicatedHostLicenseTypesPtrType = reflect.TypeOf((**DedicatedHostLicenseTypes)(nil)).Elem()

type DedicatedHostLicenseTypesPtrInput interface {
	pulumi.Input

	ToDedicatedHostLicenseTypesPtrOutput() DedicatedHostLicenseTypesPtrOutput
	ToDedicatedHostLicenseTypesPtrOutputWithContext(context.Context) DedicatedHostLicenseTypesPtrOutput
}

type dedicatedHostLicenseTypesPtr string

func DedicatedHostLicenseTypesPtr(v string) DedicatedHostLicenseTypesPtrInput {
	return (*dedicatedHostLicenseTypesPtr)(&v)
}

func (*dedicatedHostLicenseTypesPtr) ElementType() reflect.Type {
	return dedicatedHostLicenseTypesPtrType
}

func (in *dedicatedHostLicenseTypesPtr) ToDedicatedHostLicenseTypesPtrOutput() DedicatedHostLicenseTypesPtrOutput {
	return pulumi.ToOutput(in).(DedicatedHostLicenseTypesPtrOutput)
}

func (in *dedicatedHostLicenseTypesPtr) ToDedicatedHostLicenseTypesPtrOutputWithContext(ctx context.Context) DedicatedHostLicenseTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DedicatedHostLicenseTypesPtrOutput)
}

type DiffDiskOptions string

const (
	DiffDiskOptionsLocal = DiffDiskOptions("Local")
)

func (DiffDiskOptions) ElementType() reflect.Type {
	return reflect.TypeOf((*DiffDiskOptions)(nil)).Elem()
}

func (e DiffDiskOptions) ToDiffDiskOptionsOutput() DiffDiskOptionsOutput {
	return pulumi.ToOutput(e).(DiffDiskOptionsOutput)
}

func (e DiffDiskOptions) ToDiffDiskOptionsOutputWithContext(ctx context.Context) DiffDiskOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DiffDiskOptionsOutput)
}

func (e DiffDiskOptions) ToDiffDiskOptionsPtrOutput() DiffDiskOptionsPtrOutput {
	return e.ToDiffDiskOptionsPtrOutputWithContext(context.Background())
}

func (e DiffDiskOptions) ToDiffDiskOptionsPtrOutputWithContext(ctx context.Context) DiffDiskOptionsPtrOutput {
	return DiffDiskOptions(e).ToDiffDiskOptionsOutputWithContext(ctx).ToDiffDiskOptionsPtrOutputWithContext(ctx)
}

func (e DiffDiskOptions) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiffDiskOptions) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiffDiskOptions) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DiffDiskOptions) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DiffDiskOptionsOutput struct{ *pulumi.OutputState }

func (DiffDiskOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiffDiskOptions)(nil)).Elem()
}

func (o DiffDiskOptionsOutput) ToDiffDiskOptionsOutput() DiffDiskOptionsOutput {
	return o
}

func (o DiffDiskOptionsOutput) ToDiffDiskOptionsOutputWithContext(ctx context.Context) DiffDiskOptionsOutput {
	return o
}

func (o DiffDiskOptionsOutput) ToDiffDiskOptionsPtrOutput() DiffDiskOptionsPtrOutput {
	return o.ToDiffDiskOptionsPtrOutputWithContext(context.Background())
}

func (o DiffDiskOptionsOutput) ToDiffDiskOptionsPtrOutputWithContext(ctx context.Context) DiffDiskOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiffDiskOptions) *DiffDiskOptions {
		return &v
	}).(DiffDiskOptionsPtrOutput)
}

func (o DiffDiskOptionsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DiffDiskOptionsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiffDiskOptions) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DiffDiskOptionsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiffDiskOptionsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiffDiskOptions) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DiffDiskOptionsPtrOutput struct{ *pulumi.OutputState }

func (DiffDiskOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiffDiskOptions)(nil)).Elem()
}

func (o DiffDiskOptionsPtrOutput) ToDiffDiskOptionsPtrOutput() DiffDiskOptionsPtrOutput {
	return o
}

func (o DiffDiskOptionsPtrOutput) ToDiffDiskOptionsPtrOutputWithContext(ctx context.Context) DiffDiskOptionsPtrOutput {
	return o
}

func (o DiffDiskOptionsPtrOutput) Elem() DiffDiskOptionsOutput {
	return o.ApplyT(func(v *DiffDiskOptions) DiffDiskOptions {
		if v != nil {
			return *v
		}
		var ret DiffDiskOptions
		return ret
	}).(DiffDiskOptionsOutput)
}

func (o DiffDiskOptionsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiffDiskOptionsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DiffDiskOptions) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DiffDiskOptionsInput interface {
	pulumi.Input

	ToDiffDiskOptionsOutput() DiffDiskOptionsOutput
	ToDiffDiskOptionsOutputWithContext(context.Context) DiffDiskOptionsOutput
}

var diffDiskOptionsPtrType = reflect.TypeOf((**DiffDiskOptions)(nil)).Elem()

type DiffDiskOptionsPtrInput interface {
	pulumi.Input

	ToDiffDiskOptionsPtrOutput() DiffDiskOptionsPtrOutput
	ToDiffDiskOptionsPtrOutputWithContext(context.Context) DiffDiskOptionsPtrOutput
}

type diffDiskOptionsPtr string

func DiffDiskOptionsPtr(v string) DiffDiskOptionsPtrInput {
	return (*diffDiskOptionsPtr)(&v)
}

func (*diffDiskOptionsPtr) ElementType() reflect.Type {
	return diffDiskOptionsPtrType
}

func (in *diffDiskOptionsPtr) ToDiffDiskOptionsPtrOutput() DiffDiskOptionsPtrOutput {
	return pulumi.ToOutput(in).(DiffDiskOptionsPtrOutput)
}

func (in *diffDiskOptionsPtr) ToDiffDiskOptionsPtrOutputWithContext(ctx context.Context) DiffDiskOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DiffDiskOptionsPtrOutput)
}

type DiffDiskPlacement string

const (
	DiffDiskPlacementCacheDisk    = DiffDiskPlacement("CacheDisk")
	DiffDiskPlacementResourceDisk = DiffDiskPlacement("ResourceDisk")
)

func (DiffDiskPlacement) ElementType() reflect.Type {
	return reflect.TypeOf((*DiffDiskPlacement)(nil)).Elem()
}

func (e DiffDiskPlacement) ToDiffDiskPlacementOutput() DiffDiskPlacementOutput {
	return pulumi.ToOutput(e).(DiffDiskPlacementOutput)
}

func (e DiffDiskPlacement) ToDiffDiskPlacementOutputWithContext(ctx context.Context) DiffDiskPlacementOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DiffDiskPlacementOutput)
}

func (e DiffDiskPlacement) ToDiffDiskPlacementPtrOutput() DiffDiskPlacementPtrOutput {
	return e.ToDiffDiskPlacementPtrOutputWithContext(context.Background())
}

func (e DiffDiskPlacement) ToDiffDiskPlacementPtrOutputWithContext(ctx context.Context) DiffDiskPlacementPtrOutput {
	return DiffDiskPlacement(e).ToDiffDiskPlacementOutputWithContext(ctx).ToDiffDiskPlacementPtrOutputWithContext(ctx)
}

func (e DiffDiskPlacement) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiffDiskPlacement) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiffDiskPlacement) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DiffDiskPlacement) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DiffDiskPlacementOutput struct{ *pulumi.OutputState }

func (DiffDiskPlacementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiffDiskPlacement)(nil)).Elem()
}

func (o DiffDiskPlacementOutput) ToDiffDiskPlacementOutput() DiffDiskPlacementOutput {
	return o
}

func (o DiffDiskPlacementOutput) ToDiffDiskPlacementOutputWithContext(ctx context.Context) DiffDiskPlacementOutput {
	return o
}

func (o DiffDiskPlacementOutput) ToDiffDiskPlacementPtrOutput() DiffDiskPlacementPtrOutput {
	return o.ToDiffDiskPlacementPtrOutputWithContext(context.Background())
}

func (o DiffDiskPlacementOutput) ToDiffDiskPlacementPtrOutputWithContext(ctx context.Context) DiffDiskPlacementPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiffDiskPlacement) *DiffDiskPlacement {
		return &v
	}).(DiffDiskPlacementPtrOutput)
}

func (o DiffDiskPlacementOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DiffDiskPlacementOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiffDiskPlacement) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DiffDiskPlacementOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiffDiskPlacementOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiffDiskPlacement) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DiffDiskPlacementPtrOutput struct{ *pulumi.OutputState }

func (DiffDiskPlacementPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiffDiskPlacement)(nil)).Elem()
}

func (o DiffDiskPlacementPtrOutput) ToDiffDiskPlacementPtrOutput() DiffDiskPlacementPtrOutput {
	return o
}

func (o DiffDiskPlacementPtrOutput) ToDiffDiskPlacementPtrOutputWithContext(ctx context.Context) DiffDiskPlacementPtrOutput {
	return o
}

func (o DiffDiskPlacementPtrOutput) Elem() DiffDiskPlacementOutput {
	return o.ApplyT(func(v *DiffDiskPlacement) DiffDiskPlacement {
		if v != nil {
			return *v
		}
		var ret DiffDiskPlacement
		return ret
	}).(DiffDiskPlacementOutput)
}

func (o DiffDiskPlacementPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiffDiskPlacementPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DiffDiskPlacement) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DiffDiskPlacementInput interface {
	pulumi.Input

	ToDiffDiskPlacementOutput() DiffDiskPlacementOutput
	ToDiffDiskPlacementOutputWithContext(context.Context) DiffDiskPlacementOutput
}

var diffDiskPlacementPtrType = reflect.TypeOf((**DiffDiskPlacement)(nil)).Elem()

type DiffDiskPlacementPtrInput interface {
	pulumi.Input

	ToDiffDiskPlacementPtrOutput() DiffDiskPlacementPtrOutput
	ToDiffDiskPlacementPtrOutputWithContext(context.Context) DiffDiskPlacementPtrOutput
}

type diffDiskPlacementPtr string

func DiffDiskPlacementPtr(v string) DiffDiskPlacementPtrInput {
	return (*diffDiskPlacementPtr)(&v)
}

func (*diffDiskPlacementPtr) ElementType() reflect.Type {
	return diffDiskPlacementPtrType
}

func (in *diffDiskPlacementPtr) ToDiffDiskPlacementPtrOutput() DiffDiskPlacementPtrOutput {
	return pulumi.ToOutput(in).(DiffDiskPlacementPtrOutput)
}

func (in *diffDiskPlacementPtr) ToDiffDiskPlacementPtrOutputWithContext(ctx context.Context) DiffDiskPlacementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DiffDiskPlacementPtrOutput)
}

type DiskCreateOptionTypes string

const (
	DiskCreateOptionTypesFromImage = DiskCreateOptionTypes("FromImage")
	DiskCreateOptionTypesEmpty     = DiskCreateOptionTypes("Empty")
	DiskCreateOptionTypesAttach    = DiskCreateOptionTypes("Attach")
)

func (DiskCreateOptionTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskCreateOptionTypes)(nil)).Elem()
}

func (e DiskCreateOptionTypes) ToDiskCreateOptionTypesOutput() DiskCreateOptionTypesOutput {
	return pulumi.ToOutput(e).(DiskCreateOptionTypesOutput)
}

func (e DiskCreateOptionTypes) ToDiskCreateOptionTypesOutputWithContext(ctx context.Context) DiskCreateOptionTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DiskCreateOptionTypesOutput)
}

func (e DiskCreateOptionTypes) ToDiskCreateOptionTypesPtrOutput() DiskCreateOptionTypesPtrOutput {
	return e.ToDiskCreateOptionTypesPtrOutputWithContext(context.Background())
}

func (e DiskCreateOptionTypes) ToDiskCreateOptionTypesPtrOutputWithContext(ctx context.Context) DiskCreateOptionTypesPtrOutput {
	return DiskCreateOptionTypes(e).ToDiskCreateOptionTypesOutputWithContext(ctx).ToDiskCreateOptionTypesPtrOutputWithContext(ctx)
}

func (e DiskCreateOptionTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskCreateOptionTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskCreateOptionTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DiskCreateOptionTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DiskCreateOptionTypesOutput struct{ *pulumi.OutputState }

func (DiskCreateOptionTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskCreateOptionTypes)(nil)).Elem()
}

func (o DiskCreateOptionTypesOutput) ToDiskCreateOptionTypesOutput() DiskCreateOptionTypesOutput {
	return o
}

func (o DiskCreateOptionTypesOutput) ToDiskCreateOptionTypesOutputWithContext(ctx context.Context) DiskCreateOptionTypesOutput {
	return o
}

func (o DiskCreateOptionTypesOutput) ToDiskCreateOptionTypesPtrOutput() DiskCreateOptionTypesPtrOutput {
	return o.ToDiskCreateOptionTypesPtrOutputWithContext(context.Background())
}

func (o DiskCreateOptionTypesOutput) ToDiskCreateOptionTypesPtrOutputWithContext(ctx context.Context) DiskCreateOptionTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskCreateOptionTypes) *DiskCreateOptionTypes {
		return &v
	}).(DiskCreateOptionTypesPtrOutput)
}

func (o DiskCreateOptionTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DiskCreateOptionTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiskCreateOptionTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DiskCreateOptionTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiskCreateOptionTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiskCreateOptionTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DiskCreateOptionTypesPtrOutput struct{ *pulumi.OutputState }

func (DiskCreateOptionTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskCreateOptionTypes)(nil)).Elem()
}

func (o DiskCreateOptionTypesPtrOutput) ToDiskCreateOptionTypesPtrOutput() DiskCreateOptionTypesPtrOutput {
	return o
}

func (o DiskCreateOptionTypesPtrOutput) ToDiskCreateOptionTypesPtrOutputWithContext(ctx context.Context) DiskCreateOptionTypesPtrOutput {
	return o
}

func (o DiskCreateOptionTypesPtrOutput) Elem() DiskCreateOptionTypesOutput {
	return o.ApplyT(func(v *DiskCreateOptionTypes) DiskCreateOptionTypes {
		if v != nil {
			return *v
		}
		var ret DiskCreateOptionTypes
		return ret
	}).(DiskCreateOptionTypesOutput)
}

func (o DiskCreateOptionTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiskCreateOptionTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DiskCreateOptionTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DiskCreateOptionTypesInput interface {
	pulumi.Input

	ToDiskCreateOptionTypesOutput() DiskCreateOptionTypesOutput
	ToDiskCreateOptionTypesOutputWithContext(context.Context) DiskCreateOptionTypesOutput
}

var diskCreateOptionTypesPtrType = reflect.TypeOf((**DiskCreateOptionTypes)(nil)).Elem()

type DiskCreateOptionTypesPtrInput interface {
	pulumi.Input

	ToDiskCreateOptionTypesPtrOutput() DiskCreateOptionTypesPtrOutput
	ToDiskCreateOptionTypesPtrOutputWithContext(context.Context) DiskCreateOptionTypesPtrOutput
}

type diskCreateOptionTypesPtr string

func DiskCreateOptionTypesPtr(v string) DiskCreateOptionTypesPtrInput {
	return (*diskCreateOptionTypesPtr)(&v)
}

func (*diskCreateOptionTypesPtr) ElementType() reflect.Type {
	return diskCreateOptionTypesPtrType
}

func (in *diskCreateOptionTypesPtr) ToDiskCreateOptionTypesPtrOutput() DiskCreateOptionTypesPtrOutput {
	return pulumi.ToOutput(in).(DiskCreateOptionTypesPtrOutput)
}

func (in *diskCreateOptionTypesPtr) ToDiskCreateOptionTypesPtrOutputWithContext(ctx context.Context) DiskCreateOptionTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DiskCreateOptionTypesPtrOutput)
}

type HyperVGenerationTypes string

const (
	HyperVGenerationTypesV1 = HyperVGenerationTypes("V1")
	HyperVGenerationTypesV2 = HyperVGenerationTypes("V2")
)

func (HyperVGenerationTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVGenerationTypes)(nil)).Elem()
}

func (e HyperVGenerationTypes) ToHyperVGenerationTypesOutput() HyperVGenerationTypesOutput {
	return pulumi.ToOutput(e).(HyperVGenerationTypesOutput)
}

func (e HyperVGenerationTypes) ToHyperVGenerationTypesOutputWithContext(ctx context.Context) HyperVGenerationTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HyperVGenerationTypesOutput)
}

func (e HyperVGenerationTypes) ToHyperVGenerationTypesPtrOutput() HyperVGenerationTypesPtrOutput {
	return e.ToHyperVGenerationTypesPtrOutputWithContext(context.Background())
}

func (e HyperVGenerationTypes) ToHyperVGenerationTypesPtrOutputWithContext(ctx context.Context) HyperVGenerationTypesPtrOutput {
	return HyperVGenerationTypes(e).ToHyperVGenerationTypesOutputWithContext(ctx).ToHyperVGenerationTypesPtrOutputWithContext(ctx)
}

func (e HyperVGenerationTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HyperVGenerationTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HyperVGenerationTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HyperVGenerationTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HyperVGenerationTypesOutput struct{ *pulumi.OutputState }

func (HyperVGenerationTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVGenerationTypes)(nil)).Elem()
}

func (o HyperVGenerationTypesOutput) ToHyperVGenerationTypesOutput() HyperVGenerationTypesOutput {
	return o
}

func (o HyperVGenerationTypesOutput) ToHyperVGenerationTypesOutputWithContext(ctx context.Context) HyperVGenerationTypesOutput {
	return o
}

func (o HyperVGenerationTypesOutput) ToHyperVGenerationTypesPtrOutput() HyperVGenerationTypesPtrOutput {
	return o.ToHyperVGenerationTypesPtrOutputWithContext(context.Background())
}

func (o HyperVGenerationTypesOutput) ToHyperVGenerationTypesPtrOutputWithContext(ctx context.Context) HyperVGenerationTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HyperVGenerationTypes) *HyperVGenerationTypes {
		return &v
	}).(HyperVGenerationTypesPtrOutput)
}

func (o HyperVGenerationTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HyperVGenerationTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HyperVGenerationTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HyperVGenerationTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HyperVGenerationTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HyperVGenerationTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HyperVGenerationTypesPtrOutput struct{ *pulumi.OutputState }

func (HyperVGenerationTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HyperVGenerationTypes)(nil)).Elem()
}

func (o HyperVGenerationTypesPtrOutput) ToHyperVGenerationTypesPtrOutput() HyperVGenerationTypesPtrOutput {
	return o
}

func (o HyperVGenerationTypesPtrOutput) ToHyperVGenerationTypesPtrOutputWithContext(ctx context.Context) HyperVGenerationTypesPtrOutput {
	return o
}

func (o HyperVGenerationTypesPtrOutput) Elem() HyperVGenerationTypesOutput {
	return o.ApplyT(func(v *HyperVGenerationTypes) HyperVGenerationTypes {
		if v != nil {
			return *v
		}
		var ret HyperVGenerationTypes
		return ret
	}).(HyperVGenerationTypesOutput)
}

func (o HyperVGenerationTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HyperVGenerationTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HyperVGenerationTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HyperVGenerationTypesInput interface {
	pulumi.Input

	ToHyperVGenerationTypesOutput() HyperVGenerationTypesOutput
	ToHyperVGenerationTypesOutputWithContext(context.Context) HyperVGenerationTypesOutput
}

var hyperVGenerationTypesPtrType = reflect.TypeOf((**HyperVGenerationTypes)(nil)).Elem()

type HyperVGenerationTypesPtrInput interface {
	pulumi.Input

	ToHyperVGenerationTypesPtrOutput() HyperVGenerationTypesPtrOutput
	ToHyperVGenerationTypesPtrOutputWithContext(context.Context) HyperVGenerationTypesPtrOutput
}

type hyperVGenerationTypesPtr string

func HyperVGenerationTypesPtr(v string) HyperVGenerationTypesPtrInput {
	return (*hyperVGenerationTypesPtr)(&v)
}

func (*hyperVGenerationTypesPtr) ElementType() reflect.Type {
	return hyperVGenerationTypesPtrType
}

func (in *hyperVGenerationTypesPtr) ToHyperVGenerationTypesPtrOutput() HyperVGenerationTypesPtrOutput {
	return pulumi.ToOutput(in).(HyperVGenerationTypesPtrOutput)
}

func (in *hyperVGenerationTypesPtr) ToHyperVGenerationTypesPtrOutputWithContext(ctx context.Context) HyperVGenerationTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HyperVGenerationTypesPtrOutput)
}

type IPVersion string

const (
	IPVersionIPv4 = IPVersion("IPv4")
	IPVersionIPv6 = IPVersion("IPv6")
)

func (IPVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*IPVersion)(nil)).Elem()
}

func (e IPVersion) ToIPVersionOutput() IPVersionOutput {
	return pulumi.ToOutput(e).(IPVersionOutput)
}

func (e IPVersion) ToIPVersionOutputWithContext(ctx context.Context) IPVersionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IPVersionOutput)
}

func (e IPVersion) ToIPVersionPtrOutput() IPVersionPtrOutput {
	return e.ToIPVersionPtrOutputWithContext(context.Background())
}

func (e IPVersion) ToIPVersionPtrOutputWithContext(ctx context.Context) IPVersionPtrOutput {
	return IPVersion(e).ToIPVersionOutputWithContext(ctx).ToIPVersionPtrOutputWithContext(ctx)
}

func (e IPVersion) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IPVersion) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IPVersion) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IPVersion) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IPVersionOutput struct{ *pulumi.OutputState }

func (IPVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPVersion)(nil)).Elem()
}

func (o IPVersionOutput) ToIPVersionOutput() IPVersionOutput {
	return o
}

func (o IPVersionOutput) ToIPVersionOutputWithContext(ctx context.Context) IPVersionOutput {
	return o
}

func (o IPVersionOutput) ToIPVersionPtrOutput() IPVersionPtrOutput {
	return o.ToIPVersionPtrOutputWithContext(context.Background())
}

func (o IPVersionOutput) ToIPVersionPtrOutputWithContext(ctx context.Context) IPVersionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IPVersion) *IPVersion {
		return &v
	}).(IPVersionPtrOutput)
}

func (o IPVersionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IPVersionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IPVersion) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IPVersionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IPVersionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IPVersion) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IPVersionPtrOutput struct{ *pulumi.OutputState }

func (IPVersionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPVersion)(nil)).Elem()
}

func (o IPVersionPtrOutput) ToIPVersionPtrOutput() IPVersionPtrOutput {
	return o
}

func (o IPVersionPtrOutput) ToIPVersionPtrOutputWithContext(ctx context.Context) IPVersionPtrOutput {
	return o
}

func (o IPVersionPtrOutput) Elem() IPVersionOutput {
	return o.ApplyT(func(v *IPVersion) IPVersion {
		if v != nil {
			return *v
		}
		var ret IPVersion
		return ret
	}).(IPVersionOutput)
}

func (o IPVersionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IPVersionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IPVersion) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IPVersionInput interface {
	pulumi.Input

	ToIPVersionOutput() IPVersionOutput
	ToIPVersionOutputWithContext(context.Context) IPVersionOutput
}

var ipversionPtrType = reflect.TypeOf((**IPVersion)(nil)).Elem()

type IPVersionPtrInput interface {
	pulumi.Input

	ToIPVersionPtrOutput() IPVersionPtrOutput
	ToIPVersionPtrOutputWithContext(context.Context) IPVersionPtrOutput
}

type ipversionPtr string

func IPVersionPtr(v string) IPVersionPtrInput {
	return (*ipversionPtr)(&v)
}

func (*ipversionPtr) ElementType() reflect.Type {
	return ipversionPtrType
}

func (in *ipversionPtr) ToIPVersionPtrOutput() IPVersionPtrOutput {
	return pulumi.ToOutput(in).(IPVersionPtrOutput)
}

func (in *ipversionPtr) ToIPVersionPtrOutputWithContext(ctx context.Context) IPVersionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IPVersionPtrOutput)
}

type InGuestPatchMode string

const (
	InGuestPatchModeManual              = InGuestPatchMode("Manual")
	InGuestPatchModeAutomaticByOS       = InGuestPatchMode("AutomaticByOS")
	InGuestPatchModeAutomaticByPlatform = InGuestPatchMode("AutomaticByPlatform")
)

func (InGuestPatchMode) ElementType() reflect.Type {
	return reflect.TypeOf((*InGuestPatchMode)(nil)).Elem()
}

func (e InGuestPatchMode) ToInGuestPatchModeOutput() InGuestPatchModeOutput {
	return pulumi.ToOutput(e).(InGuestPatchModeOutput)
}

func (e InGuestPatchMode) ToInGuestPatchModeOutputWithContext(ctx context.Context) InGuestPatchModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InGuestPatchModeOutput)
}

func (e InGuestPatchMode) ToInGuestPatchModePtrOutput() InGuestPatchModePtrOutput {
	return e.ToInGuestPatchModePtrOutputWithContext(context.Background())
}

func (e InGuestPatchMode) ToInGuestPatchModePtrOutputWithContext(ctx context.Context) InGuestPatchModePtrOutput {
	return InGuestPatchMode(e).ToInGuestPatchModeOutputWithContext(ctx).ToInGuestPatchModePtrOutputWithContext(ctx)
}

func (e InGuestPatchMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InGuestPatchMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InGuestPatchMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InGuestPatchMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InGuestPatchModeOutput struct{ *pulumi.OutputState }

func (InGuestPatchModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InGuestPatchMode)(nil)).Elem()
}

func (o InGuestPatchModeOutput) ToInGuestPatchModeOutput() InGuestPatchModeOutput {
	return o
}

func (o InGuestPatchModeOutput) ToInGuestPatchModeOutputWithContext(ctx context.Context) InGuestPatchModeOutput {
	return o
}

func (o InGuestPatchModeOutput) ToInGuestPatchModePtrOutput() InGuestPatchModePtrOutput {
	return o.ToInGuestPatchModePtrOutputWithContext(context.Background())
}

func (o InGuestPatchModeOutput) ToInGuestPatchModePtrOutputWithContext(ctx context.Context) InGuestPatchModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InGuestPatchMode) *InGuestPatchMode {
		return &v
	}).(InGuestPatchModePtrOutput)
}

func (o InGuestPatchModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InGuestPatchModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InGuestPatchMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InGuestPatchModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InGuestPatchModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InGuestPatchMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InGuestPatchModePtrOutput struct{ *pulumi.OutputState }

func (InGuestPatchModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InGuestPatchMode)(nil)).Elem()
}

func (o InGuestPatchModePtrOutput) ToInGuestPatchModePtrOutput() InGuestPatchModePtrOutput {
	return o
}

func (o InGuestPatchModePtrOutput) ToInGuestPatchModePtrOutputWithContext(ctx context.Context) InGuestPatchModePtrOutput {
	return o
}

func (o InGuestPatchModePtrOutput) Elem() InGuestPatchModeOutput {
	return o.ApplyT(func(v *InGuestPatchMode) InGuestPatchMode {
		if v != nil {
			return *v
		}
		var ret InGuestPatchMode
		return ret
	}).(InGuestPatchModeOutput)
}

func (o InGuestPatchModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InGuestPatchModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InGuestPatchMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type InGuestPatchModeInput interface {
	pulumi.Input

	ToInGuestPatchModeOutput() InGuestPatchModeOutput
	ToInGuestPatchModeOutputWithContext(context.Context) InGuestPatchModeOutput
}

var inGuestPatchModePtrType = reflect.TypeOf((**InGuestPatchMode)(nil)).Elem()

type InGuestPatchModePtrInput interface {
	pulumi.Input

	ToInGuestPatchModePtrOutput() InGuestPatchModePtrOutput
	ToInGuestPatchModePtrOutputWithContext(context.Context) InGuestPatchModePtrOutput
}

type inGuestPatchModePtr string

func InGuestPatchModePtr(v string) InGuestPatchModePtrInput {
	return (*inGuestPatchModePtr)(&v)
}

func (*inGuestPatchModePtr) ElementType() reflect.Type {
	return inGuestPatchModePtrType
}

func (in *inGuestPatchModePtr) ToInGuestPatchModePtrOutput() InGuestPatchModePtrOutput {
	return pulumi.ToOutput(in).(InGuestPatchModePtrOutput)
}

func (in *inGuestPatchModePtr) ToInGuestPatchModePtrOutputWithContext(ctx context.Context) InGuestPatchModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InGuestPatchModePtrOutput)
}

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
	// Generalized image. Needs to be provisioned during deployment time.
	OperatingSystemStateTypesGeneralized = OperatingSystemStateTypes("Generalized")
	// Specialized image. Contains already provisioned OS Disk.
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

func (ProximityPlacementGroupType) ElementType() reflect.Type {
	return reflect.TypeOf((*ProximityPlacementGroupType)(nil)).Elem()
}

func (e ProximityPlacementGroupType) ToProximityPlacementGroupTypeOutput() ProximityPlacementGroupTypeOutput {
	return pulumi.ToOutput(e).(ProximityPlacementGroupTypeOutput)
}

func (e ProximityPlacementGroupType) ToProximityPlacementGroupTypeOutputWithContext(ctx context.Context) ProximityPlacementGroupTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProximityPlacementGroupTypeOutput)
}

func (e ProximityPlacementGroupType) ToProximityPlacementGroupTypePtrOutput() ProximityPlacementGroupTypePtrOutput {
	return e.ToProximityPlacementGroupTypePtrOutputWithContext(context.Background())
}

func (e ProximityPlacementGroupType) ToProximityPlacementGroupTypePtrOutputWithContext(ctx context.Context) ProximityPlacementGroupTypePtrOutput {
	return ProximityPlacementGroupType(e).ToProximityPlacementGroupTypeOutputWithContext(ctx).ToProximityPlacementGroupTypePtrOutputWithContext(ctx)
}

func (e ProximityPlacementGroupType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProximityPlacementGroupType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProximityPlacementGroupType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProximityPlacementGroupType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProximityPlacementGroupTypeOutput struct{ *pulumi.OutputState }

func (ProximityPlacementGroupTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProximityPlacementGroupType)(nil)).Elem()
}

func (o ProximityPlacementGroupTypeOutput) ToProximityPlacementGroupTypeOutput() ProximityPlacementGroupTypeOutput {
	return o
}

func (o ProximityPlacementGroupTypeOutput) ToProximityPlacementGroupTypeOutputWithContext(ctx context.Context) ProximityPlacementGroupTypeOutput {
	return o
}

func (o ProximityPlacementGroupTypeOutput) ToProximityPlacementGroupTypePtrOutput() ProximityPlacementGroupTypePtrOutput {
	return o.ToProximityPlacementGroupTypePtrOutputWithContext(context.Background())
}

func (o ProximityPlacementGroupTypeOutput) ToProximityPlacementGroupTypePtrOutputWithContext(ctx context.Context) ProximityPlacementGroupTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProximityPlacementGroupType) *ProximityPlacementGroupType {
		return &v
	}).(ProximityPlacementGroupTypePtrOutput)
}

func (o ProximityPlacementGroupTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProximityPlacementGroupTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProximityPlacementGroupType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProximityPlacementGroupTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProximityPlacementGroupTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProximityPlacementGroupType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProximityPlacementGroupTypePtrOutput struct{ *pulumi.OutputState }

func (ProximityPlacementGroupTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProximityPlacementGroupType)(nil)).Elem()
}

func (o ProximityPlacementGroupTypePtrOutput) ToProximityPlacementGroupTypePtrOutput() ProximityPlacementGroupTypePtrOutput {
	return o
}

func (o ProximityPlacementGroupTypePtrOutput) ToProximityPlacementGroupTypePtrOutputWithContext(ctx context.Context) ProximityPlacementGroupTypePtrOutput {
	return o
}

func (o ProximityPlacementGroupTypePtrOutput) Elem() ProximityPlacementGroupTypeOutput {
	return o.ApplyT(func(v *ProximityPlacementGroupType) ProximityPlacementGroupType {
		if v != nil {
			return *v
		}
		var ret ProximityPlacementGroupType
		return ret
	}).(ProximityPlacementGroupTypeOutput)
}

func (o ProximityPlacementGroupTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProximityPlacementGroupTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProximityPlacementGroupType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProximityPlacementGroupTypeInput interface {
	pulumi.Input

	ToProximityPlacementGroupTypeOutput() ProximityPlacementGroupTypeOutput
	ToProximityPlacementGroupTypeOutputWithContext(context.Context) ProximityPlacementGroupTypeOutput
}

var proximityPlacementGroupTypePtrType = reflect.TypeOf((**ProximityPlacementGroupType)(nil)).Elem()

type ProximityPlacementGroupTypePtrInput interface {
	pulumi.Input

	ToProximityPlacementGroupTypePtrOutput() ProximityPlacementGroupTypePtrOutput
	ToProximityPlacementGroupTypePtrOutputWithContext(context.Context) ProximityPlacementGroupTypePtrOutput
}

type proximityPlacementGroupTypePtr string

func ProximityPlacementGroupTypePtr(v string) ProximityPlacementGroupTypePtrInput {
	return (*proximityPlacementGroupTypePtr)(&v)
}

func (*proximityPlacementGroupTypePtr) ElementType() reflect.Type {
	return proximityPlacementGroupTypePtrType
}

func (in *proximityPlacementGroupTypePtr) ToProximityPlacementGroupTypePtrOutput() ProximityPlacementGroupTypePtrOutput {
	return pulumi.ToOutput(in).(ProximityPlacementGroupTypePtrOutput)
}

func (in *proximityPlacementGroupTypePtr) ToProximityPlacementGroupTypePtrOutputWithContext(ctx context.Context) ProximityPlacementGroupTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProximityPlacementGroupTypePtrOutput)
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
	StorageAccountTypes_UltraSSD_LRS    = StorageAccountTypes("UltraSSD_LRS")
)

func (StorageAccountTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountTypes)(nil)).Elem()
}

func (e StorageAccountTypes) ToStorageAccountTypesOutput() StorageAccountTypesOutput {
	return pulumi.ToOutput(e).(StorageAccountTypesOutput)
}

func (e StorageAccountTypes) ToStorageAccountTypesOutputWithContext(ctx context.Context) StorageAccountTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StorageAccountTypesOutput)
}

func (e StorageAccountTypes) ToStorageAccountTypesPtrOutput() StorageAccountTypesPtrOutput {
	return e.ToStorageAccountTypesPtrOutputWithContext(context.Background())
}

func (e StorageAccountTypes) ToStorageAccountTypesPtrOutputWithContext(ctx context.Context) StorageAccountTypesPtrOutput {
	return StorageAccountTypes(e).ToStorageAccountTypesOutputWithContext(ctx).ToStorageAccountTypesPtrOutputWithContext(ctx)
}

func (e StorageAccountTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageAccountTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageAccountTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StorageAccountTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StorageAccountTypesOutput struct{ *pulumi.OutputState }

func (StorageAccountTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountTypes)(nil)).Elem()
}

func (o StorageAccountTypesOutput) ToStorageAccountTypesOutput() StorageAccountTypesOutput {
	return o
}

func (o StorageAccountTypesOutput) ToStorageAccountTypesOutputWithContext(ctx context.Context) StorageAccountTypesOutput {
	return o
}

func (o StorageAccountTypesOutput) ToStorageAccountTypesPtrOutput() StorageAccountTypesPtrOutput {
	return o.ToStorageAccountTypesPtrOutputWithContext(context.Background())
}

func (o StorageAccountTypesOutput) ToStorageAccountTypesPtrOutputWithContext(ctx context.Context) StorageAccountTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccountTypes) *StorageAccountTypes {
		return &v
	}).(StorageAccountTypesPtrOutput)
}

func (o StorageAccountTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StorageAccountTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageAccountTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StorageAccountTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageAccountTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageAccountTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StorageAccountTypesPtrOutput struct{ *pulumi.OutputState }

func (StorageAccountTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountTypes)(nil)).Elem()
}

func (o StorageAccountTypesPtrOutput) ToStorageAccountTypesPtrOutput() StorageAccountTypesPtrOutput {
	return o
}

func (o StorageAccountTypesPtrOutput) ToStorageAccountTypesPtrOutputWithContext(ctx context.Context) StorageAccountTypesPtrOutput {
	return o
}

func (o StorageAccountTypesPtrOutput) Elem() StorageAccountTypesOutput {
	return o.ApplyT(func(v *StorageAccountTypes) StorageAccountTypes {
		if v != nil {
			return *v
		}
		var ret StorageAccountTypes
		return ret
	}).(StorageAccountTypesOutput)
}

func (o StorageAccountTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageAccountTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StorageAccountTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StorageAccountTypesInput interface {
	pulumi.Input

	ToStorageAccountTypesOutput() StorageAccountTypesOutput
	ToStorageAccountTypesOutputWithContext(context.Context) StorageAccountTypesOutput
}

var storageAccountTypesPtrType = reflect.TypeOf((**StorageAccountTypes)(nil)).Elem()

type StorageAccountTypesPtrInput interface {
	pulumi.Input

	ToStorageAccountTypesPtrOutput() StorageAccountTypesPtrOutput
	ToStorageAccountTypesPtrOutputWithContext(context.Context) StorageAccountTypesPtrOutput
}

type storageAccountTypesPtr string

func StorageAccountTypesPtr(v string) StorageAccountTypesPtrInput {
	return (*storageAccountTypesPtr)(&v)
}

func (*storageAccountTypesPtr) ElementType() reflect.Type {
	return storageAccountTypesPtrType
}

func (in *storageAccountTypesPtr) ToStorageAccountTypesPtrOutput() StorageAccountTypesPtrOutput {
	return pulumi.ToOutput(in).(StorageAccountTypesPtrOutput)
}

func (in *storageAccountTypesPtr) ToStorageAccountTypesPtrOutputWithContext(ctx context.Context) StorageAccountTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StorageAccountTypesPtrOutput)
}

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

func (VirtualMachineEvictionPolicyTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineEvictionPolicyTypes)(nil)).Elem()
}

func (e VirtualMachineEvictionPolicyTypes) ToVirtualMachineEvictionPolicyTypesOutput() VirtualMachineEvictionPolicyTypesOutput {
	return pulumi.ToOutput(e).(VirtualMachineEvictionPolicyTypesOutput)
}

func (e VirtualMachineEvictionPolicyTypes) ToVirtualMachineEvictionPolicyTypesOutputWithContext(ctx context.Context) VirtualMachineEvictionPolicyTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VirtualMachineEvictionPolicyTypesOutput)
}

func (e VirtualMachineEvictionPolicyTypes) ToVirtualMachineEvictionPolicyTypesPtrOutput() VirtualMachineEvictionPolicyTypesPtrOutput {
	return e.ToVirtualMachineEvictionPolicyTypesPtrOutputWithContext(context.Background())
}

func (e VirtualMachineEvictionPolicyTypes) ToVirtualMachineEvictionPolicyTypesPtrOutputWithContext(ctx context.Context) VirtualMachineEvictionPolicyTypesPtrOutput {
	return VirtualMachineEvictionPolicyTypes(e).ToVirtualMachineEvictionPolicyTypesOutputWithContext(ctx).ToVirtualMachineEvictionPolicyTypesPtrOutputWithContext(ctx)
}

func (e VirtualMachineEvictionPolicyTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualMachineEvictionPolicyTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualMachineEvictionPolicyTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VirtualMachineEvictionPolicyTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VirtualMachineEvictionPolicyTypesOutput struct{ *pulumi.OutputState }

func (VirtualMachineEvictionPolicyTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineEvictionPolicyTypes)(nil)).Elem()
}

func (o VirtualMachineEvictionPolicyTypesOutput) ToVirtualMachineEvictionPolicyTypesOutput() VirtualMachineEvictionPolicyTypesOutput {
	return o
}

func (o VirtualMachineEvictionPolicyTypesOutput) ToVirtualMachineEvictionPolicyTypesOutputWithContext(ctx context.Context) VirtualMachineEvictionPolicyTypesOutput {
	return o
}

func (o VirtualMachineEvictionPolicyTypesOutput) ToVirtualMachineEvictionPolicyTypesPtrOutput() VirtualMachineEvictionPolicyTypesPtrOutput {
	return o.ToVirtualMachineEvictionPolicyTypesPtrOutputWithContext(context.Background())
}

func (o VirtualMachineEvictionPolicyTypesOutput) ToVirtualMachineEvictionPolicyTypesPtrOutputWithContext(ctx context.Context) VirtualMachineEvictionPolicyTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineEvictionPolicyTypes) *VirtualMachineEvictionPolicyTypes {
		return &v
	}).(VirtualMachineEvictionPolicyTypesPtrOutput)
}

func (o VirtualMachineEvictionPolicyTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VirtualMachineEvictionPolicyTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualMachineEvictionPolicyTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VirtualMachineEvictionPolicyTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualMachineEvictionPolicyTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualMachineEvictionPolicyTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineEvictionPolicyTypesPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineEvictionPolicyTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineEvictionPolicyTypes)(nil)).Elem()
}

func (o VirtualMachineEvictionPolicyTypesPtrOutput) ToVirtualMachineEvictionPolicyTypesPtrOutput() VirtualMachineEvictionPolicyTypesPtrOutput {
	return o
}

func (o VirtualMachineEvictionPolicyTypesPtrOutput) ToVirtualMachineEvictionPolicyTypesPtrOutputWithContext(ctx context.Context) VirtualMachineEvictionPolicyTypesPtrOutput {
	return o
}

func (o VirtualMachineEvictionPolicyTypesPtrOutput) Elem() VirtualMachineEvictionPolicyTypesOutput {
	return o.ApplyT(func(v *VirtualMachineEvictionPolicyTypes) VirtualMachineEvictionPolicyTypes {
		if v != nil {
			return *v
		}
		var ret VirtualMachineEvictionPolicyTypes
		return ret
	}).(VirtualMachineEvictionPolicyTypesOutput)
}

func (o VirtualMachineEvictionPolicyTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualMachineEvictionPolicyTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VirtualMachineEvictionPolicyTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VirtualMachineEvictionPolicyTypesInput interface {
	pulumi.Input

	ToVirtualMachineEvictionPolicyTypesOutput() VirtualMachineEvictionPolicyTypesOutput
	ToVirtualMachineEvictionPolicyTypesOutputWithContext(context.Context) VirtualMachineEvictionPolicyTypesOutput
}

var virtualMachineEvictionPolicyTypesPtrType = reflect.TypeOf((**VirtualMachineEvictionPolicyTypes)(nil)).Elem()

type VirtualMachineEvictionPolicyTypesPtrInput interface {
	pulumi.Input

	ToVirtualMachineEvictionPolicyTypesPtrOutput() VirtualMachineEvictionPolicyTypesPtrOutput
	ToVirtualMachineEvictionPolicyTypesPtrOutputWithContext(context.Context) VirtualMachineEvictionPolicyTypesPtrOutput
}

type virtualMachineEvictionPolicyTypesPtr string

func VirtualMachineEvictionPolicyTypesPtr(v string) VirtualMachineEvictionPolicyTypesPtrInput {
	return (*virtualMachineEvictionPolicyTypesPtr)(&v)
}

func (*virtualMachineEvictionPolicyTypesPtr) ElementType() reflect.Type {
	return virtualMachineEvictionPolicyTypesPtrType
}

func (in *virtualMachineEvictionPolicyTypesPtr) ToVirtualMachineEvictionPolicyTypesPtrOutput() VirtualMachineEvictionPolicyTypesPtrOutput {
	return pulumi.ToOutput(in).(VirtualMachineEvictionPolicyTypesPtrOutput)
}

func (in *virtualMachineEvictionPolicyTypesPtr) ToVirtualMachineEvictionPolicyTypesPtrOutputWithContext(ctx context.Context) VirtualMachineEvictionPolicyTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VirtualMachineEvictionPolicyTypesPtrOutput)
}

type VirtualMachinePriorityTypes string

const (
	VirtualMachinePriorityTypesRegular = VirtualMachinePriorityTypes("Regular")
	VirtualMachinePriorityTypesLow     = VirtualMachinePriorityTypes("Low")
	VirtualMachinePriorityTypesSpot    = VirtualMachinePriorityTypes("Spot")
)

func (VirtualMachinePriorityTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachinePriorityTypes)(nil)).Elem()
}

func (e VirtualMachinePriorityTypes) ToVirtualMachinePriorityTypesOutput() VirtualMachinePriorityTypesOutput {
	return pulumi.ToOutput(e).(VirtualMachinePriorityTypesOutput)
}

func (e VirtualMachinePriorityTypes) ToVirtualMachinePriorityTypesOutputWithContext(ctx context.Context) VirtualMachinePriorityTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VirtualMachinePriorityTypesOutput)
}

func (e VirtualMachinePriorityTypes) ToVirtualMachinePriorityTypesPtrOutput() VirtualMachinePriorityTypesPtrOutput {
	return e.ToVirtualMachinePriorityTypesPtrOutputWithContext(context.Background())
}

func (e VirtualMachinePriorityTypes) ToVirtualMachinePriorityTypesPtrOutputWithContext(ctx context.Context) VirtualMachinePriorityTypesPtrOutput {
	return VirtualMachinePriorityTypes(e).ToVirtualMachinePriorityTypesOutputWithContext(ctx).ToVirtualMachinePriorityTypesPtrOutputWithContext(ctx)
}

func (e VirtualMachinePriorityTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualMachinePriorityTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualMachinePriorityTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VirtualMachinePriorityTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VirtualMachinePriorityTypesOutput struct{ *pulumi.OutputState }

func (VirtualMachinePriorityTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachinePriorityTypes)(nil)).Elem()
}

func (o VirtualMachinePriorityTypesOutput) ToVirtualMachinePriorityTypesOutput() VirtualMachinePriorityTypesOutput {
	return o
}

func (o VirtualMachinePriorityTypesOutput) ToVirtualMachinePriorityTypesOutputWithContext(ctx context.Context) VirtualMachinePriorityTypesOutput {
	return o
}

func (o VirtualMachinePriorityTypesOutput) ToVirtualMachinePriorityTypesPtrOutput() VirtualMachinePriorityTypesPtrOutput {
	return o.ToVirtualMachinePriorityTypesPtrOutputWithContext(context.Background())
}

func (o VirtualMachinePriorityTypesOutput) ToVirtualMachinePriorityTypesPtrOutputWithContext(ctx context.Context) VirtualMachinePriorityTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachinePriorityTypes) *VirtualMachinePriorityTypes {
		return &v
	}).(VirtualMachinePriorityTypesPtrOutput)
}

func (o VirtualMachinePriorityTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VirtualMachinePriorityTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualMachinePriorityTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VirtualMachinePriorityTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualMachinePriorityTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualMachinePriorityTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VirtualMachinePriorityTypesPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachinePriorityTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachinePriorityTypes)(nil)).Elem()
}

func (o VirtualMachinePriorityTypesPtrOutput) ToVirtualMachinePriorityTypesPtrOutput() VirtualMachinePriorityTypesPtrOutput {
	return o
}

func (o VirtualMachinePriorityTypesPtrOutput) ToVirtualMachinePriorityTypesPtrOutputWithContext(ctx context.Context) VirtualMachinePriorityTypesPtrOutput {
	return o
}

func (o VirtualMachinePriorityTypesPtrOutput) Elem() VirtualMachinePriorityTypesOutput {
	return o.ApplyT(func(v *VirtualMachinePriorityTypes) VirtualMachinePriorityTypes {
		if v != nil {
			return *v
		}
		var ret VirtualMachinePriorityTypes
		return ret
	}).(VirtualMachinePriorityTypesOutput)
}

func (o VirtualMachinePriorityTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualMachinePriorityTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VirtualMachinePriorityTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VirtualMachinePriorityTypesInput interface {
	pulumi.Input

	ToVirtualMachinePriorityTypesOutput() VirtualMachinePriorityTypesOutput
	ToVirtualMachinePriorityTypesOutputWithContext(context.Context) VirtualMachinePriorityTypesOutput
}

var virtualMachinePriorityTypesPtrType = reflect.TypeOf((**VirtualMachinePriorityTypes)(nil)).Elem()

type VirtualMachinePriorityTypesPtrInput interface {
	pulumi.Input

	ToVirtualMachinePriorityTypesPtrOutput() VirtualMachinePriorityTypesPtrOutput
	ToVirtualMachinePriorityTypesPtrOutputWithContext(context.Context) VirtualMachinePriorityTypesPtrOutput
}

type virtualMachinePriorityTypesPtr string

func VirtualMachinePriorityTypesPtr(v string) VirtualMachinePriorityTypesPtrInput {
	return (*virtualMachinePriorityTypesPtr)(&v)
}

func (*virtualMachinePriorityTypesPtr) ElementType() reflect.Type {
	return virtualMachinePriorityTypesPtrType
}

func (in *virtualMachinePriorityTypesPtr) ToVirtualMachinePriorityTypesPtrOutput() VirtualMachinePriorityTypesPtrOutput {
	return pulumi.ToOutput(in).(VirtualMachinePriorityTypesPtrOutput)
}

func (in *virtualMachinePriorityTypesPtr) ToVirtualMachinePriorityTypesPtrOutputWithContext(ctx context.Context) VirtualMachinePriorityTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VirtualMachinePriorityTypesPtrOutput)
}

type VirtualMachineScaleSetScaleInRules string

const (
	VirtualMachineScaleSetScaleInRulesDefault  = VirtualMachineScaleSetScaleInRules("Default")
	VirtualMachineScaleSetScaleInRulesOldestVM = VirtualMachineScaleSetScaleInRules("OldestVM")
	VirtualMachineScaleSetScaleInRulesNewestVM = VirtualMachineScaleSetScaleInRules("NewestVM")
)

func (VirtualMachineScaleSetScaleInRules) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetScaleInRules)(nil)).Elem()
}

func (e VirtualMachineScaleSetScaleInRules) ToVirtualMachineScaleSetScaleInRulesOutput() VirtualMachineScaleSetScaleInRulesOutput {
	return pulumi.ToOutput(e).(VirtualMachineScaleSetScaleInRulesOutput)
}

func (e VirtualMachineScaleSetScaleInRules) ToVirtualMachineScaleSetScaleInRulesOutputWithContext(ctx context.Context) VirtualMachineScaleSetScaleInRulesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VirtualMachineScaleSetScaleInRulesOutput)
}

func (e VirtualMachineScaleSetScaleInRules) ToVirtualMachineScaleSetScaleInRulesPtrOutput() VirtualMachineScaleSetScaleInRulesPtrOutput {
	return e.ToVirtualMachineScaleSetScaleInRulesPtrOutputWithContext(context.Background())
}

func (e VirtualMachineScaleSetScaleInRules) ToVirtualMachineScaleSetScaleInRulesPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetScaleInRulesPtrOutput {
	return VirtualMachineScaleSetScaleInRules(e).ToVirtualMachineScaleSetScaleInRulesOutputWithContext(ctx).ToVirtualMachineScaleSetScaleInRulesPtrOutputWithContext(ctx)
}

func (e VirtualMachineScaleSetScaleInRules) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualMachineScaleSetScaleInRules) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualMachineScaleSetScaleInRules) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VirtualMachineScaleSetScaleInRules) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VirtualMachineScaleSetScaleInRulesOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetScaleInRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetScaleInRules)(nil)).Elem()
}

func (o VirtualMachineScaleSetScaleInRulesOutput) ToVirtualMachineScaleSetScaleInRulesOutput() VirtualMachineScaleSetScaleInRulesOutput {
	return o
}

func (o VirtualMachineScaleSetScaleInRulesOutput) ToVirtualMachineScaleSetScaleInRulesOutputWithContext(ctx context.Context) VirtualMachineScaleSetScaleInRulesOutput {
	return o
}

func (o VirtualMachineScaleSetScaleInRulesOutput) ToVirtualMachineScaleSetScaleInRulesPtrOutput() VirtualMachineScaleSetScaleInRulesPtrOutput {
	return o.ToVirtualMachineScaleSetScaleInRulesPtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetScaleInRulesOutput) ToVirtualMachineScaleSetScaleInRulesPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetScaleInRulesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetScaleInRules) *VirtualMachineScaleSetScaleInRules {
		return &v
	}).(VirtualMachineScaleSetScaleInRulesPtrOutput)
}

func (o VirtualMachineScaleSetScaleInRulesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetScaleInRulesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualMachineScaleSetScaleInRules) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetScaleInRulesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetScaleInRulesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualMachineScaleSetScaleInRules) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetScaleInRulesPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetScaleInRulesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetScaleInRules)(nil)).Elem()
}

func (o VirtualMachineScaleSetScaleInRulesPtrOutput) ToVirtualMachineScaleSetScaleInRulesPtrOutput() VirtualMachineScaleSetScaleInRulesPtrOutput {
	return o
}

func (o VirtualMachineScaleSetScaleInRulesPtrOutput) ToVirtualMachineScaleSetScaleInRulesPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetScaleInRulesPtrOutput {
	return o
}

func (o VirtualMachineScaleSetScaleInRulesPtrOutput) Elem() VirtualMachineScaleSetScaleInRulesOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetScaleInRules) VirtualMachineScaleSetScaleInRules {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetScaleInRules
		return ret
	}).(VirtualMachineScaleSetScaleInRulesOutput)
}

func (o VirtualMachineScaleSetScaleInRulesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetScaleInRulesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VirtualMachineScaleSetScaleInRules) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VirtualMachineScaleSetScaleInRulesInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetScaleInRulesOutput() VirtualMachineScaleSetScaleInRulesOutput
	ToVirtualMachineScaleSetScaleInRulesOutputWithContext(context.Context) VirtualMachineScaleSetScaleInRulesOutput
}

var virtualMachineScaleSetScaleInRulesPtrType = reflect.TypeOf((**VirtualMachineScaleSetScaleInRules)(nil)).Elem()

type VirtualMachineScaleSetScaleInRulesPtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetScaleInRulesPtrOutput() VirtualMachineScaleSetScaleInRulesPtrOutput
	ToVirtualMachineScaleSetScaleInRulesPtrOutputWithContext(context.Context) VirtualMachineScaleSetScaleInRulesPtrOutput
}

type virtualMachineScaleSetScaleInRulesPtr string

func VirtualMachineScaleSetScaleInRulesPtr(v string) VirtualMachineScaleSetScaleInRulesPtrInput {
	return (*virtualMachineScaleSetScaleInRulesPtr)(&v)
}

func (*virtualMachineScaleSetScaleInRulesPtr) ElementType() reflect.Type {
	return virtualMachineScaleSetScaleInRulesPtrType
}

func (in *virtualMachineScaleSetScaleInRulesPtr) ToVirtualMachineScaleSetScaleInRulesPtrOutput() VirtualMachineScaleSetScaleInRulesPtrOutput {
	return pulumi.ToOutput(in).(VirtualMachineScaleSetScaleInRulesPtrOutput)
}

func (in *virtualMachineScaleSetScaleInRulesPtr) ToVirtualMachineScaleSetScaleInRulesPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetScaleInRulesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VirtualMachineScaleSetScaleInRulesPtrOutput)
}

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

func (VirtualMachineSizeTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineSizeTypes)(nil)).Elem()
}

func (e VirtualMachineSizeTypes) ToVirtualMachineSizeTypesOutput() VirtualMachineSizeTypesOutput {
	return pulumi.ToOutput(e).(VirtualMachineSizeTypesOutput)
}

func (e VirtualMachineSizeTypes) ToVirtualMachineSizeTypesOutputWithContext(ctx context.Context) VirtualMachineSizeTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VirtualMachineSizeTypesOutput)
}

func (e VirtualMachineSizeTypes) ToVirtualMachineSizeTypesPtrOutput() VirtualMachineSizeTypesPtrOutput {
	return e.ToVirtualMachineSizeTypesPtrOutputWithContext(context.Background())
}

func (e VirtualMachineSizeTypes) ToVirtualMachineSizeTypesPtrOutputWithContext(ctx context.Context) VirtualMachineSizeTypesPtrOutput {
	return VirtualMachineSizeTypes(e).ToVirtualMachineSizeTypesOutputWithContext(ctx).ToVirtualMachineSizeTypesPtrOutputWithContext(ctx)
}

func (e VirtualMachineSizeTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualMachineSizeTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VirtualMachineSizeTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VirtualMachineSizeTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VirtualMachineSizeTypesOutput struct{ *pulumi.OutputState }

func (VirtualMachineSizeTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineSizeTypes)(nil)).Elem()
}

func (o VirtualMachineSizeTypesOutput) ToVirtualMachineSizeTypesOutput() VirtualMachineSizeTypesOutput {
	return o
}

func (o VirtualMachineSizeTypesOutput) ToVirtualMachineSizeTypesOutputWithContext(ctx context.Context) VirtualMachineSizeTypesOutput {
	return o
}

func (o VirtualMachineSizeTypesOutput) ToVirtualMachineSizeTypesPtrOutput() VirtualMachineSizeTypesPtrOutput {
	return o.ToVirtualMachineSizeTypesPtrOutputWithContext(context.Background())
}

func (o VirtualMachineSizeTypesOutput) ToVirtualMachineSizeTypesPtrOutputWithContext(ctx context.Context) VirtualMachineSizeTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineSizeTypes) *VirtualMachineSizeTypes {
		return &v
	}).(VirtualMachineSizeTypesPtrOutput)
}

func (o VirtualMachineSizeTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VirtualMachineSizeTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualMachineSizeTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VirtualMachineSizeTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualMachineSizeTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VirtualMachineSizeTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineSizeTypesPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineSizeTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineSizeTypes)(nil)).Elem()
}

func (o VirtualMachineSizeTypesPtrOutput) ToVirtualMachineSizeTypesPtrOutput() VirtualMachineSizeTypesPtrOutput {
	return o
}

func (o VirtualMachineSizeTypesPtrOutput) ToVirtualMachineSizeTypesPtrOutputWithContext(ctx context.Context) VirtualMachineSizeTypesPtrOutput {
	return o
}

func (o VirtualMachineSizeTypesPtrOutput) Elem() VirtualMachineSizeTypesOutput {
	return o.ApplyT(func(v *VirtualMachineSizeTypes) VirtualMachineSizeTypes {
		if v != nil {
			return *v
		}
		var ret VirtualMachineSizeTypes
		return ret
	}).(VirtualMachineSizeTypesOutput)
}

func (o VirtualMachineSizeTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VirtualMachineSizeTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VirtualMachineSizeTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VirtualMachineSizeTypesInput interface {
	pulumi.Input

	ToVirtualMachineSizeTypesOutput() VirtualMachineSizeTypesOutput
	ToVirtualMachineSizeTypesOutputWithContext(context.Context) VirtualMachineSizeTypesOutput
}

var virtualMachineSizeTypesPtrType = reflect.TypeOf((**VirtualMachineSizeTypes)(nil)).Elem()

type VirtualMachineSizeTypesPtrInput interface {
	pulumi.Input

	ToVirtualMachineSizeTypesPtrOutput() VirtualMachineSizeTypesPtrOutput
	ToVirtualMachineSizeTypesPtrOutputWithContext(context.Context) VirtualMachineSizeTypesPtrOutput
}

type virtualMachineSizeTypesPtr string

func VirtualMachineSizeTypesPtr(v string) VirtualMachineSizeTypesPtrInput {
	return (*virtualMachineSizeTypesPtr)(&v)
}

func (*virtualMachineSizeTypesPtr) ElementType() reflect.Type {
	return virtualMachineSizeTypesPtrType
}

func (in *virtualMachineSizeTypesPtr) ToVirtualMachineSizeTypesPtrOutput() VirtualMachineSizeTypesPtrOutput {
	return pulumi.ToOutput(in).(VirtualMachineSizeTypesPtrOutput)
}

func (in *virtualMachineSizeTypesPtr) ToVirtualMachineSizeTypesPtrOutputWithContext(ctx context.Context) VirtualMachineSizeTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VirtualMachineSizeTypesPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CachingTypesOutput{})
	pulumi.RegisterOutputType(CachingTypesPtrOutput{})
	pulumi.RegisterOutputType(ComponentNamesOutput{})
	pulumi.RegisterOutputType(ComponentNamesPtrOutput{})
	pulumi.RegisterOutputType(DedicatedHostLicenseTypesOutput{})
	pulumi.RegisterOutputType(DedicatedHostLicenseTypesPtrOutput{})
	pulumi.RegisterOutputType(DiffDiskOptionsOutput{})
	pulumi.RegisterOutputType(DiffDiskOptionsPtrOutput{})
	pulumi.RegisterOutputType(DiffDiskPlacementOutput{})
	pulumi.RegisterOutputType(DiffDiskPlacementPtrOutput{})
	pulumi.RegisterOutputType(DiskCreateOptionTypesOutput{})
	pulumi.RegisterOutputType(DiskCreateOptionTypesPtrOutput{})
	pulumi.RegisterOutputType(HyperVGenerationTypesOutput{})
	pulumi.RegisterOutputType(HyperVGenerationTypesPtrOutput{})
	pulumi.RegisterOutputType(IPVersionOutput{})
	pulumi.RegisterOutputType(IPVersionPtrOutput{})
	pulumi.RegisterOutputType(InGuestPatchModeOutput{})
	pulumi.RegisterOutputType(InGuestPatchModePtrOutput{})
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
	pulumi.RegisterOutputType(ProximityPlacementGroupTypeOutput{})
	pulumi.RegisterOutputType(ProximityPlacementGroupTypePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(SettingNamesOutput{})
	pulumi.RegisterOutputType(SettingNamesPtrOutput{})
	pulumi.RegisterOutputType(StatusLevelTypesOutput{})
	pulumi.RegisterOutputType(StatusLevelTypesPtrOutput{})
	pulumi.RegisterOutputType(StorageAccountTypesOutput{})
	pulumi.RegisterOutputType(StorageAccountTypesPtrOutput{})
	pulumi.RegisterOutputType(UpgradeModeOutput{})
	pulumi.RegisterOutputType(UpgradeModePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineEvictionPolicyTypesOutput{})
	pulumi.RegisterOutputType(VirtualMachineEvictionPolicyTypesPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachinePriorityTypesOutput{})
	pulumi.RegisterOutputType(VirtualMachinePriorityTypesPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetScaleInRulesOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetScaleInRulesPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineSizeTypesOutput{})
	pulumi.RegisterOutputType(VirtualMachineSizeTypesPtrOutput{})
}
