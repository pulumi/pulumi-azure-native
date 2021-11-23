


package v20210404preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AutoTrackingConfiguration string

const (
	AutoTrackingConfigurationDisabled = AutoTrackingConfiguration("disabled")
	AutoTrackingConfigurationXBand    = AutoTrackingConfiguration("xBand")
	AutoTrackingConfigurationSBand    = AutoTrackingConfiguration("sBand")
)

func (AutoTrackingConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoTrackingConfiguration)(nil)).Elem()
}

func (e AutoTrackingConfiguration) ToAutoTrackingConfigurationOutput() AutoTrackingConfigurationOutput {
	return pulumi.ToOutput(e).(AutoTrackingConfigurationOutput)
}

func (e AutoTrackingConfiguration) ToAutoTrackingConfigurationOutputWithContext(ctx context.Context) AutoTrackingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AutoTrackingConfigurationOutput)
}

func (e AutoTrackingConfiguration) ToAutoTrackingConfigurationPtrOutput() AutoTrackingConfigurationPtrOutput {
	return e.ToAutoTrackingConfigurationPtrOutputWithContext(context.Background())
}

func (e AutoTrackingConfiguration) ToAutoTrackingConfigurationPtrOutputWithContext(ctx context.Context) AutoTrackingConfigurationPtrOutput {
	return AutoTrackingConfiguration(e).ToAutoTrackingConfigurationOutputWithContext(ctx).ToAutoTrackingConfigurationPtrOutputWithContext(ctx)
}

func (e AutoTrackingConfiguration) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoTrackingConfiguration) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoTrackingConfiguration) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutoTrackingConfiguration) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AutoTrackingConfigurationOutput struct{ *pulumi.OutputState }

func (AutoTrackingConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoTrackingConfiguration)(nil)).Elem()
}

func (o AutoTrackingConfigurationOutput) ToAutoTrackingConfigurationOutput() AutoTrackingConfigurationOutput {
	return o
}

func (o AutoTrackingConfigurationOutput) ToAutoTrackingConfigurationOutputWithContext(ctx context.Context) AutoTrackingConfigurationOutput {
	return o
}

func (o AutoTrackingConfigurationOutput) ToAutoTrackingConfigurationPtrOutput() AutoTrackingConfigurationPtrOutput {
	return o.ToAutoTrackingConfigurationPtrOutputWithContext(context.Background())
}

func (o AutoTrackingConfigurationOutput) ToAutoTrackingConfigurationPtrOutputWithContext(ctx context.Context) AutoTrackingConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoTrackingConfiguration) *AutoTrackingConfiguration {
		return &v
	}).(AutoTrackingConfigurationPtrOutput)
}

func (o AutoTrackingConfigurationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AutoTrackingConfigurationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoTrackingConfiguration) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AutoTrackingConfigurationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoTrackingConfigurationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoTrackingConfiguration) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AutoTrackingConfigurationPtrOutput struct{ *pulumi.OutputState }

func (AutoTrackingConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoTrackingConfiguration)(nil)).Elem()
}

func (o AutoTrackingConfigurationPtrOutput) ToAutoTrackingConfigurationPtrOutput() AutoTrackingConfigurationPtrOutput {
	return o
}

func (o AutoTrackingConfigurationPtrOutput) ToAutoTrackingConfigurationPtrOutputWithContext(ctx context.Context) AutoTrackingConfigurationPtrOutput {
	return o
}

func (o AutoTrackingConfigurationPtrOutput) Elem() AutoTrackingConfigurationOutput {
	return o.ApplyT(func(v *AutoTrackingConfiguration) AutoTrackingConfiguration {
		if v != nil {
			return *v
		}
		var ret AutoTrackingConfiguration
		return ret
	}).(AutoTrackingConfigurationOutput)
}

func (o AutoTrackingConfigurationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoTrackingConfigurationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AutoTrackingConfiguration) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AutoTrackingConfigurationInput interface {
	pulumi.Input

	ToAutoTrackingConfigurationOutput() AutoTrackingConfigurationOutput
	ToAutoTrackingConfigurationOutputWithContext(context.Context) AutoTrackingConfigurationOutput
}

