


package v20221111preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Gallery struct {
	pulumi.CustomResourceState

	GalleryResourceId pulumi.StringOutput      `pulumi:"galleryResourceId"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewGallery(ctx *pulumi.Context,
	name string, args *GalleryArgs, opts ...pulumi.ResourceOption) (*Gallery, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DevCenterName == nil {
		return nil, errors.New("invalid value for required argument 'DevCenterName'")
	}
	if args.GalleryResourceId == nil {
		return nil, errors.New("invalid value for required argument 'GalleryResourceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devcenter:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220801preview:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220901preview:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221012preview:Gallery"),
		},
	})
	opts = append(opts, aliases)
	var resource Gallery
	err := ctx.RegisterResource("azure-native:devcenter/v20221111preview:Gallery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGallery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GalleryState, opts ...pulumi.ResourceOption) (*Gallery, error) {
	var resource Gallery
	err := ctx.ReadResource("azure-native:devcenter/v20221111preview:Gallery", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type galleryState struct {
}

type GalleryState struct {
}

func (GalleryState) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryState)(nil)).Elem()
}

type galleryArgs struct {
	DevCenterName     string  `pulumi:"devCenterName"`
	GalleryName       *string `pulumi:"galleryName"`
	GalleryResourceId string  `pulumi:"galleryResourceId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type GalleryArgs struct {
	DevCenterName     pulumi.StringInput
	GalleryName       pulumi.StringPtrInput
	GalleryResourceId pulumi.StringInput
	ResourceGroupName pulumi.StringInput
}

func (GalleryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryArgs)(nil)).Elem()
}

type GalleryInput interface {
	pulumi.Input

	ToGalleryOutput() GalleryOutput
	ToGalleryOutputWithContext(ctx context.Context) GalleryOutput
}

func (*Gallery) ElementType() reflect.Type {
	return reflect.TypeOf((**Gallery)(nil)).Elem()
}

func (i *Gallery) ToGalleryOutput() GalleryOutput {
	return i.ToGalleryOutputWithContext(context.Background())
}

func (i *Gallery) ToGalleryOutputWithContext(ctx context.Context) GalleryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryOutput)
}

type GalleryOutput struct{ *pulumi.OutputState }

func (GalleryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Gallery)(nil)).Elem()
}

func (o GalleryOutput) ToGalleryOutput() GalleryOutput {
	return o
}

func (o GalleryOutput) ToGalleryOutputWithContext(ctx context.Context) GalleryOutput {
	return o
}

func (o GalleryOutput) GalleryResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Gallery) pulumi.StringOutput { return v.GalleryResourceId }).(pulumi.StringOutput)
}

func (o GalleryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Gallery) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GalleryOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Gallery) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o GalleryOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Gallery) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GalleryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Gallery) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GalleryOutput{})
}
