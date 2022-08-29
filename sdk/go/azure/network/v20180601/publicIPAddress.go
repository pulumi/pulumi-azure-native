


package v20180601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PublicIPAddress struct {
	pulumi.CustomResourceState

	DnsSettings              PublicIPAddressDnsSettingsResponsePtrOutput `pulumi:"dnsSettings"`
	Etag                     pulumi.StringPtrOutput                      `pulumi:"etag"`
	IdleTimeoutInMinutes     pulumi.IntPtrOutput                         `pulumi:"idleTimeoutInMinutes"`
	IpAddress                pulumi.StringPtrOutput                      `pulumi:"ipAddress"`
	IpConfiguration          IPConfigurationResponseOutput               `pulumi:"ipConfiguration"`
	IpTags                   IpTagResponseArrayOutput                    `pulumi:"ipTags"`
	Location                 pulumi.StringPtrOutput                      `pulumi:"location"`
	Name                     pulumi.StringOutput                         `pulumi:"name"`
	ProvisioningState        pulumi.StringPtrOutput                      `pulumi:"provisioningState"`
	PublicIPAddressVersion   pulumi.StringPtrOutput                      `pulumi:"publicIPAddressVersion"`
	PublicIPAllocationMethod pulumi.StringPtrOutput                      `pulumi:"publicIPAllocationMethod"`
	ResourceGuid             pulumi.StringPtrOutput                      `pulumi:"resourceGuid"`
	Sku                      PublicIPAddressSkuResponsePtrOutput         `pulumi:"sku"`
	Tags                     pulumi.StringMapOutput                      `pulumi:"tags"`
	Type                     pulumi.StringOutput                         `pulumi:"type"`
	Zones                    pulumi.StringArrayOutput                    `pulumi:"zones"`
}


func NewPublicIPAddress(ctx *pulumi.Context,
	name string, args *PublicIPAddressArgs, opts ...pulumi.ResourceOption) (*PublicIPAddress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150501preview:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:PublicIPAddress"),
		},
	})
	opts = append(opts, aliases)
	var resource PublicIPAddress
	err := ctx.RegisterResource("azure-native:network/v20180601:PublicIPAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPublicIPAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicIPAddressState, opts ...pulumi.ResourceOption) (*PublicIPAddress, error) {
	var resource PublicIPAddress
	err := ctx.ReadResource("azure-native:network/v20180601:PublicIPAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type publicIPAddressState struct {
}

type PublicIPAddressState struct {
}

func (PublicIPAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIPAddressState)(nil)).Elem()
}

type publicIPAddressArgs struct {
	DnsSettings              *PublicIPAddressDnsSettings `pulumi:"dnsSettings"`
	Id                       *string                     `pulumi:"id"`
	IdleTimeoutInMinutes     *int                        `pulumi:"idleTimeoutInMinutes"`
	IpAddress                *string                     `pulumi:"ipAddress"`
	IpTags                   []IpTag                     `pulumi:"ipTags"`
	Location                 *string                     `pulumi:"location"`
	ProvisioningState        *string                     `pulumi:"provisioningState"`
	PublicIPAddressVersion   *string                     `pulumi:"publicIPAddressVersion"`
	PublicIPAllocationMethod *string                     `pulumi:"publicIPAllocationMethod"`
	PublicIpAddressName      *string                     `pulumi:"publicIpAddressName"`
	ResourceGroupName        string                      `pulumi:"resourceGroupName"`
	ResourceGuid             *string                     `pulumi:"resourceGuid"`
	Sku                      *PublicIPAddressSku         `pulumi:"sku"`
	Tags                     map[string]string           `pulumi:"tags"`
	Zones                    []string                    `pulumi:"zones"`
}


type PublicIPAddressArgs struct {
	DnsSettings              PublicIPAddressDnsSettingsPtrInput
	Id                       pulumi.StringPtrInput
	IdleTimeoutInMinutes     pulumi.IntPtrInput
	IpAddress                pulumi.StringPtrInput
	IpTags                   IpTagArrayInput
	Location                 pulumi.StringPtrInput
	ProvisioningState        pulumi.StringPtrInput
	PublicIPAddressVersion   pulumi.StringPtrInput
	PublicIPAllocationMethod pulumi.StringPtrInput
	PublicIpAddressName      pulumi.StringPtrInput
	ResourceGroupName        pulumi.StringInput
	ResourceGuid             pulumi.StringPtrInput
	Sku                      PublicIPAddressSkuPtrInput
	Tags                     pulumi.StringMapInput
	Zones                    pulumi.StringArrayInput
}

func (PublicIPAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIPAddressArgs)(nil)).Elem()
}

type PublicIPAddressInput interface {
	pulumi.Input

	ToPublicIPAddressOutput() PublicIPAddressOutput
	ToPublicIPAddressOutputWithContext(ctx context.Context) PublicIPAddressOutput
}

func (*PublicIPAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddress)(nil)).Elem()
}

func (i *PublicIPAddress) ToPublicIPAddressOutput() PublicIPAddressOutput {
	return i.ToPublicIPAddressOutputWithContext(context.Background())
}

func (i *PublicIPAddress) ToPublicIPAddressOutputWithContext(ctx context.Context) PublicIPAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressOutput)
}

type PublicIPAddressOutput struct{ *pulumi.OutputState }

func (PublicIPAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddress)(nil)).Elem()
}

func (o PublicIPAddressOutput) ToPublicIPAddressOutput() PublicIPAddressOutput {
	return o
}

func (o PublicIPAddressOutput) ToPublicIPAddressOutputWithContext(ctx context.Context) PublicIPAddressOutput {
	return o
}

func (o PublicIPAddressOutput) DnsSettings() PublicIPAddressDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) PublicIPAddressDnsSettingsResponsePtrOutput { return v.DnsSettings }).(PublicIPAddressDnsSettingsResponsePtrOutput)
}

func (o PublicIPAddressOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.IntPtrOutput { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o PublicIPAddressOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringPtrOutput { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressOutput) IpConfiguration() IPConfigurationResponseOutput {
	return o.ApplyT(func(v *PublicIPAddress) IPConfigurationResponseOutput { return v.IpConfiguration }).(IPConfigurationResponseOutput)
}

func (o PublicIPAddressOutput) IpTags() IpTagResponseArrayOutput {
	return o.ApplyT(func(v *PublicIPAddress) IpTagResponseArrayOutput { return v.IpTags }).(IpTagResponseArrayOutput)
}

func (o PublicIPAddressOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PublicIPAddressOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringPtrOutput { return v.PublicIPAddressVersion }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressOutput) PublicIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringPtrOutput { return v.PublicIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringPtrOutput { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressOutput) Sku() PublicIPAddressSkuResponsePtrOutput {
	return o.ApplyT(func(v *PublicIPAddress) PublicIPAddressSkuResponsePtrOutput { return v.Sku }).(PublicIPAddressSkuResponsePtrOutput)
}

func (o PublicIPAddressOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PublicIPAddressOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o PublicIPAddressOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PublicIPAddress) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(PublicIPAddressOutput{})
}