var autoTrackingConfigurationPtrType = reflect.TypeOf((**AutoTrackingConfiguration)(nil)).Elem()

type AutoTrackingConfigurationPtrInput interface {
	pulumi.Input

	ToAutoTrackingConfigurationPtrOutput() AutoTrackingConfigurationPtrOutput
	ToAutoTrackingConfigurationPtrOutputWithContext(context.Context) AutoTrackingConfigurationPtrOutput
}

type autoTrackingConfigurationPtr string

func AutoTrackingConfigurationPtr(v string) AutoTrackingConfigurationPtrInput {
	return (*autoTrackingConfigurationPtr)(&v)
}

func (*autoTrackingConfigurationPtr) ElementType() reflect.Type {
	return autoTrackingConfigurationPtrType
}

func (in *autoTrackingConfigurationPtr) ToAutoTrackingConfigurationPtrOutput() AutoTrackingConfigurationPtrOutput {
	return pulumi.ToOutput(in).(AutoTrackingConfigurationPtrOutput)
}

func (in *autoTrackingConfigurationPtr) ToAutoTrackingConfigurationPtrOutputWithContext(ctx context.Context) AutoTrackingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AutoTrackingConfigurationPtrOutput)
}

type Direction string

const (
	DirectionUplink   = Direction("uplink")
	DirectionDownlink = Direction("downlink")
)

func (Direction) ElementType() reflect.Type {
	return reflect.TypeOf((*Direction)(nil)).Elem()
}

func (e Direction) ToDirectionOutput() DirectionOutput {
	return pulumi.ToOutput(e).(DirectionOutput)
}

func (e Direction) ToDirectionOutputWithContext(ctx context.Context) DirectionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DirectionOutput)
}

func (e Direction) ToDirectionPtrOutput() DirectionPtrOutput {
	return e.ToDirectionPtrOutputWithContext(context.Background())
}

func (e Direction) ToDirectionPtrOutputWithContext(ctx context.Context) DirectionPtrOutput {
	return Direction(e).ToDirectionOutputWithContext(ctx).ToDirectionPtrOutputWithContext(ctx)
}

func (e Direction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Direction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Direction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Direction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DirectionOutput struct{ *pulumi.OutputState }

func (DirectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Direction)(nil)).Elem()
}

func (o DirectionOutput) ToDirectionOutput() DirectionOutput {
	return o
}

func (o DirectionOutput) ToDirectionOutputWithContext(ctx context.Context) DirectionOutput {
	return o
}

func (o DirectionOutput) ToDirectionPtrOutput() DirectionPtrOutput {
	return o.ToDirectionPtrOutputWithContext(context.Background())
}

func (o DirectionOutput) ToDirectionPtrOutputWithContext(ctx context.Context) DirectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Direction) *Direction {
		return &v
	}).(DirectionPtrOutput)
}

func (o DirectionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DirectionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Direction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DirectionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DirectionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Direction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DirectionPtrOutput struct{ *pulumi.OutputState }

func (DirectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Direction)(nil)).Elem()
}

func (o DirectionPtrOutput) ToDirectionPtrOutput() DirectionPtrOutput {
	return o
}

func (o DirectionPtrOutput) ToDirectionPtrOutputWithContext(ctx context.Context) DirectionPtrOutput {
	return o
}

func (o DirectionPtrOutput) Elem() DirectionOutput {
	return o.ApplyT(func(v *Direction) Direction {
		if v != nil {
			return *v
		}
		var ret Direction
		return ret
	}).(DirectionOutput)
}

