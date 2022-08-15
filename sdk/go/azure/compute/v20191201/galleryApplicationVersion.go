


package v20191201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type GalleryApplicationVersion struct {
	pulumi.CustomResourceState

	Location          pulumi.StringOutput                                      `pulumi:"location"`
	Name              pulumi.StringOutput                                      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                                      `pulumi:"provisioningState"`
	PublishingProfile GalleryApplicationVersionPublishingProfileResponseOutput `pulumi:"publishingProfile"`
	ReplicationStatus ReplicationStatusResponseOutput                          `pulumi:"replicationStatus"`
	Tags              pulumi.StringMapOutput                                   `pulumi:"tags"`
	Type              pulumi.StringOutput                                      `pulumi:"type"`
}


func NewGalleryApplicationVersion(ctx *pulumi.Context,
	name string, args *GalleryApplicationVersionArgs, opts ...pulumi.ResourceOption) (*GalleryApplicationVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GalleryApplicationName == nil {
		return nil, errors.New("invalid value for required argument 'GalleryApplicationName'")
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
			Type: pulumi.String("azure-native:compute:GalleryApplicationVersion"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:GalleryApplicationVersion"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:GalleryApplicationVersion"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200930:GalleryApplicationVersion"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:GalleryApplicationVersion"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211001:GalleryApplicationVersion"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220103:GalleryApplicationVersion"),
		},
	})
	opts = append(opts, aliases)
	var resource GalleryApplicationVersion
	err := ctx.RegisterResource("azure-native:compute/v20191201:GalleryApplicationVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGalleryApplicationVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GalleryApplicationVersionState, opts ...pulumi.ResourceOption) (*GalleryApplicationVersion, error) {
	var resource GalleryApplicationVersion
	err := ctx.ReadResource("azure-native:compute/v20191201:GalleryApplicationVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type galleryApplicationVersionState struct {
}

type GalleryApplicationVersionState struct {
}

func (GalleryApplicationVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryApplicationVersionState)(nil)).Elem()
}

type galleryApplicationVersionArgs struct {
	GalleryApplicationName        string                                     `pulumi:"galleryApplicationName"`
	GalleryApplicationVersionName *string                                    `pulumi:"galleryApplicationVersionName"`
	GalleryName                   string                                     `pulumi:"galleryName"`
	Location                      *string                                    `pulumi:"location"`
	PublishingProfile             GalleryApplicationVersionPublishingProfile `pulumi:"publishingProfile"`
	ResourceGroupName             string                                     `pulumi:"resourceGroupName"`
	Tags                          map[string]string                          `pulumi:"tags"`
}


type GalleryApplicationVersionArgs struct {
	GalleryApplicationName        pulumi.StringInput
	GalleryApplicationVersionName pulumi.StringPtrInput
	GalleryName                   pulumi.StringInput
	Location                      pulumi.StringPtrInput
	PublishingProfile             GalleryApplicationVersionPublishingProfileInput
	ResourceGroupName             pulumi.StringInput
	Tags                          pulumi.StringMapInput
}

func (GalleryApplicationVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryApplicationVersionArgs)(nil)).Elem()
}

type GalleryApplicationVersionInput interface {
	pulumi.Input

	ToGalleryApplicationVersionOutput() GalleryApplicationVersionOutput
	ToGalleryApplicationVersionOutputWithContext(ctx context.Context) GalleryApplicationVersionOutput
}

func (*GalleryApplicationVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryApplicationVersion)(nil)).Elem()
}

func (i *GalleryApplicationVersion) ToGalleryApplicationVersionOutput() GalleryApplicationVersionOutput {
	return i.ToGalleryApplicationVersionOutputWithContext(context.Background())
}

func (i *GalleryApplicationVersion) ToGalleryApplicationVersionOutputWithContext(ctx context.Context) GalleryApplicationVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryApplicationVersionOutput)
}

type GalleryApplicationVersionOutput struct{ *pulumi.OutputState }

func (GalleryApplicationVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryApplicationVersion)(nil)).Elem()
}

func (o GalleryApplicationVersionOutput) ToGalleryApplicationVersionOutput() GalleryApplicationVersionOutput {
	return o
}

func (o GalleryApplicationVersionOutput) ToGalleryApplicationVersionOutputWithContext(ctx context.Context) GalleryApplicationVersionOutput {
	return o
}

func (o GalleryApplicationVersionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryApplicationVersion) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o GalleryApplicationVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryApplicationVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GalleryApplicationVersionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryApplicationVersion) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o GalleryApplicationVersionOutput) PublishingProfile() GalleryApplicationVersionPublishingProfileResponseOutput {
	return o.ApplyT(func(v *GalleryApplicationVersion) GalleryApplicationVersionPublishingProfileResponseOutput {
		return v.PublishingProfile
	}).(GalleryApplicationVersionPublishingProfileResponseOutput)
}

func (o GalleryApplicationVersionOutput) ReplicationStatus() ReplicationStatusResponseOutput {
	return o.ApplyT(func(v *GalleryApplicationVersion) ReplicationStatusResponseOutput { return v.ReplicationStatus }).(ReplicationStatusResponseOutput)
}

func (o GalleryApplicationVersionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GalleryApplicationVersion) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GalleryApplicationVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryApplicationVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GalleryApplicationVersionOutput{})
}
