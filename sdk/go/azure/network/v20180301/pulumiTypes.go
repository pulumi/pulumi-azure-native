


package v20180301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DnsConfig struct {
	RelativeName *string  `pulumi:"relativeName"`
	Ttl          *float64 `pulumi:"ttl"`
}





type DnsConfigInput interface {
	pulumi.Input

	ToDnsConfigOutput() DnsConfigOutput
	ToDnsConfigOutputWithContext(context.Context) DnsConfigOutput
}

type DnsConfigArgs struct {
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
	Fqdn         string   `pulumi:"fqdn"`
	RelativeName *string  `pulumi:"relativeName"`
	Ttl          *float64 `pulumi:"ttl"`
}





type DnsConfigResponseInput interface {
	pulumi.Input

	ToDnsConfigResponseOutput() DnsConfigResponseOutput
	ToDnsConfigResponseOutputWithContext(context.Context) DnsConfigResponseOutput
}

type DnsConfigResponseArgs struct {
	Fqdn         pulumi.StringInput     `pulumi:"fqdn"`
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

func (o DnsConfigResponseOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v DnsConfigResponse) string { return v.Fqdn }).(pulumi.StringOutput)
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
		return &v.Fqdn
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
	CustomHeaders         []EndpointPropertiesCustomHeaders `pulumi:"customHeaders"`
	EndpointLocation      *string                           `pulumi:"endpointLocation"`
	EndpointMonitorStatus *string                           `pulumi:"endpointMonitorStatus"`
	EndpointStatus        *string                           `pulumi:"endpointStatus"`
	GeoMapping            []string                          `pulumi:"geoMapping"`
	Id                    *string                           `pulumi:"id"`
	MinChildEndpoints     *float64                          `pulumi:"minChildEndpoints"`
	Name                  *string                           `pulumi:"name"`
	Priority              *float64                          `pulumi:"priority"`
	Target                *string                           `pulumi:"target"`
	TargetResourceId      *string                           `pulumi:"targetResourceId"`
	Type                  *string                           `pulumi:"type"`
	Weight                *float64                          `pulumi:"weight"`
}





type EndpointTypeInput interface {
	pulumi.Input

	ToEndpointTypeOutput() EndpointTypeOutput
	ToEndpointTypeOutputWithContext(context.Context) EndpointTypeOutput
}

type EndpointTypeArgs struct {
	CustomHeaders         EndpointPropertiesCustomHeadersArrayInput `pulumi:"customHeaders"`
	EndpointLocation      pulumi.StringPtrInput                     `pulumi:"endpointLocation"`
	EndpointMonitorStatus pulumi.StringPtrInput                     `pulumi:"endpointMonitorStatus"`
	EndpointStatus        pulumi.StringPtrInput                     `pulumi:"endpointStatus"`
	GeoMapping            pulumi.StringArrayInput                   `pulumi:"geoMapping"`
	Id                    pulumi.StringPtrInput                     `pulumi:"id"`
	MinChildEndpoints     pulumi.Float64PtrInput                    `pulumi:"minChildEndpoints"`
	Name                  pulumi.StringPtrInput                     `pulumi:"name"`
	Priority              pulumi.Float64PtrInput                    `pulumi:"priority"`
	Target                pulumi.StringPtrInput                     `pulumi:"target"`
	TargetResourceId      pulumi.StringPtrInput                     `pulumi:"targetResourceId"`
	Type                  pulumi.StringPtrInput                     `pulumi:"type"`
	Weight                pulumi.Float64PtrInput                    `pulumi:"weight"`
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

func (o EndpointTypeOutput) CustomHeaders() EndpointPropertiesCustomHeadersArrayOutput {
	return o.ApplyT(func(v EndpointType) []EndpointPropertiesCustomHeaders { return v.CustomHeaders }).(EndpointPropertiesCustomHeadersArrayOutput)
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

func (o EndpointTypeOutput) GeoMapping() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EndpointType) []string { return v.GeoMapping }).(pulumi.StringArrayOutput)
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

type EndpointPropertiesCustomHeaders struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type EndpointPropertiesCustomHeadersInput interface {
	pulumi.Input

	ToEndpointPropertiesCustomHeadersOutput() EndpointPropertiesCustomHeadersOutput
	ToEndpointPropertiesCustomHeadersOutputWithContext(context.Context) EndpointPropertiesCustomHeadersOutput
}

type EndpointPropertiesCustomHeadersArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (EndpointPropertiesCustomHeadersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointPropertiesCustomHeaders)(nil)).Elem()
}

func (i EndpointPropertiesCustomHeadersArgs) ToEndpointPropertiesCustomHeadersOutput() EndpointPropertiesCustomHeadersOutput {
	return i.ToEndpointPropertiesCustomHeadersOutputWithContext(context.Background())
}

func (i EndpointPropertiesCustomHeadersArgs) ToEndpointPropertiesCustomHeadersOutputWithContext(ctx context.Context) EndpointPropertiesCustomHeadersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesCustomHeadersOutput)
}





type EndpointPropertiesCustomHeadersArrayInput interface {
	pulumi.Input

	ToEndpointPropertiesCustomHeadersArrayOutput() EndpointPropertiesCustomHeadersArrayOutput
	ToEndpointPropertiesCustomHeadersArrayOutputWithContext(context.Context) EndpointPropertiesCustomHeadersArrayOutput
}

type EndpointPropertiesCustomHeadersArray []EndpointPropertiesCustomHeadersInput

func (EndpointPropertiesCustomHeadersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointPropertiesCustomHeaders)(nil)).Elem()
}

