


package v20210501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DeviceType string

const (
	DeviceTypeUnknown        = DeviceType("Unknown")
	DeviceTypeAzureStackEdge = DeviceType("AzureStackEdge")
)

func (DeviceType) ElementType() reflect.Type {
	return reflect.TypeOf((*DeviceType)(nil)).Elem()
}

func (e DeviceType) ToDeviceTypeOutput() DeviceTypeOutput {
	return pulumi.ToOutput(e).(DeviceTypeOutput)
}

func (e DeviceType) ToDeviceTypeOutputWithContext(ctx context.Context) DeviceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DeviceTypeOutput)
}

func (e DeviceType) ToDeviceTypePtrOutput() DeviceTypePtrOutput {
	return e.ToDeviceTypePtrOutputWithContext(context.Background())
}

func (e DeviceType) ToDeviceTypePtrOutputWithContext(ctx context.Context) DeviceTypePtrOutput {
	return DeviceType(e).ToDeviceTypeOutputWithContext(ctx).ToDeviceTypePtrOutputWithContext(ctx)
}

func (e DeviceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeviceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeviceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DeviceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DeviceTypeOutput struct{ *pulumi.OutputState }

func (DeviceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeviceType)(nil)).Elem()
}

func (o DeviceTypeOutput) ToDeviceTypeOutput() DeviceTypeOutput {
	return o
}

func (o DeviceTypeOutput) ToDeviceTypeOutputWithContext(ctx context.Context) DeviceTypeOutput {
	return o
}

func (o DeviceTypeOutput) ToDeviceTypePtrOutput() DeviceTypePtrOutput {
	return o.ToDeviceTypePtrOutputWithContext(context.Background())
}

func (o DeviceTypeOutput) ToDeviceTypePtrOutputWithContext(ctx context.Context) DeviceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeviceType) *DeviceType {
		return &v
	}).(DeviceTypePtrOutput)
}

func (o DeviceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DeviceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeviceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DeviceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeviceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeviceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DeviceTypePtrOutput struct{ *pulumi.OutputState }

func (DeviceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceType)(nil)).Elem()
}

func (o DeviceTypePtrOutput) ToDeviceTypePtrOutput() DeviceTypePtrOutput {
	return o
}

func (o DeviceTypePtrOutput) ToDeviceTypePtrOutputWithContext(ctx context.Context) DeviceTypePtrOutput {
	return o
}

func (o DeviceTypePtrOutput) Elem() DeviceTypeOutput {
	return o.ApplyT(func(v *DeviceType) DeviceType {
		if v != nil {
			return *v
		}
		var ret DeviceType
		return ret
	}).(DeviceTypeOutput)
}

func (o DeviceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeviceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DeviceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DeviceTypeInput interface {
	pulumi.Input

	ToDeviceTypeOutput() DeviceTypeOutput
	ToDeviceTypeOutputWithContext(context.Context) DeviceTypeOutput
}

var deviceTypePtrType = reflect.TypeOf((**DeviceType)(nil)).Elem()

type DeviceTypePtrInput interface {
	pulumi.Input

	ToDeviceTypePtrOutput() DeviceTypePtrOutput
	ToDeviceTypePtrOutputWithContext(context.Context) DeviceTypePtrOutput
}

type deviceTypePtr string

func DeviceTypePtr(v string) DeviceTypePtrInput {
	return (*deviceTypePtr)(&v)
}

func (*deviceTypePtr) ElementType() reflect.Type {
	return deviceTypePtrType
}

func (in *deviceTypePtr) ToDeviceTypePtrOutput() DeviceTypePtrOutput {
	return pulumi.ToOutput(in).(DeviceTypePtrOutput)
}

func (in *deviceTypePtr) ToDeviceTypePtrOutputWithContext(ctx context.Context) DeviceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DeviceTypePtrOutput)
}

type DiskCreateOptionTypes string

const (
	DiskCreateOptionTypesUnknown = DiskCreateOptionTypes("Unknown")
	DiskCreateOptionTypesEmpty   = DiskCreateOptionTypes("Empty")
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

type IPAllocationMethod string

const (
	IPAllocationMethodUnknown = IPAllocationMethod("Unknown")
	IPAllocationMethodStatic  = IPAllocationMethod("Static")
	IPAllocationMethodDynamic = IPAllocationMethod("Dynamic")
)

func (IPAllocationMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*IPAllocationMethod)(nil)).Elem()
}

func (e IPAllocationMethod) ToIPAllocationMethodOutput() IPAllocationMethodOutput {
	return pulumi.ToOutput(e).(IPAllocationMethodOutput)
}

func (e IPAllocationMethod) ToIPAllocationMethodOutputWithContext(ctx context.Context) IPAllocationMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IPAllocationMethodOutput)
}

func (e IPAllocationMethod) ToIPAllocationMethodPtrOutput() IPAllocationMethodPtrOutput {
	return e.ToIPAllocationMethodPtrOutputWithContext(context.Background())
}

