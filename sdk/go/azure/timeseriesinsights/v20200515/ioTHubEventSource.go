


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
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20200515:IoTHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights:IoTHubEventSource"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights:IoTHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20170228preview:IoTHubEventSource"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20170228preview:IoTHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20171115:IoTHubEventSource"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20171115:IoTHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20180815preview:IoTHubEventSource"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20180815preview:IoTHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210331preview:IoTHubEventSource"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20210331preview:IoTHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210630preview:IoTHubEventSource"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20210630preview:IoTHubEventSource"),
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
	return reflect.TypeOf((*IoTHubEventSource)(nil))
}

func (i *IoTHubEventSource) ToIoTHubEventSourceOutput() IoTHubEventSourceOutput {
	return i.ToIoTHubEventSourceOutputWithContext(context.Background())
}

func (i *IoTHubEventSource) ToIoTHubEventSourceOutputWithContext(ctx context.Context) IoTHubEventSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTHubEventSourceOutput)
}

type IoTHubEventSourceOutput struct{ *pulumi.OutputState }

func (IoTHubEventSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IoTHubEventSource)(nil))
}

func (o IoTHubEventSourceOutput) ToIoTHubEventSourceOutput() IoTHubEventSourceOutput {
	return o
}

func (o IoTHubEventSourceOutput) ToIoTHubEventSourceOutputWithContext(ctx context.Context) IoTHubEventSourceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IoTHubEventSourceOutput{})
}
