


package v20200717preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AddonType string

const (
	AddonTypeSRM = AddonType("SRM")
	AddonTypeVR  = AddonType("VR")
)

func (AddonType) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonType)(nil)).Elem()
}

func (e AddonType) ToAddonTypeOutput() AddonTypeOutput {
	return pulumi.ToOutput(e).(AddonTypeOutput)
}

func (e AddonType) ToAddonTypeOutputWithContext(ctx context.Context) AddonTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AddonTypeOutput)
}

func (e AddonType) ToAddonTypePtrOutput() AddonTypePtrOutput {
	return e.ToAddonTypePtrOutputWithContext(context.Background())
}

func (e AddonType) ToAddonTypePtrOutputWithContext(ctx context.Context) AddonTypePtrOutput {
	return AddonType(e).ToAddonTypeOutputWithContext(ctx).ToAddonTypePtrOutputWithContext(ctx)
}

func (e AddonType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AddonType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AddonType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AddonType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AddonTypeOutput struct{ *pulumi.OutputState }

func (AddonTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonType)(nil)).Elem()
}

func (o AddonTypeOutput) ToAddonTypeOutput() AddonTypeOutput {
	return o
}

func (o AddonTypeOutput) ToAddonTypeOutputWithContext(ctx context.Context) AddonTypeOutput {
	return o
}

func (o AddonTypeOutput) ToAddonTypePtrOutput() AddonTypePtrOutput {
	return o.ToAddonTypePtrOutputWithContext(context.Background())
}

func (o AddonTypeOutput) ToAddonTypePtrOutputWithContext(ctx context.Context) AddonTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AddonType) *AddonType {
		return &v
	}).(AddonTypePtrOutput)
}

func (o AddonTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AddonTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AddonType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AddonTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AddonTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AddonType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AddonTypePtrOutput struct{ *pulumi.OutputState }

func (AddonTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddonType)(nil)).Elem()
}

func (o AddonTypePtrOutput) ToAddonTypePtrOutput() AddonTypePtrOutput {
	return o
}

func (o AddonTypePtrOutput) ToAddonTypePtrOutputWithContext(ctx context.Context) AddonTypePtrOutput {
	return o
}

func (o AddonTypePtrOutput) Elem() AddonTypeOutput {
	return o.ApplyT(func(v *AddonType) AddonType {
		if v != nil {
			return *v
		}
		var ret AddonType
		return ret
	}).(AddonTypeOutput)
}

func (o AddonTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AddonTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AddonType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AddonTypeInput interface {
	pulumi.Input

	ToAddonTypeOutput() AddonTypeOutput
	ToAddonTypeOutputWithContext(context.Context) AddonTypeOutput
}

var addonTypePtrType = reflect.TypeOf((**AddonType)(nil)).Elem()

type AddonTypePtrInput interface {
	pulumi.Input

	ToAddonTypePtrOutput() AddonTypePtrOutput
	ToAddonTypePtrOutputWithContext(context.Context) AddonTypePtrOutput
}

type addonTypePtr string

func AddonTypePtr(v string) AddonTypePtrInput {
	return (*addonTypePtr)(&v)
}

func (*addonTypePtr) ElementType() reflect.Type {
	return addonTypePtrType
}

func (in *addonTypePtr) ToAddonTypePtrOutput() AddonTypePtrOutput {
	return pulumi.ToOutput(in).(AddonTypePtrOutput)
}

func (in *addonTypePtr) ToAddonTypePtrOutputWithContext(ctx context.Context) AddonTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AddonTypePtrOutput)
}

type DhcpTypeEnum string

const (
	DhcpTypeEnum_SERVER_RELAY = DhcpTypeEnum("SERVER, RELAY")
)

func (DhcpTypeEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*DhcpTypeEnum)(nil)).Elem()
}

func (e DhcpTypeEnum) ToDhcpTypeEnumOutput() DhcpTypeEnumOutput {
	return pulumi.ToOutput(e).(DhcpTypeEnumOutput)
}