func (e IPAllocationMethod) ToIPAllocationMethodPtrOutputWithContext(ctx context.Context) IPAllocationMethodPtrOutput {
	return IPAllocationMethod(e).ToIPAllocationMethodOutputWithContext(ctx).ToIPAllocationMethodPtrOutputWithContext(ctx)
}

func (e IPAllocationMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IPAllocationMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IPAllocationMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IPAllocationMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IPAllocationMethodOutput struct{ *pulumi.OutputState }

func (IPAllocationMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPAllocationMethod)(nil)).Elem()
}

func (o IPAllocationMethodOutput) ToIPAllocationMethodOutput() IPAllocationMethodOutput {
	return o
}

func (o IPAllocationMethodOutput) ToIPAllocationMethodOutputWithContext(ctx context.Context) IPAllocationMethodOutput {
	return o
}

func (o IPAllocationMethodOutput) ToIPAllocationMethodPtrOutput() IPAllocationMethodPtrOutput {
	return o.ToIPAllocationMethodPtrOutputWithContext(context.Background())
}

func (o IPAllocationMethodOutput) ToIPAllocationMethodPtrOutputWithContext(ctx context.Context) IPAllocationMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IPAllocationMethod) *IPAllocationMethod {
		return &v
	}).(IPAllocationMethodPtrOutput)
}

func (o IPAllocationMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IPAllocationMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IPAllocationMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IPAllocationMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IPAllocationMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IPAllocationMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IPAllocationMethodPtrOutput struct{ *pulumi.OutputState }

func (IPAllocationMethodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPAllocationMethod)(nil)).Elem()
}

func (o IPAllocationMethodPtrOutput) ToIPAllocationMethodPtrOutput() IPAllocationMethodPtrOutput {
	return o
}

func (o IPAllocationMethodPtrOutput) ToIPAllocationMethodPtrOutputWithContext(ctx context.Context) IPAllocationMethodPtrOutput {
	return o
}

func (o IPAllocationMethodPtrOutput) Elem() IPAllocationMethodOutput {
	return o.ApplyT(func(v *IPAllocationMethod) IPAllocationMethod {
		if v != nil {
			return *v
		}
		var ret IPAllocationMethod
		return ret
	}).(IPAllocationMethodOutput)
}

func (o IPAllocationMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IPAllocationMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IPAllocationMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IPAllocationMethodInput interface {
	pulumi.Input

	ToIPAllocationMethodOutput() IPAllocationMethodOutput
	ToIPAllocationMethodOutputWithContext(context.Context) IPAllocationMethodOutput
}

var ipallocationMethodPtrType = reflect.TypeOf((**IPAllocationMethod)(nil)).Elem()

type IPAllocationMethodPtrInput interface {
	pulumi.Input

	ToIPAllocationMethodPtrOutput() IPAllocationMethodPtrOutput
	ToIPAllocationMethodPtrOutputWithContext(context.Context) IPAllocationMethodPtrOutput
}

type ipallocationMethodPtr string

func IPAllocationMethodPtr(v string) IPAllocationMethodPtrInput {
	return (*ipallocationMethodPtr)(&v)
}

func (*ipallocationMethodPtr) ElementType() reflect.Type {
	return ipallocationMethodPtrType
}

func (in *ipallocationMethodPtr) ToIPAllocationMethodPtrOutput() IPAllocationMethodPtrOutput {
	return pulumi.ToOutput(in).(IPAllocationMethodPtrOutput)
}

func (in *ipallocationMethodPtr) ToIPAllocationMethodPtrOutputWithContext(ctx context.Context) IPAllocationMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IPAllocationMethodPtrOutput)
}

type IPVersion string

const (
	IPVersionUnknown = IPVersion("Unknown")
	IPVersionIPv4    = IPVersion("IPv4")
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

type NetworkFunctionRoleConfigurationType string

const (
	NetworkFunctionRoleConfigurationTypeUnknown        = NetworkFunctionRoleConfigurationType("Unknown")
	NetworkFunctionRoleConfigurationTypeVirtualMachine = NetworkFunctionRoleConfigurationType("VirtualMachine")
)

func (NetworkFunctionRoleConfigurationType) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkFunctionRoleConfigurationType)(nil)).Elem()
}

func (e NetworkFunctionRoleConfigurationType) ToNetworkFunctionRoleConfigurationTypeOutput() NetworkFunctionRoleConfigurationTypeOutput {
	return pulumi.ToOutput(e).(NetworkFunctionRoleConfigurationTypeOutput)
}

func (e NetworkFunctionRoleConfigurationType) ToNetworkFunctionRoleConfigurationTypeOutputWithContext(ctx context.Context) NetworkFunctionRoleConfigurationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NetworkFunctionRoleConfigurationTypeOutput)
}

func (e NetworkFunctionRoleConfigurationType) ToNetworkFunctionRoleConfigurationTypePtrOutput() NetworkFunctionRoleConfigurationTypePtrOutput {
	return e.ToNetworkFunctionRoleConfigurationTypePtrOutputWithContext(context.Background())
}

