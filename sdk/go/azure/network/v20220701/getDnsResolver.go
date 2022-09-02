


package v20220701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDnsResolver(ctx *pulumi.Context, args *LookupDnsResolverArgs, opts ...pulumi.InvokeOption) (*LookupDnsResolverResult, error) {
	var rv LookupDnsResolverResult
	err := ctx.Invoke("azure-native:network/v20220701:getDnsResolver", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDnsResolverArgs struct {
	DnsResolverName   string `pulumi:"dnsResolverName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDnsResolverResult struct {
	DnsResolverState  string              `pulumi:"dnsResolverState"`
	Etag              string              `pulumi:"etag"`
	Id                string              `pulumi:"id"`
	Location          string              `pulumi:"location"`
	Name              string              `pulumi:"name"`
	ProvisioningState string              `pulumi:"provisioningState"`
	ResourceGuid      string              `pulumi:"resourceGuid"`
	SystemData        SystemDataResponse  `pulumi:"systemData"`
	Tags              map[string]string   `pulumi:"tags"`
	Type              string              `pulumi:"type"`
	VirtualNetwork    SubResourceResponse `pulumi:"virtualNetwork"`
}

func LookupDnsResolverOutput(ctx *pulumi.Context, args LookupDnsResolverOutputArgs, opts ...pulumi.InvokeOption) LookupDnsResolverResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDnsResolverResult, error) {
			args := v.(LookupDnsResolverArgs)
			r, err := LookupDnsResolver(ctx, &args, opts...)
			var s LookupDnsResolverResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDnsResolverResultOutput)
}

type LookupDnsResolverOutputArgs struct {
	DnsResolverName   pulumi.StringInput `pulumi:"dnsResolverName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDnsResolverOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDnsResolverArgs)(nil)).Elem()
}


type LookupDnsResolverResultOutput struct{ *pulumi.OutputState }

func (LookupDnsResolverResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDnsResolverResult)(nil)).Elem()
}

func (o LookupDnsResolverResultOutput) ToLookupDnsResolverResultOutput() LookupDnsResolverResultOutput {
	return o
}

func (o LookupDnsResolverResultOutput) ToLookupDnsResolverResultOutputWithContext(ctx context.Context) LookupDnsResolverResultOutput {
	return o
}

func (o LookupDnsResolverResultOutput) DnsResolverState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResolverResult) string { return v.DnsResolverState }).(pulumi.StringOutput)
}

func (o LookupDnsResolverResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResolverResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupDnsResolverResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResolverResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDnsResolverResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResolverResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDnsResolverResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResolverResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDnsResolverResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResolverResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDnsResolverResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResolverResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupDnsResolverResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDnsResolverResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDnsResolverResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDnsResolverResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDnsResolverResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResolverResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDnsResolverResultOutput) VirtualNetwork() SubResourceResponseOutput {
	return o.ApplyT(func(v LookupDnsResolverResult) SubResourceResponse { return v.VirtualNetwork }).(SubResourceResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDnsResolverResultOutput{})
}
