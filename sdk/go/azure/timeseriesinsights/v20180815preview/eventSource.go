


package v20180815preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type EventSource struct {
	pulumi.CustomResourceState

	Kind     pulumi.StringOutput    `pulumi:"kind"`
	Location pulumi.StringOutput    `pulumi:"location"`
	Name     pulumi.StringOutput    `pulumi:"name"`
	Tags     pulumi.StringMapOutput `pulumi:"tags"`
	Type     pulumi.StringOutput    `pulumi:"type"`
}


func NewEventSource(ctx *pulumi.Context,
	name string, args *EventSourceArgs, opts ...pulumi.ResourceOption) (*EventSource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:timeseriesinsights:EventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20170228preview:EventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20171115:EventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20200515:EventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210331preview:EventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210630preview:EventSource"),
		},
	})
	opts = append(opts, aliases)
	var resource EventSource
	err := ctx.RegisterResource("azure-native:timeseriesinsights/v20180815preview:EventSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEventSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventSourceState, opts ...pulumi.ResourceOption) (*EventSource, error) {
	var resource EventSource
	err := ctx.ReadResource("azure-native:timeseriesinsights/v20180815preview:EventSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type eventSourceState struct {
}

type EventSourceState struct {
}

func (EventSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSourceState)(nil)).Elem()
}

type eventSourceArgs struct {
	EnvironmentName   string            `pulumi:"environmentName"`
	EventSourceName   *string           `pulumi:"eventSourceName"`
	Kind              string            `pulumi:"kind"`
	LocalTimestamp    *LocalTimestamp   `pulumi:"localTimestamp"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}


type EventSourceArgs struct {
	EnvironmentName   pulumi.StringInput
	EventSourceName   pulumi.StringPtrInput
	Kind              pulumi.StringInput
	LocalTimestamp    LocalTimestampPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (EventSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSourceArgs)(nil)).Elem()
}

type EventSourceInput interface {
	pulumi.Input

	ToEventSourceOutput() EventSourceOutput
	ToEventSourceOutputWithContext(ctx context.Context) EventSourceOutput
}

func (*EventSource) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSource)(nil)).Elem()
}

func (i *EventSource) ToEventSourceOutput() EventSourceOutput {
	return i.ToEventSourceOutputWithContext(context.Background())
}

func (i *EventSource) ToEventSourceOutputWithContext(ctx context.Context) EventSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSourceOutput)
}

type EventSourceOutput struct{ *pulumi.OutputState }

func (EventSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSource)(nil)).Elem()
}

func (o EventSourceOutput) ToEventSourceOutput() EventSourceOutput {
	return o
}

func (o EventSourceOutput) ToEventSourceOutputWithContext(ctx context.Context) EventSourceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EventSourceOutput{})
}