func (i EndpointPropertiesCustomHeadersArray) ToEndpointPropertiesCustomHeadersArrayOutput() EndpointPropertiesCustomHeadersArrayOutput {
	return i.ToEndpointPropertiesCustomHeadersArrayOutputWithContext(context.Background())
}

func (i EndpointPropertiesCustomHeadersArray) ToEndpointPropertiesCustomHeadersArrayOutputWithContext(ctx context.Context) EndpointPropertiesCustomHeadersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesCustomHeadersArrayOutput)
}

type EndpointPropertiesCustomHeadersOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesCustomHeadersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointPropertiesCustomHeaders)(nil)).Elem()
}

func (o EndpointPropertiesCustomHeadersOutput) ToEndpointPropertiesCustomHeadersOutput() EndpointPropertiesCustomHeadersOutput {
	return o
}

func (o EndpointPropertiesCustomHeadersOutput) ToEndpointPropertiesCustomHeadersOutputWithContext(ctx context.Context) EndpointPropertiesCustomHeadersOutput {
	return o
}

func (o EndpointPropertiesCustomHeadersOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointPropertiesCustomHeaders) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EndpointPropertiesCustomHeadersOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointPropertiesCustomHeaders) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type EndpointPropertiesCustomHeadersArrayOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesCustomHeadersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointPropertiesCustomHeaders)(nil)).Elem()
}

func (o EndpointPropertiesCustomHeadersArrayOutput) ToEndpointPropertiesCustomHeadersArrayOutput() EndpointPropertiesCustomHeadersArrayOutput {
	return o
}

func (o EndpointPropertiesCustomHeadersArrayOutput) ToEndpointPropertiesCustomHeadersArrayOutputWithContext(ctx context.Context) EndpointPropertiesCustomHeadersArrayOutput {
	return o
}

func (o EndpointPropertiesCustomHeadersArrayOutput) Index(i pulumi.IntInput) EndpointPropertiesCustomHeadersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointPropertiesCustomHeaders {
		return vs[0].([]EndpointPropertiesCustomHeaders)[vs[1].(int)]
	}).(EndpointPropertiesCustomHeadersOutput)
}

type EndpointPropertiesResponseCustomHeaders struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type EndpointPropertiesResponseCustomHeadersInput interface {
	pulumi.Input

	ToEndpointPropertiesResponseCustomHeadersOutput() EndpointPropertiesResponseCustomHeadersOutput
	ToEndpointPropertiesResponseCustomHeadersOutputWithContext(context.Context) EndpointPropertiesResponseCustomHeadersOutput
}

type EndpointPropertiesResponseCustomHeadersArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (EndpointPropertiesResponseCustomHeadersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointPropertiesResponseCustomHeaders)(nil)).Elem()
}

func (i EndpointPropertiesResponseCustomHeadersArgs) ToEndpointPropertiesResponseCustomHeadersOutput() EndpointPropertiesResponseCustomHeadersOutput {
	return i.ToEndpointPropertiesResponseCustomHeadersOutputWithContext(context.Background())
}

func (i EndpointPropertiesResponseCustomHeadersArgs) ToEndpointPropertiesResponseCustomHeadersOutputWithContext(ctx context.Context) EndpointPropertiesResponseCustomHeadersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesResponseCustomHeadersOutput)
}





type EndpointPropertiesResponseCustomHeadersArrayInput interface {
	pulumi.Input

	ToEndpointPropertiesResponseCustomHeadersArrayOutput() EndpointPropertiesResponseCustomHeadersArrayOutput
	ToEndpointPropertiesResponseCustomHeadersArrayOutputWithContext(context.Context) EndpointPropertiesResponseCustomHeadersArrayOutput
}

type EndpointPropertiesResponseCustomHeadersArray []EndpointPropertiesResponseCustomHeadersInput

func (EndpointPropertiesResponseCustomHeadersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointPropertiesResponseCustomHeaders)(nil)).Elem()
}

func (i EndpointPropertiesResponseCustomHeadersArray) ToEndpointPropertiesResponseCustomHeadersArrayOutput() EndpointPropertiesResponseCustomHeadersArrayOutput {
	return i.ToEndpointPropertiesResponseCustomHeadersArrayOutputWithContext(context.Background())
}

func (i EndpointPropertiesResponseCustomHeadersArray) ToEndpointPropertiesResponseCustomHeadersArrayOutputWithContext(ctx context.Context) EndpointPropertiesResponseCustomHeadersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesResponseCustomHeadersArrayOutput)
}

type EndpointPropertiesResponseCustomHeadersOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesResponseCustomHeadersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointPropertiesResponseCustomHeaders)(nil)).Elem()
}

func (o EndpointPropertiesResponseCustomHeadersOutput) ToEndpointPropertiesResponseCustomHeadersOutput() EndpointPropertiesResponseCustomHeadersOutput {
	return o
}

func (o EndpointPropertiesResponseCustomHeadersOutput) ToEndpointPropertiesResponseCustomHeadersOutputWithContext(ctx context.Context) EndpointPropertiesResponseCustomHeadersOutput {
	return o
}

func (o EndpointPropertiesResponseCustomHeadersOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointPropertiesResponseCustomHeaders) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EndpointPropertiesResponseCustomHeadersOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointPropertiesResponseCustomHeaders) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type EndpointPropertiesResponseCustomHeadersArrayOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesResponseCustomHeadersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointPropertiesResponseCustomHeaders)(nil)).Elem()
}

func (o EndpointPropertiesResponseCustomHeadersArrayOutput) ToEndpointPropertiesResponseCustomHeadersArrayOutput() EndpointPropertiesResponseCustomHeadersArrayOutput {
	return o
}

func (o EndpointPropertiesResponseCustomHeadersArrayOutput) ToEndpointPropertiesResponseCustomHeadersArrayOutputWithContext(ctx context.Context) EndpointPropertiesResponseCustomHeadersArrayOutput {
	return o
}

