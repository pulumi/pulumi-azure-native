


package v20150521preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type VirtualNetworkResource struct {
	pulumi.CustomResourceState

	AllowedSubnets             SubnetResponseArrayOutput         `pulumi:"allowedSubnets"`
	Description                pulumi.StringPtrOutput            `pulumi:"description"`
	ExternalProviderResourceId pulumi.StringPtrOutput            `pulumi:"externalProviderResourceId"`
	Location                   pulumi.StringPtrOutput            `pulumi:"location"`
	Name                       pulumi.StringPtrOutput            `pulumi:"name"`
	ProvisioningState          pulumi.StringPtrOutput            `pulumi:"provisioningState"`
	SubnetOverrides            SubnetOverrideResponseArrayOutput `pulumi:"subnetOverrides"`
	Tags                       pulumi.StringMapOutput            `pulumi:"tags"`
	Type                       pulumi.StringPtrOutput            `pulumi:"type"`
}


func NewVirtualNetworkResource(ctx *pulumi.Context,
	name string, args *VirtualNetworkResourceArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkResource, error) {
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
			Type: pulumi.String("azure-native:devtestlab:VirtualNetworkResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:VirtualNetworkResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:VirtualNetworkResource"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualNetworkResource
	err := ctx.RegisterResource("azure-native:devtestlab/v20150521preview:VirtualNetworkResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualNetworkResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkResourceState, opts ...pulumi.ResourceOption) (*VirtualNetworkResource, error) {
	var resource VirtualNetworkResource
	err := ctx.ReadResource("azure-native:devtestlab/v20150521preview:VirtualNetworkResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualNetworkResourceState struct {
}

type VirtualNetworkResourceState struct {
}

func (VirtualNetworkResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkResourceState)(nil)).Elem()
}

type virtualNetworkResourceArgs struct {
	AllowedSubnets             []Subnet          `pulumi:"allowedSubnets"`
	Description                *string           `pulumi:"description"`
	ExternalProviderResourceId *string           `pulumi:"externalProviderResourceId"`
	Id                         *string           `pulumi:"id"`
	LabName                    string            `pulumi:"labName"`
	Location                   *string           `pulumi:"location"`
	Name                       *string           `pulumi:"name"`
	ProvisioningState          *string           `pulumi:"provisioningState"`
	ResourceGroupName          string            `pulumi:"resourceGroupName"`
	SubnetOverrides            []SubnetOverride  `pulumi:"subnetOverrides"`
	Tags                       map[string]string `pulumi:"tags"`
	Type                       *string           `pulumi:"type"`
}


type VirtualNetworkResourceArgs struct {
	AllowedSubnets             SubnetArrayInput
	Description                pulumi.StringPtrInput
	ExternalProviderResourceId pulumi.StringPtrInput
	Id                         pulumi.StringPtrInput
	LabName                    pulumi.StringInput
	Location                   pulumi.StringPtrInput
	Name                       pulumi.StringPtrInput
	ProvisioningState          pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	SubnetOverrides            SubnetOverrideArrayInput
	Tags                       pulumi.StringMapInput
	Type                       pulumi.StringPtrInput
}

func (VirtualNetworkResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkResourceArgs)(nil)).Elem()
}

type VirtualNetworkResourceInput interface {
	pulumi.Input

	ToVirtualNetworkResourceOutput() VirtualNetworkResourceOutput
	ToVirtualNetworkResourceOutputWithContext(ctx context.Context) VirtualNetworkResourceOutput
}

func (*VirtualNetworkResource) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkResource)(nil)).Elem()
}

func (i *VirtualNetworkResource) ToVirtualNetworkResourceOutput() VirtualNetworkResourceOutput {
	return i.ToVirtualNetworkResourceOutputWithContext(context.Background())
}

func (i *VirtualNetworkResource) ToVirtualNetworkResourceOutputWithContext(ctx context.Context) VirtualNetworkResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkResourceOutput)
}

type VirtualNetworkResourceOutput struct{ *pulumi.OutputState }

func (VirtualNetworkResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkResource)(nil)).Elem()
}

func (o VirtualNetworkResourceOutput) ToVirtualNetworkResourceOutput() VirtualNetworkResourceOutput {
	return o
}

func (o VirtualNetworkResourceOutput) ToVirtualNetworkResourceOutputWithContext(ctx context.Context) VirtualNetworkResourceOutput {
	return o
}

func (o VirtualNetworkResourceOutput) AllowedSubnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetworkResource) SubnetResponseArrayOutput { return v.AllowedSubnets }).(SubnetResponseArrayOutput)
}

func (o VirtualNetworkResourceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkResource) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkResourceOutput) ExternalProviderResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkResource) pulumi.StringPtrOutput { return v.ExternalProviderResourceId }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkResourceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkResource) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkResourceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkResource) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkResourceOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkResource) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkResourceOutput) SubnetOverrides() SubnetOverrideResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetworkResource) SubnetOverrideResponseArrayOutput { return v.SubnetOverrides }).(SubnetOverrideResponseArrayOutput)
}

func (o VirtualNetworkResourceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetworkResource) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualNetworkResourceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkResource) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkResourceOutput{})
}