func (e NetworkFunctionRoleConfigurationType) ToNetworkFunctionRoleConfigurationTypePtrOutputWithContext(ctx context.Context) NetworkFunctionRoleConfigurationTypePtrOutput {
	return NetworkFunctionRoleConfigurationType(e).ToNetworkFunctionRoleConfigurationTypeOutputWithContext(ctx).ToNetworkFunctionRoleConfigurationTypePtrOutputWithContext(ctx)
}

func (e NetworkFunctionRoleConfigurationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkFunctionRoleConfigurationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkFunctionRoleConfigurationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NetworkFunctionRoleConfigurationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NetworkFunctionRoleConfigurationTypeOutput struct{ *pulumi.OutputState }

func (NetworkFunctionRoleConfigurationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkFunctionRoleConfigurationType)(nil)).Elem()
}

func (o NetworkFunctionRoleConfigurationTypeOutput) ToNetworkFunctionRoleConfigurationTypeOutput() NetworkFunctionRoleConfigurationTypeOutput {
	return o
}

func (o NetworkFunctionRoleConfigurationTypeOutput) ToNetworkFunctionRoleConfigurationTypeOutputWithContext(ctx context.Context) NetworkFunctionRoleConfigurationTypeOutput {
	return o
}

func (o NetworkFunctionRoleConfigurationTypeOutput) ToNetworkFunctionRoleConfigurationTypePtrOutput() NetworkFunctionRoleConfigurationTypePtrOutput {
	return o.ToNetworkFunctionRoleConfigurationTypePtrOutputWithContext(context.Background())
}

func (o NetworkFunctionRoleConfigurationTypeOutput) ToNetworkFunctionRoleConfigurationTypePtrOutputWithContext(ctx context.Context) NetworkFunctionRoleConfigurationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkFunctionRoleConfigurationType) *NetworkFunctionRoleConfigurationType {
		return &v
	}).(NetworkFunctionRoleConfigurationTypePtrOutput)
}

func (o NetworkFunctionRoleConfigurationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NetworkFunctionRoleConfigurationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkFunctionRoleConfigurationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NetworkFunctionRoleConfigurationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkFunctionRoleConfigurationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkFunctionRoleConfigurationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NetworkFunctionRoleConfigurationTypePtrOutput struct{ *pulumi.OutputState }

func (NetworkFunctionRoleConfigurationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkFunctionRoleConfigurationType)(nil)).Elem()
}

func (o NetworkFunctionRoleConfigurationTypePtrOutput) ToNetworkFunctionRoleConfigurationTypePtrOutput() NetworkFunctionRoleConfigurationTypePtrOutput {
	return o
}

func (o NetworkFunctionRoleConfigurationTypePtrOutput) ToNetworkFunctionRoleConfigurationTypePtrOutputWithContext(ctx context.Context) NetworkFunctionRoleConfigurationTypePtrOutput {
	return o
}

func (o NetworkFunctionRoleConfigurationTypePtrOutput) Elem() NetworkFunctionRoleConfigurationTypeOutput {
	return o.ApplyT(func(v *NetworkFunctionRoleConfigurationType) NetworkFunctionRoleConfigurationType {
		if v != nil {
			return *v
		}
		var ret NetworkFunctionRoleConfigurationType
		return ret
	}).(NetworkFunctionRoleConfigurationTypeOutput)
}

func (o NetworkFunctionRoleConfigurationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkFunctionRoleConfigurationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NetworkFunctionRoleConfigurationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NetworkFunctionRoleConfigurationTypeInput interface {
	pulumi.Input

	ToNetworkFunctionRoleConfigurationTypeOutput() NetworkFunctionRoleConfigurationTypeOutput
	ToNetworkFunctionRoleConfigurationTypeOutputWithContext(context.Context) NetworkFunctionRoleConfigurationTypeOutput
}

var networkFunctionRoleConfigurationTypePtrType = reflect.TypeOf((**NetworkFunctionRoleConfigurationType)(nil)).Elem()

type NetworkFunctionRoleConfigurationTypePtrInput interface {
	pulumi.Input

	ToNetworkFunctionRoleConfigurationTypePtrOutput() NetworkFunctionRoleConfigurationTypePtrOutput
	ToNetworkFunctionRoleConfigurationTypePtrOutputWithContext(context.Context) NetworkFunctionRoleConfigurationTypePtrOutput
}

type networkFunctionRoleConfigurationTypePtr string

func NetworkFunctionRoleConfigurationTypePtr(v string) NetworkFunctionRoleConfigurationTypePtrInput {
	return (*networkFunctionRoleConfigurationTypePtr)(&v)
}

func (*networkFunctionRoleConfigurationTypePtr) ElementType() reflect.Type {
	return networkFunctionRoleConfigurationTypePtrType
}

func (in *networkFunctionRoleConfigurationTypePtr) ToNetworkFunctionRoleConfigurationTypePtrOutput() NetworkFunctionRoleConfigurationTypePtrOutput {
	return pulumi.ToOutput(in).(NetworkFunctionRoleConfigurationTypePtrOutput)
}

func (in *networkFunctionRoleConfigurationTypePtr) ToNetworkFunctionRoleConfigurationTypePtrOutputWithContext(ctx context.Context) NetworkFunctionRoleConfigurationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NetworkFunctionRoleConfigurationTypePtrOutput)
}

