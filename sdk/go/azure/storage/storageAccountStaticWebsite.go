


package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StorageAccountStaticWebsite struct {
	pulumi.CustomResourceState

	ContainerName    pulumi.StringOutput    `pulumi:"containerName"`
	Error404Document pulumi.StringPtrOutput `pulumi:"error404Document"`
	IndexDocument    pulumi.StringPtrOutput `pulumi:"indexDocument"`
}


func NewStorageAccountStaticWebsite(ctx *pulumi.Context,
	name string, args *StorageAccountStaticWebsiteArgs, opts ...pulumi.ResourceOption) (*StorageAccountStaticWebsite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource StorageAccountStaticWebsite
	err := ctx.RegisterResource("azure-native:storage:StorageAccountStaticWebsite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStorageAccountStaticWebsite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageAccountStaticWebsiteState, opts ...pulumi.ResourceOption) (*StorageAccountStaticWebsite, error) {
	var resource StorageAccountStaticWebsite
	err := ctx.ReadResource("azure-native:storage:StorageAccountStaticWebsite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type storageAccountStaticWebsiteState struct {
}

type StorageAccountStaticWebsiteState struct {
}

func (StorageAccountStaticWebsiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAccountStaticWebsiteState)(nil)).Elem()
}

type storageAccountStaticWebsiteArgs struct {
	AccountName       string  `pulumi:"accountName"`
	Error404Document  *string `pulumi:"error404Document"`
	IndexDocument     *string `pulumi:"indexDocument"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type StorageAccountStaticWebsiteArgs struct {
	AccountName       pulumi.StringInput
	Error404Document  pulumi.StringPtrInput
	IndexDocument     pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
}

func (StorageAccountStaticWebsiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAccountStaticWebsiteArgs)(nil)).Elem()
}

type StorageAccountStaticWebsiteInput interface {
	pulumi.Input

	ToStorageAccountStaticWebsiteOutput() StorageAccountStaticWebsiteOutput
	ToStorageAccountStaticWebsiteOutputWithContext(ctx context.Context) StorageAccountStaticWebsiteOutput
}

func (*StorageAccountStaticWebsite) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountStaticWebsite)(nil))
}

func (i *StorageAccountStaticWebsite) ToStorageAccountStaticWebsiteOutput() StorageAccountStaticWebsiteOutput {
	return i.ToStorageAccountStaticWebsiteOutputWithContext(context.Background())
}

func (i *StorageAccountStaticWebsite) ToStorageAccountStaticWebsiteOutputWithContext(ctx context.Context) StorageAccountStaticWebsiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountStaticWebsiteOutput)
}

type StorageAccountStaticWebsiteOutput struct{ *pulumi.OutputState }

func (StorageAccountStaticWebsiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountStaticWebsite)(nil))
}

func (o StorageAccountStaticWebsiteOutput) ToStorageAccountStaticWebsiteOutput() StorageAccountStaticWebsiteOutput {
	return o
}

func (o StorageAccountStaticWebsiteOutput) ToStorageAccountStaticWebsiteOutputWithContext(ctx context.Context) StorageAccountStaticWebsiteOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAccountStaticWebsiteInput)(nil)).Elem(), &StorageAccountStaticWebsite{})
	pulumi.RegisterOutputType(StorageAccountStaticWebsiteOutput{})
}
