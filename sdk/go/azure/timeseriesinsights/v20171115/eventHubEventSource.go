


package v20171115

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EventHubEventSource struct {
	pulumi.CustomResourceState

	ConsumerGroupName     pulumi.StringOutput    `pulumi:"consumerGroupName"`
	CreationTime          pulumi.StringOutput    `pulumi:"creationTime"`
	EventHubName          pulumi.StringOutput    `pulumi:"eventHubName"`
	EventSourceResourceId pulumi.StringOutput    `pulumi:"eventSourceResourceId"`
	KeyName               pulumi.StringOutput    `pulumi:"keyName"`
	Kind                  pulumi.StringOutput    `pulumi:"kind"`
	Location              pulumi.StringOutput    `pulumi:"location"`
	Name                  pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState     pulumi.StringOutput    `pulumi:"provisioningState"`
	ServiceBusNamespace   pulumi.StringOutput    `pulumi:"serviceBusNamespace"`
	Tags                  pulumi.StringMapOutput `pulumi:"tags"`
	TimestampPropertyName pulumi.StringPtrOutput `pulumi:"timestampPropertyName"`
	Type                  pulumi.StringOutput    `pulumi:"type"`
}


func NewEventHubEventSource(ctx *pulumi.Context,
	name string, args *EventHubEventSourceArgs, opts ...pulumi.ResourceOption) (*EventHubEventSource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsumerGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerGroupName'")
	}
	if args.EnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentName'")
	}
	if args.EventHubName == nil {
		return nil, errors.New("invalid value for required argument 'EventHubName'")
	}
	if args.EventSourceResourceId == nil {
		return nil, errors.New("invalid value for required argument 'EventSourceResourceId'")
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
	if args.ServiceBusNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ServiceBusNamespace'")
	}
	if args.SharedAccessKey == nil {
		return nil, errors.New("invalid value for required argument 'SharedAccessKey'")
	}
	args.Kind = pulumi.String("Microsoft.EventHub")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:timeseriesinsights:EventHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20170228preview:EventHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20180815preview:EventHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20200515:EventHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210331preview:EventHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210630preview:EventHubEventSource"),
		},
	})
	opts = append(opts, aliases)
	var resource EventHubEventSource
	err := ctx.RegisterResource("azure-native:timeseriesinsights/v20171115:EventHubEventSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEventHubEventSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventHubEventSourceState, opts ...pulumi.ResourceOption) (*EventHubEventSource, error) {
	var resource EventHubEventSource
	err := ctx.ReadResource("azure-native:timeseriesinsights/v20171115:EventHubEventSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type eventHubEventSourceState struct {
}

type EventHubEventSourceState struct {
}

func (EventHubEventSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubEventSourceState)(nil)).Elem()
}

type eventHubEventSourceArgs struct {
	ConsumerGroupName     string            `pulumi:"consumerGroupName"`
	EnvironmentName       string            `pulumi:"environmentName"`
	EventHubName          string            `pulumi:"eventHubName"`
	EventSourceName       *string           `pulumi:"eventSourceName"`
	EventSourceResourceId string            `pulumi:"eventSourceResourceId"`
	KeyName               string            `pulumi:"keyName"`
	Kind                  string            `pulumi:"kind"`
	Location              *string           `pulumi:"location"`
	ResourceGroupName     string            `pulumi:"resourceGroupName"`
	ServiceBusNamespace   string            `pulumi:"serviceBusNamespace"`
	SharedAccessKey       string            `pulumi:"sharedAccessKey"`
	Tags                  map[string]string `pulumi:"tags"`
	TimestampPropertyName *string           `pulumi:"timestampPropertyName"`
}


type EventHubEventSourceArgs struct {
	ConsumerGroupName     pulumi.StringInput
	EnvironmentName       pulumi.StringInput
	EventHubName          pulumi.StringInput
	EventSourceName       pulumi.StringPtrInput
	EventSourceResourceId pulumi.StringInput
	KeyName               pulumi.StringInput
	Kind                  pulumi.StringInput
	Location              pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	ServiceBusNamespace   pulumi.StringInput
	SharedAccessKey       pulumi.StringInput
	Tags                  pulumi.StringMapInput
	TimestampPropertyName pulumi.StringPtrInput
}

func (EventHubEventSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubEventSourceArgs)(nil)).Elem()
}

type EventHubEventSourceInput interface {
	pulumi.Input

	ToEventHubEventSourceOutput() EventHubEventSourceOutput
	ToEventHubEventSourceOutputWithContext(ctx context.Context) EventHubEventSourceOutput
}

func (*EventHubEventSource) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubEventSource)(nil))
}

func (i *EventHubEventSource) ToEventHubEventSourceOutput() EventHubEventSourceOutput {
	return i.ToEventHubEventSourceOutputWithContext(context.Background())
}

func (i *EventHubEventSource) ToEventHubEventSourceOutputWithContext(ctx context.Context) EventHubEventSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubEventSourceOutput)
}

type EventHubEventSourceOutput struct{ *pulumi.OutputState }

func (EventHubEventSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubEventSource)(nil))
}

func (o EventHubEventSourceOutput) ToEventHubEventSourceOutput() EventHubEventSourceOutput {
	return o
}

func (o EventHubEventSourceOutput) ToEventHubEventSourceOutputWithContext(ctx context.Context) EventHubEventSourceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EventHubEventSourceOutput{})
}
