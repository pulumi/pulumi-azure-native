


package v20201101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PublicIPPrefix struct {
	pulumi.CustomResourceState

	CustomIPPrefix                      SubResourceResponsePtrOutput                 `pulumi:"customIPPrefix"`
	Etag                                pulumi.StringOutput                          `pulumi:"etag"`
	ExtendedLocation                    ExtendedLocationResponsePtrOutput            `pulumi:"extendedLocation"`
	IpPrefix                            pulumi.StringOutput                          `pulumi:"ipPrefix"`
	IpTags                              IpTagResponseArrayOutput                     `pulumi:"ipTags"`
	LoadBalancerFrontendIpConfiguration SubResourceResponseOutput                    `pulumi:"loadBalancerFrontendIpConfiguration"`
	Location                            pulumi.StringPtrOutput                       `pulumi:"location"`
	Name                                pulumi.StringOutput                          `pulumi:"name"`
	NatGateway                          NatGatewayResponsePtrOutput                  `pulumi:"natGateway"`
	PrefixLength                        pulumi.IntPtrOutput                          `pulumi:"prefixLength"`
	ProvisioningState                   pulumi.StringOutput                          `pulumi:"provisioningState"`
	PublicIPAddressVersion              pulumi.StringPtrOutput                       `pulumi:"publicIPAddressVersion"`
	PublicIPAddresses                   ReferencedPublicIpAddressResponseArrayOutput `pulumi:"publicIPAddresses"`
	ResourceGuid                        pulumi.StringOutput                          `pulumi:"resourceGuid"`
	Sku                                 PublicIPPrefixSkuResponsePtrOutput           `pulumi:"sku"`
	Tags                                pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                                pulumi.StringOutput                          `pulumi:"type"`
	Zones                               pulumi.StringArrayOutput                     `pulumi:"zones"`
}


func NewPublicIPPrefix(ctx *pulumi.Context,
	name string, args *PublicIPPrefixArgs, opts ...pulumi.ResourceOption) (*PublicIPPrefix, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:PublicIPPrefix"),
		},
	})
	opts = append(opts, aliases)
	var resource PublicIPPrefix
	err := ctx.RegisterResource("azure-native:network/v20201101:PublicIPPrefix", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPublicIPPrefix(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicIPPrefixState, opts ...pulumi.ResourceOption) (*PublicIPPrefix, error) {
	var resource PublicIPPrefix
	err := ctx.ReadResource("azure-native:network/v20201101:PublicIPPrefix", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type publicIPPrefixState struct {
}

type PublicIPPrefixState struct {
}

func (PublicIPPrefixState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIPPrefixState)(nil)).Elem()
}

type publicIPPrefixArgs struct {
	CustomIPPrefix         *SubResource       `pulumi:"customIPPrefix"`
	ExtendedLocation       *ExtendedLocation  `pulumi:"extendedLocation"`
	Id                     *string            `pulumi:"id"`
	IpTags                 []IpTag            `pulumi:"ipTags"`
	Location               *string            `pulumi:"location"`
	NatGateway             *NatGatewayType    `pulumi:"natGateway"`
	PrefixLength           *int               `pulumi:"prefixLength"`
	PublicIPAddressVersion *string            `pulumi:"publicIPAddressVersion"`
	PublicIpPrefixName     *string            `pulumi:"publicIpPrefixName"`
	ResourceGroupName      string             `pulumi:"resourceGroupName"`
	Sku                    *PublicIPPrefixSku `pulumi:"sku"`
	Tags                   map[string]string  `pulumi:"tags"`
	Zones                  []string           `pulumi:"zones"`
}


type PublicIPPrefixArgs struct {
	CustomIPPrefix         SubResourcePtrInput
	ExtendedLocation       ExtendedLocationPtrInput
	Id                     pulumi.StringPtrInput
	IpTags                 IpTagArrayInput
	Location               pulumi.StringPtrInput
	NatGateway             NatGatewayTypePtrInput
	PrefixLength           pulumi.IntPtrInput
	PublicIPAddressVersion pulumi.StringPtrInput
	PublicIpPrefixName     pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	Sku                    PublicIPPrefixSkuPtrInput
	Tags                   pulumi.StringMapInput
	Zones                  pulumi.StringArrayInput
}

func (PublicIPPrefixArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIPPrefixArgs)(nil)).Elem()
}

