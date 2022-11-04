


package v20180201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type NetworkInterface struct {
	pulumi.CustomResourceState

	DnsSettings                 NetworkInterfaceDnsSettingsResponsePtrOutput       `pulumi:"dnsSettings"`
	EnableAcceleratedNetworking pulumi.BoolPtrOutput                               `pulumi:"enableAcceleratedNetworking"`
	EnableIPForwarding          pulumi.BoolPtrOutput                               `pulumi:"enableIPForwarding"`
	Etag                        pulumi.StringPtrOutput                             `pulumi:"etag"`
	IpConfigurations            NetworkInterfaceIPConfigurationResponseArrayOutput `pulumi:"ipConfigurations"`
	Location                    pulumi.StringPtrOutput                             `pulumi:"location"`
	MacAddress                  pulumi.StringPtrOutput                             `pulumi:"macAddress"`
	Name                        pulumi.StringOutput                                `pulumi:"name"`
	NetworkSecurityGroup        NetworkSecurityGroupResponsePtrOutput              `pulumi:"networkSecurityGroup"`
	Primary                     pulumi.BoolPtrOutput                               `pulumi:"primary"`
	ProvisioningState           pulumi.StringPtrOutput                             `pulumi:"provisioningState"`
	ResourceGuid                pulumi.StringPtrOutput                             `pulumi:"resourceGuid"`
	Tags                        pulumi.StringMapOutput                             `pulumi:"tags"`
	Type                        pulumi.StringOutput                                `pulumi:"type"`
	VirtualMachine              SubResourceResponsePtrOutput                       `pulumi:"virtualMachine"`
}


func NewNetworkInterface(ctx *pulumi.Context,
	name string, args *NetworkInterfaceArgs, opts ...pulumi.ResourceOption) (*NetworkInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
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
			Type: pulumi.String("azure-native:network/v20200801:NetworkInterface"),
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
	err := ctx.RegisterResource("azure-native:network/v20180201:NetworkInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNetworkInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceState, opts ...pulumi.ResourceOption) (*NetworkInterface, error) {
	var resource NetworkInterface
	err := ctx.ReadResource("azure-native:network/v20180201:NetworkInterface", name, id, state, &resource, opts...)
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
	Id                          *string                           `pulumi:"id"`
	IpConfigurations            []NetworkInterfaceIPConfiguration `pulumi:"ipConfigurations"`
	Location                    *string                           `pulumi:"location"`
	MacAddress                  *string                           `pulumi:"macAddress"`
	NetworkInterfaceName        *string                           `pulumi:"networkInterfaceName"`
	NetworkSecurityGroup        *NetworkSecurityGroupType         `pulumi:"networkSecurityGroup"`
	Primary                     *bool                             `pulumi:"primary"`
	ProvisioningState           *string                           `pulumi:"provisioningState"`
	ResourceGroupName           string                            `pulumi:"resourceGroupName"`
	ResourceGuid                *string                           `pulumi:"resourceGuid"`
	Tags                        map[string]string                 `pulumi:"tags"`
	VirtualMachine              *SubResource                      `pulumi:"virtualMachine"`
}


type NetworkInterfaceArgs struct {
	DnsSettings                 NetworkInterfaceDnsSettingsPtrInput
	EnableAcceleratedNetworking pulumi.BoolPtrInput
	EnableIPForwarding          pulumi.BoolPtrInput
	Id                          pulumi.StringPtrInput
	IpConfigurations            NetworkInterfaceIPConfigurationArrayInput
	Location                    pulumi.StringPtrInput
	MacAddress                  pulumi.StringPtrInput
	NetworkInterfaceName        pulumi.StringPtrInput
	NetworkSecurityGroup        NetworkSecurityGroupTypePtrInput
	Primary                     pulumi.BoolPtrInput
	ProvisioningState           pulumi.StringPtrInput
	ResourceGroupName           pulumi.StringInput
	ResourceGuid                pulumi.StringPtrInput
	Tags                        pulumi.StringMapInput
	VirtualMachine              SubResourcePtrInput
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

func (o NetworkInterfaceOutput) EnableAcceleratedNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.BoolPtrOutput { return v.EnableAcceleratedNetworking }).(pulumi.BoolPtrOutput)
}

func (o NetworkInterfaceOutput) EnableIPForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.BoolPtrOutput { return v.EnableIPForwarding }).(pulumi.BoolPtrOutput)
}

func (o NetworkInterfaceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceOutput) IpConfigurations() NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *NetworkInterface) NetworkInterfaceIPConfigurationResponseArrayOutput {
		return v.IpConfigurations
	}).(NetworkInterfaceIPConfigurationResponseArrayOutput)
}

func (o NetworkInterfaceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.MacAddress }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkInterfaceOutput) NetworkSecurityGroup() NetworkSecurityGroupResponsePtrOutput {
	return o.ApplyT(func(v *NetworkInterface) NetworkSecurityGroupResponsePtrOutput { return v.NetworkSecurityGroup }).(NetworkSecurityGroupResponsePtrOutput)
}

func (o NetworkInterfaceOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.BoolPtrOutput { return v.Primary }).(pulumi.BoolPtrOutput)
}

func (o NetworkInterfaceOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringPtrOutput { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NetworkInterfaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterface) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o NetworkInterfaceOutput) VirtualMachine() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *NetworkInterface) SubResourceResponsePtrOutput { return v.VirtualMachine }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkInterfaceOutput{})
}
