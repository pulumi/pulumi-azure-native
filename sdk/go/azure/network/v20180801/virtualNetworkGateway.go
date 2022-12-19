


package v20180801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualNetworkGateway struct {
	pulumi.CustomResourceState

	ActiveActive           pulumi.BoolPtrOutput                                    `pulumi:"activeActive"`
	BgpSettings            BgpSettingsResponsePtrOutput                            `pulumi:"bgpSettings"`
	EnableBgp              pulumi.BoolPtrOutput                                    `pulumi:"enableBgp"`
	Etag                   pulumi.StringPtrOutput                                  `pulumi:"etag"`
	GatewayDefaultSite     SubResourceResponsePtrOutput                            `pulumi:"gatewayDefaultSite"`
	GatewayType            pulumi.StringPtrOutput                                  `pulumi:"gatewayType"`
	IpConfigurations       VirtualNetworkGatewayIPConfigurationResponseArrayOutput `pulumi:"ipConfigurations"`
	Location               pulumi.StringPtrOutput                                  `pulumi:"location"`
	Name                   pulumi.StringOutput                                     `pulumi:"name"`
	ProvisioningState      pulumi.StringOutput                                     `pulumi:"provisioningState"`
	ResourceGuid           pulumi.StringPtrOutput                                  `pulumi:"resourceGuid"`
	Sku                    VirtualNetworkGatewaySkuResponsePtrOutput               `pulumi:"sku"`
	Tags                   pulumi.StringMapOutput                                  `pulumi:"tags"`
	Type                   pulumi.StringOutput                                     `pulumi:"type"`
	VpnClientConfiguration VpnClientConfigurationResponsePtrOutput                 `pulumi:"vpnClientConfiguration"`
	VpnType                pulumi.StringPtrOutput                                  `pulumi:"vpnType"`
}


func NewVirtualNetworkGateway(ctx *pulumi.Context,
	name string, args *VirtualNetworkGatewayArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:VirtualNetworkGateway"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualNetworkGateway
	err := ctx.RegisterResource("azure-native:network/v20180801:VirtualNetworkGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualNetworkGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkGatewayState, opts ...pulumi.ResourceOption) (*VirtualNetworkGateway, error) {
	var resource VirtualNetworkGateway
	err := ctx.ReadResource("azure-native:network/v20180801:VirtualNetworkGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualNetworkGatewayState struct {
}

type VirtualNetworkGatewayState struct {
}

func (VirtualNetworkGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkGatewayState)(nil)).Elem()
}

type virtualNetworkGatewayArgs struct {
	ActiveActive              *bool                                  `pulumi:"activeActive"`
	BgpSettings               *BgpSettings                           `pulumi:"bgpSettings"`
	EnableBgp                 *bool                                  `pulumi:"enableBgp"`
	GatewayDefaultSite        *SubResource                           `pulumi:"gatewayDefaultSite"`
	GatewayType               *string                                `pulumi:"gatewayType"`
	Id                        *string                                `pulumi:"id"`
	IpConfigurations          []VirtualNetworkGatewayIPConfiguration `pulumi:"ipConfigurations"`
	Location                  *string                                `pulumi:"location"`
	ResourceGroupName         string                                 `pulumi:"resourceGroupName"`
	ResourceGuid              *string                                `pulumi:"resourceGuid"`
	Sku                       *VirtualNetworkGatewaySku              `pulumi:"sku"`
	Tags                      map[string]string                      `pulumi:"tags"`
	VirtualNetworkGatewayName *string                                `pulumi:"virtualNetworkGatewayName"`
	VpnClientConfiguration    *VpnClientConfiguration                `pulumi:"vpnClientConfiguration"`
	VpnType                   *string                                `pulumi:"vpnType"`
}


type VirtualNetworkGatewayArgs struct {
	ActiveActive              pulumi.BoolPtrInput
	BgpSettings               BgpSettingsPtrInput
	EnableBgp                 pulumi.BoolPtrInput
	GatewayDefaultSite        SubResourcePtrInput
	GatewayType               pulumi.StringPtrInput
	Id                        pulumi.StringPtrInput
	IpConfigurations          VirtualNetworkGatewayIPConfigurationArrayInput
	Location                  pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	ResourceGuid              pulumi.StringPtrInput
	Sku                       VirtualNetworkGatewaySkuPtrInput
	Tags                      pulumi.StringMapInput
	VirtualNetworkGatewayName pulumi.StringPtrInput
	VpnClientConfiguration    VpnClientConfigurationPtrInput
	VpnType                   pulumi.StringPtrInput
}

func (VirtualNetworkGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkGatewayArgs)(nil)).Elem()
}

type VirtualNetworkGatewayInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayOutput() VirtualNetworkGatewayOutput
	ToVirtualNetworkGatewayOutputWithContext(ctx context.Context) VirtualNetworkGatewayOutput
}

func (*VirtualNetworkGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGateway)(nil)).Elem()
}

func (i *VirtualNetworkGateway) ToVirtualNetworkGatewayOutput() VirtualNetworkGatewayOutput {
	return i.ToVirtualNetworkGatewayOutputWithContext(context.Background())
}

func (i *VirtualNetworkGateway) ToVirtualNetworkGatewayOutputWithContext(ctx context.Context) VirtualNetworkGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayOutput)
}

type VirtualNetworkGatewayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGateway)(nil)).Elem()
}

func (o VirtualNetworkGatewayOutput) ToVirtualNetworkGatewayOutput() VirtualNetworkGatewayOutput {
	return o
}

func (o VirtualNetworkGatewayOutput) ToVirtualNetworkGatewayOutputWithContext(ctx context.Context) VirtualNetworkGatewayOutput {
	return o
}

func (o VirtualNetworkGatewayOutput) ActiveActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.BoolPtrOutput { return v.ActiveActive }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayOutput) BgpSettings() BgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) BgpSettingsResponsePtrOutput { return v.BgpSettings }).(BgpSettingsResponsePtrOutput)
}

func (o VirtualNetworkGatewayOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.BoolPtrOutput { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayOutput) GatewayDefaultSite() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) SubResourceResponsePtrOutput { return v.GatewayDefaultSite }).(SubResourceResponsePtrOutput)
}

func (o VirtualNetworkGatewayOutput) GatewayType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.StringPtrOutput { return v.GatewayType }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayOutput) IpConfigurations() VirtualNetworkGatewayIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) VirtualNetworkGatewayIPConfigurationResponseArrayOutput {
		return v.IpConfigurations
	}).(VirtualNetworkGatewayIPConfigurationResponseArrayOutput)
}

func (o VirtualNetworkGatewayOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.StringPtrOutput { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayOutput) Sku() VirtualNetworkGatewaySkuResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) VirtualNetworkGatewaySkuResponsePtrOutput { return v.Sku }).(VirtualNetworkGatewaySkuResponsePtrOutput)
}

func (o VirtualNetworkGatewayOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualNetworkGatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayOutput) VpnClientConfiguration() VpnClientConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) VpnClientConfigurationResponsePtrOutput {
		return v.VpnClientConfiguration
	}).(VpnClientConfigurationResponsePtrOutput)
}

func (o VirtualNetworkGatewayOutput) VpnType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.StringPtrOutput { return v.VpnType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkGatewayOutput{})
}
