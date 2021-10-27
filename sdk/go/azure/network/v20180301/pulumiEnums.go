


package v20180301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EndpointMonitorStatus string

const (
	EndpointMonitorStatusCheckingEndpoint = EndpointMonitorStatus("CheckingEndpoint")
	EndpointMonitorStatusOnline           = EndpointMonitorStatus("Online")
	EndpointMonitorStatusDegraded         = EndpointMonitorStatus("Degraded")
	EndpointMonitorStatusDisabled         = EndpointMonitorStatus("Disabled")
	EndpointMonitorStatusInactive         = EndpointMonitorStatus("Inactive")
	EndpointMonitorStatusStopped          = EndpointMonitorStatus("Stopped")
)

func (EndpointMonitorStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointMonitorStatus)(nil)).Elem()
}

func (e EndpointMonitorStatus) ToEndpointMonitorStatusOutput() EndpointMonitorStatusOutput {
	return pulumi.ToOutput(e).(EndpointMonitorStatusOutput)
}

func (e EndpointMonitorStatus) ToEndpointMonitorStatusOutputWithContext(ctx context.Context) EndpointMonitorStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EndpointMonitorStatusOutput)
}

func (e EndpointMonitorStatus) ToEndpointMonitorStatusPtrOutput() EndpointMonitorStatusPtrOutput {
	return e.ToEndpointMonitorStatusPtrOutputWithContext(context.Background())
}

func (e EndpointMonitorStatus) ToEndpointMonitorStatusPtrOutputWithContext(ctx context.Context) EndpointMonitorStatusPtrOutput {
	return EndpointMonitorStatus(e).ToEndpointMonitorStatusOutputWithContext(ctx).ToEndpointMonitorStatusPtrOutputWithContext(ctx)
}

func (e EndpointMonitorStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EndpointMonitorStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EndpointMonitorStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EndpointMonitorStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EndpointMonitorStatusOutput struct{ *pulumi.OutputState }

func (EndpointMonitorStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointMonitorStatus)(nil)).Elem()
}

func (o EndpointMonitorStatusOutput) ToEndpointMonitorStatusOutput() EndpointMonitorStatusOutput {
	return o
}

func (o EndpointMonitorStatusOutput) ToEndpointMonitorStatusOutputWithContext(ctx context.Context) EndpointMonitorStatusOutput {
	return o
}

func (o EndpointMonitorStatusOutput) ToEndpointMonitorStatusPtrOutput() EndpointMonitorStatusPtrOutput {
	return o.ToEndpointMonitorStatusPtrOutputWithContext(context.Background())
}

func (o EndpointMonitorStatusOutput) ToEndpointMonitorStatusPtrOutputWithContext(ctx context.Context) EndpointMonitorStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointMonitorStatus) *EndpointMonitorStatus {
		return &v
	}).(EndpointMonitorStatusPtrOutput)
}

func (o EndpointMonitorStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EndpointMonitorStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EndpointMonitorStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EndpointMonitorStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EndpointMonitorStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EndpointMonitorStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EndpointMonitorStatusPtrOutput struct{ *pulumi.OutputState }

func (EndpointMonitorStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointMonitorStatus)(nil)).Elem()
}

func (o EndpointMonitorStatusPtrOutput) ToEndpointMonitorStatusPtrOutput() EndpointMonitorStatusPtrOutput {
	return o
}

func (o EndpointMonitorStatusPtrOutput) ToEndpointMonitorStatusPtrOutputWithContext(ctx context.Context) EndpointMonitorStatusPtrOutput {
	return o
}

func (o EndpointMonitorStatusPtrOutput) Elem() EndpointMonitorStatusOutput {
	return o.ApplyT(func(v *EndpointMonitorStatus) EndpointMonitorStatus {
		if v != nil {
			return *v
		}
		var ret EndpointMonitorStatus
		return ret
	}).(EndpointMonitorStatusOutput)
}

