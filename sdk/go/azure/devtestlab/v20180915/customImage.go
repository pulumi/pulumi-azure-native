


package v20180915

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
			Type: pulumi.String("azure-native:devtestlab:CustomImage"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20150521preview:CustomImage"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:CustomImage"),
		},
	})
	opts = append(opts, aliases)
	var resource CustomImage
	err := ctx.RegisterResource("azure-native:devtestlab/v20180915:CustomImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCustomImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomImageState, opts ...pulumi.ResourceOption) (*CustomImage, error) {
	var resource CustomImage
	err := ctx.ReadResource("azure-native:devtestlab/v20180915:CustomImage", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((**CustomImage)(nil)).Elem()
}

func (i *CustomImage) ToCustomImageOutput() CustomImageOutput {
	return i.ToCustomImageOutputWithContext(context.Background())
}

func (i *CustomImage) ToCustomImageOutputWithContext(ctx context.Context) CustomImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImageOutput)
}

type CustomImageOutput struct{ *pulumi.OutputState }

func (CustomImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomImage)(nil)).Elem()
}

func (o CustomImageOutput) ToCustomImageOutput() CustomImageOutput {
	return o
}

func (o CustomImageOutput) ToCustomImageOutputWithContext(ctx context.Context) CustomImageOutput {
	return o
}

func (o CustomImageOutput) Author() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.StringPtrOutput { return v.Author }).(pulumi.StringPtrOutput)
}

func (o CustomImageOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

func (o CustomImageOutput) CustomImagePlan() CustomImagePropertiesFromPlanResponsePtrOutput {
	return o.ApplyT(func(v *CustomImage) CustomImagePropertiesFromPlanResponsePtrOutput { return v.CustomImagePlan }).(CustomImagePropertiesFromPlanResponsePtrOutput)
}

func (o CustomImageOutput) DataDiskStorageInfo() DataDiskStorageTypeInfoResponseArrayOutput {
	return o.ApplyT(func(v *CustomImage) DataDiskStorageTypeInfoResponseArrayOutput { return v.DataDiskStorageInfo }).(DataDiskStorageTypeInfoResponseArrayOutput)
}

func (o CustomImageOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CustomImageOutput) IsPlanAuthorized() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.BoolPtrOutput { return v.IsPlanAuthorized }).(pulumi.BoolPtrOutput)
}

func (o CustomImageOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o CustomImageOutput) ManagedImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.StringPtrOutput { return v.ManagedImageId }).(pulumi.StringPtrOutput)
}

func (o CustomImageOutput) ManagedSnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.StringPtrOutput { return v.ManagedSnapshotId }).(pulumi.StringPtrOutput)
}

func (o CustomImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CustomImageOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o CustomImageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o CustomImageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o CustomImageOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomImage) pulumi.StringOutput { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func (o CustomImageOutput) Vhd() CustomImagePropertiesCustomResponsePtrOutput {
	return o.ApplyT(func(v *CustomImage) CustomImagePropertiesCustomResponsePtrOutput { return v.Vhd }).(CustomImagePropertiesCustomResponsePtrOutput)
}

func (o CustomImageOutput) Vm() CustomImagePropertiesFromVmResponsePtrOutput {
	return o.ApplyT(func(v *CustomImage) CustomImagePropertiesFromVmResponsePtrOutput { return v.Vm }).(CustomImagePropertiesFromVmResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomImageOutput{})
}
