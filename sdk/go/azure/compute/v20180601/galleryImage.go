


package v20180601

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
			Type: pulumi.String("azure-native:compute/v20190301:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200930:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211001:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220103:GalleryImage"),
		},
	})
	opts = append(opts, aliases)
	var resource GalleryImage
	err := ctx.RegisterResource("azure-native:compute/v20180601:GalleryImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGalleryImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GalleryImageState, opts ...pulumi.ResourceOption) (*GalleryImage, error) {
	var resource GalleryImage
	err := ctx.ReadResource("azure-native:compute/v20180601:GalleryImage", name, id, state, &resource, opts...)
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

func (o GalleryImageOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GalleryImageOutput) Disallowed() DisallowedResponsePtrOutput {
	return o.ApplyT(func(v *GalleryImage) DisallowedResponsePtrOutput { return v.Disallowed }).(DisallowedResponsePtrOutput)
}

func (o GalleryImageOutput) EndOfLifeDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringPtrOutput { return v.EndOfLifeDate }).(pulumi.StringPtrOutput)
}

func (o GalleryImageOutput) Eula() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringPtrOutput { return v.Eula }).(pulumi.StringPtrOutput)
}

func (o GalleryImageOutput) Identifier() GalleryImageIdentifierResponseOutput {
	return o.ApplyT(func(v *GalleryImage) GalleryImageIdentifierResponseOutput { return v.Identifier }).(GalleryImageIdentifierResponseOutput)
}

func (o GalleryImageOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o GalleryImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GalleryImageOutput) OsState() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.OsState }).(pulumi.StringOutput)
}

func (o GalleryImageOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.OsType }).(pulumi.StringOutput)
}

func (o GalleryImageOutput) PrivacyStatementUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringPtrOutput { return v.PrivacyStatementUri }).(pulumi.StringPtrOutput)
}

func (o GalleryImageOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o GalleryImageOutput) PurchasePlan() ImagePurchasePlanResponsePtrOutput {
	return o.ApplyT(func(v *GalleryImage) ImagePurchasePlanResponsePtrOutput { return v.PurchasePlan }).(ImagePurchasePlanResponsePtrOutput)
}

func (o GalleryImageOutput) Recommended() RecommendedMachineConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *GalleryImage) RecommendedMachineConfigurationResponsePtrOutput { return v.Recommended }).(RecommendedMachineConfigurationResponsePtrOutput)
}

func (o GalleryImageOutput) ReleaseNoteUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringPtrOutput { return v.ReleaseNoteUri }).(pulumi.StringPtrOutput)
}

func (o GalleryImageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GalleryImageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GalleryImageOutput{})
}