func (o EndpointMonitorStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EndpointMonitorStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EndpointMonitorStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EndpointMonitorStatusInput interface {
	pulumi.Input

	ToEndpointMonitorStatusOutput() EndpointMonitorStatusOutput
	ToEndpointMonitorStatusOutputWithContext(context.Context) EndpointMonitorStatusOutput
}

var endpointMonitorStatusPtrType = reflect.TypeOf((**EndpointMonitorStatus)(nil)).Elem()

type EndpointMonitorStatusPtrInput interface {
	pulumi.Input

	ToEndpointMonitorStatusPtrOutput() EndpointMonitorStatusPtrOutput
	ToEndpointMonitorStatusPtrOutputWithContext(context.Context) EndpointMonitorStatusPtrOutput
}

type endpointMonitorStatusPtr string

func EndpointMonitorStatusPtr(v string) EndpointMonitorStatusPtrInput {
	return (*endpointMonitorStatusPtr)(&v)
}

func (*endpointMonitorStatusPtr) ElementType() reflect.Type {
	return endpointMonitorStatusPtrType
}

func (in *endpointMonitorStatusPtr) ToEndpointMonitorStatusPtrOutput() EndpointMonitorStatusPtrOutput {
	return pulumi.ToOutput(in).(EndpointMonitorStatusPtrOutput)
}

func (in *endpointMonitorStatusPtr) ToEndpointMonitorStatusPtrOutputWithContext(ctx context.Context) EndpointMonitorStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EndpointMonitorStatusPtrOutput)
}

type EndpointStatus string

const (
	EndpointStatusEnabled  = EndpointStatus("Enabled")
	EndpointStatusDisabled = EndpointStatus("Disabled")
)

func (EndpointStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointStatus)(nil)).Elem()
}

func (e EndpointStatus) ToEndpointStatusOutput() EndpointStatusOutput {
	return pulumi.ToOutput(e).(EndpointStatusOutput)
}

func (e EndpointStatus) ToEndpointStatusOutputWithContext(ctx context.Context) EndpointStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EndpointStatusOutput)
}

func (e EndpointStatus) ToEndpointStatusPtrOutput() EndpointStatusPtrOutput {
	return e.ToEndpointStatusPtrOutputWithContext(context.Background())
}

func (e EndpointStatus) ToEndpointStatusPtrOutputWithContext(ctx context.Context) EndpointStatusPtrOutput {
	return EndpointStatus(e).ToEndpointStatusOutputWithContext(ctx).ToEndpointStatusPtrOutputWithContext(ctx)
}

func (e EndpointStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EndpointStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EndpointStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EndpointStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EndpointStatusOutput struct{ *pulumi.OutputState }

func (EndpointStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointStatus)(nil)).Elem()
}

func (o EndpointStatusOutput) ToEndpointStatusOutput() EndpointStatusOutput {
	return o
}

func (o EndpointStatusOutput) ToEndpointStatusOutputWithContext(ctx context.Context) EndpointStatusOutput {
	return o
}

func (o EndpointStatusOutput) ToEndpointStatusPtrOutput() EndpointStatusPtrOutput {
	return o.ToEndpointStatusPtrOutputWithContext(context.Background())
}

func (o EndpointStatusOutput) ToEndpointStatusPtrOutputWithContext(ctx context.Context) EndpointStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointStatus) *EndpointStatus {
		return &v
	}).(EndpointStatusPtrOutput)
}

func (o EndpointStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EndpointStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EndpointStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EndpointStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EndpointStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EndpointStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EndpointStatusPtrOutput struct{ *pulumi.OutputState }

func (EndpointStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointStatus)(nil)).Elem()
}

func (o EndpointStatusPtrOutput) ToEndpointStatusPtrOutput() EndpointStatusPtrOutput {
	return o
}

func (o EndpointStatusPtrOutput) ToEndpointStatusPtrOutputWithContext(ctx context.Context) EndpointStatusPtrOutput {
	return o
}

func (o EndpointStatusPtrOutput) Elem() EndpointStatusOutput {
	return o.ApplyT(func(v *EndpointStatus) EndpointStatus {
		if v != nil {
			return *v
		}
		var ret EndpointStatus
		return ret
	}).(EndpointStatusOutput)
}

func (o EndpointStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EndpointStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EndpointStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EndpointStatusInput interface {
	pulumi.Input

	ToEndpointStatusOutput() EndpointStatusOutput
	ToEndpointStatusOutputWithContext(context.Context) EndpointStatusOutput
}

var endpointStatusPtrType = reflect.TypeOf((**EndpointStatus)(nil)).Elem()

type EndpointStatusPtrInput interface {
	pulumi.Input

	ToEndpointStatusPtrOutput() EndpointStatusPtrOutput
	ToEndpointStatusPtrOutputWithContext(context.Context) EndpointStatusPtrOutput
}

type endpointStatusPtr string

func EndpointStatusPtr(v string) EndpointStatusPtrInput {
	return (*endpointStatusPtr)(&v)
}

func (*endpointStatusPtr) ElementType() reflect.Type {
	return endpointStatusPtrType
}

func (in *endpointStatusPtr) ToEndpointStatusPtrOutput() EndpointStatusPtrOutput {
	return pulumi.ToOutput(in).(EndpointStatusPtrOutput)
}

func (in *endpointStatusPtr) ToEndpointStatusPtrOutputWithContext(ctx context.Context) EndpointStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EndpointStatusPtrOutput)
}

