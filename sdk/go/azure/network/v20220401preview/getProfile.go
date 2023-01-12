


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProfile(ctx *pulumi.Context, args *LookupProfileArgs, opts ...pulumi.InvokeOption) (*LookupProfileResult, error) {
	var rv LookupProfileResult
	err := ctx.Invoke("azure-native:network/v20220401preview:getProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProfileArgs struct {
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupProfileResult struct {
	AllowedEndpointRecordTypes  []string               `pulumi:"allowedEndpointRecordTypes"`
	DnsConfig                   *DnsConfigResponse     `pulumi:"dnsConfig"`
	Endpoints                   []EndpointResponse     `pulumi:"endpoints"`
	Id                          *string                `pulumi:"id"`
	Location                    *string                `pulumi:"location"`
	MaxReturn                   *float64               `pulumi:"maxReturn"`
	MonitorConfig               *MonitorConfigResponse `pulumi:"monitorConfig"`
	Name                        *string                `pulumi:"name"`
	ProfileStatus               *string                `pulumi:"profileStatus"`
	Tags                        map[string]string      `pulumi:"tags"`
	TrafficRoutingMethod        *string                `pulumi:"trafficRoutingMethod"`
	TrafficViewEnrollmentStatus *string                `pulumi:"trafficViewEnrollmentStatus"`
	Type                        *string                `pulumi:"type"`
}

func LookupProfileOutput(ctx *pulumi.Context, args LookupProfileOutputArgs, opts ...pulumi.InvokeOption) LookupProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProfileResult, error) {
			args := v.(LookupProfileArgs)
			r, err := LookupProfile(ctx, &args, opts...)
			var s LookupProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProfileResultOutput)
}

type LookupProfileOutputArgs struct {
	ProfileName       pulumi.StringInput `pulumi:"profileName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfileArgs)(nil)).Elem()
}


type LookupProfileResultOutput struct{ *pulumi.OutputState }

func (LookupProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfileResult)(nil)).Elem()
}

func (o LookupProfileResultOutput) ToLookupProfileResultOutput() LookupProfileResultOutput {
	return o
}

func (o LookupProfileResultOutput) ToLookupProfileResultOutputWithContext(ctx context.Context) LookupProfileResultOutput {
	return o
}

func (o LookupProfileResultOutput) AllowedEndpointRecordTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupProfileResult) []string { return v.AllowedEndpointRecordTypes }).(pulumi.StringArrayOutput)
}

func (o LookupProfileResultOutput) DnsConfig() DnsConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *DnsConfigResponse { return v.DnsConfig }).(DnsConfigResponsePtrOutput)
}

func (o LookupProfileResultOutput) Endpoints() EndpointResponseArrayOutput {
	return o.ApplyT(func(v LookupProfileResult) []EndpointResponse { return v.Endpoints }).(EndpointResponseArrayOutput)
}

func (o LookupProfileResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupProfileResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupProfileResultOutput) MaxReturn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *float64 { return v.MaxReturn }).(pulumi.Float64PtrOutput)
}

func (o LookupProfileResultOutput) MonitorConfig() MonitorConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *MonitorConfigResponse { return v.MonitorConfig }).(MonitorConfigResponsePtrOutput)
}

func (o LookupProfileResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupProfileResultOutput) ProfileStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.ProfileStatus }).(pulumi.StringPtrOutput)
}

func (o LookupProfileResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProfileResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupProfileResultOutput) TrafficRoutingMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.TrafficRoutingMethod }).(pulumi.StringPtrOutput)
}

func (o LookupProfileResultOutput) TrafficViewEnrollmentStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.TrafficViewEnrollmentStatus }).(pulumi.StringPtrOutput)
}

func (o LookupProfileResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProfileResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProfileResultOutput{})
}
