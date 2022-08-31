


package v20200605preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualNetwork struct {
	pulumi.CustomResourceState

	ExtendedLocation  ExtendedLocationResponseOutput `pulumi:"extendedLocation"`
	InventoryItemId   pulumi.StringPtrOutput         `pulumi:"inventoryItemId"`
	Location          pulumi.StringOutput            `pulumi:"location"`
	Name              pulumi.StringOutput            `pulumi:"name"`
	NetworkName       pulumi.StringOutput            `pulumi:"networkName"`
	ProvisioningState pulumi.StringOutput            `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput       `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput         `pulumi:"tags"`
	Type              pulumi.StringOutput            `pulumi:"type"`
	Uuid              pulumi.StringPtrOutput         `pulumi:"uuid"`
	VmmServerId       pulumi.StringPtrOutput         `pulumi:"vmmServerId"`
}


func NewVirtualNetwork(ctx *pulumi.Context,
	name string, args *VirtualNetworkArgs, opts ...pulumi.ResourceOption) (*VirtualNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:scvmm:VirtualNetwork"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualNetwork
	err := ctx.RegisterResource("azure-native:scvmm/v20200605preview:VirtualNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkState, opts ...pulumi.ResourceOption) (*VirtualNetwork, error) {
	var resource VirtualNetwork
	err := ctx.ReadResource("azure-native:scvmm/v20200605preview:VirtualNetwork", name, id, state, &resource, opts...)
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
	ExtendedLocation   ExtendedLocation  `pulumi:"extendedLocation"`
	InventoryItemId    *string           `pulumi:"inventoryItemId"`
	Location           *string           `pulumi:"location"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	Tags               map[string]string `pulumi:"tags"`
	Uuid               *string           `pulumi:"uuid"`
	VirtualNetworkName *string           `pulumi:"virtualNetworkName"`
	VmmServerId        *string           `pulumi:"vmmServerId"`
}


type VirtualNetworkArgs struct {
	ExtendedLocation   ExtendedLocationInput
	InventoryItemId    pulumi.StringPtrInput
	Location           pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	Tags               pulumi.StringMapInput
	Uuid               pulumi.StringPtrInput
	VirtualNetworkName pulumi.StringPtrInput
	VmmServerId        pulumi.StringPtrInput
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

func (o VirtualNetworkOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *VirtualNetwork) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

func (o VirtualNetworkOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringPtrOutput { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VirtualNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualNetworkOutput) NetworkName() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.NetworkName }).(pulumi.StringOutput)
}

func (o VirtualNetworkOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualNetworkOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VirtualNetwork) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o VirtualNetworkOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualNetworkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualNetworkOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringPtrOutput { return v.Uuid }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkOutput) VmmServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetwork) pulumi.StringPtrOutput { return v.VmmServerId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkOutput{})
}
