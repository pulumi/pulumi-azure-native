


package v20180401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupEndpoint(ctx *pulumi.Context, args *LookupEndpointArgs, opts ...pulumi.InvokeOption) (*LookupEndpointResult, error) {
	var rv LookupEndpointResult
	err := ctx.Invoke("azure-native:network/v20180401:getEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEndpointArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	EndpointType      string `pulumi:"endpointType"`
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupEndpointResult struct {
	CustomHeaders         []EndpointPropertiesResponseCustomHeaders `pulumi:"customHeaders"`
	EndpointLocation      *string                                   `pulumi:"endpointLocation"`
	EndpointMonitorStatus *string                                   `pulumi:"endpointMonitorStatus"`
	EndpointStatus        *string                                   `pulumi:"endpointStatus"`
	GeoMapping            []string                                  `pulumi:"geoMapping"`
	Id                    *string                                   `pulumi:"id"`
	MinChildEndpoints     *float64                                  `pulumi:"minChildEndpoints"`
	Name                  *string                                   `pulumi:"name"`
	Priority              *float64                                  `pulumi:"priority"`
	Subnets               []EndpointPropertiesResponseSubnets       `pulumi:"subnets"`
	Target                *string                                   `pulumi:"target"`
	TargetResourceId      *string                                   `pulumi:"targetResourceId"`
	Type                  *string                                   `pulumi:"type"`
	Weight                *float64                                  `pulumi:"weight"`
}

func LookupEndpointOutput(ctx *pulumi.Context, args LookupEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEndpointResult, error) {
			args := v.(LookupEndpointArgs)
			r, err := LookupEndpoint(ctx, &args, opts...)
			var s LookupEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEndpointResultOutput)
}

type LookupEndpointOutputArgs struct {
	EndpointName      pulumi.StringInput `pulumi:"endpointName"`
	EndpointType      pulumi.StringInput `pulumi:"endpointType"`
	ProfileName       pulumi.StringInput `pulumi:"profileName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointArgs)(nil)).Elem()
}


type LookupEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointResult)(nil)).Elem()
}

func (o LookupEndpointResultOutput) ToLookupEndpointResultOutput() LookupEndpointResultOutput {
	return o
}

func (o LookupEndpointResultOutput) ToLookupEndpointResultOutputWithContext(ctx context.Context) LookupEndpointResultOutput {
	return o
}

func (o LookupEndpointResultOutput) CustomHeaders() EndpointPropertiesResponseCustomHeadersArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []EndpointPropertiesResponseCustomHeaders { return v.CustomHeaders }).(EndpointPropertiesResponseCustomHeadersArrayOutput)
}

func (o LookupEndpointResultOutput) EndpointLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.EndpointLocation }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) EndpointMonitorStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.EndpointMonitorStatus }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) EndpointStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.EndpointStatus }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) GeoMapping() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []string { return v.GeoMapping }).(pulumi.StringArrayOutput)
}

func (o LookupEndpointResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) MinChildEndpoints() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *float64 { return v.MinChildEndpoints }).(pulumi.Float64PtrOutput)
}

func (o LookupEndpointResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) Priority() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *float64 { return v.Priority }).(pulumi.Float64PtrOutput)
}

func (o LookupEndpointResultOutput) Subnets() EndpointPropertiesResponseSubnetsArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []EndpointPropertiesResponseSubnets { return v.Subnets }).(EndpointPropertiesResponseSubnetsArrayOutput)
}

func (o LookupEndpointResultOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.TargetResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) Weight() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *float64 { return v.Weight }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEndpointResultOutput{})
}
