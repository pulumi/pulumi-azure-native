


package v20220701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOutboundEndpoint(ctx *pulumi.Context, args *LookupOutboundEndpointArgs, opts ...pulumi.InvokeOption) (*LookupOutboundEndpointResult, error) {
	var rv LookupOutboundEndpointResult
	err := ctx.Invoke("azure-native:network/v20220701:getOutboundEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOutboundEndpointArgs struct {
	DnsResolverName      string `pulumi:"dnsResolverName"`
	OutboundEndpointName string `pulumi:"outboundEndpointName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupOutboundEndpointResult struct {
	Etag              string              `pulumi:"etag"`
	Id                string              `pulumi:"id"`
	Location          string              `pulumi:"location"`
	Name              string              `pulumi:"name"`
	ProvisioningState string              `pulumi:"provisioningState"`
	ResourceGuid      string              `pulumi:"resourceGuid"`
	Subnet            SubResourceResponse `pulumi:"subnet"`
	SystemData        SystemDataResponse  `pulumi:"systemData"`
	Tags              map[string]string   `pulumi:"tags"`
	Type              string              `pulumi:"type"`
}

func LookupOutboundEndpointOutput(ctx *pulumi.Context, args LookupOutboundEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupOutboundEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOutboundEndpointResult, error) {
			args := v.(LookupOutboundEndpointArgs)
			r, err := LookupOutboundEndpoint(ctx, &args, opts...)
			var s LookupOutboundEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOutboundEndpointResultOutput)
}

type LookupOutboundEndpointOutputArgs struct {
	DnsResolverName      pulumi.StringInput `pulumi:"dnsResolverName"`
	OutboundEndpointName pulumi.StringInput `pulumi:"outboundEndpointName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupOutboundEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOutboundEndpointArgs)(nil)).Elem()
}


type LookupOutboundEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupOutboundEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOutboundEndpointResult)(nil)).Elem()
}

func (o LookupOutboundEndpointResultOutput) ToLookupOutboundEndpointResultOutput() LookupOutboundEndpointResultOutput {
	return o
}

func (o LookupOutboundEndpointResultOutput) ToLookupOutboundEndpointResultOutputWithContext(ctx context.Context) LookupOutboundEndpointResultOutput {
	return o
}

func (o LookupOutboundEndpointResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutboundEndpointResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupOutboundEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutboundEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOutboundEndpointResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutboundEndpointResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupOutboundEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutboundEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOutboundEndpointResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutboundEndpointResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupOutboundEndpointResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutboundEndpointResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupOutboundEndpointResultOutput) Subnet() SubResourceResponseOutput {
	return o.ApplyT(func(v LookupOutboundEndpointResult) SubResourceResponse { return v.Subnet }).(SubResourceResponseOutput)
}

func (o LookupOutboundEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOutboundEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupOutboundEndpointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOutboundEndpointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupOutboundEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutboundEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOutboundEndpointResultOutput{})
}
