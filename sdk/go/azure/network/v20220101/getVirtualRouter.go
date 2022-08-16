


package v20220101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualRouter(ctx *pulumi.Context, args *LookupVirtualRouterArgs, opts ...pulumi.InvokeOption) (*LookupVirtualRouterResult, error) {
	var rv LookupVirtualRouterResult
	err := ctx.Invoke("azure-native:network/v20220101:getVirtualRouter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualRouterArgs struct {
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	VirtualRouterName string  `pulumi:"virtualRouterName"`
}


type LookupVirtualRouterResult struct {
	Etag              string                `pulumi:"etag"`
	HostedGateway     *SubResourceResponse  `pulumi:"hostedGateway"`
	HostedSubnet      *SubResourceResponse  `pulumi:"hostedSubnet"`
	Id                *string               `pulumi:"id"`
	Location          *string               `pulumi:"location"`
	Name              string                `pulumi:"name"`
	Peerings          []SubResourceResponse `pulumi:"peerings"`
	ProvisioningState string                `pulumi:"provisioningState"`
	Tags              map[string]string     `pulumi:"tags"`
	Type              string                `pulumi:"type"`
	VirtualRouterAsn  *float64              `pulumi:"virtualRouterAsn"`
	VirtualRouterIps  []string              `pulumi:"virtualRouterIps"`
}

func LookupVirtualRouterOutput(ctx *pulumi.Context, args LookupVirtualRouterOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualRouterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualRouterResult, error) {
			args := v.(LookupVirtualRouterArgs)
			r, err := LookupVirtualRouter(ctx, &args, opts...)
			var s LookupVirtualRouterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualRouterResultOutput)
}

type LookupVirtualRouterOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	VirtualRouterName pulumi.StringInput    `pulumi:"virtualRouterName"`
}

func (LookupVirtualRouterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualRouterArgs)(nil)).Elem()
}


type LookupVirtualRouterResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualRouterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualRouterResult)(nil)).Elem()
}

func (o LookupVirtualRouterResultOutput) ToLookupVirtualRouterResultOutput() LookupVirtualRouterResultOutput {
	return o
}

func (o LookupVirtualRouterResultOutput) ToLookupVirtualRouterResultOutputWithContext(ctx context.Context) LookupVirtualRouterResultOutput {
	return o
}

func (o LookupVirtualRouterResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupVirtualRouterResultOutput) HostedGateway() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) *SubResourceResponse { return v.HostedGateway }).(SubResourceResponsePtrOutput)
}

func (o LookupVirtualRouterResultOutput) HostedSubnet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) *SubResourceResponse { return v.HostedSubnet }).(SubResourceResponsePtrOutput)
}

func (o LookupVirtualRouterResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualRouterResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualRouterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualRouterResultOutput) Peerings() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) []SubResourceResponse { return v.Peerings }).(SubResourceResponseArrayOutput)
}

func (o LookupVirtualRouterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualRouterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualRouterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualRouterResultOutput) VirtualRouterAsn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) *float64 { return v.VirtualRouterAsn }).(pulumi.Float64PtrOutput)
}

func (o LookupVirtualRouterResultOutput) VirtualRouterIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) []string { return v.VirtualRouterIps }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualRouterResultOutput{})
}
