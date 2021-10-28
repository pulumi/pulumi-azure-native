


package labservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GalleryImage struct {
	pulumi.CustomResourceState

	Author                pulumi.StringOutput                 `pulumi:"author"`
	CreatedDate           pulumi.StringOutput                 `pulumi:"createdDate"`
	Description           pulumi.StringOutput                 `pulumi:"description"`
	Icon                  pulumi.StringOutput                 `pulumi:"icon"`
	ImageReference        GalleryImageReferenceResponseOutput `pulumi:"imageReference"`
	IsEnabled             pulumi.BoolPtrOutput                `pulumi:"isEnabled"`
	IsOverride            pulumi.BoolPtrOutput                `pulumi:"isOverride"`
	IsPlanAuthorized      pulumi.BoolPtrOutput                `pulumi:"isPlanAuthorized"`
	LatestOperationResult LatestOperationResultResponseOutput `pulumi:"latestOperationResult"`
	Location              pulumi.StringPtrOutput              `pulumi:"location"`
	Name                  pulumi.StringOutput                 `pulumi:"name"`
	PlanId                pulumi.StringOutput                 `pulumi:"planId"`
	ProvisioningState     pulumi.StringPtrOutput              `pulumi:"provisioningState"`
	Tags                  pulumi.StringMapOutput              `pulumi:"tags"`
	Type                  pulumi.StringOutput                 `pulumi:"type"`
	UniqueIdentifier      pulumi.StringPtrOutput              `pulumi:"uniqueIdentifier"`
}


func NewGalleryImage(ctx *pulumi.Context,
	name string, args *GalleryImageArgs, opts ...pulumi.ResourceOption) (*GalleryImage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabAccountName == nil {
		return nil, errors.New("invalid value for required argument 'LabAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:labservices:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:labservices/v20181015:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-nextgen:labservices/v20181015:GalleryImage"),
		},
	})
	opts = append(opts, aliases)
	var resource GalleryImage
	err := ctx.RegisterResource("azure-native:labservices:GalleryImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGalleryImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GalleryImageState, opts ...pulumi.ResourceOption) (*GalleryImage, error) {
	var resource GalleryImage
	err := ctx.ReadResource("azure-native:labservices:GalleryImage", name, id, state, &resource, opts...)
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
	GalleryImageName  *string           `pulumi:"galleryImageName"`
	IsEnabled         *bool             `pulumi:"isEnabled"`
	IsOverride        *bool             `pulumi:"isOverride"`
	IsPlanAuthorized  *bool             `pulumi:"isPlanAuthorized"`
	LabAccountName    string            `pulumi:"labAccountName"`
	Location          *string           `pulumi:"location"`
	ProvisioningState *string           `pulumi:"provisioningState"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	UniqueIdentifier  *string           `pulumi:"uniqueIdentifier"`
}


type GalleryImageArgs struct {
	GalleryImageName  pulumi.StringPtrInput
	IsEnabled         pulumi.BoolPtrInput
	IsOverride        pulumi.BoolPtrInput
	IsPlanAuthorized  pulumi.BoolPtrInput
	LabAccountName    pulumi.StringInput
	Location          pulumi.StringPtrInput
	ProvisioningState pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	UniqueIdentifier  pulumi.StringPtrInput
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
	return reflect.TypeOf((*GalleryImage)(nil))
}

func (i *GalleryImage) ToGalleryImageOutput() GalleryImageOutput {
	return i.ToGalleryImageOutputWithContext(context.Background())
}

func (i *GalleryImage) ToGalleryImageOutputWithContext(ctx context.Context) GalleryImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageOutput)
}

type GalleryImageOutput struct{ *pulumi.OutputState }

func (GalleryImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImage)(nil))
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
