


package v20151101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DnsConfig struct {
	Fqdn         *string  `pulumi:"fqdn"`
	RelativeName *string  `pulumi:"relativeName"`
	Ttl          *float64 `pulumi:"ttl"`
}





type DnsConfigInput interface {
	pulumi.Input

	ToDnsConfigOutput() DnsConfigOutput
	ToDnsConfigOutputWithContext(context.Context) DnsConfigOutput
}

type DnsConfigArgs struct {
	Fqdn         pulumi.StringPtrInput  `pulumi:"fqdn"`
	RelativeName pulumi.StringPtrInput  `pulumi:"relativeName"`
	Ttl          pulumi.Float64PtrInput `pulumi:"ttl"`
}

func (DnsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsConfig)(nil)).Elem()
}

func (i DnsConfigArgs) ToDnsConfigOutput() DnsConfigOutput {
	return i.ToDnsConfigOutputWithContext(context.Background())
}

func (i DnsConfigArgs) ToDnsConfigOutputWithContext(ctx context.Context) DnsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsConfigOutput)
}

func (i DnsConfigArgs) ToDnsConfigPtrOutput() DnsConfigPtrOutput {
	return i.ToDnsConfigPtrOutputWithContext(context.Background())
}

func (i DnsConfigArgs) ToDnsConfigPtrOutputWithContext(ctx context.Context) DnsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsConfigOutput).ToDnsConfigPtrOutputWithContext(ctx)
}









type DnsConfigPtrInput interface {
	pulumi.Input

	ToDnsConfigPtrOutput() DnsConfigPtrOutput
	ToDnsConfigPtrOutputWithContext(context.Context) DnsConfigPtrOutput
}

type dnsConfigPtrType DnsConfigArgs

func DnsConfigPtr(v *DnsConfigArgs) DnsConfigPtrInput {
	return (*dnsConfigPtrType)(v)
}

func (*dnsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsConfig)(nil)).Elem()
}

func (i *dnsConfigPtrType) ToDnsConfigPtrOutput() DnsConfigPtrOutput {
	return i.ToDnsConfigPtrOutputWithContext(context.Background())
}

func (i *dnsConfigPtrType) ToDnsConfigPtrOutputWithContext(ctx context.Context) DnsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsConfigPtrOutput)
}

type DnsConfigOutput struct{ *pulumi.OutputState }

func (DnsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsConfig)(nil)).Elem()
}

func (o DnsConfigOutput) ToDnsConfigOutput() DnsConfigOutput {
	return o
}

func (o DnsConfigOutput) ToDnsConfigOutputWithContext(ctx context.Context) DnsConfigOutput {
	return o
}

func (o DnsConfigOutput) ToDnsConfigPtrOutput() DnsConfigPtrOutput {
	return o.ToDnsConfigPtrOutputWithContext(context.Background())
}

func (o DnsConfigOutput) ToDnsConfigPtrOutputWithContext(ctx context.Context) DnsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DnsConfig) *DnsConfig {
		return &v
	}).(DnsConfigPtrOutput)
}

func (o DnsConfigOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DnsConfig) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o DnsConfigOutput) RelativeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DnsConfig) *string { return v.RelativeName }).(pulumi.StringPtrOutput)
}

func (o DnsConfigOutput) Ttl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DnsConfig) *float64 { return v.Ttl }).(pulumi.Float64PtrOutput)
}

type DnsConfigPtrOutput struct{ *pulumi.OutputState }

func (DnsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsConfig)(nil)).Elem()
}

func (o DnsConfigPtrOutput) ToDnsConfigPtrOutput() DnsConfigPtrOutput {
	return o
}

func (o DnsConfigPtrOutput) ToDnsConfigPtrOutputWithContext(ctx context.Context) DnsConfigPtrOutput {
	return o
}

func (o DnsConfigPtrOutput) Elem() DnsConfigOutput {
	return o.ApplyT(func(v *DnsConfig) DnsConfig {
		if v != nil {
			return *v
		}
		var ret DnsConfig
		return ret
	}).(DnsConfigOutput)
}

func (o DnsConfigPtrOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsConfig) *string {
		if v == nil {
			return nil
		}
		return v.Fqdn
	}).(pulumi.StringPtrOutput)
}

