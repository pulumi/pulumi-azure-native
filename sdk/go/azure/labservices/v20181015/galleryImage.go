


package v20181015

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
			Type: pulumi.String("azure-native:labservices:GalleryImage"),
		},
	})
	opts = append(opts, aliases)
	var resource GalleryImage
	err := ctx.RegisterResource("azure-native:labservices/v20181015:GalleryImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGalleryImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GalleryImageState, opts ...pulumi.ResourceOption) (*GalleryImage, error) {
	var resource GalleryImage
	err := ctx.ReadResource("azure-native:labservices/v20181015:GalleryImage", name, id, state, &resource, opts...)
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

func (o GalleryImageOutput) Author() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.Author }).(pulumi.StringOutput)
}

func (o GalleryImageOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o GalleryImageOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o GalleryImageOutput) Icon() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.Icon }).(pulumi.StringOutput)
}

func (o GalleryImageOutput) ImageReference() GalleryImageReferenceResponseOutput {
	return o.ApplyT(func(v *GalleryImage) GalleryImageReferenceResponseOutput { return v.ImageReference }).(GalleryImageReferenceResponseOutput)
}

func (o GalleryImageOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.BoolPtrOutput { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

func (o GalleryImageOutput) IsOverride() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.BoolPtrOutput { return v.IsOverride }).(pulumi.BoolPtrOutput)
}

func (o GalleryImageOutput) IsPlanAuthorized() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.BoolPtrOutput { return v.IsPlanAuthorized }).(pulumi.BoolPtrOutput)
}

func (o GalleryImageOutput) LatestOperationResult() LatestOperationResultResponseOutput {
	return o.ApplyT(func(v *GalleryImage) LatestOperationResultResponseOutput { return v.LatestOperationResult }).(LatestOperationResultResponseOutput)
}

func (o GalleryImageOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o GalleryImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GalleryImageOutput) PlanId() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.PlanId }).(pulumi.StringOutput)
}

func (o GalleryImageOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o GalleryImageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GalleryImageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o GalleryImageOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringPtrOutput { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GalleryImageOutput{})
}
