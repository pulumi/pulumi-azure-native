


package v20180701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualWAN(ctx *pulumi.Context, args *LookupVirtualWANArgs, opts ...pulumi.InvokeOption) (*LookupVirtualWANResult, error) {
	var rv LookupVirtualWANResult
	err := ctx.Invoke("azure-native:network/v20180701:getVirtualWAN", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualWANArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VirtualWANName    string `pulumi:"virtualWANName"`
}


type LookupVirtualWANResult struct {
	DisableVpnEncryption *bool                 `pulumi:"disableVpnEncryption"`
	Etag                 string                `pulumi:"etag"`
	Id                   *string               `pulumi:"id"`
	Location             string                `pulumi:"location"`
	Name                 string                `pulumi:"name"`
	ProvisioningState    string                `pulumi:"provisioningState"`
	Tags                 map[string]string     `pulumi:"tags"`
	Type                 string                `pulumi:"type"`
	VirtualHubs          []SubResourceResponse `pulumi:"virtualHubs"`
	VpnSites             []SubResourceResponse `pulumi:"vpnSites"`
}

func LookupVirtualWANOutput(ctx *pulumi.Context, args LookupVirtualWANOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualWANResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualWANResult, error) {
			args := v.(LookupVirtualWANArgs)
			r, err := LookupVirtualWAN(ctx, &args, opts...)
			var s LookupVirtualWANResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualWANResultOutput)
}

type LookupVirtualWANOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualWANName    pulumi.StringInput `pulumi:"virtualWANName"`
}

func (LookupVirtualWANOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualWANArgs)(nil)).Elem()
}


type LookupVirtualWANResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualWANResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualWANResult)(nil)).Elem()
}

func (o LookupVirtualWANResultOutput) ToLookupVirtualWANResultOutput() LookupVirtualWANResultOutput {
	return o
}

func (o LookupVirtualWANResultOutput) ToLookupVirtualWANResultOutputWithContext(ctx context.Context) LookupVirtualWANResultOutput {
	return o
}

func (o LookupVirtualWANResultOutput) DisableVpnEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualWANResult) *bool { return v.DisableVpnEncryption }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualWANResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWANResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupVirtualWANResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualWANResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualWANResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWANResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVirtualWANResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWANResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualWANResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWANResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualWANResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualWANResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualWANResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWANResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualWANResultOutput) VirtualHubs() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualWANResult) []SubResourceResponse { return v.VirtualHubs }).(SubResourceResponseArrayOutput)
}

func (o LookupVirtualWANResultOutput) VpnSites() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualWANResult) []SubResourceResponse { return v.VpnSites }).(SubResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualWANResultOutput{})
}