func (e DhcpTypeEnum) ToDhcpTypeEnumOutputWithContext(ctx context.Context) DhcpTypeEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DhcpTypeEnumOutput)
}

func (e DhcpTypeEnum) ToDhcpTypeEnumPtrOutput() DhcpTypeEnumPtrOutput {
	return e.ToDhcpTypeEnumPtrOutputWithContext(context.Background())
}

func (e DhcpTypeEnum) ToDhcpTypeEnumPtrOutputWithContext(ctx context.Context) DhcpTypeEnumPtrOutput {
	return DhcpTypeEnum(e).ToDhcpTypeEnumOutputWithContext(ctx).ToDhcpTypeEnumPtrOutputWithContext(ctx)
}

func (e DhcpTypeEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DhcpTypeEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DhcpTypeEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DhcpTypeEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DhcpTypeEnumOutput struct{ *pulumi.OutputState }

func (DhcpTypeEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DhcpTypeEnum)(nil)).Elem()
}

func (o DhcpTypeEnumOutput) ToDhcpTypeEnumOutput() DhcpTypeEnumOutput {
	return o
}

func (o DhcpTypeEnumOutput) ToDhcpTypeEnumOutputWithContext(ctx context.Context) DhcpTypeEnumOutput {
	return o
}

func (o DhcpTypeEnumOutput) ToDhcpTypeEnumPtrOutput() DhcpTypeEnumPtrOutput {
	return o.ToDhcpTypeEnumPtrOutputWithContext(context.Background())
}

func (o DhcpTypeEnumOutput) ToDhcpTypeEnumPtrOutputWithContext(ctx context.Context) DhcpTypeEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DhcpTypeEnum) *DhcpTypeEnum {
		return &v
	}).(DhcpTypeEnumPtrOutput)
}

func (o DhcpTypeEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DhcpTypeEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DhcpTypeEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DhcpTypeEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DhcpTypeEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DhcpTypeEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DhcpTypeEnumPtrOutput struct{ *pulumi.OutputState }

func (DhcpTypeEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpTypeEnum)(nil)).Elem()
}

func (o DhcpTypeEnumPtrOutput) ToDhcpTypeEnumPtrOutput() DhcpTypeEnumPtrOutput {
	return o
}

func (o DhcpTypeEnumPtrOutput) ToDhcpTypeEnumPtrOutputWithContext(ctx context.Context) DhcpTypeEnumPtrOutput {
	return o
}

func (o DhcpTypeEnumPtrOutput) Elem() DhcpTypeEnumOutput {
	return o.ApplyT(func(v *DhcpTypeEnum) DhcpTypeEnum {
		if v != nil {
			return *v
		}
		var ret DhcpTypeEnum
		return ret
	}).(DhcpTypeEnumOutput)
}

func (o DhcpTypeEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DhcpTypeEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DhcpTypeEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DhcpTypeEnumInput interface {
	pulumi.Input

	ToDhcpTypeEnumOutput() DhcpTypeEnumOutput
	ToDhcpTypeEnumOutputWithContext(context.Context) DhcpTypeEnumOutput
}

var dhcpTypeEnumPtrType = reflect.TypeOf((**DhcpTypeEnum)(nil)).Elem()

type DhcpTypeEnumPtrInput interface {
	pulumi.Input

	ToDhcpTypeEnumPtrOutput() DhcpTypeEnumPtrOutput
	ToDhcpTypeEnumPtrOutputWithContext(context.Context) DhcpTypeEnumPtrOutput
}

type dhcpTypeEnumPtr string

func DhcpTypeEnumPtr(v string) DhcpTypeEnumPtrInput {
	return (*dhcpTypeEnumPtr)(&v)
}

func (*dhcpTypeEnumPtr) ElementType() reflect.Type {
	return dhcpTypeEnumPtrType
}

func (in *dhcpTypeEnumPtr) ToDhcpTypeEnumPtrOutput() DhcpTypeEnumPtrOutput {
	return pulumi.ToOutput(in).(DhcpTypeEnumPtrOutput)
}

func (in *dhcpTypeEnumPtr) ToDhcpTypeEnumPtrOutputWithContext(ctx context.Context) DhcpTypeEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DhcpTypeEnumPtrOutput)
}

