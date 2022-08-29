


package connectedvmwarevsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InventoryItem struct {
	pulumi.CustomResourceState

	InventoryType     pulumi.StringOutput      `pulumi:"inventoryType"`
	Kind              pulumi.StringPtrOutput   `pulumi:"kind"`
	ManagedResourceId pulumi.StringPtrOutput   `pulumi:"managedResourceId"`
	MoName            pulumi.StringPtrOutput   `pulumi:"moName"`
	MoRefId           pulumi.StringPtrOutput   `pulumi:"moRefId"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewInventoryItem(ctx *pulumi.Context,
	name string, args *InventoryItemArgs, opts ...pulumi.ResourceOption) (*InventoryItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InventoryType == nil {
		return nil, errors.New("invalid value for required argument 'InventoryType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VcenterName == nil {
		return nil, errors.New("invalid value for required argument 'VcenterName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20201001preview:InventoryItem"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220110preview:InventoryItem"),
		},
	})
	opts = append(opts, aliases)
	var resource InventoryItem
	err := ctx.RegisterResource("azure-native:connectedvmwarevsphere:InventoryItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetInventoryItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InventoryItemState, opts ...pulumi.ResourceOption) (*InventoryItem, error) {
	var resource InventoryItem
	err := ctx.ReadResource("azure-native:connectedvmwarevsphere:InventoryItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type inventoryItemState struct {
}

type InventoryItemState struct {
}

func (InventoryItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*inventoryItemState)(nil)).Elem()
}

type inventoryItemArgs struct {
	InventoryItemName *string `pulumi:"inventoryItemName"`
	InventoryType     string  `pulumi:"inventoryType"`
	Kind              *string `pulumi:"kind"`
	ManagedResourceId *string `pulumi:"managedResourceId"`
	MoName            *string `pulumi:"moName"`
	MoRefId           *string `pulumi:"moRefId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	VcenterName       string  `pulumi:"vcenterName"`
}


type InventoryItemArgs struct {
	InventoryItemName pulumi.StringPtrInput
	InventoryType     pulumi.StringInput
	Kind              pulumi.StringPtrInput
	ManagedResourceId pulumi.StringPtrInput
	MoName            pulumi.StringPtrInput
	MoRefId           pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	VcenterName       pulumi.StringInput
}

func (InventoryItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inventoryItemArgs)(nil)).Elem()
}

type InventoryItemInput interface {
	pulumi.Input

	ToInventoryItemOutput() InventoryItemOutput
	ToInventoryItemOutputWithContext(ctx context.Context) InventoryItemOutput
}

func (*InventoryItem) ElementType() reflect.Type {
	return reflect.TypeOf((**InventoryItem)(nil)).Elem()
}

func (i *InventoryItem) ToInventoryItemOutput() InventoryItemOutput {
	return i.ToInventoryItemOutputWithContext(context.Background())
}

func (i *InventoryItem) ToInventoryItemOutputWithContext(ctx context.Context) InventoryItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InventoryItemOutput)
}

type InventoryItemOutput struct{ *pulumi.OutputState }

func (InventoryItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InventoryItem)(nil)).Elem()
}

func (o InventoryItemOutput) ToInventoryItemOutput() InventoryItemOutput {
	return o
}

func (o InventoryItemOutput) ToInventoryItemOutputWithContext(ctx context.Context) InventoryItemOutput {
	return o
}

func (o InventoryItemOutput) InventoryType() pulumi.StringOutput {
	return o.ApplyT(func(v *InventoryItem) pulumi.StringOutput { return v.InventoryType }).(pulumi.StringOutput)
}

func (o InventoryItemOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InventoryItem) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o InventoryItemOutput) ManagedResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InventoryItem) pulumi.StringPtrOutput { return v.ManagedResourceId }).(pulumi.StringPtrOutput)
}

func (o InventoryItemOutput) MoName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InventoryItem) pulumi.StringPtrOutput { return v.MoName }).(pulumi.StringPtrOutput)
}

func (o InventoryItemOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InventoryItem) pulumi.StringPtrOutput { return v.MoRefId }).(pulumi.StringPtrOutput)
}

func (o InventoryItemOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InventoryItem) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o InventoryItemOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *InventoryItem) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o InventoryItemOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *InventoryItem) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o InventoryItemOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *InventoryItem) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(InventoryItemOutput{})
}
