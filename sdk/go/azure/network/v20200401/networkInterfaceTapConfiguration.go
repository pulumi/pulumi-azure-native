


package v20200401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkInterfaceTapConfiguration struct {
	pulumi.CustomResourceState

	Etag              pulumi.StringOutput                `pulumi:"etag"`
	Name              pulumi.StringPtrOutput             `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                `pulumi:"provisioningState"`
	Type              pulumi.StringOutput                `pulumi:"type"`
	VirtualNetworkTap VirtualNetworkTapResponsePtrOutput `pulumi:"virtualNetworkTap"`
}


func NewNetworkInterfaceTapConfiguration(ctx *pulumi.Context,
	name string, args *NetworkInterfaceTapConfigurationArgs, opts ...pulumi.ResourceOption) (*NetworkInterfaceTapConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkInterfaceName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkInterfaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NetworkInterfaceTapConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:NetworkInterfaceTapConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:NetworkInterfaceTapConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:NetworkInterfaceTapConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:NetworkInterfaceTapConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:NetworkInterfaceTapConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:NetworkInterfaceTapConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:NetworkInterfaceTapConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:NetworkInterfaceTapConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:NetworkInterfaceTapConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:NetworkInterfaceTapConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:NetworkInterfaceTapConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:NetworkInterfaceTapConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:NetworkInterfaceTapConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:NetworkInterfaceTapConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:NetworkInterfaceTapConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:NetworkInterfaceTapConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:NetworkInterfaceTapConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:NetworkInterfaceTapConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:NetworkInterfaceTapConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:NetworkInterfaceTapConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:NetworkInterfaceTapConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkInterfaceTapConfiguration
	err := ctx.RegisterResource("azure-native:network/v20200401:NetworkInterfaceTapConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNetworkInterfaceTapConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceTapConfigurationState, opts ...pulumi.ResourceOption) (*NetworkInterfaceTapConfiguration, error) {
	var resource NetworkInterfaceTapConfiguration
	err := ctx.ReadResource("azure-native:network/v20200401:NetworkInterfaceTapConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type networkInterfaceTapConfigurationState struct {
}

type NetworkInterfaceTapConfigurationState struct {
}

func (NetworkInterfaceTapConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceTapConfigurationState)(nil)).Elem()
}

type networkInterfaceTapConfigurationArgs struct {
	Id                   *string                `pulumi:"id"`
	Name                 *string                `pulumi:"name"`
	NetworkInterfaceName string                 `pulumi:"networkInterfaceName"`
	ResourceGroupName    string                 `pulumi:"resourceGroupName"`
	TapConfigurationName *string                `pulumi:"tapConfigurationName"`
	VirtualNetworkTap    *VirtualNetworkTapType `pulumi:"virtualNetworkTap"`
}


type NetworkInterfaceTapConfigurationArgs struct {
	Id                   pulumi.StringPtrInput
	Name                 pulumi.StringPtrInput
	NetworkInterfaceName pulumi.StringInput
	ResourceGroupName    pulumi.StringInput
	TapConfigurationName pulumi.StringPtrInput
	VirtualNetworkTap    VirtualNetworkTapTypePtrInput
}

func (NetworkInterfaceTapConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceTapConfigurationArgs)(nil)).Elem()
}

type NetworkInterfaceTapConfigurationInput interface {
	pulumi.Input

	ToNetworkInterfaceTapConfigurationOutput() NetworkInterfaceTapConfigurationOutput
	ToNetworkInterfaceTapConfigurationOutputWithContext(ctx context.Context) NetworkInterfaceTapConfigurationOutput
}

func (*NetworkInterfaceTapConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceTapConfiguration)(nil))
}

func (i *NetworkInterfaceTapConfiguration) ToNetworkInterfaceTapConfigurationOutput() NetworkInterfaceTapConfigurationOutput {
	return i.ToNetworkInterfaceTapConfigurationOutputWithContext(context.Background())
}

func (i *NetworkInterfaceTapConfiguration) ToNetworkInterfaceTapConfigurationOutputWithContext(ctx context.Context) NetworkInterfaceTapConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceTapConfigurationOutput)
}

type NetworkInterfaceTapConfigurationOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceTapConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceTapConfiguration)(nil))
}

func (o NetworkInterfaceTapConfigurationOutput) ToNetworkInterfaceTapConfigurationOutput() NetworkInterfaceTapConfigurationOutput {
	return o
}

func (o NetworkInterfaceTapConfigurationOutput) ToNetworkInterfaceTapConfigurationOutputWithContext(ctx context.Context) NetworkInterfaceTapConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NetworkInterfaceTapConfigurationOutput{})
}
