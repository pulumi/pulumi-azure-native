


package v20190201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExpressRoutePort(ctx *pulumi.Context, args *LookupExpressRoutePortArgs, opts ...pulumi.InvokeOption) (*LookupExpressRoutePortResult, error) {
	var rv LookupExpressRoutePortResult
	err := ctx.Invoke("azure-native:network/v20190201:getExpressRoutePort", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRoutePortArgs struct {
	ExpressRoutePortName string `pulumi:"expressRoutePortName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupExpressRoutePortResult struct {
	AllocationDate             string                     `pulumi:"allocationDate"`
	BandwidthInGbps            *int                       `pulumi:"bandwidthInGbps"`
	Circuits                   []SubResourceResponse      `pulumi:"circuits"`
	Encapsulation              *string                    `pulumi:"encapsulation"`
	Etag                       string                     `pulumi:"etag"`
	EtherType                  string                     `pulumi:"etherType"`
	Id                         *string                    `pulumi:"id"`
	Links                      []ExpressRouteLinkResponse `pulumi:"links"`
	Location                   *string                    `pulumi:"location"`
	Mtu                        string                     `pulumi:"mtu"`
	Name                       string                     `pulumi:"name"`
	PeeringLocation            *string                    `pulumi:"peeringLocation"`
	ProvisionedBandwidthInGbps float64                    `pulumi:"provisionedBandwidthInGbps"`
	ProvisioningState          string                     `pulumi:"provisioningState"`
	ResourceGuid               *string                    `pulumi:"resourceGuid"`
	Tags                       map[string]string          `pulumi:"tags"`
	Type                       string                     `pulumi:"type"`
}

func LookupExpressRoutePortOutput(ctx *pulumi.Context, args LookupExpressRoutePortOutputArgs, opts ...pulumi.InvokeOption) LookupExpressRoutePortResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExpressRoutePortResult, error) {
			args := v.(LookupExpressRoutePortArgs)
			r, err := LookupExpressRoutePort(ctx, &args, opts...)
			var s LookupExpressRoutePortResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExpressRoutePortResultOutput)
}

type LookupExpressRoutePortOutputArgs struct {
	ExpressRoutePortName pulumi.StringInput `pulumi:"expressRoutePortName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExpressRoutePortOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRoutePortArgs)(nil)).Elem()
}


type LookupExpressRoutePortResultOutput struct{ *pulumi.OutputState }

func (LookupExpressRoutePortResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRoutePortResult)(nil)).Elem()
}

func (o LookupExpressRoutePortResultOutput) ToLookupExpressRoutePortResultOutput() LookupExpressRoutePortResultOutput {
	return o
}

func (o LookupExpressRoutePortResultOutput) ToLookupExpressRoutePortResultOutputWithContext(ctx context.Context) LookupExpressRoutePortResultOutput {
	return o
}

func (o LookupExpressRoutePortResultOutput) AllocationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) string { return v.AllocationDate }).(pulumi.StringOutput)
}

func (o LookupExpressRoutePortResultOutput) BandwidthInGbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) *int { return v.BandwidthInGbps }).(pulumi.IntPtrOutput)
}

func (o LookupExpressRoutePortResultOutput) Circuits() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) []SubResourceResponse { return v.Circuits }).(SubResourceResponseArrayOutput)
}

func (o LookupExpressRoutePortResultOutput) Encapsulation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) *string { return v.Encapsulation }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRoutePortResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupExpressRoutePortResultOutput) EtherType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) string { return v.EtherType }).(pulumi.StringOutput)
}

func (o LookupExpressRoutePortResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRoutePortResultOutput) Links() ExpressRouteLinkResponseArrayOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) []ExpressRouteLinkResponse { return v.Links }).(ExpressRouteLinkResponseArrayOutput)
}

func (o LookupExpressRoutePortResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRoutePortResultOutput) Mtu() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) string { return v.Mtu }).(pulumi.StringOutput)
}

func (o LookupExpressRoutePortResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupExpressRoutePortResultOutput) PeeringLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) *string { return v.PeeringLocation }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRoutePortResultOutput) ProvisionedBandwidthInGbps() pulumi.Float64Output {
	return o.ApplyT(func(v LookupExpressRoutePortResult) float64 { return v.ProvisionedBandwidthInGbps }).(pulumi.Float64Output)
}

func (o LookupExpressRoutePortResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupExpressRoutePortResultOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRoutePortResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupExpressRoutePortResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExpressRoutePortResultOutput{})
}