type DnsServiceLogLevelEnum string

const (
	DnsServiceLogLevelEnumDEBUG   = DnsServiceLogLevelEnum("DEBUG")
	DnsServiceLogLevelEnumINFO    = DnsServiceLogLevelEnum("INFO")
	DnsServiceLogLevelEnumWARNING = DnsServiceLogLevelEnum("WARNING")
	DnsServiceLogLevelEnumERROR   = DnsServiceLogLevelEnum("ERROR")
	DnsServiceLogLevelEnumFATAL   = DnsServiceLogLevelEnum("FATAL")
)

func (DnsServiceLogLevelEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsServiceLogLevelEnum)(nil)).Elem()
}

func (e DnsServiceLogLevelEnum) ToDnsServiceLogLevelEnumOutput() DnsServiceLogLevelEnumOutput {
	return pulumi.ToOutput(e).(DnsServiceLogLevelEnumOutput)
}

func (e DnsServiceLogLevelEnum) ToDnsServiceLogLevelEnumOutputWithContext(ctx context.Context) DnsServiceLogLevelEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DnsServiceLogLevelEnumOutput)
}

func (e DnsServiceLogLevelEnum) ToDnsServiceLogLevelEnumPtrOutput() DnsServiceLogLevelEnumPtrOutput {
	return e.ToDnsServiceLogLevelEnumPtrOutputWithContext(context.Background())
}

func (e DnsServiceLogLevelEnum) ToDnsServiceLogLevelEnumPtrOutputWithContext(ctx context.Context) DnsServiceLogLevelEnumPtrOutput {
	return DnsServiceLogLevelEnum(e).ToDnsServiceLogLevelEnumOutputWithContext(ctx).ToDnsServiceLogLevelEnumPtrOutputWithContext(ctx)
}

func (e DnsServiceLogLevelEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DnsServiceLogLevelEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DnsServiceLogLevelEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DnsServiceLogLevelEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DnsServiceLogLevelEnumOutput struct{ *pulumi.OutputState }

func (DnsServiceLogLevelEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsServiceLogLevelEnum)(nil)).Elem()
}

func (o DnsServiceLogLevelEnumOutput) ToDnsServiceLogLevelEnumOutput() DnsServiceLogLevelEnumOutput {
	return o
}

func (o DnsServiceLogLevelEnumOutput) ToDnsServiceLogLevelEnumOutputWithContext(ctx context.Context) DnsServiceLogLevelEnumOutput {
	return o
}

func (o DnsServiceLogLevelEnumOutput) ToDnsServiceLogLevelEnumPtrOutput() DnsServiceLogLevelEnumPtrOutput {
	return o.ToDnsServiceLogLevelEnumPtrOutputWithContext(context.Background())
}

func (o DnsServiceLogLevelEnumOutput) ToDnsServiceLogLevelEnumPtrOutputWithContext(ctx context.Context) DnsServiceLogLevelEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DnsServiceLogLevelEnum) *DnsServiceLogLevelEnum {
		return &v
	}).(DnsServiceLogLevelEnumPtrOutput)
}

func (o DnsServiceLogLevelEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DnsServiceLogLevelEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DnsServiceLogLevelEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DnsServiceLogLevelEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DnsServiceLogLevelEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DnsServiceLogLevelEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DnsServiceLogLevelEnumPtrOutput struct{ *pulumi.OutputState }

func (DnsServiceLogLevelEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsServiceLogLevelEnum)(nil)).Elem()
}

func (o DnsServiceLogLevelEnumPtrOutput) ToDnsServiceLogLevelEnumPtrOutput() DnsServiceLogLevelEnumPtrOutput {
	return o
}

