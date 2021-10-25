


package v20200801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualNetworkGateway struct {
	pulumi.CustomResourceState

	ActiveActive                   pulumi.BoolPtrOutput                                    `pulumi:"activeActive"`
	BgpSettings                    BgpSettingsResponsePtrOutput                            `pulumi:"bgpSettings"`
	CustomRoutes                   AddressSpaceResponsePtrOutput                           `pulumi:"customRoutes"`
	EnableBgp                      pulumi.BoolPtrOutput                                    `pulumi:"enableBgp"`
	EnableDnsForwarding            pulumi.BoolPtrOutput                                    `pulumi:"enableDnsForwarding"`
	EnablePrivateIpAddress         pulumi.BoolPtrOutput                                    `pulumi:"enablePrivateIpAddress"`
	Etag                           pulumi.StringOutput                                     `pulumi:"etag"`
	ExtendedLocation               ExtendedLocationResponsePtrOutput                       `pulumi:"extendedLocation"`
	GatewayDefaultSite             SubResourceResponsePtrOutput                            `pulumi:"gatewayDefaultSite"`
	GatewayType                    pulumi.StringPtrOutput                                  `pulumi:"gatewayType"`
	InboundDnsForwardingEndpoint   pulumi.StringOutput                                     `pulumi:"inboundDnsForwardingEndpoint"`
	IpConfigurations               VirtualNetworkGatewayIPConfigurationResponseArrayOutput `pulumi:"ipConfigurations"`
	Location                       pulumi.StringPtrOutput                                  `pulumi:"location"`
	Name                           pulumi.StringOutput                                     `pulumi:"name"`
	ProvisioningState              pulumi.StringOutput                                     `pulumi:"provisioningState"`
	ResourceGuid                   pulumi.StringOutput                                     `pulumi:"resourceGuid"`
	Sku                            VirtualNetworkGatewaySkuResponsePtrOutput               `pulumi:"sku"`
	Tags                           pulumi.StringMapOutput                                  `pulumi:"tags"`
	Type                           pulumi.StringOutput                                     `pulumi:"type"`
	VNetExtendedLocationResourceId pulumi.StringPtrOutput                                  `pulumi:"vNetExtendedLocationResourceId"`
	VpnClientConfiguration         VpnClientConfigurationResponsePtrOutput                 `pulumi:"vpnClientConfiguration"`
	VpnGatewayGeneration           pulumi.StringPtrOutput                                  `pulumi:"vpnGatewayGeneration"`
	VpnType                        pulumi.StringPtrOutput                                  `pulumi:"vpnType"`
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
			Type: pulumi.String("azure-nextgen:network/v20200801:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20150615:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20160330:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20160601:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20160901:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20161201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170301:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170601:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170801:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170901:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20171001:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20171101:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180101:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180401:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180601:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180701:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180801:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181001:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181101:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190401:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190601:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20201101:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210301:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210501:VirtualNetworkGateway"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualNetworkGateway
	err := ctx.RegisterResource("azure-native:network/v20200801:VirtualNetworkGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualNetworkGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkGatewayState, opts ...pulumi.ResourceOption) (*VirtualNetworkGateway, error) {
	var resource VirtualNetworkGateway
	err := ctx.ReadResource("azure-native:network/v20200801:VirtualNetworkGateway", name, id, state, &resource, opts...)
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
	ActiveActive                   *bool                                  `pulumi:"activeActive"`
	BgpSettings                    *BgpSettings                           `pulumi:"bgpSettings"`
	CustomRoutes                   *AddressSpace                          `pulumi:"customRoutes"`
	EnableBgp                      *bool                                  `pulumi:"enableBgp"`
	EnableDnsForwarding            *bool                                  `pulumi:"enableDnsForwarding"`
	EnablePrivateIpAddress         *bool                                  `pulumi:"enablePrivateIpAddress"`
	ExtendedLocation               *ExtendedLocation                      `pulumi:"extendedLocation"`
	GatewayDefaultSite             *SubResource                           `pulumi:"gatewayDefaultSite"`
	GatewayType                    *string                                `pulumi:"gatewayType"`
	Id                             *string                                `pulumi:"id"`
	IpConfigurations               []VirtualNetworkGatewayIPConfiguration `pulumi:"ipConfigurations"`
	Location                       *string                                `pulumi:"location"`
	ResourceGroupName              string                                 `pulumi:"resourceGroupName"`
	Sku                            *VirtualNetworkGatewaySku              `pulumi:"sku"`
	Tags                           map[string]string                      `pulumi:"tags"`
	VNetExtendedLocationResourceId *string                                `pulumi:"vNetExtendedLocationResourceId"`
	VirtualNetworkGatewayName      *string                                `pulumi:"virtualNetworkGatewayName"`
	VpnClientConfiguration         *VpnClientConfiguration                `pulumi:"vpnClientConfiguration"`
	VpnGatewayGeneration           *string                                `pulumi:"vpnGatewayGeneration"`
	VpnType                        *string                                `pulumi:"vpnType"`
}


type VirtualNetworkGatewayArgs struct {
	ActiveActive                   pulumi.BoolPtrInput
	BgpSettings                    BgpSettingsPtrInput
	CustomRoutes                   AddressSpacePtrInput
	EnableBgp                      pulumi.BoolPtrInput
	EnableDnsForwarding            pulumi.BoolPtrInput
	EnablePrivateIpAddress         pulumi.BoolPtrInput
	ExtendedLocation               ExtendedLocationPtrInput
	GatewayDefaultSite             SubResourcePtrInput
	GatewayType                    pulumi.StringPtrInput
	Id                             pulumi.StringPtrInput
	IpConfigurations               VirtualNetworkGatewayIPConfigurationArrayInput
	Location                       pulumi.StringPtrInput
	ResourceGroupName              pulumi.StringInput
	Sku                            VirtualNetworkGatewaySkuPtrInput
	Tags                           pulumi.StringMapInput
	VNetExtendedLocationResourceId pulumi.StringPtrInput
	VirtualNetworkGatewayName      pulumi.StringPtrInput
	VpnClientConfiguration         VpnClientConfigurationPtrInput
	VpnGatewayGeneration           pulumi.StringPtrInput
	VpnType                        pulumi.StringPtrInput
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
	return reflect.TypeOf((*VirtualNetworkGateway)(nil))
}

func (i *VirtualNetworkGateway) ToVirtualNetworkGatewayOutput() VirtualNetworkGatewayOutput {
	return i.ToVirtualNetworkGatewayOutputWithContext(context.Background())
}

func (i *VirtualNetworkGateway) ToVirtualNetworkGatewayOutputWithContext(ctx context.Context) VirtualNetworkGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayOutput)
}

type VirtualNetworkGatewayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGateway)(nil))
}

func (o VirtualNetworkGatewayOutput) ToVirtualNetworkGatewayOutput() VirtualNetworkGatewayOutput {
	return o
}

func (o VirtualNetworkGatewayOutput) ToVirtualNetworkGatewayOutputWithContext(ctx context.Context) VirtualNetworkGatewayOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkGatewayOutput{})
}