func (o EndpointPropertiesResponseCustomHeadersArrayOutput) Index(i pulumi.IntInput) EndpointPropertiesResponseCustomHeadersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointPropertiesResponseCustomHeaders {
		return vs[0].([]EndpointPropertiesResponseCustomHeaders)[vs[1].(int)]
	}).(EndpointPropertiesResponseCustomHeadersOutput)
}

type EndpointResponse struct {
	CustomHeaders         []EndpointPropertiesResponseCustomHeaders `pulumi:"customHeaders"`
	EndpointLocation      *string                                   `pulumi:"endpointLocation"`
	EndpointMonitorStatus *string                                   `pulumi:"endpointMonitorStatus"`
	EndpointStatus        *string                                   `pulumi:"endpointStatus"`
	GeoMapping            []string                                  `pulumi:"geoMapping"`
	Id                    *string                                   `pulumi:"id"`
	MinChildEndpoints     *float64                                  `pulumi:"minChildEndpoints"`
	Name                  *string                                   `pulumi:"name"`
	Priority              *float64                                  `pulumi:"priority"`
	Target                *string                                   `pulumi:"target"`
	TargetResourceId      *string                                   `pulumi:"targetResourceId"`
	Type                  *string                                   `pulumi:"type"`
	Weight                *float64                                  `pulumi:"weight"`
}





type EndpointResponseInput interface {
	pulumi.Input

	ToEndpointResponseOutput() EndpointResponseOutput
	ToEndpointResponseOutputWithContext(context.Context) EndpointResponseOutput
}

type EndpointResponseArgs struct {
	CustomHeaders         EndpointPropertiesResponseCustomHeadersArrayInput `pulumi:"customHeaders"`
	EndpointLocation      pulumi.StringPtrInput                             `pulumi:"endpointLocation"`
	EndpointMonitorStatus pulumi.StringPtrInput                             `pulumi:"endpointMonitorStatus"`
	EndpointStatus        pulumi.StringPtrInput                             `pulumi:"endpointStatus"`
	GeoMapping            pulumi.StringArrayInput                           `pulumi:"geoMapping"`
	Id                    pulumi.StringPtrInput                             `pulumi:"id"`
	MinChildEndpoints     pulumi.Float64PtrInput                            `pulumi:"minChildEndpoints"`
	Name                  pulumi.StringPtrInput                             `pulumi:"name"`
	Priority              pulumi.Float64PtrInput                            `pulumi:"priority"`
	Target                pulumi.StringPtrInput                             `pulumi:"target"`
	TargetResourceId      pulumi.StringPtrInput                             `pulumi:"targetResourceId"`
	Type                  pulumi.StringPtrInput                             `pulumi:"type"`
	Weight                pulumi.Float64PtrInput                            `pulumi:"weight"`
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

func (o EndpointResponseOutput) CustomHeaders() EndpointPropertiesResponseCustomHeadersArrayOutput {
	return o.ApplyT(func(v EndpointResponse) []EndpointPropertiesResponseCustomHeaders { return v.CustomHeaders }).(EndpointPropertiesResponseCustomHeadersArrayOutput)
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

func (o EndpointResponseOutput) GeoMapping() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EndpointResponse) []string { return v.GeoMapping }).(pulumi.StringArrayOutput)
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
	CustomHeaders             []MonitorConfigCustomHeaders            `pulumi:"customHeaders"`
	ExpectedStatusCodeRanges  []MonitorConfigExpectedStatusCodeRanges `pulumi:"expectedStatusCodeRanges"`
	IntervalInSeconds         *float64                                `pulumi:"intervalInSeconds"`
	Path                      *string                                 `pulumi:"path"`
	Port                      *float64                                `pulumi:"port"`
	ProfileMonitorStatus      *string                                 `pulumi:"profileMonitorStatus"`
	Protocol                  *string                                 `pulumi:"protocol"`
	TimeoutInSeconds          *float64                                `pulumi:"timeoutInSeconds"`
	ToleratedNumberOfFailures *float64                                `pulumi:"toleratedNumberOfFailures"`
}





type MonitorConfigInput interface {
	pulumi.Input

	ToMonitorConfigOutput() MonitorConfigOutput
	ToMonitorConfigOutputWithContext(context.Context) MonitorConfigOutput
}

type MonitorConfigArgs struct {
	CustomHeaders             MonitorConfigCustomHeadersArrayInput            `pulumi:"customHeaders"`
	ExpectedStatusCodeRanges  MonitorConfigExpectedStatusCodeRangesArrayInput `pulumi:"expectedStatusCodeRanges"`
	IntervalInSeconds         pulumi.Float64PtrInput                          `pulumi:"intervalInSeconds"`
	Path                      pulumi.StringPtrInput                           `pulumi:"path"`
	Port                      pulumi.Float64PtrInput                          `pulumi:"port"`
	ProfileMonitorStatus      pulumi.StringPtrInput                           `pulumi:"profileMonitorStatus"`
	Protocol                  pulumi.StringPtrInput                           `pulumi:"protocol"`
	TimeoutInSeconds          pulumi.Float64PtrInput                          `pulumi:"timeoutInSeconds"`
	ToleratedNumberOfFailures pulumi.Float64PtrInput                          `pulumi:"toleratedNumberOfFailures"`
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

func (o MonitorConfigOutput) CustomHeaders() MonitorConfigCustomHeadersArrayOutput {
	return o.ApplyT(func(v MonitorConfig) []MonitorConfigCustomHeaders { return v.CustomHeaders }).(MonitorConfigCustomHeadersArrayOutput)
}

func (o MonitorConfigOutput) ExpectedStatusCodeRanges() MonitorConfigExpectedStatusCodeRangesArrayOutput {
	return o.ApplyT(func(v MonitorConfig) []MonitorConfigExpectedStatusCodeRanges { return v.ExpectedStatusCodeRanges }).(MonitorConfigExpectedStatusCodeRangesArrayOutput)
}

func (o MonitorConfigOutput) IntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MonitorConfig) *float64 { return v.IntervalInSeconds }).(pulumi.Float64PtrOutput)
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