func (o DnsServiceLogLevelEnumPtrOutput) ToDnsServiceLogLevelEnumPtrOutputWithContext(ctx context.Context) DnsServiceLogLevelEnumPtrOutput {
	return o
}

func (o DnsServiceLogLevelEnumPtrOutput) Elem() DnsServiceLogLevelEnumOutput {
	return o.ApplyT(func(v *DnsServiceLogLevelEnum) DnsServiceLogLevelEnum {
		if v != nil {
			return *v
		}
		var ret DnsServiceLogLevelEnum
		return ret
	}).(DnsServiceLogLevelEnumOutput)
}

func (o DnsServiceLogLevelEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DnsServiceLogLevelEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DnsServiceLogLevelEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DnsServiceLogLevelEnumInput interface {
	pulumi.Input

	ToDnsServiceLogLevelEnumOutput() DnsServiceLogLevelEnumOutput
	ToDnsServiceLogLevelEnumOutputWithContext(context.Context) DnsServiceLogLevelEnumOutput
}

var dnsServiceLogLevelEnumPtrType = reflect.TypeOf((**DnsServiceLogLevelEnum)(nil)).Elem()

type DnsServiceLogLevelEnumPtrInput interface {
	pulumi.Input

	ToDnsServiceLogLevelEnumPtrOutput() DnsServiceLogLevelEnumPtrOutput
	ToDnsServiceLogLevelEnumPtrOutputWithContext(context.Context) DnsServiceLogLevelEnumPtrOutput
}

type dnsServiceLogLevelEnumPtr string

func DnsServiceLogLevelEnumPtr(v string) DnsServiceLogLevelEnumPtrInput {
	return (*dnsServiceLogLevelEnumPtr)(&v)
}

func (*dnsServiceLogLevelEnumPtr) ElementType() reflect.Type {
	return dnsServiceLogLevelEnumPtrType
}

func (in *dnsServiceLogLevelEnumPtr) ToDnsServiceLogLevelEnumPtrOutput() DnsServiceLogLevelEnumPtrOutput {
	return pulumi.ToOutput(in).(DnsServiceLogLevelEnumPtrOutput)
}

func (in *dnsServiceLogLevelEnumPtr) ToDnsServiceLogLevelEnumPtrOutputWithContext(ctx context.Context) DnsServiceLogLevelEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DnsServiceLogLevelEnumPtrOutput)
}

type InternetEnum string

const (
	InternetEnumEnabled  = InternetEnum("Enabled")
	InternetEnumDisabled = InternetEnum("Disabled")
)

func (InternetEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*InternetEnum)(nil)).Elem()
}

func (e InternetEnum) ToInternetEnumOutput() InternetEnumOutput {
	return pulumi.ToOutput(e).(InternetEnumOutput)
}

func (e InternetEnum) ToInternetEnumOutputWithContext(ctx context.Context) InternetEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InternetEnumOutput)
}

func (e InternetEnum) ToInternetEnumPtrOutput() InternetEnumPtrOutput {
	return e.ToInternetEnumPtrOutputWithContext(context.Background())
}

func (e InternetEnum) ToInternetEnumPtrOutputWithContext(ctx context.Context) InternetEnumPtrOutput {
	return InternetEnum(e).ToInternetEnumOutputWithContext(ctx).ToInternetEnumPtrOutputWithContext(ctx)
}

func (e InternetEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InternetEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InternetEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InternetEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InternetEnumOutput struct{ *pulumi.OutputState }

func (InternetEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InternetEnum)(nil)).Elem()
}

func (o InternetEnumOutput) ToInternetEnumOutput() InternetEnumOutput {
	return o
}

func (o InternetEnumOutput) ToInternetEnumOutputWithContext(ctx context.Context) InternetEnumOutput {
	return o
}

func (o InternetEnumOutput) ToInternetEnumPtrOutput() InternetEnumPtrOutput {
	return o.ToInternetEnumPtrOutputWithContext(context.Background())
}

