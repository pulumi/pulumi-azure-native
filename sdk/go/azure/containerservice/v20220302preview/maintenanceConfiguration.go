


package v20220302preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MaintenanceConfiguration struct {
	pulumi.CustomResourceState

	Name           pulumi.StringOutput           `pulumi:"name"`
	NotAllowedTime TimeSpanResponseArrayOutput   `pulumi:"notAllowedTime"`
	SystemData     SystemDataResponseOutput      `pulumi:"systemData"`
	TimeInWeek     TimeInWeekResponseArrayOutput `pulumi:"timeInWeek"`
	Type           pulumi.StringOutput           `pulumi:"type"`
}


func NewMaintenanceConfiguration(ctx *pulumi.Context,
	name string, args *MaintenanceConfigurationArgs, opts ...pulumi.ResourceOption) (*MaintenanceConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerservice:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20201201:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210201:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210301:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210501:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210701:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210801:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210901:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20211001:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20211101preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220101:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220102preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220201:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220202preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220301:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220401:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220402preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220502preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220601:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220602preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220701:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220702preview:MaintenanceConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220802preview:MaintenanceConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource MaintenanceConfiguration
	err := ctx.RegisterResource("azure-native:containerservice/v20220302preview:MaintenanceConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMaintenanceConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MaintenanceConfigurationState, opts ...pulumi.ResourceOption) (*MaintenanceConfiguration, error) {
	var resource MaintenanceConfiguration
	err := ctx.ReadResource("azure-native:containerservice/v20220302preview:MaintenanceConfiguration", name, id, state, &resource, opts...)
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
	ConfigName        *string      `pulumi:"configName"`
	NotAllowedTime    []TimeSpan   `pulumi:"notAllowedTime"`
	ResourceGroupName string       `pulumi:"resourceGroupName"`
	ResourceName      string       `pulumi:"resourceName"`
	TimeInWeek        []TimeInWeek `pulumi:"timeInWeek"`
}


type MaintenanceConfigurationArgs struct {
	ConfigName        pulumi.StringPtrInput
	NotAllowedTime    TimeSpanArrayInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringInput
	TimeInWeek        TimeInWeekArrayInput
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

func (o MaintenanceConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MaintenanceConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MaintenanceConfigurationOutput) NotAllowedTime() TimeSpanResponseArrayOutput {
	return o.ApplyT(func(v *MaintenanceConfiguration) TimeSpanResponseArrayOutput { return v.NotAllowedTime }).(TimeSpanResponseArrayOutput)
}

func (o MaintenanceConfigurationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MaintenanceConfiguration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MaintenanceConfigurationOutput) TimeInWeek() TimeInWeekResponseArrayOutput {
	return o.ApplyT(func(v *MaintenanceConfiguration) TimeInWeekResponseArrayOutput { return v.TimeInWeek }).(TimeInWeekResponseArrayOutput)
}

func (o MaintenanceConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MaintenanceConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MaintenanceConfigurationOutput{})
}