func (o MonitorConfigOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MonitorConfig) *float64 { return v.TimeoutInSeconds }).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigOutput) ToleratedNumberOfFailures() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MonitorConfig) *float64 { return v.ToleratedNumberOfFailures }).(pulumi.Float64PtrOutput)
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

func (o MonitorConfigPtrOutput) CustomHeaders() MonitorConfigCustomHeadersArrayOutput {
	return o.ApplyT(func(v *MonitorConfig) []MonitorConfigCustomHeaders {
		if v == nil {
			return nil
		}
		return v.CustomHeaders
	}).(MonitorConfigCustomHeadersArrayOutput)
}

func (o MonitorConfigPtrOutput) ExpectedStatusCodeRanges() MonitorConfigExpectedStatusCodeRangesArrayOutput {
	return o.ApplyT(func(v *MonitorConfig) []MonitorConfigExpectedStatusCodeRanges {
		if v == nil {
			return nil
		}
		return v.ExpectedStatusCodeRanges
	}).(MonitorConfigExpectedStatusCodeRangesArrayOutput)
}

func (o MonitorConfigPtrOutput) IntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MonitorConfig) *float64 {
		if v == nil {
			return nil
		}
		return v.IntervalInSeconds
	}).(pulumi.Float64PtrOutput)
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

func (o MonitorConfigPtrOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MonitorConfig) *float64 {
		if v == nil {
			return nil
		}
		return v.TimeoutInSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigPtrOutput) ToleratedNumberOfFailures() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MonitorConfig) *float64 {
		if v == nil {
			return nil
		}
		return v.ToleratedNumberOfFailures
	}).(pulumi.Float64PtrOutput)
}

type MonitorConfigCustomHeaders struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type MonitorConfigCustomHeadersInput interface {
	pulumi.Input

	ToMonitorConfigCustomHeadersOutput() MonitorConfigCustomHeadersOutput
	ToMonitorConfigCustomHeadersOutputWithContext(context.Context) MonitorConfigCustomHeadersOutput
}

type MonitorConfigCustomHeadersArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (MonitorConfigCustomHeadersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfigCustomHeaders)(nil)).Elem()
}

func (i MonitorConfigCustomHeadersArgs) ToMonitorConfigCustomHeadersOutput() MonitorConfigCustomHeadersOutput {
	return i.ToMonitorConfigCustomHeadersOutputWithContext(context.Background())
}

func (i MonitorConfigCustomHeadersArgs) ToMonitorConfigCustomHeadersOutputWithContext(ctx context.Context) MonitorConfigCustomHeadersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigCustomHeadersOutput)
}





type MonitorConfigCustomHeadersArrayInput interface {
	pulumi.Input

	ToMonitorConfigCustomHeadersArrayOutput() MonitorConfigCustomHeadersArrayOutput
	ToMonitorConfigCustomHeadersArrayOutputWithContext(context.Context) MonitorConfigCustomHeadersArrayOutput
}

type MonitorConfigCustomHeadersArray []MonitorConfigCustomHeadersInput

func (MonitorConfigCustomHeadersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonitorConfigCustomHeaders)(nil)).Elem()
}

func (i MonitorConfigCustomHeadersArray) ToMonitorConfigCustomHeadersArrayOutput() MonitorConfigCustomHeadersArrayOutput {
	return i.ToMonitorConfigCustomHeadersArrayOutputWithContext(context.Background())
}

func (i MonitorConfigCustomHeadersArray) ToMonitorConfigCustomHeadersArrayOutputWithContext(ctx context.Context) MonitorConfigCustomHeadersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigCustomHeadersArrayOutput)
}

type MonitorConfigCustomHeadersOutput struct{ *pulumi.OutputState }

func (MonitorConfigCustomHeadersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfigCustomHeaders)(nil)).Elem()
}

func (o MonitorConfigCustomHeadersOutput) ToMonitorConfigCustomHeadersOutput() MonitorConfigCustomHeadersOutput {
	return o
}

func (o MonitorConfigCustomHeadersOutput) ToMonitorConfigCustomHeadersOutputWithContext(ctx context.Context) MonitorConfigCustomHeadersOutput {
	return o
}

func (o MonitorConfigCustomHeadersOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfigCustomHeaders) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o MonitorConfigCustomHeadersOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfigCustomHeaders) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type MonitorConfigCustomHeadersArrayOutput struct{ *pulumi.OutputState }

func (MonitorConfigCustomHeadersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonitorConfigCustomHeaders)(nil)).Elem()
}

func (o MonitorConfigCustomHeadersArrayOutput) ToMonitorConfigCustomHeadersArrayOutput() MonitorConfigCustomHeadersArrayOutput {
	return o
}

func (o MonitorConfigCustomHeadersArrayOutput) ToMonitorConfigCustomHeadersArrayOutputWithContext(ctx context.Context) MonitorConfigCustomHeadersArrayOutput {
	return o
}

func (o MonitorConfigCustomHeadersArrayOutput) Index(i pulumi.IntInput) MonitorConfigCustomHeadersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MonitorConfigCustomHeaders {
		return vs[0].([]MonitorConfigCustomHeaders)[vs[1].(int)]
	}).(MonitorConfigCustomHeadersOutput)
}

