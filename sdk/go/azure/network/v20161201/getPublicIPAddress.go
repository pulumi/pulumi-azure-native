


package v20161201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupPublicIPAddress(ctx *pulumi.Context, args *LookupPublicIPAddressArgs, opts ...pulumi.InvokeOption) (*LookupPublicIPAddressResult, error) {
	var rv LookupPublicIPAddressResult
	err := ctx.Invoke("azure-native:network/v20161201:getPublicIPAddress", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPublicIPAddressArgs struct {
	Expand              *string `pulumi:"expand"`
	PublicIpAddressName string  `pulumi:"publicIpAddressName"`
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
}


type LookupPublicIPAddressResult struct {
	DnsSettings              *PublicIPAddressDnsSettingsResponse `pulumi:"dnsSettings"`
	Etag                     *string                             `pulumi:"etag"`
	Id                       *string                             `pulumi:"id"`
	IdleTimeoutInMinutes     *int                                `pulumi:"idleTimeoutInMinutes"`
	IpAddress                *string                             `pulumi:"ipAddress"`
	IpConfiguration          IPConfigurationResponse             `pulumi:"ipConfiguration"`
	Location                 *string                             `pulumi:"location"`
	Name                     string                              `pulumi:"name"`
	ProvisioningState        *string                             `pulumi:"provisioningState"`
	PublicIPAddressVersion   *string                             `pulumi:"publicIPAddressVersion"`
	PublicIPAllocationMethod *string                             `pulumi:"publicIPAllocationMethod"`
	ResourceGuid             *string                             `pulumi:"resourceGuid"`
	Tags                     map[string]string                   `pulumi:"tags"`
	Type                     string                              `pulumi:"type"`
}

func LookupPublicIPAddressOutput(ctx *pulumi.Context, args LookupPublicIPAddressOutputArgs, opts ...pulumi.InvokeOption) LookupPublicIPAddressResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPublicIPAddressResult, error) {
			args := v.(LookupPublicIPAddressArgs)
			r, err := LookupPublicIPAddress(ctx, &args, opts...)
			var s LookupPublicIPAddressResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPublicIPAddressResultOutput)
}

type LookupPublicIPAddressOutputArgs struct {
	Expand              pulumi.StringPtrInput `pulumi:"expand"`
	PublicIpAddressName pulumi.StringInput    `pulumi:"publicIpAddressName"`
	ResourceGroupName   pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupPublicIPAddressOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicIPAddressArgs)(nil)).Elem()
}


type LookupPublicIPAddressResultOutput struct{ *pulumi.OutputState }

func (LookupPublicIPAddressResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicIPAddressResult)(nil)).Elem()
}

func (o LookupPublicIPAddressResultOutput) ToLookupPublicIPAddressResultOutput() LookupPublicIPAddressResultOutput {
	return o
}

func (o LookupPublicIPAddressResultOutput) ToLookupPublicIPAddressResultOutputWithContext(ctx context.Context) LookupPublicIPAddressResultOutput {
	return o
}

func (o LookupPublicIPAddressResultOutput) DnsSettings() PublicIPAddressDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *PublicIPAddressDnsSettingsResponse { return v.DnsSettings }).(PublicIPAddressDnsSettingsResponsePtrOutput)
}

func (o LookupPublicIPAddressResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupPublicIPAddressResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupPublicIPAddressResultOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o LookupPublicIPAddressResultOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o LookupPublicIPAddressResultOutput) IpConfiguration() IPConfigurationResponseOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) IPConfigurationResponse { return v.IpConfiguration }).(IPConfigurationResponseOutput)
}

func (o LookupPublicIPAddressResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupPublicIPAddressResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPublicIPAddressResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupPublicIPAddressResultOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *string { return v.PublicIPAddressVersion }).(pulumi.StringPtrOutput)
}

func (o LookupPublicIPAddressResultOutput) PublicIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *string { return v.PublicIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o LookupPublicIPAddressResultOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o LookupPublicIPAddressResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPublicIPAddressResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPublicIPAddressResultOutput{})
}