type MonitorProtocol string

const (
	MonitorProtocolHTTP  = MonitorProtocol("HTTP")
	MonitorProtocolHTTPS = MonitorProtocol("HTTPS")
	MonitorProtocolTCP   = MonitorProtocol("TCP")
)

func (MonitorProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorProtocol)(nil)).Elem()
}

func (e MonitorProtocol) ToMonitorProtocolOutput() MonitorProtocolOutput {
	return pulumi.ToOutput(e).(MonitorProtocolOutput)
}

func (e MonitorProtocol) ToMonitorProtocolOutputWithContext(ctx context.Context) MonitorProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MonitorProtocolOutput)
}

func (e MonitorProtocol) ToMonitorProtocolPtrOutput() MonitorProtocolPtrOutput {
	return e.ToMonitorProtocolPtrOutputWithContext(context.Background())
}

func (e MonitorProtocol) ToMonitorProtocolPtrOutputWithContext(ctx context.Context) MonitorProtocolPtrOutput {
	return MonitorProtocol(e).ToMonitorProtocolOutputWithContext(ctx).ToMonitorProtocolPtrOutputWithContext(ctx)
}

func (e MonitorProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MonitorProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MonitorProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MonitorProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MonitorProtocolOutput struct{ *pulumi.OutputState }

func (MonitorProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorProtocol)(nil)).Elem()
}

func (o MonitorProtocolOutput) ToMonitorProtocolOutput() MonitorProtocolOutput {
	return o
}

func (o MonitorProtocolOutput) ToMonitorProtocolOutputWithContext(ctx context.Context) MonitorProtocolOutput {
	return o
}

func (o MonitorProtocolOutput) ToMonitorProtocolPtrOutput() MonitorProtocolPtrOutput {
	return o.ToMonitorProtocolPtrOutputWithContext(context.Background())
}

func (o MonitorProtocolOutput) ToMonitorProtocolPtrOutputWithContext(ctx context.Context) MonitorProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonitorProtocol) *MonitorProtocol {
		return &v
	}).(MonitorProtocolPtrOutput)
}

func (o MonitorProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MonitorProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MonitorProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MonitorProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MonitorProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MonitorProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MonitorProtocolPtrOutput struct{ *pulumi.OutputState }

func (MonitorProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorProtocol)(nil)).Elem()
}

func (o MonitorProtocolPtrOutput) ToMonitorProtocolPtrOutput() MonitorProtocolPtrOutput {
	return o
}

func (o MonitorProtocolPtrOutput) ToMonitorProtocolPtrOutputWithContext(ctx context.Context) MonitorProtocolPtrOutput {
	return o
}

func (o MonitorProtocolPtrOutput) Elem() MonitorProtocolOutput {
	return o.ApplyT(func(v *MonitorProtocol) MonitorProtocol {
		if v != nil {
			return *v
		}
		var ret MonitorProtocol
		return ret
	}).(MonitorProtocolOutput)
}

func (o MonitorProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MonitorProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MonitorProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MonitorProtocolInput interface {
	pulumi.Input

	ToMonitorProtocolOutput() MonitorProtocolOutput
	ToMonitorProtocolOutputWithContext(context.Context) MonitorProtocolOutput
}

var monitorProtocolPtrType = reflect.TypeOf((**MonitorProtocol)(nil)).Elem()

type MonitorProtocolPtrInput interface {
	pulumi.Input

	ToMonitorProtocolPtrOutput() MonitorProtocolPtrOutput
	ToMonitorProtocolPtrOutputWithContext(context.Context) MonitorProtocolPtrOutput
}

type monitorProtocolPtr string

func MonitorProtocolPtr(v string) MonitorProtocolPtrInput {
	return (*monitorProtocolPtr)(&v)
}

func (*monitorProtocolPtr) ElementType() reflect.Type {
	return monitorProtocolPtrType
}

func (in *monitorProtocolPtr) ToMonitorProtocolPtrOutput() MonitorProtocolPtrOutput {
	return pulumi.ToOutput(in).(MonitorProtocolPtrOutput)
}

func (in *monitorProtocolPtr) ToMonitorProtocolPtrOutputWithContext(ctx context.Context) MonitorProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MonitorProtocolPtrOutput)
}

