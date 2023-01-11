


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StoragecontainerRetrieve struct {
	pulumi.CustomResourceState

	AvailableSizeMB   pulumi.Float64Output                               `pulumi:"availableSizeMB"`
	ContainerSizeMB   pulumi.Float64Output                               `pulumi:"containerSizeMB"`
	ExtendedLocation  StoragecontainersResponseExtendedLocationPtrOutput `pulumi:"extendedLocation"`
	Location          pulumi.StringOutput                                `pulumi:"location"`
	Name              pulumi.StringOutput                                `pulumi:"name"`
	Path              pulumi.StringPtrOutput                             `pulumi:"path"`
	ProvisioningState pulumi.StringPtrOutput                             `pulumi:"provisioningState"`
	ResourceName      pulumi.StringPtrOutput                             `pulumi:"resourceName"`
	Status            StorageContainerStatusResponseOutput               `pulumi:"status"`
	SystemData        SystemDataResponseOutput                           `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                             `pulumi:"tags"`
	Type              pulumi.StringOutput                                `pulumi:"type"`
}


func NewStoragecontainerRetrieve(ctx *pulumi.Context,
	name string, args *StoragecontainerRetrieveArgs, opts ...pulumi.ResourceOption) (*StoragecontainerRetrieve, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource StoragecontainerRetrieve
	err := ctx.RegisterResource("azure-native:azurestackhci/v20210901preview:storagecontainerRetrieve", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStoragecontainerRetrieve(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StoragecontainerRetrieveState, opts ...pulumi.ResourceOption) (*StoragecontainerRetrieve, error) {
	var resource StoragecontainerRetrieve
	err := ctx.ReadResource("azure-native:azurestackhci/v20210901preview:storagecontainerRetrieve", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type storagecontainerRetrieveState struct {
}

type StoragecontainerRetrieveState struct {
}

func (StoragecontainerRetrieveState) ElementType() reflect.Type {
	return reflect.TypeOf((*storagecontainerRetrieveState)(nil)).Elem()
}

type storagecontainerRetrieveArgs struct {
	ExtendedLocation      *StoragecontainersExtendedLocation `pulumi:"extendedLocation"`
	Location              *string                            `pulumi:"location"`
	Path                  *string                            `pulumi:"path"`
	ProvisioningState     *string                            `pulumi:"provisioningState"`
	ResourceGroupName     string                             `pulumi:"resourceGroupName"`
	ResourceName          *string                            `pulumi:"resourceName"`
	StoragecontainersName *string                            `pulumi:"storagecontainersName"`
	Tags                  map[string]string                  `pulumi:"tags"`
}


type StoragecontainerRetrieveArgs struct {
	ExtendedLocation      StoragecontainersExtendedLocationPtrInput
	Location              pulumi.StringPtrInput
	Path                  pulumi.StringPtrInput
	ProvisioningState     pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	ResourceName          pulumi.StringPtrInput
	StoragecontainersName pulumi.StringPtrInput
	Tags                  pulumi.StringMapInput
}

func (StoragecontainerRetrieveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storagecontainerRetrieveArgs)(nil)).Elem()
}

type StoragecontainerRetrieveInput interface {
	pulumi.Input

	ToStoragecontainerRetrieveOutput() StoragecontainerRetrieveOutput
	ToStoragecontainerRetrieveOutputWithContext(ctx context.Context) StoragecontainerRetrieveOutput
}

func (*StoragecontainerRetrieve) ElementType() reflect.Type {
	return reflect.TypeOf((**StoragecontainerRetrieve)(nil)).Elem()
}

func (i *StoragecontainerRetrieve) ToStoragecontainerRetrieveOutput() StoragecontainerRetrieveOutput {
	return i.ToStoragecontainerRetrieveOutputWithContext(context.Background())
}

func (i *StoragecontainerRetrieve) ToStoragecontainerRetrieveOutputWithContext(ctx context.Context) StoragecontainerRetrieveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoragecontainerRetrieveOutput)
}

type StoragecontainerRetrieveOutput struct{ *pulumi.OutputState }

func (StoragecontainerRetrieveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StoragecontainerRetrieve)(nil)).Elem()
}

func (o StoragecontainerRetrieveOutput) ToStoragecontainerRetrieveOutput() StoragecontainerRetrieveOutput {
	return o
}

func (o StoragecontainerRetrieveOutput) ToStoragecontainerRetrieveOutputWithContext(ctx context.Context) StoragecontainerRetrieveOutput {
	return o
}

func (o StoragecontainerRetrieveOutput) AvailableSizeMB() pulumi.Float64Output {
	return o.ApplyT(func(v *StoragecontainerRetrieve) pulumi.Float64Output { return v.AvailableSizeMB }).(pulumi.Float64Output)
}

func (o StoragecontainerRetrieveOutput) ContainerSizeMB() pulumi.Float64Output {
	return o.ApplyT(func(v *StoragecontainerRetrieve) pulumi.Float64Output { return v.ContainerSizeMB }).(pulumi.Float64Output)
}

func (o StoragecontainerRetrieveOutput) ExtendedLocation() StoragecontainersResponseExtendedLocationPtrOutput {
	return o.ApplyT(func(v *StoragecontainerRetrieve) StoragecontainersResponseExtendedLocationPtrOutput {
		return v.ExtendedLocation
	}).(StoragecontainersResponseExtendedLocationPtrOutput)
}

func (o StoragecontainerRetrieveOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *StoragecontainerRetrieve) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o StoragecontainerRetrieveOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StoragecontainerRetrieve) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StoragecontainerRetrieveOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StoragecontainerRetrieve) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

func (o StoragecontainerRetrieveOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StoragecontainerRetrieve) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o StoragecontainerRetrieveOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StoragecontainerRetrieve) pulumi.StringPtrOutput { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o StoragecontainerRetrieveOutput) Status() StorageContainerStatusResponseOutput {
	return o.ApplyT(func(v *StoragecontainerRetrieve) StorageContainerStatusResponseOutput { return v.Status }).(StorageContainerStatusResponseOutput)
}

func (o StoragecontainerRetrieveOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *StoragecontainerRetrieve) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o StoragecontainerRetrieveOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StoragecontainerRetrieve) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o StoragecontainerRetrieveOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StoragecontainerRetrieve) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StoragecontainerRetrieveOutput{})
}
