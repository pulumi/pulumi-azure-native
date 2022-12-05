


package v20181001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type ContainerGroup struct {
	pulumi.CustomResourceState

	Containers               ContainerResponseArrayOutput                  `pulumi:"containers"`
	Diagnostics              ContainerGroupDiagnosticsResponsePtrOutput    `pulumi:"diagnostics"`
	DnsConfig                DnsConfigurationResponsePtrOutput             `pulumi:"dnsConfig"`
	Identity                 ContainerGroupIdentityResponsePtrOutput       `pulumi:"identity"`
	ImageRegistryCredentials ImageRegistryCredentialResponseArrayOutput    `pulumi:"imageRegistryCredentials"`
	InstanceView             ContainerGroupResponseInstanceViewOutput      `pulumi:"instanceView"`
	IpAddress                IpAddressResponsePtrOutput                    `pulumi:"ipAddress"`
	Location                 pulumi.StringPtrOutput                        `pulumi:"location"`
	Name                     pulumi.StringOutput                           `pulumi:"name"`
	NetworkProfile           ContainerGroupNetworkProfileResponsePtrOutput `pulumi:"networkProfile"`
	OsType                   pulumi.StringOutput                           `pulumi:"osType"`
	ProvisioningState        pulumi.StringOutput                           `pulumi:"provisioningState"`
	RestartPolicy            pulumi.StringPtrOutput                        `pulumi:"restartPolicy"`
	Tags                     pulumi.StringMapOutput                        `pulumi:"tags"`
	Type                     pulumi.StringOutput                           `pulumi:"type"`
	Volumes                  VolumeResponseArrayOutput                     `pulumi:"volumes"`
}


func NewContainerGroup(ctx *pulumi.Context,
	name string, args *ContainerGroupArgs, opts ...pulumi.ResourceOption) (*ContainerGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Containers == nil {
		return nil, errors.New("invalid value for required argument 'Containers'")
	}
	if args.OsType == nil {
		return nil, errors.New("invalid value for required argument 'OsType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerinstance:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20170801preview:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20171001preview:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20171201preview:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20180201preview:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20180401:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20180601:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20180901:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20191201:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20201101:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20210301:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20210701:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20210901:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20211001:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20220901:ContainerGroup"),
		},
		{
			Type: pulumi.String("azure-native:containerinstance/v20221001preview:ContainerGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource ContainerGroup
	err := ctx.RegisterResource("azure-native:containerinstance/v20181001:ContainerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetContainerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerGroupState, opts ...pulumi.ResourceOption) (*ContainerGroup, error) {
	var resource ContainerGroup
	err := ctx.ReadResource("azure-native:containerinstance/v20181001:ContainerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type containerGroupState struct {
}

type ContainerGroupState struct {
}

func (ContainerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerGroupState)(nil)).Elem()
}

type containerGroupArgs struct {
	ContainerGroupName       *string                       `pulumi:"containerGroupName"`
	Containers               []Container                   `pulumi:"containers"`
	Diagnostics              *ContainerGroupDiagnostics    `pulumi:"diagnostics"`
	DnsConfig                *DnsConfiguration             `pulumi:"dnsConfig"`
	Identity                 *ContainerGroupIdentity       `pulumi:"identity"`
	ImageRegistryCredentials []ImageRegistryCredential     `pulumi:"imageRegistryCredentials"`
	IpAddress                *IpAddress                    `pulumi:"ipAddress"`
	Location                 *string                       `pulumi:"location"`
	NetworkProfile           *ContainerGroupNetworkProfile `pulumi:"networkProfile"`
	OsType                   string                        `pulumi:"osType"`
	ResourceGroupName        string                        `pulumi:"resourceGroupName"`
	RestartPolicy            *string                       `pulumi:"restartPolicy"`
	Tags                     map[string]string             `pulumi:"tags"`
	Volumes                  []Volume                      `pulumi:"volumes"`
}


type ContainerGroupArgs struct {
	ContainerGroupName       pulumi.StringPtrInput
	Containers               ContainerArrayInput
	Diagnostics              ContainerGroupDiagnosticsPtrInput
	DnsConfig                DnsConfigurationPtrInput
	Identity                 ContainerGroupIdentityPtrInput
	ImageRegistryCredentials ImageRegistryCredentialArrayInput
	IpAddress                IpAddressPtrInput
	Location                 pulumi.StringPtrInput
	NetworkProfile           ContainerGroupNetworkProfilePtrInput
	OsType                   pulumi.StringInput
	ResourceGroupName        pulumi.StringInput
	RestartPolicy            pulumi.StringPtrInput
	Tags                     pulumi.StringMapInput
	Volumes                  VolumeArrayInput
}

func (ContainerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerGroupArgs)(nil)).Elem()
}

