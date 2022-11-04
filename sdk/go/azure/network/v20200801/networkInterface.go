


package v20200801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkInterface struct {
	pulumi.CustomResourceState

	DnsSettings                 NetworkInterfaceDnsSettingsResponsePtrOutput        `pulumi:"dnsSettings"`
	DscpConfiguration           SubResourceResponseOutput                           `pulumi:"dscpConfiguration"`
	EnableAcceleratedNetworking pulumi.BoolPtrOutput                                `pulumi:"enableAcceleratedNetworking"`
	EnableIPForwarding          pulumi.BoolPtrOutput                                `pulumi:"enableIPForwarding"`
	Etag                        pulumi.StringOutput                                 `pulumi:"etag"`
	ExtendedLocation            ExtendedLocationResponsePtrOutput                   `pulumi:"extendedLocation"`
	HostedWorkloads             pulumi.StringArrayOutput                            `pulumi:"hostedWorkloads"`
	IpConfigurations            NetworkInterfaceIPConfigurationResponseArrayOutput  `pulumi:"ipConfigurations"`
	Location                    pulumi.StringPtrOutput                              `pulumi:"location"`
	MacAddress                  pulumi.StringOutput                                 `pulumi:"macAddress"`
	MigrationPhase              pulumi.StringPtrOutput                              `pulumi:"migrationPhase"`
	Name                        pulumi.StringOutput                                 `pulumi:"name"`
	NetworkSecurityGroup        NetworkSecurityGroupResponsePtrOutput               `pulumi:"networkSecurityGroup"`
	NicType                     pulumi.StringPtrOutput                              `pulumi:"nicType"`
	Primary                     pulumi.BoolOutput                                   `pulumi:"primary"`
	PrivateEndpoint             PrivateEndpointResponseOutput                       `pulumi:"privateEndpoint"`
	PrivateLinkService          PrivateLinkServiceResponsePtrOutput                 `pulumi:"privateLinkService"`
	ProvisioningState           pulumi.StringOutput                                 `pulumi:"provisioningState"`
	ResourceGuid                pulumi.StringOutput                                 `pulumi:"resourceGuid"`
	Tags                        pulumi.StringMapOutput                              `pulumi:"tags"`
	TapConfigurations           NetworkInterfaceTapConfigurationResponseArrayOutput `pulumi:"tapConfigurations"`
	Type                        pulumi.StringOutput                                 `pulumi:"type"`
	VirtualMachine              SubResourceResponseOutput                           `pulumi:"virtualMachine"`
}


func NewNetworkInterface(ctx *pulumi.Context,
	name string, args *NetworkInterfaceArgs, opts ...pulumi.ResourceOption) (*NetworkInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.NicType) {
		args.NicType = pulumi.StringPtr("Standard")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150501preview:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:NetworkInterface"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:NetworkInterface"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkInterface
	err := ctx.RegisterResource("azure-native:network/v20200801:NetworkInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNetworkInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceState, opts ...pulumi.ResourceOption) (*NetworkInterface, error) {
	var resource NetworkInterface
	err := ctx.ReadResource("azure-native:network/v20200801:NetworkInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type networkInterfaceState struct {
}

type NetworkInterfaceState struct {
}

func (NetworkInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceState)(nil)).Elem()
}

type networkInterfaceArgs struct {
	DnsSettings                 *NetworkInterfaceDnsSettings      `pulumi:"dnsSettings"`
	EnableAcceleratedNetworking *bool                             `pulumi:"enableAcceleratedNetworking"`
	EnableIPForwarding          *bool                             `pulumi:"enableIPForwarding"`
	ExtendedLocation            *ExtendedLocation                 `pulumi:"extendedLocation"`
	Id                          *string                           `pulumi:"id"`
	IpConfigurations            []NetworkInterfaceIPConfiguration `pulumi:"ipConfigurations"`
	Location                    *string                           `pulumi:"location"`
	MigrationPhase              *string                           `pulumi:"migrationPhase"`
	NetworkInterfaceName        *string                           `pulumi:"networkInterfaceName"`
	NetworkSecurityGroup        *NetworkSecurityGroupType         `pulumi:"networkSecurityGroup"`
	NicType                     *string                           `pulumi:"nicType"`
	PrivateLinkService          *PrivateLinkServiceType           `pulumi:"privateLinkService"`
	ResourceGroupName           string                            `pulumi:"resourceGroupName"`
	Tags                        map[string]string                 `pulumi:"tags"`
}