func (o InternetEnumOutput) ToInternetEnumPtrOutputWithContext(ctx context.Context) InternetEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InternetEnum) *InternetEnum {
		return &v
	}).(InternetEnumPtrOutput)
}

func (o InternetEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InternetEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InternetEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InternetEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InternetEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InternetEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InternetEnumPtrOutput struct{ *pulumi.OutputState }

func (InternetEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InternetEnum)(nil)).Elem()
}

func (o InternetEnumPtrOutput) ToInternetEnumPtrOutput() InternetEnumPtrOutput {
	return o
}

func (o InternetEnumPtrOutput) ToInternetEnumPtrOutputWithContext(ctx context.Context) InternetEnumPtrOutput {
	return o
}

func (o InternetEnumPtrOutput) Elem() InternetEnumOutput {
	return o.ApplyT(func(v *InternetEnum) InternetEnum {
		if v != nil {
			return *v
		}
		var ret InternetEnum
		return ret
	}).(InternetEnumOutput)
}

func (o InternetEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InternetEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InternetEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type InternetEnumInput interface {
	pulumi.Input

	ToInternetEnumOutput() InternetEnumOutput
	ToInternetEnumOutputWithContext(context.Context) InternetEnumOutput
}

var internetEnumPtrType = reflect.TypeOf((**InternetEnum)(nil)).Elem()

type InternetEnumPtrInput interface {
	pulumi.Input

	ToInternetEnumPtrOutput() InternetEnumPtrOutput
	ToInternetEnumPtrOutputWithContext(context.Context) InternetEnumPtrOutput
}

type internetEnumPtr string

func InternetEnumPtr(v string) InternetEnumPtrInput {
	return (*internetEnumPtr)(&v)
}

func (*internetEnumPtr) ElementType() reflect.Type {
	return internetEnumPtrType
}

func (in *internetEnumPtr) ToInternetEnumPtrOutput() InternetEnumPtrOutput {
	return pulumi.ToOutput(in).(InternetEnumPtrOutput)
}

func (in *internetEnumPtr) ToInternetEnumPtrOutputWithContext(ctx context.Context) InternetEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InternetEnumPtrOutput)
}

type PortMirroringDirectionEnum string

const (
	PortMirroringDirectionEnum_INGRESS_EGRESS_BIDIRECTIONAL = PortMirroringDirectionEnum("INGRESS, EGRESS, BIDIRECTIONAL")
)

func (PortMirroringDirectionEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*PortMirroringDirectionEnum)(nil)).Elem()
}

func (e PortMirroringDirectionEnum) ToPortMirroringDirectionEnumOutput() PortMirroringDirectionEnumOutput {
	return pulumi.ToOutput(e).(PortMirroringDirectionEnumOutput)
}

func (e PortMirroringDirectionEnum) ToPortMirroringDirectionEnumOutputWithContext(ctx context.Context) PortMirroringDirectionEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PortMirroringDirectionEnumOutput)
}

func (e PortMirroringDirectionEnum) ToPortMirroringDirectionEnumPtrOutput() PortMirroringDirectionEnumPtrOutput {
	return e.ToPortMirroringDirectionEnumPtrOutputWithContext(context.Background())
}

func (e PortMirroringDirectionEnum) ToPortMirroringDirectionEnumPtrOutputWithContext(ctx context.Context) PortMirroringDirectionEnumPtrOutput {
	return PortMirroringDirectionEnum(e).ToPortMirroringDirectionEnumOutputWithContext(ctx).ToPortMirroringDirectionEnumPtrOutputWithContext(ctx)
}

func (e PortMirroringDirectionEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PortMirroringDirectionEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PortMirroringDirectionEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PortMirroringDirectionEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PortMirroringDirectionEnumOutput struct{ *pulumi.OutputState }

func (PortMirroringDirectionEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PortMirroringDirectionEnum)(nil)).Elem()
}

func (o PortMirroringDirectionEnumOutput) ToPortMirroringDirectionEnumOutput() PortMirroringDirectionEnumOutput {
	return o
}