func (o DnsConfigPtrOutput) RelativeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsConfig) *string {
		if v == nil {
			return nil
		}
		return v.RelativeName
	}).(pulumi.StringPtrOutput)
}

func (o DnsConfigPtrOutput) Ttl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DnsConfig) *float64 {
		if v == nil {
			return nil
		}
		return v.Ttl
	}).(pulumi.Float64PtrOutput)
}

type DnsConfigResponse struct {
	Fqdn         *string  `pulumi:"fqdn"`
	RelativeName *string  `pulumi:"relativeName"`
	Ttl          *float64 `pulumi:"ttl"`
}





type DnsConfigResponseInput interface {
	pulumi.Input

	ToDnsConfigResponseOutput() DnsConfigResponseOutput
	ToDnsConfigResponseOutputWithContext(context.Context) DnsConfigResponseOutput
}

type DnsConfigResponseArgs struct {
	Fqdn         pulumi.StringPtrInput  `pulumi:"fqdn"`
	RelativeName pulumi.StringPtrInput  `pulumi:"relativeName"`
	Ttl          pulumi.Float64PtrInput `pulumi:"ttl"`
}

func (DnsConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsConfigResponse)(nil)).Elem()
}

func (i DnsConfigResponseArgs) ToDnsConfigResponseOutput() DnsConfigResponseOutput {
	return i.ToDnsConfigResponseOutputWithContext(context.Background())
}

func (i DnsConfigResponseArgs) ToDnsConfigResponseOutputWithContext(ctx context.Context) DnsConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsConfigResponseOutput)
}

func (i DnsConfigResponseArgs) ToDnsConfigResponsePtrOutput() DnsConfigResponsePtrOutput {
	return i.ToDnsConfigResponsePtrOutputWithContext(context.Background())
}

func (i DnsConfigResponseArgs) ToDnsConfigResponsePtrOutputWithContext(ctx context.Context) DnsConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsConfigResponseOutput).ToDnsConfigResponsePtrOutputWithContext(ctx)
}









type DnsConfigResponsePtrInput interface {
	pulumi.Input

	ToDnsConfigResponsePtrOutput() DnsConfigResponsePtrOutput
	ToDnsConfigResponsePtrOutputWithContext(context.Context) DnsConfigResponsePtrOutput
}

type dnsConfigResponsePtrType DnsConfigResponseArgs

func DnsConfigResponsePtr(v *DnsConfigResponseArgs) DnsConfigResponsePtrInput {
	return (*dnsConfigResponsePtrType)(v)
}

func (*dnsConfigResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsConfigResponse)(nil)).Elem()
}

func (i *dnsConfigResponsePtrType) ToDnsConfigResponsePtrOutput() DnsConfigResponsePtrOutput {
	return i.ToDnsConfigResponsePtrOutputWithContext(context.Background())
}

func (i *dnsConfigResponsePtrType) ToDnsConfigResponsePtrOutputWithContext(ctx context.Context) DnsConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsConfigResponsePtrOutput)
}

type DnsConfigResponseOutput struct{ *pulumi.OutputState }

func (DnsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsConfigResponse)(nil)).Elem()
}

func (o DnsConfigResponseOutput) ToDnsConfigResponseOutput() DnsConfigResponseOutput {
	return o
}

func (o DnsConfigResponseOutput) ToDnsConfigResponseOutputWithContext(ctx context.Context) DnsConfigResponseOutput {
	return o
}

func (o DnsConfigResponseOutput) ToDnsConfigResponsePtrOutput() DnsConfigResponsePtrOutput {
	return o.ToDnsConfigResponsePtrOutputWithContext(context.Background())
}

func (o DnsConfigResponseOutput) ToDnsConfigResponsePtrOutputWithContext(ctx context.Context) DnsConfigResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DnsConfigResponse) *DnsConfigResponse {
		return &v
	}).(DnsConfigResponsePtrOutput)
}

func (o DnsConfigResponseOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DnsConfigResponse) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o DnsConfigResponseOutput) RelativeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DnsConfigResponse) *string { return v.RelativeName }).(pulumi.StringPtrOutput)
}

func (o DnsConfigResponseOutput) Ttl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DnsConfigResponse) *float64 { return v.Ttl }).(pulumi.Float64PtrOutput)
}

type DnsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (DnsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsConfigResponse)(nil)).Elem()
}

func (o DnsConfigResponsePtrOutput) ToDnsConfigResponsePtrOutput() DnsConfigResponsePtrOutput {
	return o
}

func (o DnsConfigResponsePtrOutput) ToDnsConfigResponsePtrOutputWithContext(ctx context.Context) DnsConfigResponsePtrOutput {
	return o
}

func (o DnsConfigResponsePtrOutput) Elem() DnsConfigResponseOutput {
	return o.ApplyT(func(v *DnsConfigResponse) DnsConfigResponse {
		if v != nil {
			return *v
		}
		var ret DnsConfigResponse
		return ret
	}).(DnsConfigResponseOutput)
}

func (o DnsConfigResponsePtrOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Fqdn
	}).(pulumi.StringPtrOutput)
}

func (o DnsConfigResponsePtrOutput) RelativeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.RelativeName
	}).(pulumi.StringPtrOutput)
}

func (o DnsConfigResponsePtrOutput) Ttl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DnsConfigResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Ttl
	}).(pulumi.Float64PtrOutput)
}

type EndpointType struct {
	EndpointLocation      *string  `pulumi:"endpointLocation"`
	EndpointMonitorStatus *string  `pulumi:"endpointMonitorStatus"`
	EndpointStatus        *string  `pulumi:"endpointStatus"`
	Id                    *string  `pulumi:"id"`
	MinChildEndpoints     *float64 `pulumi:"minChildEndpoints"`
	Name                  *string  `pulumi:"name"`
	Priority              *float64 `pulumi:"priority"`
	Target                *string  `pulumi:"target"`
	TargetResourceId      *string  `pulumi:"targetResourceId"`
	Type                  *string  `pulumi:"type"`
	Weight                *float64 `pulumi:"weight"`
}





type EndpointTypeInput interface {
	pulumi.Input

	ToEndpointTypeOutput() EndpointTypeOutput
	ToEndpointTypeOutputWithContext(context.Context) EndpointTypeOutput
}

type EndpointTypeArgs struct {
	EndpointLocation      pulumi.StringPtrInput  `pulumi:"endpointLocation"`
	EndpointMonitorStatus pulumi.StringPtrInput  `pulumi:"endpointMonitorStatus"`
	EndpointStatus        pulumi.StringPtrInput  `pulumi:"endpointStatus"`
	Id                    pulumi.StringPtrInput  `pulumi:"id"`
	MinChildEndpoints     pulumi.Float64PtrInput `pulumi:"minChildEndpoints"`
	Name                  pulumi.StringPtrInput  `pulumi:"name"`
	Priority              pulumi.Float64PtrInput `pulumi:"priority"`
	Target                pulumi.StringPtrInput  `pulumi:"target"`
	TargetResourceId      pulumi.StringPtrInput  `pulumi:"targetResourceId"`
	Type                  pulumi.StringPtrInput  `pulumi:"type"`
	Weight                pulumi.Float64PtrInput `pulumi:"weight"`
}

func (EndpointTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointType)(nil)).Elem()
}

func (i EndpointTypeArgs) ToEndpointTypeOutput() EndpointTypeOutput {
	return i.ToEndpointTypeOutputWithContext(context.Background())
}

func (i EndpointTypeArgs) ToEndpointTypeOutputWithContext(ctx context.Context) EndpointTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointTypeOutput)
}





type EndpointTypeArrayInput interface {
	pulumi.Input

	ToEndpointTypeArrayOutput() EndpointTypeArrayOutput
	ToEndpointTypeArrayOutputWithContext(context.Context) EndpointTypeArrayOutput
}

type EndpointTypeArray []EndpointTypeInput

func (EndpointTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointType)(nil)).Elem()
}

func (i EndpointTypeArray) ToEndpointTypeArrayOutput() EndpointTypeArrayOutput {
	return i.ToEndpointTypeArrayOutputWithContext(context.Background())
}

func (i EndpointTypeArray) ToEndpointTypeArrayOutputWithContext(ctx context.Context) EndpointTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointTypeArrayOutput)
}

type EndpointTypeOutput struct{ *pulumi.OutputState }

func (EndpointTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointType)(nil)).Elem()
}

