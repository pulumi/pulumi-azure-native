


package hybridcontainerservice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StorageSpaceRetrieve struct {
	pulumi.CustomResourceState

	ExtendedLocation StorageSpacesResponseExtendedLocationPtrOutput `pulumi:"extendedLocation"`
	Location         pulumi.StringOutput                            `pulumi:"location"`
	Name             pulumi.StringOutput                            `pulumi:"name"`
	Properties       StorageSpacesPropertiesResponseOutput          `pulumi:"properties"`
	SystemData       SystemDataResponseOutput                       `pulumi:"systemData"`
	Tags             pulumi.StringMapOutput                         `pulumi:"tags"`
	Type             pulumi.StringOutput                            `pulumi:"type"`
}


func NewStorageSpaceRetrieve(ctx *pulumi.Context,
	name string, args *StorageSpaceRetrieveArgs, opts ...pulumi.ResourceOption) (*StorageSpaceRetrieve, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridcontainerservice/v20220501preview:storageSpaceRetrieve"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageSpaceRetrieve
	err := ctx.RegisterResource("azure-native:hybridcontainerservice:storageSpaceRetrieve", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStorageSpaceRetrieve(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageSpaceRetrieveState, opts ...pulumi.ResourceOption) (*StorageSpaceRetrieve, error) {
	var resource StorageSpaceRetrieve
	err := ctx.ReadResource("azure-native:hybridcontainerservice:storageSpaceRetrieve", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type storageSpaceRetrieveState struct {
}

type StorageSpaceRetrieveState struct {
}

func (StorageSpaceRetrieveState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageSpaceRetrieveState)(nil)).Elem()
}

type storageSpaceRetrieveArgs struct {
	ExtendedLocation  *StorageSpacesExtendedLocation `pulumi:"extendedLocation"`
	Location          *string                        `pulumi:"location"`
	Properties        *StorageSpacesProperties       `pulumi:"properties"`
	ResourceGroupName string                         `pulumi:"resourceGroupName"`
	StorageSpacesName *string                        `pulumi:"storageSpacesName"`
	Tags              map[string]string              `pulumi:"tags"`
}


type StorageSpaceRetrieveArgs struct {
	ExtendedLocation  StorageSpacesExtendedLocationPtrInput
	Location          pulumi.StringPtrInput
	Properties        StorageSpacesPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	StorageSpacesName pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (StorageSpaceRetrieveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageSpaceRetrieveArgs)(nil)).Elem()
}

type StorageSpaceRetrieveInput interface {
	pulumi.Input

	ToStorageSpaceRetrieveOutput() StorageSpaceRetrieveOutput
	ToStorageSpaceRetrieveOutputWithContext(ctx context.Context) StorageSpaceRetrieveOutput
}

func (*StorageSpaceRetrieve) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpaceRetrieve)(nil)).Elem()
}

func (i *StorageSpaceRetrieve) ToStorageSpaceRetrieveOutput() StorageSpaceRetrieveOutput {
	return i.ToStorageSpaceRetrieveOutputWithContext(context.Background())
}

func (i *StorageSpaceRetrieve) ToStorageSpaceRetrieveOutputWithContext(ctx context.Context) StorageSpaceRetrieveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageSpaceRetrieveOutput)
}

type StorageSpaceRetrieveOutput struct{ *pulumi.OutputState }

func (StorageSpaceRetrieveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageSpaceRetrieve)(nil)).Elem()
}

func (o StorageSpaceRetrieveOutput) ToStorageSpaceRetrieveOutput() StorageSpaceRetrieveOutput {
	return o
}

func (o StorageSpaceRetrieveOutput) ToStorageSpaceRetrieveOutputWithContext(ctx context.Context) StorageSpaceRetrieveOutput {
	return o
}

func (o StorageSpaceRetrieveOutput) ExtendedLocation() StorageSpacesResponseExtendedLocationPtrOutput {
	return o.ApplyT(func(v *StorageSpaceRetrieve) StorageSpacesResponseExtendedLocationPtrOutput {
		return v.ExtendedLocation
	}).(StorageSpacesResponseExtendedLocationPtrOutput)
}

func (o StorageSpaceRetrieveOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageSpaceRetrieve) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o StorageSpaceRetrieveOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageSpaceRetrieve) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StorageSpaceRetrieveOutput) Properties() StorageSpacesPropertiesResponseOutput {
	return o.ApplyT(func(v *StorageSpaceRetrieve) StorageSpacesPropertiesResponseOutput { return v.Properties }).(StorageSpacesPropertiesResponseOutput)
}

func (o StorageSpaceRetrieveOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *StorageSpaceRetrieve) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o StorageSpaceRetrieveOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StorageSpaceRetrieve) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o StorageSpaceRetrieveOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageSpaceRetrieve) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageSpaceRetrieveOutput{})
}
