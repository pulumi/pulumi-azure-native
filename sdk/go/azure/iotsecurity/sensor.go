


package iotsecurity

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Sensor struct {
	pulumi.CustomResourceState

	ConnectivityTime   pulumi.StringOutput      `pulumi:"connectivityTime"`
	DynamicLearning    pulumi.BoolOutput        `pulumi:"dynamicLearning"`
	LearningMode       pulumi.BoolOutput        `pulumi:"learningMode"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	SensorStatus       pulumi.StringOutput      `pulumi:"sensorStatus"`
	SensorType         pulumi.StringPtrOutput   `pulumi:"sensorType"`
	SensorVersion      pulumi.StringOutput      `pulumi:"sensorVersion"`
	SystemData         SystemDataResponseOutput `pulumi:"systemData"`
	TiAutomaticUpdates pulumi.BoolPtrOutput     `pulumi:"tiAutomaticUpdates"`
	TiStatus           pulumi.StringOutput      `pulumi:"tiStatus"`
	TiVersion          pulumi.StringOutput      `pulumi:"tiVersion"`
	Type               pulumi.StringOutput      `pulumi:"type"`
	Zone               pulumi.StringPtrOutput   `pulumi:"zone"`
}


func NewSensor(ctx *pulumi.Context,
	name string, args *SensorArgs, opts ...pulumi.ResourceOption) (*Sensor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:iotsecurity:Sensor"),
		},
		{
			Type: pulumi.String("azure-native:iotsecurity/v20210201preview:Sensor"),
		},
		{
			Type: pulumi.String("azure-nextgen:iotsecurity/v20210201preview:Sensor"),
		},
		{
			Type: pulumi.String("azure-native:iotsecurity/v20210901preview:Sensor"),
		},
		{
			Type: pulumi.String("azure-nextgen:iotsecurity/v20210901preview:Sensor"),
		},
	})
	opts = append(opts, aliases)
	var resource Sensor
	err := ctx.RegisterResource("azure-native:iotsecurity:Sensor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSensor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SensorState, opts ...pulumi.ResourceOption) (*Sensor, error) {
	var resource Sensor
	err := ctx.ReadResource("azure-native:iotsecurity:Sensor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sensorState struct {
}

type SensorState struct {
}

func (SensorState) ElementType() reflect.Type {
	return reflect.TypeOf((*sensorState)(nil)).Elem()
}

type sensorArgs struct {
	Scope              string  `pulumi:"scope"`
	SensorName         *string `pulumi:"sensorName"`
	SensorType         *string `pulumi:"sensorType"`
	TiAutomaticUpdates *bool   `pulumi:"tiAutomaticUpdates"`
	Zone               *string `pulumi:"zone"`
}


type SensorArgs struct {
	Scope              pulumi.StringInput
	SensorName         pulumi.StringPtrInput
	SensorType         pulumi.StringPtrInput
	TiAutomaticUpdates pulumi.BoolPtrInput
	Zone               pulumi.StringPtrInput
}

func (SensorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sensorArgs)(nil)).Elem()
}

type SensorInput interface {
	pulumi.Input

	ToSensorOutput() SensorOutput
	ToSensorOutputWithContext(ctx context.Context) SensorOutput
}

func (*Sensor) ElementType() reflect.Type {
	return reflect.TypeOf((*Sensor)(nil))
}

func (i *Sensor) ToSensorOutput() SensorOutput {
	return i.ToSensorOutputWithContext(context.Background())
}

func (i *Sensor) ToSensorOutputWithContext(ctx context.Context) SensorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SensorOutput)
}

type SensorOutput struct{ *pulumi.OutputState }

func (SensorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sensor)(nil))
}

func (o SensorOutput) ToSensorOutput() SensorOutput {
	return o
}

func (o SensorOutput) ToSensorOutputWithContext(ctx context.Context) SensorOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SensorInput)(nil)).Elem(), &Sensor{})
	pulumi.RegisterOutputType(SensorOutput{})
}