type ProfileMonitorStatus string

const (
	ProfileMonitorStatusCheckingEndpoints = ProfileMonitorStatus("CheckingEndpoints")
	ProfileMonitorStatusOnline            = ProfileMonitorStatus("Online")
	ProfileMonitorStatusDegraded          = ProfileMonitorStatus("Degraded")
	ProfileMonitorStatusDisabled          = ProfileMonitorStatus("Disabled")
	ProfileMonitorStatusInactive          = ProfileMonitorStatus("Inactive")
)

func (ProfileMonitorStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileMonitorStatus)(nil)).Elem()
}

func (e ProfileMonitorStatus) ToProfileMonitorStatusOutput() ProfileMonitorStatusOutput {
	return pulumi.ToOutput(e).(ProfileMonitorStatusOutput)
}

func (e ProfileMonitorStatus) ToProfileMonitorStatusOutputWithContext(ctx context.Context) ProfileMonitorStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProfileMonitorStatusOutput)
}

func (e ProfileMonitorStatus) ToProfileMonitorStatusPtrOutput() ProfileMonitorStatusPtrOutput {
	return e.ToProfileMonitorStatusPtrOutputWithContext(context.Background())
}

func (e ProfileMonitorStatus) ToProfileMonitorStatusPtrOutputWithContext(ctx context.Context) ProfileMonitorStatusPtrOutput {
	return ProfileMonitorStatus(e).ToProfileMonitorStatusOutputWithContext(ctx).ToProfileMonitorStatusPtrOutputWithContext(ctx)
}

func (e ProfileMonitorStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProfileMonitorStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProfileMonitorStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProfileMonitorStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProfileMonitorStatusOutput struct{ *pulumi.OutputState }

func (ProfileMonitorStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileMonitorStatus)(nil)).Elem()
}

func (o ProfileMonitorStatusOutput) ToProfileMonitorStatusOutput() ProfileMonitorStatusOutput {
	return o
}

func (o ProfileMonitorStatusOutput) ToProfileMonitorStatusOutputWithContext(ctx context.Context) ProfileMonitorStatusOutput {
	return o
}

func (o ProfileMonitorStatusOutput) ToProfileMonitorStatusPtrOutput() ProfileMonitorStatusPtrOutput {
	return o.ToProfileMonitorStatusPtrOutputWithContext(context.Background())
}

func (o ProfileMonitorStatusOutput) ToProfileMonitorStatusPtrOutputWithContext(ctx context.Context) ProfileMonitorStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProfileMonitorStatus) *ProfileMonitorStatus {
		return &v
	}).(ProfileMonitorStatusPtrOutput)
}

func (o ProfileMonitorStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProfileMonitorStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProfileMonitorStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProfileMonitorStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProfileMonitorStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProfileMonitorStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProfileMonitorStatusPtrOutput struct{ *pulumi.OutputState }

func (ProfileMonitorStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfileMonitorStatus)(nil)).Elem()
}

func (o ProfileMonitorStatusPtrOutput) ToProfileMonitorStatusPtrOutput() ProfileMonitorStatusPtrOutput {
	return o
}

func (o ProfileMonitorStatusPtrOutput) ToProfileMonitorStatusPtrOutputWithContext(ctx context.Context) ProfileMonitorStatusPtrOutput {
	return o
}

func (o ProfileMonitorStatusPtrOutput) Elem() ProfileMonitorStatusOutput {
	return o.ApplyT(func(v *ProfileMonitorStatus) ProfileMonitorStatus {
		if v != nil {
			return *v
		}
		var ret ProfileMonitorStatus
		return ret
	}).(ProfileMonitorStatusOutput)
}

func (o ProfileMonitorStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProfileMonitorStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProfileMonitorStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProfileMonitorStatusInput interface {
	pulumi.Input

	ToProfileMonitorStatusOutput() ProfileMonitorStatusOutput
	ToProfileMonitorStatusOutputWithContext(context.Context) ProfileMonitorStatusOutput
}

var profileMonitorStatusPtrType = reflect.TypeOf((**ProfileMonitorStatus)(nil)).Elem()

type ProfileMonitorStatusPtrInput interface {
	pulumi.Input

	ToProfileMonitorStatusPtrOutput() ProfileMonitorStatusPtrOutput
	ToProfileMonitorStatusPtrOutputWithContext(context.Context) ProfileMonitorStatusPtrOutput
}

type profileMonitorStatusPtr string

func ProfileMonitorStatusPtr(v string) ProfileMonitorStatusPtrInput {
	return (*profileMonitorStatusPtr)(&v)
}

func (*profileMonitorStatusPtr) ElementType() reflect.Type {
	return profileMonitorStatusPtrType
}

func (in *profileMonitorStatusPtr) ToProfileMonitorStatusPtrOutput() ProfileMonitorStatusPtrOutput {
	return pulumi.ToOutput(in).(ProfileMonitorStatusPtrOutput)
}

func (in *profileMonitorStatusPtr) ToProfileMonitorStatusPtrOutputWithContext(ctx context.Context) ProfileMonitorStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProfileMonitorStatusPtrOutput)
}

