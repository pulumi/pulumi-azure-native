


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GalleryimageRetrieve struct {
	pulumi.CustomResourceState

	CloudInitDataSource pulumi.StringPtrOutput                  `pulumi:"cloudInitDataSource"`
	ContainerName       pulumi.StringPtrOutput                  `pulumi:"containerName"`
	ExtendedLocation    ExtendedLocationResponsePtrOutput       `pulumi:"extendedLocation"`
	HyperVGeneration    pulumi.StringPtrOutput                  `pulumi:"hyperVGeneration"`
	Identifier          GalleryImageIdentifierResponsePtrOutput `pulumi:"identifier"`
	ImagePath           pulumi.StringPtrOutput                  `pulumi:"imagePath"`
	Location            pulumi.StringOutput                     `pulumi:"location"`
	Name                pulumi.StringOutput                     `pulumi:"name"`
	OsType              pulumi.StringPtrOutput                  `pulumi:"osType"`
	ProvisioningState   pulumi.StringOutput                     `pulumi:"provisioningState"`
	ResourceName        pulumi.StringPtrOutput                  `pulumi:"resourceName"`
	Status              GalleryImageStatusResponseOutput        `pulumi:"status"`
	SystemData          SystemDataResponseOutput                `pulumi:"systemData"`
	Tags                pulumi.StringMapOutput                  `pulumi:"tags"`
	Type                pulumi.StringOutput                     `pulumi:"type"`
	Version             GalleryImageVersionResponsePtrOutput    `pulumi:"version"`
}


func NewGalleryimageRetrieve(ctx *pulumi.Context,
	name string, args *GalleryimageRetrieveArgs, opts ...pulumi.ResourceOption) (*GalleryimageRetrieve, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210701preview:galleryimageRetrieve"),
		},
	})
	opts = append(opts, aliases)
	var resource GalleryimageRetrieve
	err := ctx.RegisterResource("azure-native:azurestackhci/v20210901preview:galleryimageRetrieve", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGalleryimageRetrieve(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GalleryimageRetrieveState, opts ...pulumi.ResourceOption) (*GalleryimageRetrieve, error) {
	var resource GalleryimageRetrieve
	err := ctx.ReadResource("azure-native:azurestackhci/v20210901preview:galleryimageRetrieve", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type galleryimageRetrieveState struct {
}

type GalleryimageRetrieveState struct {
}

func (GalleryimageRetrieveState) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryimageRetrieveState)(nil)).Elem()
}

type galleryimageRetrieveArgs struct {
	CloudInitDataSource *string                 `pulumi:"cloudInitDataSource"`
	ContainerName       *string                 `pulumi:"containerName"`
	ExtendedLocation    *ExtendedLocation       `pulumi:"extendedLocation"`
	GalleryimagesName   *string                 `pulumi:"galleryimagesName"`
	HyperVGeneration    *string                 `pulumi:"hyperVGeneration"`
	Identifier          *GalleryImageIdentifier `pulumi:"identifier"`
	ImagePath           *string                 `pulumi:"imagePath"`
	Location            *string                 `pulumi:"location"`
	OsType              *OperatingSystemTypes   `pulumi:"osType"`
	ResourceGroupName   string                  `pulumi:"resourceGroupName"`
	ResourceName        *string                 `pulumi:"resourceName"`
	Tags                map[string]string       `pulumi:"tags"`
	Version             *GalleryImageVersion    `pulumi:"version"`
}


type GalleryimageRetrieveArgs struct {
	CloudInitDataSource pulumi.StringPtrInput
	ContainerName       pulumi.StringPtrInput
	ExtendedLocation    ExtendedLocationPtrInput
	GalleryimagesName   pulumi.StringPtrInput
	HyperVGeneration    pulumi.StringPtrInput
	Identifier          GalleryImageIdentifierPtrInput
	ImagePath           pulumi.StringPtrInput
	Location            pulumi.StringPtrInput
	OsType              OperatingSystemTypesPtrInput
	ResourceGroupName   pulumi.StringInput
	ResourceName        pulumi.StringPtrInput
	Tags                pulumi.StringMapInput
	Version             GalleryImageVersionPtrInput
}

func (GalleryimageRetrieveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryimageRetrieveArgs)(nil)).Elem()
}

type GalleryimageRetrieveInput interface {
	pulumi.Input

	ToGalleryimageRetrieveOutput() GalleryimageRetrieveOutput
	ToGalleryimageRetrieveOutputWithContext(ctx context.Context) GalleryimageRetrieveOutput
}

func (*GalleryimageRetrieve) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryimageRetrieve)(nil)).Elem()
}

func (i *GalleryimageRetrieve) ToGalleryimageRetrieveOutput() GalleryimageRetrieveOutput {
	return i.ToGalleryimageRetrieveOutputWithContext(context.Background())
}

func (i *GalleryimageRetrieve) ToGalleryimageRetrieveOutputWithContext(ctx context.Context) GalleryimageRetrieveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryimageRetrieveOutput)
}

type GalleryimageRetrieveOutput struct{ *pulumi.OutputState }

func (GalleryimageRetrieveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryimageRetrieve)(nil)).Elem()
}

func (o GalleryimageRetrieveOutput) ToGalleryimageRetrieveOutput() GalleryimageRetrieveOutput {
	return o
}

func (o GalleryimageRetrieveOutput) ToGalleryimageRetrieveOutputWithContext(ctx context.Context) GalleryimageRetrieveOutput {
	return o
}

func (o GalleryimageRetrieveOutput) CloudInitDataSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) pulumi.StringPtrOutput { return v.CloudInitDataSource }).(pulumi.StringPtrOutput)
}

func (o GalleryimageRetrieveOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) pulumi.StringPtrOutput { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o GalleryimageRetrieveOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o GalleryimageRetrieveOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) pulumi.StringPtrOutput { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

func (o GalleryimageRetrieveOutput) Identifier() GalleryImageIdentifierResponsePtrOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) GalleryImageIdentifierResponsePtrOutput { return v.Identifier }).(GalleryImageIdentifierResponsePtrOutput)
}

func (o GalleryimageRetrieveOutput) ImagePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) pulumi.StringPtrOutput { return v.ImagePath }).(pulumi.StringPtrOutput)
}

func (o GalleryimageRetrieveOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o GalleryimageRetrieveOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GalleryimageRetrieveOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o GalleryimageRetrieveOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o GalleryimageRetrieveOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) pulumi.StringPtrOutput { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o GalleryimageRetrieveOutput) Status() GalleryImageStatusResponseOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) GalleryImageStatusResponseOutput { return v.Status }).(GalleryImageStatusResponseOutput)
}

func (o GalleryimageRetrieveOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GalleryimageRetrieveOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GalleryimageRetrieveOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o GalleryimageRetrieveOutput) Version() GalleryImageVersionResponsePtrOutput {
	return o.ApplyT(func(v *GalleryimageRetrieve) GalleryImageVersionResponsePtrOutput { return v.Version }).(GalleryImageVersionResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GalleryimageRetrieveOutput{})
}
