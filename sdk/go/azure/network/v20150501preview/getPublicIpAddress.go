


package v20150501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupPublicIpAddress(ctx *pulumi.Context, args *LookupPublicIpAddressArgs, opts ...pulumi.InvokeOption) (*LookupPublicIpAddressResult, error) {
	var rv LookupPublicIpAddressResult
	err := ctx.Invoke("azure-native:network/v20150501preview:getPublicIpAddress", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPublicIpAddressArgs struct {
	PublicIpAddressName string `pulumi:"publicIpAddressName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupPublicIpAddressResult struct {
	DnsSettings              *PublicIpAddressDnsSettingsResponse `pulumi:"dnsSettings"`
	Etag                     *string                             `pulumi:"etag"`
	Id                       string                              `pulumi:"id"`
	IdleTimeoutInMinutes     *int                                `pulumi:"idleTimeoutInMinutes"`
	IpAddress                *string                             `pulumi:"ipAddress"`
	IpConfiguration          *SubResourceResponse                `pulumi:"ipConfiguration"`
	Location                 string                              `pulumi:"location"`
	Name                     string                              `pulumi:"name"`
	ProvisioningState        *string                             `pulumi:"provisioningState"`
	PublicIPAllocationMethod string                              `pulumi:"publicIPAllocationMethod"`
	ResourceGuid             *string                             `pulumi:"resourceGuid"`
	Tags                     map[string]string                   `pulumi:"tags"`
	Type                     string                              `pulumi:"type"`
}

func LookupPublicIpAddressOutput(ctx *pulumi.Context, args LookupPublicIpAddressOutputArgs, opts ...pulumi.InvokeOption) LookupPublicIpAddressResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPublicIpAddressResult, error) {
			args := v.(LookupPublicIpAddressArgs)
			r, err := LookupPublicIpAddress(ctx, &args, opts...)
			var s LookupPublicIpAddressResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPublicIpAddressResultOutput)
}

type LookupPublicIpAddressOutputArgs struct {
	PublicIpAddressName pulumi.StringInput `pulumi:"publicIpAddressName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPublicIpAddressOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicIpAddressArgs)(nil)).Elem()
}


type LookupPublicIpAddressResultOutput struct{ *pulumi.OutputState }

func (LookupPublicIpAddressResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicIpAddressResult)(nil)).Elem()
}

func (o LookupPublicIpAddressResultOutput) ToLookupPublicIpAddressResultOutput() LookupPublicIpAddressResultOutput {
	return o
}

func (o LookupPublicIpAddressResultOutput) ToLookupPublicIpAddressResultOutputWithContext(ctx context.Context) LookupPublicIpAddressResultOutput {
	return o
}

func (o LookupPublicIpAddressResultOutput) DnsSettings() PublicIpAddressDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIpAddressResult) *PublicIpAddressDnsSettingsResponse { return v.DnsSettings }).(PublicIpAddressDnsSettingsResponsePtrOutput)
}

func (o LookupPublicIpAddressResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIpAddressResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupPublicIpAddressResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIpAddressResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPublicIpAddressResultOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPublicIpAddressResult) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o LookupPublicIpAddressResultOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIpAddressResult) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o LookupPublicIpAddressResultOutput) IpConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIpAddressResult) *SubResourceResponse { return v.IpConfiguration }).(SubResourceResponsePtrOutput)
}

func (o LookupPublicIpAddressResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIpAddressResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPublicIpAddressResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIpAddressResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPublicIpAddressResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIpAddressResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupPublicIpAddressResultOutput) PublicIPAllocationMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIpAddressResult) string { return v.PublicIPAllocationMethod }).(pulumi.StringOutput)
}

func (o LookupPublicIpAddressResultOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIpAddressResult) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o LookupPublicIpAddressResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPublicIpAddressResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPublicIpAddressResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIpAddressResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPublicIpAddressResultOutput{})
}