type MonitorConfigExpectedStatusCodeRanges struct {
	Max *int `pulumi:"max"`
	Min *int `pulumi:"min"`
}





type MonitorConfigExpectedStatusCodeRangesInput interface {
	pulumi.Input

	ToMonitorConfigExpectedStatusCodeRangesOutput() MonitorConfigExpectedStatusCodeRangesOutput
	ToMonitorConfigExpectedStatusCodeRangesOutputWithContext(context.Context) MonitorConfigExpectedStatusCodeRangesOutput
}

type MonitorConfigExpectedStatusCodeRangesArgs struct {
	Max pulumi.IntPtrInput `pulumi:"max"`
	Min pulumi.IntPtrInput `pulumi:"min"`
}

func (MonitorConfigExpectedStatusCodeRangesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfigExpectedStatusCodeRanges)(nil)).Elem()
}

func (i MonitorConfigExpectedStatusCodeRangesArgs) ToMonitorConfigExpectedStatusCodeRangesOutput() MonitorConfigExpectedStatusCodeRangesOutput {
	return i.ToMonitorConfigExpectedStatusCodeRangesOutputWithContext(context.Background())
}

func (i MonitorConfigExpectedStatusCodeRangesArgs) ToMonitorConfigExpectedStatusCodeRangesOutputWithContext(ctx context.Context) MonitorConfigExpectedStatusCodeRangesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigExpectedStatusCodeRangesOutput)
}





type MonitorConfigExpectedStatusCodeRangesArrayInput interface {
	pulumi.Input

	ToMonitorConfigExpectedStatusCodeRangesArrayOutput() MonitorConfigExpectedStatusCodeRangesArrayOutput
	ToMonitorConfigExpectedStatusCodeRangesArrayOutputWithContext(context.Context) MonitorConfigExpectedStatusCodeRangesArrayOutput
}

type MonitorConfigExpectedStatusCodeRangesArray []MonitorConfigExpectedStatusCodeRangesInput

func (MonitorConfigExpectedStatusCodeRangesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonitorConfigExpectedStatusCodeRanges)(nil)).Elem()
}

func (i MonitorConfigExpectedStatusCodeRangesArray) ToMonitorConfigExpectedStatusCodeRangesArrayOutput() MonitorConfigExpectedStatusCodeRangesArrayOutput {
	return i.ToMonitorConfigExpectedStatusCodeRangesArrayOutputWithContext(context.Background())
}

func (i MonitorConfigExpectedStatusCodeRangesArray) ToMonitorConfigExpectedStatusCodeRangesArrayOutputWithContext(ctx context.Context) MonitorConfigExpectedStatusCodeRangesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigExpectedStatusCodeRangesArrayOutput)
}

type MonitorConfigExpectedStatusCodeRangesOutput struct{ *pulumi.OutputState }

func (MonitorConfigExpectedStatusCodeRangesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfigExpectedStatusCodeRanges)(nil)).Elem()
}

func (o MonitorConfigExpectedStatusCodeRangesOutput) ToMonitorConfigExpectedStatusCodeRangesOutput() MonitorConfigExpectedStatusCodeRangesOutput {
	return o
}

func (o MonitorConfigExpectedStatusCodeRangesOutput) ToMonitorConfigExpectedStatusCodeRangesOutputWithContext(ctx context.Context) MonitorConfigExpectedStatusCodeRangesOutput {
	return o
}

func (o MonitorConfigExpectedStatusCodeRangesOutput) Max() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MonitorConfigExpectedStatusCodeRanges) *int { return v.Max }).(pulumi.IntPtrOutput)
}

func (o MonitorConfigExpectedStatusCodeRangesOutput) Min() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MonitorConfigExpectedStatusCodeRanges) *int { return v.Min }).(pulumi.IntPtrOutput)
}

type MonitorConfigExpectedStatusCodeRangesArrayOutput struct{ *pulumi.OutputState }

func (MonitorConfigExpectedStatusCodeRangesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonitorConfigExpectedStatusCodeRanges)(nil)).Elem()
}

func (o MonitorConfigExpectedStatusCodeRangesArrayOutput) ToMonitorConfigExpectedStatusCodeRangesArrayOutput() MonitorConfigExpectedStatusCodeRangesArrayOutput {
	return o
}

func (o MonitorConfigExpectedStatusCodeRangesArrayOutput) ToMonitorConfigExpectedStatusCodeRangesArrayOutputWithContext(ctx context.Context) MonitorConfigExpectedStatusCodeRangesArrayOutput {
	return o
}

func (o MonitorConfigExpectedStatusCodeRangesArrayOutput) Index(i pulumi.IntInput) MonitorConfigExpectedStatusCodeRangesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MonitorConfigExpectedStatusCodeRanges {
		return vs[0].([]MonitorConfigExpectedStatusCodeRanges)[vs[1].(int)]
	}).(MonitorConfigExpectedStatusCodeRangesOutput)
}

type MonitorConfigResponse struct {
	CustomHeaders             []MonitorConfigResponseCustomHeaders            `pulumi:"customHeaders"`
	ExpectedStatusCodeRanges  []MonitorConfigResponseExpectedStatusCodeRanges `pulumi:"expectedStatusCodeRanges"`
	IntervalInSeconds         *float64                                        `pulumi:"intervalInSeconds"`
	Path                      *string                                         `pulumi:"path"`
	Port                      *float64                                        `pulumi:"port"`
	ProfileMonitorStatus      *string                                         `pulumi:"profileMonitorStatus"`
	Protocol                  *string                                         `pulumi:"protocol"`
	TimeoutInSeconds          *float64                                        `pulumi:"timeoutInSeconds"`
	ToleratedNumberOfFailures *float64                                        `pulumi:"toleratedNumberOfFailures"`
}





