


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
			Type: pulumi.String("azure-nextgen:connectedvmwarevsphere:InventoryItem"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20201001preview:InventoryItem"),
		},
		{
			Type: pulumi.String("azure-nextgen:connectedvmwarevsphere/v20201001preview:InventoryItem"),
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
	return reflect.TypeOf((*InventoryItem)(nil))
}

func (i *InventoryItem) ToInventoryItemOutput() InventoryItemOutput {
	return i.ToInventoryItemOutputWithContext(context.Background())
}

func (i *InventoryItem) ToInventoryItemOutputWithContext(ctx context.Context) InventoryItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InventoryItemOutput)
}

type InventoryItemOutput struct{ *pulumi.OutputState }

func (InventoryItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InventoryItem)(nil))
}

func (o InventoryItemOutput) ToInventoryItemOutput() InventoryItemOutput {
	return o
}

func (o InventoryItemOutput) ToInventoryItemOutputWithContext(ctx context.Context) InventoryItemOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InventoryItemInput)(nil)).Elem(), &InventoryItem{})
	pulumi.RegisterOutputType(InventoryItemOutput{})
}
