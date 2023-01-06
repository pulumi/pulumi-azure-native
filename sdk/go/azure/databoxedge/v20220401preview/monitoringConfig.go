


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MonitoringConfig struct {
	pulumi.CustomResourceState

	MetricConfigurations MetricConfigurationResponseArrayOutput `pulumi:"metricConfigurations"`
	Name                 pulumi.StringOutput                    `pulumi:"name"`
	SystemData           SystemDataResponseOutput               `pulumi:"systemData"`
	Type                 pulumi.StringOutput                    `pulumi:"type"`
}


func NewMonitoringConfig(ctx *pulumi.Context,
	name string, args *MonitoringConfigArgs, opts ...pulumi.ResourceOption) (*MonitoringConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.MetricConfigurations == nil {
		return nil, errors.New("invalid value for required argument 'MetricConfigurations'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RoleName == nil {
		return nil, errors.New("invalid value for required argument 'RoleName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge:MonitoringConfig"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:MonitoringConfig"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:MonitoringConfig"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:MonitoringConfig"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:MonitoringConfig"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:MonitoringConfig"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:MonitoringConfig"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:MonitoringConfig"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:MonitoringConfig"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20221201preview:MonitoringConfig"),
		},
	})
	opts = append(opts, aliases)
	var resource MonitoringConfig
	err := ctx.RegisterResource("azure-native:databoxedge/v20220401preview:MonitoringConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMonitoringConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitoringConfigState, opts ...pulumi.ResourceOption) (*MonitoringConfig, error) {
	var resource MonitoringConfig
	err := ctx.ReadResource("azure-native:databoxedge/v20220401preview:MonitoringConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type monitoringConfigState struct {
}

type MonitoringConfigState struct {
}

func (MonitoringConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringConfigState)(nil)).Elem()
}

type monitoringConfigArgs struct {
	DeviceName           string                `pulumi:"deviceName"`
	MetricConfigurations []MetricConfiguration `pulumi:"metricConfigurations"`
	ResourceGroupName    string                `pulumi:"resourceGroupName"`
	RoleName             string                `pulumi:"roleName"`
}


type MonitoringConfigArgs struct {
	DeviceName           pulumi.StringInput
	MetricConfigurations MetricConfigurationArrayInput
	ResourceGroupName    pulumi.StringInput
	RoleName             pulumi.StringInput
}

func (MonitoringConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringConfigArgs)(nil)).Elem()
}

type MonitoringConfigInput interface {
	pulumi.Input

	ToMonitoringConfigOutput() MonitoringConfigOutput
	ToMonitoringConfigOutputWithContext(ctx context.Context) MonitoringConfigOutput
}

func (*MonitoringConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringConfig)(nil)).Elem()
}

func (i *MonitoringConfig) ToMonitoringConfigOutput() MonitoringConfigOutput {
	return i.ToMonitoringConfigOutputWithContext(context.Background())
}

func (i *MonitoringConfig) ToMonitoringConfigOutputWithContext(ctx context.Context) MonitoringConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringConfigOutput)
}

type MonitoringConfigOutput struct{ *pulumi.OutputState }

func (MonitoringConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringConfig)(nil)).Elem()
}

func (o MonitoringConfigOutput) ToMonitoringConfigOutput() MonitoringConfigOutput {
	return o
}

func (o MonitoringConfigOutput) ToMonitoringConfigOutputWithContext(ctx context.Context) MonitoringConfigOutput {
	return o
}

func (o MonitoringConfigOutput) MetricConfigurations() MetricConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *MonitoringConfig) MetricConfigurationResponseArrayOutput { return v.MetricConfigurations }).(MetricConfigurationResponseArrayOutput)
}

func (o MonitoringConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MonitoringConfigOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MonitoringConfig) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MonitoringConfigOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringConfig) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MonitoringConfigOutput{})
}
