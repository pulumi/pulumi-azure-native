


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Storage struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput          `pulumi:"name"`
	Properties StorageAccountResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput     `pulumi:"systemData"`
	Type       pulumi.StringOutput          `pulumi:"type"`
}


func NewStorage(ctx *pulumi.Context,
	name string, args *StorageArgs, opts ...pulumi.ResourceOption) (*Storage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:Storage"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210901preview:Storage"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:Storage"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:Storage"),
		},
	})
	opts = append(opts, aliases)
	var resource Storage
	err := ctx.RegisterResource("azure-native:appplatform/v20220101preview:Storage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStorage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageState, opts ...pulumi.ResourceOption) (*Storage, error) {
	var resource Storage
	err := ctx.ReadResource("azure-native:appplatform/v20220101preview:Storage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type storageState struct {
}

type StorageState struct {
}

func (StorageState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageState)(nil)).Elem()
}

type storageArgs struct {
	Properties        *StorageAccount `pulumi:"properties"`
	ResourceGroupName string          `pulumi:"resourceGroupName"`
	ServiceName       string          `pulumi:"serviceName"`
	StorageName       *string         `pulumi:"storageName"`
}


type StorageArgs struct {
	Properties        StorageAccountPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
	StorageName       pulumi.StringPtrInput
}

func (StorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageArgs)(nil)).Elem()
}

type StorageInput interface {
	pulumi.Input

	ToStorageOutput() StorageOutput
	ToStorageOutputWithContext(ctx context.Context) StorageOutput
}

func (*Storage) ElementType() reflect.Type {
	return reflect.TypeOf((**Storage)(nil)).Elem()
}

func (i *Storage) ToStorageOutput() StorageOutput {
	return i.ToStorageOutputWithContext(context.Background())
}

func (i *Storage) ToStorageOutputWithContext(ctx context.Context) StorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageOutput)
}

type StorageOutput struct{ *pulumi.OutputState }

func (StorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Storage)(nil)).Elem()
}

func (o StorageOutput) ToStorageOutput() StorageOutput {
	return o
}

func (o StorageOutput) ToStorageOutputWithContext(ctx context.Context) StorageOutput {
	return o
}

func (o StorageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Storage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StorageOutput) Properties() StorageAccountResponseOutput {
	return o.ApplyT(func(v *Storage) StorageAccountResponseOutput { return v.Properties }).(StorageAccountResponseOutput)
}

func (o StorageOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Storage) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o StorageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Storage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageOutput{})
}