func (o DirectionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DirectionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Direction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DirectionInput interface {
	pulumi.Input

	ToDirectionOutput() DirectionOutput
	ToDirectionOutputWithContext(context.Context) DirectionOutput
}

var directionPtrType = reflect.TypeOf((**Direction)(nil)).Elem()

type DirectionPtrInput interface {
	pulumi.Input

	ToDirectionPtrOutput() DirectionPtrOutput
	ToDirectionPtrOutputWithContext(context.Context) DirectionPtrOutput
}

type directionPtr string

func DirectionPtr(v string) DirectionPtrInput {
	return (*directionPtr)(&v)
}

func (*directionPtr) ElementType() reflect.Type {
	return directionPtrType
}

func (in *directionPtr) ToDirectionPtrOutput() DirectionPtrOutput {
	return pulumi.ToOutput(in).(DirectionPtrOutput)
}

func (in *directionPtr) ToDirectionPtrOutputWithContext(ctx context.Context) DirectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DirectionPtrOutput)
}

type Polarization string

const (
	PolarizationRHCP             = Polarization("RHCP")
	PolarizationLHCP             = Polarization("LHCP")
	PolarizationDualRhcpLhcp     = Polarization("dualRhcpLhcp")
	PolarizationLinearVertical   = Polarization("linearVertical")
	PolarizationLinearHorizontal = Polarization("linearHorizontal")
)

func (Polarization) ElementType() reflect.Type {
	return reflect.TypeOf((*Polarization)(nil)).Elem()
}

func (e Polarization) ToPolarizationOutput() PolarizationOutput {
	return pulumi.ToOutput(e).(PolarizationOutput)
}

func (e Polarization) ToPolarizationOutputWithContext(ctx context.Context) PolarizationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PolarizationOutput)
}

func (e Polarization) ToPolarizationPtrOutput() PolarizationPtrOutput {
	return e.ToPolarizationPtrOutputWithContext(context.Background())
}

func (e Polarization) ToPolarizationPtrOutputWithContext(ctx context.Context) PolarizationPtrOutput {
	return Polarization(e).ToPolarizationOutputWithContext(ctx).ToPolarizationPtrOutputWithContext(ctx)
}

func (e Polarization) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Polarization) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Polarization) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Polarization) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PolarizationOutput struct{ *pulumi.OutputState }

func (PolarizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Polarization)(nil)).Elem()
}

func (o PolarizationOutput) ToPolarizationOutput() PolarizationOutput {
	return o
}

func (o PolarizationOutput) ToPolarizationOutputWithContext(ctx context.Context) PolarizationOutput {
	return o
}

func (o PolarizationOutput) ToPolarizationPtrOutput() PolarizationPtrOutput {
	return o.ToPolarizationPtrOutputWithContext(context.Background())
}

func (o PolarizationOutput) ToPolarizationPtrOutputWithContext(ctx context.Context) PolarizationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Polarization) *Polarization {
		return &v
	}).(PolarizationPtrOutput)
}

func (o PolarizationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PolarizationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Polarization) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PolarizationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolarizationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Polarization) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PolarizationPtrOutput struct{ *pulumi.OutputState }

func (PolarizationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Polarization)(nil)).Elem()
}

func (o PolarizationPtrOutput) ToPolarizationPtrOutput() PolarizationPtrOutput {
	return o
}

func (o PolarizationPtrOutput) ToPolarizationPtrOutputWithContext(ctx context.Context) PolarizationPtrOutput {
	return o
}

func (o PolarizationPtrOutput) Elem() PolarizationOutput {
	return o.ApplyT(func(v *Polarization) Polarization {
		if v != nil {
			return *v
		}
		var ret Polarization
		return ret
	}).(PolarizationOutput)
}

func (o PolarizationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolarizationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Polarization) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PolarizationInput interface {
	pulumi.Input

	ToPolarizationOutput() PolarizationOutput
	ToPolarizationOutputWithContext(context.Context) PolarizationOutput
}

var polarizationPtrType = reflect.TypeOf((**Polarization)(nil)).Elem()

type PolarizationPtrInput interface {
	pulumi.Input

	ToPolarizationPtrOutput() PolarizationPtrOutput
	ToPolarizationPtrOutputWithContext(context.Context) PolarizationPtrOutput
}

type polarizationPtr string

func PolarizationPtr(v string) PolarizationPtrInput {
	return (*polarizationPtr)(&v)
}

