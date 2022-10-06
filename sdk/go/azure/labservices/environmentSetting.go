


package labservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnvironmentSetting struct {
	pulumi.CustomResourceState

	ConfigurationState    pulumi.StringPtrOutput              `pulumi:"configurationState"`
	Description           pulumi.StringPtrOutput              `pulumi:"description"`
	LastChanged           pulumi.StringOutput                 `pulumi:"lastChanged"`
	LastPublished         pulumi.StringOutput                 `pulumi:"lastPublished"`
	LatestOperationResult LatestOperationResultResponseOutput `pulumi:"latestOperationResult"`
	Location              pulumi.StringPtrOutput              `pulumi:"location"`
	Name                  pulumi.StringOutput                 `pulumi:"name"`
	ProvisioningState     pulumi.StringPtrOutput              `pulumi:"provisioningState"`
	PublishingState       pulumi.StringOutput                 `pulumi:"publishingState"`
	ResourceSettings      ResourceSettingsResponseOutput      `pulumi:"resourceSettings"`
	Tags                  pulumi.StringMapOutput              `pulumi:"tags"`
	Title                 pulumi.StringPtrOutput              `pulumi:"title"`
	Type                  pulumi.StringOutput                 `pulumi:"type"`
	UniqueIdentifier      pulumi.StringPtrOutput              `pulumi:"uniqueIdentifier"`
}


func NewEnvironmentSetting(ctx *pulumi.Context,
	name string, args *EnvironmentSettingArgs, opts ...pulumi.ResourceOption) (*EnvironmentSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabAccountName == nil {
		return nil, errors.New("invalid value for required argument 'LabAccountName'")
	}
	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceSettings == nil {
		return nil, errors.New("invalid value for required argument 'ResourceSettings'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:labservices/v20181015:EnvironmentSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource EnvironmentSetting
	err := ctx.RegisterResource("azure-native:labservices:EnvironmentSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEnvironmentSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentSettingState, opts ...pulumi.ResourceOption) (*EnvironmentSetting, error) {
	var resource EnvironmentSetting
	err := ctx.ReadResource("azure-native:labservices:EnvironmentSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type environmentSettingState struct {
}

type EnvironmentSettingState struct {
}

func (EnvironmentSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentSettingState)(nil)).Elem()
}

type environmentSettingArgs struct {
	ConfigurationState     *string           `pulumi:"configurationState"`
	Description            *string           `pulumi:"description"`
	EnvironmentSettingName *string           `pulumi:"environmentSettingName"`
	LabAccountName         string            `pulumi:"labAccountName"`
	LabName                string            `pulumi:"labName"`
	Location               *string           `pulumi:"location"`
	ProvisioningState      *string           `pulumi:"provisioningState"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	ResourceSettings       ResourceSettings  `pulumi:"resourceSettings"`
	Tags                   map[string]string `pulumi:"tags"`
	Title                  *string           `pulumi:"title"`
	UniqueIdentifier       *string           `pulumi:"uniqueIdentifier"`
}


type EnvironmentSettingArgs struct {
	ConfigurationState     pulumi.StringPtrInput
	Description            pulumi.StringPtrInput
	EnvironmentSettingName pulumi.StringPtrInput
	LabAccountName         pulumi.StringInput
	LabName                pulumi.StringInput
	Location               pulumi.StringPtrInput
	ProvisioningState      pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	ResourceSettings       ResourceSettingsInput
	Tags                   pulumi.StringMapInput
	Title                  pulumi.StringPtrInput
	UniqueIdentifier       pulumi.StringPtrInput
}

func (EnvironmentSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentSettingArgs)(nil)).Elem()
}

type EnvironmentSettingInput interface {
	pulumi.Input

	ToEnvironmentSettingOutput() EnvironmentSettingOutput
	ToEnvironmentSettingOutputWithContext(ctx context.Context) EnvironmentSettingOutput
}

func (*EnvironmentSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentSetting)(nil)).Elem()
}

func (i *EnvironmentSetting) ToEnvironmentSettingOutput() EnvironmentSettingOutput {
	return i.ToEnvironmentSettingOutputWithContext(context.Background())
}

func (i *EnvironmentSetting) ToEnvironmentSettingOutputWithContext(ctx context.Context) EnvironmentSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentSettingOutput)
}

type EnvironmentSettingOutput struct{ *pulumi.OutputState }

func (EnvironmentSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentSetting)(nil)).Elem()
}

func (o EnvironmentSettingOutput) ToEnvironmentSettingOutput() EnvironmentSettingOutput {
	return o
}

func (o EnvironmentSettingOutput) ToEnvironmentSettingOutputWithContext(ctx context.Context) EnvironmentSettingOutput {
	return o
}

func (o EnvironmentSettingOutput) ConfigurationState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringPtrOutput { return v.ConfigurationState }).(pulumi.StringPtrOutput)
}

func (o EnvironmentSettingOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o EnvironmentSettingOutput) LastChanged() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringOutput { return v.LastChanged }).(pulumi.StringOutput)
}

func (o EnvironmentSettingOutput) LastPublished() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringOutput { return v.LastPublished }).(pulumi.StringOutput)
}

func (o EnvironmentSettingOutput) LatestOperationResult() LatestOperationResultResponseOutput {
	return o.ApplyT(func(v *EnvironmentSetting) LatestOperationResultResponseOutput { return v.LatestOperationResult }).(LatestOperationResultResponseOutput)
}

func (o EnvironmentSettingOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o EnvironmentSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EnvironmentSettingOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o EnvironmentSettingOutput) PublishingState() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringOutput { return v.PublishingState }).(pulumi.StringOutput)
}

func (o EnvironmentSettingOutput) ResourceSettings() ResourceSettingsResponseOutput {
	return o.ApplyT(func(v *EnvironmentSetting) ResourceSettingsResponseOutput { return v.ResourceSettings }).(ResourceSettingsResponseOutput)
}

func (o EnvironmentSettingOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o EnvironmentSettingOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringPtrOutput { return v.Title }).(pulumi.StringPtrOutput)
}

func (o EnvironmentSettingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o EnvironmentSettingOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringPtrOutput { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EnvironmentSettingOutput{})
}