type PublicIPPrefixInput interface {
	pulumi.Input

	ToPublicIPPrefixOutput() PublicIPPrefixOutput
	ToPublicIPPrefixOutputWithContext(ctx context.Context) PublicIPPrefixOutput
}

func (*PublicIPPrefix) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPPrefix)(nil)).Elem()
}

func (i *PublicIPPrefix) ToPublicIPPrefixOutput() PublicIPPrefixOutput {
	return i.ToPublicIPPrefixOutputWithContext(context.Background())
}

func (i *PublicIPPrefix) ToPublicIPPrefixOutputWithContext(ctx context.Context) PublicIPPrefixOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPPrefixOutput)
}

type PublicIPPrefixOutput struct{ *pulumi.OutputState }

func (PublicIPPrefixOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPPrefix)(nil)).Elem()
}

func (o PublicIPPrefixOutput) ToPublicIPPrefixOutput() PublicIPPrefixOutput {
	return o
}

func (o PublicIPPrefixOutput) ToPublicIPPrefixOutputWithContext(ctx context.Context) PublicIPPrefixOutput {
	return o
}

func (o PublicIPPrefixOutput) CustomIPPrefix() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *PublicIPPrefix) SubResourceResponsePtrOutput { return v.CustomIPPrefix }).(SubResourceResponsePtrOutput)
}

func (o PublicIPPrefixOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIPPrefix) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o PublicIPPrefixOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *PublicIPPrefix) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o PublicIPPrefixOutput) IpPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIPPrefix) pulumi.StringOutput { return v.IpPrefix }).(pulumi.StringOutput)
}

func (o PublicIPPrefixOutput) IpTags() IpTagResponseArrayOutput {
	return o.ApplyT(func(v *PublicIPPrefix) IpTagResponseArrayOutput { return v.IpTags }).(IpTagResponseArrayOutput)
}

func (o PublicIPPrefixOutput) LoadBalancerFrontendIpConfiguration() SubResourceResponseOutput {
	return o.ApplyT(func(v *PublicIPPrefix) SubResourceResponseOutput { return v.LoadBalancerFrontendIpConfiguration }).(SubResourceResponseOutput)
}

func (o PublicIPPrefixOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPPrefix) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o PublicIPPrefixOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIPPrefix) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PublicIPPrefixOutput) NatGateway() NatGatewayResponsePtrOutput {
	return o.ApplyT(func(v *PublicIPPrefix) NatGatewayResponsePtrOutput { return v.NatGateway }).(NatGatewayResponsePtrOutput)
}

func (o PublicIPPrefixOutput) PrefixLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PublicIPPrefix) pulumi.IntPtrOutput { return v.PrefixLength }).(pulumi.IntPtrOutput)
}

func (o PublicIPPrefixOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIPPrefix) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PublicIPPrefixOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPPrefix) pulumi.StringPtrOutput { return v.PublicIPAddressVersion }).(pulumi.StringPtrOutput)
}

func (o PublicIPPrefixOutput) PublicIPAddresses() ReferencedPublicIpAddressResponseArrayOutput {
	return o.ApplyT(func(v *PublicIPPrefix) ReferencedPublicIpAddressResponseArrayOutput { return v.PublicIPAddresses }).(ReferencedPublicIpAddressResponseArrayOutput)
}

func (o PublicIPPrefixOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIPPrefix) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o PublicIPPrefixOutput) Sku() PublicIPPrefixSkuResponsePtrOutput {
	return o.ApplyT(func(v *PublicIPPrefix) PublicIPPrefixSkuResponsePtrOutput { return v.Sku }).(PublicIPPrefixSkuResponsePtrOutput)
}

func (o PublicIPPrefixOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PublicIPPrefix) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PublicIPPrefixOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIPPrefix) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o PublicIPPrefixOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PublicIPPrefix) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(PublicIPPrefixOutput{})
}