type MonitorConfigResponseInput interface {
	pulumi.Input

	ToMonitorConfigResponseOutput() MonitorConfigResponseOutput
	ToMonitorConfigResponseOutputWithContext(context.Context) MonitorConfigResponseOutput
}

type MonitorConfigResponseArgs struct {
	CustomHeaders             MonitorConfigResponseCustomHeadersArrayInput            `pulumi:"customHeaders"`
	ExpectedStatusCodeRanges  MonitorConfigResponseExpectedStatusCodeRangesArrayInput `pulumi:"expectedStatusCodeRanges"`
	IntervalInSeconds         pulumi.Float64PtrInput                                  `pulumi:"intervalInSeconds"`
	Path                      pulumi.StringPtrInput                                   `pulumi:"path"`
	Port                      pulumi.Float64PtrInput                                  `pulumi:"port"`
	ProfileMonitorStatus      pulumi.StringPtrInput                                   `pulumi:"profileMonitorStatus"`
	Protocol                  pulumi.StringPtrInput                                   `pulumi:"protocol"`
	TimeoutInSeconds          pulumi.Float64PtrInput                                  `pulumi:"timeoutInSeconds"`
	ToleratedNumberOfFailures pulumi.Float64PtrInput                                  `pulumi:"toleratedNumberOfFailures"`
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

func (o MonitorConfigResponseOutput) CustomHeaders() MonitorConfigResponseCustomHeadersArrayOutput {
	return o.ApplyT(func(v MonitorConfigResponse) []MonitorConfigResponseCustomHeaders { return v.CustomHeaders }).(MonitorConfigResponseCustomHeadersArrayOutput)
}

func (o MonitorConfigResponseOutput) ExpectedStatusCodeRanges() MonitorConfigResponseExpectedStatusCodeRangesArrayOutput {
	return o.ApplyT(func(v MonitorConfigResponse) []MonitorConfigResponseExpectedStatusCodeRanges {
		return v.ExpectedStatusCodeRanges
	}).(MonitorConfigResponseExpectedStatusCodeRangesArrayOutput)
}

func (o MonitorConfigResponseOutput) IntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MonitorConfigResponse) *float64 { return v.IntervalInSeconds }).(pulumi.Float64PtrOutput)
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

func (o MonitorConfigResponseOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MonitorConfigResponse) *float64 { return v.TimeoutInSeconds }).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigResponseOutput) ToleratedNumberOfFailures() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MonitorConfigResponse) *float64 { return v.ToleratedNumberOfFailures }).(pulumi.Float64PtrOutput)
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

func (o MonitorConfigResponsePtrOutput) CustomHeaders() MonitorConfigResponseCustomHeadersArrayOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) []MonitorConfigResponseCustomHeaders {
		if v == nil {
			return nil
		}
		return v.CustomHeaders
	}).(MonitorConfigResponseCustomHeadersArrayOutput)
}

func (o MonitorConfigResponsePtrOutput) ExpectedStatusCodeRanges() MonitorConfigResponseExpectedStatusCodeRangesArrayOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) []MonitorConfigResponseExpectedStatusCodeRanges {
		if v == nil {
			return nil
		}
		return v.ExpectedStatusCodeRanges
	}).(MonitorConfigResponseExpectedStatusCodeRangesArrayOutput)
}

func (o MonitorConfigResponsePtrOutput) IntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.IntervalInSeconds
	}).(pulumi.Float64PtrOutput)
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

func (o MonitorConfigResponsePtrOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.TimeoutInSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigResponsePtrOutput) ToleratedNumberOfFailures() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.ToleratedNumberOfFailures
	}).(pulumi.Float64PtrOutput)
}

type MonitorConfigResponseCustomHeaders struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type MonitorConfigResponseCustomHeadersInput interface {
	pulumi.Input

	ToMonitorConfigResponseCustomHeadersOutput() MonitorConfigResponseCustomHeadersOutput
	ToMonitorConfigResponseCustomHeadersOutputWithContext(context.Context) MonitorConfigResponseCustomHeadersOutput
}

type MonitorConfigResponseCustomHeadersArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (MonitorConfigResponseCustomHeadersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfigResponseCustomHeaders)(nil)).Elem()
}

func (i MonitorConfigResponseCustomHeadersArgs) ToMonitorConfigResponseCustomHeadersOutput() MonitorConfigResponseCustomHeadersOutput {
	return i.ToMonitorConfigResponseCustomHeadersOutputWithContext(context.Background())
}

func (i MonitorConfigResponseCustomHeadersArgs) ToMonitorConfigResponseCustomHeadersOutputWithContext(ctx context.Context) MonitorConfigResponseCustomHeadersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigResponseCustomHeadersOutput)
}





type MonitorConfigResponseCustomHeadersArrayInput interface {
	pulumi.Input

	ToMonitorConfigResponseCustomHeadersArrayOutput() MonitorConfigResponseCustomHeadersArrayOutput
	ToMonitorConfigResponseCustomHeadersArrayOutputWithContext(context.Context) MonitorConfigResponseCustomHeadersArrayOutput
}

type MonitorConfigResponseCustomHeadersArray []MonitorConfigResponseCustomHeadersInput

func (MonitorConfigResponseCustomHeadersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonitorConfigResponseCustomHeaders)(nil)).Elem()
}

func (i MonitorConfigResponseCustomHeadersArray) ToMonitorConfigResponseCustomHeadersArrayOutput() MonitorConfigResponseCustomHeadersArrayOutput {
	return i.ToMonitorConfigResponseCustomHeadersArrayOutputWithContext(context.Background())
}

