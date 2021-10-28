


package v20180915

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualNetwork struct {
	pulumi.CustomResourceState

	AllowedSubnets             SubnetResponseArrayOutput         `pulumi:"allowedSubnets"`
	CreatedDate                pulumi.StringOutput               `pulumi:"createdDate"`
	Description                pulumi.StringPtrOutput            `pulumi:"description"`
	ExternalProviderResourceId pulumi.StringPtrOutput            `pulumi:"externalProviderResourceId"`
	ExternalSubnets            ExternalSubnetResponseArrayOutput `pulumi:"externalSubnets"`
	Location                   pulumi.StringPtrOutput            `pulumi:"location"`
	Name                       pulumi.StringOutput               `pulumi:"name"`
	ProvisioningState          pulumi.StringOutput               `pulumi:"provisioningState"`
	SubnetOverrides            SubnetOverrideResponseArrayOutput `pulumi:"subnetOverrides"`
	Tags                       pulumi.StringMapOutput            `pulumi:"tags"`
	Type                       pulumi.StringOutput               `pulumi:"type"`
	UniqueIdentifier           pulumi.StringOutput               `pulumi:"uniqueIdentifier"`
}


func NewVirtualNetwork(ctx *pulumi.Context,
	name string, args *VirtualNetworkArgs, opts ...pulumi.ResourceOption) (*VirtualNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20180915:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20150521preview:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20150521preview:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20160515:VirtualNetwork"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualNetwork
	err := ctx.RegisterResource("azure-native:devtestlab/v20180915:VirtualNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkState, opts ...pulumi.ResourceOption) (*VirtualNetwork, error) {
	var resource VirtualNetwork
	err := ctx.ReadResource("azure-native:devtestlab/v20180915:VirtualNetwork", name, id, state, &resource, opts...)
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
	AllowedSubnets             []Subnet          `pulumi:"allowedSubnets"`
	Description                *string           `pulumi:"description"`
	ExternalProviderResourceId *string           `pulumi:"externalProviderResourceId"`
	LabName                    string            `pulumi:"labName"`
	Location                   *string           `pulumi:"location"`
	Name                       *string           `pulumi:"name"`
	ResourceGroupName          string            `pulumi:"resourceGroupName"`
	SubnetOverrides            []SubnetOverride  `pulumi:"subnetOverrides"`
	Tags                       map[string]string `pulumi:"tags"`
}


type VirtualNetworkArgs struct {
	AllowedSubnets             SubnetArrayInput
	Description                pulumi.StringPtrInput
	ExternalProviderResourceId pulumi.StringPtrInput
	LabName                    pulumi.StringInput
	Location                   pulumi.StringPtrInput
	Name                       pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	SubnetOverrides            SubnetOverrideArrayInput
	Tags                       pulumi.StringMapInput
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
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNetworkInput)(nil)).Elem(), &VirtualNetwork{})
	pulumi.RegisterOutputType(VirtualNetworkOutput{})
}