type ProfileStatus string

const (
	ProfileStatusEnabled  = ProfileStatus("Enabled")
	ProfileStatusDisabled = ProfileStatus("Disabled")
)

func (ProfileStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileStatus)(nil)).Elem()
}

func (e ProfileStatus) ToProfileStatusOutput() ProfileStatusOutput {
	return pulumi.ToOutput(e).(ProfileStatusOutput)
}

func (e ProfileStatus) ToProfileStatusOutputWithContext(ctx context.Context) ProfileStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProfileStatusOutput)
}

func (e ProfileStatus) ToProfileStatusPtrOutput() ProfileStatusPtrOutput {
	return e.ToProfileStatusPtrOutputWithContext(context.Background())
}

func (e ProfileStatus) ToProfileStatusPtrOutputWithContext(ctx context.Context) ProfileStatusPtrOutput {
	return ProfileStatus(e).ToProfileStatusOutputWithContext(ctx).ToProfileStatusPtrOutputWithContext(ctx)
}

func (e ProfileStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProfileStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProfileStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProfileStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProfileStatusOutput struct{ *pulumi.OutputState }

func (ProfileStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileStatus)(nil)).Elem()
}

func (o ProfileStatusOutput) ToProfileStatusOutput() ProfileStatusOutput {
	return o
}

func (o ProfileStatusOutput) ToProfileStatusOutputWithContext(ctx context.Context) ProfileStatusOutput {
	return o
}

func (o ProfileStatusOutput) ToProfileStatusPtrOutput() ProfileStatusPtrOutput {
	return o.ToProfileStatusPtrOutputWithContext(context.Background())
}

func (o ProfileStatusOutput) ToProfileStatusPtrOutputWithContext(ctx context.Context) ProfileStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProfileStatus) *ProfileStatus {
		return &v
	}).(ProfileStatusPtrOutput)
}

func (o ProfileStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProfileStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProfileStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProfileStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProfileStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProfileStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProfileStatusPtrOutput struct{ *pulumi.OutputState }

func (ProfileStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfileStatus)(nil)).Elem()
}

func (o ProfileStatusPtrOutput) ToProfileStatusPtrOutput() ProfileStatusPtrOutput {
	return o
}

func (o ProfileStatusPtrOutput) ToProfileStatusPtrOutputWithContext(ctx context.Context) ProfileStatusPtrOutput {
	return o
}

func (o ProfileStatusPtrOutput) Elem() ProfileStatusOutput {
	return o.ApplyT(func(v *ProfileStatus) ProfileStatus {
		if v != nil {
			return *v
		}
		var ret ProfileStatus
		return ret
	}).(ProfileStatusOutput)
}

func (o ProfileStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProfileStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProfileStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProfileStatusInput interface {
	pulumi.Input

	ToProfileStatusOutput() ProfileStatusOutput
	ToProfileStatusOutputWithContext(context.Context) ProfileStatusOutput
}

var profileStatusPtrType = reflect.TypeOf((**ProfileStatus)(nil)).Elem()

type ProfileStatusPtrInput interface {
	pulumi.Input

	ToProfileStatusPtrOutput() ProfileStatusPtrOutput
	ToProfileStatusPtrOutputWithContext(context.Context) ProfileStatusPtrOutput
}

type profileStatusPtr string

func ProfileStatusPtr(v string) ProfileStatusPtrInput {
	return (*profileStatusPtr)(&v)
}

func (*profileStatusPtr) ElementType() reflect.Type {
	return profileStatusPtrType
}

func (in *profileStatusPtr) ToProfileStatusPtrOutput() ProfileStatusPtrOutput {
	return pulumi.ToOutput(in).(ProfileStatusPtrOutput)
}

func (in *profileStatusPtr) ToProfileStatusPtrOutputWithContext(ctx context.Context) ProfileStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProfileStatusPtrOutput)
}