type NetworkFunctionType string

const (
	NetworkFunctionTypeUnknown                      = NetworkFunctionType("Unknown")
	NetworkFunctionTypeVirtualNetworkFunction       = NetworkFunctionType("VirtualNetworkFunction")
	NetworkFunctionTypeContainerizedNetworkFunction = NetworkFunctionType("ContainerizedNetworkFunction")
)

func (NetworkFunctionType) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkFunctionType)(nil)).Elem()
}

func (e NetworkFunctionType) ToNetworkFunctionTypeOutput() NetworkFunctionTypeOutput {
	return pulumi.ToOutput(e).(NetworkFunctionTypeOutput)
}

func (e NetworkFunctionType) ToNetworkFunctionTypeOutputWithContext(ctx context.Context) NetworkFunctionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NetworkFunctionTypeOutput)
}

func (e NetworkFunctionType) ToNetworkFunctionTypePtrOutput() NetworkFunctionTypePtrOutput {
	return e.ToNetworkFunctionTypePtrOutputWithContext(context.Background())
}

func (e NetworkFunctionType) ToNetworkFunctionTypePtrOutputWithContext(ctx context.Context) NetworkFunctionTypePtrOutput {
	return NetworkFunctionType(e).ToNetworkFunctionTypeOutputWithContext(ctx).ToNetworkFunctionTypePtrOutputWithContext(ctx)
}

func (e NetworkFunctionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkFunctionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkFunctionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NetworkFunctionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NetworkFunctionTypeOutput struct{ *pulumi.OutputState }

func (NetworkFunctionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkFunctionType)(nil)).Elem()
}

func (o NetworkFunctionTypeOutput) ToNetworkFunctionTypeOutput() NetworkFunctionTypeOutput {
	return o
}

func (o NetworkFunctionTypeOutput) ToNetworkFunctionTypeOutputWithContext(ctx context.Context) NetworkFunctionTypeOutput {
	return o
}

func (o NetworkFunctionTypeOutput) ToNetworkFunctionTypePtrOutput() NetworkFunctionTypePtrOutput {
	return o.ToNetworkFunctionTypePtrOutputWithContext(context.Background())
}

func (o NetworkFunctionTypeOutput) ToNetworkFunctionTypePtrOutputWithContext(ctx context.Context) NetworkFunctionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkFunctionType) *NetworkFunctionType {
		return &v
	}).(NetworkFunctionTypePtrOutput)
}

func (o NetworkFunctionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NetworkFunctionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkFunctionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NetworkFunctionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkFunctionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkFunctionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NetworkFunctionTypePtrOutput struct{ *pulumi.OutputState }

func (NetworkFunctionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkFunctionType)(nil)).Elem()
}

func (o NetworkFunctionTypePtrOutput) ToNetworkFunctionTypePtrOutput() NetworkFunctionTypePtrOutput {
	return o
}

func (o NetworkFunctionTypePtrOutput) ToNetworkFunctionTypePtrOutputWithContext(ctx context.Context) NetworkFunctionTypePtrOutput {
	return o
}

func (o NetworkFunctionTypePtrOutput) Elem() NetworkFunctionTypeOutput {
	return o.ApplyT(func(v *NetworkFunctionType) NetworkFunctionType {
		if v != nil {
			return *v
		}
		var ret NetworkFunctionType
		return ret
	}).(NetworkFunctionTypeOutput)
}

func (o NetworkFunctionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkFunctionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NetworkFunctionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NetworkFunctionTypeInput interface {
	pulumi.Input

	ToNetworkFunctionTypeOutput() NetworkFunctionTypeOutput
	ToNetworkFunctionTypeOutputWithContext(context.Context) NetworkFunctionTypeOutput
}

var networkFunctionTypePtrType = reflect.TypeOf((**NetworkFunctionType)(nil)).Elem()

type NetworkFunctionTypePtrInput interface {
	pulumi.Input

	ToNetworkFunctionTypePtrOutput() NetworkFunctionTypePtrOutput
	ToNetworkFunctionTypePtrOutputWithContext(context.Context) NetworkFunctionTypePtrOutput
}

type networkFunctionTypePtr string

func NetworkFunctionTypePtr(v string) NetworkFunctionTypePtrInput {
	return (*networkFunctionTypePtr)(&v)
}

func (*networkFunctionTypePtr) ElementType() reflect.Type {
	return networkFunctionTypePtrType
}

func (in *networkFunctionTypePtr) ToNetworkFunctionTypePtrOutput() NetworkFunctionTypePtrOutput {
	return pulumi.ToOutput(in).(NetworkFunctionTypePtrOutput)
}

func (in *networkFunctionTypePtr) ToNetworkFunctionTypePtrOutputWithContext(ctx context.Context) NetworkFunctionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NetworkFunctionTypePtrOutput)
}

