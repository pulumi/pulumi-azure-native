


package v20160901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type LocalNetworkGateway struct {
	pulumi.CustomResourceState

	BgpSettings              BgpSettingsResponsePtrOutput `pulumi:"bgpSettings"`
	Etag                     pulumi.StringPtrOutput       `pulumi:"etag"`
	GatewayIpAddress         pulumi.StringPtrOutput       `pulumi:"gatewayIpAddress"`
	LocalNetworkAddressSpace AddressSpaceResponseOutput   `pulumi:"localNetworkAddressSpace"`
	Location                 pulumi.StringPtrOutput       `pulumi:"location"`
	Name                     pulumi.StringOutput          `pulumi:"name"`
	ProvisioningState        pulumi.StringOutput          `pulumi:"provisioningState"`
	ResourceGuid             pulumi.StringPtrOutput       `pulumi:"resourceGuid"`
	Tags                     pulumi.StringMapOutput       `pulumi:"tags"`
	Type                     pulumi.StringOutput          `pulumi:"type"`
}


func NewLocalNetworkGateway(ctx *pulumi.Context,
	name string, args *LocalNetworkGatewayArgs, opts ...pulumi.ResourceOption) (*LocalNetworkGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LocalNetworkAddressSpace == nil {
		return nil, errors.New("invalid value for required argument 'LocalNetworkAddressSpace'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:LocalNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:LocalNetworkGateway"),
		},
	})
	opts = append(opts, aliases)
	var resource LocalNetworkGateway
	err := ctx.RegisterResource("azure-native:network/v20160901:LocalNetworkGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLocalNetworkGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalNetworkGatewayState, opts ...pulumi.ResourceOption) (*LocalNetworkGateway, error) {
	var resource LocalNetworkGateway
	err := ctx.ReadResource("azure-native:network/v20160901:LocalNetworkGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type localNetworkGatewayState struct {
}

type LocalNetworkGatewayState struct {
}

func (LocalNetworkGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*localNetworkGatewayState)(nil)).Elem()
}

type localNetworkGatewayArgs struct {
	BgpSettings              *BgpSettings      `pulumi:"bgpSettings"`
	GatewayIpAddress         *string           `pulumi:"gatewayIpAddress"`
	Id                       *string           `pulumi:"id"`
	LocalNetworkAddressSpace AddressSpace      `pulumi:"localNetworkAddressSpace"`
	LocalNetworkGatewayName  *string           `pulumi:"localNetworkGatewayName"`
	Location                 *string           `pulumi:"location"`
	ResourceGroupName        string            `pulumi:"resourceGroupName"`
	ResourceGuid             *string           `pulumi:"resourceGuid"`
	Tags                     map[string]string `pulumi:"tags"`
}


type LocalNetworkGatewayArgs struct {
	BgpSettings              BgpSettingsPtrInput
	GatewayIpAddress         pulumi.StringPtrInput
	Id                       pulumi.StringPtrInput
	LocalNetworkAddressSpace AddressSpaceInput
	LocalNetworkGatewayName  pulumi.StringPtrInput
	Location                 pulumi.StringPtrInput
	ResourceGroupName        pulumi.StringInput
	ResourceGuid             pulumi.StringPtrInput
	Tags                     pulumi.StringMapInput
}

func (LocalNetworkGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localNetworkGatewayArgs)(nil)).Elem()
}

type LocalNetworkGatewayInput interface {
	pulumi.Input

	ToLocalNetworkGatewayOutput() LocalNetworkGatewayOutput
	ToLocalNetworkGatewayOutputWithContext(ctx context.Context) LocalNetworkGatewayOutput
}

func (*LocalNetworkGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalNetworkGateway)(nil)).Elem()
}

func (i *LocalNetworkGateway) ToLocalNetworkGatewayOutput() LocalNetworkGatewayOutput {
	return i.ToLocalNetworkGatewayOutputWithContext(context.Background())
}

func (i *LocalNetworkGateway) ToLocalNetworkGatewayOutputWithContext(ctx context.Context) LocalNetworkGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNetworkGatewayOutput)
}

type LocalNetworkGatewayOutput struct{ *pulumi.OutputState }

func (LocalNetworkGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalNetworkGateway)(nil)).Elem()
}

func (o LocalNetworkGatewayOutput) ToLocalNetworkGatewayOutput() LocalNetworkGatewayOutput {
	return o
}

func (o LocalNetworkGatewayOutput) ToLocalNetworkGatewayOutputWithContext(ctx context.Context) LocalNetworkGatewayOutput {
	return o
}

func (o LocalNetworkGatewayOutput) BgpSettings() BgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v *LocalNetworkGateway) BgpSettingsResponsePtrOutput { return v.BgpSettings }).(BgpSettingsResponsePtrOutput)
}

func (o LocalNetworkGatewayOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGateway) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayOutput) GatewayIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGateway) pulumi.StringPtrOutput { return v.GatewayIpAddress }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayOutput) LocalNetworkAddressSpace() AddressSpaceResponseOutput {
	return o.ApplyT(func(v *LocalNetworkGateway) AddressSpaceResponseOutput { return v.LocalNetworkAddressSpace }).(AddressSpaceResponseOutput)
}

func (o LocalNetworkGatewayOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGateway) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalNetworkGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LocalNetworkGatewayOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalNetworkGateway) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LocalNetworkGatewayOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGateway) pulumi.StringPtrOutput { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LocalNetworkGateway) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LocalNetworkGatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalNetworkGateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LocalNetworkGatewayOutput{})
}