func (o EndpointTypeOutput) ToEndpointTypeOutput() EndpointTypeOutput {
	return o
}

func (o EndpointTypeOutput) ToEndpointTypeOutputWithContext(ctx context.Context) EndpointTypeOutput {
	return o
}

func (o EndpointTypeOutput) EndpointLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.EndpointLocation }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) EndpointMonitorStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.EndpointMonitorStatus }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) EndpointStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.EndpointStatus }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) MinChildEndpoints() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EndpointType) *float64 { return v.MinChildEndpoints }).(pulumi.Float64PtrOutput)
}

func (o EndpointTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) Priority() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EndpointType) *float64 { return v.Priority }).(pulumi.Float64PtrOutput)
}

func (o EndpointTypeOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.TargetResourceId }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) Weight() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EndpointType) *float64 { return v.Weight }).(pulumi.Float64PtrOutput)
}

type EndpointTypeArrayOutput struct{ *pulumi.OutputState }

func (EndpointTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointType)(nil)).Elem()
}

func (o EndpointTypeArrayOutput) ToEndpointTypeArrayOutput() EndpointTypeArrayOutput {
	return o
}

func (o EndpointTypeArrayOutput) ToEndpointTypeArrayOutputWithContext(ctx context.Context) EndpointTypeArrayOutput {
	return o
}

func (o EndpointTypeArrayOutput) Index(i pulumi.IntInput) EndpointTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointType {
		return vs[0].([]EndpointType)[vs[1].(int)]
	}).(EndpointTypeOutput)
}

type EndpointResponse struct {
	EndpointLocation      *string  `pulumi:"endpointLocation"`
	EndpointMonitorStatus *string  `pulumi:"endpointMonitorStatus"`
	EndpointStatus        *string  `pulumi:"endpointStatus"`
	Id                    *string  `pulumi:"id"`
	MinChildEndpoints     *float64 `pulumi:"minChildEndpoints"`
	Name                  *string  `pulumi:"name"`
	Priority              *float64 `pulumi:"priority"`
	Target                *string  `pulumi:"target"`
	TargetResourceId      *string  `pulumi:"targetResourceId"`
	Type                  *string  `pulumi:"type"`
	Weight                *float64 `pulumi:"weight"`
}





type EndpointResponseInput interface {
	pulumi.Input

	ToEndpointResponseOutput() EndpointResponseOutput
	ToEndpointResponseOutputWithContext(context.Context) EndpointResponseOutput
}

type EndpointResponseArgs struct {
	EndpointLocation      pulumi.StringPtrInput  `pulumi:"endpointLocation"`
	EndpointMonitorStatus pulumi.StringPtrInput  `pulumi:"endpointMonitorStatus"`
	EndpointStatus        pulumi.StringPtrInput  `pulumi:"endpointStatus"`
	Id                    pulumi.StringPtrInput  `pulumi:"id"`
	MinChildEndpoints     pulumi.Float64PtrInput `pulumi:"minChildEndpoints"`
	Name                  pulumi.StringPtrInput  `pulumi:"name"`
	Priority              pulumi.Float64PtrInput `pulumi:"priority"`
	Target                pulumi.StringPtrInput  `pulumi:"target"`
	TargetResourceId      pulumi.StringPtrInput  `pulumi:"targetResourceId"`
	Type                  pulumi.StringPtrInput  `pulumi:"type"`
	Weight                pulumi.Float64PtrInput `pulumi:"weight"`
}

func (EndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointResponse)(nil)).Elem()
}

func (i EndpointResponseArgs) ToEndpointResponseOutput() EndpointResponseOutput {
	return i.ToEndpointResponseOutputWithContext(context.Background())
}

func (i EndpointResponseArgs) ToEndpointResponseOutputWithContext(ctx context.Context) EndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointResponseOutput)
}





type EndpointResponseArrayInput interface {
	pulumi.Input

	ToEndpointResponseArrayOutput() EndpointResponseArrayOutput
	ToEndpointResponseArrayOutputWithContext(context.Context) EndpointResponseArrayOutput
}

type EndpointResponseArray []EndpointResponseInput

func (EndpointResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointResponse)(nil)).Elem()
}