func (o PortMirroringDirectionEnumOutput) ToPortMirroringDirectionEnumOutputWithContext(ctx context.Context) PortMirroringDirectionEnumOutput {
	return o
}

func (o PortMirroringDirectionEnumOutput) ToPortMirroringDirectionEnumPtrOutput() PortMirroringDirectionEnumPtrOutput {
	return o.ToPortMirroringDirectionEnumPtrOutputWithContext(context.Background())
}

func (o PortMirroringDirectionEnumOutput) ToPortMirroringDirectionEnumPtrOutputWithContext(ctx context.Context) PortMirroringDirectionEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PortMirroringDirectionEnum) *PortMirroringDirectionEnum {
		return &v
	}).(PortMirroringDirectionEnumPtrOutput)
}

func (o PortMirroringDirectionEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PortMirroringDirectionEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PortMirroringDirectionEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PortMirroringDirectionEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PortMirroringDirectionEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PortMirroringDirectionEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PortMirroringDirectionEnumPtrOutput struct{ *pulumi.OutputState }

func (PortMirroringDirectionEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PortMirroringDirectionEnum)(nil)).Elem()
}

func (o PortMirroringDirectionEnumPtrOutput) ToPortMirroringDirectionEnumPtrOutput() PortMirroringDirectionEnumPtrOutput {
	return o
}

func (o PortMirroringDirectionEnumPtrOutput) ToPortMirroringDirectionEnumPtrOutputWithContext(ctx context.Context) PortMirroringDirectionEnumPtrOutput {
	return o
}

func (o PortMirroringDirectionEnumPtrOutput) Elem() PortMirroringDirectionEnumOutput {
	return o.ApplyT(func(v *PortMirroringDirectionEnum) PortMirroringDirectionEnum {
		if v != nil {
			return *v
		}
		var ret PortMirroringDirectionEnum
		return ret
	}).(PortMirroringDirectionEnumOutput)
}

func (o PortMirroringDirectionEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PortMirroringDirectionEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PortMirroringDirectionEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PortMirroringDirectionEnumInput interface {
	pulumi.Input

	ToPortMirroringDirectionEnumOutput() PortMirroringDirectionEnumOutput
	ToPortMirroringDirectionEnumOutputWithContext(context.Context) PortMirroringDirectionEnumOutput
}

var portMirroringDirectionEnumPtrType = reflect.TypeOf((**PortMirroringDirectionEnum)(nil)).Elem()

type PortMirroringDirectionEnumPtrInput interface {
	pulumi.Input

	ToPortMirroringDirectionEnumPtrOutput() PortMirroringDirectionEnumPtrOutput
	ToPortMirroringDirectionEnumPtrOutputWithContext(context.Context) PortMirroringDirectionEnumPtrOutput
}

type portMirroringDirectionEnumPtr string

func PortMirroringDirectionEnumPtr(v string) PortMirroringDirectionEnumPtrInput {
	return (*portMirroringDirectionEnumPtr)(&v)
}

func (*portMirroringDirectionEnumPtr) ElementType() reflect.Type {
	return portMirroringDirectionEnumPtrType
}

func (in *portMirroringDirectionEnumPtr) ToPortMirroringDirectionEnumPtrOutput() PortMirroringDirectionEnumPtrOutput {
	return pulumi.ToOutput(in).(PortMirroringDirectionEnumPtrOutput)
}

func (in *portMirroringDirectionEnumPtr) ToPortMirroringDirectionEnumPtrOutputWithContext(ctx context.Context) PortMirroringDirectionEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PortMirroringDirectionEnumPtrOutput)
}

type SslEnum string

const (
	SslEnumEnabled  = SslEnum("Enabled")
	SslEnumDisabled = SslEnum("Disabled")
)

func (SslEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*SslEnum)(nil)).Elem()
}

func (e SslEnum) ToSslEnumOutput() SslEnumOutput {
	return pulumi.ToOutput(e).(SslEnumOutput)
}

func (e SslEnum) ToSslEnumOutputWithContext(ctx context.Context) SslEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SslEnumOutput)
}

