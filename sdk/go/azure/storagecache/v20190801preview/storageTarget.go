


package v20190801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StorageTarget struct {
	pulumi.CustomResourceState

	Clfs              ClfsTargetResponsePtrOutput          `pulumi:"clfs"`
	Junctions         NamespaceJunctionResponseArrayOutput `pulumi:"junctions"`
	Name              pulumi.StringOutput                  `pulumi:"name"`
	Nfs3              Nfs3TargetResponsePtrOutput          `pulumi:"nfs3"`
	ProvisioningState pulumi.StringPtrOutput               `pulumi:"provisioningState"`
	TargetType        pulumi.StringPtrOutput               `pulumi:"targetType"`
	Type              pulumi.StringOutput                  `pulumi:"type"`
	Unknown           UnknownTargetResponsePtrOutput       `pulumi:"unknown"`
}


func NewStorageTarget(ctx *pulumi.Context,
	name string, args *StorageTargetArgs, opts ...pulumi.ResourceOption) (*StorageTarget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CacheName == nil {
		return nil, errors.New("invalid value for required argument 'CacheName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storagecache:StorageTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20191101:StorageTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20200301:StorageTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20201001:StorageTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20210301:StorageTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20210501:StorageTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20210901:StorageTarget"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageTarget
	err := ctx.RegisterResource("azure-native:storagecache/v20190801preview:StorageTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStorageTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageTargetState, opts ...pulumi.ResourceOption) (*StorageTarget, error) {
	var resource StorageTarget
	err := ctx.ReadResource("azure-native:storagecache/v20190801preview:StorageTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type storageTargetState struct {
}

type StorageTargetState struct {
}

func (StorageTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageTargetState)(nil)).Elem()
}

type storageTargetArgs struct {
	CacheName         string              `pulumi:"cacheName"`
	Clfs              *ClfsTarget         `pulumi:"clfs"`
	Junctions         []NamespaceJunction `pulumi:"junctions"`
	Nfs3              *Nfs3Target         `pulumi:"nfs3"`
	ProvisioningState *string             `pulumi:"provisioningState"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
	StorageTargetName *string             `pulumi:"storageTargetName"`
	TargetType        *string             `pulumi:"targetType"`
	Unknown           *UnknownTarget      `pulumi:"unknown"`
}


type StorageTargetArgs struct {
	CacheName         pulumi.StringInput
	Clfs              ClfsTargetPtrInput
	Junctions         NamespaceJunctionArrayInput
	Nfs3              Nfs3TargetPtrInput
	ProvisioningState pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	StorageTargetName pulumi.StringPtrInput
	TargetType        pulumi.StringPtrInput
	Unknown           UnknownTargetPtrInput
}

func (StorageTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageTargetArgs)(nil)).Elem()
}

type StorageTargetInput interface {
	pulumi.Input

	ToStorageTargetOutput() StorageTargetOutput
	ToStorageTargetOutputWithContext(ctx context.Context) StorageTargetOutput
}

func (*StorageTarget) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageTarget)(nil)).Elem()
}

func (i *StorageTarget) ToStorageTargetOutput() StorageTargetOutput {
	return i.ToStorageTargetOutputWithContext(context.Background())
}

func (i *StorageTarget) ToStorageTargetOutputWithContext(ctx context.Context) StorageTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageTargetOutput)
}

type StorageTargetOutput struct{ *pulumi.OutputState }

func (StorageTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageTarget)(nil)).Elem()
}

func (o StorageTargetOutput) ToStorageTargetOutput() StorageTargetOutput {
	return o
}

func (o StorageTargetOutput) ToStorageTargetOutputWithContext(ctx context.Context) StorageTargetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(StorageTargetOutput{})
}
