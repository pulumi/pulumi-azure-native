


package appplatform

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MonitoringSetting struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                       `pulumi:"name"`
	Properties MonitoringSettingPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                       `pulumi:"type"`
}


func NewMonitoringSetting(ctx *pulumi.Context,
	name string, args *MonitoringSettingArgs, opts ...pulumi.ResourceOption) (*MonitoringSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform/v20200701:MonitoringSetting"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20201101preview:MonitoringSetting"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210601preview:MonitoringSetting"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210901preview:MonitoringSetting"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:MonitoringSetting"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:MonitoringSetting"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220401:MonitoringSetting"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:MonitoringSetting"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:MonitoringSetting"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:MonitoringSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource MonitoringSetting
	err := ctx.RegisterResource("azure-native:appplatform:MonitoringSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMonitoringSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitoringSettingState, opts ...pulumi.ResourceOption) (*MonitoringSetting, error) {
	var resource MonitoringSetting
	err := ctx.ReadResource("azure-native:appplatform:MonitoringSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type monitoringSettingState struct {
}

type MonitoringSettingState struct {
}

func (MonitoringSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringSettingState)(nil)).Elem()
}

type monitoringSettingArgs struct {
	Properties        *MonitoringSettingProperties `pulumi:"properties"`
	ResourceGroupName string                       `pulumi:"resourceGroupName"`
	ServiceName       string                       `pulumi:"serviceName"`
}


type MonitoringSettingArgs struct {
	Properties        MonitoringSettingPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (MonitoringSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringSettingArgs)(nil)).Elem()
}

type MonitoringSettingInput interface {
	pulumi.Input

	ToMonitoringSettingOutput() MonitoringSettingOutput
	ToMonitoringSettingOutputWithContext(ctx context.Context) MonitoringSettingOutput
}

func (*MonitoringSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringSetting)(nil)).Elem()
}

func (i *MonitoringSetting) ToMonitoringSettingOutput() MonitoringSettingOutput {
	return i.ToMonitoringSettingOutputWithContext(context.Background())
}

func (i *MonitoringSetting) ToMonitoringSettingOutputWithContext(ctx context.Context) MonitoringSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringSettingOutput)
}

type MonitoringSettingOutput struct{ *pulumi.OutputState }

func (MonitoringSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringSetting)(nil)).Elem()
}

func (o MonitoringSettingOutput) ToMonitoringSettingOutput() MonitoringSettingOutput {
	return o
}

func (o MonitoringSettingOutput) ToMonitoringSettingOutputWithContext(ctx context.Context) MonitoringSettingOutput {
	return o
}

func (o MonitoringSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringSetting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MonitoringSettingOutput) Properties() MonitoringSettingPropertiesResponseOutput {
	return o.ApplyT(func(v *MonitoringSetting) MonitoringSettingPropertiesResponseOutput { return v.Properties }).(MonitoringSettingPropertiesResponseOutput)
}

func (o MonitoringSettingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringSetting) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MonitoringSettingOutput{})
}
