


package v20160430preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Image struct {
	pulumi.CustomResourceState

	Location             pulumi.StringOutput                  `pulumi:"location"`
	Name                 pulumi.StringOutput                  `pulumi:"name"`
	ProvisioningState    pulumi.StringOutput                  `pulumi:"provisioningState"`
	SourceVirtualMachine SubResourceResponsePtrOutput         `pulumi:"sourceVirtualMachine"`
	StorageProfile       ImageStorageProfileResponsePtrOutput `pulumi:"storageProfile"`
	Tags                 pulumi.StringMapOutput               `pulumi:"tags"`
	Type                 pulumi.StringOutput                  `pulumi:"type"`
}


func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:compute/v20160430preview:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute:Image"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20170330:Image"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20170330:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20171201:Image"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20171201:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180401:Image"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20180401:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:Image"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20180601:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20181001:Image"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20181001:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:Image"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20190301:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:Image"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20190701:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:Image"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20191201:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:Image"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20200601:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:Image"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20201201:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:Image"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20210301:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:Image"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20210401:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:Image"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20210701:Image"),
		},
	})
	opts = append(opts, aliases)
	var resource Image
	err := ctx.RegisterResource("azure-native:compute/v20160430preview:Image", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageState, opts ...pulumi.ResourceOption) (*Image, error) {
	var resource Image
	err := ctx.ReadResource("azure-native:compute/v20160430preview:Image", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type imageState struct {
}

type ImageState struct {
}

func (ImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageState)(nil)).Elem()
}

type imageArgs struct {
	ImageName            *string              `pulumi:"imageName"`
	Location             *string              `pulumi:"location"`
	ResourceGroupName    string               `pulumi:"resourceGroupName"`
	SourceVirtualMachine *SubResource         `pulumi:"sourceVirtualMachine"`
	StorageProfile       *ImageStorageProfile `pulumi:"storageProfile"`
	Tags                 map[string]string    `pulumi:"tags"`
}


type ImageArgs struct {
	ImageName            pulumi.StringPtrInput
	Location             pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	SourceVirtualMachine SubResourcePtrInput
	StorageProfile       ImageStorageProfilePtrInput
	Tags                 pulumi.StringMapInput
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}

type ImageInput interface {
	pulumi.Input

	ToImageOutput() ImageOutput
	ToImageOutputWithContext(ctx context.Context) ImageOutput
}

func (*Image) ElementType() reflect.Type {
	return reflect.TypeOf((*Image)(nil))
}

func (i *Image) ToImageOutput() ImageOutput {
	return i.ToImageOutputWithContext(context.Background())
}

func (i *Image) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOutput)
}

type ImageOutput struct{ *pulumi.OutputState }

func (ImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Image)(nil))
}

func (o ImageOutput) ToImageOutput() ImageOutput {
	return o
}

func (o ImageOutput) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageInput)(nil)).Elem(), &Image{})
	pulumi.RegisterOutputType(ImageOutput{})
}