type TrafficRoutingMethod string

const (
	TrafficRoutingMethodPerformance = TrafficRoutingMethod("Performance")
	TrafficRoutingMethodPriority    = TrafficRoutingMethod("Priority")
	TrafficRoutingMethodWeighted    = TrafficRoutingMethod("Weighted")
	TrafficRoutingMethodGeographic  = TrafficRoutingMethod("Geographic")
)

func (TrafficRoutingMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficRoutingMethod)(nil)).Elem()
}

func (e TrafficRoutingMethod) ToTrafficRoutingMethodOutput() TrafficRoutingMethodOutput {
	return pulumi.ToOutput(e).(TrafficRoutingMethodOutput)
}

func (e TrafficRoutingMethod) ToTrafficRoutingMethodOutputWithContext(ctx context.Context) TrafficRoutingMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TrafficRoutingMethodOutput)
}

func (e TrafficRoutingMethod) ToTrafficRoutingMethodPtrOutput() TrafficRoutingMethodPtrOutput {
	return e.ToTrafficRoutingMethodPtrOutputWithContext(context.Background())
}

func (e TrafficRoutingMethod) ToTrafficRoutingMethodPtrOutputWithContext(ctx context.Context) TrafficRoutingMethodPtrOutput {
	return TrafficRoutingMethod(e).ToTrafficRoutingMethodOutputWithContext(ctx).ToTrafficRoutingMethodPtrOutputWithContext(ctx)
}

func (e TrafficRoutingMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TrafficRoutingMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TrafficRoutingMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TrafficRoutingMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TrafficRoutingMethodOutput struct{ *pulumi.OutputState }

func (TrafficRoutingMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficRoutingMethod)(nil)).Elem()
}

func (o TrafficRoutingMethodOutput) ToTrafficRoutingMethodOutput() TrafficRoutingMethodOutput {
	return o
}

func (o TrafficRoutingMethodOutput) ToTrafficRoutingMethodOutputWithContext(ctx context.Context) TrafficRoutingMethodOutput {
	return o
}

func (o TrafficRoutingMethodOutput) ToTrafficRoutingMethodPtrOutput() TrafficRoutingMethodPtrOutput {
	return o.ToTrafficRoutingMethodPtrOutputWithContext(context.Background())
}

func (o TrafficRoutingMethodOutput) ToTrafficRoutingMethodPtrOutputWithContext(ctx context.Context) TrafficRoutingMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TrafficRoutingMethod) *TrafficRoutingMethod {
		return &v
	}).(TrafficRoutingMethodPtrOutput)
}

func (o TrafficRoutingMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TrafficRoutingMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TrafficRoutingMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TrafficRoutingMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TrafficRoutingMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TrafficRoutingMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TrafficRoutingMethodPtrOutput struct{ *pulumi.OutputState }

func (TrafficRoutingMethodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficRoutingMethod)(nil)).Elem()
}

func (o TrafficRoutingMethodPtrOutput) ToTrafficRoutingMethodPtrOutput() TrafficRoutingMethodPtrOutput {
	return o
}

func (o TrafficRoutingMethodPtrOutput) ToTrafficRoutingMethodPtrOutputWithContext(ctx context.Context) TrafficRoutingMethodPtrOutput {
	return o
}

func (o TrafficRoutingMethodPtrOutput) Elem() TrafficRoutingMethodOutput {
	return o.ApplyT(func(v *TrafficRoutingMethod) TrafficRoutingMethod {
		if v != nil {
			return *v
		}
		var ret TrafficRoutingMethod
		return ret
	}).(TrafficRoutingMethodOutput)
}

func (o TrafficRoutingMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TrafficRoutingMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TrafficRoutingMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TrafficRoutingMethodInput interface {
	pulumi.Input

	ToTrafficRoutingMethodOutput() TrafficRoutingMethodOutput
	ToTrafficRoutingMethodOutputWithContext(context.Context) TrafficRoutingMethodOutput
}

var trafficRoutingMethodPtrType = reflect.TypeOf((**TrafficRoutingMethod)(nil)).Elem()

type TrafficRoutingMethodPtrInput interface {
	pulumi.Input

	ToTrafficRoutingMethodPtrOutput() TrafficRoutingMethodPtrOutput
	ToTrafficRoutingMethodPtrOutputWithContext(context.Context) TrafficRoutingMethodPtrOutput
}

type trafficRoutingMethodPtr string

func TrafficRoutingMethodPtr(v string) TrafficRoutingMethodPtrInput {
	return (*trafficRoutingMethodPtr)(&v)
}

func (*trafficRoutingMethodPtr) ElementType() reflect.Type {
	return trafficRoutingMethodPtrType
}

func (in *trafficRoutingMethodPtr) ToTrafficRoutingMethodPtrOutput() TrafficRoutingMethodPtrOutput {
	return pulumi.ToOutput(in).(TrafficRoutingMethodPtrOutput)
}

func (in *trafficRoutingMethodPtr) ToTrafficRoutingMethodPtrOutputWithContext(ctx context.Context) TrafficRoutingMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TrafficRoutingMethodPtrOutput)
}