func (i EndpointResponseArray) ToEndpointResponseArrayOutput() EndpointResponseArrayOutput {
	return i.ToEndpointResponseArrayOutputWithContext(context.Background())
}

func (i EndpointResponseArray) ToEndpointResponseArrayOutputWithContext(ctx context.Context) EndpointResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointResponseArrayOutput)
}

type EndpointResponseOutput struct{ *pulumi.OutputState }

func (EndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointResponse)(nil)).Elem()
}

func (o EndpointResponseOutput) ToEndpointResponseOutput() EndpointResponseOutput {
	return o
}

func (o EndpointResponseOutput) ToEndpointResponseOutputWithContext(ctx context.Context) EndpointResponseOutput {
	return o
}

func (o EndpointResponseOutput) EndpointLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.EndpointLocation }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) EndpointMonitorStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.EndpointMonitorStatus }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) EndpointStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.EndpointStatus }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) MinChildEndpoints() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EndpointResponse) *float64 { return v.MinChildEndpoints }).(pulumi.Float64PtrOutput)
}

func (o EndpointResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) Priority() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EndpointResponse) *float64 { return v.Priority }).(pulumi.Float64PtrOutput)
}

func (o EndpointResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.TargetResourceId }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) Weight() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EndpointResponse) *float64 { return v.Weight }).(pulumi.Float64PtrOutput)
}

type EndpointResponseArrayOutput struct{ *pulumi.OutputState }

func (EndpointResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointResponse)(nil)).Elem()
}

func (o EndpointResponseArrayOutput) ToEndpointResponseArrayOutput() EndpointResponseArrayOutput {
	return o
}

func (o EndpointResponseArrayOutput) ToEndpointResponseArrayOutputWithContext(ctx context.Context) EndpointResponseArrayOutput {
	return o
}

func (o EndpointResponseArrayOutput) Index(i pulumi.IntInput) EndpointResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointResponse {
		return vs[0].([]EndpointResponse)[vs[1].(int)]
	}).(EndpointResponseOutput)
}

type MonitorConfig struct {
	Path                 *string  `pulumi:"path"`
	Port                 *float64 `pulumi:"port"`
	ProfileMonitorStatus *string  `pulumi:"profileMonitorStatus"`
	Protocol             *string  `pulumi:"protocol"`
}





type MonitorConfigInput interface {
	pulumi.Input

	ToMonitorConfigOutput() MonitorConfigOutput
	ToMonitorConfigOutputWithContext(context.Context) MonitorConfigOutput
}

type MonitorConfigArgs struct {
	Path                 pulumi.StringPtrInput  `pulumi:"path"`
	Port                 pulumi.Float64PtrInput `pulumi:"port"`
	ProfileMonitorStatus pulumi.StringPtrInput  `pulumi:"profileMonitorStatus"`
	Protocol             pulumi.StringPtrInput  `pulumi:"protocol"`
}

func (MonitorConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfig)(nil)).Elem()
}

func (i MonitorConfigArgs) ToMonitorConfigOutput() MonitorConfigOutput {
	return i.ToMonitorConfigOutputWithContext(context.Background())
}

func (i MonitorConfigArgs) ToMonitorConfigOutputWithContext(ctx context.Context) MonitorConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigOutput)
}

func (i MonitorConfigArgs) ToMonitorConfigPtrOutput() MonitorConfigPtrOutput {
	return i.ToMonitorConfigPtrOutputWithContext(context.Background())
}

func (i MonitorConfigArgs) ToMonitorConfigPtrOutputWithContext(ctx context.Context) MonitorConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigOutput).ToMonitorConfigPtrOutputWithContext(ctx)
}









type MonitorConfigPtrInput interface {
	pulumi.Input

	ToMonitorConfigPtrOutput() MonitorConfigPtrOutput
	ToMonitorConfigPtrOutputWithContext(context.Context) MonitorConfigPtrOutput
}

type monitorConfigPtrType MonitorConfigArgs

func MonitorConfigPtr(v *MonitorConfigArgs) MonitorConfigPtrInput {
	return (*monitorConfigPtrType)(v)
}

func (*monitorConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorConfig)(nil)).Elem()
}

func (i *monitorConfigPtrType) ToMonitorConfigPtrOutput() MonitorConfigPtrOutput {
	return i.ToMonitorConfigPtrOutputWithContext(context.Background())
}

