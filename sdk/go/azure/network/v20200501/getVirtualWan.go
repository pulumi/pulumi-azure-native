


package v20200501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualWan(ctx *pulumi.Context, args *LookupVirtualWanArgs, opts ...pulumi.InvokeOption) (*LookupVirtualWanResult, error) {
	var rv LookupVirtualWanResult
	err := ctx.Invoke("azure-native:network/v20200501:getVirtualWan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualWanArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VirtualWANName    string `pulumi:"virtualWANName"`
}


type LookupVirtualWanResult struct {
	AllowBranchToBranchTraffic     *bool                 `pulumi:"allowBranchToBranchTraffic"`
	AllowVnetToVnetTraffic         *bool                 `pulumi:"allowVnetToVnetTraffic"`
	DisableVpnEncryption           *bool                 `pulumi:"disableVpnEncryption"`
	Etag                           string                `pulumi:"etag"`
	Id                             *string               `pulumi:"id"`
	Location                       string                `pulumi:"location"`
	Name                           string                `pulumi:"name"`
	Office365LocalBreakoutCategory string                `pulumi:"office365LocalBreakoutCategory"`
	ProvisioningState              string                `pulumi:"provisioningState"`
	Tags                           map[string]string     `pulumi:"tags"`
	Type                           string                `pulumi:"type"`
	VirtualHubs                    []SubResourceResponse `pulumi:"virtualHubs"`
	VpnSites                       []SubResourceResponse `pulumi:"vpnSites"`
}

func LookupVirtualWanOutput(ctx *pulumi.Context, args LookupVirtualWanOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualWanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualWanResult, error) {
			args := v.(LookupVirtualWanArgs)
			r, err := LookupVirtualWan(ctx, &args, opts...)
			var s LookupVirtualWanResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualWanResultOutput)
}

type LookupVirtualWanOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualWANName    pulumi.StringInput `pulumi:"virtualWANName"`
}

func (LookupVirtualWanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualWanArgs)(nil)).Elem()
}


type LookupVirtualWanResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualWanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualWanResult)(nil)).Elem()
}

func (o LookupVirtualWanResultOutput) ToLookupVirtualWanResultOutput() LookupVirtualWanResultOutput {
	return o
}

func (o LookupVirtualWanResultOutput) ToLookupVirtualWanResultOutputWithContext(ctx context.Context) LookupVirtualWanResultOutput {
	return o
}

func (o LookupVirtualWanResultOutput) AllowBranchToBranchTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) *bool { return v.AllowBranchToBranchTraffic }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualWanResultOutput) AllowVnetToVnetTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) *bool { return v.AllowVnetToVnetTraffic }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualWanResultOutput) DisableVpnEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) *bool { return v.DisableVpnEncryption }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualWanResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupVirtualWanResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualWanResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVirtualWanResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualWanResultOutput) Office365LocalBreakoutCategory() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) string { return v.Office365LocalBreakoutCategory }).(pulumi.StringOutput)
}

func (o LookupVirtualWanResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualWanResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualWanResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualWanResultOutput) VirtualHubs() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) []SubResourceResponse { return v.VirtualHubs }).(SubResourceResponseArrayOutput)
}

func (o LookupVirtualWanResultOutput) VpnSites() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualWanResult) []SubResourceResponse { return v.VpnSites }).(SubResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualWanResultOutput{})
}
