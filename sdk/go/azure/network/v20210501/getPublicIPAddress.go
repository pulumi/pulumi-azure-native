


package v20210501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPublicIPAddress(ctx *pulumi.Context, args *LookupPublicIPAddressArgs, opts ...pulumi.InvokeOption) (*LookupPublicIPAddressResult, error) {
	var rv LookupPublicIPAddressResult
	err := ctx.Invoke("azure-native:network/v20210501:getPublicIPAddress", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPublicIPAddressArgs struct {
	Expand              *string `pulumi:"expand"`
	PublicIpAddressName string  `pulumi:"publicIpAddressName"`
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
}


type LookupPublicIPAddressResult struct {
	DdosSettings             *DdosSettingsResponse               `pulumi:"ddosSettings"`
	DeleteOption             *string                             `pulumi:"deleteOption"`
	DnsSettings              *PublicIPAddressDnsSettingsResponse `pulumi:"dnsSettings"`
	Etag                     string                              `pulumi:"etag"`
	ExtendedLocation         *ExtendedLocationResponse           `pulumi:"extendedLocation"`
	Id                       *string                             `pulumi:"id"`
	IdleTimeoutInMinutes     *int                                `pulumi:"idleTimeoutInMinutes"`
	IpAddress                *string                             `pulumi:"ipAddress"`
	IpConfiguration          IPConfigurationResponse             `pulumi:"ipConfiguration"`
	IpTags                   []IpTagResponse                     `pulumi:"ipTags"`
	LinkedPublicIPAddress    *PublicIPAddressResponse            `pulumi:"linkedPublicIPAddress"`
	Location                 *string                             `pulumi:"location"`
	MigrationPhase           *string                             `pulumi:"migrationPhase"`
	Name                     string                              `pulumi:"name"`
	NatGateway               *NatGatewayResponse                 `pulumi:"natGateway"`
	ProvisioningState        string                              `pulumi:"provisioningState"`
	PublicIPAddressVersion   *string                             `pulumi:"publicIPAddressVersion"`
	PublicIPAllocationMethod *string                             `pulumi:"publicIPAllocationMethod"`
	PublicIPPrefix           *SubResourceResponse                `pulumi:"publicIPPrefix"`
	ResourceGuid             string                              `pulumi:"resourceGuid"`
	ServicePublicIPAddress   *PublicIPAddressResponse            `pulumi:"servicePublicIPAddress"`
	Sku                      *PublicIPAddressSkuResponse         `pulumi:"sku"`
	Tags                     map[string]string                   `pulumi:"tags"`
	Type                     string                              `pulumi:"type"`
	Zones                    []string                            `pulumi:"zones"`
}


func (val *LookupPublicIPAddressResult) Defaults() *LookupPublicIPAddressResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.IpConfiguration = *tmp.IpConfiguration.Defaults()

	tmp.LinkedPublicIPAddress = tmp.LinkedPublicIPAddress.Defaults()

	tmp.ServicePublicIPAddress = tmp.ServicePublicIPAddress.Defaults()

	return &tmp
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

func (o LookupPublicIPAddressResultOutput) DdosSettings() DdosSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *DdosSettingsResponse { return v.DdosSettings }).(DdosSettingsResponsePtrOutput)
}

func (o LookupPublicIPAddressResultOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *string { return v.DeleteOption }).(pulumi.StringPtrOutput)
}

func (o LookupPublicIPAddressResultOutput) DnsSettings() PublicIPAddressDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *PublicIPAddressDnsSettingsResponse { return v.DnsSettings }).(PublicIPAddressDnsSettingsResponsePtrOutput)
}

func (o LookupPublicIPAddressResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupPublicIPAddressResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
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

func (o LookupPublicIPAddressResultOutput) IpTags() IpTagResponseArrayOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) []IpTagResponse { return v.IpTags }).(IpTagResponseArrayOutput)
}

func (o LookupPublicIPAddressResultOutput) LinkedPublicIPAddress() PublicIPAddressResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *PublicIPAddressResponse { return v.LinkedPublicIPAddress }).(PublicIPAddressResponsePtrOutput)
}

func (o LookupPublicIPAddressResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupPublicIPAddressResultOutput) MigrationPhase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *string { return v.MigrationPhase }).(pulumi.StringPtrOutput)
}

func (o LookupPublicIPAddressResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPublicIPAddressResultOutput) NatGateway() NatGatewayResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *NatGatewayResponse { return v.NatGateway }).(NatGatewayResponsePtrOutput)
}

func (o LookupPublicIPAddressResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPublicIPAddressResultOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *string { return v.PublicIPAddressVersion }).(pulumi.StringPtrOutput)
}

func (o LookupPublicIPAddressResultOutput) PublicIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *string { return v.PublicIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o LookupPublicIPAddressResultOutput) PublicIPPrefix() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *SubResourceResponse { return v.PublicIPPrefix }).(SubResourceResponsePtrOutput)
}

func (o LookupPublicIPAddressResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupPublicIPAddressResultOutput) ServicePublicIPAddress() PublicIPAddressResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *PublicIPAddressResponse { return v.ServicePublicIPAddress }).(PublicIPAddressResponsePtrOutput)
}

func (o LookupPublicIPAddressResultOutput) Sku() PublicIPAddressSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) *PublicIPAddressSkuResponse { return v.Sku }).(PublicIPAddressSkuResponsePtrOutput)
}

func (o LookupPublicIPAddressResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPublicIPAddressResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupPublicIPAddressResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPublicIPAddressResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPublicIPAddressResultOutput{})
}
