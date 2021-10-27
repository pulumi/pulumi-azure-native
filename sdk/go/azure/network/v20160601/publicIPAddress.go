


package v20160601

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
	Location                 pulumi.StringPtrOutput                      `pulumi:"location"`
	Name                     pulumi.StringOutput                         `pulumi:"name"`
	ProvisioningState        pulumi.StringPtrOutput                      `pulumi:"provisioningState"`
	PublicIPAddressVersion   pulumi.StringPtrOutput                      `pulumi:"publicIPAddressVersion"`
	PublicIPAllocationMethod pulumi.StringPtrOutput                      `pulumi:"publicIPAllocationMethod"`
	ResourceGuid             pulumi.StringPtrOutput                      `pulumi:"resourceGuid"`
	Tags                     pulumi.StringMapOutput                      `pulumi:"tags"`
	Type                     pulumi.StringOutput                         `pulumi:"type"`
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
			Type: pulumi.String("azure-nextgen:network/v20160601:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150501preview:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20150501preview:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20150615:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20160330:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20160901:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20161201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170301:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170601:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170801:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170901:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20171001:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20171101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180401:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180601:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180701:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180801:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181001:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190401:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190601:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200801:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20201101:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210201:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210301:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:PublicIPAddress"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210501:PublicIPAddress"),
		},
	})
	opts = append(opts, aliases)
	var resource PublicIPAddress
	err := ctx.RegisterResource("azure-native:network/v20160601:PublicIPAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPublicIPAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicIPAddressState, opts ...pulumi.ResourceOption) (*PublicIPAddress, error) {
	var resource PublicIPAddress
	err := ctx.ReadResource("azure-native:network/v20160601:PublicIPAddress", name, id, state, &resource, opts...)
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
	Etag                     *string                     `pulumi:"etag"`
	Id                       *string                     `pulumi:"id"`
	IdleTimeoutInMinutes     *int                        `pulumi:"idleTimeoutInMinutes"`
	IpAddress                *string                     `pulumi:"ipAddress"`
	Location                 *string                     `pulumi:"location"`
	ProvisioningState        *string                     `pulumi:"provisioningState"`
	PublicIPAddressVersion   *string                     `pulumi:"publicIPAddressVersion"`
	PublicIPAllocationMethod *string                     `pulumi:"publicIPAllocationMethod"`
	PublicIpAddressName      *string                     `pulumi:"publicIpAddressName"`
	ResourceGroupName        string                      `pulumi:"resourceGroupName"`
	ResourceGuid             *string                     `pulumi:"resourceGuid"`
	Tags                     map[string]string           `pulumi:"tags"`
}


type PublicIPAddressArgs struct {
	DnsSettings              PublicIPAddressDnsSettingsPtrInput
	Etag                     pulumi.StringPtrInput
	Id                       pulumi.StringPtrInput
	IdleTimeoutInMinutes     pulumi.IntPtrInput
	IpAddress                pulumi.StringPtrInput
	Location                 pulumi.StringPtrInput
	ProvisioningState        pulumi.StringPtrInput
	PublicIPAddressVersion   pulumi.StringPtrInput
	PublicIPAllocationMethod pulumi.StringPtrInput
	PublicIpAddressName      pulumi.StringPtrInput
	ResourceGroupName        pulumi.StringInput
	ResourceGuid             pulumi.StringPtrInput
	Tags                     pulumi.StringMapInput
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
	return reflect.TypeOf((*PublicIPAddress)(nil))
}

func (i *PublicIPAddress) ToPublicIPAddressOutput() PublicIPAddressOutput {
	return i.ToPublicIPAddressOutputWithContext(context.Background())
}

func (i *PublicIPAddress) ToPublicIPAddressOutputWithContext(ctx context.Context) PublicIPAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressOutput)
}

type PublicIPAddressOutput struct{ *pulumi.OutputState }

func (PublicIPAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddress)(nil))
}

func (o PublicIPAddressOutput) ToPublicIPAddressOutput() PublicIPAddressOutput {
	return o
}

func (o PublicIPAddressOutput) ToPublicIPAddressOutputWithContext(ctx context.Context) PublicIPAddressOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PublicIPAddressInput)(nil)).Elem(), &PublicIPAddress{})
	pulumi.RegisterOutputType(PublicIPAddressOutput{})
}