type ContainerGroupInput interface {
	pulumi.Input

	ToContainerGroupOutput() ContainerGroupOutput
	ToContainerGroupOutputWithContext(ctx context.Context) ContainerGroupOutput
}

func (*ContainerGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerGroup)(nil)).Elem()
}

func (i *ContainerGroup) ToContainerGroupOutput() ContainerGroupOutput {
	return i.ToContainerGroupOutputWithContext(context.Background())
}

func (i *ContainerGroup) ToContainerGroupOutputWithContext(ctx context.Context) ContainerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerGroupOutput)
}

type ContainerGroupOutput struct{ *pulumi.OutputState }

func (ContainerGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerGroup)(nil)).Elem()
}

func (o ContainerGroupOutput) ToContainerGroupOutput() ContainerGroupOutput {
	return o
}

func (o ContainerGroupOutput) ToContainerGroupOutputWithContext(ctx context.Context) ContainerGroupOutput {
	return o
}

func (o ContainerGroupOutput) Containers() ContainerResponseArrayOutput {
	return o.ApplyT(func(v *ContainerGroup) ContainerResponseArrayOutput { return v.Containers }).(ContainerResponseArrayOutput)
}

func (o ContainerGroupOutput) Diagnostics() ContainerGroupDiagnosticsResponsePtrOutput {
	return o.ApplyT(func(v *ContainerGroup) ContainerGroupDiagnosticsResponsePtrOutput { return v.Diagnostics }).(ContainerGroupDiagnosticsResponsePtrOutput)
}

func (o ContainerGroupOutput) DnsConfig() DnsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ContainerGroup) DnsConfigurationResponsePtrOutput { return v.DnsConfig }).(DnsConfigurationResponsePtrOutput)
}

func (o ContainerGroupOutput) Identity() ContainerGroupIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ContainerGroup) ContainerGroupIdentityResponsePtrOutput { return v.Identity }).(ContainerGroupIdentityResponsePtrOutput)
}

func (o ContainerGroupOutput) ImageRegistryCredentials() ImageRegistryCredentialResponseArrayOutput {
	return o.ApplyT(func(v *ContainerGroup) ImageRegistryCredentialResponseArrayOutput { return v.ImageRegistryCredentials }).(ImageRegistryCredentialResponseArrayOutput)
}

func (o ContainerGroupOutput) InstanceView() ContainerGroupResponseInstanceViewOutput {
	return o.ApplyT(func(v *ContainerGroup) ContainerGroupResponseInstanceViewOutput { return v.InstanceView }).(ContainerGroupResponseInstanceViewOutput)
}

func (o ContainerGroupOutput) IpAddress() IpAddressResponsePtrOutput {
	return o.ApplyT(func(v *ContainerGroup) IpAddressResponsePtrOutput { return v.IpAddress }).(IpAddressResponsePtrOutput)
}

func (o ContainerGroupOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerGroup) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ContainerGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerGroupOutput) NetworkProfile() ContainerGroupNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *ContainerGroup) ContainerGroupNetworkProfileResponsePtrOutput { return v.NetworkProfile }).(ContainerGroupNetworkProfileResponsePtrOutput)
}

func (o ContainerGroupOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerGroup) pulumi.StringOutput { return v.OsType }).(pulumi.StringOutput)
}

func (o ContainerGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ContainerGroupOutput) RestartPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerGroup) pulumi.StringPtrOutput { return v.RestartPolicy }).(pulumi.StringPtrOutput)
}

func (o ContainerGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ContainerGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ContainerGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ContainerGroupOutput) Volumes() VolumeResponseArrayOutput {
	return o.ApplyT(func(v *ContainerGroup) VolumeResponseArrayOutput { return v.Volumes }).(VolumeResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerGroupOutput{})
}
