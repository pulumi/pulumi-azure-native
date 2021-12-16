


package v20191201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GalleryImage struct {
	pulumi.CustomResourceState

	Description         pulumi.StringPtrOutput                           `pulumi:"description"`
	Disallowed          DisallowedResponsePtrOutput                      `pulumi:"disallowed"`
	EndOfLifeDate       pulumi.StringPtrOutput                           `pulumi:"endOfLifeDate"`
	Eula                pulumi.StringPtrOutput                           `pulumi:"eula"`
	HyperVGeneration    pulumi.StringPtrOutput                           `pulumi:"hyperVGeneration"`
	Identifier          GalleryImageIdentifierResponseOutput             `pulumi:"identifier"`
	Location            pulumi.StringOutput                              `pulumi:"location"`
	Name                pulumi.StringOutput                              `pulumi:"name"`
	OsState             pulumi.StringOutput                              `pulumi:"osState"`
	OsType              pulumi.StringOutput                              `pulumi:"osType"`
	PrivacyStatementUri pulumi.StringPtrOutput                           `pulumi:"privacyStatementUri"`
	ProvisioningState   pulumi.StringOutput                              `pulumi:"provisioningState"`
	PurchasePlan        ImagePurchasePlanResponsePtrOutput               `pulumi:"purchasePlan"`
	Recommended         RecommendedMachineConfigurationResponsePtrOutput `pulumi:"recommended"`
	ReleaseNoteUri      pulumi.StringPtrOutput                           `pulumi:"releaseNoteUri"`
	Tags                pulumi.StringMapOutput                           `pulumi:"tags"`
	Type                pulumi.StringOutput                              `pulumi:"type"`
}


func NewGalleryImage(ctx *pulumi.Context,
	name string, args *GalleryImageArgs, opts ...pulumi.ResourceOption) (*GalleryImage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GalleryName == nil {
		return nil, errors.New("invalid value for required argument 'GalleryName'")
	}
	if args.Identifier == nil {
		return nil, errors.New("invalid value for required argument 'Identifier'")
	}
	if args.OsState == nil {
		return nil, errors.New("invalid value for required argument 'OsState'")
	}
	if args.OsType == nil {
		return nil, errors.New("invalid value for required argument 'OsType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200930:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:GalleryImage"),
		},
	})
	opts = append(opts, aliases)
	var resource GalleryImage
	err := ctx.RegisterResource("azure-native:compute/v20191201:GalleryImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGalleryImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GalleryImageState, opts ...pulumi.ResourceOption) (*GalleryImage, error) {
	var resource GalleryImage
	err := ctx.ReadResource("azure-native:compute/v20191201:GalleryImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type galleryImageState struct {
}

type GalleryImageState struct {
}

func (GalleryImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryImageState)(nil)).Elem()
}

type galleryImageArgs struct {
	Description         *string                          `pulumi:"description"`
	Disallowed          *Disallowed                      `pulumi:"disallowed"`
	EndOfLifeDate       *string                          `pulumi:"endOfLifeDate"`
	Eula                *string                          `pulumi:"eula"`
	GalleryImageName    *string                          `pulumi:"galleryImageName"`
	GalleryName         string                           `pulumi:"galleryName"`
	HyperVGeneration    *string                          `pulumi:"hyperVGeneration"`
	Identifier          GalleryImageIdentifier           `pulumi:"identifier"`
	Location            *string                          `pulumi:"location"`
	OsState             OperatingSystemStateTypes        `pulumi:"osState"`
	OsType              OperatingSystemTypes             `pulumi:"osType"`
	PrivacyStatementUri *string                          `pulumi:"privacyStatementUri"`
	PurchasePlan        *ImagePurchasePlan               `pulumi:"purchasePlan"`
	Recommended         *RecommendedMachineConfiguration `pulumi:"recommended"`
	ReleaseNoteUri      *string                          `pulumi:"releaseNoteUri"`
	ResourceGroupName   string                           `pulumi:"resourceGroupName"`
	Tags                map[string]string                `pulumi:"tags"`
}


type GalleryImageArgs struct {
	Description         pulumi.StringPtrInput
	Disallowed          DisallowedPtrInput
	EndOfLifeDate       pulumi.StringPtrInput
	Eula                pulumi.StringPtrInput
	GalleryImageName    pulumi.StringPtrInput
	GalleryName         pulumi.StringInput
	HyperVGeneration    pulumi.StringPtrInput
	Identifier          GalleryImageIdentifierInput
	Location            pulumi.StringPtrInput
	OsState             OperatingSystemStateTypesInput
	OsType              OperatingSystemTypesInput
	PrivacyStatementUri pulumi.StringPtrInput
	PurchasePlan        ImagePurchasePlanPtrInput
	Recommended         RecommendedMachineConfigurationPtrInput
	ReleaseNoteUri      pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	Tags                pulumi.StringMapInput
}

func (GalleryImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryImageArgs)(nil)).Elem()
}

type GalleryImageInput interface {
	pulumi.Input

	ToGalleryImageOutput() GalleryImageOutput
	ToGalleryImageOutputWithContext(ctx context.Context) GalleryImageOutput
}

func (*GalleryImage) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImage)(nil)).Elem()
}

func (i *GalleryImage) ToGalleryImageOutput() GalleryImageOutput {
	return i.ToGalleryImageOutputWithContext(context.Background())
}

func (i *GalleryImage) ToGalleryImageOutputWithContext(ctx context.Context) GalleryImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageOutput)
}

type GalleryImageOutput struct{ *pulumi.OutputState }

func (GalleryImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImage)(nil)).Elem()
}

func (o GalleryImageOutput) ToGalleryImageOutput() GalleryImageOutput {
	return o
}

func (o GalleryImageOutput) ToGalleryImageOutputWithContext(ctx context.Context) GalleryImageOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GalleryImageOutput{})
}