type TrafficViewEnrollmentStatus string

const (
	TrafficViewEnrollmentStatusEnabled  = TrafficViewEnrollmentStatus("Enabled")
	TrafficViewEnrollmentStatusDisabled = TrafficViewEnrollmentStatus("Disabled")
)

func (TrafficViewEnrollmentStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficViewEnrollmentStatus)(nil)).Elem()
}

func (e TrafficViewEnrollmentStatus) ToTrafficViewEnrollmentStatusOutput() TrafficViewEnrollmentStatusOutput {
	return pulumi.ToOutput(e).(TrafficViewEnrollmentStatusOutput)
}

func (e TrafficViewEnrollmentStatus) ToTrafficViewEnrollmentStatusOutputWithContext(ctx context.Context) TrafficViewEnrollmentStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TrafficViewEnrollmentStatusOutput)
}

func (e TrafficViewEnrollmentStatus) ToTrafficViewEnrollmentStatusPtrOutput() TrafficViewEnrollmentStatusPtrOutput {
	return e.ToTrafficViewEnrollmentStatusPtrOutputWithContext(context.Background())
}

func (e TrafficViewEnrollmentStatus) ToTrafficViewEnrollmentStatusPtrOutputWithContext(ctx context.Context) TrafficViewEnrollmentStatusPtrOutput {
	return TrafficViewEnrollmentStatus(e).ToTrafficViewEnrollmentStatusOutputWithContext(ctx).ToTrafficViewEnrollmentStatusPtrOutputWithContext(ctx)
}

func (e TrafficViewEnrollmentStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TrafficViewEnrollmentStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TrafficViewEnrollmentStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TrafficViewEnrollmentStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TrafficViewEnrollmentStatusOutput struct{ *pulumi.OutputState }

func (TrafficViewEnrollmentStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficViewEnrollmentStatus)(nil)).Elem()
}

func (o TrafficViewEnrollmentStatusOutput) ToTrafficViewEnrollmentStatusOutput() TrafficViewEnrollmentStatusOutput {
	return o
}

func (o TrafficViewEnrollmentStatusOutput) ToTrafficViewEnrollmentStatusOutputWithContext(ctx context.Context) TrafficViewEnrollmentStatusOutput {
	return o
}

func (o TrafficViewEnrollmentStatusOutput) ToTrafficViewEnrollmentStatusPtrOutput() TrafficViewEnrollmentStatusPtrOutput {
	return o.ToTrafficViewEnrollmentStatusPtrOutputWithContext(context.Background())
}

func (o TrafficViewEnrollmentStatusOutput) ToTrafficViewEnrollmentStatusPtrOutputWithContext(ctx context.Context) TrafficViewEnrollmentStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TrafficViewEnrollmentStatus) *TrafficViewEnrollmentStatus {
		return &v
	}).(TrafficViewEnrollmentStatusPtrOutput)
}

func (o TrafficViewEnrollmentStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TrafficViewEnrollmentStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TrafficViewEnrollmentStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TrafficViewEnrollmentStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TrafficViewEnrollmentStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TrafficViewEnrollmentStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TrafficViewEnrollmentStatusPtrOutput struct{ *pulumi.OutputState }

func (TrafficViewEnrollmentStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficViewEnrollmentStatus)(nil)).Elem()
}

func (o TrafficViewEnrollmentStatusPtrOutput) ToTrafficViewEnrollmentStatusPtrOutput() TrafficViewEnrollmentStatusPtrOutput {
	return o
}

func (o TrafficViewEnrollmentStatusPtrOutput) ToTrafficViewEnrollmentStatusPtrOutputWithContext(ctx context.Context) TrafficViewEnrollmentStatusPtrOutput {
	return o
}

func (o TrafficViewEnrollmentStatusPtrOutput) Elem() TrafficViewEnrollmentStatusOutput {
	return o.ApplyT(func(v *TrafficViewEnrollmentStatus) TrafficViewEnrollmentStatus {
		if v != nil {
			return *v
		}
		var ret TrafficViewEnrollmentStatus
		return ret
	}).(TrafficViewEnrollmentStatusOutput)
}

func (o TrafficViewEnrollmentStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TrafficViewEnrollmentStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TrafficViewEnrollmentStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TrafficViewEnrollmentStatusInput interface {
	pulumi.Input

	ToTrafficViewEnrollmentStatusOutput() TrafficViewEnrollmentStatusOutput
	ToTrafficViewEnrollmentStatusOutputWithContext(context.Context) TrafficViewEnrollmentStatusOutput
}

var trafficViewEnrollmentStatusPtrType = reflect.TypeOf((**TrafficViewEnrollmentStatus)(nil)).Elem()

type TrafficViewEnrollmentStatusPtrInput interface {
	pulumi.Input

	ToTrafficViewEnrollmentStatusPtrOutput() TrafficViewEnrollmentStatusPtrOutput
	ToTrafficViewEnrollmentStatusPtrOutputWithContext(context.Context) TrafficViewEnrollmentStatusPtrOutput
}

type trafficViewEnrollmentStatusPtr string

func TrafficViewEnrollmentStatusPtr(v string) TrafficViewEnrollmentStatusPtrInput {
	return (*trafficViewEnrollmentStatusPtr)(&v)
}

func (*trafficViewEnrollmentStatusPtr) ElementType() reflect.Type {
	return trafficViewEnrollmentStatusPtrType
}

func (in *trafficViewEnrollmentStatusPtr) ToTrafficViewEnrollmentStatusPtrOutput() TrafficViewEnrollmentStatusPtrOutput {
	return pulumi.ToOutput(in).(TrafficViewEnrollmentStatusPtrOutput)
}

func (in *trafficViewEnrollmentStatusPtr) ToTrafficViewEnrollmentStatusPtrOutputWithContext(ctx context.Context) TrafficViewEnrollmentStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TrafficViewEnrollmentStatusPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointMonitorStatusInput)(nil)).Elem(), EndpointMonitorStatus("CheckingEndpoint"))
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointMonitorStatusPtrInput)(nil)).Elem(), EndpointMonitorStatus("CheckingEndpoint"))
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointStatusInput)(nil)).Elem(), EndpointStatus("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointStatusPtrInput)(nil)).Elem(), EndpointStatus("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorProtocolInput)(nil)).Elem(), MonitorProtocol("HTTP"))
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorProtocolPtrInput)(nil)).Elem(), MonitorProtocol("HTTP"))
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileMonitorStatusInput)(nil)).Elem(), ProfileMonitorStatus("CheckingEndpoints"))
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileMonitorStatusPtrInput)(nil)).Elem(), ProfileMonitorStatus("CheckingEndpoints"))
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileStatusInput)(nil)).Elem(), ProfileStatus("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*ProfileStatusPtrInput)(nil)).Elem(), ProfileStatus("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficRoutingMethodInput)(nil)).Elem(), TrafficRoutingMethod("Performance"))
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficRoutingMethodPtrInput)(nil)).Elem(), TrafficRoutingMethod("Performance"))
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficViewEnrollmentStatusInput)(nil)).Elem(), TrafficViewEnrollmentStatus("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficViewEnrollmentStatusPtrInput)(nil)).Elem(), TrafficViewEnrollmentStatus("Enabled"))
	pulumi.RegisterOutputType(EndpointMonitorStatusOutput{})
	pulumi.RegisterOutputType(EndpointMonitorStatusPtrOutput{})
	pulumi.RegisterOutputType(EndpointStatusOutput{})
	pulumi.RegisterOutputType(EndpointStatusPtrOutput{})
	pulumi.RegisterOutputType(MonitorProtocolOutput{})
	pulumi.RegisterOutputType(MonitorProtocolPtrOutput{})
	pulumi.RegisterOutputType(ProfileMonitorStatusOutput{})
	pulumi.RegisterOutputType(ProfileMonitorStatusPtrOutput{})
	pulumi.RegisterOutputType(ProfileStatusOutput{})
	pulumi.RegisterOutputType(ProfileStatusPtrOutput{})
	pulumi.RegisterOutputType(TrafficRoutingMethodOutput{})
	pulumi.RegisterOutputType(TrafficRoutingMethodPtrOutput{})
	pulumi.RegisterOutputType(TrafficViewEnrollmentStatusOutput{})
	pulumi.RegisterOutputType(TrafficViewEnrollmentStatusPtrOutput{})
}
