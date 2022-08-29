


package v20181101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkProfile struct {
	pulumi.CustomResourceState

	ContainerNetworkInterfaceConfigurations ContainerNetworkInterfaceConfigurationResponseArrayOutput `pulumi:"containerNetworkInterfaceConfigurations"`
	ContainerNetworkInterfaces              ContainerNetworkInterfaceResponseArrayOutput              `pulumi:"containerNetworkInterfaces"`
	Etag                                    pulumi.StringPtrOutput                                    `pulumi:"etag"`
	Location                                pulumi.StringPtrOutput                                    `pulumi:"location"`
	Name                                    pulumi.StringOutput                                       `pulumi:"name"`
	ProvisioningState                       pulumi.StringOutput                                       `pulumi:"provisioningState"`
	ResourceGuid                            pulumi.StringOutput                                       `pulumi:"resourceGuid"`
	Tags                                    pulumi.StringMapOutput                                    `pulumi:"tags"`
	Type                                    pulumi.StringOutput                                       `pulumi:"type"`
}


func NewNetworkProfile(ctx *pulumi.Context,
	name string, args *NetworkProfileArgs, opts ...pulumi.ResourceOption) (*NetworkProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:NetworkProfile"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:NetworkProfile"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkProfile
	err := ctx.RegisterResource("azure-native:network/v20181101:NetworkProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNetworkProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkProfileState, opts ...pulumi.ResourceOption) (*NetworkProfile, error) {
	var resource NetworkProfile
	err := ctx.ReadResource("azure-native:network/v20181101:NetworkProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type networkProfileState struct {
}

type NetworkProfileState struct {
}

func (NetworkProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkProfileState)(nil)).Elem()
}

type networkProfileArgs struct {
	ContainerNetworkInterfaceConfigurations []ContainerNetworkInterfaceConfiguration `pulumi:"containerNetworkInterfaceConfigurations"`
	ContainerNetworkInterfaces              []ContainerNetworkInterface              `pulumi:"containerNetworkInterfaces"`
	Id                                      *string                                  `pulumi:"id"`
	Location                                *string                                  `pulumi:"location"`
	NetworkProfileName                      *string                                  `pulumi:"networkProfileName"`
	ResourceGroupName                       string                                   `pulumi:"resourceGroupName"`
	Tags                                    map[string]string                        `pulumi:"tags"`
}


type NetworkProfileArgs struct {
	ContainerNetworkInterfaceConfigurations ContainerNetworkInterfaceConfigurationArrayInput
	ContainerNetworkInterfaces              ContainerNetworkInterfaceArrayInput
	Id                                      pulumi.StringPtrInput
	Location                                pulumi.StringPtrInput
	NetworkProfileName                      pulumi.StringPtrInput
	ResourceGroupName                       pulumi.StringInput
	Tags                                    pulumi.StringMapInput
}

func (NetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkProfileArgs)(nil)).Elem()
}

type NetworkProfileInput interface {
	pulumi.Input

	ToNetworkProfileOutput() NetworkProfileOutput
	ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput
}

func (*NetworkProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (i *NetworkProfile) ToNetworkProfileOutput() NetworkProfileOutput {
	return i.ToNetworkProfileOutputWithContext(context.Background())
}

func (i *NetworkProfile) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput)
}

type NetworkProfileOutput struct{ *pulumi.OutputState }

func (NetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (o NetworkProfileOutput) ToNetworkProfileOutput() NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ContainerNetworkInterfaceConfigurations() ContainerNetworkInterfaceConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *NetworkProfile) ContainerNetworkInterfaceConfigurationResponseArrayOutput {
		return v.ContainerNetworkInterfaceConfigurations
	}).(ContainerNetworkInterfaceConfigurationResponseArrayOutput)
}

func (o NetworkProfileOutput) ContainerNetworkInterfaces() ContainerNetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v *NetworkProfile) ContainerNetworkInterfaceResponseArrayOutput {
		return v.ContainerNetworkInterfaces
	}).(ContainerNetworkInterfaceResponseArrayOutput)
}

func (o NetworkProfileOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkProfileOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o NetworkProfileOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o NetworkProfileOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NetworkProfileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkProfileOutput{})
}
