


package v20200701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MaintenanceConfiguration struct {
	pulumi.CustomResourceState

	Duration            pulumi.StringPtrOutput `pulumi:"duration"`
	ExpirationDateTime  pulumi.StringPtrOutput `pulumi:"expirationDateTime"`
	ExtensionProperties pulumi.StringMapOutput `pulumi:"extensionProperties"`
	Location            pulumi.StringPtrOutput `pulumi:"location"`
	MaintenanceScope    pulumi.StringPtrOutput `pulumi:"maintenanceScope"`
	Name                pulumi.StringOutput    `pulumi:"name"`
	Namespace           pulumi.StringPtrOutput `pulumi:"namespace"`
	RecurEvery          pulumi.StringPtrOutput `pulumi:"recurEvery"`
	StartDateTime       pulumi.StringPtrOutput `pulumi:"startDateTime"`
	Tags                pulumi.StringMapOutput `pulumi:"tags"`
	TimeZone            pulumi.StringPtrOutput `pulumi:"timeZone"`
	Type                pulumi.StringOutput    `pulumi:"type"`
	Visibility          pulumi.StringPtrOutput `pulumi:"visibility"`
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
			Type: pulumi.String("azure-nextgen:maintenance/v20200701preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:maintenance:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-nextgen:maintenance:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20180601preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-nextgen:maintenance/v20180601preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20200401:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-nextgen:maintenance/v20200401:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20210401preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-nextgen:maintenance/v20210401preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20210501:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-nextgen:maintenance/v20210501:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20210901preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-nextgen:maintenance/v20210901preview:MaintenanceConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource MaintenanceConfiguration
	err := ctx.RegisterResource("azure-native:maintenance/v20200701preview:MaintenanceConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMaintenanceConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MaintenanceConfigurationState, opts ...pulumi.ResourceOption) (*MaintenanceConfiguration, error) {
	var resource MaintenanceConfiguration
	err := ctx.ReadResource("azure-native:maintenance/v20200701preview:MaintenanceConfiguration", name, id, state, &resource, opts...)
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
	Duration            *string           `pulumi:"duration"`
	ExpirationDateTime  *string           `pulumi:"expirationDateTime"`
	ExtensionProperties map[string]string `pulumi:"extensionProperties"`
	Location            *string           `pulumi:"location"`
	MaintenanceScope    *string           `pulumi:"maintenanceScope"`
	Namespace           *string           `pulumi:"namespace"`
	RecurEvery          *string           `pulumi:"recurEvery"`
	ResourceGroupName   string            `pulumi:"resourceGroupName"`
	ResourceName        *string           `pulumi:"resourceName"`
	StartDateTime       *string           `pulumi:"startDateTime"`
	Tags                map[string]string `pulumi:"tags"`
	TimeZone            *string           `pulumi:"timeZone"`
	Visibility          *string           `pulumi:"visibility"`
}


type MaintenanceConfigurationArgs struct {
	Duration            pulumi.StringPtrInput
	ExpirationDateTime  pulumi.StringPtrInput
	ExtensionProperties pulumi.StringMapInput
	Location            pulumi.StringPtrInput
	MaintenanceScope    pulumi.StringPtrInput
	Namespace           pulumi.StringPtrInput
	RecurEvery          pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	ResourceName        pulumi.StringPtrInput
	StartDateTime       pulumi.StringPtrInput
	Tags                pulumi.StringMapInput
	TimeZone            pulumi.StringPtrInput
	Visibility          pulumi.StringPtrInput
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
	return reflect.TypeOf((*MaintenanceConfiguration)(nil))
}

func (i *MaintenanceConfiguration) ToMaintenanceConfigurationOutput() MaintenanceConfigurationOutput {
	return i.ToMaintenanceConfigurationOutputWithContext(context.Background())
}

func (i *MaintenanceConfiguration) ToMaintenanceConfigurationOutputWithContext(ctx context.Context) MaintenanceConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceConfigurationOutput)
}

type MaintenanceConfigurationOutput struct{ *pulumi.OutputState }

func (MaintenanceConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceConfiguration)(nil))
}

func (o MaintenanceConfigurationOutput) ToMaintenanceConfigurationOutput() MaintenanceConfigurationOutput {
	return o
}

func (o MaintenanceConfigurationOutput) ToMaintenanceConfigurationOutputWithContext(ctx context.Context) MaintenanceConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MaintenanceConfigurationOutput{})
}
