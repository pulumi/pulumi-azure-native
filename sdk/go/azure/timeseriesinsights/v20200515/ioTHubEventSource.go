


package v20200515

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IoTHubEventSource struct {
	pulumi.CustomResourceState

	ConsumerGroupName     pulumi.StringOutput             `pulumi:"consumerGroupName"`
	CreationTime          pulumi.StringOutput             `pulumi:"creationTime"`
	EventSourceResourceId pulumi.StringOutput             `pulumi:"eventSourceResourceId"`
	IotHubName            pulumi.StringOutput             `pulumi:"iotHubName"`
	KeyName               pulumi.StringOutput             `pulumi:"keyName"`
	Kind                  pulumi.StringOutput             `pulumi:"kind"`
	LocalTimestamp        LocalTimestampResponsePtrOutput `pulumi:"localTimestamp"`
	Location              pulumi.StringOutput             `pulumi:"location"`
	Name                  pulumi.StringOutput             `pulumi:"name"`
	ProvisioningState     pulumi.StringOutput             `pulumi:"provisioningState"`
	Tags                  pulumi.StringMapOutput          `pulumi:"tags"`
	Time                  pulumi.StringPtrOutput          `pulumi:"time"`
	TimestampPropertyName pulumi.StringPtrOutput          `pulumi:"timestampPropertyName"`
	Type                  pulumi.StringOutput             `pulumi:"type"`
}


func NewIoTHubEventSource(ctx *pulumi.Context,
	name string, args *IoTHubEventSourceArgs, opts ...pulumi.ResourceOption) (*IoTHubEventSource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsumerGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerGroupName'")
	}
	if args.EnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentName'")
	}
	if args.EventSourceResourceId == nil {
		return nil, errors.New("invalid value for required argument 'EventSourceResourceId'")
	}
	if args.IotHubName == nil {
		return nil, errors.New("invalid value for required argument 'IotHubName'")
	}
	if args.KeyName == nil {
		return nil, errors.New("invalid value for required argument 'KeyName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SharedAccessKey == nil {
		return nil, errors.New("invalid value for required argument 'SharedAccessKey'")
	}
	args.Kind = pulumi.String("Microsoft.IoTHub")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:timeseriesinsights:IoTHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20170228preview:IoTHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20171115:IoTHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20180815preview:IoTHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210331preview:IoTHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210630preview:IoTHubEventSource"),
		},
	})
	opts = append(opts, aliases)
	var resource IoTHubEventSource
	err := ctx.RegisterResource("azure-native:timeseriesinsights/v20200515:IoTHubEventSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIoTHubEventSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IoTHubEventSourceState, opts ...pulumi.ResourceOption) (*IoTHubEventSource, error) {
	var resource IoTHubEventSource
	err := ctx.ReadResource("azure-native:timeseriesinsights/v20200515:IoTHubEventSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type ioTHubEventSourceState struct {
}

type IoTHubEventSourceState struct {
}

func (IoTHubEventSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*ioTHubEventSourceState)(nil)).Elem()
}

type ioTHubEventSourceArgs struct {
	ConsumerGroupName     string            `pulumi:"consumerGroupName"`
	EnvironmentName       string            `pulumi:"environmentName"`
	EventSourceName       *string           `pulumi:"eventSourceName"`
	EventSourceResourceId string            `pulumi:"eventSourceResourceId"`
	IotHubName            string            `pulumi:"iotHubName"`
	KeyName               string            `pulumi:"keyName"`
	Kind                  string            `pulumi:"kind"`
	LocalTimestamp        *LocalTimestamp   `pulumi:"localTimestamp"`
	Location              *string           `pulumi:"location"`
	ResourceGroupName     string            `pulumi:"resourceGroupName"`
	SharedAccessKey       string            `pulumi:"sharedAccessKey"`
	Tags                  map[string]string `pulumi:"tags"`
	Time                  *string           `pulumi:"time"`
	TimestampPropertyName *string           `pulumi:"timestampPropertyName"`
	Type                  *string           `pulumi:"type"`
}


type IoTHubEventSourceArgs struct {
	ConsumerGroupName     pulumi.StringInput
	EnvironmentName       pulumi.StringInput
	EventSourceName       pulumi.StringPtrInput
	EventSourceResourceId pulumi.StringInput
	IotHubName            pulumi.StringInput
	KeyName               pulumi.StringInput
	Kind                  pulumi.StringInput
	LocalTimestamp        LocalTimestampPtrInput
	Location              pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	SharedAccessKey       pulumi.StringInput
	Tags                  pulumi.StringMapInput
	Time                  pulumi.StringPtrInput
	TimestampPropertyName pulumi.StringPtrInput
	Type                  pulumi.StringPtrInput
}

func (IoTHubEventSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ioTHubEventSourceArgs)(nil)).Elem()
}

type IoTHubEventSourceInput interface {
	pulumi.Input

	ToIoTHubEventSourceOutput() IoTHubEventSourceOutput
	ToIoTHubEventSourceOutputWithContext(ctx context.Context) IoTHubEventSourceOutput
}

func (*IoTHubEventSource) ElementType() reflect.Type {
	return reflect.TypeOf((**IoTHubEventSource)(nil)).Elem()
}

func (i *IoTHubEventSource) ToIoTHubEventSourceOutput() IoTHubEventSourceOutput {
	return i.ToIoTHubEventSourceOutputWithContext(context.Background())
}

func (i *IoTHubEventSource) ToIoTHubEventSourceOutputWithContext(ctx context.Context) IoTHubEventSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTHubEventSourceOutput)
}

type IoTHubEventSourceOutput struct{ *pulumi.OutputState }

func (IoTHubEventSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IoTHubEventSource)(nil)).Elem()
}

func (o IoTHubEventSourceOutput) ToIoTHubEventSourceOutput() IoTHubEventSourceOutput {
	return o
}

func (o IoTHubEventSourceOutput) ToIoTHubEventSourceOutputWithContext(ctx context.Context) IoTHubEventSourceOutput {
	return o
}

func (o IoTHubEventSourceOutput) ConsumerGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringOutput { return v.ConsumerGroupName }).(pulumi.StringOutput)
}

func (o IoTHubEventSourceOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o IoTHubEventSourceOutput) EventSourceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringOutput { return v.EventSourceResourceId }).(pulumi.StringOutput)
}

func (o IoTHubEventSourceOutput) IotHubName() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringOutput { return v.IotHubName }).(pulumi.StringOutput)
}

func (o IoTHubEventSourceOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringOutput { return v.KeyName }).(pulumi.StringOutput)
}

func (o IoTHubEventSourceOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o IoTHubEventSourceOutput) LocalTimestamp() LocalTimestampResponsePtrOutput {
	return o.ApplyT(func(v *IoTHubEventSource) LocalTimestampResponsePtrOutput { return v.LocalTimestamp }).(LocalTimestampResponsePtrOutput)
}

func (o IoTHubEventSourceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o IoTHubEventSourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IoTHubEventSourceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o IoTHubEventSourceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o IoTHubEventSourceOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringPtrOutput { return v.Time }).(pulumi.StringPtrOutput)
}

func (o IoTHubEventSourceOutput) TimestampPropertyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringPtrOutput { return v.TimestampPropertyName }).(pulumi.StringPtrOutput)
}

func (o IoTHubEventSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IoTHubEventSource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IoTHubEventSourceOutput{})
}