type NetworkInterfaceArgs struct {
	DnsSettings                 NetworkInterfaceDnsSettingsPtrInput
	EnableAcceleratedNetworking pulumi.BoolPtrInput
	EnableIPForwarding          pulumi.BoolPtrInput
	ExtendedLocation            ExtendedLocationPtrInput
	Id                          pulumi.StringPtrInput
	IpConfigurations            NetworkInterfaceIPConfigurationArrayInput
	Location                    pulumi.StringPtrInput
	MigrationPhase              pulumi.StringPtrInput
	NetworkInterfaceName        pulumi.StringPtrInput
	NetworkSecurityGroup        NetworkSecurityGroupTypePtrInput
	NicType                     pulumi.StringPtrInput
	PrivateLinkService          PrivateLinkServiceTypePtrInput
	ResourceGroupName           pulumi.StringInput
	Tags                        pulumi.StringMapInput
}

func (NetworkInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceArgs)(nil)).Elem()
}

type NetworkInterfaceInput interface {
	pulumi.Input

	ToNetworkInterfaceOutput() NetworkInterfaceOutput
	ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput
}

func (*NetworkInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterface)(nil)).Elem()
}

func (i *NetworkInterface) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return i.ToNetworkInterfaceOutputWithContext(context.Background())
}

func (i *NetworkInterface) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceOutput)
}

type NetworkInterfaceOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterface)(nil)).Elem()
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return o
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return o
}

func (o NetworkInterfaceOutput) DnsSettings() NetworkInterfaceDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *NetworkInterface) NetworkInterfaceDnsSettingsResponsePtrOutput { return v.DnsSettings }).(NetworkInterfaceDnsSettingsResponsePtrOutput)
}

func (o NetworkInterfaceOutput) DscpConfiguration() SubResourceResponseOutput {
	return o.ApplyT(func(v *NetworkInterface) SubResourceResponseOutput { return v.DscpConfiguration }).(SubResourceResponseOutput)
}

func (o NetworkInterfaceOutput) EnableAcceleratedNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.BoolPtrOutput { return v.EnableAcceleratedNetworking }).(pulumi.BoolPtrOutput)
}

func (o NetworkInterfaceOutput) EnableIPForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.BoolPtrOutput { return v.EnableIPForwarding }).(pulumi.BoolPtrOutput)
}

func (o NetworkInterfaceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o NetworkInterfaceOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *NetworkInterface) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o NetworkInterfaceOutput) HostedWorkloads() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringArrayOutput { return v.HostedWorkloads }).(pulumi.StringArrayOutput)
}

func (o NetworkInterfaceOutput) IpConfigurations() NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *NetworkInterface) NetworkInterfaceIPConfigurationResponseArrayOutput {
		return v.IpConfigurations
	}).(NetworkInterfaceIPConfigurationResponseArrayOutput)
}

func (o NetworkInterfaceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.MacAddress }).(pulumi.StringOutput)
}

func (o NetworkInterfaceOutput) MigrationPhase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.MigrationPhase }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkInterfaceOutput) NetworkSecurityGroup() NetworkSecurityGroupResponsePtrOutput {
	return o.ApplyT(func(v *NetworkInterface) NetworkSecurityGroupResponsePtrOutput { return v.NetworkSecurityGroup }).(NetworkSecurityGroupResponsePtrOutput)
}

func (o NetworkInterfaceOutput) NicType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.NicType }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceOutput) Primary() pulumi.BoolOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.BoolOutput { return v.Primary }).(pulumi.BoolOutput)
}

func (o NetworkInterfaceOutput) PrivateEndpoint() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *NetworkInterface) PrivateEndpointResponseOutput { return v.PrivateEndpoint }).(PrivateEndpointResponseOutput)
}

func (o NetworkInterfaceOutput) PrivateLinkService() PrivateLinkServiceResponsePtrOutput {
	return o.ApplyT(func(v *NetworkInterface) PrivateLinkServiceResponsePtrOutput { return v.PrivateLinkService }).(PrivateLinkServiceResponsePtrOutput)
}

func (o NetworkInterfaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o NetworkInterfaceOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o NetworkInterfaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NetworkInterfaceOutput) TapConfigurations() NetworkInterfaceTapConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *NetworkInterface) NetworkInterfaceTapConfigurationResponseArrayOutput {
		return v.TapConfigurations
	}).(NetworkInterfaceTapConfigurationResponseArrayOutput)
}

func (o NetworkInterfaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o NetworkInterfaceOutput) VirtualMachine() SubResourceResponseOutput {
	return o.ApplyT(func(v *NetworkInterface) SubResourceResponseOutput { return v.VirtualMachine }).(SubResourceResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkInterfaceOutput{})
}