func (i *monitorConfigPtrType) ToMonitorConfigPtrOutputWithContext(ctx context.Context) MonitorConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigPtrOutput)
}

type MonitorConfigOutput struct{ *pulumi.OutputState }

func (MonitorConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfig)(nil)).Elem()
}

func (o MonitorConfigOutput) ToMonitorConfigOutput() MonitorConfigOutput {
	return o
}

func (o MonitorConfigOutput) ToMonitorConfigOutputWithContext(ctx context.Context) MonitorConfigOutput {
	return o
}

func (o MonitorConfigOutput) ToMonitorConfigPtrOutput() MonitorConfigPtrOutput {
	return o.ToMonitorConfigPtrOutputWithContext(context.Background())
}

func (o MonitorConfigOutput) ToMonitorConfigPtrOutputWithContext(ctx context.Context) MonitorConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonitorConfig) *MonitorConfig {
		return &v
	}).(MonitorConfigPtrOutput)
}

func (o MonitorConfigOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfig) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o MonitorConfigOutput) Port() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MonitorConfig) *float64 { return v.Port }).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigOutput) ProfileMonitorStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfig) *string { return v.ProfileMonitorStatus }).(pulumi.StringPtrOutput)
}

func (o MonitorConfigOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfig) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

type MonitorConfigPtrOutput struct{ *pulumi.OutputState }

func (MonitorConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorConfig)(nil)).Elem()
}

func (o MonitorConfigPtrOutput) ToMonitorConfigPtrOutput() MonitorConfigPtrOutput {
	return o
}

func (o MonitorConfigPtrOutput) ToMonitorConfigPtrOutputWithContext(ctx context.Context) MonitorConfigPtrOutput {
	return o
}

func (o MonitorConfigPtrOutput) Elem() MonitorConfigOutput {
	return o.ApplyT(func(v *MonitorConfig) MonitorConfig {
		if v != nil {
			return *v
		}
		var ret MonitorConfig
		return ret
	}).(MonitorConfigOutput)
}

func (o MonitorConfigPtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorConfig) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o MonitorConfigPtrOutput) Port() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MonitorConfig) *float64 {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigPtrOutput) ProfileMonitorStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorConfig) *string {
		if v == nil {
			return nil
		}
		return v.ProfileMonitorStatus
	}).(pulumi.StringPtrOutput)
}

func (o MonitorConfigPtrOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorConfig) *string {
		if v == nil {
			return nil
		}
		return v.Protocol
	}).(pulumi.StringPtrOutput)
}

type MonitorConfigResponse struct {
	Path                 *string  `pulumi:"path"`
	Port                 *float64 `pulumi:"port"`
	ProfileMonitorStatus *string  `pulumi:"profileMonitorStatus"`
	Protocol             *string  `pulumi:"protocol"`
}





type MonitorConfigResponseInput interface {
	pulumi.Input

	ToMonitorConfigResponseOutput() MonitorConfigResponseOutput
	ToMonitorConfigResponseOutputWithContext(context.Context) MonitorConfigResponseOutput
}

type MonitorConfigResponseArgs struct {
	Path                 pulumi.StringPtrInput  `pulumi:"path"`
	Port                 pulumi.Float64PtrInput `pulumi:"port"`
	ProfileMonitorStatus pulumi.StringPtrInput  `pulumi:"profileMonitorStatus"`
	Protocol             pulumi.StringPtrInput  `pulumi:"protocol"`
}

func (MonitorConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfigResponse)(nil)).Elem()
}

func (i MonitorConfigResponseArgs) ToMonitorConfigResponseOutput() MonitorConfigResponseOutput {
	return i.ToMonitorConfigResponseOutputWithContext(context.Background())
}

func (i MonitorConfigResponseArgs) ToMonitorConfigResponseOutputWithContext(ctx context.Context) MonitorConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigResponseOutput)
}

func (i MonitorConfigResponseArgs) ToMonitorConfigResponsePtrOutput() MonitorConfigResponsePtrOutput {
	return i.ToMonitorConfigResponsePtrOutputWithContext(context.Background())
}

