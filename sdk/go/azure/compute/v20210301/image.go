


package v20210301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Image struct {
	pulumi.CustomResourceState

	ExtendedLocation     ExtendedLocationResponsePtrOutput    `pulumi:"extendedLocation"`
	HyperVGeneration     pulumi.StringPtrOutput               `pulumi:"hyperVGeneration"`
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
			Type: pulumi.String("azure-native:compute:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160430preview:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20170330:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20171201:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180401:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20181001:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:Image"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:Image"),
		},
	})
	opts = append(opts, aliases)
	var resource Image
	err := ctx.RegisterResource("azure-native:compute/v20210301:Image", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageState, opts ...pulumi.ResourceOption) (*Image, error) {
	var resource Image
	err := ctx.ReadResource("azure-native:compute/v20210301:Image", name, id, state, &resource, opts...)
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
	ExtendedLocation     *ExtendedLocation    `pulumi:"extendedLocation"`
	HyperVGeneration     *string              `pulumi:"hyperVGeneration"`
	ImageName            *string              `pulumi:"imageName"`
	Location             *string              `pulumi:"location"`
	ResourceGroupName    string               `pulumi:"resourceGroupName"`
	SourceVirtualMachine *SubResource         `pulumi:"sourceVirtualMachine"`
	StorageProfile       *ImageStorageProfile `pulumi:"storageProfile"`
	Tags                 map[string]string    `pulumi:"tags"`
}


type ImageArgs struct {
	ExtendedLocation     ExtendedLocationPtrInput
	HyperVGeneration     pulumi.StringPtrInput
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
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (i *Image) ToImageOutput() ImageOutput {
	return i.ToImageOutputWithContext(context.Background())
}

func (i *Image) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOutput)
}

type ImageOutput struct{ *pulumi.OutputState }

func (ImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (o ImageOutput) ToImageOutput() ImageOutput {
	return o
}

func (o ImageOutput) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return o
}

func (o ImageOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *Image) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o ImageOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

func (o ImageOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ImageOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ImageOutput) SourceVirtualMachine() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *Image) SubResourceResponsePtrOutput { return v.SourceVirtualMachine }).(SubResourceResponsePtrOutput)
}

func (o ImageOutput) StorageProfile() ImageStorageProfileResponsePtrOutput {
	return o.ApplyT(func(v *Image) ImageStorageProfileResponsePtrOutput { return v.StorageProfile }).(ImageStorageProfileResponsePtrOutput)
}

func (o ImageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Image) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ImageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ImageOutput{})
}
