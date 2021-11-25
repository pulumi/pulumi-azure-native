


package kusto

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EventGridDataConnection struct {
	pulumi.CustomResourceState

	BlobStorageEventType     pulumi.StringPtrOutput `pulumi:"blobStorageEventType"`
	ConsumerGroup            pulumi.StringOutput    `pulumi:"consumerGroup"`
	DataFormat               pulumi.StringPtrOutput `pulumi:"dataFormat"`
	EventHubResourceId       pulumi.StringOutput    `pulumi:"eventHubResourceId"`
	IgnoreFirstRecord        pulumi.BoolPtrOutput   `pulumi:"ignoreFirstRecord"`
	Kind                     pulumi.StringOutput    `pulumi:"kind"`
	Location                 pulumi.StringPtrOutput `pulumi:"location"`
	MappingRuleName          pulumi.StringPtrOutput `pulumi:"mappingRuleName"`
	Name                     pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState        pulumi.StringOutput    `pulumi:"provisioningState"`
	StorageAccountResourceId pulumi.StringOutput    `pulumi:"storageAccountResourceId"`
	TableName                pulumi.StringPtrOutput `pulumi:"tableName"`
	Type                     pulumi.StringOutput    `pulumi:"type"`
}


func NewEventGridDataConnection(ctx *pulumi.Context,
	name string, args *EventGridDataConnectionArgs, opts ...pulumi.ResourceOption) (*EventGridDataConnection, error) {
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
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageAccountResourceId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountResourceId'")
	}
	args.Kind = pulumi.String("EventGrid")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto/v20190121:EventGridDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190515:EventGridDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190907:EventGridDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20191109:EventGridDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200215:EventGridDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200614:EventGridDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200918:EventGridDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210101:EventGridDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210827:EventGridDataConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource EventGridDataConnection
	err := ctx.RegisterResource("azure-native:kusto:EventGridDataConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEventGridDataConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventGridDataConnectionState, opts ...pulumi.ResourceOption) (*EventGridDataConnection, error) {
	var resource EventGridDataConnection
	err := ctx.ReadResource("azure-native:kusto:EventGridDataConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type eventGridDataConnectionState struct {
}

type EventGridDataConnectionState struct {
}

func (EventGridDataConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventGridDataConnectionState)(nil)).Elem()
}

type eventGridDataConnectionArgs struct {
	BlobStorageEventType     *string `pulumi:"blobStorageEventType"`
	ClusterName              string  `pulumi:"clusterName"`
	ConsumerGroup            string  `pulumi:"consumerGroup"`
	DataConnectionName       *string `pulumi:"dataConnectionName"`
	DataFormat               *string `pulumi:"dataFormat"`
	DatabaseName             string  `pulumi:"databaseName"`
	EventHubResourceId       string  `pulumi:"eventHubResourceId"`
	IgnoreFirstRecord        *bool   `pulumi:"ignoreFirstRecord"`
	Kind                     string  `pulumi:"kind"`
	Location                 *string `pulumi:"location"`
	MappingRuleName          *string `pulumi:"mappingRuleName"`
	ResourceGroupName        string  `pulumi:"resourceGroupName"`
	StorageAccountResourceId string  `pulumi:"storageAccountResourceId"`
	TableName                *string `pulumi:"tableName"`
}


type EventGridDataConnectionArgs struct {
	BlobStorageEventType     pulumi.StringPtrInput
	ClusterName              pulumi.StringInput
	ConsumerGroup            pulumi.StringInput
	DataConnectionName       pulumi.StringPtrInput
	DataFormat               pulumi.StringPtrInput
	DatabaseName             pulumi.StringInput
	EventHubResourceId       pulumi.StringInput
	IgnoreFirstRecord        pulumi.BoolPtrInput
	Kind                     pulumi.StringInput
	Location                 pulumi.StringPtrInput
	MappingRuleName          pulumi.StringPtrInput
	ResourceGroupName        pulumi.StringInput
	StorageAccountResourceId pulumi.StringInput
	TableName                pulumi.StringPtrInput
}

func (EventGridDataConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventGridDataConnectionArgs)(nil)).Elem()
}

type EventGridDataConnectionInput interface {
	pulumi.Input

	ToEventGridDataConnectionOutput() EventGridDataConnectionOutput
	ToEventGridDataConnectionOutputWithContext(ctx context.Context) EventGridDataConnectionOutput
}

func (*EventGridDataConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*EventGridDataConnection)(nil))
}

func (i *EventGridDataConnection) ToEventGridDataConnectionOutput() EventGridDataConnectionOutput {
	return i.ToEventGridDataConnectionOutputWithContext(context.Background())
}

func (i *EventGridDataConnection) ToEventGridDataConnectionOutputWithContext(ctx context.Context) EventGridDataConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventGridDataConnectionOutput)
}

type EventGridDataConnectionOutput struct{ *pulumi.OutputState }

func (EventGridDataConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventGridDataConnection)(nil))
}

func (o EventGridDataConnectionOutput) ToEventGridDataConnectionOutput() EventGridDataConnectionOutput {
	return o
}

func (o EventGridDataConnectionOutput) ToEventGridDataConnectionOutputWithContext(ctx context.Context) EventGridDataConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EventGridDataConnectionOutput{})
}
