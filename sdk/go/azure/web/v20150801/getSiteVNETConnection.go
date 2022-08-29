


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSiteVNETConnection(ctx *pulumi.Context, args *LookupSiteVNETConnectionArgs, opts ...pulumi.InvokeOption) (*LookupSiteVNETConnectionResult, error) {
	var rv LookupSiteVNETConnectionResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteVNETConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteVNETConnectionArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VnetName          string `pulumi:"vnetName"`
}


type LookupSiteVNETConnectionResult struct {
	CertBlob       *string             `pulumi:"certBlob"`
	CertThumbprint *string             `pulumi:"certThumbprint"`
	DnsServers     *string             `pulumi:"dnsServers"`
	Id             *string             `pulumi:"id"`
	Kind           *string             `pulumi:"kind"`
	Location       string              `pulumi:"location"`
	Name           *string             `pulumi:"name"`
	ResyncRequired *bool               `pulumi:"resyncRequired"`
	Routes         []VnetRouteResponse `pulumi:"routes"`
	Tags           map[string]string   `pulumi:"tags"`
	Type           *string             `pulumi:"type"`
	VnetResourceId *string             `pulumi:"vnetResourceId"`
}

func LookupSiteVNETConnectionOutput(ctx *pulumi.Context, args LookupSiteVNETConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupSiteVNETConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSiteVNETConnectionResult, error) {
			args := v.(LookupSiteVNETConnectionArgs)
			r, err := LookupSiteVNETConnection(ctx, &args, opts...)
			var s LookupSiteVNETConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSiteVNETConnectionResultOutput)
}

type LookupSiteVNETConnectionOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VnetName          pulumi.StringInput `pulumi:"vnetName"`
}

func (LookupSiteVNETConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteVNETConnectionArgs)(nil)).Elem()
}


type LookupSiteVNETConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupSiteVNETConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteVNETConnectionResult)(nil)).Elem()
}

func (o LookupSiteVNETConnectionResultOutput) ToLookupSiteVNETConnectionResultOutput() LookupSiteVNETConnectionResultOutput {
	return o
}

func (o LookupSiteVNETConnectionResultOutput) ToLookupSiteVNETConnectionResultOutputWithContext(ctx context.Context) LookupSiteVNETConnectionResultOutput {
	return o
}

func (o LookupSiteVNETConnectionResultOutput) CertBlob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionResult) *string { return v.CertBlob }).(pulumi.StringPtrOutput)
}

func (o LookupSiteVNETConnectionResultOutput) CertThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionResult) *string { return v.CertThumbprint }).(pulumi.StringPtrOutput)
}

func (o LookupSiteVNETConnectionResultOutput) DnsServers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionResult) *string { return v.DnsServers }).(pulumi.StringPtrOutput)
}

func (o LookupSiteVNETConnectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSiteVNETConnectionResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupSiteVNETConnectionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSiteVNETConnectionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupSiteVNETConnectionResultOutput) ResyncRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionResult) *bool { return v.ResyncRequired }).(pulumi.BoolPtrOutput)
}

func (o LookupSiteVNETConnectionResultOutput) Routes() VnetRouteResponseArrayOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionResult) []VnetRouteResponse { return v.Routes }).(VnetRouteResponseArrayOutput)
}

func (o LookupSiteVNETConnectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSiteVNETConnectionResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupSiteVNETConnectionResultOutput) VnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionResult) *string { return v.VnetResourceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSiteVNETConnectionResultOutput{})
}
