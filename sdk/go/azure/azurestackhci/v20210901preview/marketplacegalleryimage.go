


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Marketplacegalleryimage struct {
	pulumi.CustomResourceState

	CloudInitDataSource pulumi.StringPtrOutput                      `pulumi:"cloudInitDataSource"`
	ContainerName       pulumi.StringPtrOutput                      `pulumi:"containerName"`
	ExtendedLocation    ExtendedLocationResponsePtrOutput           `pulumi:"extendedLocation"`
	HyperVGeneration    pulumi.StringPtrOutput                      `pulumi:"hyperVGeneration"`
	Identifier          GalleryImageIdentifierResponsePtrOutput     `pulumi:"identifier"`
	Location            pulumi.StringOutput                         `pulumi:"location"`
	Name                pulumi.StringOutput                         `pulumi:"name"`
	OsType              pulumi.StringPtrOutput                      `pulumi:"osType"`
	ProvisioningState   pulumi.StringOutput                         `pulumi:"provisioningState"`
	ResourceName        pulumi.StringPtrOutput                      `pulumi:"resourceName"`
	Status              MarketplaceGalleryImageStatusResponseOutput `pulumi:"status"`
	SystemData          SystemDataResponseOutput                    `pulumi:"systemData"`
	Tags                pulumi.StringMapOutput                      `pulumi:"tags"`
	Type                pulumi.StringOutput                         `pulumi:"type"`
	Version             GalleryImageVersionResponsePtrOutput        `pulumi:"version"`
}


func NewMarketplacegalleryimage(ctx *pulumi.Context,
	name string, args *MarketplacegalleryimageArgs, opts ...pulumi.ResourceOption) (*Marketplacegalleryimage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Marketplacegalleryimage
	err := ctx.RegisterResource("azure-native:azurestackhci/v20210901preview:marketplacegalleryimage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMarketplacegalleryimage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MarketplacegalleryimageState, opts ...pulumi.ResourceOption) (*Marketplacegalleryimage, error) {
	var resource Marketplacegalleryimage
	err := ctx.ReadResource("azure-native:azurestackhci/v20210901preview:marketplacegalleryimage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type marketplacegalleryimageState struct {
}

type MarketplacegalleryimageState struct {
}

func (MarketplacegalleryimageState) ElementType() reflect.Type {
	return reflect.TypeOf((*marketplacegalleryimageState)(nil)).Elem()
}

type marketplacegalleryimageArgs struct {
	CloudInitDataSource          *string                 `pulumi:"cloudInitDataSource"`
	ContainerName                *string                 `pulumi:"containerName"`
	ExtendedLocation             *ExtendedLocation       `pulumi:"extendedLocation"`
	HyperVGeneration             *string                 `pulumi:"hyperVGeneration"`
	Identifier                   *GalleryImageIdentifier `pulumi:"identifier"`
	Location                     *string                 `pulumi:"location"`
	MarketplacegalleryimagesName *string                 `pulumi:"marketplacegalleryimagesName"`
	OsType                       *OperatingSystemTypes   `pulumi:"osType"`
	ResourceGroupName            string                  `pulumi:"resourceGroupName"`
	ResourceName                 *string                 `pulumi:"resourceName"`
	Tags                         map[string]string       `pulumi:"tags"`
	Version                      *GalleryImageVersion    `pulumi:"version"`
}


type MarketplacegalleryimageArgs struct {
	CloudInitDataSource          pulumi.StringPtrInput
	ContainerName                pulumi.StringPtrInput
	ExtendedLocation             ExtendedLocationPtrInput
	HyperVGeneration             pulumi.StringPtrInput
	Identifier                   GalleryImageIdentifierPtrInput
	Location                     pulumi.StringPtrInput
	MarketplacegalleryimagesName pulumi.StringPtrInput
	OsType                       OperatingSystemTypesPtrInput
	ResourceGroupName            pulumi.StringInput
	ResourceName                 pulumi.StringPtrInput
	Tags                         pulumi.StringMapInput
	Version                      GalleryImageVersionPtrInput
}

func (MarketplacegalleryimageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*marketplacegalleryimageArgs)(nil)).Elem()
}

type MarketplacegalleryimageInput interface {
	pulumi.Input

	ToMarketplacegalleryimageOutput() MarketplacegalleryimageOutput
	ToMarketplacegalleryimageOutputWithContext(ctx context.Context) MarketplacegalleryimageOutput
}

func (*Marketplacegalleryimage) ElementType() reflect.Type {
	return reflect.TypeOf((**Marketplacegalleryimage)(nil)).Elem()
}

func (i *Marketplacegalleryimage) ToMarketplacegalleryimageOutput() MarketplacegalleryimageOutput {
	return i.ToMarketplacegalleryimageOutputWithContext(context.Background())
}

func (i *Marketplacegalleryimage) ToMarketplacegalleryimageOutputWithContext(ctx context.Context) MarketplacegalleryimageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarketplacegalleryimageOutput)
}

type MarketplacegalleryimageOutput struct{ *pulumi.OutputState }

func (MarketplacegalleryimageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Marketplacegalleryimage)(nil)).Elem()
}

func (o MarketplacegalleryimageOutput) ToMarketplacegalleryimageOutput() MarketplacegalleryimageOutput {
	return o
}

func (o MarketplacegalleryimageOutput) ToMarketplacegalleryimageOutputWithContext(ctx context.Context) MarketplacegalleryimageOutput {
	return o
}

func (o MarketplacegalleryimageOutput) CloudInitDataSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) pulumi.StringPtrOutput { return v.CloudInitDataSource }).(pulumi.StringPtrOutput)
}

func (o MarketplacegalleryimageOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) pulumi.StringPtrOutput { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o MarketplacegalleryimageOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o MarketplacegalleryimageOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) pulumi.StringPtrOutput { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

func (o MarketplacegalleryimageOutput) Identifier() GalleryImageIdentifierResponsePtrOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) GalleryImageIdentifierResponsePtrOutput { return v.Identifier }).(GalleryImageIdentifierResponsePtrOutput)
}

func (o MarketplacegalleryimageOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o MarketplacegalleryimageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MarketplacegalleryimageOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o MarketplacegalleryimageOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o MarketplacegalleryimageOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) pulumi.StringPtrOutput { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o MarketplacegalleryimageOutput) Status() MarketplaceGalleryImageStatusResponseOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) MarketplaceGalleryImageStatusResponseOutput { return v.Status }).(MarketplaceGalleryImageStatusResponseOutput)
}

func (o MarketplacegalleryimageOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MarketplacegalleryimageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o MarketplacegalleryimageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o MarketplacegalleryimageOutput) Version() GalleryImageVersionResponsePtrOutput {
	return o.ApplyT(func(v *Marketplacegalleryimage) GalleryImageVersionResponsePtrOutput { return v.Version }).(GalleryImageVersionResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MarketplacegalleryimageOutput{})
}