type OperatingSystemTypes string

const (
	OperatingSystemTypesUnknown = OperatingSystemTypes("Unknown")
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

type SkuDeploymentMode string

const (
	SkuDeploymentModeUnknown         = SkuDeploymentMode("Unknown")
	SkuDeploymentModeAzure           = SkuDeploymentMode("Azure")
	SkuDeploymentModePrivateEdgeZone = SkuDeploymentMode("PrivateEdgeZone")
)

func (SkuDeploymentMode) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuDeploymentMode)(nil)).Elem()
}

func (e SkuDeploymentMode) ToSkuDeploymentModeOutput() SkuDeploymentModeOutput {
	return pulumi.ToOutput(e).(SkuDeploymentModeOutput)
}

func (e SkuDeploymentMode) ToSkuDeploymentModeOutputWithContext(ctx context.Context) SkuDeploymentModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuDeploymentModeOutput)
}

func (e SkuDeploymentMode) ToSkuDeploymentModePtrOutput() SkuDeploymentModePtrOutput {
	return e.ToSkuDeploymentModePtrOutputWithContext(context.Background())
}

func (e SkuDeploymentMode) ToSkuDeploymentModePtrOutputWithContext(ctx context.Context) SkuDeploymentModePtrOutput {
	return SkuDeploymentMode(e).ToSkuDeploymentModeOutputWithContext(ctx).ToSkuDeploymentModePtrOutputWithContext(ctx)
}

func (e SkuDeploymentMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuDeploymentMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuDeploymentMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuDeploymentMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuDeploymentModeOutput struct{ *pulumi.OutputState }

func (SkuDeploymentModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuDeploymentMode)(nil)).Elem()
}

func (o SkuDeploymentModeOutput) ToSkuDeploymentModeOutput() SkuDeploymentModeOutput {
	return o
}

func (o SkuDeploymentModeOutput) ToSkuDeploymentModeOutputWithContext(ctx context.Context) SkuDeploymentModeOutput {
	return o
}

func (o SkuDeploymentModeOutput) ToSkuDeploymentModePtrOutput() SkuDeploymentModePtrOutput {
	return o.ToSkuDeploymentModePtrOutputWithContext(context.Background())
}

func (o SkuDeploymentModeOutput) ToSkuDeploymentModePtrOutputWithContext(ctx context.Context) SkuDeploymentModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuDeploymentMode) *SkuDeploymentMode {
		return &v
	}).(SkuDeploymentModePtrOutput)
}

func (o SkuDeploymentModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuDeploymentModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuDeploymentMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuDeploymentModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuDeploymentModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuDeploymentMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuDeploymentModePtrOutput struct{ *pulumi.OutputState }

func (SkuDeploymentModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuDeploymentMode)(nil)).Elem()
}

func (o SkuDeploymentModePtrOutput) ToSkuDeploymentModePtrOutput() SkuDeploymentModePtrOutput {
	return o
}

func (o SkuDeploymentModePtrOutput) ToSkuDeploymentModePtrOutputWithContext(ctx context.Context) SkuDeploymentModePtrOutput {
	return o
}

func (o SkuDeploymentModePtrOutput) Elem() SkuDeploymentModeOutput {
	return o.ApplyT(func(v *SkuDeploymentMode) SkuDeploymentMode {
		if v != nil {
			return *v
		}
		var ret SkuDeploymentMode
		return ret
	}).(SkuDeploymentModeOutput)
}

func (o SkuDeploymentModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuDeploymentModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuDeploymentMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SkuDeploymentModeInput interface {
	pulumi.Input

	ToSkuDeploymentModeOutput() SkuDeploymentModeOutput
	ToSkuDeploymentModeOutputWithContext(context.Context) SkuDeploymentModeOutput
}

var skuDeploymentModePtrType = reflect.TypeOf((**SkuDeploymentMode)(nil)).Elem()

type SkuDeploymentModePtrInput interface {
	pulumi.Input

	ToSkuDeploymentModePtrOutput() SkuDeploymentModePtrOutput
	ToSkuDeploymentModePtrOutputWithContext(context.Context) SkuDeploymentModePtrOutput
}

type skuDeploymentModePtr string

func SkuDeploymentModePtr(v string) SkuDeploymentModePtrInput {
	return (*skuDeploymentModePtr)(&v)
}

func (*skuDeploymentModePtr) ElementType() reflect.Type {
	return skuDeploymentModePtrType
}

func (in *skuDeploymentModePtr) ToSkuDeploymentModePtrOutput() SkuDeploymentModePtrOutput {
	return pulumi.ToOutput(in).(SkuDeploymentModePtrOutput)
}

func (in *skuDeploymentModePtr) ToSkuDeploymentModePtrOutputWithContext(ctx context.Context) SkuDeploymentModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuDeploymentModePtrOutput)
}

type SkuType string

