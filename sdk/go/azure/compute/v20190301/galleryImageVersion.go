


package v20190301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GalleryImageVersion struct {
	pulumi.CustomResourceState

	Location          pulumi.StringOutput                                `pulumi:"location"`
	Name              pulumi.StringOutput                                `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                                `pulumi:"provisioningState"`
	PublishingProfile GalleryImageVersionPublishingProfileResponseOutput `pulumi:"publishingProfile"`
	ReplicationStatus ReplicationStatusResponseOutput                    `pulumi:"replicationStatus"`
	StorageProfile    GalleryImageVersionStorageProfileResponseOutput    `pulumi:"storageProfile"`
	Tags              pulumi.StringMapOutput                             `pulumi:"tags"`
	Type              pulumi.StringOutput                                `pulumi:"type"`
}


func NewGalleryImageVersion(ctx *pulumi.Context,
	name string, args *GalleryImageVersionArgs, opts ...pulumi.ResourceOption) (*GalleryImageVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GalleryImageName == nil {
		return nil, errors.New("invalid value for required argument 'GalleryImageName'")
	}
	if args.GalleryName == nil {
		return nil, errors.New("invalid value for required argument 'GalleryName'")
	}
	if args.PublishingProfile == nil {
		return nil, errors.New("invalid value for required argument 'PublishingProfile'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:GalleryImageVersion"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:GalleryImageVersion"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:GalleryImageVersion"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:GalleryImageVersion"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200930:GalleryImageVersion"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:GalleryImageVersion"),
		},
	})
	opts = append(opts, aliases)
	var resource GalleryImageVersion
	err := ctx.RegisterResource("azure-native:compute/v20190301:GalleryImageVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGalleryImageVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GalleryImageVersionState, opts ...pulumi.ResourceOption) (*GalleryImageVersion, error) {
	var resource GalleryImageVersion
	err := ctx.ReadResource("azure-native:compute/v20190301:GalleryImageVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type galleryImageVersionState struct {
}

type GalleryImageVersionState struct {
}

func (GalleryImageVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryImageVersionState)(nil)).Elem()
}

type galleryImageVersionArgs struct {
	GalleryImageName        string                               `pulumi:"galleryImageName"`
	GalleryImageVersionName *string                              `pulumi:"galleryImageVersionName"`
	GalleryName             string                               `pulumi:"galleryName"`
	Location                *string                              `pulumi:"location"`
	PublishingProfile       GalleryImageVersionPublishingProfile `pulumi:"publishingProfile"`
	ResourceGroupName       string                               `pulumi:"resourceGroupName"`
	Tags                    map[string]string                    `pulumi:"tags"`
}


type GalleryImageVersionArgs struct {
	GalleryImageName        pulumi.StringInput
	GalleryImageVersionName pulumi.StringPtrInput
	GalleryName             pulumi.StringInput
	Location                pulumi.StringPtrInput
	PublishingProfile       GalleryImageVersionPublishingProfileInput
	ResourceGroupName       pulumi.StringInput
	Tags                    pulumi.StringMapInput
}

func (GalleryImageVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryImageVersionArgs)(nil)).Elem()
}

type GalleryImageVersionInput interface {
	pulumi.Input

	ToGalleryImageVersionOutput() GalleryImageVersionOutput
	ToGalleryImageVersionOutputWithContext(ctx context.Context) GalleryImageVersionOutput
}

func (*GalleryImageVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImageVersion)(nil)).Elem()
}

func (i *GalleryImageVersion) ToGalleryImageVersionOutput() GalleryImageVersionOutput {
	return i.ToGalleryImageVersionOutputWithContext(context.Background())
}

func (i *GalleryImageVersion) ToGalleryImageVersionOutputWithContext(ctx context.Context) GalleryImageVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageVersionOutput)
}

type GalleryImageVersionOutput struct{ *pulumi.OutputState }

func (GalleryImageVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImageVersion)(nil)).Elem()
}

func (o GalleryImageVersionOutput) ToGalleryImageVersionOutput() GalleryImageVersionOutput {
	return o
}

func (o GalleryImageVersionOutput) ToGalleryImageVersionOutputWithContext(ctx context.Context) GalleryImageVersionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GalleryImageVersionOutput{})
}
