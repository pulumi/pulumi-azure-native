


package v20180907preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EventHubConnection struct {
	pulumi.CustomResourceState

	ConsumerGroup      pulumi.StringOutput    `pulumi:"consumerGroup"`
	DataFormat         pulumi.StringPtrOutput `pulumi:"dataFormat"`
	EventHubResourceId pulumi.StringOutput    `pulumi:"eventHubResourceId"`
	Location           pulumi.StringPtrOutput `pulumi:"location"`
	MappingRuleName    pulumi.StringPtrOutput `pulumi:"mappingRuleName"`
	Name               pulumi.StringOutput    `pulumi:"name"`
	TableName          pulumi.StringPtrOutput `pulumi:"tableName"`
	Type               pulumi.StringOutput    `pulumi:"type"`
}


func NewEventHubConnection(ctx *pulumi.Context,
	name string, args *EventHubConnectionArgs, opts ...pulumi.ResourceOption) (*EventHubConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ConsumerGroup == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerGroup'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.EventHubResourceId == nil {
		return nil, errors.New("invalid value for required argument 'EventHubResourceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:kusto/v20180907preview:EventHubConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto:EventHubConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto:EventHubConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20170907privatepreview:EventHubConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:kusto/v20170907privatepreview:EventHubConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource EventHubConnection
	err := ctx.RegisterResource("azure-native:kusto/v20180907preview:EventHubConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEventHubConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventHubConnectionState, opts ...pulumi.ResourceOption) (*EventHubConnection, error) {
	var resource EventHubConnection
	err := ctx.ReadResource("azure-native:kusto/v20180907preview:EventHubConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type eventHubConnectionState struct {
}

type EventHubConnectionState struct {
}

func (EventHubConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubConnectionState)(nil)).Elem()
}

type eventHubConnectionArgs struct {
	ClusterName            string  `pulumi:"clusterName"`
	ConsumerGroup          string  `pulumi:"consumerGroup"`
	DataFormat             *string `pulumi:"dataFormat"`
	DatabaseName           string  `pulumi:"databaseName"`
	EventHubConnectionName *string `pulumi:"eventHubConnectionName"`
	EventHubResourceId     string  `pulumi:"eventHubResourceId"`
	Location               *string `pulumi:"location"`
	MappingRuleName        *string `pulumi:"mappingRuleName"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
	TableName              *string `pulumi:"tableName"`
}


type EventHubConnectionArgs struct {
	ClusterName            pulumi.StringInput
	ConsumerGroup          pulumi.StringInput
	DataFormat             pulumi.StringPtrInput
	DatabaseName           pulumi.StringInput
	EventHubConnectionName pulumi.StringPtrInput
	EventHubResourceId     pulumi.StringInput
	Location               pulumi.StringPtrInput
	MappingRuleName        pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	TableName              pulumi.StringPtrInput
}

func (EventHubConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubConnectionArgs)(nil)).Elem()
}

type EventHubConnectionInput interface {
	pulumi.Input

	ToEventHubConnectionOutput() EventHubConnectionOutput
	ToEventHubConnectionOutputWithContext(ctx context.Context) EventHubConnectionOutput
}

func (*EventHubConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubConnection)(nil))
}

func (i *EventHubConnection) ToEventHubConnectionOutput() EventHubConnectionOutput {
	return i.ToEventHubConnectionOutputWithContext(context.Background())
}

func (i *EventHubConnection) ToEventHubConnectionOutputWithContext(ctx context.Context) EventHubConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubConnectionOutput)
}

type EventHubConnectionOutput struct{ *pulumi.OutputState }

func (EventHubConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubConnection)(nil))
}

func (o EventHubConnectionOutput) ToEventHubConnectionOutput() EventHubConnectionOutput {
	return o
}

func (o EventHubConnectionOutput) ToEventHubConnectionOutputWithContext(ctx context.Context) EventHubConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EventHubConnectionOutput{})
}