const (
	SkuTypeUnknown           = SkuType("Unknown")
	SkuTypeEvolvedPacketCore = SkuType("EvolvedPacketCore")
	SkuTypeSDWAN             = SkuType("SDWAN")
	SkuTypeFirewall          = SkuType("Firewall")
)

func (SkuType) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuType)(nil)).Elem()
}

func (e SkuType) ToSkuTypeOutput() SkuTypeOutput {
	return pulumi.ToOutput(e).(SkuTypeOutput)
}

func (e SkuType) ToSkuTypeOutputWithContext(ctx context.Context) SkuTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuTypeOutput)
}

func (e SkuType) ToSkuTypePtrOutput() SkuTypePtrOutput {
	return e.ToSkuTypePtrOutputWithContext(context.Background())
}

func (e SkuType) ToSkuTypePtrOutputWithContext(ctx context.Context) SkuTypePtrOutput {
	return SkuType(e).ToSkuTypeOutputWithContext(ctx).ToSkuTypePtrOutputWithContext(ctx)
}

func (e SkuType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuTypeOutput struct{ *pulumi.OutputState }

func (SkuTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuType)(nil)).Elem()
}

func (o SkuTypeOutput) ToSkuTypeOutput() SkuTypeOutput {
	return o
}

func (o SkuTypeOutput) ToSkuTypeOutputWithContext(ctx context.Context) SkuTypeOutput {
	return o
}

func (o SkuTypeOutput) ToSkuTypePtrOutput() SkuTypePtrOutput {
	return o.ToSkuTypePtrOutputWithContext(context.Background())
}

func (o SkuTypeOutput) ToSkuTypePtrOutputWithContext(ctx context.Context) SkuTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuType) *SkuType {
		return &v
	}).(SkuTypePtrOutput)
}

func (o SkuTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuTypePtrOutput struct{ *pulumi.OutputState }

func (SkuTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuType)(nil)).Elem()
}

func (o SkuTypePtrOutput) ToSkuTypePtrOutput() SkuTypePtrOutput {
	return o
}

func (o SkuTypePtrOutput) ToSkuTypePtrOutputWithContext(ctx context.Context) SkuTypePtrOutput {
	return o
}

func (o SkuTypePtrOutput) Elem() SkuTypeOutput {
	return o.ApplyT(func(v *SkuType) SkuType {
		if v != nil {
			return *v
		}
		var ret SkuType
		return ret
	}).(SkuTypeOutput)
}

func (o SkuTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SkuTypeInput interface {
	pulumi.Input

	ToSkuTypeOutput() SkuTypeOutput
	ToSkuTypeOutputWithContext(context.Context) SkuTypeOutput
}

var skuTypePtrType = reflect.TypeOf((**SkuType)(nil)).Elem()

type SkuTypePtrInput interface {
	pulumi.Input

	ToSkuTypePtrOutput() SkuTypePtrOutput
	ToSkuTypePtrOutputWithContext(context.Context) SkuTypePtrOutput
}

type skuTypePtr string

func SkuTypePtr(v string) SkuTypePtrInput {
	return (*skuTypePtr)(&v)
}

func (*skuTypePtr) ElementType() reflect.Type {
	return skuTypePtrType
}

func (in *skuTypePtr) ToSkuTypePtrOutput() SkuTypePtrOutput {
	return pulumi.ToOutput(in).(SkuTypePtrOutput)
}

func (in *skuTypePtr) ToSkuTypePtrOutputWithContext(ctx context.Context) SkuTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuTypePtrOutput)
}

type VMSwitchType string

const (
	VMSwitchTypeUnknown    = VMSwitchType("Unknown")
	VMSwitchTypeManagement = VMSwitchType("Management")
	VMSwitchTypeWan        = VMSwitchType("Wan")
	VMSwitchTypeLan        = VMSwitchType("Lan")
)

func (VMSwitchType) ElementType() reflect.Type {
	return reflect.TypeOf((*VMSwitchType)(nil)).Elem()
}

func (e VMSwitchType) ToVMSwitchTypeOutput() VMSwitchTypeOutput {
	return pulumi.ToOutput(e).(VMSwitchTypeOutput)
}

func (e VMSwitchType) ToVMSwitchTypeOutputWithContext(ctx context.Context) VMSwitchTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VMSwitchTypeOutput)
}

func (e VMSwitchType) ToVMSwitchTypePtrOutput() VMSwitchTypePtrOutput {
	return e.ToVMSwitchTypePtrOutputWithContext(context.Background())
}

func (e VMSwitchType) ToVMSwitchTypePtrOutputWithContext(ctx context.Context) VMSwitchTypePtrOutput {
	return VMSwitchType(e).ToVMSwitchTypeOutputWithContext(ctx).ToVMSwitchTypePtrOutputWithContext(ctx)
}

func (e VMSwitchType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VMSwitchType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VMSwitchType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VMSwitchType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VMSwitchTypeOutput struct{ *pulumi.OutputState }

func (VMSwitchTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMSwitchType)(nil)).Elem()
}

