


package v20160515

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomImage struct {
	pulumi.CustomResourceState

	Author            pulumi.StringPtrOutput                       `pulumi:"author"`
	CreationDate      pulumi.StringOutput                          `pulumi:"creationDate"`
	Description       pulumi.StringPtrOutput                       `pulumi:"description"`
	Location          pulumi.StringPtrOutput                       `pulumi:"location"`
	ManagedImageId    pulumi.StringPtrOutput                       `pulumi:"managedImageId"`
	Name              pulumi.StringOutput                          `pulumi:"name"`
	ProvisioningState pulumi.StringPtrOutput                       `pulumi:"provisioningState"`
	Tags              pulumi.StringMapOutput                       `pulumi:"tags"`
	Type              pulumi.StringOutput                          `pulumi:"type"`
	UniqueIdentifier  pulumi.StringPtrOutput                       `pulumi:"uniqueIdentifier"`
	Vhd               CustomImagePropertiesCustomResponsePtrOutput `pulumi:"vhd"`
	Vm                CustomImagePropertiesFromVmResponsePtrOutput `pulumi:"vm"`
}


func NewCustomImage(ctx *pulumi.Context,
	name string, args *CustomImageArgs, opts ...pulumi.ResourceOption) (*CustomImage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20160515:CustomImage"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab:CustomImage"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab:CustomImage"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20150521preview:CustomImage"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20150521preview:CustomImage"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:CustomImage"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20180915:CustomImage"),
		},
	})
	opts = append(opts, aliases)
	var resource CustomImage
	err := ctx.RegisterResource("azure-native:devtestlab/v20160515:CustomImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCustomImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomImageState, opts ...pulumi.ResourceOption) (*CustomImage, error) {
	var resource CustomImage
	err := ctx.ReadResource("azure-native:devtestlab/v20160515:CustomImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type customImageState struct {
}

type CustomImageState struct {
}

func (CustomImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*customImageState)(nil)).Elem()
}

type customImageArgs struct {
	Author            *string                      `pulumi:"author"`
	Description       *string                      `pulumi:"description"`
	LabName           string                       `pulumi:"labName"`
	Location          *string                      `pulumi:"location"`
	ManagedImageId    *string                      `pulumi:"managedImageId"`
	Name              *string                      `pulumi:"name"`
	ProvisioningState *string                      `pulumi:"provisioningState"`
	ResourceGroupName string                       `pulumi:"resourceGroupName"`
	Tags              map[string]string            `pulumi:"tags"`
	UniqueIdentifier  *string                      `pulumi:"uniqueIdentifier"`
	Vhd               *CustomImagePropertiesCustom `pulumi:"vhd"`
	Vm                *CustomImagePropertiesFromVm `pulumi:"vm"`
}


type CustomImageArgs struct {
	Author            pulumi.StringPtrInput
	Description       pulumi.StringPtrInput
	LabName           pulumi.StringInput
	Location          pulumi.StringPtrInput
	ManagedImageId    pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	ProvisioningState pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	UniqueIdentifier  pulumi.StringPtrInput
	Vhd               CustomImagePropertiesCustomPtrInput
	Vm                CustomImagePropertiesFromVmPtrInput
}

func (CustomImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customImageArgs)(nil)).Elem()
}

type CustomImageInput interface {
	pulumi.Input

	ToCustomImageOutput() CustomImageOutput
	ToCustomImageOutputWithContext(ctx context.Context) CustomImageOutput
}

func (*CustomImage) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomImage)(nil))
}

func (i *CustomImage) ToCustomImageOutput() CustomImageOutput {
	return i.ToCustomImageOutputWithContext(context.Background())
}

func (i *CustomImage) ToCustomImageOutputWithContext(ctx context.Context) CustomImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImageOutput)
}

type CustomImageOutput struct{ *pulumi.OutputState }

func (CustomImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomImage)(nil))
}

func (o CustomImageOutput) ToCustomImageOutput() CustomImageOutput {
	return o
}

func (o CustomImageOutput) ToCustomImageOutputWithContext(ctx context.Context) CustomImageOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomImageInput)(nil)).Elem(), &CustomImage{})
	pulumi.RegisterOutputType(CustomImageOutput{})
}
