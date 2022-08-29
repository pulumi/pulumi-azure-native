


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSiteVNETConnectionSlot(ctx *pulumi.Context, args *LookupSiteVNETConnectionSlotArgs, opts ...pulumi.InvokeOption) (*LookupSiteVNETConnectionSlotResult, error) {
	var rv LookupSiteVNETConnectionSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteVNETConnectionSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteVNETConnectionSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
	VnetName          string `pulumi:"vnetName"`
}


type LookupSiteVNETConnectionSlotResult struct {
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

func LookupSiteVNETConnectionSlotOutput(ctx *pulumi.Context, args LookupSiteVNETConnectionSlotOutputArgs, opts ...pulumi.InvokeOption) LookupSiteVNETConnectionSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSiteVNETConnectionSlotResult, error) {
			args := v.(LookupSiteVNETConnectionSlotArgs)
			r, err := LookupSiteVNETConnectionSlot(ctx, &args, opts...)
			var s LookupSiteVNETConnectionSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSiteVNETConnectionSlotResultOutput)
}

type LookupSiteVNETConnectionSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
	VnetName          pulumi.StringInput `pulumi:"vnetName"`
}

func (LookupSiteVNETConnectionSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteVNETConnectionSlotArgs)(nil)).Elem()
}


type LookupSiteVNETConnectionSlotResultOutput struct{ *pulumi.OutputState }

func (LookupSiteVNETConnectionSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteVNETConnectionSlotResult)(nil)).Elem()
}

func (o LookupSiteVNETConnectionSlotResultOutput) ToLookupSiteVNETConnectionSlotResultOutput() LookupSiteVNETConnectionSlotResultOutput {
	return o
}

func (o LookupSiteVNETConnectionSlotResultOutput) ToLookupSiteVNETConnectionSlotResultOutputWithContext(ctx context.Context) LookupSiteVNETConnectionSlotResultOutput {
	return o
}

func (o LookupSiteVNETConnectionSlotResultOutput) CertBlob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionSlotResult) *string { return v.CertBlob }).(pulumi.StringPtrOutput)
}

func (o LookupSiteVNETConnectionSlotResultOutput) CertThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionSlotResult) *string { return v.CertThumbprint }).(pulumi.StringPtrOutput)
}

func (o LookupSiteVNETConnectionSlotResultOutput) DnsServers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionSlotResult) *string { return v.DnsServers }).(pulumi.StringPtrOutput)
}

func (o LookupSiteVNETConnectionSlotResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionSlotResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSiteVNETConnectionSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupSiteVNETConnectionSlotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionSlotResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSiteVNETConnectionSlotResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionSlotResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupSiteVNETConnectionSlotResultOutput) ResyncRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionSlotResult) *bool { return v.ResyncRequired }).(pulumi.BoolPtrOutput)
}

func (o LookupSiteVNETConnectionSlotResultOutput) Routes() VnetRouteResponseArrayOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionSlotResult) []VnetRouteResponse { return v.Routes }).(VnetRouteResponseArrayOutput)
}

func (o LookupSiteVNETConnectionSlotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionSlotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSiteVNETConnectionSlotResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionSlotResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupSiteVNETConnectionSlotResultOutput) VnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteVNETConnectionSlotResult) *string { return v.VnetResourceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSiteVNETConnectionSlotResultOutput{})
}
