


package v20220701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupInboundEndpoint(ctx *pulumi.Context, args *LookupInboundEndpointArgs, opts ...pulumi.InvokeOption) (*LookupInboundEndpointResult, error) {
	var rv LookupInboundEndpointResult
	err := ctx.Invoke("azure-native:network/v20220701:getInboundEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInboundEndpointArgs struct {
	DnsResolverName     string `pulumi:"dnsResolverName"`
	InboundEndpointName string `pulumi:"inboundEndpointName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupInboundEndpointResult struct {
	Etag              string                                   `pulumi:"etag"`
	Id                string                                   `pulumi:"id"`
	IpConfigurations  []InboundEndpointIPConfigurationResponse `pulumi:"ipConfigurations"`
	Location          string                                   `pulumi:"location"`
	Name              string                                   `pulumi:"name"`
	ProvisioningState string                                   `pulumi:"provisioningState"`
	ResourceGuid      string                                   `pulumi:"resourceGuid"`
	SystemData        SystemDataResponse                       `pulumi:"systemData"`
	Tags              map[string]string                        `pulumi:"tags"`
	Type              string                                   `pulumi:"type"`
}

func LookupInboundEndpointOutput(ctx *pulumi.Context, args LookupInboundEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupInboundEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInboundEndpointResult, error) {
			args := v.(LookupInboundEndpointArgs)
			r, err := LookupInboundEndpoint(ctx, &args, opts...)
			var s LookupInboundEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInboundEndpointResultOutput)
}

type LookupInboundEndpointOutputArgs struct {
	DnsResolverName     pulumi.StringInput `pulumi:"dnsResolverName"`
	InboundEndpointName pulumi.StringInput `pulumi:"inboundEndpointName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupInboundEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInboundEndpointArgs)(nil)).Elem()
}


type LookupInboundEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupInboundEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInboundEndpointResult)(nil)).Elem()
}

func (o LookupInboundEndpointResultOutput) ToLookupInboundEndpointResultOutput() LookupInboundEndpointResultOutput {
	return o
}

func (o LookupInboundEndpointResultOutput) ToLookupInboundEndpointResultOutputWithContext(ctx context.Context) LookupInboundEndpointResultOutput {
	return o
}

func (o LookupInboundEndpointResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundEndpointResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupInboundEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupInboundEndpointResultOutput) IpConfigurations() InboundEndpointIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupInboundEndpointResult) []InboundEndpointIPConfigurationResponse {
		return v.IpConfigurations
	}).(InboundEndpointIPConfigurationResponseArrayOutput)
}

func (o LookupInboundEndpointResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundEndpointResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupInboundEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupInboundEndpointResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundEndpointResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupInboundEndpointResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundEndpointResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupInboundEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupInboundEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupInboundEndpointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInboundEndpointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupInboundEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInboundEndpointResultOutput{})
}
