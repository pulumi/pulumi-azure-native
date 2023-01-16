


package v20221201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Update struct {
	pulumi.CustomResourceState

	AdditionalProperties pulumi.StringPtrOutput                `pulumi:"additionalProperties"`
	AvailabilityType     pulumi.StringPtrOutput                `pulumi:"availabilityType"`
	Description          pulumi.StringPtrOutput                `pulumi:"description"`
	DisplayName          pulumi.StringPtrOutput                `pulumi:"displayName"`
	HealthCheckDate      pulumi.StringPtrOutput                `pulumi:"healthCheckDate"`
	InstalledDate        pulumi.StringPtrOutput                `pulumi:"installedDate"`
	Location             pulumi.StringPtrOutput                `pulumi:"location"`
	Name                 pulumi.StringOutput                   `pulumi:"name"`
	NotifyMessage        pulumi.StringPtrOutput                `pulumi:"notifyMessage"`
	PackagePath          pulumi.StringPtrOutput                `pulumi:"packagePath"`
	PackageSizeInMb      pulumi.Float64PtrOutput               `pulumi:"packageSizeInMb"`
	PackageType          pulumi.StringPtrOutput                `pulumi:"packageType"`
	Prerequisites        UpdatePrerequisiteResponseArrayOutput `pulumi:"prerequisites"`
	ProgressPercentage   pulumi.Float64PtrOutput               `pulumi:"progressPercentage"`
	ProvisioningState    pulumi.StringOutput                   `pulumi:"provisioningState"`
	Publisher            pulumi.StringPtrOutput                `pulumi:"publisher"`
	ReleaseLink          pulumi.StringPtrOutput                `pulumi:"releaseLink"`
	State                pulumi.StringPtrOutput                `pulumi:"state"`
	SystemData           SystemDataResponseOutput              `pulumi:"systemData"`
	Type                 pulumi.StringOutput                   `pulumi:"type"`
	Version              pulumi.StringPtrOutput                `pulumi:"version"`
}


func NewUpdate(ctx *pulumi.Context,
	name string, args *UpdateArgs, opts ...pulumi.ResourceOption) (*Update, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Update
	err := ctx.RegisterResource("azure-native:azurestackhci/v20221201:Update", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetUpdate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UpdateState, opts ...pulumi.ResourceOption) (*Update, error) {
	var resource Update
	err := ctx.ReadResource("azure-native:azurestackhci/v20221201:Update", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type updateState struct {
}

type UpdateState struct {
}

func (UpdateState) ElementType() reflect.Type {
	return reflect.TypeOf((*updateState)(nil)).Elem()
}

type updateArgs struct {
	AdditionalProperties *string              `pulumi:"additionalProperties"`
	AvailabilityType     *string              `pulumi:"availabilityType"`
	ClusterName          string               `pulumi:"clusterName"`
	Description          *string              `pulumi:"description"`
	DisplayName          *string              `pulumi:"displayName"`
	HealthCheckDate      *string              `pulumi:"healthCheckDate"`
	InstalledDate        *string              `pulumi:"installedDate"`
	Location             *string              `pulumi:"location"`
	NotifyMessage        *string              `pulumi:"notifyMessage"`
	PackagePath          *string              `pulumi:"packagePath"`
	PackageSizeInMb      *float64             `pulumi:"packageSizeInMb"`
	PackageType          *string              `pulumi:"packageType"`
	Prerequisites        []UpdatePrerequisite `pulumi:"prerequisites"`
	ProgressPercentage   *float64             `pulumi:"progressPercentage"`
	Publisher            *string              `pulumi:"publisher"`
	ReleaseLink          *string              `pulumi:"releaseLink"`
	ResourceGroupName    string               `pulumi:"resourceGroupName"`
	State                *string              `pulumi:"state"`
	UpdateName           *string              `pulumi:"updateName"`
	Version              *string              `pulumi:"version"`
}


type UpdateArgs struct {
	AdditionalProperties pulumi.StringPtrInput
	AvailabilityType     pulumi.StringPtrInput
	ClusterName          pulumi.StringInput
	Description          pulumi.StringPtrInput
	DisplayName          pulumi.StringPtrInput
	HealthCheckDate      pulumi.StringPtrInput
	InstalledDate        pulumi.StringPtrInput
	Location             pulumi.StringPtrInput
	NotifyMessage        pulumi.StringPtrInput
	PackagePath          pulumi.StringPtrInput
	PackageSizeInMb      pulumi.Float64PtrInput
	PackageType          pulumi.StringPtrInput
	Prerequisites        UpdatePrerequisiteArrayInput
	ProgressPercentage   pulumi.Float64PtrInput
	Publisher            pulumi.StringPtrInput
	ReleaseLink          pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	State                pulumi.StringPtrInput
	UpdateName           pulumi.StringPtrInput
	Version              pulumi.StringPtrInput
}

func (UpdateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*updateArgs)(nil)).Elem()
}

type UpdateInput interface {
	pulumi.Input

	ToUpdateOutput() UpdateOutput
	ToUpdateOutputWithContext(ctx context.Context) UpdateOutput
}

func (*Update) ElementType() reflect.Type {
	return reflect.TypeOf((**Update)(nil)).Elem()
}

func (i *Update) ToUpdateOutput() UpdateOutput {
	return i.ToUpdateOutputWithContext(context.Background())
}

func (i *Update) ToUpdateOutputWithContext(ctx context.Context) UpdateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpdateOutput)
}

type UpdateOutput struct{ *pulumi.OutputState }

func (UpdateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Update)(nil)).Elem()
}

func (o UpdateOutput) ToUpdateOutput() UpdateOutput {
	return o
}

func (o UpdateOutput) ToUpdateOutputWithContext(ctx context.Context) UpdateOutput {
	return o
}

func (o UpdateOutput) AdditionalProperties() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.AdditionalProperties }).(pulumi.StringPtrOutput)
}

func (o UpdateOutput) AvailabilityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.AvailabilityType }).(pulumi.StringPtrOutput)
}

func (o UpdateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o UpdateOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o UpdateOutput) HealthCheckDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.HealthCheckDate }).(pulumi.StringPtrOutput)
}

func (o UpdateOutput) InstalledDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.InstalledDate }).(pulumi.StringPtrOutput)
}

func (o UpdateOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o UpdateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Update) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o UpdateOutput) NotifyMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.NotifyMessage }).(pulumi.StringPtrOutput)
}

func (o UpdateOutput) PackagePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.PackagePath }).(pulumi.StringPtrOutput)
}

func (o UpdateOutput) PackageSizeInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Update) pulumi.Float64PtrOutput { return v.PackageSizeInMb }).(pulumi.Float64PtrOutput)
}

func (o UpdateOutput) PackageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.PackageType }).(pulumi.StringPtrOutput)
}

func (o UpdateOutput) Prerequisites() UpdatePrerequisiteResponseArrayOutput {
	return o.ApplyT(func(v *Update) UpdatePrerequisiteResponseArrayOutput { return v.Prerequisites }).(UpdatePrerequisiteResponseArrayOutput)
}

func (o UpdateOutput) ProgressPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Update) pulumi.Float64PtrOutput { return v.ProgressPercentage }).(pulumi.Float64PtrOutput)
}

func (o UpdateOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Update) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o UpdateOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o UpdateOutput) ReleaseLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.ReleaseLink }).(pulumi.StringPtrOutput)
}

func (o UpdateOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

func (o UpdateOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Update) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o UpdateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Update) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o UpdateOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Update) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(UpdateOutput{})
}
