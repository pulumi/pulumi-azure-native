


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SiteSensor struct {
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


func NewSiteSensor(ctx *pulumi.Context,
	name string, args *SiteSensorArgs, opts ...pulumi.ResourceOption) (*SiteSensor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IotDefenderLocation == nil {
		return nil, errors.New("invalid value for required argument 'IotDefenderLocation'")
	}
	if args.SiteName == nil {
		return nil, errors.New("invalid value for required argument 'SiteName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:iotsecurity/v20210901preview:SiteSensor"),
		},
		{
			Type: pulumi.String("azure-native:iotsecurity:SiteSensor"),
		},
		{
			Type: pulumi.String("azure-nextgen:iotsecurity:SiteSensor"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteSensor
	err := ctx.RegisterResource("azure-native:iotsecurity/v20210901preview:SiteSensor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteSensor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteSensorState, opts ...pulumi.ResourceOption) (*SiteSensor, error) {
	var resource SiteSensor
	err := ctx.ReadResource("azure-native:iotsecurity/v20210901preview:SiteSensor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteSensorState struct {
}

type SiteSensorState struct {
}

func (SiteSensorState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteSensorState)(nil)).Elem()
}

type siteSensorArgs struct {
	IotDefenderLocation string  `pulumi:"iotDefenderLocation"`
	SensorName          *string `pulumi:"sensorName"`
	SensorType          *string `pulumi:"sensorType"`
	SiteName            string  `pulumi:"siteName"`
	TiAutomaticUpdates  *bool   `pulumi:"tiAutomaticUpdates"`
	Zone                *string `pulumi:"zone"`
}


type SiteSensorArgs struct {
	IotDefenderLocation pulumi.StringInput
	SensorName          pulumi.StringPtrInput
	SensorType          pulumi.StringPtrInput
	SiteName            pulumi.StringInput
	TiAutomaticUpdates  pulumi.BoolPtrInput
	Zone                pulumi.StringPtrInput
}

func (SiteSensorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteSensorArgs)(nil)).Elem()
}

type SiteSensorInput interface {
	pulumi.Input

	ToSiteSensorOutput() SiteSensorOutput
	ToSiteSensorOutputWithContext(ctx context.Context) SiteSensorOutput
}

func (*SiteSensor) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteSensor)(nil))
}

func (i *SiteSensor) ToSiteSensorOutput() SiteSensorOutput {
	return i.ToSiteSensorOutputWithContext(context.Background())
}

func (i *SiteSensor) ToSiteSensorOutputWithContext(ctx context.Context) SiteSensorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteSensorOutput)
}

type SiteSensorOutput struct{ *pulumi.OutputState }

func (SiteSensorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteSensor)(nil))
}

func (o SiteSensorOutput) ToSiteSensorOutput() SiteSensorOutput {
	return o
}

func (o SiteSensorOutput) ToSiteSensorOutputWithContext(ctx context.Context) SiteSensorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SiteSensorOutput{})
}