func (*polarizationPtr) ElementType() reflect.Type {
	return polarizationPtrType
}

func (in *polarizationPtr) ToPolarizationPtrOutput() PolarizationPtrOutput {
	return pulumi.ToOutput(in).(PolarizationPtrOutput)
}

func (in *polarizationPtr) ToPolarizationPtrOutputWithContext(ctx context.Context) PolarizationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PolarizationPtrOutput)
}

type Protocol string

const (
	ProtocolTCP = Protocol("TCP")
	ProtocolUDP = Protocol("UDP")
)

func (Protocol) ElementType() reflect.Type {
	return reflect.TypeOf((*Protocol)(nil)).Elem()
}

func (e Protocol) ToProtocolOutput() ProtocolOutput {
	return pulumi.ToOutput(e).(ProtocolOutput)
}

func (e Protocol) ToProtocolOutputWithContext(ctx context.Context) ProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProtocolOutput)
}

func (e Protocol) ToProtocolPtrOutput() ProtocolPtrOutput {
	return e.ToProtocolPtrOutputWithContext(context.Background())
}

func (e Protocol) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return Protocol(e).ToProtocolOutputWithContext(ctx).ToProtocolPtrOutputWithContext(ctx)
}

func (e Protocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Protocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Protocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Protocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProtocolOutput struct{ *pulumi.OutputState }

func (ProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Protocol)(nil)).Elem()
}

func (o ProtocolOutput) ToProtocolOutput() ProtocolOutput {
	return o
}

func (o ProtocolOutput) ToProtocolOutputWithContext(ctx context.Context) ProtocolOutput {
	return o
}

func (o ProtocolOutput) ToProtocolPtrOutput() ProtocolPtrOutput {
	return o.ToProtocolPtrOutputWithContext(context.Background())
}

func (o ProtocolOutput) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Protocol) *Protocol {
		return &v
	}).(ProtocolPtrOutput)
}

func (o ProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Protocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Protocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProtocolPtrOutput struct{ *pulumi.OutputState }

func (ProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Protocol)(nil)).Elem()
}

func (o ProtocolPtrOutput) ToProtocolPtrOutput() ProtocolPtrOutput {
	return o
}

func (o ProtocolPtrOutput) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return o
}

func (o ProtocolPtrOutput) Elem() ProtocolOutput {
	return o.ApplyT(func(v *Protocol) Protocol {
		if v != nil {
			return *v
		}
		var ret Protocol
		return ret
	}).(ProtocolOutput)
}

func (o ProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Protocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProtocolInput interface {
	pulumi.Input

	ToProtocolOutput() ProtocolOutput
	ToProtocolOutputWithContext(context.Context) ProtocolOutput
}

var protocolPtrType = reflect.TypeOf((**Protocol)(nil)).Elem()

type ProtocolPtrInput interface {
	pulumi.Input

	ToProtocolPtrOutput() ProtocolPtrOutput
	ToProtocolPtrOutputWithContext(context.Context) ProtocolPtrOutput
}

type protocolPtr string

func ProtocolPtr(v string) ProtocolPtrInput {
	return (*protocolPtr)(&v)
}

func (*protocolPtr) ElementType() reflect.Type {
	return protocolPtrType
}

func (in *protocolPtr) ToProtocolPtrOutput() ProtocolPtrOutput {
	return pulumi.ToOutput(in).(ProtocolPtrOutput)
}

func (in *protocolPtr) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProtocolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AutoTrackingConfigurationOutput{})
	pulumi.RegisterOutputType(AutoTrackingConfigurationPtrOutput{})
	pulumi.RegisterOutputType(DirectionOutput{})
	pulumi.RegisterOutputType(DirectionPtrOutput{})
	pulumi.RegisterOutputType(PolarizationOutput{})
	pulumi.RegisterOutputType(PolarizationPtrOutput{})
	pulumi.RegisterOutputType(ProtocolOutput{})
	pulumi.RegisterOutputType(ProtocolPtrOutput{})
}