func (e SslEnum) ToSslEnumPtrOutput() SslEnumPtrOutput {
	return e.ToSslEnumPtrOutputWithContext(context.Background())
}

func (e SslEnum) ToSslEnumPtrOutputWithContext(ctx context.Context) SslEnumPtrOutput {
	return SslEnum(e).ToSslEnumOutputWithContext(ctx).ToSslEnumPtrOutputWithContext(ctx)
}

func (e SslEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SslEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SslEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SslEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SslEnumOutput struct{ *pulumi.OutputState }

func (SslEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SslEnum)(nil)).Elem()
}

func (o SslEnumOutput) ToSslEnumOutput() SslEnumOutput {
	return o
}

func (o SslEnumOutput) ToSslEnumOutputWithContext(ctx context.Context) SslEnumOutput {
	return o
}

func (o SslEnumOutput) ToSslEnumPtrOutput() SslEnumPtrOutput {
	return o.ToSslEnumPtrOutputWithContext(context.Background())
}

func (o SslEnumOutput) ToSslEnumPtrOutputWithContext(ctx context.Context) SslEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SslEnum) *SslEnum {
		return &v
	}).(SslEnumPtrOutput)
}

func (o SslEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SslEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SslEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SslEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SslEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SslEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SslEnumPtrOutput struct{ *pulumi.OutputState }

func (SslEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SslEnum)(nil)).Elem()
}

func (o SslEnumPtrOutput) ToSslEnumPtrOutput() SslEnumPtrOutput {
	return o
}

func (o SslEnumPtrOutput) ToSslEnumPtrOutputWithContext(ctx context.Context) SslEnumPtrOutput {
	return o
}

func (o SslEnumPtrOutput) Elem() SslEnumOutput {
	return o.ApplyT(func(v *SslEnum) SslEnum {
		if v != nil {
			return *v
		}
		var ret SslEnum
		return ret
	}).(SslEnumOutput)
}

func (o SslEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SslEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SslEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SslEnumInput interface {
	pulumi.Input

	ToSslEnumOutput() SslEnumOutput
	ToSslEnumOutputWithContext(context.Context) SslEnumOutput
}

var sslEnumPtrType = reflect.TypeOf((**SslEnum)(nil)).Elem()

type SslEnumPtrInput interface {
	pulumi.Input

	ToSslEnumPtrOutput() SslEnumPtrOutput
	ToSslEnumPtrOutputWithContext(context.Context) SslEnumPtrOutput
}

type sslEnumPtr string

func SslEnumPtr(v string) SslEnumPtrInput {
	return (*sslEnumPtr)(&v)
}

func (*sslEnumPtr) ElementType() reflect.Type {
	return sslEnumPtrType
}

func (in *sslEnumPtr) ToSslEnumPtrOutput() SslEnumPtrOutput {
	return pulumi.ToOutput(in).(SslEnumPtrOutput)
}

func (in *sslEnumPtr) ToSslEnumPtrOutputWithContext(ctx context.Context) SslEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SslEnumPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AddonTypeOutput{})
	pulumi.RegisterOutputType(AddonTypePtrOutput{})
	pulumi.RegisterOutputType(DhcpTypeEnumOutput{})
	pulumi.RegisterOutputType(DhcpTypeEnumPtrOutput{})
	pulumi.RegisterOutputType(DnsServiceLogLevelEnumOutput{})
	pulumi.RegisterOutputType(DnsServiceLogLevelEnumPtrOutput{})
	pulumi.RegisterOutputType(InternetEnumOutput{})
	pulumi.RegisterOutputType(InternetEnumPtrOutput{})
	pulumi.RegisterOutputType(PortMirroringDirectionEnumOutput{})
	pulumi.RegisterOutputType(PortMirroringDirectionEnumPtrOutput{})
	pulumi.RegisterOutputType(SslEnumOutput{})
	pulumi.RegisterOutputType(SslEnumPtrOutput{})
}
