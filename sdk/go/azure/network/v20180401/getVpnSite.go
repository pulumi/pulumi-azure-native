


package v20180401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupVpnSite(ctx *pulumi.Context, args *LookupVpnSiteArgs, opts ...pulumi.InvokeOption) (*LookupVpnSiteResult, error) {
	var rv LookupVpnSiteResult
	err := ctx.Invoke("azure-native:network/v20180401:getVpnSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpnSiteArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VpnSiteName       string `pulumi:"vpnSiteName"`
}


type LookupVpnSiteResult struct {
	AddressSpace      *AddressSpaceResponse     `pulumi:"addressSpace"`
	BgpProperties     *BgpSettingsResponse      `pulumi:"bgpProperties"`
	DeviceProperties  *DevicePropertiesResponse `pulumi:"deviceProperties"`
	Etag              string                    `pulumi:"etag"`
	Id                *string                   `pulumi:"id"`
	IpAddress         *string                   `pulumi:"ipAddress"`
	Location          string                    `pulumi:"location"`
	Name              string                    `pulumi:"name"`
	ProvisioningState string                    `pulumi:"provisioningState"`
	SiteKey           *string                   `pulumi:"siteKey"`
	Tags              map[string]string         `pulumi:"tags"`
	Type              string                    `pulumi:"type"`
	VirtualWAN        *SubResourceResponse      `pulumi:"virtualWAN"`
}

func LookupVpnSiteOutput(ctx *pulumi.Context, args LookupVpnSiteOutputArgs, opts ...pulumi.InvokeOption) LookupVpnSiteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpnSiteResult, error) {
			args := v.(LookupVpnSiteArgs)
			r, err := LookupVpnSite(ctx, &args, opts...)
			var s LookupVpnSiteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpnSiteResultOutput)
}

type LookupVpnSiteOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VpnSiteName       pulumi.StringInput `pulumi:"vpnSiteName"`
}

func (LookupVpnSiteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnSiteArgs)(nil)).Elem()
}


type LookupVpnSiteResultOutput struct{ *pulumi.OutputState }

func (LookupVpnSiteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnSiteResult)(nil)).Elem()
}

func (o LookupVpnSiteResultOutput) ToLookupVpnSiteResultOutput() LookupVpnSiteResultOutput {
	return o
}

func (o LookupVpnSiteResultOutput) ToLookupVpnSiteResultOutputWithContext(ctx context.Context) LookupVpnSiteResultOutput {
	return o
}

func (o LookupVpnSiteResultOutput) AddressSpace() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) *AddressSpaceResponse { return v.AddressSpace }).(AddressSpaceResponsePtrOutput)
}

func (o LookupVpnSiteResultOutput) BgpProperties() BgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) *BgpSettingsResponse { return v.BgpProperties }).(BgpSettingsResponsePtrOutput)
}

func (o LookupVpnSiteResultOutput) DeviceProperties() DevicePropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) *DevicePropertiesResponse { return v.DeviceProperties }).(DevicePropertiesResponsePtrOutput)
}

func (o LookupVpnSiteResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupVpnSiteResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVpnSiteResultOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o LookupVpnSiteResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVpnSiteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVpnSiteResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVpnSiteResultOutput) SiteKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) *string { return v.SiteKey }).(pulumi.StringPtrOutput)
}

func (o LookupVpnSiteResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVpnSiteResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVpnSiteResultOutput) VirtualWAN() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVpnSiteResult) *SubResourceResponse { return v.VirtualWAN }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpnSiteResultOutput{})
}
