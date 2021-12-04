


package v20181201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualNetwork struct {
	pulumi.CustomResourceState

	AddressSpace           AddressSpaceResponsePtrOutput            `pulumi:"addressSpace"`
	DdosProtectionPlan     SubResourceResponsePtrOutput             `pulumi:"ddosProtectionPlan"`
	DhcpOptions            DhcpOptionsResponsePtrOutput             `pulumi:"dhcpOptions"`
	EnableDdosProtection   pulumi.BoolPtrOutput                     `pulumi:"enableDdosProtection"`
	EnableVmProtection     pulumi.BoolPtrOutput                     `pulumi:"enableVmProtection"`
	Etag                   pulumi.StringPtrOutput                   `pulumi:"etag"`
	Location               pulumi.StringPtrOutput                   `pulumi:"location"`
	Name                   pulumi.StringOutput                      `pulumi:"name"`
	ProvisioningState      pulumi.StringPtrOutput                   `pulumi:"provisioningState"`
	ResourceGuid           pulumi.StringPtrOutput                   `pulumi:"resourceGuid"`
	Subnets                SubnetResponseArrayOutput                `pulumi:"subnets"`
	Tags                   pulumi.StringMapOutput                   `pulumi:"tags"`
	Type                   pulumi.StringOutput                      `pulumi:"type"`
	VirtualNetworkPeerings VirtualNetworkPeeringResponseArrayOutput `pulumi:"virtualNetworkPeerings"`
}


func NewVirtualNetwork(ctx *pulumi.Context,
	name string, args *VirtualNetworkArgs, opts ...pulumi.ResourceOption) (*VirtualNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.EnableDdosProtection == nil {
		args.EnableDdosProtection = pulumi.BoolPtr(false)
	}
	if args.EnableVmProtection == nil {
		args.EnableVmProtection = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150501preview:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualNetwork"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualNetwork
	err := ctx.RegisterResource("azure-native:network/v20181201:VirtualNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkState, opts ...pulumi.ResourceOption) (*VirtualNetwork, error) {
	var resource VirtualNetwork
	err := ctx.ReadResource("azure-native:network/v20181201:VirtualNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualNetworkState struct {
}

type VirtualNetworkState struct {
}

func (VirtualNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkState)(nil)).Elem()
}

type virtualNetworkArgs struct {
	AddressSpace           *AddressSpace               `pulumi:"addressSpace"`
	DdosProtectionPlan     *SubResource                `pulumi:"ddosProtectionPlan"`
	DhcpOptions            *DhcpOptions                `pulumi:"dhcpOptions"`
	EnableDdosProtection   *bool                       `pulumi:"enableDdosProtection"`
	EnableVmProtection     *bool                       `pulumi:"enableVmProtection"`
	Etag                   *string                     `pulumi:"etag"`
	Id                     *string                     `pulumi:"id"`
	Location               *string                     `pulumi:"location"`
	ProvisioningState      *string                     `pulumi:"provisioningState"`
	ResourceGroupName      string                      `pulumi:"resourceGroupName"`
	ResourceGuid           *string                     `pulumi:"resourceGuid"`
	Subnets                []SubnetType                `pulumi:"subnets"`
	Tags                   map[string]string           `pulumi:"tags"`
	VirtualNetworkName     *string                     `pulumi:"virtualNetworkName"`
	VirtualNetworkPeerings []VirtualNetworkPeeringType `pulumi:"virtualNetworkPeerings"`
}


type VirtualNetworkArgs struct {
	AddressSpace           AddressSpacePtrInput
	DdosProtectionPlan     SubResourcePtrInput
	DhcpOptions            DhcpOptionsPtrInput
	EnableDdosProtection   pulumi.BoolPtrInput
	EnableVmProtection     pulumi.BoolPtrInput
	Etag                   pulumi.StringPtrInput
	Id                     pulumi.StringPtrInput
	Location               pulumi.StringPtrInput
	ProvisioningState      pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	ResourceGuid           pulumi.StringPtrInput
	Subnets                SubnetTypeArrayInput
	Tags                   pulumi.StringMapInput
	VirtualNetworkName     pulumi.StringPtrInput
	VirtualNetworkPeerings VirtualNetworkPeeringTypeArrayInput
}

func (VirtualNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkArgs)(nil)).Elem()
}

type VirtualNetworkInput interface {
	pulumi.Input

	ToVirtualNetworkOutput() VirtualNetworkOutput
	ToVirtualNetworkOutputWithContext(ctx context.Context) VirtualNetworkOutput
}

func (*VirtualNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetwork)(nil))
}

func (i *VirtualNetwork) ToVirtualNetworkOutput() VirtualNetworkOutput {
	return i.ToVirtualNetworkOutputWithContext(context.Background())
}

func (i *VirtualNetwork) ToVirtualNetworkOutputWithContext(ctx context.Context) VirtualNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkOutput)
}

type VirtualNetworkOutput struct{ *pulumi.OutputState }

func (VirtualNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetwork)(nil))
}

func (o VirtualNetworkOutput) ToVirtualNetworkOutput() VirtualNetworkOutput {
	return o
}

func (o VirtualNetworkOutput) ToVirtualNetworkOutputWithContext(ctx context.Context) VirtualNetworkOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkOutput{})
}
