


package v20170605preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type StorageSyncService struct {
	pulumi.CustomResourceState

	Location                 pulumi.StringPtrOutput `pulumi:"location"`
	Name                     pulumi.StringOutput    `pulumi:"name"`
	StorageSyncServiceStatus pulumi.IntOutput       `pulumi:"storageSyncServiceStatus"`
	StorageSyncServiceUid    pulumi.StringOutput    `pulumi:"storageSyncServiceUid"`
	Tags                     pulumi.AnyOutput       `pulumi:"tags"`
	Type                     pulumi.StringOutput    `pulumi:"type"`
}


func NewStorageSyncService(ctx *pulumi.Context,
	name string, args *StorageSyncServiceArgs, opts ...pulumi.ResourceOption) (*StorageSyncService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storagesync:StorageSyncService"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20180402:StorageSyncService"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20180701:StorageSyncService"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20181001:StorageSyncService"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190201:StorageSyncService"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190301:StorageSyncService"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20190601:StorageSyncService"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20191001:StorageSyncService"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20200301:StorageSyncService"),
		},
		{
			Type: pulumi.String("azure-native:storagesync/v20200901:StorageSyncService"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageSyncService
	err := ctx.RegisterResource("azure-native:storagesync/v20170605preview:StorageSyncService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStorageSyncService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageSyncServiceState, opts ...pulumi.ResourceOption) (*StorageSyncService, error) {
	var resource StorageSyncService
	err := ctx.ReadResource("azure-native:storagesync/v20170605preview:StorageSyncService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type storageSyncServiceState struct {
}

type StorageSyncServiceState struct {
}

func (StorageSyncServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageSyncServiceState)(nil)).Elem()
}

type storageSyncServiceArgs struct {
	Location               *string     `pulumi:"location"`
	ResourceGroupName      string      `pulumi:"resourceGroupName"`
	StorageSyncServiceName *string     `pulumi:"storageSyncServiceName"`
	Tags                   interface{} `pulumi:"tags"`
}


type StorageSyncServiceArgs struct {
	Location               pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	StorageSyncServiceName pulumi.StringPtrInput
	Tags                   pulumi.Input
}

func (StorageSyncServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageSyncServiceArgs)(nil)).Elem()
}

type StorageSyncServiceInput interface {
	pulumi.Input

	ToStorageSyncServiceOutput() StorageSyncServiceOutput
	ToStorageSyncServiceOutputWithContext(ctx context.Context) StorageSyncServiceOutput
}

func (*StorageSyncService) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSyncService)(nil)).Elem()
}

func (i *StorageSyncService) ToStorageSyncServiceOutput() StorageSyncServiceOutput {
	return i.ToStorageSyncServiceOutputWithContext(context.Background())
}

func (i *StorageSyncService) ToStorageSyncServiceOutputWithContext(ctx context.Context) StorageSyncServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSyncServiceOutput)
}

type StorageSyncServiceOutput struct{ *pulumi.OutputState }

func (StorageSyncServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSyncService)(nil)).Elem()
}

func (o StorageSyncServiceOutput) ToStorageSyncServiceOutput() StorageSyncServiceOutput {
	return o
}

func (o StorageSyncServiceOutput) ToStorageSyncServiceOutputWithContext(ctx context.Context) StorageSyncServiceOutput {
	return o
}

func (o StorageSyncServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageSyncService) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o StorageSyncServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageSyncService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StorageSyncServiceOutput) StorageSyncServiceStatus() pulumi.IntOutput {
	return o.ApplyT(func(v *StorageSyncService) pulumi.IntOutput { return v.StorageSyncServiceStatus }).(pulumi.IntOutput)
}

func (o StorageSyncServiceOutput) StorageSyncServiceUid() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageSyncService) pulumi.StringOutput { return v.StorageSyncServiceUid }).(pulumi.StringOutput)
}

func (o StorageSyncServiceOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *StorageSyncService) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

func (o StorageSyncServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageSyncService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageSyncServiceOutput{})
}