func (o VMSwitchTypeOutput) ToVMSwitchTypeOutput() VMSwitchTypeOutput {
	return o
}

func (o VMSwitchTypeOutput) ToVMSwitchTypeOutputWithContext(ctx context.Context) VMSwitchTypeOutput {
	return o
}

func (o VMSwitchTypeOutput) ToVMSwitchTypePtrOutput() VMSwitchTypePtrOutput {
	return o.ToVMSwitchTypePtrOutputWithContext(context.Background())
}

func (o VMSwitchTypeOutput) ToVMSwitchTypePtrOutputWithContext(ctx context.Context) VMSwitchTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VMSwitchType) *VMSwitchType {
		return &v
	}).(VMSwitchTypePtrOutput)
}

func (o VMSwitchTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VMSwitchTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VMSwitchType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VMSwitchTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VMSwitchTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VMSwitchType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VMSwitchTypePtrOutput struct{ *pulumi.OutputState }

func (VMSwitchTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VMSwitchType)(nil)).Elem()
}

func (o VMSwitchTypePtrOutput) ToVMSwitchTypePtrOutput() VMSwitchTypePtrOutput {
	return o
}

func (o VMSwitchTypePtrOutput) ToVMSwitchTypePtrOutputWithContext(ctx context.Context) VMSwitchTypePtrOutput {
	return o
}

func (o VMSwitchTypePtrOutput) Elem() VMSwitchTypeOutput {
	return o.ApplyT(func(v *VMSwitchType) VMSwitchType {
		if v != nil {
			return *v
		}
		var ret VMSwitchType
		return ret
	}).(VMSwitchTypeOutput)
}

func (o VMSwitchTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VMSwitchTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VMSwitchType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type VMSwitchTypeInput interface {
	pulumi.Input

	ToVMSwitchTypeOutput() VMSwitchTypeOutput
	ToVMSwitchTypeOutputWithContext(context.Context) VMSwitchTypeOutput
}

var vmswitchTypePtrType = reflect.TypeOf((**VMSwitchType)(nil)).Elem()

type VMSwitchTypePtrInput interface {
	pulumi.Input

	ToVMSwitchTypePtrOutput() VMSwitchTypePtrOutput
	ToVMSwitchTypePtrOutputWithContext(context.Context) VMSwitchTypePtrOutput
}

type vmswitchTypePtr string

func VMSwitchTypePtr(v string) VMSwitchTypePtrInput {
	return (*vmswitchTypePtr)(&v)
}

func (*vmswitchTypePtr) ElementType() reflect.Type {
	return vmswitchTypePtrType
}

func (in *vmswitchTypePtr) ToVMSwitchTypePtrOutput() VMSwitchTypePtrOutput {
	return pulumi.ToOutput(in).(VMSwitchTypePtrOutput)
}

func (in *vmswitchTypePtr) ToVMSwitchTypePtrOutputWithContext(ctx context.Context) VMSwitchTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VMSwitchTypePtrOutput)
}

type VirtualMachineSizeTypes string

