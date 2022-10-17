


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type MaintenanceConfiguration struct {
	pulumi.CustomResourceState

	ExtensionProperties pulumi.StringMapOutput `pulumi:"extensionProperties"`
	Location            pulumi.StringPtrOutput `pulumi:"location"`
	MaintenanceScope    pulumi.StringPtrOutput `pulumi:"maintenanceScope"`
	Name                pulumi.StringOutput    `pulumi:"name"`
	Namespace           pulumi.StringPtrOutput `pulumi:"namespace"`
	Tags                pulumi.StringMapOutput `pulumi:"tags"`
	Type                pulumi.StringOutput    `pulumi:"type"`
}


func NewMaintenanceConfiguration(ctx *pulumi.Context,
	name string, args *MaintenanceConfigurationArgs, opts ...pulumi.ResourceOption) (*MaintenanceConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:maintenance:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20200401:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20200701preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20210401preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20210501:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20210901preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20220701preview:MaintenanceConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource MaintenanceConfiguration
	err := ctx.RegisterResource("azure-native:maintenance/v20180601preview:MaintenanceConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMaintenanceConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MaintenanceConfigurationState, opts ...pulumi.ResourceOption) (*MaintenanceConfiguration, error) {
	var resource MaintenanceConfiguration
	err := ctx.ReadResource("azure-native:maintenance/v20180601preview:MaintenanceConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type maintenanceConfigurationState struct {
}

type MaintenanceConfigurationState struct {
}

func (MaintenanceConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceConfigurationState)(nil)).Elem()
}

type maintenanceConfigurationArgs struct {
	ExtensionProperties map[string]string `pulumi:"extensionProperties"`
	Location            *string           `pulumi:"location"`
	MaintenanceScope    *string           `pulumi:"maintenanceScope"`
	Namespace           *string           `pulumi:"namespace"`
	ResourceGroupName   string            `pulumi:"resourceGroupName"`
	ResourceName        *string           `pulumi:"resourceName"`
	Tags                map[string]string `pulumi:"tags"`
}


type MaintenanceConfigurationArgs struct {
	ExtensionProperties pulumi.StringMapInput
	Location            pulumi.StringPtrInput
	MaintenanceScope    pulumi.StringPtrInput
	Namespace           pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	ResourceName        pulumi.StringPtrInput
	Tags                pulumi.StringMapInput
}

func (MaintenanceConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceConfigurationArgs)(nil)).Elem()
}

type MaintenanceConfigurationInput interface {
	pulumi.Input

	ToMaintenanceConfigurationOutput() MaintenanceConfigurationOutput
	ToMaintenanceConfigurationOutputWithContext(ctx context.Context) MaintenanceConfigurationOutput
}

func (*MaintenanceConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceConfiguration)(nil)).Elem()
}

func (i *MaintenanceConfiguration) ToMaintenanceConfigurationOutput() MaintenanceConfigurationOutput {
	return i.ToMaintenanceConfigurationOutputWithContext(context.Background())
}

func (i *MaintenanceConfiguration) ToMaintenanceConfigurationOutputWithContext(ctx context.Context) MaintenanceConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceConfigurationOutput)
}

type MaintenanceConfigurationOutput struct{ *pulumi.OutputState }

func (MaintenanceConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceConfiguration)(nil)).Elem()
}

func (o MaintenanceConfigurationOutput) ToMaintenanceConfigurationOutput() MaintenanceConfigurationOutput {
	return o
}

func (o MaintenanceConfigurationOutput) ToMaintenanceConfigurationOutputWithContext(ctx context.Context) MaintenanceConfigurationOutput {
	return o
}

func (o MaintenanceConfigurationOutput) ExtensionProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MaintenanceConfiguration) pulumi.StringMapOutput { return v.ExtensionProperties }).(pulumi.StringMapOutput)
}

func (o MaintenanceConfigurationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceConfiguration) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o MaintenanceConfigurationOutput) MaintenanceScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceConfiguration) pulumi.StringPtrOutput { return v.MaintenanceScope }).(pulumi.StringPtrOutput)
}

func (o MaintenanceConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MaintenanceConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MaintenanceConfigurationOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceConfiguration) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o MaintenanceConfigurationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MaintenanceConfiguration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o MaintenanceConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MaintenanceConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MaintenanceConfigurationOutput{})
}
