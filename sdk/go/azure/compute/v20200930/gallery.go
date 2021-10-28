


package v20200930

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Gallery struct {
	pulumi.CustomResourceState

	Description       pulumi.StringPtrOutput             `pulumi:"description"`
	Identifier        GalleryIdentifierResponsePtrOutput `pulumi:"identifier"`
	Location          pulumi.StringOutput                `pulumi:"location"`
	Name              pulumi.StringOutput                `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                `pulumi:"provisioningState"`
	SharingProfile    SharingProfileResponsePtrOutput    `pulumi:"sharingProfile"`
	Tags              pulumi.StringMapOutput             `pulumi:"tags"`
	Type              pulumi.StringOutput                `pulumi:"type"`
}


func NewGallery(ctx *pulumi.Context,
	name string, args *GalleryArgs, opts ...pulumi.ResourceOption) (*Gallery, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:compute/v20200930:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:compute:Gallery"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:Gallery"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20180601:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:Gallery"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20190301:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:Gallery"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20190701:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:Gallery"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20191201:Gallery"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:Gallery"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20210701:Gallery"),
		},
	})
	opts = append(opts, aliases)
	var resource Gallery
	err := ctx.RegisterResource("azure-native:compute/v20200930:Gallery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGallery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GalleryState, opts ...pulumi.ResourceOption) (*Gallery, error) {
	var resource Gallery
	err := ctx.ReadResource("azure-native:compute/v20200930:Gallery", name, id, state, &resource, opts...)
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
	Description       *string           `pulumi:"description"`
	GalleryName       *string           `pulumi:"galleryName"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	SharingProfile    *SharingProfile   `pulumi:"sharingProfile"`
	Tags              map[string]string `pulumi:"tags"`
}


type GalleryArgs struct {
	Description       pulumi.StringPtrInput
	GalleryName       pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	SharingProfile    SharingProfilePtrInput
	Tags              pulumi.StringMapInput
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
	return reflect.TypeOf((*Gallery)(nil))
}

func (i *Gallery) ToGalleryOutput() GalleryOutput {
	return i.ToGalleryOutputWithContext(context.Background())
}

func (i *Gallery) ToGalleryOutputWithContext(ctx context.Context) GalleryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryOutput)
}

type GalleryOutput struct{ *pulumi.OutputState }

func (GalleryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Gallery)(nil))
}

func (o GalleryOutput) ToGalleryOutput() GalleryOutput {
	return o
}

func (o GalleryOutput) ToGalleryOutputWithContext(ctx context.Context) GalleryOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GalleryOutput{})
}