const (
	VirtualMachineSizeTypesUnknown           = VirtualMachineSizeTypes("Unknown")
	VirtualMachineSizeTypes_Standard_D1_v2   = VirtualMachineSizeTypes("Standard_D1_v2")
	VirtualMachineSizeTypes_Standard_D2_v2   = VirtualMachineSizeTypes("Standard_D2_v2")
	VirtualMachineSizeTypes_Standard_D3_v2   = VirtualMachineSizeTypes("Standard_D3_v2")
	VirtualMachineSizeTypes_Standard_D4_v2   = VirtualMachineSizeTypes("Standard_D4_v2")
	VirtualMachineSizeTypes_Standard_D5_v2   = VirtualMachineSizeTypes("Standard_D5_v2")
	VirtualMachineSizeTypes_Standard_D11_v2  = VirtualMachineSizeTypes("Standard_D11_v2")
	VirtualMachineSizeTypes_Standard_D12_v2  = VirtualMachineSizeTypes("Standard_D12_v2")
	VirtualMachineSizeTypes_Standard_D13_v2  = VirtualMachineSizeTypes("Standard_D13_v2")
	VirtualMachineSizeTypes_Standard_DS1_v2  = VirtualMachineSizeTypes("Standard_DS1_v2")
	VirtualMachineSizeTypes_Standard_DS2_v2  = VirtualMachineSizeTypes("Standard_DS2_v2")
	VirtualMachineSizeTypes_Standard_DS3_v2  = VirtualMachineSizeTypes("Standard_DS3_v2")
	VirtualMachineSizeTypes_Standard_DS4_v2  = VirtualMachineSizeTypes("Standard_DS4_v2")
	VirtualMachineSizeTypes_Standard_DS5_v2  = VirtualMachineSizeTypes("Standard_DS5_v2")
	VirtualMachineSizeTypes_Standard_DS11_v2 = VirtualMachineSizeTypes("Standard_DS11_v2")
	VirtualMachineSizeTypes_Standard_DS12_v2 = VirtualMachineSizeTypes("Standard_DS12_v2")
	VirtualMachineSizeTypes_Standard_DS13_v2 = VirtualMachineSizeTypes("Standard_DS13_v2")
	VirtualMachineSizeTypes_Standard_F1      = VirtualMachineSizeTypes("Standard_F1")
	VirtualMachineSizeTypes_Standard_F2      = VirtualMachineSizeTypes("Standard_F2")
	VirtualMachineSizeTypes_Standard_F4      = VirtualMachineSizeTypes("Standard_F4")
	VirtualMachineSizeTypes_Standard_F8      = VirtualMachineSizeTypes("Standard_F8")
	VirtualMachineSizeTypes_Standard_F16     = VirtualMachineSizeTypes("Standard_F16")
	VirtualMachineSizeTypes_Standard_F1s     = VirtualMachineSizeTypes("Standard_F1s")
	VirtualMachineSizeTypes_Standard_F2s     = VirtualMachineSizeTypes("Standard_F2s")
	VirtualMachineSizeTypes_Standard_F4s     = VirtualMachineSizeTypes("Standard_F4s")
	VirtualMachineSizeTypes_Standard_F8s     = VirtualMachineSizeTypes("Standard_F8s")
	VirtualMachineSizeTypes_Standard_F16s    = VirtualMachineSizeTypes("Standard_F16s")
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
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceTypeInput)(nil)).Elem(), DeviceType("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceTypePtrInput)(nil)).Elem(), DeviceType("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*DiskCreateOptionTypesInput)(nil)).Elem(), DiskCreateOptionTypes("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*DiskCreateOptionTypesPtrInput)(nil)).Elem(), DiskCreateOptionTypes("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*IPAllocationMethodInput)(nil)).Elem(), IPAllocationMethod("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*IPAllocationMethodPtrInput)(nil)).Elem(), IPAllocationMethod("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*IPVersionInput)(nil)).Elem(), IPVersion("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*IPVersionPtrInput)(nil)).Elem(), IPVersion("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkFunctionRoleConfigurationTypeInput)(nil)).Elem(), NetworkFunctionRoleConfigurationType("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkFunctionRoleConfigurationTypePtrInput)(nil)).Elem(), NetworkFunctionRoleConfigurationType("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkFunctionTypeInput)(nil)).Elem(), NetworkFunctionType("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkFunctionTypePtrInput)(nil)).Elem(), NetworkFunctionType("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*OperatingSystemTypesInput)(nil)).Elem(), OperatingSystemTypes("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*OperatingSystemTypesPtrInput)(nil)).Elem(), OperatingSystemTypes("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuDeploymentModeInput)(nil)).Elem(), SkuDeploymentMode("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuDeploymentModePtrInput)(nil)).Elem(), SkuDeploymentMode("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuTypeInput)(nil)).Elem(), SkuType("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuTypePtrInput)(nil)).Elem(), SkuType("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*VMSwitchTypeInput)(nil)).Elem(), VMSwitchType("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*VMSwitchTypePtrInput)(nil)).Elem(), VMSwitchType("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineSizeTypesInput)(nil)).Elem(), VirtualMachineSizeTypes("Unknown"))
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineSizeTypesPtrInput)(nil)).Elem(), VirtualMachineSizeTypes("Unknown"))
	pulumi.RegisterOutputType(DeviceTypeOutput{})
	pulumi.RegisterOutputType(DeviceTypePtrOutput{})
	pulumi.RegisterOutputType(DiskCreateOptionTypesOutput{})
	pulumi.RegisterOutputType(DiskCreateOptionTypesPtrOutput{})
	pulumi.RegisterOutputType(IPAllocationMethodOutput{})
	pulumi.RegisterOutputType(IPAllocationMethodPtrOutput{})
	pulumi.RegisterOutputType(IPVersionOutput{})
	pulumi.RegisterOutputType(IPVersionPtrOutput{})
	pulumi.RegisterOutputType(NetworkFunctionRoleConfigurationTypeOutput{})
	pulumi.RegisterOutputType(NetworkFunctionRoleConfigurationTypePtrOutput{})
	pulumi.RegisterOutputType(NetworkFunctionTypeOutput{})
	pulumi.RegisterOutputType(NetworkFunctionTypePtrOutput{})
	pulumi.RegisterOutputType(OperatingSystemTypesOutput{})
	pulumi.RegisterOutputType(OperatingSystemTypesPtrOutput{})
	pulumi.RegisterOutputType(SkuDeploymentModeOutput{})
	pulumi.RegisterOutputType(SkuDeploymentModePtrOutput{})
	pulumi.RegisterOutputType(SkuTypeOutput{})
	pulumi.RegisterOutputType(SkuTypePtrOutput{})
	pulumi.RegisterOutputType(VMSwitchTypeOutput{})
	pulumi.RegisterOutputType(VMSwitchTypePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineSizeTypesOutput{})
	pulumi.RegisterOutputType(VirtualMachineSizeTypesPtrOutput{})
}
