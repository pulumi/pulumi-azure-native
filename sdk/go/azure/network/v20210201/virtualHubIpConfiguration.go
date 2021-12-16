


package v20210201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualHubIpConfiguration struct {
	pulumi.CustomResourceState

	Etag                      pulumi.StringOutput              `pulumi:"etag"`
	Name                      pulumi.StringPtrOutput           `pulumi:"name"`
	PrivateIPAddress          pulumi.StringPtrOutput           `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod pulumi.StringPtrOutput           `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         pulumi.StringOutput              `pulumi:"provisioningState"`
	PublicIPAddress           PublicIPAddressResponsePtrOutput `pulumi:"publicIPAddress"`
	Subnet                    SubnetResponsePtrOutput          `pulumi:"subnet"`
	Type                      pulumi.StringOutput              `pulumi:"type"`
}


func NewVirtualHubIpConfiguration(ctx *pulumi.Context,
	name string, args *VirtualHubIpConfigurationArgs, opts ...pulumi.ResourceOption) (*VirtualHubIpConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualHubName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualHubName'")
	}
	subnetApplier := func(v SubnetType) *SubnetType { return v.Defaults() }
	if args.Subnet != nil {
		args.Subnet = args.Subnet.ToSubnetTypePtrOutput().Elem().ApplyT(subnetApplier).(SubnetTypePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:VirtualHubIpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualHubIpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualHubIpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualHubIpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualHubIpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualHubIpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualHubIpConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualHubIpConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualHubIpConfiguration
	err := ctx.RegisterResource("azure-native:network/v20210201:VirtualHubIpConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualHubIpConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualHubIpConfigurationState, opts ...pulumi.ResourceOption) (*VirtualHubIpConfiguration, error) {
	var resource VirtualHubIpConfiguration
	err := ctx.ReadResource("azure-native:network/v20210201:VirtualHubIpConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualHubIpConfigurationState struct {
}

type VirtualHubIpConfigurationState struct {
}

func (VirtualHubIpConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubIpConfigurationState)(nil)).Elem()
}

type virtualHubIpConfigurationArgs struct {
	Id                        *string              `pulumi:"id"`
	IpConfigName              *string              `pulumi:"ipConfigName"`
	Name                      *string              `pulumi:"name"`
	PrivateIPAddress          *string              `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod *string              `pulumi:"privateIPAllocationMethod"`
	PublicIPAddress           *PublicIPAddressType `pulumi:"publicIPAddress"`
	ResourceGroupName         string               `pulumi:"resourceGroupName"`
	Subnet                    *SubnetType          `pulumi:"subnet"`
	VirtualHubName            string               `pulumi:"virtualHubName"`
}


type VirtualHubIpConfigurationArgs struct {
	Id                        pulumi.StringPtrInput
	IpConfigName              pulumi.StringPtrInput
	Name                      pulumi.StringPtrInput
	PrivateIPAddress          pulumi.StringPtrInput
	PrivateIPAllocationMethod pulumi.StringPtrInput
	PublicIPAddress           PublicIPAddressTypePtrInput
	ResourceGroupName         pulumi.StringInput
	Subnet                    SubnetTypePtrInput
	VirtualHubName            pulumi.StringInput
}

func (VirtualHubIpConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubIpConfigurationArgs)(nil)).Elem()
}

type VirtualHubIpConfigurationInput interface {
	pulumi.Input

	ToVirtualHubIpConfigurationOutput() VirtualHubIpConfigurationOutput
	ToVirtualHubIpConfigurationOutputWithContext(ctx context.Context) VirtualHubIpConfigurationOutput
}

func (*VirtualHubIpConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHubIpConfiguration)(nil)).Elem()
}

func (i *VirtualHubIpConfiguration) ToVirtualHubIpConfigurationOutput() VirtualHubIpConfigurationOutput {
	return i.ToVirtualHubIpConfigurationOutputWithContext(context.Background())
}

func (i *VirtualHubIpConfiguration) ToVirtualHubIpConfigurationOutputWithContext(ctx context.Context) VirtualHubIpConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubIpConfigurationOutput)
}

type VirtualHubIpConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualHubIpConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHubIpConfiguration)(nil)).Elem()
}

func (o VirtualHubIpConfigurationOutput) ToVirtualHubIpConfigurationOutput() VirtualHubIpConfigurationOutput {
	return o
}

func (o VirtualHubIpConfigurationOutput) ToVirtualHubIpConfigurationOutputWithContext(ctx context.Context) VirtualHubIpConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VirtualHubIpConfigurationOutput{})
}
