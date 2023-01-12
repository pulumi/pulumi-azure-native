


package storagemover

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StorageMover struct {
	pulumi.CustomResourceState

	Description       pulumi.StringPtrOutput   `pulumi:"description"`
	Location          pulumi.StringOutput      `pulumi:"location"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput   `pulumi:"tags"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewStorageMover(ctx *pulumi.Context,
	name string, args *StorageMoverArgs, opts ...pulumi.ResourceOption) (*StorageMover, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storagemover/v20220701preview:StorageMover"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageMover
	err := ctx.RegisterResource("azure-native:storagemover:StorageMover", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStorageMover(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageMoverState, opts ...pulumi.ResourceOption) (*StorageMover, error) {
	var resource StorageMover
	err := ctx.ReadResource("azure-native:storagemover:StorageMover", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type storageMoverState struct {
}

type StorageMoverState struct {
}

func (StorageMoverState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageMoverState)(nil)).Elem()
}

type storageMoverArgs struct {
	Description       *string           `pulumi:"description"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	StorageMoverName  *string           `pulumi:"storageMoverName"`
	Tags              map[string]string `pulumi:"tags"`
}


type StorageMoverArgs struct {
	Description       pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	StorageMoverName  pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (StorageMoverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageMoverArgs)(nil)).Elem()
}

type StorageMoverInput interface {
	pulumi.Input

	ToStorageMoverOutput() StorageMoverOutput
	ToStorageMoverOutputWithContext(ctx context.Context) StorageMoverOutput
}

func (*StorageMover) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageMover)(nil)).Elem()
}

func (i *StorageMover) ToStorageMoverOutput() StorageMoverOutput {
	return i.ToStorageMoverOutputWithContext(context.Background())
}

func (i *StorageMover) ToStorageMoverOutputWithContext(ctx context.Context) StorageMoverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageMoverOutput)
}

type StorageMoverOutput struct{ *pulumi.OutputState }

func (StorageMoverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageMover)(nil)).Elem()
}

func (o StorageMoverOutput) ToStorageMoverOutput() StorageMoverOutput {
	return o
}

func (o StorageMoverOutput) ToStorageMoverOutputWithContext(ctx context.Context) StorageMoverOutput {
	return o
}

func (o StorageMoverOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageMover) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o StorageMoverOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageMover) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o StorageMoverOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageMover) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StorageMoverOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageMover) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o StorageMoverOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *StorageMover) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o StorageMoverOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StorageMover) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o StorageMoverOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageMover) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageMoverOutput{})
}
