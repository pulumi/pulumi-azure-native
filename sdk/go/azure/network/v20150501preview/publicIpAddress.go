


package v20150501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type PublicIpAddress struct {
	pulumi.CustomResourceState

	DnsSettings              PublicIpAddressDnsSettingsResponsePtrOutput `pulumi:"dnsSettings"`
	Etag                     pulumi.StringPtrOutput                      `pulumi:"etag"`
	IdleTimeoutInMinutes     pulumi.IntPtrOutput                         `pulumi:"idleTimeoutInMinutes"`
	IpAddress                pulumi.StringPtrOutput                      `pulumi:"ipAddress"`
	IpConfiguration          SubResourceResponsePtrOutput                `pulumi:"ipConfiguration"`
	Location                 pulumi.StringOutput                         `pulumi:"location"`
	Name                     pulumi.StringOutput                         `pulumi:"name"`
	ProvisioningState        pulumi.StringPtrOutput                      `pulumi:"provisioningState"`
	PublicIPAllocationMethod pulumi.StringOutput                         `pulumi:"publicIPAllocationMethod"`
	ResourceGuid             pulumi.StringPtrOutput                      `pulumi:"resourceGuid"`
	Tags                     pulumi.StringMapOutput                      `pulumi:"tags"`
	Type                     pulumi.StringOutput                         `pulumi:"type"`
}


func NewPublicIpAddress(ctx *pulumi.Context,
	name string, args *PublicIpAddressArgs, opts ...pulumi.ResourceOption) (*PublicIpAddress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PublicIPAllocationMethod == nil {
		return nil, errors.New("invalid value for required argument 'PublicIPAllocationMethod'")
	}
	if args.PublicIpAddressName == nil {
		return nil, errors.New("invalid value for required argument 'PublicIpAddressName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:PublicIpAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:PublicIpAddress"),
		},
	})
	opts = append(opts, aliases)
	var resource PublicIpAddress
	err := ctx.RegisterResource("azure-native:network/v20150501preview:PublicIpAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPublicIpAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicIpAddressState, opts ...pulumi.ResourceOption) (*PublicIpAddress, error) {
	var resource PublicIpAddress
	err := ctx.ReadResource("azure-native:network/v20150501preview:PublicIpAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type publicIpAddressState struct {
}

type PublicIpAddressState struct {
}

func (PublicIpAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIpAddressState)(nil)).Elem()
}

type publicIpAddressArgs struct {
	DnsSettings              *PublicIpAddressDnsSettings `pulumi:"dnsSettings"`
	IdleTimeoutInMinutes     *int                        `pulumi:"idleTimeoutInMinutes"`
	IpAddress                *string                     `pulumi:"ipAddress"`
	IpConfiguration          *SubResource                `pulumi:"ipConfiguration"`
	Location                 *string                     `pulumi:"location"`
	ProvisioningState        *string                     `pulumi:"provisioningState"`
	PublicIPAllocationMethod string                      `pulumi:"publicIPAllocationMethod"`
	PublicIpAddressName      string                      `pulumi:"publicIpAddressName"`
	ResourceGroupName        string                      `pulumi:"resourceGroupName"`
	ResourceGuid             *string                     `pulumi:"resourceGuid"`
	Tags                     map[string]string           `pulumi:"tags"`
}


type PublicIpAddressArgs struct {
	DnsSettings              PublicIpAddressDnsSettingsPtrInput
	IdleTimeoutInMinutes     pulumi.IntPtrInput
	IpAddress                pulumi.StringPtrInput
	IpConfiguration          SubResourcePtrInput
	Location                 pulumi.StringPtrInput
	ProvisioningState        pulumi.StringPtrInput
	PublicIPAllocationMethod pulumi.StringInput
	PublicIpAddressName      pulumi.StringInput
	ResourceGroupName        pulumi.StringInput
	ResourceGuid             pulumi.StringPtrInput
	Tags                     pulumi.StringMapInput
}

func (PublicIpAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIpAddressArgs)(nil)).Elem()
}

type PublicIpAddressInput interface {
	pulumi.Input

	ToPublicIpAddressOutput() PublicIpAddressOutput
	ToPublicIpAddressOutputWithContext(ctx context.Context) PublicIpAddressOutput
}

func (*PublicIpAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIpAddress)(nil)).Elem()
}

func (i *PublicIpAddress) ToPublicIpAddressOutput() PublicIpAddressOutput {
	return i.ToPublicIpAddressOutputWithContext(context.Background())
}

func (i *PublicIpAddress) ToPublicIpAddressOutputWithContext(ctx context.Context) PublicIpAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpAddressOutput)
}

type PublicIpAddressOutput struct{ *pulumi.OutputState }

func (PublicIpAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIpAddress)(nil)).Elem()
}

func (o PublicIpAddressOutput) ToPublicIpAddressOutput() PublicIpAddressOutput {
	return o
}

func (o PublicIpAddressOutput) ToPublicIpAddressOutputWithContext(ctx context.Context) PublicIpAddressOutput {
	return o
}

func (o PublicIpAddressOutput) DnsSettings() PublicIpAddressDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *PublicIpAddress) PublicIpAddressDnsSettingsResponsePtrOutput { return v.DnsSettings }).(PublicIpAddressDnsSettingsResponsePtrOutput)
}

func (o PublicIpAddressOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIpAddress) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o PublicIpAddressOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PublicIpAddress) pulumi.IntPtrOutput { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o PublicIpAddressOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIpAddress) pulumi.StringPtrOutput { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o PublicIpAddressOutput) IpConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *PublicIpAddress) SubResourceResponsePtrOutput { return v.IpConfiguration }).(SubResourceResponsePtrOutput)
}

func (o PublicIpAddressOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIpAddress) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PublicIpAddressOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIpAddress) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PublicIpAddressOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIpAddress) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o PublicIpAddressOutput) PublicIPAllocationMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIpAddress) pulumi.StringOutput { return v.PublicIPAllocationMethod }).(pulumi.StringOutput)
}

func (o PublicIpAddressOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIpAddress) pulumi.StringPtrOutput { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o PublicIpAddressOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PublicIpAddress) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PublicIpAddressOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIpAddress) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PublicIpAddressOutput{})
}
