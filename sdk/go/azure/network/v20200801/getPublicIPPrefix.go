


package v20200801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPublicIPPrefix(ctx *pulumi.Context, args *LookupPublicIPPrefixArgs, opts ...pulumi.InvokeOption) (*LookupPublicIPPrefixResult, error) {
	var rv LookupPublicIPPrefixResult
	err := ctx.Invoke("azure-native:network/v20200801:getPublicIPPrefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPublicIPPrefixArgs struct {
	Expand             *string `pulumi:"expand"`
	PublicIpPrefixName string  `pulumi:"publicIpPrefixName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
}


type LookupPublicIPPrefixResult struct {
	CustomIPPrefix                      *SubResourceResponse                `pulumi:"customIPPrefix"`
	Etag                                string                              `pulumi:"etag"`
	ExtendedLocation                    *ExtendedLocationResponse           `pulumi:"extendedLocation"`
	Id                                  *string                             `pulumi:"id"`
	IpPrefix                            string                              `pulumi:"ipPrefix"`
	IpTags                              []IpTagResponse                     `pulumi:"ipTags"`
	LoadBalancerFrontendIpConfiguration SubResourceResponse                 `pulumi:"loadBalancerFrontendIpConfiguration"`
	Location                            *string                             `pulumi:"location"`
	Name                                string                              `pulumi:"name"`
	NatGateway                          *NatGatewayResponse                 `pulumi:"natGateway"`
	PrefixLength                        *int                                `pulumi:"prefixLength"`
	ProvisioningState                   string                              `pulumi:"provisioningState"`
	PublicIPAddressVersion              *string                             `pulumi:"publicIPAddressVersion"`
	PublicIPAddresses                   []ReferencedPublicIpAddressResponse `pulumi:"publicIPAddresses"`
	ResourceGuid                        string                              `pulumi:"resourceGuid"`
	Sku                                 *PublicIPPrefixSkuResponse          `pulumi:"sku"`
	Tags                                map[string]string                   `pulumi:"tags"`
	Type                                string                              `pulumi:"type"`
	Zones                               []string                            `pulumi:"zones"`
}

func LookupPublicIPPrefixOutput(ctx *pulumi.Context, args LookupPublicIPPrefixOutputArgs, opts ...pulumi.InvokeOption) LookupPublicIPPrefixResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPublicIPPrefixResult, error) {
			args := v.(LookupPublicIPPrefixArgs)
			r, err := LookupPublicIPPrefix(ctx, &args, opts...)
			var s LookupPublicIPPrefixResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPublicIPPrefixResultOutput)
}

type LookupPublicIPPrefixOutputArgs struct {
	Expand             pulumi.StringPtrInput `pulumi:"expand"`
	PublicIpPrefixName pulumi.StringInput    `pulumi:"publicIpPrefixName"`
	ResourceGroupName  pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupPublicIPPrefixOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicIPPrefixArgs)(nil)).Elem()
}


type LookupPublicIPPrefixResultOutput struct{ *pulumi.OutputState }

func (LookupPublicIPPrefixResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicIPPrefixResult)(nil)).Elem()
}

func (o LookupPublicIPPrefixResultOutput) ToLookupPublicIPPrefixResultOutput() LookupPublicIPPrefixResultOutput {
	return o
}

func (o LookupPublicIPPrefixResultOutput) ToLookupPublicIPPrefixResultOutputWithContext(ctx context.Context) LookupPublicIPPrefixResultOutput {
	return o
}

func (o LookupPublicIPPrefixResultOutput) CustomIPPrefix() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) *SubResourceResponse { return v.CustomIPPrefix }).(SubResourceResponsePtrOutput)
}

func (o LookupPublicIPPrefixResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupPublicIPPrefixResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupPublicIPPrefixResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupPublicIPPrefixResultOutput) IpPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) string { return v.IpPrefix }).(pulumi.StringOutput)
}

func (o LookupPublicIPPrefixResultOutput) IpTags() IpTagResponseArrayOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) []IpTagResponse { return v.IpTags }).(IpTagResponseArrayOutput)
}

func (o LookupPublicIPPrefixResultOutput) LoadBalancerFrontendIpConfiguration() SubResourceResponseOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) SubResourceResponse { return v.LoadBalancerFrontendIpConfiguration }).(SubResourceResponseOutput)
}

func (o LookupPublicIPPrefixResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupPublicIPPrefixResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPublicIPPrefixResultOutput) NatGateway() NatGatewayResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) *NatGatewayResponse { return v.NatGateway }).(NatGatewayResponsePtrOutput)
}

func (o LookupPublicIPPrefixResultOutput) PrefixLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) *int { return v.PrefixLength }).(pulumi.IntPtrOutput)
}

func (o LookupPublicIPPrefixResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPublicIPPrefixResultOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) *string { return v.PublicIPAddressVersion }).(pulumi.StringPtrOutput)
}

func (o LookupPublicIPPrefixResultOutput) PublicIPAddresses() ReferencedPublicIpAddressResponseArrayOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) []ReferencedPublicIpAddressResponse { return v.PublicIPAddresses }).(ReferencedPublicIpAddressResponseArrayOutput)
}

func (o LookupPublicIPPrefixResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupPublicIPPrefixResultOutput) Sku() PublicIPPrefixSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) *PublicIPPrefixSkuResponse { return v.Sku }).(PublicIPPrefixSkuResponsePtrOutput)
}

func (o LookupPublicIPPrefixResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPublicIPPrefixResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupPublicIPPrefixResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPublicIPPrefixResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPublicIPPrefixResultOutput{})
}