func (i MonitorConfigResponseCustomHeadersArray) ToMonitorConfigResponseCustomHeadersArrayOutputWithContext(ctx context.Context) MonitorConfigResponseCustomHeadersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigResponseCustomHeadersArrayOutput)
}

type MonitorConfigResponseCustomHeadersOutput struct{ *pulumi.OutputState }

func (MonitorConfigResponseCustomHeadersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfigResponseCustomHeaders)(nil)).Elem()
}

func (o MonitorConfigResponseCustomHeadersOutput) ToMonitorConfigResponseCustomHeadersOutput() MonitorConfigResponseCustomHeadersOutput {
	return o
}

func (o MonitorConfigResponseCustomHeadersOutput) ToMonitorConfigResponseCustomHeadersOutputWithContext(ctx context.Context) MonitorConfigResponseCustomHeadersOutput {
	return o
}

func (o MonitorConfigResponseCustomHeadersOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfigResponseCustomHeaders) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o MonitorConfigResponseCustomHeadersOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfigResponseCustomHeaders) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type MonitorConfigResponseCustomHeadersArrayOutput struct{ *pulumi.OutputState }

func (MonitorConfigResponseCustomHeadersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonitorConfigResponseCustomHeaders)(nil)).Elem()
}

func (o MonitorConfigResponseCustomHeadersArrayOutput) ToMonitorConfigResponseCustomHeadersArrayOutput() MonitorConfigResponseCustomHeadersArrayOutput {
	return o
}

func (o MonitorConfigResponseCustomHeadersArrayOutput) ToMonitorConfigResponseCustomHeadersArrayOutputWithContext(ctx context.Context) MonitorConfigResponseCustomHeadersArrayOutput {
	return o
}

func (o MonitorConfigResponseCustomHeadersArrayOutput) Index(i pulumi.IntInput) MonitorConfigResponseCustomHeadersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MonitorConfigResponseCustomHeaders {
		return vs[0].([]MonitorConfigResponseCustomHeaders)[vs[1].(int)]
	}).(MonitorConfigResponseCustomHeadersOutput)
}

type MonitorConfigResponseExpectedStatusCodeRanges struct {
	Max *int `pulumi:"max"`
	Min *int `pulumi:"min"`
}





type MonitorConfigResponseExpectedStatusCodeRangesInput interface {
	pulumi.Input

	ToMonitorConfigResponseExpectedStatusCodeRangesOutput() MonitorConfigResponseExpectedStatusCodeRangesOutput
	ToMonitorConfigResponseExpectedStatusCodeRangesOutputWithContext(context.Context) MonitorConfigResponseExpectedStatusCodeRangesOutput
}

type MonitorConfigResponseExpectedStatusCodeRangesArgs struct {
	Max pulumi.IntPtrInput `pulumi:"max"`
	Min pulumi.IntPtrInput `pulumi:"min"`
}

func (MonitorConfigResponseExpectedStatusCodeRangesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfigResponseExpectedStatusCodeRanges)(nil)).Elem()
}

func (i MonitorConfigResponseExpectedStatusCodeRangesArgs) ToMonitorConfigResponseExpectedStatusCodeRangesOutput() MonitorConfigResponseExpectedStatusCodeRangesOutput {
	return i.ToMonitorConfigResponseExpectedStatusCodeRangesOutputWithContext(context.Background())
}

func (i MonitorConfigResponseExpectedStatusCodeRangesArgs) ToMonitorConfigResponseExpectedStatusCodeRangesOutputWithContext(ctx context.Context) MonitorConfigResponseExpectedStatusCodeRangesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigResponseExpectedStatusCodeRangesOutput)
}





type MonitorConfigResponseExpectedStatusCodeRangesArrayInput interface {
	pulumi.Input

	ToMonitorConfigResponseExpectedStatusCodeRangesArrayOutput() MonitorConfigResponseExpectedStatusCodeRangesArrayOutput
	ToMonitorConfigResponseExpectedStatusCodeRangesArrayOutputWithContext(context.Context) MonitorConfigResponseExpectedStatusCodeRangesArrayOutput
}

type MonitorConfigResponseExpectedStatusCodeRangesArray []MonitorConfigResponseExpectedStatusCodeRangesInput

func (MonitorConfigResponseExpectedStatusCodeRangesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonitorConfigResponseExpectedStatusCodeRanges)(nil)).Elem()
}

func (i MonitorConfigResponseExpectedStatusCodeRangesArray) ToMonitorConfigResponseExpectedStatusCodeRangesArrayOutput() MonitorConfigResponseExpectedStatusCodeRangesArrayOutput {
	return i.ToMonitorConfigResponseExpectedStatusCodeRangesArrayOutputWithContext(context.Background())
}

func (i MonitorConfigResponseExpectedStatusCodeRangesArray) ToMonitorConfigResponseExpectedStatusCodeRangesArrayOutputWithContext(ctx context.Context) MonitorConfigResponseExpectedStatusCodeRangesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigResponseExpectedStatusCodeRangesArrayOutput)
}

type MonitorConfigResponseExpectedStatusCodeRangesOutput struct{ *pulumi.OutputState }

func (MonitorConfigResponseExpectedStatusCodeRangesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfigResponseExpectedStatusCodeRanges)(nil)).Elem()
}

func (o MonitorConfigResponseExpectedStatusCodeRangesOutput) ToMonitorConfigResponseExpectedStatusCodeRangesOutput() MonitorConfigResponseExpectedStatusCodeRangesOutput {
	return o
}

func (o MonitorConfigResponseExpectedStatusCodeRangesOutput) ToMonitorConfigResponseExpectedStatusCodeRangesOutputWithContext(ctx context.Context) MonitorConfigResponseExpectedStatusCodeRangesOutput {
	return o
}