func (i MonitorConfigResponseArgs) ToMonitorConfigResponsePtrOutputWithContext(ctx context.Context) MonitorConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigResponseOutput).ToMonitorConfigResponsePtrOutputWithContext(ctx)
}









type MonitorConfigResponsePtrInput interface {
	pulumi.Input

	ToMonitorConfigResponsePtrOutput() MonitorConfigResponsePtrOutput
	ToMonitorConfigResponsePtrOutputWithContext(context.Context) MonitorConfigResponsePtrOutput
}

type monitorConfigResponsePtrType MonitorConfigResponseArgs

func MonitorConfigResponsePtr(v *MonitorConfigResponseArgs) MonitorConfigResponsePtrInput {
	return (*monitorConfigResponsePtrType)(v)
}

func (*monitorConfigResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorConfigResponse)(nil)).Elem()
}

func (i *monitorConfigResponsePtrType) ToMonitorConfigResponsePtrOutput() MonitorConfigResponsePtrOutput {
	return i.ToMonitorConfigResponsePtrOutputWithContext(context.Background())
}

func (i *monitorConfigResponsePtrType) ToMonitorConfigResponsePtrOutputWithContext(ctx context.Context) MonitorConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigResponsePtrOutput)
}

type MonitorConfigResponseOutput struct{ *pulumi.OutputState }

func (MonitorConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfigResponse)(nil)).Elem()
}

func (o MonitorConfigResponseOutput) ToMonitorConfigResponseOutput() MonitorConfigResponseOutput {
	return o
}

func (o MonitorConfigResponseOutput) ToMonitorConfigResponseOutputWithContext(ctx context.Context) MonitorConfigResponseOutput {
	return o
}

func (o MonitorConfigResponseOutput) ToMonitorConfigResponsePtrOutput() MonitorConfigResponsePtrOutput {
	return o.ToMonitorConfigResponsePtrOutputWithContext(context.Background())
}

func (o MonitorConfigResponseOutput) ToMonitorConfigResponsePtrOutputWithContext(ctx context.Context) MonitorConfigResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonitorConfigResponse) *MonitorConfigResponse {
		return &v
	}).(MonitorConfigResponsePtrOutput)
}

func (o MonitorConfigResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfigResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o MonitorConfigResponseOutput) Port() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MonitorConfigResponse) *float64 { return v.Port }).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigResponseOutput) ProfileMonitorStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfigResponse) *string { return v.ProfileMonitorStatus }).(pulumi.StringPtrOutput)
}

func (o MonitorConfigResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfigResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

type MonitorConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (MonitorConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorConfigResponse)(nil)).Elem()
}

func (o MonitorConfigResponsePtrOutput) ToMonitorConfigResponsePtrOutput() MonitorConfigResponsePtrOutput {
	return o
}

func (o MonitorConfigResponsePtrOutput) ToMonitorConfigResponsePtrOutputWithContext(ctx context.Context) MonitorConfigResponsePtrOutput {
	return o
}

func (o MonitorConfigResponsePtrOutput) Elem() MonitorConfigResponseOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) MonitorConfigResponse {
		if v != nil {
			return *v
		}
		var ret MonitorConfigResponse
		return ret
	}).(MonitorConfigResponseOutput)
}

func (o MonitorConfigResponsePtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o MonitorConfigResponsePtrOutput) Port() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigResponsePtrOutput) ProfileMonitorStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProfileMonitorStatus
	}).(pulumi.StringPtrOutput)
}

func (o MonitorConfigResponsePtrOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Protocol
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DnsConfigOutput{})
	pulumi.RegisterOutputType(DnsConfigPtrOutput{})
	pulumi.RegisterOutputType(DnsConfigResponseOutput{})
	pulumi.RegisterOutputType(DnsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(EndpointTypeOutput{})
	pulumi.RegisterOutputType(EndpointTypeArrayOutput{})
	pulumi.RegisterOutputType(EndpointResponseOutput{})
	pulumi.RegisterOutputType(EndpointResponseArrayOutput{})
	pulumi.RegisterOutputType(MonitorConfigOutput{})
	pulumi.RegisterOutputType(MonitorConfigPtrOutput{})
	pulumi.RegisterOutputType(MonitorConfigResponseOutput{})
	pulumi.RegisterOutputType(MonitorConfigResponsePtrOutput{})
}
