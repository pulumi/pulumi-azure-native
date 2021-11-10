


package devtestlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomImage struct {
	pulumi.CustomResourceState

	Author              pulumi.StringPtrOutput                         `pulumi:"author"`
	CreationDate        pulumi.StringOutput                            `pulumi:"creationDate"`
	CustomImagePlan     CustomImagePropertiesFromPlanResponsePtrOutput `pulumi:"customImagePlan"`
	DataDiskStorageInfo DataDiskStorageTypeInfoResponseArrayOutput     `pulumi:"dataDiskStorageInfo"`
	Description         pulumi.StringPtrOutput                         `pulumi:"description"`
	IsPlanAuthorized    pulumi.BoolPtrOutput                           `pulumi:"isPlanAuthorized"`
	Location            pulumi.StringPtrOutput                         `pulumi:"location"`
	ManagedImageId      pulumi.StringPtrOutput                         `pulumi:"managedImageId"`
	ManagedSnapshotId   pulumi.StringPtrOutput                         `pulumi:"managedSnapshotId"`
	Name                pulumi.StringOutput                            `pulumi:"name"`
	ProvisioningState   pulumi.StringOutput                            `pulumi:"provisioningState"`
	Tags                pulumi.StringMapOutput                         `pulumi:"tags"`
	Type                pulumi.StringOutput                            `pulumi:"type"`
	UniqueIdentifier    pulumi.StringOutput                            `pulumi:"uniqueIdentifier"`
	Vhd                 CustomImagePropertiesCustomResponsePtrOutput   `pulumi:"vhd"`
	Vm                  CustomImagePropertiesFromVmResponsePtrOutput   `pulumi:"vm"`
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
			Type: pulumi.String("azure-native:devtestlab/v20150521preview:CustomImage"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:CustomImage"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:CustomImage"),
		},
	})
	opts = append(opts, aliases)
	var resource CustomImage
	err := ctx.RegisterResource("azure-native:devtestlab:CustomImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCustomImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomImageState, opts ...pulumi.ResourceOption) (*CustomImage, error) {
	var resource CustomImage
	err := ctx.ReadResource("azure-native:devtestlab:CustomImage", name, id, state, &resource, opts...)
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
	Author              *string                        `pulumi:"author"`
	CustomImagePlan     *CustomImagePropertiesFromPlan `pulumi:"customImagePlan"`
	DataDiskStorageInfo []DataDiskStorageTypeInfo      `pulumi:"dataDiskStorageInfo"`
	Description         *string                        `pulumi:"description"`
	IsPlanAuthorized    *bool                          `pulumi:"isPlanAuthorized"`
	LabName             string                         `pulumi:"labName"`
	Location            *string                        `pulumi:"location"`
	ManagedImageId      *string                        `pulumi:"managedImageId"`
	ManagedSnapshotId   *string                        `pulumi:"managedSnapshotId"`
	Name                *string                        `pulumi:"name"`
	ResourceGroupName   string                         `pulumi:"resourceGroupName"`
	Tags                map[string]string              `pulumi:"tags"`
	Vhd                 *CustomImagePropertiesCustom   `pulumi:"vhd"`
	Vm                  *CustomImagePropertiesFromVm   `pulumi:"vm"`
}


type CustomImageArgs struct {
	Author              pulumi.StringPtrInput
	CustomImagePlan     CustomImagePropertiesFromPlanPtrInput
	DataDiskStorageInfo DataDiskStorageTypeInfoArrayInput
	Description         pulumi.StringPtrInput
	IsPlanAuthorized    pulumi.BoolPtrInput
	LabName             pulumi.StringInput
	Location            pulumi.StringPtrInput
	ManagedImageId      pulumi.StringPtrInput
	ManagedSnapshotId   pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	Tags                pulumi.StringMapInput
	Vhd                 CustomImagePropertiesCustomPtrInput
	Vm                  CustomImagePropertiesFromVmPtrInput
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
	pulumi.RegisterOutputType(CustomImageOutput{})
}