func (o MonitorConfigResponseExpectedStatusCodeRangesOutput) Max() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MonitorConfigResponseExpectedStatusCodeRanges) *int { return v.Max }).(pulumi.IntPtrOutput)
}

func (o MonitorConfigResponseExpectedStatusCodeRangesOutput) Min() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MonitorConfigResponseExpectedStatusCodeRanges) *int { return v.Min }).(pulumi.IntPtrOutput)
}

type MonitorConfigResponseExpectedStatusCodeRangesArrayOutput struct{ *pulumi.OutputState }

func (MonitorConfigResponseExpectedStatusCodeRangesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonitorConfigResponseExpectedStatusCodeRanges)(nil)).Elem()
}

func (o MonitorConfigResponseExpectedStatusCodeRangesArrayOutput) ToMonitorConfigResponseExpectedStatusCodeRangesArrayOutput() MonitorConfigResponseExpectedStatusCodeRangesArrayOutput {
	return o
}

func (o MonitorConfigResponseExpectedStatusCodeRangesArrayOutput) ToMonitorConfigResponseExpectedStatusCodeRangesArrayOutputWithContext(ctx context.Context) MonitorConfigResponseExpectedStatusCodeRangesArrayOutput {
	return o
}

func (o MonitorConfigResponseExpectedStatusCodeRangesArrayOutput) Index(i pulumi.IntInput) MonitorConfigResponseExpectedStatusCodeRangesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MonitorConfigResponseExpectedStatusCodeRanges {
		return vs[0].([]MonitorConfigResponseExpectedStatusCodeRanges)[vs[1].(int)]
	}).(MonitorConfigResponseExpectedStatusCodeRangesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsConfigInput)(nil)).Elem(), DnsConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsConfigPtrInput)(nil)).Elem(), DnsConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsConfigResponseInput)(nil)).Elem(), DnsConfigResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsConfigResponsePtrInput)(nil)).Elem(), DnsConfigResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointTypeInput)(nil)).Elem(), EndpointTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointTypeArrayInput)(nil)).Elem(), EndpointTypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointPropertiesCustomHeadersInput)(nil)).Elem(), EndpointPropertiesCustomHeadersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointPropertiesCustomHeadersArrayInput)(nil)).Elem(), EndpointPropertiesCustomHeadersArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointPropertiesResponseCustomHeadersInput)(nil)).Elem(), EndpointPropertiesResponseCustomHeadersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointPropertiesResponseCustomHeadersArrayInput)(nil)).Elem(), EndpointPropertiesResponseCustomHeadersArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointResponseInput)(nil)).Elem(), EndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointResponseArrayInput)(nil)).Elem(), EndpointResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorConfigInput)(nil)).Elem(), MonitorConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorConfigPtrInput)(nil)).Elem(), MonitorConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorConfigCustomHeadersInput)(nil)).Elem(), MonitorConfigCustomHeadersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorConfigCustomHeadersArrayInput)(nil)).Elem(), MonitorConfigCustomHeadersArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorConfigExpectedStatusCodeRangesInput)(nil)).Elem(), MonitorConfigExpectedStatusCodeRangesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorConfigExpectedStatusCodeRangesArrayInput)(nil)).Elem(), MonitorConfigExpectedStatusCodeRangesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorConfigResponseInput)(nil)).Elem(), MonitorConfigResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorConfigResponsePtrInput)(nil)).Elem(), MonitorConfigResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorConfigResponseCustomHeadersInput)(nil)).Elem(), MonitorConfigResponseCustomHeadersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorConfigResponseCustomHeadersArrayInput)(nil)).Elem(), MonitorConfigResponseCustomHeadersArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorConfigResponseExpectedStatusCodeRangesInput)(nil)).Elem(), MonitorConfigResponseExpectedStatusCodeRangesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorConfigResponseExpectedStatusCodeRangesArrayInput)(nil)).Elem(), MonitorConfigResponseExpectedStatusCodeRangesArray{})
	pulumi.RegisterOutputType(DnsConfigOutput{})
	pulumi.RegisterOutputType(DnsConfigPtrOutput{})
	pulumi.RegisterOutputType(DnsConfigResponseOutput{})
	pulumi.RegisterOutputType(DnsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(EndpointTypeOutput{})
	pulumi.RegisterOutputType(EndpointTypeArrayOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesCustomHeadersOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesCustomHeadersArrayOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesResponseCustomHeadersOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesResponseCustomHeadersArrayOutput{})
	pulumi.RegisterOutputType(EndpointResponseOutput{})
	pulumi.RegisterOutputType(EndpointResponseArrayOutput{})
	pulumi.RegisterOutputType(MonitorConfigOutput{})
	pulumi.RegisterOutputType(MonitorConfigPtrOutput{})
	pulumi.RegisterOutputType(MonitorConfigCustomHeadersOutput{})
	pulumi.RegisterOutputType(MonitorConfigCustomHeadersArrayOutput{})
	pulumi.RegisterOutputType(MonitorConfigExpectedStatusCodeRangesOutput{})
	pulumi.RegisterOutputType(MonitorConfigExpectedStatusCodeRangesArrayOutput{})
	pulumi.RegisterOutputType(MonitorConfigResponseOutput{})
	pulumi.RegisterOutputType(MonitorConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(MonitorConfigResponseCustomHeadersOutput{})
	pulumi.RegisterOutputType(MonitorConfigResponseCustomHeadersArrayOutput{})
	pulumi.RegisterOutputType(MonitorConfigResponseExpectedStatusCodeRangesOutput{})
	pulumi.RegisterOutputType(MonitorConfigResponseExpectedStatusCodeRangesArrayOutput{})
}
