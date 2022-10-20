


package v20160515

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
	ProvisioningState          pulumi.StringPtrOutput            `pulumi:"provisioningState"`
	SubnetOverrides            SubnetOverrideResponseArrayOutput `pulumi:"subnetOverrides"`
	Tags                       pulumi.StringMapOutput            `pulumi:"tags"`
	Type                       pulumi.StringOutput               `pulumi:"type"`
	UniqueIdentifier           pulumi.StringPtrOutput            `pulumi:"uniqueIdentifier"`
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
			Type: pulumi.String("azure-native:devtestlab:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20150521preview:VirtualNetwork"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:VirtualNetwork"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualNetwork
	err := ctx.RegisterResource("azure-native:devtestlab/v20160515:VirtualNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkState, opts ...pulumi.ResourceOption) (*VirtualNetwork, error) {
	var resource VirtualNetwork
	err := ctx.ReadResource("azure-native:devtestlab/v20160515:VirtualNetwork", name, id, state, &resource, opts...)
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
	ExternalSubnets            []ExternalSubnet  `pulumi:"externalSubnets"`
	LabName                    string            `pulumi:"labName"`
	Location                   *string           `pulumi:"location"`
	Name                       *string           `pulumi:"name"`
	ProvisioningState          *string           `pulumi:"provisioningState"`
	ResourceGroupName          string            `pulumi:"resourceGroupName"`
	SubnetOverrides            []SubnetOverride  `pulumi:"subnetOverrides"`
	Tags                       map[string]string `pulumi:"tags"`
	UniqueIdentifier           *string           `pulumi:"uniqueIdentifier"`
}


type VirtualNetworkArgs struct {
	AllowedSubnets             SubnetArrayInput
	Description                pulumi.StringPtrInput
	ExternalProviderResourceId pulumi.StringPtrInput
	ExternalSubnets            ExternalSubnetArrayInput
	LabName                    pulumi.StringInput
	Location                   pulumi.StringPtrInput
	Name                       pulumi.StringPtrInput
	ProvisioningState          pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	SubnetOverrides            SubnetOverrideArrayInput
	Tags                       pulumi.StringMapInput
	UniqueIdentifier           pulumi.StringPtrInput
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
	return reflect.TypeOf((**VirtualNetwork)(nil)).Elem()
}

func (i *VirtualNetwork) ToVirtualNetworkOutput() VirtualNetworkOutput {
	return i.ToVirtualNetworkOutputWithContext(context.Background())
}

func (i *VirtualNetwork) ToVirtualNetworkOutputWithContext(ctx context.Context) VirtualNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkOutput)
}

type VirtualNetworkOutput struct{ *pulumi.OutputState }

func (VirtualNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetwork)(nil)).Elem()
}

func (o VirtualNetworkOutput) ToVirtualNetworkOutput() VirtualNetworkOutput {
	return o
}

func (o VirtualNetworkOutput) ToVirtualNetworkOutputWithContext(ctx context.Context) VirtualNetworkOutput {
	return o
}

func (o VirtualNetworkOutput) AllowedSubnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetwork) SubnetResponseArrayOutput { return v.AllowedSubnets }).(SubnetResponseArrayOutput)
}

func (o VirtualNetworkOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o VirtualNetworkOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkOutput) ExternalProviderResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringPtrOutput { return v.ExternalProviderResourceId }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkOutput) ExternalSubnets() ExternalSubnetResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetwork) ExternalSubnetResponseArrayOutput { return v.ExternalSubnets }).(ExternalSubnetResponseArrayOutput)
}

func (o VirtualNetworkOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualNetworkOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkOutput) SubnetOverrides() SubnetOverrideResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetwork) SubnetOverrideResponseArrayOutput { return v.SubnetOverrides }).(SubnetOverrideResponseArrayOutput)
}

func (o VirtualNetworkOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualNetworkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualNetworkOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringPtrOutput { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkOutput{})
}
